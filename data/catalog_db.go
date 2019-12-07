package data

/*
 Copyright 2019 Crunchy Data Solutions, Inc.
 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at
      http://www.apache.org/licenses/LICENSE-2.0
 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"sort"
	"strconv"
	"text/template"

	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type catalogDB struct {
	dbconn     *pgxpool.Pool
	layers     map[string]*Layer
	layersSort []*Layer
}

var instanceDB catalogDB

// CatDBInstance tbd
func CatDBInstance() Catalog {
	// TODO: make a singleton
	instanceDB = newCatalogDB()
	return &instanceDB
}

func newCatalogDB() catalogDB {
	conn := dbConnect()
	cat := catalogDB{
		dbconn: conn,
	}
	return cat
}

func dbConnect() *pgxpool.Pool {
	var err error
	db, err := pgxpool.Connect(context.Background(), "host=localhost")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func (cat *catalogDB) Layers() ([]*Layer, error) {
	instanceDB.loadLayers()
	return cat.layersSort, nil
}

func (cat *catalogDB) LayerByName(name string) (*Layer, error) {
	layer, ok := cat.layers[name]
	if !ok {
		return nil, fmt.Errorf(errMsgBadLayerName, name)
	}
	return layer, nil
}

func (cat *catalogDB) LayerFeatures(name string) ([]string, error) {
	layer, err := cat.LayerByName(name)
	if err != nil {
		return nil, err
	}

	sql := fmt.Sprintf(sqlFeatures, layer.GeometryColumn, layer.IDColumn, layer.ID)
	log.Println(sql)

	features := readFeatures(cat.dbconn, layer, sql)
	return features, nil
}

func (cat *catalogDB) LayerFeaturesOLD(name string) ([]string, error) {
	_, err := cat.LayerByName(name)
	if err != nil {
		return []string{}, err
	}
	return featuresMock, nil
}

func (cat *catalogDB) LayerFeature(name string, id string) (string, error) {
	index, err := strconv.Atoi(id)
	if err != nil {
		return "", fmt.Errorf(errMsgNoBadFeatureID, id)
	}

	//fmt.Println("LayerFeatures: " + name)
	//fmt.Println(layerData)
	return featuresMock[index], nil
}

func (cat *catalogDB) loadLayers() {
	cat.layers = readLayerTables(cat.dbconn)
	cat.layersSort = layersSorted(cat.layers)
}

func layersSorted(layers map[string]*Layer) []*Layer {
	var lsort []*Layer
	for key := range layers {
		lsort = append(lsort, layers[key])
	}
	sort.SliceStable(lsort, func(i, j int) bool {
		return lsort[i].Title < lsort[j].Title
	})
	return lsort
}

func readLayerTables(db *pgxpool.Pool) map[string]*Layer {
	rows, err := db.Query(context.Background(), sqlLayers)
	if err != nil {
		log.Fatal(err)
	}
	layers := make(map[string]*Layer)
	for rows.Next() {
		layer := readLayer(rows)
		layers[layer.ID] = layer
	}
	// Check for errors from iterating over rows.
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	rows.Close()
	return layers
}

func readLayer(rows pgx.Rows) *Layer {
	var (
		schema, table, description, geometryCol string
		srid                                    int
		geometryType, idColumn                  string
		//props                                   [][]string
		props pgtype.TextArray
	)

	err := rows.Scan(&schema, &table, &description, &geometryCol,
		&srid, &geometryType, &idColumn, &props)
	if err != nil {
		log.Fatal(err)
	}

	// Use https://godoc.org/github.com/jackc/pgtype#TextArray
	// here to scan the text[][] map of attribute name/type
	// created in the query. It gets a little ugly demapping the
	// pgx TextArray type, but it is at least native handling of
	// the array. It's complex because of PgSQL ARRAY generality
	// really, no fault of pgx
	properties := make(map[string]string)

	arrLen := props.Dimensions[0].Length
	arrStart := props.Dimensions[0].LowerBound - 1
	elmLen := props.Dimensions[1].Length
	for i := arrStart; i < arrLen; i++ {
		elmPos := i * elmLen
		properties[props.Elements[elmPos].String] = props.Elements[elmPos+1].String
	}

	// "schema.tablename" is the unique key for table layers
	id := fmt.Sprintf("%s.%s", schema, table)

	// Synthesize a title for now
	title := id
	// synthesize a description if none provided
	if description == "" {
		description = fmt.Sprintf("Data for table %v", id)
	}

	return &Layer{
		ID:             id,
		Schema:         schema,
		Table:          table,
		Title:          title,
		Description:    description,
		GeometryColumn: geometryCol,
		Srid:           srid,
		GeometryType:   geometryType,
		IDColumn:       idColumn,
		//Properties:     properties,
	}
}

func readFeatures(db *pgxpool.Pool, layer *Layer, sqlFeatures string) []string {
	rows, err := db.Query(context.Background(), sqlFeatures)
	if err != nil {
		log.Fatal(err)
	}
	var features []string
	for rows.Next() {
		feature := readFeature(rows)
		log.Println(feature)
		features = append(features, feature)
	}
	// Check for errors from iterating over rows.
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	rows.Close()
	return features
}

func readFeature(rows pgx.Rows) string {
	var id, geom string
	err := rows.Scan(&geom, &id)
	if err != nil {
		log.Fatal(err)
	}
	return makeFeature(id, geom)
}

type featureData struct {
	ID   string
	Geom string
	Val  string
}

var templateFeature = `{ "type": "Feature", "id": {{ .ID }},
"geometry": {{ .Geom }},
"properties": { "value": "{{ .Val }}"  } }
`

func makeFeature(id string, geom string) string {
	tmpl, err := template.New("feature").Parse(templateFeature)
	if err != nil {
		panic(err)
	}
	val := fmt.Sprintf("data value %v", id)
	featData := featureData{id, geom, val}
	var tempOut bytes.Buffer
	tmpl.Execute(&tempOut, featData)
	feature := tempOut.String()
	//fmt.Println(features[index])
	return feature
}
