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
* declare an SRID (spatial reference ID)

Each feature collection can report metadata about its definition,
and can be queried to return datasets of features.
It is also possible to query individual features by **id** in tables which have
defined primary keys.

Exposed tables and views are also limited by the database access permissions
defined for the service database user.
See the [Security]({{< relref "security" >}}) section for more information.



## List Feature Collections


`/collections`

- response is JSON containing list of collections

## Query Feature Collection metadata

`/collections/{collid}`

- response is JSON containing metadata about collection

## Query Feature Collection features

`/collections/{collid}/items`

- response is GeoJSON for result dataset

### Limiting and paging results

`limit=N`

`offset=N`

### Ordering results

`orderBy=PROP`

`orderBy=PROP:A`

`orderBy=PROP:D`

### Filter by bbox

`bbox=MINX,MINY,MAXX,MAXY`

- extent is in lon/lat (4326)

### Specify properties in result

`properties=PROP1,PROP2,PROP3...`

## Query single feature

`/collections/{collid}/items/{fid}`

### Specify properties in result

`properties=PROP1,PROP2,PROP3...`
