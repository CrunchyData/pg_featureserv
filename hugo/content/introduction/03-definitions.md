---
title: "Definitions"
date:
draft: false
weight: 150
---

### Feature
A representation of a real-world spatial phenomenon which can be modelled by a geometry and zero or more scalar-valued properties.

### Feature collection
A set of **features** from a spatial dataset.  In `pg_featureserv`, these are mapped to database tables and views.

### Spatial database
A database that includes a "geometry" column type. The PostGIS extension to PostgreSQL adds a geometry column type, as well as hundreds of functions to operate on that type.  For example, it provides the [ST_AsGeoJSON()](https://postgis.net/docs/ST_AsGeoJSON.html) function that `pg_featureserv` uses.

### Web API
An Application Program Interface (API) allows client software to make programmatic requests to a service and retrieve information from it. 

A Web API is an API which is founded on the technologies of the Web.
These include:

* use of the HTTP protocol to provide high-level semantics for operations, as well as efficient mechanisms for querying, security and transporting data to clients
* following the REST paradigm to simplify the model of interacting with data
* using the standard JSON and GeoJSON formats as the primary way of encoding data
