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
http://localhost:9000/collections/ne.admin_0_countries/items
```

Additional query parameters can be appended to the basic query
to provide control over what sets of features are returned.

These are analagous to using SQL statement clauses to control
the results of a query.
In fact, the service
implements these parameters by generating the equivalent SQL.
This provides maximum performance since it allows
the Postgres SQL engine to optimize the query execution plan.

### Filter by Bounding Box

The query parameter `bbox=MINX,MINY,MAXX,MAXY`
is used to limit the features returned to those that intersect
a specified bounding box.
The bounding box is specified in geographic coordinates
(longitude/latitude, SRID = 4326).
If the source data has a non-geographic coordinate system
the bounding box is transformed to the source coordinate system
to perform the query.

#### Example
```
http://localhost:9000/collections/ne.admin_0_countries/items?bbox=10.4,43.3,26.4,47.7
```

### Restrict Properties

The query parameter `properties=PROP1,PROP2,PROP3...`
can be used to restrict the properties which are returned
in the response.
This allows reducing the response size of feature collections
which have a large number of properties.

#### Example
```
http://localhost:9000/collections/ne.admin_0_countries/items?properties=name,abbrev,pop_est
```

### Limiting and Paging

The query parameter `limit=N` can be added to the query to control
the maximum number of features returned in a response document.
There is also a server-defined maximum which cannot be exceeded.

The query parameter `offset=N` specifies the offset in the
actual query result at which the response feature set starts.

Used together these two parameters allow paging through large result
sets.

#### Example
```
http://localhost:9000/collections/ne.admin_0_countries/items?limit=50&offset=200
```

Even if the `limit` parameter is not specified
the response feature count is limited,
to avoid overloading the server and client.
The default number of features in a response
is set by the configuration parameter `LimitDefault`.
The maximum number of features which can be requested in the `limit` parameter
is set by the configuration parameters `LimitMax`.

### Ordering

The result set can be ordered by any property it contains.
This allows performing "greatest N" or "smallest N" queries.

* `orderBy=PROP` orders results by `PROP` in ascending order

The sort order can be specified by appending `:A` (ascending)
or `:D` (descending) to the ordering property name.
The default is ascending order.

* `orderBy=PROP:A` orders results by `PROP` in ascending order
* `orderBy=PROP:D` orders results by `PROP` in descending order

#### Example
```
http://localhost:9000/collections/ne.admin_0_countries/items?orderBy=name
```



## Query a Single Feature

The path `/collections/{collid}/items/{fid}`
allows querying a single feature in a feature collection
by specifying its ID.

The response is a GeoJSON feature containing the result.

#### Example
```
http://localhost:9000/collections/ne.admin_0_countries/items/23
```


### Restrict Properties

The query parameter `properties=PROP1,PROP2,PROP3...`
can be used to restrict the properties which are returned
in the response.

#### Example
```
http://localhost:9000/collections/ne.admin_0_countries/items/23?properties=name,abbrev,pop_est
```
