---
title: "Deploying"
date:
draft: false
weight: 250
---

### Basic Operation

The service can be run with minimal configuration.
Only the database connection information is required.
(Even that can be omitted
if run with the `--test` option.)
The database connection information can be provided in an environment variable
`DATABASE_URL` containing a Postgres [connection string](https://www.postgresql.org/docs/12/libpq-connect.html#LIBPQ-CONNSTRING)

#### Linux or OSX
```sh
export DATABASE_URL=postgresql://username:password@host/dbname
./pg_featureserv
```

#### Windows
```
SET DATABASE_URL=postgresql://username:password@host/dbname
pg_featureserv.exe
```

### Command options

|  Option  |  Description  |
|-------------|-----------|
| `-?` | show command usage |
| `--config <file>.toml` | specify configuration file to use |
| `--debug` | set logging level to TRACE (can also be set in config file) |
| `--devel`| run in development mode (i.e. assets are reloaded for every request) |
| `--test` | run in test mode, using an internal catalog of sample tables and data rather than a database |


### Configuration File

The configuration file is automatically read from the file `config.toml`
in the directory the application starts in, if it exists.

If you want to specify a different file, use the `--config` commandline parameter to pass in a full path to the configuration file.  When using the `--config` option the local configuration file is ignored.

```sh
./pg_featureserv --config /opt/pg_featureserv/pg_featureserv.toml
```

If no configuration is specified, the server runs using internal defaults
(which are the same as provided in the example configuration file).
Where possible the program autodetects things like the `UrlBase`.

The only required configuration is the `DbConnection` setting,
if not provided in the environment variable `DATABASE_URL`.
Even this can be omitted if the server is run with the `--test` flag.

The default configuration file is shown below.

```toml
[Server]
# The hostname to use in links
HttpHost = "0.0.0.0"

# The IP port to listen on
HttpPort = 9000

# Advertise URLs relative to this server name
# default is to look this up from incoming request headers
# UrlBase = "http://localhost:9000/"

# String to return for Access-Control-Allow-Origin header
# CORSOrigins = "*"

# set Debug to true to run in debug mode (can also be done on cmd-line)
# Debug = true

# Read html templates from this directory
AssetsPath = "./assets"

[Database]
# Database connection
# postgresql://username:password@host/dbname
# DbConnection = "postgresql://username:password@host/dbname"

# Close pooled connections after this interval
# 1d, 1h, 1m, 1s, see https://golang.org/pkg/time/#ParseDuration
# DbPoolMaxConnLifeTime = "1h"

# Hold no more than this number of connections in the database pool
# DbPoolMaxConns = 4

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
```
