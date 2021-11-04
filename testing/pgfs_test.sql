--===============================================
-- DDL for database objects to test pg_featureserv
--
--===============================================

CREATE SCHEMA pgfs_test;

-- DROP SCHEMA pgfs_test CASCADE;

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
