---
title: "Deploying"
date:
draft: false
weight: 300
---

## Basic Operation

The service can be run with minimal configuration.
Only the database connection information is required.
(Even that can be omitted
if run with the `--test` option.)
The database connection information can be provided in an environment variable
`DATABASE_URL` containing a Postgres [connection string](https://www.postgresql.org/docs/12/libpq-connect.html#LIBPQ-CONNSTRING).

### Linux or OSX
```sh
export DATABASE_URL=postgresql://username:password@host/dbname
./pg_featureserv
```

### Windows
```
SET DATABASE_URL=postgresql://username:password@host/dbname
pg_featureserv.exe
```

## Command options

|  Option  |  Description  |
|-------------|-----------|
| `-?` | Show command usage |
| `--config <file>.toml` | Specify configuration file to use. |
| `--debug` | Set logging level to TRACE (can also be set in config file). |
| `--devel`| Run in development mode.  Assets are reloaded on every request. |
| `--test` | Run in test mode.  Uses an internal catalog of sample tables and data.  Does not require a database. |
