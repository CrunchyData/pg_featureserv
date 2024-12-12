<p align="center">
  <a href="https://access.crunchydata.com/documentation/pg_featureserv/latest/"><img width="180" height="180" src="./hugo/static/crunchy-spatial-logo.png?raw=true" /></a>
</p>

# pg_featureserv

[![.github/workflows/ci.yml](https://github.com/CrunchyData/pg_featureserv/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/CrunchyData/pg_featureserv/actions/workflows/ci.yml)

A lightweight RESTful geospatial feature server for [PostGIS](https://postgis.net/), written in [Go](https://golang.org/).
It supports the [*OGC API - Features*](https://ogcapi.ogc.org/features/) REST API standard.

See also our companion project [`pg_tileserv`](https://github.com/CrunchyData/pg_tileserv).

## Features

* Implements the [*OGC API - Features*](https://ogcapi.ogc.org/features/) standard.
  * Standard query parameters: `limit`, `bbox`, `bbox-crs`, property filtering, `sortby`, `crs`
  * Query parameters `filter` and `filter-crs` allow [CQL filtering](https://portal.ogc.org/files/96288), with spatial support
  * Extended query parameters: `offset`, `properties`, `transform`, `precision`, `groupby`
* Data responses are formatted in JSON and [GeoJSON](https://www.rfc-editor.org/rfc/rfc7946.txt)
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

### Relevant Standards

* [*OGC API - Features - Part 1: Core*](http://docs.ogc.org/is/17-069r3/17-069r3.html)
* [*OGC API - Features - Part 2: Coordinate Reference Systems by Reference*](https://docs.ogc.org/is/18-058/18-058.html)
* [**DRAFT** *OGC API - Features - Part 3: Filtering*](http://docs.ogc.org/DRAFTS/19-079r1.html)
* [**DRAFT** *Common Query Language (CQL2)*](https://docs.ogc.org/DRAFTS/21-065.html)
* [*GeoJSON*](https://www.rfc-editor.org/rfc/rfc7946.txt)

## Download

Builds of the latest code:

* [Linux](https://postgisftw.s3.amazonaws.com/pg_featureserv_latest_linux.zip)
* [Windows](https://postgisftw.s3.amazonaws.com/pg_featureserv_latest_windows.zip)
* [MacOS](https://postgisftw.s3.amazonaws.com/pg_featureserv_latest_macos.zip)
* [Docker](https://hub.docker.com/r/pramsey/pg_featureserv)


## Build from Source

`pg_featureserv` is developed under Go 1.13.  It may also work with earlier versions.

In the following, replace version `<VERSION>` with the `pg_featureserv` version are building against.

### Without a Go environment

Without `go` installed, you can build `pg_featureserv` in a docker image:

* Download or clone this repository into `$GOPATH/src/github.com/CrunchyData/pg_featureserv`
* Run the following command in `pg_featureserv/`:
  ```bash
  make APPVERSION=<VERSION> clean build-in-docker
  ```

### In Go environment

* Download or clone this repository into `$GOPATH/src/github.com/CrunchyData/pg_featureserv`
* To build the executable, run the following commands:
  ```bash
  cd $GOPATH/src/github.com/CrunchyData/pg_featureserv/
  go build
  ```

* This creates a `pg_featureserv` executable in the application directory
* (Optional) Run the unit tests using `go test ./...`

### Docker image of `pg_featureserv`

#### Build the image

```bash
make APPVERSION=<VERSION> clean docker
```

#### Run the image

To run using an image built above:

```bash
docker run --rm -dt -e DATABASE_URL=postgres://user:pass@host/dbname -p 9000:9000 pramsey/pg_featureserv:<VERSION>
```

## Configure the service

The [configuration file](config/pg_featureserv.toml.example) is automatically read from the following locations, if it exists:

* In the system configuration directory, at `/etc/pg_featureserv.toml`
* Relative to the directory from which the program is run, `./config/pg_featureserv.toml`
* In a root volume at `/config/pg_featureserv.toml`

To specify a configuration file directly use the `--config` commandline parameter.
In this case configuration files in other locations are ignored.

### Configuration Using Environment Variables

To set the database connection the environment variable `DATABASE_URL`
can be used with a
Postgres [connection string](https://www.postgresql.org/docs/12/libpq-connect.html#LIBPQ-CONNSTRING):
```bash
export DATABASE_URL="host=localhost user=postgres"
```

Other parameters in the configuration file can be over-ridden in the environment.
Prepend the upper-cased parameter name with `PGFS_section_` to set the value.
For example, to change the HTTP port, function schemas, and service title:
```bash
export PGFS_SERVER_HTTPPORT=8889
export PGFS_DATABASE_FUNCTIONINCLUDES="postgisftw,my_funs"
export PGFS_METADATA_TITLE="My PGFS"
```

### SSL
For SSL support, a server **private key** and an **authority certificate** are needed.
For testing purposes you can generate a **self-signed key/cert pair** using `openssl`:
```bash
openssl req  -nodes -new -x509  -keyout server.key -out server.crt
```
These are set in the configuration file:
```
TlsServerCertificateFile = "/path/server.crt"
TlsServerPrivateKeyFile = "/path/server.key"
```

## Run the service

* Change to the application directory:
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

## Troubleshooting

To get detailed information about service operation
run with the `--debug` commandline parameter.
```sh
./pg_featureserv --debug
```
Debugging can also be enabled via the configuration file (`Server.Debug=true`),
or in the environment:
```sh
export PGFS_SERVER_DEBUG=true
 ```

## Requests Overview

Features are identified by a _collection name_ and _feature id_ pair.

The default response is in JSON/GeoJSON format.
Append `.html` to the request path to see the UI page for the resource.
In a web browser, to request a JSON response append `.json` to the path (which overrides the browser `Accept` header).

The example requests assume that the service is running locally and configured
to listen on port 9000.

- Landing page (HTML or JSON): http://localhost:9000/
- Landing page (HTML): http://localhost:9000/index.html
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
