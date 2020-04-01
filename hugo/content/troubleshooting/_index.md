---
title: "Troubleshooting"
date:
draft: false
weight: 50
---

## Approaches

If the service isn't behaving as expected, there are a few approaches you can use to determine the issue.

### Service Logging

The service outputs logging information to the console.
By default, the log level is set to show errors and warnings only.
To get more information about what is going on behind the scenes,
run the server with the `--debug` commandline parameter:
```sh
./pg_featureserv --debug
```
You can also turn on debug logging in the [configuration file](/installation/configuration/):
```
Debug = true
```

### HTTP Response

Hitting the service endpoints with a command-line utility like [curl](https://curl.haxx.se/)
can yield useful information:
```sh
curl -I http://localhost:9000/home.json
```

### SQL Logging

The `debug` mode of the server logs the SQL that is being emitted to the database.
If you have access to the database that the service is querying, it can
be useful to try manually executing the SQL.
This can provide more detailed database error reporting.

For issues involving access permissions, it may be useful to
connect as the same user that the service is using.

To delve more deeply into the SQL that is being run on the database, you can turn on [statement logging](https://www.postgresql.org/docs/current/runtime-config-logging.html#GUC-LOG-STATEMENT) in PostgreSQL by editing the `postgresql.conf` file for your database and restarting.

## Bug reporting

If you find an issue with the feature server, it can be reported on the GitHub issue tracker:

* https://github.com/crunchydata/pg_featureserv/issues

When reporting an issue, please provide the software version being used.
This can be obtained from the service log, or by running:
```sh
./pg_featureserv --version
```
