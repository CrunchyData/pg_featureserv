---
title: "Architecture"
date:
draft: false
weight: 200
---

`pg_featureserv` has a simple architecture.
It consists of a single server application, written in Go.
It is configured via static (read-only) information sourced from a file, command-line and/or environment variables.

`pg_featureserv` can run stand-alone or inside a containerized environment.
It connects to a Postgres database using an internal database pool
(which can itself connect to a database load-balancer such as `pgbouncer`).
It contains an integrated web server which provides the HTTP interface to clients.
The interface provides both a data-centric REST API and a HTML-based user interface.

The `pg_featureserv` server integrates with the following:

* a PostGIS-enabled Postgres database instance or cluster, containing the data being served
and the catalog metadata describing it.
* client software which accesses the HTTP interface.
Typically this is a web-mapping application running in a web browser,
but it could also be a non-browser application (ranging from a simple data access utility such as `curl` or `OGR`
to a desktop GIS application such as `QGIS`), or a web proxy mediating access to the service.


The context diagram below shows `pg_featureserv` running alongside `pg_tileserv` to
provide a PostGIS-centric "platform for the spatial web".

![pg_feaureserv Architecture](/pg_fs_architecture.png)
