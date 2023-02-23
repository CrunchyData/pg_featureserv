package data

/*
 Copyright 2022 Crunchy Data Solutions, Inc.
 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at
      http://www.apache.org/licenses/LICENSE-2.0
 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.

 Date     : October 2022
 Authors  : Benoit De Mezzo (benoit dot de dot mezzo at oslandia dot com)
        	Amaury Zarzelli (amaury dot zarzelli at ign dot fr)
			Jean-philippe Bazonnais (jean-philippe dot bazonnais at ign dot fr)
			Nicolas Revelant (nicolas dot revelant at ign dot fr)
*/

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/CrunchyData/pg_featureserv/internal/api"
	log "github.com/sirupsen/logrus"
)

const sqlTables = `SELECT
	Format('%I.%I', n.nspname, c.relname) AS id,
	n.nspname AS schema,
	c.relname AS table,
	coalesce(d.description, '') AS description,
	a.attname AS geometry_column,
	postgis_typmod_srid(a.atttypmod) AS srid,
	postgis_typmod_type(a.atttypmod) AS geometry_type,
	coalesce(ia.attname, '') AS id_column,
	(
		SELECT array_agg(ARRAY[sa.attname, st.typname, coalesce(da.description,''), sa.attnum::text, sa.attnotnull]::text[] ORDER BY sa.attnum)
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
LEFT JOIN pg_description d ON (c.oid = d.objoid AND d.objsubid = 0)
LEFT JOIN pg_index i ON (c.oid = i.indrelid AND i.indisprimary
AND i.indnatts = 1)
LEFT JOIN pg_attribute ia ON (ia.attrelid = i.indexrelid)
LEFT JOIN pg_type it ON (ia.atttypid = it.oid AND it.typname in ('int2', 'int4', 'int8'))
WHERE c.relkind IN ('r', 'v', 'm', 'p', 'f')
AND t.typname IN ('geometry', 'geography')
AND has_table_privilege(c.oid, 'select')
AND postgis_typmod_srid(a.atttypmod) > 0
ORDER BY id
`
const sqlFunctionsTemplate = `WITH
proargs AS (
	SELECT p.oid,
		generate_subscripts(p.proallargtypes, 1) AS argorder,
		unnest(p.proallargtypes) AS argtype,
		unnest(p.proargmodes) AS argmode,
		unnest(p.proargnames) AS argname
	FROM pg_proc p
	JOIN pg_namespace n ON (p.pronamespace = n.oid)
	WHERE n.nspname IN (#SCHEMAS#)
		AND array_length(p.proargnames, 1) = array_length(p.proargmodes, 1)
		AND array_length(p.proargmodes, 1) = array_length(p.proallargtypes, 1)
		AND has_schema_privilege(n.oid, 'usage')
		AND has_function_privilege(Format('%I.%I(%s)', n.nspname, p.proname, oidvectortypes(proargtypes)), 'execute')
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
	n.nspname || '.' || p.proname AS id,
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

const sqlNotifyFunction = `CREATE OR REPLACE FUNCTION %s.notify_event() RETURNS TRIGGER AS $$
DECLARE
		data json;
		notification json;
		new_xmin text;
		old_xmin text;
		id text;
BEGIN
		IF (TG_OP = 'DELETE') THEN
			data = row_to_json(OLD);
			old_xmin = OLD.xmin::text;
			new_xmin = NULL;
		ELSIF (TG_OP = 'INSERT') THEN
			data = row_to_json(NEW);
			old_xmin = NULL;
			new_xmin = NEW.xmin::text;
		ELSIF (TG_OP = 'UPDATE') THEN
			data = row_to_json(NEW);
			old_xmin = OLD.xmin::text;
			new_xmin = NEW.xmin::text;
		END IF;

		id := Format('%%I.%%I', TG_TABLE_SCHEMA, TG_TABLE_NAME);

		-- Contruct the notification as a JSON string.
		notification = json_build_object(
						'id', id,
						'schema',TG_TABLE_SCHEMA,
						'table',TG_TABLE_NAME,
						'action', TG_OP,
						'old_xmin', old_xmin,
						'new_xmin', new_xmin,
						'data', data);
		PERFORM pg_notify('table_update', notification::text);
		RETURN NULL;
END;

$$ LANGUAGE plpgsql;
`

func sqlFunctions(funSchemas []string) string {
	inSchemas := quotedList(funSchemas)
	return strings.Replace(sqlFunctionsTemplate, "#SCHEMAS#", inSchemas, 1)
}

func quotedList(names []string) string {
	itemsJoin := strings.Join(names, "','")
	return "'" + itemsJoin + "'"
}

//const sqlFmtExtentEst = `WITH ext AS (SELECT ST_Transform(ST_SetSRID(ST_EstimatedExtent('%s', '%s', '%s'), %d), 4326) AS geom)
//      SELECT ST_XMin(ext.geom) AS xmin, ST_YMin(ext.geom) AS ymin, ST_XMax(ext.geom) AS xmax, ST_YMax(ext.geom) AS ymax FROM ext;`

const sqlFmtExtentEst = `SELECT ST_XMin(ext.geom) AS xmin, ST_YMin(ext.geom) AS ymin, ST_XMax(ext.geom) AS xmax, ST_YMax(ext.geom) AS ymax
FROM ( SELECT ST_Transform(ST_SetSRID(ST_EstimatedExtent('%s', '%s', '%s'), %d), 4326) AS geom ) AS ext;`

func sqlExtentEstimated(tbl *api.Table) string {
	return fmt.Sprintf(sqlFmtExtentEst, tbl.Schema, tbl.Table, tbl.GeometryColumn, tbl.Srid)
}

const sqlFmtExtentExact = `SELECT ST_XMin(ext.geom) AS xmin, ST_YMin(ext.geom) AS ymin, ST_XMax(ext.geom) AS xmax, ST_YMax(ext.geom) AS ymax
FROM (SELECT coalesce( ST_Transform(ST_SetSRID(ST_Extent("%s"), %d), 4326),	ST_MakeEnvelope(-180, -90, 180, 90, 4326)) AS geom FROM "%s"."%s" ) AS ext;`

func sqlExtentExact(tbl *api.Table) string {
	return fmt.Sprintf(sqlFmtExtentExact, tbl.GeometryColumn, tbl.Srid, tbl.Schema, tbl.Table)
}

// xmin is used as weak eTag value
const sqlFmtFeatures = "SELECT %v, xmin AS eTag, %v FROM \"%s\".\"%s\" %v %v %v %s;"

func sqlFeatures(tbl *api.Table, param *QueryParam) (string, []interface{}) {
	geomCol := sqlGeomCol(tbl.GeometryColumn, tbl.Srid, param)

	propCols := sqlColListFromColumnMap(param.Columns, tbl.DbTypes)
	bboxFilter := sqlBBoxFilter(tbl.GeometryColumn, tbl.Srid, param.Bbox, param.BboxCrs)
	attrFilter, attrVals := sqlAttrFilter(param.Filter)
	cqlFilter := sqlCqlFilter(param.FilterSql)
	sqlWhere := sqlWhere(bboxFilter, attrFilter, cqlFilter)
	sqlGroupBy := sqlGroupBy(param.GroupBy)
	sqlOrderBy := sqlOrderBy(param.SortBy)
	sqlLimitOffset := sqlLimitOffset(param.Limit, param.Offset)
	sql := fmt.Sprintf(sqlFmtFeatures, geomCol, propCols, tbl.Schema, tbl.Table, sqlWhere, sqlGroupBy, sqlOrderBy, sqlLimitOffset)
	return sql, attrVals
}

// sqlColList creates a comma-separated column list, or blank if no columns
// If addLeadingComma is true, a leading comma is added, for use when the target SQL has columns defined before
func sqlColListFromColumnMap(names []string, dbtypes map[string]api.Column) string {
	if len(names) == 0 {
		return "null"
	}

	var cols []string
	for _, col := range names {
		colExpr := sqlColExpr(col, dbtypes[col].Type)
		cols = append(cols, colExpr)
	}
	return strings.Join(cols, ",")
}

// sqlColListFromPGTypeMap creates a comma-separated column list, or blank if no columns
func sqlColListFromStringMap(names []string, dbtypes map[string]api.PGType) string {
	if len(names) == 0 {
		return "null"
	}

	var cols []string
	for _, col := range names {
		colExpr := sqlColExpr(col, dbtypes[col])
		cols = append(cols, colExpr)
	}
	return strings.Join(cols, ",")
}

// makeSQLColExpr casts a column to text if type is unknown to PGX
func sqlColExpr(name string, dbtype api.PGType) string {

	name = strconv.Quote(name)

	// TODO: make this more data-driven / configurable
	// for properties that will be treated as a string in the JSON response,
	// cast to text.  This allows displaying data types that pgx
	// does not support out of the box, as long as it can be cast to text.
	if dbtype.ToJSONType() == api.JSONTypeString {
		return fmt.Sprintf("%s::text", name)
	}

	return name
}

// xmin is used as weak eTag value
const sqlFmtFeature = "SELECT %v, xmin AS eTag, %v FROM \"%s\".\"%s\" WHERE \"%v\" = $1 LIMIT 1"

func sqlFeature(tbl *api.Table, param *QueryParam) string {
	geomCol := sqlGeomCol(tbl.GeometryColumn, tbl.Srid, param)

	propCols := sqlColListFromColumnMap(param.Columns, tbl.DbTypes)
	sql := fmt.Sprintf(sqlFmtFeature, geomCol, propCols, tbl.Schema, tbl.Table, tbl.IDColumn)
	return sql
}

func sqlCqlFilter(sql string) string {
	//log.Debug("SQL = " + sql)
	if len(sql) == 0 {
		return ""
	}
	return "(" + sql + ")"
}

func sqlWhere(cond1 string, cond2 string, cond3 string) string {
	var condList []string
	if len(cond1) > 0 {
		condList = append(condList, cond1)
	}
	if len(cond2) > 0 {
		condList = append(condList, cond2)
	}
	if len(cond3) > 0 {
		condList = append(condList, cond3)
	}
	where := strings.Join(condList, " AND ")
	if len(where) > 0 {
		where = " WHERE " + where
	}
	return where
}

func sqlAttrFilter(filterConds []*PropertyFilter) (string, []interface{}) {
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

const sqlFmtBBoxTransformFilter = ` ST_Intersects("%v", ST_Transform( ST_MakeEnvelope(%v, %v, %v, %v, %v), %v)) `
const sqlFmtBBoxGeoFilter = ` ST_Intersects("%v", ST_MakeEnvelope(%v, %v, %v, %v, %v)) `

func sqlBBoxFilter(geomCol string, srcSRID int, bbox *api.Extent, bboxSRID int) string {
	if bbox == nil {
		return ""
	}
	if srcSRID == bboxSRID {
		return fmt.Sprintf(sqlFmtBBoxGeoFilter, geomCol,
			bbox.Minx, bbox.Miny, bbox.Maxx, bbox.Maxy, bboxSRID)
	}
	//-- transform bbox to src CRS so spatial index is used
	return fmt.Sprintf(sqlFmtBBoxTransformFilter, geomCol,
		bbox.Minx, bbox.Miny, bbox.Maxx, bbox.Maxy, bboxSRID,
		srcSRID)
}

const sqlFmtGeomCol = `ST_AsGeoJSON( %v %v ) AS _geojson`

func sqlGeomCol(geomCol string, sourceSRID int, param *QueryParam) string {
	geomColSafe := strconv.Quote(geomCol)
	geomExpr := applyTransform(param.TransformFuns, geomColSafe)
	geomOutExpr := transformToOutCrs(geomExpr, sourceSRID, param.Crs)
	sql := fmt.Sprintf(sqlFmtGeomCol, geomOutExpr, sqlPrecisionArg(param.Precision))
	return sql
}

func transformToOutCrs(geomExpr string, sourceSRID, outSRID int) string {
	if sourceSRID == outSRID {
		return geomExpr
	}
	//-- ST_Transform only accepts geometry, so cast to make sure
	return fmt.Sprintf("ST_Transform( (%v)::geometry, %v)", geomExpr, outSRID)
}

func sqlPrecisionArg(precision int) string {
	if precision < 0 {
		return ""
	}
	sqlPrecision := fmt.Sprintf(",%v", precision)
	return sqlPrecision
}

const sqlFmtOrderBy = `ORDER BY "%v" %v`

func sqlOrderBy(ordering []api.Sorting) string {
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

const sqlFmtGroupBy = `GROUP BY "%v"`

func sqlGroupBy(groupBy []string) string {
	if len(groupBy) <= 0 {
		return ""
	}
	// TODO: support more than one grouping
	col := groupBy[0]
	sql := fmt.Sprintf(sqlFmtGroupBy, col)
	log.Debugf("group by: %s", sql)
	return sql
}

func sqlLimitOffset(limit int, offset int) string {
	sqlLim := ""
	if limit >= 0 {
		sqlLim = fmt.Sprintf(" LIMIT %d", limit)
	}
	sqlOff := ""
	if offset > 0 {
		sqlOff = fmt.Sprintf(" OFFSET %d", offset)
	}
	return sqlLim + sqlOff
}

func applyTransform(funs []api.TransformFunction, expr string) string {
	if funs == nil {
		return expr
	}
	for _, fun := range funs {
		expr = fun.Apply(expr)
	}
	return expr
}

const sqlFmtGeomFunction = "SELECT %s, %s FROM \"%s\".\"%s\"( %v ) %v %v %s;"

func sqlGeomFunction(fn *api.Function, args map[string]string, propCols []string, param *QueryParam) (string, []interface{}) {
	sqlArgs, argVals := sqlFunctionArgs(fn, args)
	sqlGeomCol := sqlGeomCol(fn.GeometryColumn, SRID_UNKNOWN, param)
	sqlPropCols := sqlColListFromStringMap(propCols, fn.Types)
	//-- SRS of function output is unknown, so have to assume 4326
	bboxFilter := sqlBBoxFilter(fn.GeometryColumn, SRID_4326, param.Bbox, param.BboxCrs)
	cqlFilter := sqlCqlFilter(param.FilterSql)
	sqlWhere := sqlWhere(bboxFilter, cqlFilter, "")
	sqlOrderBy := sqlOrderBy(param.SortBy)
	sqlLimitOffset := sqlLimitOffset(param.Limit, param.Offset)
	sql := fmt.Sprintf(sqlFmtGeomFunction, sqlGeomCol, sqlPropCols, fn.Schema, fn.Name, sqlArgs, sqlWhere, sqlOrderBy, sqlLimitOffset)
	return sql, argVals
}

const sqlFmtFunction = "SELECT %v FROM \"%s\".\"%s\"( %v ) %v %v %s;"

func sqlFunction(fn *api.Function, args map[string]string, propCols []string, param *QueryParam) (string, []interface{}) {
	sqlArgs, argVals := sqlFunctionArgs(fn, args)
	sqlPropCols := sqlColListFromStringMap(propCols, fn.Types)
	cqlFilter := sqlCqlFilter(param.FilterSql)
	sqlWhere := sqlWhere(cqlFilter, "", "")
	sqlOrderBy := sqlOrderBy(param.SortBy)
	sqlLimitOffset := sqlLimitOffset(param.Limit, param.Offset)
	sql := fmt.Sprintf(sqlFmtFunction, sqlPropCols, fn.Schema, fn.Name, sqlArgs, sqlWhere, sqlOrderBy, sqlLimitOffset)
	return sql, argVals
}

func sqlFunctionArgs(fn *api.Function, argValues map[string]string) (string, []interface{}) {
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
