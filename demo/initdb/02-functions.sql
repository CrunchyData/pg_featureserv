-- Lightweight demo functions that do not depend on external datasets
CREATE SCHEMA IF NOT EXISTS postgisftw;

-- us_grid with id
CREATE OR REPLACE FUNCTION postgisftw.us_grid(
    num_x integer DEFAULT 10,
    num_y integer DEFAULT 10)
RETURNS TABLE(id text, geom geometry)
AS $$
DECLARE
    lon_min CONSTANT numeric := -128;
    lon_max CONSTANT numeric := -64;
    lat_min CONSTANT numeric := 24;
    lat_max CONSTANT numeric := 49;
    dlon numeric;
    dlat numeric;
BEGIN
    dlon := (lon_max - lon_min) / num_x;
    dlat := (lat_max - lat_min) / num_y;
    RETURN QUERY
        SELECT
            x.x::text || '_' || y.y::text AS fid,
            ST_MakeEnvelope(
                lon_min + (x.x - 1) * dlon, lat_min + (y.y - 1) * dlat,
                lon_min + x.x * dlon,       lat_min + y.y * dlat, 4326
            ) AS geom
        FROM generate_series(1, num_x) AS x(x)
        CROSS JOIN generate_series(1, num_y) AS y(y);
END;
$$ LANGUAGE plpgsql STABLE STRICT;

COMMENT ON FUNCTION postgisftw.us_grid IS 'Generates a grid of rectangles covering the USA';

-- us_grid without id
CREATE OR REPLACE FUNCTION postgisftw.us_grid_noid(
    num_x integer DEFAULT 10,
    num_y integer DEFAULT 10)
RETURNS TABLE(geom geometry)
AS $$
DECLARE
    lon_min CONSTANT numeric := -128;
    lon_max CONSTANT numeric := -64;
    lat_min CONSTANT numeric := 24;
    lat_max CONSTANT numeric := 49;
    dlon numeric;
    dlat numeric;
BEGIN
    dlon := (lon_max - lon_min) / num_x;
    dlat := (lat_max - lat_min) / num_y;
    RETURN QUERY
        SELECT
            ST_MakeEnvelope(
                lon_min + (x.x - 1) * dlon, lat_min + (y.y - 1) * dlat,
                lon_min + x.x * dlon,       lat_min + y.y * dlat, 4326
            ) AS geom
        FROM generate_series(1, num_x) AS x(x)
        CROSS JOIN generate_series(1, num_y) AS y(y);
END;
$$ LANGUAGE plpgsql STABLE STRICT;

-- geo_grid over arbitrary extent
CREATE OR REPLACE FUNCTION postgisftw.geo_grid(
    num_x integer DEFAULT 10,
    num_y integer DEFAULT 10,
    lon_min numeric DEFAULT -180.0,
    lat_min numeric DEFAULT -90.0,
    lon_max numeric DEFAULT 180.0,
    lat_max numeric DEFAULT 90.0)
RETURNS TABLE(id text, geom geometry)
AS $$
DECLARE
    dlon numeric;
    dlat numeric;
BEGIN
    dlon := (lon_max - lon_min) / num_x;
    dlat := (lat_max - lat_min) / num_y;
    RETURN QUERY
        SELECT
            x.x::text || '_' || y.y::text AS id,
            ST_MakeEnvelope(
                lon_min + (x.x - 1) * dlon, lat_min + (y.y - 1) * dlat,
                lon_min + x.x * dlon,       lat_min + y.y * dlat, 4326
            ) AS geom
        FROM generate_series(1, num_x) AS x(x)
        CROSS JOIN generate_series(1, num_y) AS y(y);
END;
$$ LANGUAGE plpgsql IMMUTABLE STRICT;

COMMENT ON FUNCTION postgisftw.geo_grid IS 'Generates a grid of rectangles over a geographic extent';

-- Simple buffer function (no external deps)
CREATE OR REPLACE FUNCTION postgisftw.buffer(
    input geometry DEFAULT 'POINT(0 0)'::geometry,
    dist numeric DEFAULT 10)
RETURNS TABLE(geom geometry)
AS $$
BEGIN
    RETURN QUERY SELECT ST_Buffer(ST_SetSRID(input, 4326), dist) AS geom;
END;
$$ LANGUAGE plpgsql STABLE STRICT;

