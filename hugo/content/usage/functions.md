---
title: "Functions"
date:
draft: false
weight: 200
---

A powerful feature of Postgres is the ability to define functions
which encapsulate complex logic and return sets of records.
Functions encapulate complex logic behind a simple
interface (that of providing some number of arguments
and getting a set of data records with some number of columns).
This makes them easy to expose via a simple web API.

Functions can perform any data processing that is
possible to perform with PostGIS, and thus
provide a powerful extension to the capabilities of
the `pg_featureserv` API.

As with feature collections, available functions can be listed,
and each function can supply metadata describing it.

## Expose Database Functions



## List Functions

The path `/functions` returns a JSON document
containing a list of the functions available in the service.

#### Example
```
http://localhost:9000/functions
```

Each listed function is described by a subset of its metadata,
including name and description.
A set of links provide URLs for accessing:

* `self` - the function metadata
* `alt` - the function metadata as an HTML view
* `items` - the function data items


## Describe Function metadata

The path `/functions/{funid}` returns a JSON object describing
the metadata for a database function.
The `{funid}` is the name of the function.

#### Example
```
http://localhost:9000/functions/geonames_geom
```

The response is a JSON document ontaining metadata about the function, including:

* The geometry column name
* The geometry type
* The geometry spatial reference code (SRID)
* The extent of the feature collection (if available)
* The column name providing the feature identifiers (if any)
* A list of the properties and their JSON types

A set of links provide URLs for accessing:

* `self` - the function metadata
* `alt` - the function metadata as an HTML view
* `items` - the data items returned by querying the function
