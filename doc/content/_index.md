---
title: "pg_featureserv"
date:
draft: false
---

# pg_featureserv
`pg_featureserv` is a [PostGIS](https://postgis.net/)-based feature server written in [Go](https://golang.org/).
It is a lightweight, low-configuration RESTful web service that provides
access to spatial data stored in PostGIS tables, as well as spatial processing capability
based on PostGIS spatial functions.
The service API follows the [OGC API for Features Version 1.0](http://docs.opengeospatial.org/is/17-069r3/17-069r3.html) standard.
It also provides extensions
that expose more of the powerful spatial capabilities of PostGIS.

This guide walks you through how to install and use `pg_featureserv` for your spatial applications. The [Usage](./usage) section goes in-depth on how the service works. It also includes some [basic examples](./examples) of web map applications that source feature data from `pg_featureserv`.
