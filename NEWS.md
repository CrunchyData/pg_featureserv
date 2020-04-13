# pg_featureserv Version History

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
