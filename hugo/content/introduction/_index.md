---
title: "About pg_featureserv"
date:
draft: false
weight: 10
---

## Motivation

There are numerous services available that can be used to serve features, such as [Geoserver](https://geoserver.org), [Mapserver](https://mapserver.org), and [pygeoapi](https://pygeoapi.io/). These applications typically provide the capability to read from multiple data sources
and generate feature datasets in various formats.
They also tend to be large, complex applications which require significant expertise
to install, configure, secure and tune.

## PostGIS-Only

In contrast, `pg_featureserv` works exclusively with PostGIS, which allows for greater flexibility of usage.
By targetting PostGIS as the sole data provider, `pg_featureserv` gains significant capabilties:

* **Automatic configuration.** Just point the server at a PostgreSQL / PostGIS database, and the server discovers and automatically publishes all tables it has access to.

    The Postgres system catalog provides all the metadata needed to support publishing datasets (such as primary key columns and table descriptions). Changes to the database are then published automatically without needing to restart the service. You can also take advantage of Postgres' clustering capabilites to provide scale-out and high availability.

* **Full SQL power.** The server relies on the database to conduct all data operations, including converting geometry records into GeoJSON. Since the database is optimized to perform operations such as filtering and sorting, this increases your application's performance.

    By using functions as data sources, the server can run any SQL at all to generate features. Any data processing, feature filtering, or record aggregation that you can express in SQL can be published as feature datasets. Function parameters are also exposed as URL query parameters, which allows dynamically changing the data returned.

    Using the full power of SQL means that it is easy to publish any existing database functionality via the service, and the learning curve for developers can be minimized.

* **Database security model.** You can restrict access to tables and functions using standard database access control. This means you can also use advanced access control techniques like row-level security to dynamically filter access based on the login role.

By using a single powerful spatial data source, the `pg_featureserv` codebase is significantly smaller and simpler.
This means more rapid development, fewer software defects, a more secure interface, and easier deployment on a wider variety of platforms.

## Modern web service architecture

`pg_featureserv` follows the modern architectural paradigm of web-friendly, RESTful microservices.

As noted in the W3C/OGC [Spatial Data on the Web Best Practices](https://www.w3.org/TR/sdw-bp/), exposing spatial data using modern web standards improves spatial data discoverability, accessibility and interoperability.

A key benefit of following the lightweight [OGC API for Features Core](http://docs.opengeospatial.org/is/17-069r3/17-069r3.html) standard is the ease of extending it to expose service-specific capabilities, including the powerful spatial capabilities of PostGIS. For instance, with `pg_featureserv` you can query spatial functions as well as static collections, using a similar API.

By focussing on the single aspect of serving spatial features, `pg_featureserv` makes it easier to deploy, provision, manage, and secure feature services within a containerized environment.

## PostGIS for the Web

`pg_featureserv` is one component of _PostGIS for the Web_ (aka "PostGIS FTW"), a growing family of spatial micro-services. Database-centric applications naturally have a central source of coordinating state, the database, which allows otherwise independent microservices to provide HTTP-level access to the data with minimal middleware complexity.

* [pg_tileserv](https://access.crunchydata.com/documentation/pg_tileserv/latest/) provides MVT tiles for interactive clients and smooth rendering
* [pg_featureserv](/) provides GeoJSON feature services for reading and writing vector and attribute data from tables

_PostGIS for the Web_ makes it possible to stand up a spatial services architecture of stateless microservices surrounding a PostgreSQL/PostGIS database cluster, in a standard container environment, on any cloud platform or internal datacenter.
