# pg_featureserv

[![Travis Build Status][travisbuild]](https://travis-ci.org/CrunchyData/pg_featureserv)

[travisbuild]: https://api.travis-ci.org/CrunchyData/pg_featureserv.svg?branch=master "Travis CI"

A lightweight RESTful geospatial feature server for [PostGIS](https://postgis.net/), written in [Go](https://golang.org/).
It supports the [OGC API - Features](http://docs.opengeospatial.org/is/17-069r3/17-069r3.html) REST API standard.

## Features

* Implements the [OGC API for Features](http://docs.opengeospatial.org/is/17-069r3/17-069r3.html) standard.
* Data reponses are formatted in JSON and GeoJSON
* Provides a simple HTML user interface, with web maps to view spatial data
* Uses the power of PostgreSQL to reduce the amount of code
  and to make data definition easy and familiar.
  * Feature collections are defined by database objects (tables and views)
* Uses PostGIS to provide geospatial functionality:
  * Transforming geometry data into the output coordinate system
  * Marshalling feature data into GeoJSON
* CORS support with configurable Allowed Origins
* GZIP response encoding

## Download

Builds of the latest code:

* [Linux](https://postgisftw.s3.amazonaws.com/pg_featureserv_latest_linux.zip)
* [Windows](https://postgisftw.s3.amazonaws.com/pg_featureserv_latest_windows.zip)
* [OSX](https://postgisftw.s3.amazonaws.com/pg_featureserv_latest_osx.zip)
* [Docker](https://hub.docker.com/repository/docker/pramsey/pg_featureserv)

## Building from Source

pg_featurserv is developed under Go 1.13.  It may also work with earlier versions.

* Ensure the Go compiler is installed
* Download or clone this repository into `$GOPATH/src/github.com/CrunchyData/pg_featureserv`
* To build the executable, run the following commands:
```bash
cd $GOPATH/src/github.com/CrunchyData/pg_featureserv/
go build
```
* This should create a `pg_featureserv` executable in the application directory

## Configuring the service

* Set the environment variable `DATABASE_URL` with a Postgres [connection string](https://www.postgresql.org/docs/12/libpq-connect.html#LIBPQ-CONNSTRING)
  * Example: `export DATABASE_URL="host=localhost"`
* Edit the configuration file `config.toml`, located in the application directory

## Running the service

* If not already done, move to the application directory:
  * `cd $GOPATH/src/github.com/CrunchyData/pg_featureserv`
* Start the server:
  * `./pg_featureserv`
* Open the service home page in a browser:
  * `http:/localhost:9000/home.html`

### Command-line options

* `-?` - show command usage
* `--config file.toml` - specify configuaration file to use
* `--debug` - set logging level to TRACE (can also be set in config file)
* `--devel` - run in development mode (e.g. HTML templates reloaded every query)
* `--test` - run in test mode, with an internal catalog of tables and data

## Requests Overview

Features are identified by a _collection name_ and _feature id_ pair.

The default response is in JSON/GeoJSON format.
Append `.html` to the request path to see the UI page for the resource.

The example requests assume that the service is running locally and configured
to listen on port 9000.

- API landing: http://localhost:9000/
- API definition: http://localhost:9000/api
- Conformance: http://localhost:9000/conformance
- Collections: http://localhost:9000/collections
- Feature collection metadata: http://localhost:9000/collections/{name}
- Features from a single feature collection: http://localhost:9000/collections/{name}/items
- Single feature from a feature collection: http://localhost:9000/collections/{name}/items/{featureid}

See [API Summary](API.md) for a summary of the API
