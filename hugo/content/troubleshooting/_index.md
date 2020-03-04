---
title: "Troubleshooting"
date:
draft: false
weight: 50
---

## Feature Server

The server outputs logging information to the console.
By default, the log level is set to show errors and warnings only.

To get more information about what is going on behind the scenes, run the server with the `--debug` commandline parameter on
```sh
./pg_featureserv --debug
```
Or, turn on debugging in the configuration file:
```
Debug = true
```

## Web Layer

Hitting the service endpoints with a command-line utility like [curl](https://curl.haxx.se/) can also yield useful information:
```sh
curl -I http://localhost:9000/home.json
```

## Database Layer

The debug mode of the feature server returns the SQL that is being called on the database. If you want to delve more deeply into the SQL that is being run on the database, you can turn on [statement logging](https://www.postgresql.org/docs/current/runtime-config-logging.html#GUC-LOG-STATEMENT) in PostgreSQL by editing the `postgresql.conf` file for your database and restarting.

## Bug Reporting

If you find an issue with the feature server, bugs can be reported on GitHub at the issue tracker:

* https://github.com/crunchydata/pg_featureserv/issues
