---
title: "Tables and Views"
date:
draft: false
weight: 100
---

## Expose Tables and Views as Feature Collections

`pg_featureserv` exposes database spatial tables and views in the API as **feature collections**.

Spatial tables and views are those which:

* include a geometry column;
* declare a geometry type; and,
* declare an SRID (spatial reference ID).

Each feature collection can report metadata about its definition,
and can be queried to return datasets of features.
It is also possible to query individual features by **id** in tables which have
defined primary keys.

Exposed tables and views are also limited by the database access permissions
defined for the service database user.
See the [Security](/usage/security/) section for more information.

## List Feature Collections

`/collections`

- Response is JSON containing list of collections.

## Query Feature Collection Metadata

`/collections/{collid}`

- Response is JSON containing metadata about collection.

## Query Feature Collection Features

`/collections/{collid}/items`

- Response is GeoJSON for result dataset.

### Limiting and paging results

- `limit=N` specifies the maximum number of records returned. The default is 10.

- `offset=N` specifies the number of rows to skip before beginning to return rows.

### Specify properties in result

`properties=PROP1,PROP2,PROP3...`

### Ordering results

`orderBy=PROP`

- `orderBy=PROP:A` returns results based on the specified property in ascending order.

- `orderBy=PROP:D` returns results based on the specified property in descending order.

### Filter by bbox

`bbox=MINX,MINY,MAXX,MAXY`

- Extent is in lon/lat (4326)

## Query a Single Feature

`/collections/{collid}/items/{fid}`

### Specify properties in result

`properties=PROP1,PROP2,PROP3...`
