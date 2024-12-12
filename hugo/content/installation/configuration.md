---
title: "Configuration"
date:
draft: false
weight: 400
---

`pg_featureserv` provides a number of configuration options to control its operation.
Configuration can be provided in a number of ways:

* via a configuration file
* by environment variables
* on the command-line

## Configuration file

The configuration file is automatically read from from the following locations, if it exists:

* In the system configuration directory, at `/etc/pg_featureserv.toml`
* Relative to the directory from which the program is run, `./config/pg_featureserv.toml`
* In a root volume at `/config/pg_featureserv.toml`

If you want to specify a different file, use the `--config` commandline parameter to pass in a full path to the configuration file.  When using the `--config` option, the local configuration file is ignored.

```sh
./pg_featureserv --config /opt/pg_featureserv/config.toml
```

If no configuration is specified, the server runs using internal defaults
(which are the same as provided in the example configuration file below).
Where possible, the program autodetects values such as the `UrlBase`.

The only required configuration is the `DbConnection` setting,
if not provided in the environment variable `DATABASE_URL`.
(This is not required if the server is run with the `--test` flag.)

### Configuration Using Environment Variables

To set the database connection the environment variable `DATABASE_URL`
can be used with a
Postgres [connection string](https://www.postgresql.org/docs/current/libpq-connect.html#LIBPQ-CONNSTRING):
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

### Example Configuration

An example configuration file is shown below.

```toml
[Server]
# Accept connections on this subnet (default accepts on all)
HttpHost = "0.0.0.0"

# Accept connections on this port
HttpPort = 9000
#HttpsPort = 9001

# HTTPS configuration - TLS server certificate full chain and key
# If these are not specified, the TLS server will not be started
#TlsServerCertificateFile = "cert.pem"
#TlsServerPrivateKeyFile = "key.pem"

# Advertise URLs relative to this server name and path
# default is to look this up from incoming request headers
# Note: do not add a trailing slash.
# UrlBase = "http://localhost:9000/"

# Optional path to add to the service base URL
# If set, all routes will be prefixed with this path
# (e.g. "/pg_featureserv", "/services/pg_featureserv", etc.)
# Note: No trailing slash is necessary, pg_featureserv automatically
# adds a trailing slash for you.
# BasePath = "/"

# String to return for Access-Control-Allow-Origin header
# CORSOrigins = "*"

# set Debug to true to run in debug mode (can also be set on cmd-line)
# Debug = true

# Read html templates from this directory
AssetsPath = "/usr/share/pg_featurserv/assets"

# Maximum duration for reading entire request (in seconds)
ReadTimeoutSec = 1

# Maximum duration for writing response (in seconds)
# Also controls maximum time for processing request
WriteTimeoutSec = 30

[Database]
# Database connection
# postgresql://username:password@host/dbname
# DbConnection = "postgresql://username:password@host/dbname"

# Close pooled connections after this interval
# 1d, 1h, 1m, 1s, see https://golang.org/pkg/time/#ParseDuration
# DbPoolMaxConnLifeTime = "1h"

# Hold no more than this number of connections in the database pool
# DbPoolMaxConns = 4

# Publish only these schemas and tables (default is to publish all spatial tables)
# TableIncludes = [ "public", "priv_schema.tbl" ]

# Do not publish these schemas and tables
# TableExcludes = [ "priv_schema", "public.my_tbl" ]

# Publish functions from these schemas (default is publish postgisftw)
# FunctionIncludes = [ "postgisftw", "schema2" ]

[Paging]
# The default number of features in a response
LimitDefault = 20
# Maxium number of features in a response
LimitMax = 10000

[Metadata]
# Title for this service
#Title = "pg-featureserv"
# Description of this service
#Description = "Crunchy Data Feature Server for PostGIS"

[Website]
# URL for the map view basemap
BasemapUrl = "https://maps.wikimedia.org/osm-intl/{z}/{x}/{y}.png"
```

### Configuration options

#### HttpHost

The IP address at which connections are accepted.

#### HttpPort and HttpsPort

The IP ports at which connections are accepted.

#### HTTPS configuration

If the `TlsServerCertificateFile` and `TlsServerPrivateKeyFile`
are specified then HTTPS support will be enabled,
at the port specified by `HttpsPort`.

#### UrlBase

The Base URL is the URL endpoint at which the service is advertised.
It is also used for any URL paths published by the service
(such as URLs for links in response documents).

The `UrlBase` parameter specifies a value for the Base URL. This accomodates running the service behind a reverse proxy.

The provided URL should not have a trailing slash.

##### Example
```
UrlBase = https://my-server.org/features
```

If `UrlBase` is not set, `pg_featureserv` dynamically detects the base URL.
Also, if the HTTP headers `Forwarded` or `X-Forwarded-Proto` and `X-Forwarded-Host` are present, they are respected.
Otherwise the base URL is determined by inspecting the incoming request.

#### BasePath

The BasePath allows an user to change the endpoint where pg_featureserv is served from.
If `BasePath` is not set, all traffic is served from the root of the website `localhost:9000/`.

##### Example
```
BasePath = "/services/pg_featureserv"
```

All traffic would then be served from `localhost:9000/services/pg_featurserv`.

#### CORSOrigins

The string to return in the `Access-Control-Allow-Origin` HTTP header,
which allows providing **Cross-Origin Resource Sharing** (CORS).

#### Debug

Set to `true` to run in debug mode.  This provides debug-level logging.

#### AssetsPath

The directory containing file assets used by the service (such as the HTML templates). It may be more convenient to deploy the asset files
in a location which is not relative to the service application path.

#### ReadTimeoutSec

The maximum duration (in seconds) the service allows for reading the HTTP request.
This can be relatively short, since service requests are small.

#### WriteTimeoutSec

The maximum duration (in seconds) the service allows for
processing and writing the HTTP response.
This should be long enough to allow expected requests to complete,
but not so long that the service can be saturated
by long-running requests.
Long request times may be caused by long execution times for database queries or functions,
or by returning very large responses.

#### DbConnection

The connection to the database can be set in this parameter,
using a Postgres [connection string](https://www.postgresql.org/docs/12/libpq-connect.html#LIBPQ-CONNSTRING).
The database connection can also be set via the `DATABASE_URL` environment variable, which takes precedence over this parameter.

#### DbPoolMaxConnLifeTime

The maximum duration for the lifetime for a pooled connection.
Specified using a Go [duration constant](https://golang.org/pkg/time/#ParseDuration)
such as `1d`, `2.5h`, or `30m`.

#### DbPoolMaxConns

The maximum number of database connections held in the connection pool.

#### TableIncludes

A list of the schemas and tables to publish feature collections from.
The default is to publish all geometry tables.

#### TableExcludes

A list of schemas and tables not to publish.
Overrides items specified in `TableIncludes`.

#### FunctionIncludes

A list of the schemas to publish functions from.
The default is to publish functions in the `postgisftw` schema.

#### LimitDefault

The default number of features in a response,
if not specified by the `limit` [query parameter](/usage/query_data/).

#### LimitMax

The maximum number of features that can be returned in a response.
This cannot be overridden by the `limit` query paramater.

#### Title

The title for the service.
Appears in the HTML web pages, JSON responses, and the log.

#### Description

The description for the service.
Appears in the HTML web pages and JSON responses.

#### BasemapUrl

The URL template for the basemap used in the web UI map views.
Must be a URL template suitable for the OpenLayers OSM class.
