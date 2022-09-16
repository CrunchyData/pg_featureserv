---
title: "Feature Collections"
date:
draft: false
weight: 100
---

Following the OGC Features information model, the service API publishes
PostGIS tables and views as **feature collections**.

The service API allows listing available feature collections.
Each feature collection can report metadata about its definition,
and can be queried to return data sets of features.
For tables which have a primary key defined
it is possible to query individual features by id.

## Publish tables and views as feature collections

`pg_featureserv` publishes all **spatial** tables and views (including materialized views)
which are visible in the database.

Visible spatial tables and views are those which:

* include a **geometry column**;
* declare a **geometry type**;
* declare an **SRID** (spatial reference ID);
* and the service database connection has `SELECT` privileges for
  (see the [Security](/usage/security/) section for more detail).

If the table or view has a **primary key column** it will
be used as the id for features in the collection.

Non-spatial columns are published as feature properties.
The following Postgres column data types are supported:

* `text`
* `integer`, `smallint`, `bigint`, `double precision`, `real`, `numeric`
* `boolean`
* `text[]`
* `integer[]`, `smallint[]`, `bigint[]`, `double precision[]`, `real[]`, `numeric[]`
* `boolean[]`
* `json`
* other data types *may* be supported, with output as strings

#### *Example of a spatial table*

Here is an example of defining a spatial table
which contains polygon geometries using coordinate system SRID = 4326,
a primary key column,
and two attribute columns `pid` and `address`.
(See the PostGIS documentation for more information about
[creating spatial tables](https://postgis.net/docs/manual-3.0/using_postgis_dbmanagement.html#Create_Spatial_Table)
and using [spatial reference systems](https://postgis.net/docs/manual-3.0/using_postgis_dbmanagement.html#spatial_ref_sys).)

```sql
CREATE TABLE mytable (
    id integer primary key,
    geom Geometry(Polygon, 4326),
    pid text,
    address text
);
```

#### Spatial Views
If a view directly uses the geometry column of an underlying table,
the spatial column metadata is inherited for the view.
But if a view column is defined as the result of a spatial function,
then the column must be explicitly cast to a geometry type providing the type and SRID.
Depending on the spatial function used, it may also be necessary to
explicitly set the SRID of the created geometry.

#### *Example of a spatial view definition*
```sql
CREATE VIEW my_points AS
  SELECT ST_SetSRID(
           ST_MakePoint( lon, lat ), 4326)::geometry(Point, 4326) AS geom
  FROM my_geo_table AS t;
```

#### Feature collection metadata

The service uses the database catalog information to provide metadata about a feature collection backed by a table or view:

* The **feature collection ID** is the schema-qualified name of the table or view.
* The **feature collection description** is provided by the comment on the table or view.
* The **feature geometry** is provided by the spatial column of the table or view.
* The **identifier** for features is provided by the primary key column for a table (if any).
* The **property names and types** are provided by the non-spatial columns of the table or view.
* The **description for properties** is provided by the column comment.

#### *Example of comments on a table*
```sql
COMMENT ON TABLE mytable IS 'This is my spatial table';
COMMENT ON COLUMN mytable.geom IS 'The geometry column contains polygons in SRS 4326';
COMMENT ON COLUMN mytable.pid IS 'The Parcel Identifier is the primary key';
COMMENT ON COLUMN mytable.address IS 'The address of the Parcel';
```

#### Access Control

Tables and views are visible when they are available for access
based on the database access permissions defined for the service database user (role).
See the [Security](./security/) section for examples of setting role privileges.


## List feature collections

The path `/collections` returns a JSON document containing a list of the feature collections published by the service.

#### *Example*
```
http://localhost:9000/collections
```

Each listed feature collection is described by a subset of its metadata,
including name, title, description and extent.
A list of links provide URLs for accessing:

* `self` - the feature collection metadata
* `alternate` - the feature collection metadata as an HTML view
* `items` - the feature collection data items


## Describe feature collection metadata

The path `/collections/{coll-name}` returns a JSON object describing
the metadata for a feature collection.
`{coll-name}` is the schema-qualified name of the database table or view
backing the feature collection.

#### *Example*
```
http://localhost:9000/collections/ne.admin_0_countries
```

The response is a JSON document containing metadata about the collection, including:

* The geometry column name
* The geometry type
* The geometry spatial reference code (SRID)
* The extent of the feature collection (if available)
* The column name providing the feature identifiers (if any)
* A list of the properties and their JSON types

A list of links provide URLs for accessing:

* `self` - the feature collection metadata
* `alternate` - the feature collection metadata as an HTML view
* `items` - the data items returned by querying the feature collection

## Modify collection feature

You can create, replace, patch or delete a feature of a collection.

### Create feature

To create a feature you need to provide a valid JSON document containing all the data of the feature to create.

The JSON document must match the JSON schema provided by the path `/collections/{coll-name}/schema?type=create`. Once the JSON document you can send it to the collection by using the `POST` HTTP method with the path `/collections/{coll-name}/items`.

#### *Example*

Using `curl` tool, you will have this type of request with your JSON document saved locally as `data.json` file:

```bash
curl -X POST "http://localhost:9000/collections/e.admin_0_countries/items" \
     -H "accept: */*" \
     -H "Content-Type: application/json" \
     -d "@data.json"
```

You should receive a 201 HTTP response with in the header the url of the newly created feature like:

```raw
access-control-allow-origin: *  
content-encoding: gzip  
content-length: 23  
location: http://localhost:9000/collections/e.admin_0_countries/items/10  
```
