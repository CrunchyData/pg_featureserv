package data

import (
	"fmt"
	"strings"
)

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

const forceTextTSVECTOR = "tsvector"

const sqlLayers = `SELECT
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
WHERE c.relkind IN ('r') -- add 'v' for views
AND t.typname = 'geometry'
AND has_table_privilege(c.oid, 'select')
AND postgis_typmod_srid(a.atttypmod) > 0
`

const sqlFmtFeatures = "SELECT %v, %v::text AS id, %v FROM %v %v LIMIT %v;"

func makeSQLFeatures(layer *Layer, param QueryParam) string {
	geomCol := makeGeomCol(layer, param)
	propCols := makeSQLColList(layer)
	sqlWhere := makeBBoxFilter(layer, param)
	sql := fmt.Sprintf(sqlFmtFeatures, geomCol, layer.IDColumn, propCols, layer.ID, sqlWhere, param.Limit)
	return sql
}

func makeSQLColList(layer *Layer) string {
	var cols []string
	for _, col := range layer.Columns {
		colExpr := makeSQLColExpr(col, layer.Types[col])
		cols = append(cols, colExpr)
	}
	return strings.Join(cols, ",")
}

// makeSQLColExpr casts a column to text if type is unknown to PGX
func makeSQLColExpr(name string, dbtype string) string {
	// TODO: make this more data-driven / configurable
	switch dbtype {
	case forceTextTSVECTOR:
		return fmt.Sprintf("%s::text", name)
	}
	return name
}

const sqlFeature = "SELECT %v, %v::text AS id, %v FROM %v WHERE %v = $1 LIMIT 1"

func makeSQLFeature(layer *Layer, param QueryParam) string {
	geomCol := makeGeomCol(layer, param)
	propCols := makeSQLColList(layer)
	sql := fmt.Sprintf(sqlFeature, geomCol, layer.IDColumn, propCols, layer.ID, layer.IDColumn)
	return sql
}

const sqlFmtBBoxFilter = " WHERE ST_Intersects(%v, ST_Transform( ST_MakeEnvelope(%v, %v, %v, %v, 4326), %v)) "

func makeBBoxFilter(layer *Layer, param QueryParam) string {
	if param.Bbox == nil {
		return ""
	}
	sql := fmt.Sprintf(sqlFmtBBoxFilter, layer.GeometryColumn,
		param.Bbox.Minx, param.Bbox.Miny, param.Bbox.Maxx, param.Bbox.Maxy,
		layer.Srid)
	return sql
}

const sqlGeomCol = "ST_AsGeoJSON( ST_Transform(%v, 4326) %v ) AS _geojson"

func makeGeomCol(layer *Layer, param QueryParam) string {
	geomExpr := applyFunctions(param.TransformFuns, layer.GeometryColumn)
	precision := ""
	if param.Precision >= 0 {
		precision = fmt.Sprintf(",%v", param.Precision)
	}
	sql := fmt.Sprintf(sqlGeomCol, geomExpr, precision)
	return sql
}

func applyFunctions(funs []TransformFunction, expr string) string {
	if funs == nil {
		return expr
	}
	for _, fun := range funs {
		expr = fun.apply(expr)
	}
	return expr
}
