--===============================================
-- DDL for database objects to test pg_featureserv
--
--===============================================

CREATE SCHEMA pgfs_test;

-- DROP SCHEMA pgfs_test CASCADE;

--=====================================================================

CREATE TABLE pgfs_test.test_crs
(
    id integer primary key,
    geom geometry(polygon, 3005),
    name text
);

-- DROP TABLE pgfs_test.test_crs;
-- DELETE FROM pgfs_test.test_crs;

INSERT INTO pgfs_test.test_crs
SELECT ROW_NUMBER() OVER () AS id,
        ST_MakeEnvelope(1000000.0 + 20000 * x, 400000.0 + 20000 * y,
                        1000000.0 + 20000 * (x + 1), 400000.0 + 20000 * (y + 1),
            3005) AS geom,
        x || '_' || y AS name
  FROM generate_series(0, 9) AS x(x)
  CROSS JOIN generate_series(0, 9) AS y(y);

--=====================================================================

CREATE TABLE pgfs_test.test_srid0
(
    id integer primary key,
    geom geometry(polygon, 0),
    name text
);

-- DROP TABLE pgfs_test.test_srid0;
-- DELETE FROM pgfs_test.test_srid0;

INSERT INTO pgfs_test.test_srid0
SELECT ROW_NUMBER() OVER () AS id,
        ST_MakeEnvelope(1.0 + 2 * x, 4.0 + 2 * y,
                        1.0 + 2 * (x + 1), 4.0 + 2 * (y + 1),
            0) AS geom,
        x || '_' || y AS name
  FROM generate_series(0, 9) AS x(x)
  CROSS JOIN generate_series(0, 9) AS y(y);


--=====================================================================

CREATE TABLE pgfs_test.test_json
(
    id integer primary key,
    geom geometry(point, 4326),
    val_json json
);

-- DROP TABLE pgfs_test.test_json;

INSERT INTO pgfs_test.test_json
VALUES
  (1, 'SRID=4326;POINT(1 1)', '["a", "b", "c"]'),
  (2, 'SRID=4326;POINT(2 2)', '{"p1": 1, "p2": 2.3, "p3": [1, 2, 3]}');

--=====================================================================

-- Test a table with mixed geometry types

CREATE TABLE pgfs_test.test_geom
(
    id integer primary key,
    geom geometry(Geometry, 4326),
    data text
);

-- DROP TABLE pgfs_test.test_json;

INSERT INTO pgfs_test.test_geom
VALUES
  (1, 'SRID=4326;POINT(1 1)', 'aaa'),
  (2, 'SRID=4326;LINESTRING(1 1, 2 2)', 'bbb');

--=====================================================================
CREATE TABLE pgfs_test.test_arr
(
    id integer primary key,
    geom geometry(point, 4326),
    val_bool boolean[],
    val_int integer[],
    val_dp double precision[],
    val_txt text[]
);

-- DROP TABLE pgfs_test.test_arr;

INSERT INTO pgfs_test.test_arr
    VALUES (1, 'SRID=4326;POINT(1 1)',
            '{ true, true, false }',
            '{ 1, 2, 3 }',
            '{ 1.1, 2.2, 3.3 }',
            '{ "a", "bb", "ccc" }' );

--=====================================================================
-- Test functions

-- Function which raises:    ERROR:  Shell is not a line
CREATE OR REPLACE FUNCTION postgisftw.error_postgis()
RETURNS TABLE(geom geometry)
AS $$
BEGIN
	RETURN QUERY
		SELECT ST_Polygon('POINT(0 0)', 3857) AS geom;
END;
$$
LANGUAGE 'plpgsql' STABLE STRICT;
