# pg_featureserv

A lightweight RESTful feature server for [PostGIS](https://postgis.net/), written in [Go](https://golang.org/).
It supports the [OGC API - Features](http://docs.opengeospatial.org/is/17-069r3/17-069r3.html) standard.

## Features

* Implements the [OGC API for Features](http://docs.opengeospatial.org/is/17-069r3/17-069r3.html) standard.
* Emits JSON / GeoJSON
* Provides a simple HTML user interface, with maps to view geometry data
* Uses Postgres and PostGIS to manage the data catalog
* Uses PostGIS to marshall features as GeoJSON

## Building

* `go build` in application root directory

## Configuration

* Set environment variable `DATABASE_URL` with a Postgres [connection string](https://www.postgresql.org/docs/12/libpq-connect.html#LIBPQ-CONNSTRING)
  * Example: `export DATABASE_URL="host=localhost"`
* (FUTURE) Edit configuration file

## Running

* Start server: `./pg_featureserv`
* Open in a browser: `http:/localhost:9000/home.html`

## Requests Overview

Features are identified by a _collection name_ and _feature id_ pair.
Append `.html` to request path to see UI.

- API landing: http://localhost:9000/
- API definition: http://localhost:9000/api
- Conformance: http://localhost:9000/conformance
- Collections: http://localhost:9000/collections
- Feature collection metadata: http://localhost:9000/collections/{name}
- Features from a single feature collection: http://localhost:9000/collections/{name}/items
- Single feature from a feature collection: http://localhost:9000/collections/{name}/items/{featureid}

See [API Summary](API.md) for a summary of the API
