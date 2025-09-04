## Demos

Contains example code to demonstrate the capabilities of `pg_featureserv`.

### Quick start (Docker Compose)

```sh
cd demo
docker compose up -d
```

Access the UI: `http://localhost:9000`

The SQL commands in demo/initdb will be run once on database creation. To re-run the init scripts, tear down the database and recreate.

### Routing example

A demo using OpenLayers and PgRouting.
See [instructions](routing/openlayers-pgrouting.md).

### Demo functions

Various PL/pgSQL functions
which demonstrate the capability of exposing functions.
The functions are deployed in the `postgisftw` schema.
