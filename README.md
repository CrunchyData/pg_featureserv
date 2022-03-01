<p align="center">
  <a href="https://access.crunchydata.com/documentation/pg_tileserv/latest/"><img width="180" height="180" src="./hugo/static/crunchy-spatial-logo.png?raw=true" /></a>
</p>

# pg_featureserv

[![.github/workflows/ci.yml](https://github.com/CrunchyData/pg_featureserv/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/CrunchyData/pg_featureserv/actions/workflows/ci.yml)

A lightweight RESTful geospatial feature server for [PostGIS](https://postgis.net/), written in [Go](https://golang.org/).
It supports the [*OGC API - Features*](https://ogcapi.ogc.org/features/) REST API standard.

## Features

* Implements the [*OGC API - Features*](https://ogcapi.ogc.org/features/) standard.
  * Supports standard request parameters `limit`, `bbox`, property filtering, `sortby`, `crs`, `bbox-crs`
  * Provides [CQL filtering](https://portal.ogc.org/files/96288), including spatial support, via parameters `filter` and `filter-crs`
  * Extended parameters include `offset`, `properties`, `transform`, `precision`, `groupby`
* Data responses are formatted in JSON and GeoJSON
* Provides a simple HTML user interface, with web maps to view spatial data
* Uses the power of PostgreSQL to reduce the amount of code
  and to make data definition easy and familiar.
  * Feature collections are defined by database objects (tables and views)
  * Filters are executed in the database, and use indexes where defined
* Uses PostGIS to provide geospatial functionality:
  * Spatial filtering  
  * Transforming geometry data into the output coordinate system
  * Marshalling feature data into GeoJSON
* Full-featured HTTP support
  * CORS support with configurable Allowed Origins
  * GZIP response encoding
  * HTTP and HTTPS support

For a full list of software capabilities see [FEATURES](FEATURES.md).

## Documentation

* [User Guide](https://access.crunchydata.com/documentation/pg_featureserv/latest/)
* [FEATURES](FEATURES.md) - full list of software capabilities
* [API](API.md) - summary of the web service API


## Download

Builds of the latest code:

* [Linux](https://postgisftw.s3.amazonaws.com/pg_featureserv_latest_linux.zip)
* [Windows](https://postgisftw.s3.amazonaws.com/pg_featureserv_latest_windows.zip)
* [MacOS](https://postgisftw.s3.amazonaws.com/pg_featureserv_latest_macos.zip)
* [Docker](https://hub.docker.com/r/pramsey/pg_featureserv)


## Building from Source

`pg_featureserv` is developed under Go 1.13.  It may also work with earlier versions.

* Ensure the Go compiler is installed
* Download or clone this repository into `$GOPATH/src/github.com/CrunchyData/pg_featureserv`
* To build the executable, run the following commands:
```bash
cd $GOPATH/src/github.com/CrunchyData/pg_featureserv/
go build
```
* This creates a `pg_featureserv` executable in the application directory
* (Optional) Run the unit tests using `go test ./...`

### Building a Docker image

* Build the `pg_featureserv` executable with the command:
```bash
CGO_ENABLED=0 go build
```
to avoid the runtime error `/lib64/libc.so.6: version 'GLIBC_2.XX' not found (required by ./pg_featureserv)`.

* In the `$GOPATH/src/github.com/CrunchyData/pg_featureserv/` directory, build the Docker image with:
```bash
docker build -f container/Dockerfile --build-arg VERSION=<VERSION> -t crunchydata/pg_featureserv:<VERSION> ./
```
Replace version `<VERSION>` with the `pg_featureserv` version you are building against.

## Configuring the service

* Set the environment variable `DATABASE_URL` with a Postgres [connection string](https://www.postgresql.org/docs/12/libpq-connect.html#LIBPQ-CONNSTRING)
  * Example: `export DATABASE_URL="host=localhost user=postgres"`
  * Database URL can also be set in the configuration file
* Edit the configuration file `pg_featureserv.toml`, located in `./config`, `/config`, or `/etc`

## Running the service

* If not already done, move to the application directory:
  * `cd $GOPATH/src/github.com/CrunchyData/pg_featureserv`
* Start the server:
  * `./pg_featureserv`
* Open the service home page in a browser:
  * `http:/localhost:9000/home.html`

### Command-line options

* `-?` - show command usage
* `--config file.toml` - specify configuration file to use
* `--debug` - set logging level to TRACE (can also be set in config file)
* `--devel` - run in development mode (e.g. HTML templates reloaded every query)
* `--test` - run in test mode, with an internal catalog of tables and data
* `--version` - display the version number

## Requests Overview

Features are identified by a _collection name_ and _feature id_ pair.

The default response is in JSON/GeoJSON format.
Append `.html` to the request path to see the UI page for the resource.
In a web browser, to request a JSON response append `.json` to the path (which overrides the browser `Accept` header).

The example requests assume that the service is running locally and configured
to listen on port 9000.

- Landing page (HTML or JSON): http://localhost:9000/
- Landing page (HTML): http://localhost:9000/index,html
- Landing page (JSON): http://localhost:9000/index.json
- OpenAPI definition: http://localhost:9000/api
- OpenAPI test UI: http://localhost:9000/api.html
- Conformance: http://localhost:9000/conformance
- Collections: http://localhost:9000/collections
- Collections UI: http://localhost:9000/collections.html
- Feature collection metadata: http://localhost:9000/collections/{name}
- Feature collection UI: http://localhost:9000/collections/{name}.html
- Features from a single feature collection: http://localhost:9000/collections/{name}/items
- Features from a single feature collection (Map UI): http://localhost:9000/collections/{name}/items.html
- Single feature from a feature collection: http://localhost:9000/collections/{name}/items/{featureid}
- Functions (JSON): http://localhost:9000/functions
- Functions UI: http://localhost:9000/functions.html
- Function metadata: http://localhost:9000/functions/{name}
- Function UI: http://localhost:9000/functions/{name}.html
- Features from a function (JSON): http://localhost:9000/functions/{name}/items
- Features from a function (Map UI): http://localhost:9000/functions/{name}/items.html

See [API Summary](API.md) for a summary of the web service API.
