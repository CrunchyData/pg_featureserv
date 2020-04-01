---
title: "Feature Collections"
date:
draft: false
weight: 100
---

Following the OGC Features information model, the service API publishes
PostGIS tables and views as **feature collections**.

The available feature collections are listed.
Each feature collection can report metadata about its definition,
and can be queried to return data sets of features.
It is also possible to query individual features in tables which have
defined primary keys.

## Publish tables and views as feature collections

`pg_featureserv` publishes all spatial tables and views which are visible in the database.

Spatial tables and views are those which:

* include a geometry column;
* declare a geometry type; and,
* declare an SRID (spatial reference ID).

#### Example of a spatial table

Here is a simple example of defining a spatial table
which contains polygon geometries using coordinate system SRID = 4326.
(See the PostGIS documentation for more information about
[creating spatial tables](https://postgis.net/docs/manual-3.0/using_postgis_dbmanagement.html#Create_Spatial_Table)
and using [spatial reference systems](https://postgis.net/docs/manual-3.0/using_postgis_dbmanagement.html#spatial_ref_sys).)

```sql
CREATE TABLE mytable (
    geom Geometry(Polygon, 4326),
    pid text,
    address text
);
```

Tables and views are visible when they are available for access
based on the database access permissions defined for the service database user (role).
See the [Security](./security/) section for examples of setting role priviledges.

If a view uses the geometry column of an underlying table directly
the spatial column metadata is inherited for the view.***
But if a view column is defined as the result of a spatial function
then the column must be explicitly cast to a geometry type providing the type and SRID.
Also, depending on the spatial function used, it may be necessary to
explicitly set the SRID of the created geometry.

#### Example of a view definition
```sql
CREATE VIEW my_points AS
  SELECT ST_SetSRID( ST_MakePoint( lon, lat ), 4326)::geometry(Point, 4326)
  FROM my_geo_table AS t;
```

The service uses the database catalog information to provide metadata about a feature collection backed by a table or view:

* The feature collection ID is the schema-qualified name of the table or view.
* The feature collection description is provided by the comment on the table or view.
* The feature geometry is provided by the spatial column of the table or view.
* The identifier for features is provided by the primary key column for a table (if any).
* The property names and types are provided by the non-spatial columns of the table or view.
* The description for properties is provided by the comments on table/view columns.

#### Example of comments on a table
```sql
COMMENT ON TABLE mytable IS 'This is my spatial table';
COMMENT ON COLUMN mytable.geom IS 'The geometry column contains polygons in SRS 4326';
COMMENT ON COLUMN mytable.pid IS 'The Parcel Identifier is the primary key';
COMMENT ON COLUMN mytable.address IS 'The address of the Parcel';
```

## List feature collections

The path `/collections` returns a JSON document containing a list of the feature collections published by the service.

#### Example
```
http://localhost:9000/collections
```

Each listed feature collection is described by a subset of its metadata,
including name, title, description and extent.
A list of links provide URLs for accessing: ***need to specify where they can see?

* `self` - the feature collection metadata
* `alternate` - the feature collection metadata as an HTML view
* `items` - the feature collection data items


## Describe feature collection metadata

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
* The extent*** of the feature collection (if available)
* The column name providing the feature identifiers (if any)
* A list of the properties and their JSON types

A list of links provide URLs for accessing:

* `self` - the feature collection metadata
* `alternate` - the feature collection metadata as an HTML view
* `items` - the data items returned by querying the feature collection
