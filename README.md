<p align="center">
  <a href="https://access.crunchydata.com/documentation/pg_tileserv/latest/"><img width="180" height="180" src="./hugo/static/crunchy-spatial-logo.png?raw=true" /></a>
</p>

# pg_featureserv

[![.github/workflows/ci.yml](https://github.com/CrunchyData/pg_featureserv/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/CrunchyData/pg_featureserv/actions/workflows/ci.yml)

A lightweight RESTful geospatial feature server for [PostGIS](https://postgis.net/), written in [Go](https://golang.org/).
It supports the [*OGC API - Features*](https://ogcapi.ogc.org/features/) REST API standard.

See also our companion project [`pg_tileserv`](https://github.com/CrunchyData/pg_tileserv).

## Main features

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

For a full list of software capabilities is available [here](hugo/content/roadmap/capabilities.md).

## Documentation

The full [User Guide](https://access.crunchydata.com/documentation/pg_featureserv/latest/) contains the following chapters:

* [Quick Start](hugo/content/quickstart/_index.md)
* [About pg_featureserv](hugo/content/introduction/_index.md)
* [Summary of the web service API](hugo/content/usage/api.md)
* [Installation](hugo/content/installation/installing.md)
* [Usage](hugo/content/usage/_index.md)
* [Examples](hugo/content/examples/)
* [Troubleshooting](hugo/content/troubleshooting/_index.md)
* [ChangeLog](hugo/content/roadmap/changelog.md)
* [Learn More](hugo/content/learn-more/_index.md)
