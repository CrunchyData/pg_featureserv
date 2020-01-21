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

const sqlTables = `SELECT
	Format('%s.%s', n.nspname, c.relname) AS id,
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
ORDER BY id
`
const sqlFunctions = `WITH
proargs AS (
	SELECT p.oid,
		generate_subscripts(p.proallargtypes, 1) AS argorder,
		unnest(p.proallargtypes) AS argtype,
		unnest(p.proargmodes) AS argmode,
		unnest(p.proargnames) AS argname
	FROM pg_proc p
	JOIN pg_namespace n ON (p.pronamespace = n.oid)
	WHERE n.nspname = 'postgisftw'
		AND array_length(p.proargnames, 1) = array_length(p.proargmodes, 1)
		AND array_length(p.proargmodes, 1) = array_length(p.proallargtypes, 1)
),
proargarrays AS (
	SELECT p.oid,
		array_agg(p.argname) FILTER (WHERE p.argmode = 'i') AS arginnames,
		array_agg(t.typname) FILTER (WHERE p.argmode = 'i') AS argintypes,
		array_agg(p.argname) FILTER (WHERE p.argmode IN ('o','t')) AS argoutnames,
		array_agg(t.typname) FILTER (WHERE p.argmode IN ('o','t')) AS argouttypes
	FROM proargs p
	JOIN pg_type t ON (p.argtype = t.oid)
	GROUP BY 1
)
SELECT
	p.proname AS id,
	n.nspname AS schema,
	p.proname AS function,
	coalesce(d.description, '') AS description,
	aa.arginnames AS input_names,
	aa.argintypes AS input_types,
    coalesce(string_to_array(regexp_replace(pg_get_expr(p.proargdefaults, 0::Oid), '''([a-zA-Z0-9_]+)''::text', '\1'),', '), ARRAY[]::text[]) AS input_defaults,
    argoutnames AS output_names,
    argouttypes AS output_types
FROM pg_proc p
JOIN pg_namespace n ON (p.pronamespace = n.oid)
JOIN proargarrays aa ON (p.oid = aa.oid)
LEFT JOIN pg_description d ON (p.oid = d.objoid)
ORDER BY id`

const sqlFmtFeatures = "SELECT %v %v FROM %v %v LIMIT %v OFFSET %v;"

func sqlFeatures(tbl *Table, param QueryParam) string {
	geomCol := sqlGeomCol(tbl.GeometryColumn, param)
	propCols := sqlColList(tbl.Columns, tbl.Types, true)
	sqlWhere := sqlBBoxFilter(tbl, param)
	sql := fmt.Sprintf(sqlFmtFeatures, geomCol, propCols, tbl.ID, sqlWhere, param.Limit, param.Offset)
	return sql
}

// sqlColList creates a comma-separated column list, or blank if no columns
// If addLeadingComma is true, a leading comma is added, for use when the target SQL has columns defined before
func sqlColList(names []string, dbtypes map[string]string, addLeadingComma bool) string {
	if len(names) == 0 {
		return ""
	}

	var cols []string
	for _, col := range names {
		colExpr := sqlColExpr(col, dbtypes[col])
		cols = append(cols, colExpr)
	}
	colsStr := strings.Join(cols, ",")
	if addLeadingComma {
		return ", " + colsStr
	}
	return colsStr
}

// makeSQLColExpr casts a column to text if type is unknown to PGX
func sqlColExpr(name string, dbtype string) string {
	// TODO: make this more data-driven / configurable
	switch dbtype {
	case forceTextTSVECTOR:
		return fmt.Sprintf("%s::text", name)
	}
	return name
}

const sqlFmtFeature = "SELECT %v, %v FROM %v WHERE %v = $1 LIMIT 1"

func sqlFeature(tbl *Table, param QueryParam) string {
	geomCol := sqlGeomCol(tbl.GeometryColumn, param)
	propCols := sqlColList(tbl.Columns, tbl.Types, true)
	sql := fmt.Sprintf(sqlFmtFeature, geomCol, propCols, tbl.ID, tbl.IDColumn)
	return sql
}

const sqlFmtBBoxFilter = " WHERE ST_Intersects(%v, ST_Transform( ST_MakeEnvelope(%v, %v, %v, %v, 4326), %v)) "

func sqlBBoxFilter(tbl *Table, param QueryParam) string {
	if param.Bbox == nil {
		return ""
	}
	sql := fmt.Sprintf(sqlFmtBBoxFilter, tbl.GeometryColumn,
		param.Bbox.Minx, param.Bbox.Miny, param.Bbox.Maxx, param.Bbox.Maxy,
		tbl.Srid)
	return sql
}

const sqlFmtBBoxGeoFilter = " WHERE ST_Intersects(%v, ST_MakeEnvelope(%v, %v, %v, %v, 4326)) "

func sqlBBoxGeoFilter(geomCol string, param QueryParam) string {
	if param.Bbox == nil {
		return ""
	}
	sql := fmt.Sprintf(sqlFmtBBoxGeoFilter, geomCol,
		param.Bbox.Minx, param.Bbox.Miny, param.Bbox.Maxx, param.Bbox.Maxy)
	return sql
}

const sqlFmtGeomCol = "ST_AsGeoJSON( ST_Transform(%v, 4326) %v ) AS _geojson"

func sqlGeomCol(geomCol string, param QueryParam) string {
	geomExpr := applyTransform(param.TransformFuns, geomCol)
	precision := ""
	if param.Precision >= 0 {
		precision = fmt.Sprintf(",%v", param.Precision)
	}
	sql := fmt.Sprintf(sqlFmtGeomCol, geomExpr, precision)
	return sql
}

func applyTransform(funs []TransformFunction, expr string) string {
	if funs == nil {
		return expr
	}
	for _, fun := range funs {
		expr = fun.apply(expr)
	}
	return expr
}

const sqlFmtGeomFunction = "SELECT %v %v FROM %v.%v( %v ) %v LIMIT %v;"

func sqlGeomFunction(fn *Function, args map[string]string, propCols []string, param QueryParam) (string, []interface{}) {
	sqlArgs, argVals := sqlFunctionArgs(fn, args)
	sqlGeomCol := sqlGeomCol(fn.GeometryColumn, param)
	sqlPropCols := sqlColList(propCols, fn.Types, true)
	sqlWhere := sqlBBoxGeoFilter(fn.GeometryColumn, param)
	sql := fmt.Sprintf(sqlFmtGeomFunction, sqlGeomCol, sqlPropCols, fn.Schema, fn.Name, sqlArgs, sqlWhere, param.Limit)
	return sql, argVals
}

const sqlFmtFunction = "SELECT %v FROM %v.%v( %v ) LIMIT %v;"

func sqlFunction(fn *Function, args map[string]string, propCols []string, param QueryParam) (string, []interface{}) {
	sqlArgs, argVals := sqlFunctionArgs(fn, args)
	sqlPropCols := sqlColList(propCols, fn.Types, false)
	sql := fmt.Sprintf(sqlFmtFunction, sqlPropCols, fn.Schema, fn.Name, sqlArgs, param.Limit)
	return sql, argVals
}

func sqlFunctionArgs(fn *Function, argValues map[string]string) (string, []interface{}) {
	var vals []interface{}
	var argItems []string
	i := 1
	for argName := range argValues {
		argItem := fmt.Sprintf("%v => $%v", argName, i)
		argItems = append(argItems, argItem)
		i++
		vals = append(vals, argValues[argName])
	}
	sql := strings.Join(argItems, ",")
	return sql, vals
}
