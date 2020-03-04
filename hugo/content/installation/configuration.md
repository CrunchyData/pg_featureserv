---
title: "Configuration"
date:
draft: false
weight: 400
---


## Configuration File

The configuration file is automatically read from the file `config/pg_featureserv.toml`
in the directory the application starts in, if it exists.

If you want to specify a different file, use the `--config` commandline parameter to pass in a full path to the configuration file.  When using the `--config` option the local configuration file is ignored.

```sh
./pg_featureserv --config /opt/pg_featureserv/config.toml
```

If no configuration is specified, the server runs using internal defaults
(which are the same as provided in the example configuration file).
Where possible, the program autodetects values such as the `UrlBase`.

The only required configuration is the `DbConnection` setting,
if not provided in the environment variable `DATABASE_URL`.
(Even this can be omitted if the server is run with the `--test` flag.)

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
AssetsPath = "/usr/share/pg_featurserv/assets"

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

## Configuration Options

### UrlBase

The Base URL is the URL endpoint at which users access the service.
It is also used for any URL paths returned by the service (such as response links).

The `UrlBase` can specify a value for the Base URL.
This accomodates running the service behind a reverse proxy.

If `UrlBase` is not set, `pg_featureserv` dynamically detects the base URL.
Also, if the HTTP headers `Forwarded` or `X-Forwarded-Proto` and `X-Forwarded-Host` are present they are respected.
Otherwise the base URL is determined by inspecting the incoming request.
