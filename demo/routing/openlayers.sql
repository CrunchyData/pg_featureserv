
--
-- UTILITY FUNCTION
--
CREATE OR REPLACE
FUNCTION public.boston_nearest_id(geom geometry)
RETURNS bigint
AS $$
    SELECT node.id
    FROM ways_vertices_pgr node
    JOIN ways edg
      ON (node.id = edg.source OR    -- Only return node that is
          node.id = edg.target)      --   an edge source or target.
    WHERE edg.source != edg.target   -- Drop circular edges.
    ORDER BY node.the_geom <-> $1    -- Find nearest node.
    LIMIT 1;
$$ LANGUAGE 'sql'
STABLE
STRICT
PARALLEL SAFE;


--
-- ROUTING FUNCTION
--
CREATE OR REPLACE
FUNCTION postgisftw.boston_find_route(
    from_lon FLOAT8 DEFAULT -71.07246980438231,
    from_lat FLOAT8 DEFAULT 42.3439930733156,
    to_lon FLOAT8 DEFAULT -71.06028184661864,
    to_lat FLOAT8 DEFAULT 42.354491297186655)
RETURNS
  TABLE(path_seq integer,
        edge bigint,
        cost double precision,
        agg_cost double precision,
        geom geometry)
AS $$
    BEGIN
    RETURN QUERY
    WITH clicks AS (
    SELECT
        ST_SetSRID(ST_Point(from_lon, from_lat), 4326) AS start,
        ST_SetSRID(ST_Point(to_lon, to_lat), 4326) AS stop
    )
    SELECT dijk.path_seq, dijk.edge, dijk.cost, dijk.agg_cost, ways.the_geom AS geom
    FROM ways
    CROSS JOIN clicks
    JOIN pgr_dijkstra(
        'SELECT gid as id, source, target, length_m as cost, length_m as reverse_cost FROM ways',
        -- source
        boston_nearest_id(clicks.start),
        -- target
        boston_nearest_id(clicks.stop)
        ) AS dijk
        ON ways.gid = dijk.edge;
    END;
$$ LANGUAGE 'plpgsql'
STABLE
STRICT
PARALLEL SAFE;
