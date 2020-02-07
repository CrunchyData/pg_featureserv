---
title: "Introduction"
date:
draft: false
weight: 10
---

## Motivation

There are numerous services available which can be used to serve features.
Popular ones are [Geoserver](https://geoserver.org) and [Mapserver](https://mapserver.org)); others include [pygeoapi]().  All of these services provide the capability to read from multiple data sources
and generate feature datasets in various formats.
In exchange for that flexibility of source format, they provide less flexibility of usage.

### PostGIS-Only

By targetting PostGIS as the sole data provider, `pg_featureserv` gains significant capabilties:

* **Automatic configuration.** Just point the server at a PostgreSQL / PostGIS database, and the server discovers and automatically publishes all tables it has access to.
The Postgres system catalog provides all the metadata needed to support publishing
datasets (including things like primary key columns and table descriptions).
This has the benefit that changes to the database are published automatically without needing to restar the service.
Also, it is straightforward to take advantage of Postgres' highly-evolved clustering capabilites to provide scale-out and High Availability.
* **Full SQL power.** The server relies on the database to perform all data operations, even including converting geometry records into GeoJSON.
This provides maximum peformance, since the database is highly optimized to perform data operations such as filtering and sorting.
Also, by using [function sources]() the server can run any SQL at all to generate features.
Any data processing or feature filtering or record aggregation you can express in SQL can be exposed as feature datasets.
Moreover, function parameters are exposed as URL query parameters, which allows dynamically changing the data returned.
Using the full power of SQL means that it is easy to expose any already-developed database functionality
via the service, and minimizes the learning curve for developers.
* **Database security model.** You can restrict access to tables and functions using standard database access control. This means you can also use advanced access control techniques like row-level security to dynamically filter access based on the login role.

Moreover, by utilizing a single powerful spatial data source, the `pg_featureserv` codebase is significantly smaller and simpler.
This gives the advantages of more rapid development, fewer bugs, a more secure API, and easier deployment on a wider variety of platforms.

### Modern Web Service Architecture

`pg_featureserv` follows the modern architectural paradigm of web-friendly, RESTful microservices.
The emerging OGC API suite of specifications provides a suite simple, REST-based spatial APIs
which is aligned with standard web practice.
This philosophy is described in the W3C/OGC [Spatial Data on the Web Best Practices](https://www.w3.org/TR/sdw-bp/).
The OGC API standards are expected to see wide and rapid adoption
due to their ease of impmentation on both server and client sides.
The first of these standards to be approved is the [OGC API for Features Core](http://docs.opengeospatial.org/is/17-069r3/17-069r3.html).

A key benefit of this lighter-weight standard is the ease of extending it to expose service-specific capabilities,
and `pg_featureserv` takes full advantage of this.
For instance, `pg_featureserv` allows querying spatial functions as well as static collections, using a similar API.
Other extension provide functionality which has not yet be codified in the OGC standard,
or which exposes functionality native to PostGIS.

By focussing on the single aspect of serving spatial features, `pg_featureserv` makes it easier to deploy, provision, manage and secure feature services within a containerized environment.

### PostGIS for the Web

`pg_featureserv` is one component of **PostGIS for the Web** (aka "PostGIS FTW"), a growing family of Go spatial micro-services. Database-centric applications naturally have a central source of coordinating state, the database, which allows otherwise independent micro-services to provide HTTP-level access to the database with minimal middleware complexity.

* [pg_tileserv](https://github.com/crunchydata/pg_tileserv) provides MVT tiles for interactive clients and smooth rendering
* [pg_featureserv](https://github.com/crunchydata/pg_featureserv) provides GeoJSON feature services for reading and writing vector and attribute data from tables
* [pg_importserv]() (TBD) will provide an import API for ingesting arbitrary GIS data files

_PostGIS for the Web_ makes it possible to stand up a spatial services architecture of stateless microservices surrounding a PostgreSQL/PostGIS database cluster, in a standard container environment, on any cloud platform or internal datacenter.
