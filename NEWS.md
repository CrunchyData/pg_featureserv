# pg_featureserv Version History

## Version NEXT
*Released: TBD*

### New Features

* Support for HTTPS
* Provide the `groupBy` query parameter, and ability to aggregate query features
* Add support for partitioned tables and foreign tables

### Bug Fixes

* Fix response value for empty query results for collections and functions
* Fix OpenAPI document to respect configuration `UrlBase`
* Quote query table and function names to support mixed case
* Fix `collection` reponse to include actual extent (#54)
* Fix `collection` response JSON `extent` property to be OAPIF-compliant (#54)
* Fix landing page document to include links to resources `service-desc` and `conformance`
* Fix links in JSON responses to not include JSON format (required for QGIS)
* Fix function feature requests to respect the `offset` parameter (#65)

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
