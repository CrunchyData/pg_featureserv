# PgRouting / pg_featureserv / pg_tileserv

This example demonstrates generating a point-to-point route in [PgRouting](https://pgrouting.org) and exposing the routing functionality to a web map via [pg_featureserv](https://github.com/CrunchyData/pg_featureserv).

## Server Setup

You will need PgRouting installed in order to to point-to-point routing. As "root", install the pg_routing package and utilities:
```
yum install osm2pgrouting_12 pgrouting_12 postgresql12-server
```
Create a database to hold the data.
```
createdb routing
```
Turn on the "postgis" and "pgrouting" extensions in the database.
```sql
CREATE EXTENSION postgis;
CREATE EXTENSION pgrouting;
```

## Load Data

For this example we will use a small snapshot of [OpenStreetMap](http://openstreetmap.org) data from Boston to route over.

<img src 'img/ways.png' />

```
wget http://download.osgeo.org/livedvd/data/osm/Boston_MA/Boston_MA.osm.bz2
bunzip2 Boston_MA.osm.bz2
```
Then use the "osm2pgrouting" loader tool to convert the raw data into a routable network in the database.
```
osm2pgrouting \
  --username postgres \
  --password postgres \
  --host 127.0.0.1 \
  --dbname routing \
  --file Boston_MA.osm
```

The loader tool creates three tables of data.

* The **pointsofinterest** table, which in this case is empty.
* The **ways_vertices_pgr** which is a table of nodes where the edges join. Note that every node has an `id` number.
* The **ways** table, which is a table of road/path and other "movement" oriented edges.

## Create Utility Function

PgRouting is a network router, to it routes from node-to-node on a graph. We will be clicking on a map to generate "from" and "to" routing points, and those points will not necessarily be on a network node. We need a utility function to convert a click point to a graph node id.

```sql
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
```

## Create Routing Function

Now that all the pieces are in place, we can expose routing functionality via `pg_featureserv` using the [function publication feature](https://access.crunchydata.com/documentation/pg_featureserv/latest/usage/functions/) that exposes table-valued functions in the "postgisftw" schema.

```sql
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
```


## Web Map Interface

The web map interface in full is available.

* [openlayers-pgrouting.html](openlayers-pgrouting.html)

Things to note:

* The routing functionality is all done by the server, by constructing a URL that calls the function in `pg_featureserv`

```js
function routeUrl(coord1, coord2) {
    var url = "http://localhost:9000/functions/boston_find_route/items.json";
    url += "?from_lon=" + coord1[0];
    url += "&from_lat=" + coord1[1];
    url += "&to_lon=" + coord2[0];
    url += "&to_lat=" + coord2[1];
    return url;
}
```

## Running it All

* Turn on `pg_featureserv` with `DATABASE_URL` environment variable pointing to the `routing` database, and confirm is is connecting by pointing your browser at the admin page.
* Turn on `pg_tileserv` with `DATABASE_URL` environment variable pointing to the `routing` database, and confirm is is connecting by pointing your browser at the admin page.
* Make sure the `vectorUrl` and `routeUrl` in the HTML file are pointing at your tile and feature servers.
* Open up the HTML page in your browers.
