-- Given coordinates in the hexagon tiling that has this
-- edge size, return the built-out hexagon
CREATE OR REPLACE
FUNCTION hexagon(i integer, j integer, edge float8)
RETURNS geometry
AS $$
DECLARE
h float8 := edge*cos(pi()/6.0);
cx float8 := 1.5*i*edge;
cy float8 := h*(2*j+abs(i%2));
BEGIN
RETURN ST_MakePolygon(ST_MakeLine(ARRAY[
            ST_MakePoint(cx - 1.0*edge, cy + 0),
            ST_MakePoint(cx - 0.5*edge, cy + -1*h),
            ST_MakePoint(cx + 0.5*edge, cy + -1*h),
            ST_MakePoint(cx + 1.0*edge, cy + 0),
            ST_MakePoint(cx + 0.5*edge, cy + h),
            ST_MakePoint(cx - 0.5*edge, cy + h),
            ST_MakePoint(cx - 1.0*edge, cy + 0)
        ]));
END;
$$
LANGUAGE 'plpgsql'
IMMUTABLE
STRICT
PARALLEL SAFE;

SELECT ST_AsText(hexagon(2, 2, 10.0));

-- Given a square bounds, find all the hexagonal cells
-- of a hex tiling (determined by edge size)
-- that might cover that square (slightly over-determined)
CREATE OR REPLACE
FUNCTION hexagoncoordinates(bounds geometry, edge float8,
                            OUT i integer, OUT j integer)
RETURNS SETOF record
AS $$
    DECLARE
        h float8 := edge*cos(pi()/6);
        mini integer := floor(st_xmin(bounds) / (1.5*edge));
        minj integer := floor(st_ymin(bounds) / (2*h));
        maxi integer := ceil(st_xmax(bounds) / (1.5*edge));
        maxj integer := ceil(st_ymax(bounds) / (2*h));
    BEGIN
    FOR i, j IN
    SELECT a, b
    FROM generate_series(mini, maxi) a,
         generate_series(minj, maxj) b
    LOOP
        RETURN NEXT;
    END LOOP;
    END;
$$
LANGUAGE 'plpgsql'
IMMUTABLE
STRICT
PARALLEL SAFE;

SELECT * FROM hexagoncoordinates(ST_TileEnvelope(15, 1, 1), 1000.0);


CREATE OR REPLACE
FUNCTION postgisftw.hex_grid(
  edge_len numeric DEFAULT 1,
  lon_min numeric DEFAULT -180.0,
  lat_min numeric DEFAULT -80.0,
  lon_max numeric DEFAULT 180.0,
  lat_max numeric DEFAULT 80.0
)
RETURNS TABLE(geom geometry(Polygon, 3857), i integer, j integer)
AS $$
DECLARE
    boxwm geometry;
BEGIN
  boxwm:= ST_Transform( ST_MakeEnvelope( lon_min, lat_min, lon_max, lat_max, 4326 ), 3857);
  RETURN QUERY
  SELECT ST_SetSRID(hexagon(h.i, h.j, edge_len), 3857), h.i, h.j
  FROM (SELECT * FROM hexagoncoordinates( boxwm, edge_len )) AS h;
END;
$$
LANGUAGE 'plpgsql' IMMUTABLE STRICT PARALLEL SAFE;

SELECT * FROM hex_grid(1.0, 0,0, 10, 10);
