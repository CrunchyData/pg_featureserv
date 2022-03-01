---
title: "Querying Features"
date:
draft: false
weight: 150
---

Feature collections can be queried to provide sets of features,
or to return a single feature.

## Query features

The path `/collections/{collid}/items` is the basic query to return
a set of features from a feature collection.
The response is a GeoJSON feature collection containing the result.

#### Example
```
http://localhost:9000/collections/ne.countries/items
```

Additional query parameters can be appended to the basic query
to provide control over what sets of features are returned.

These are similar to using SQL statement clauses to control
the results of a query.
In fact, the service
implements these parameters by generating the equivalent SQL.
This allows the Postgres SQL engine to optimize the query execution plan.

### Filter by bounding box

The query parameter `bbox=MINX,MINY,MAXX,MAXY`
limits the features returned to those that intersect
a specified bounding box.
The bounding box is specified in geographic coordinates
(longitude/latitude, SRID = 4326).
If the source data has a non-geographic coordinate system,
the bounding box is transformed to the source coordinate system
to perform the query.

A bounding box in a different coordinate system may be specified
by adding the `bbox-crs=SRID` query parameter.

#### Example
```
http://localhost:9000/collections/ne.countries/items?bbox=10.4,43.3,26.4,47.7
```

```
http://localhost:9000/collections/ne.countries/items?bbox-crs=3005&bbox=1000000,400000,1001000,401000
```

### Filter by properties

The response feature set can be filtered to include
only features which have a given value for one or more properties.
This is done by including query parameters which have the same name as the property
to be filtered.  The value of the parameter is the desired property value.

#### Example
```
http://localhost:9000/collections/ne.countries/items?continent=Europe
```

### Filter by CQL expression

The response feature set can be filtered to include
only features which satisfy a logical expression written in
the Common Query Languae (CQL).
See the [CQL section](/query_data/cql/) for more details.

#### Example
```
http://localhost:9000/collections/ne.countries/items?filter=continent='Europe' AND pop_est<2000000
```

### Filter geometry coordinate system

By default the coordinate system of geometry literals in the filter expressionis
is assumed to be 4326 (geodetic).
A different coordinate system
can be specified by using the query parameter `filter-crs=SRID`.

#### Example
```
http://localhost:9000/collections/ebc.voting_area/items.json?filter=DWITHIN(geom,POINT(1209000+477000),1000)&filter-crs=3005
```

### Response properties

The query parameter `properties=PROP1,PROP2,PROP3...`
specifies the feature properties returned in the response.
This can reduce the response size of feature collections
which have a large number of properties.
If the parameter is specified with an empty list,
no feature properties are returned.

#### Example
```
http://localhost:9000/collections/ne.countries/items?properties=name,abbrev,pop_est
```

### Response coordinate system

The query parameter `crs=SRID`
specifies the coordinate system to be used for the
feature geometry in the response.
The SRID must be a coordinate system which is defined in the PostGIS instance.
By default data is returned in WGS84 (SRID=4326) geodetic coordinate system.

Note: GeoJSON technically does not support coordinate systems other than 4326,
but the OGC API standard allows non-geodetic data to be encoded in GeoJSON.
However, this data may not be compatible with other systems.

#### Example
```
http://localhost:9000/collections/bc.rivers/items?crs=3005
```

### Limiting and paging

The query parameter `limit=N` controls
the maximum number of features returned in a response document.
There is also a [server-defined maximum](/installation/configuration/) which cannot be exceeded.

The query parameter `offset=N` specifies the offset in the
actual query result at which the response feature set starts.

When used together, these two parameters allow paging through large result
sets.

#### Example
```
http://localhost:9000/collections/ne.countries/items?limit=50&offset=200
```

Even if the `limit` parameter is not specified, the response feature count is limited to avoid overloading the server and client.
The default number of features in a response
is set by the configuration parameter `LimitDefault`.
The maximum number of features which can be requested in the `limit` parameter
is set by the configuration parameters `LimitMax`.

### Sorting

The result set can be sorted by any property it contains.
This allows performing "greatest N" or "smallest N" queries.

* `sortby=PROP` orders results by `PROP` in ascending order

The sort order can be specified by prefixing `+` (ascending)
or `-` (descending) to the ordering property name.
The default is ascending order.

* `sortby=+PROP` orders results by `PROP` in ascending order
* `sortby=-PROP` orders results by `PROP` in descending order

**NOTE:** if used, `+` needs to be URL-encoded as `%2B`.

#### Example
```
http://localhost:9000/collections/ne.countries/items?sortby=name
```


## Query a single feature

The path `/collections/{collid}/items/{fid}`
allows querying a single feature in a feature collection
by specifying its ID.

The response is a GeoJSON feature containing the result.

#### Example
```
http://localhost:9000/collections/ne.countries/items/23
```

### Specify response properties

The query parameter `properties=PROP1,PROP2,PROP3...`
specifies the feature properties which are returned
in the response.

#### Example
```
http://localhost:9000/collections/ne.countries/items/23?properties=name,abbrev,pop_est
```

### Specify responses coordinate system

The query parameter `crs=SRID`
can be included to specify the coordinate system to be used for the
feature geometry in the response.

#### Example
```
http://localhost:9000/collections/bc.rivers/items/23?crs=3005
```
