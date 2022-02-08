---
title: "Definitions"
date:
draft: false
weight: 100
---

### Feature

A representation of a real-world spatial phenomenon which can be modelled by a geometry and zero or more scalar-valued properties.

### Feature collection

A set of **features** from a spatial dataset.  In `pg_featureserv`, these are mapped to database tables and views.

### Spatial database

A database that includes a "geometry" column type. The PostGIS extension to PostgreSQL adds a geometry column type, as well as hundreds of functions to operate on that type. For example, it provides the [ST_AsGeoJSON()](https://postgis.net/docs/ST_AsGeoJSON.html) function that `pg_featureserv` uses.

### Web API

An **Application Program Interface** (API) allows client software to make programmatic requests to a service and retrieve information from it.

A Web API is an API founded on Web technologies. These include:

* Use of the HTTP protocol to provide high-level semantics for operations, as well as efficient mechanisms for querying, security and transporting data to clients
* Following the REST paradigm to simplify the model of interacting with data
* Using the standard JSON and GeoJSON formats as the primary way of encoding data

### CRS

A **Coordinate Reference System** (CRS) specifies how coordinate values in feature geometries map to locations on the earth's surface.  In PostGIS coordinate systems are identified by numeric SRID values (Spatial Reference Identifiers).
The available SRIDS are defined in the
[SPATIAL_REF_SYS table](https://postgis.net/docs/using_postgis_dbmanagement.html#spatial_ref_sys_table).
By default `pg_featureserv` provides data in the WGS84 geodetic coordinate system (SRID=4326).
Other coordinates systems can be used in `bbox` queries and for response data.
