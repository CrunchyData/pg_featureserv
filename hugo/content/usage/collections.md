---
title: "Feature Collections"
date:
draft: false
weight: 100
---

Following the OCG Features information model, the service API exposes
PostGIS tables and views as **feature collections**.

The available feature collections can be listed.
Each feature collection can report metadata about its definition,
and can be queried to return datasets of features.
It is also possible to query individual features in tables which have
defined primary keys.

## Expose Tables and Views as Feature Collections

`pg_featureserv` exposes all spatial tables and views which are visible in the database.

Spatial tables and views are those which:

* include a geometry column;
* declare a geometry type; and,
* declare an SRID (spatial reference ID)

Visible tables and views are the available for access by virtue of by the database access permissions defined for the service database user.
See the [Security]({{< relref "security" >}}) section for more information.

The service relies on the database catalog information to provide metadata about a table or view.
The metadata includes:

* The feature collection id is the schema-qualified name of the table or view
* The feature collection description is provided by the comment on the table or view
* The feature geometry is provided by the spatial column of the table or view
* The identifier for features is provided by the primary key column for a table (if any)
* The property names and types are provided by the non-spatial columsn of the table or view
* The description for properties is provided by the comments on table/view columns

## List Feature Collections

The path `/collections` returns a JSON document
containing a list of the feature collections
available in the service.

#### Example
```
http://localhost:9000/collections
```

Each listed feature collection is described by a subset of its metadata,
including name, title, description and extent.
A set of links provide URLs for accessing:

* `self` - the feature collection metadata
* `alt` - the feature collection metadata as an HTML view
* `items` - the feature collection data items


## Describe Feature Collection metadata

The path `/collections/{coll-name}` returns a JSON object describing
the metadata for a feature collection.
`{coll-name}` is the schema-qualified name of the database table or view
backing the feature collection.

#### Example
```
http://localhost:9000/collections/ne.admin_0_countries
```

The response is a JSON document ontaining metadata about the collection, including:

* The geometry column name
* The geometry type
* The geometry spatial reference code (SRID)
* The extent of the feature collection (if available)
* The column name providing the feature identifiers (if any)
* A list of the properties and their JSON types

A set of links provide URLs for accessing:

* `self` - the feature collection metadata
* `alt` - the feature collection metadata as an HTML view
* `items` - the the data items returned by querying the feature collection
