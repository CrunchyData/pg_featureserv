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

## Create Utility Function

PgRouting is a network router, to it routes from node-to-node on a graph. We will be clicking on a map to generate "from" and "to" routing points, and those points will not necessarily be on a network node. We need a utility function to convert a click point to a graph node id.

```sql
CREATE OR REPLACE
FUNCTION public.boston_nearest_id(geom geometry)
RETURNS bigint
AS $$
    SELECT vertices.id
    FROM (
        SELECT * FROM ways_vertices_pgr
        WHERE id IN (SELECT source FROM ways UNION SELECT target FROM ways)
    ) AS vertices
    ORDER BY vertices.the_geom <-> $1
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
    WITH clicks AS (
    SELECT
        ST_SetSRID(ST_Point(from_lon, from_lat), 4326) AS start,
        ST_SetSRID(ST_Point(to_lon, to_lat), 4326) AS stop
    )
    SELECT dijk.path_seq, dijk.edge, dijk.cost, dijk.agg_cost, ways.the_geom AS geom
    FROM ways
    CROSS JOIN clicks
    JOIN pgr_dijkstra(
        'SELECT gid as id, source, target, length as cost, length as reverse_cost FROM ways',
        -- source
        nearest_id(clicks.start),
        -- target
        nearest_id(clicks.stop)
        ) AS dijk
        ON ways.gid = dijk.edge;
$$ LANGUAGE 'sql'
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

