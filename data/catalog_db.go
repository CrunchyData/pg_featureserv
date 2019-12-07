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
	"fmt"
	"log"
	"sort"
	"strconv"

	"github.com/jackc/pgtype"
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
	for _, lyr := range cat.layers {
		if lyr.ID == name {
			return lyr, nil
		}
	}
	// not found
	return nil, fmt.Errorf(errMsgBadLayerName, name)
}

func (cat *catalogDB) LayerFeatures(name string) ([]string, error) {
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
	cat.layers = loadLayerTables(cat.dbconn)
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

func loadLayerTables(db *pgxpool.Pool) map[string]*Layer {

	sqlLayer := `
		SELECT
			n.nspname AS schema,
			c.relname AS table,
			coalesce(d.description, '') AS description,
			a.attname AS geometry_column,
			postgis_typmod_srid(a.atttypmod) AS srid,
			postgis_typmod_type(a.atttypmod) AS geometry_type,
			coalesce(ia.attname, '') AS id_column,
			(
				SELECT array_agg(ARRAY[sa.attname, st.typname]::text[])
				FROM pg_attribute sa
				JOIN pg_type st ON sa.atttypid = st.oid
				WHERE sa.attrelid = c.oid
				AND sa.attnum > 0
				AND NOT sa.attisdropped
				AND st.typname NOT IN ('geometry', 'geography')
			) AS props
		FROM pg_class c
		JOIN pg_namespace n ON (c.relnamespace = n.oid)
		JOIN pg_attribute a ON (a.attrelid = c.oid)
		JOIN pg_type t ON (a.atttypid = t.oid)
		LEFT JOIN pg_description d ON (c.oid = d.objoid)
		LEFT JOIN pg_index i ON (c.oid = i.indrelid AND i.indisprimary
		AND i.indnatts = 1)
		LEFT JOIN pg_attribute ia ON (ia.attrelid = i.indexrelid)
		LEFT JOIN pg_type it ON (ia.atttypid = it.oid AND it.typname in ('int2', 'int4', 'int8'))
		WHERE c.relkind = 'r'
		AND t.typname = 'geometry'
		AND has_table_privilege(c.oid, 'select')
		AND postgis_typmod_srid(a.atttypmod) > 0
		`

	rows, err := db.Query(context.Background(), sqlLayer)
	if err != nil {
		log.Fatal(err)
	}

	// Reset array of layers
	layers := make(map[string]*Layer)
	for rows.Next() {

		var (
			schema, table, description, geometry_column string
			srid                                        int
			geometry_type, id_column                    string
			// props                                       [][]string
			props pgtype.TextArray
		)

		err := rows.Scan(&schema, &table, &description, &geometry_column,
			&srid, &geometry_type, &id_column, &props)
		if err != nil {
			log.Fatal(err)
		}

		// We use https://godoc.org/github.com/jackc/pgtype#TextArray
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

		// "schema.tablename" is our unique key for table layers
		id := fmt.Sprintf("%s.%s", schema, table)

		// synthesize a description if none provided
		if description == "" {
			description = fmt.Sprintf("Data for table %v", id)
		}
		title := id

		lyr := Layer{
			ID:             id,
			Schema:         schema,
			Table:          table,
			Title:          title,
			Description:    description,
			GeometryColumn: geometry_column,
			Srid:           srid,
			GeometryType:   geometry_type,
			IDColumn:       id_column,
			//Properties:     properties,
		}

		layers[id] = &lyr
	}
	// Check for errors from iterating over rows.
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	rows.Close()
	return layers
}
