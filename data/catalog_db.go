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
	"context"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"text/template"

	"github.com/CrunchyData/pg_featureserv/config"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"
)

type catalogDB struct {
	dbconn     *pgxpool.Pool
	layers     map[string]*Layer
	layersSort []*Layer
}

var instanceDB catalogDB
var templateFeature *template.Template

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
	config, err := pgxpool.ParseConfig(os.Getenv(config.AppConfig.EnvDBURL))
	db, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatal(err)
	}
	connLog := fmt.Sprintf("Connected to DB: host=%v, user=%v, database=%v",
		config.ConnConfig.Host, config.ConnConfig.User, config.ConnConfig.Database)
	log.Printf(connLog)
	return db
}

func (cat *catalogDB) Layers() ([]*Layer, error) {
	cat.refreshLayers()
	return cat.layersSort, nil
}

func (cat *catalogDB) LayerByName(name string) (*Layer, error) {
	cat.refreshLayers()
	layer, ok := cat.layers[name]
	if !ok {
		return nil, nil
	}
	return layer, nil
}

func (cat *catalogDB) LayerFeatures(name string, param QueryParam) ([]string, error) {
	layer, err := cat.LayerByName(name)
	if err != nil || layer == nil {
		return nil, err
	}
	sql := makeSQLFeatures(layer, param)
	log.Println(sql)
	features := readFeatures(cat.dbconn, layer, sql)
	return features, nil
}

func (cat *catalogDB) LayerFeature(name string, id string, param QueryParam) (string, error) {
	layer, err := cat.LayerByName(name)
	if err != nil {
		return "", err
	}

	sql := fmt.Sprintf(sqlFeature, layer.GeometryColumn, layer.IDColumn, layer.ID, layer.IDColumn)
	log.Println(sql)

	args := make([]interface{}, 0)
	args = append(args, id)
	features := readFeaturesWithArgs(cat.dbconn, layer, sql, args)

	if len(features) == 0 {
		return "", nil
	}
	return features[0], nil
}

func (cat *catalogDB) refreshLayers() {
	// TODO: refresh on timed basis?
	// for now this just loads the layers once
	if cat.layers == nil {
		cat.loadLayers()
	}
	instanceDB.loadLayers()
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
		props                                   pgtype.TextArray
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

	arrLen := props.Dimensions[0].Length
	arrStart := props.Dimensions[0].LowerBound - 1
	elmLen := props.Dimensions[1].Length

	// TODO: query columns in table-defined order

	// Since Go map order is random, list columns in array
	columns := make([]string, arrLen)
	datatypes := make(map[string]string)

	for i := arrStart; i < arrLen; i++ {
		elmPos := i * elmLen
		name := props.Elements[elmPos].String
		datatype := props.Elements[elmPos+1].String
		columns[i] = name
		datatypes[name] = datatype
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
		Columns:        columns,
		Types:          datatypes,
	}
}

func readFeatures(db *pgxpool.Pool, layer *Layer, sql string) []string {
	return readFeaturesWithArgs(db, layer, sql, nil)
}

func readFeaturesWithArgs(db *pgxpool.Pool, layer *Layer, sql string, args []interface{}) []string {
	rows, err := db.Query(context.Background(), sql, args...)
	if err != nil {
		log.Warn(err)
		return nil
	}
	var features []string
	for rows.Next() {
		feature := readFeature(rows, layer)
		//log.Println(feature)
		features = append(features, feature)
	}
	// Check for errors from iterating over rows.
	if err := rows.Err(); err != nil {
		log.Warn(err)
	}
	rows.Close()
	return features
}

func readFeature(rows pgx.Rows, layer *Layer) string {
	var id, geom string
	vals, err := rows.Values()
	if err != nil {
		log.Warn(err)
		return ""
	}
	//fmt.Println(vals)
	id = fmt.Sprintf("%v", vals[1])
	geom = fmt.Sprintf("%v", vals[0])
	//fmt.Println(geom)
	props := extractProperties(vals, layer)
	return makeFeatureJSON(id, geom, props)
}

func extractProperties(vals []interface{}, layer *Layer) map[string]interface{} {
	props := make(map[string]interface{})
	for i := range layer.Columns {
		name := layer.Columns[i]
		// offset vals index by 2 to skip geom, id
		props[name] = toJSONValue(vals[i+2])
		//fmt.Printf("%v: %v\n", name, val)
	}
	return props
}

// toJSONValue convert PG types to JSON values if needed
func toJSONValue(value interface{}) interface{} {
	switch v := value.(type) {
	case *pgtype.Numeric:
		var num float64
		// TODO: handle error
		v.AssignTo(&num)
		return num
		// TODO: handle other conversions?
	}
	// for now all other values are returned  as is
	return value
}

type featureData struct {
	Type  string                 `json:"type"`
	ID    string                 `json:"id"`
	Geom  *json.RawMessage       `json:"geometry"`
	Props map[string]interface{} `json:"properties"`
}

func makeFeatureJSON(id string, geom string, props map[string]interface{}) string {
	geomRaw := json.RawMessage(geom)
	featData := featureData{"Feature", id, &geomRaw, props}
	json, err := json.Marshal(featData)
	if err != nil {
		log.Error(err)
		return ""
	}
	jsonStr := string(json)
	//fmt.Println(jsonStr)
	return jsonStr
}
