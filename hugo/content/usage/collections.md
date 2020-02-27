---
title: "Feature Collections"
date:
draft: false
weight: 100
---

Following
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

## Describe Feature Collection metadata

`/collections/{collid}`

- response is JSON containing metadata about collection
