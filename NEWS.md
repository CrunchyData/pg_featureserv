# pg_featureserv Version History

## Version ??
*Released: TBD*

### Bug Fixes

* Fix CQL parser to allow multiple AND/OR terms (#162)
* Fix collection JSON output to use nested extent bbox as per spec (#175)

## Version 1.3.1
*Released: 2024 Feb 6*

### Improvements

* Container build update (#154)
* Improve map UI layout (#146)

### Bug Fixes

* Fix CQL parser to allow scientific notation using `e` (#141)


## Version 1.3
*Released: 2023 July 10*

### New Features

* Support for HTTPS
* Support Postgres data types: JSON, array of text/int/float/numeric
* Add support for partitioned tables and foreign tables
* Add support for tables storing `geography`
* Add configuration to include/exclude published schemas and tables
* Add configuration for web UI map view basemap URL template
* Add configuration to set base path (#88)
* Add configuration via enviroment variables (#104)
* Allow configuring published function schemas (#99)
* Add the `groupby` query parameter, and ability to aggregate query features
* Support OGC API query parameter `sortby` (`orderBy` is deprecated)
* Add OGC API query parameters `crs` and `bbox-crs` (#98)
* Add CQL filtering query parameters `filter` and `filter-crs`  (#101, #102, #103, #105)

### Performance Improvements

* Improve performance of page template loading using caching

### Bug Fixes

* Fix response value for empty query results for collections and functions
* Fix OpenAPI document to respect configuration `UrlBase`
* Quote query table and function names to support mixed case
* Fix `collection` reponse to include actual extent (#54)
* Fix `collection` response JSON `extent` property to be OAPIF-compliant (#54)
* Fix landing page document to include links to resources `service-desc` and `conformance`
* Fix links in JSON responses to not include JSON format (required for QGIS)
* Fix function feature requests to respect the `offset` parameter (#65)
* Fix Swagger request to work under HTTPS (#71)
* Fix Collection description to be populated from table description (#70)
* Fix output of NULL geometry values to be JSON null
* Fix `offset` parameter to allow any non-negative value
* Fix encoding of primary key column names in item requests (#80)
* Fix error handling during Feature reading from database (#96)
* Check for empty strings in Forwarded header parsing

## Version 1.2
*Released: 2020 Dec 1*

### New Features

* Display materialized views as collections
* Improve logging and error messages

### Bug Fixes

* Handle specialty column types by casting to text (#52)

## Version 1.1
*Released: 2020 Apr 6*

### New Features

* Query parameters for filtering by property value
* Configuration write timeout cancels service processing (but not DB process)
* Handle client cancellation
* Improved collection metadata JSON doc
* Data-driven styling for points
* Data-driven styling column `style_radius`
* Configuration whitelist for geometry transformation functions
* Quote database columns in queries to allow upper case, reserved SQL words
* More detailed debug logging


### Bug Fixes

* Fix handling Accept headers to work for all paths
* Fix configuration loading to give priority to env var DATABASE_URL over config file
* Fix function argument ordering in database metadata query
* Fix function metadata query to only show functions with `execute` privilege
* Fix handling of functions with no default parameters
* Fix Home page JSON link

## Version 1.0.0
*Released: 2020 Feb 29*
