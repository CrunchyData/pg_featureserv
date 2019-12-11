package data

import "fmt"

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

const sqlFeatures = `SELECT ST_AsGeoJSON( ST_Transform(%v,4326) ) AS _geojson, %v::text AS id FROM %v LIMIT %v;`

const sqlFeature = `SELECT ST_AsGeoJSON( ST_Transform(%v,4326) ) AS _geojson, %v::text AS id FROM %v WHERE %v = $1 LIMIT 1`

func applyFunctions(funs []TransformFunction, expr string) string {
	if funs == nil {
		return expr
	}
	for _, fun := range funs {
		expr = applyFun(fun, expr)
	}
	return expr
}

func applyFun(fun TransformFunction, expr string) string {
	if fun.Name == "" {
		return expr
	}
	if fun.Arg == "" {
		return fmt.Sprintf("%v( %v )", fun.Name, expr)
	}
	return fmt.Sprintf("%v( %v, %v )", fun.Name, expr, fun.Arg)
}
