package data

import (
	"fmt"
	"strconv"
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
		SELECT array_agg(ARRAY[sa.attname, st.typname, coalesce(da.description,''), sa.attnum::text]::text[] ORDER BY sa.attnum)
		FROM pg_attribute sa
		JOIN pg_type st ON sa.atttypid = st.oid
		LEFT JOIN pg_description da ON (c.oid = da.objoid and sa.attnum = da.objsubid)
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
WHERE c.relkind IN ('r', 'v') -- add 'v' for views
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
		AND has_function_privilege(Format('%s.%s(%s)', n.nspname, p.proname, oidvectortypes(proargtypes)), 'execute')
),
proargarrays AS (
	SELECT p.oid,
		array_agg(p.argname ORDER BY argorder) FILTER (WHERE p.argmode = 'i') AS arginnames,
		array_agg(t.typname ORDER BY argorder) FILTER (WHERE p.argmode = 'i') AS argintypes,
		array_agg(p.argname ORDER BY argorder) FILTER (WHERE p.argmode IN ('o','t')) AS argoutnames,
		array_agg(t.typname ORDER BY argorder) FILTER (WHERE p.argmode IN ('o','t')) AS argouttypes
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
	coalesce(string_to_array(regexp_replace(pg_get_expr(p.proargdefaults, 0::Oid),
		'''([a-zA-Z0-9_\-\.]+)''::\w+', '\1', 'g'),', '), ARRAY[]::text[]) AS argdefaults,
	argoutnames AS output_names,
    argouttypes AS output_types
FROM pg_proc p
JOIN pg_namespace n ON (p.pronamespace = n.oid)
JOIN proargarrays aa ON (p.oid = aa.oid)
LEFT JOIN pg_description d ON (p.oid = d.objoid)
ORDER BY id`

const sqlFmtFeatures = "SELECT %v %v FROM %v %v %v LIMIT %v OFFSET %v;"

func sqlFeatures(tbl *Table, param *QueryParam) (string, []interface{}) {
	geomCol := sqlGeomCol(tbl.GeometryColumn, param)
	propCols := sqlColList(param.Columns, tbl.DbTypes, true)
	bboxFilter := sqlBBoxFilter(tbl, param.Bbox)
	attrFilter, attrVals := sqlAttrFilter(param.Filter)
	sqlWhere := sqlWhere(bboxFilter, attrFilter)
	sqlOrderBy := sqlOrderBy(param.OrderBy)
	sql := fmt.Sprintf(sqlFmtFeatures, geomCol, propCols, tbl.ID, sqlWhere, sqlOrderBy, param.Limit, param.Offset)
	return sql, attrVals
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

	name = strconv.Quote(name)

	// TODO: make this more data-driven / configurable
	switch dbtype {
	case forceTextTSVECTOR:
		return fmt.Sprintf("%s::text", name)
	}

	// for properties that will be treated as a string in the JSON response,
	// cast to text.  This allows displaying data types that pgx
	// does not support out of the box, as long as it can be cast to text.
	if toJSONTypeFromPG(dbtype) == JSONTypeString {
		return fmt.Sprintf("%s::text", name)
	}

	return name
}

const sqlFmtFeature = "SELECT %v %v FROM %v WHERE %v = $1 LIMIT 1"

func sqlFeature(tbl *Table, param *QueryParam) string {
	geomCol := sqlGeomCol(tbl.GeometryColumn, param)
	propCols := sqlColList(param.Columns, tbl.DbTypes, true)
	sql := fmt.Sprintf(sqlFmtFeature, geomCol, propCols, tbl.ID, tbl.IDColumn)
	return sql
}

func sqlWhere(cond1 string, cond2 string) string {
	var condList []string
	if len(cond1) > 0 {
		condList = append(condList, cond1)
	}
	if len(cond2) > 0 {
		condList = append(condList, cond2)
	}
	where := strings.Join(condList, " AND ")
	if len(where) > 0 {
		where = " WHERE " + where
	}
	return where
}

func sqlAttrFilter(filterConds []*FilterCond) (string, []interface{}) {
	var vals []interface{}
	var exprItems []string
	for i, cond := range filterConds {
		sqlCond := fmt.Sprintf("\"%v\" = $%v", cond.Name, i+1)
		exprItems = append(exprItems, sqlCond)
		vals = append(vals, cond.Value)
	}
	sql := strings.Join(exprItems, " AND ")
	return sql, vals
}

const sqlFmtBBoxFilter = ` ST_Intersects("%v", ST_Transform( ST_MakeEnvelope(%v, %v, %v, %v, 4326), %v)) `

func sqlBBoxFilter(tbl *Table, bbox *Extent) string {
	if bbox == nil {
		return ""
	}
	sql := fmt.Sprintf(sqlFmtBBoxFilter, tbl.GeometryColumn,
		bbox.Minx, bbox.Miny, bbox.Maxx, bbox.Maxy,
		tbl.Srid)
	return sql
}

const sqlFmtBBoxGeoFilter = ` ST_Intersects("%v", ST_MakeEnvelope(%v, %v, %v, %v, 4326)) `

func sqlBBoxGeoFilter(geomCol string, bbox *Extent) string {
	if bbox == nil {
		return ""
	}
	sql := fmt.Sprintf(sqlFmtBBoxGeoFilter, geomCol,
		bbox.Minx, bbox.Miny, bbox.Maxx, bbox.Maxy)
	return sql
}

const sqlFmtGeomCol = `ST_AsGeoJSON( ST_Transform( %v, 4326) %v ) AS _geojson`

func sqlGeomCol(geomCol string, param *QueryParam) string {
	geomColSafe := strconv.Quote(geomCol)
	geomExpr := applyTransform(param.TransformFuns, geomColSafe)
	sql := fmt.Sprintf(sqlFmtGeomCol, geomExpr, sqlPrecisionArg(param.Precision))
	return sql
}

func sqlPrecisionArg(precision int) string {
	if precision < 0 {
		return ""
	}
	sqlPrecision := fmt.Sprintf(",%v", precision)
	return sqlPrecision
}

const sqlFmtOrderBy = `ORDER By "%v" %v`

func sqlOrderBy(ordering []Ordering) string {
	if len(ordering) <= 0 {
		return ""
	}
	// TODO: support more than one ordering
	col := ordering[0].Name
	dir := ""
	if ordering[0].IsDesc {
		dir = "DESC"
	}
	sql := fmt.Sprintf(sqlFmtOrderBy, col, dir)
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

const sqlFmtGeomFunction = "SELECT %v %v FROM %v.%v( %v ) %v %v LIMIT %v;"

func sqlGeomFunction(fn *Function, args map[string]string, propCols []string, param *QueryParam) (string, []interface{}) {
	sqlArgs, argVals := sqlFunctionArgs(fn, args)
	sqlGeomCol := sqlGeomCol(fn.GeometryColumn, param)
	sqlPropCols := sqlColList(propCols, fn.Types, true)
	bboxFilter := sqlBBoxGeoFilter(fn.GeometryColumn, param.Bbox)
	sqlWhere := sqlWhere(bboxFilter, "")
	sqlOrderBy := sqlOrderBy(param.OrderBy)
	sql := fmt.Sprintf(sqlFmtGeomFunction, sqlGeomCol, sqlPropCols, fn.Schema, fn.Name, sqlArgs, sqlWhere, sqlOrderBy, param.Limit)
	return sql, argVals
}

const sqlFmtFunction = "SELECT %v FROM %v.%v( %v ) %v LIMIT %v;"

func sqlFunction(fn *Function, args map[string]string, propCols []string, param *QueryParam) (string, []interface{}) {
	sqlArgs, argVals := sqlFunctionArgs(fn, args)
	sqlPropCols := sqlColList(propCols, fn.Types, false)
	sqlOrderBy := sqlOrderBy(param.OrderBy)
	sql := fmt.Sprintf(sqlFmtFunction, sqlPropCols, fn.Schema, fn.Name, sqlArgs, sqlOrderBy, param.Limit)
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
