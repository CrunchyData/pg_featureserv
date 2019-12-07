package data

const sqlLayers = `
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

const sqlFeatures = `
SELECT ST_AsGeoJSON( ST_Transform(%v,4326) ) AS geom, %v::text AS id from %v LIMIT 100;
`
