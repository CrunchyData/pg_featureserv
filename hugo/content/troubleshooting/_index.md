---
title: "Troubleshooting"
date:
draft: false
weight: 50
---

## Approaches

If the service isn't behaving as expected, there are a few approaches you can use to determine the issue.

### HTTP Response

The service indicates the status of reponses using standard HTTP status codes,
along with text messages.  See the [API](/usage/api/) section
for details of status codes and their meanings.

HTTP status codes and headers returned in service responses can be displayed
by querying them with a command-line utility like [curl](https://curl.haxx.se/):
```sh
curl -I http://localhost:9000/home.json
```
Alternatively, most web browsers provide a debugger which can display detailed response information.

### Service Logging

The service outputs logging information to the console.
By default, the log level is set to show errors and warnings only.
Running the service with debug level logging will
provide more information about request processing.
This can include the actual SQL emitted to the database,
SQL errors, and timing of queries and responses.

To invoke debug mode, run the server with the `--debug` command-line parameter:
```sh
./pg_featureserv --debug
```
You can also turn on debug logging in the [configuration file](/installation/configuration/):
```
Debug = true
```
or via the environment:
```sh
export PGFS_SERVER_DEBUG=true
```

### SQL Logging

The `debug` mode of the server logs the SQL that is being emitted to the database.
If you have access to the database that the service is querying, it can
be useful to try manually executing the SQL.
This can provide more detailed database error reporting.

For issues involving access permissions, it may be useful to
connect as the same user that the service is using.

To delve more deeply into the SQL that is being run on the database, you can turn on [statement logging](https://www.postgresql.org/docs/current/runtime-config-logging.html#GUC-LOG-STATEMENT) in PostgreSQL by editing the `postgresql.conf` file for your database and restarting.

## Missing Tables

If tables are not being discovered as expected ensure that the tables have the correct metadata visible in the geometry views. In particular, note that recent versions of the [`AddGeometryColumn()`](https://postgis.net/docs/AddGeometryColumn.html) method do not update this metadata and will not be discovered by the pg_featureserv catalog.

## Bug reporting

If you find an issue with the feature server, it can be reported on the GitHub issue tracker:

* https://github.com/crunchydata/pg_featureserv/issues

When reporting an issue, please provide the software version being used.
This can be obtained from the service log, or by running:
```sh
./pg_featureserv --version
```
