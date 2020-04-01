---
title: "Troubleshooting"
date:
draft: false
weight: 50
---

## Approaches

If the service is not behaving as expected, there are several approaches
to determine what is happening.

### HTTP Response

The service indicates the status of reponses using standard HTTP status codes,
along with text messages.  See the [API section]({{< relref "api" >}})
for details of status codes and their meanings.

HTTP status codes and headers returned in service responses can be displayed
by running them with a command-line utility like [curl](https://curl.haxx.se/):
```sh
curl -I http://localhost:9000/home.json
```
Alternatively, most web browsers provide a debugger which can display detailed response information.

### Service Logging

The service outputs logging information to the console.
By default, the log level is set to show errors and warnings only.
Running the service with debug level logging will
provide more information about request processing.
This includes things like the actual SQL emitted to the database,
SQL errors, and timing of queries and responses.

To invoke debug mode, run the server with the `--debug` commandline parameter:
```sh
./pg_featureserv --debug
```
Debug logging can also be turned on in the configuration file:
```
Debug = true
```

### SQL Logging

The `debug` mode of the server logs the SQL that is being emitted to the database.
If you have access to the database that the service is querying it can
be useful to manually execute the SQL.
This can provide more detailed database error reporting.
For issues involving access permissions it may be useful to
connect as the same user that the service is using.

If you want to delve more deeply into the SQL that is being run on the database, you can turn on [statement logging](https://www.postgresql.org/docs/current/runtime-config-logging.html#GUC-LOG-STATEMENT) in PostgreSQL by editing the `postgresql.conf` file for your database and restarting.

## Bug Reporting

If you find an issue with the feature server, it can be reported on the GitHub issue tracker:

* https://github.com/crunchydata/pg_featureserv/issues

When reporting an issue please provide the software version being used.
This can be obtained from the service log, or by running:
```sh
./pg_featureserv --version
```
