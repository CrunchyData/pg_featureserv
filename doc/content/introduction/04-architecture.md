---
title: "Architecture"
date:
draft: false
weight: 175
---

`pg_featureserv` has a simple architecture.  It consists of a single Go application, together with static (read-only) configuration sourced from
a file, the command-line and/or environment variables.

The service application integrates with the following:

* a PostGIS-enabled Postgres database instance or cluster, containing the data being served
and the catalog metadata describing it.
* client software which utilize an HTTP API.  Typically this is a web-mapping application running in a web browser,
but it could also be another application (ranging from a simple data access utility such as `curl` or `OGR`
to a desktop GIS application such as `QGIS`), or a web proxy mediating access to the service.

The `pg_featureserv` component connects to the database using a database pool.
It contains an integrated web server which provides an HTTP interface to clients.

The `pg_featureserv` application can run stand-alone or inside a containerized environment.

The context diagram below shows `pg_featureserv` running alongside `pg_tileserv` to
provide a PostGIS-centric "platform for the spatial web".

![pg_feaureserv Architecture](/pg_fs_architecture.png)
