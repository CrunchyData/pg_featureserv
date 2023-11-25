# `pg-featureserv` Functionality

This is a list of current and planned functionality.
It includes [*OGC API - Features*](http://docs.opengeospatial.org/is/17-069r3/17-069r3.html) core requirements (although some of them are not on the current development roadmap).

## API

- [x] determine response format from request headers `Content-Type`, `Accept`
- [x] CORS support
- [x] GZIP encoding
- [x] HTTPS support
- [x] Proxy support via configurable base URL

### Schema

- [x] `/api` - OpenAPI schema endpoint
- [x] `/conformance` - OAF Conformance endpoint

### Resources
- [x] `/` landing page
- [x] `/collections`
- [x] `/collections/id`
- [x] `/collections/id/items`
- [x] `/collections/id/items/id`
- [x] `/functions`
- [x] `/functions/id`
- [x] `/functions/id/items`

### Resource Metadata
- [x] `/collections/id` JSON includes property names/types
- [x] `/functions/id` JSON includes parameter names/types/defaults and property names/types

### Query parameters - Standard
- [x] `limit=n`
- [x] `offset=n`
- [x] `crs=srid`
- [x] `bbox=x1,y1,x2,y2`
- [ ] `bbox` (6 numbers)
- [x] `bbox-crs=srid`
- [ ] `datetime`
- [x] `properties` list
  - restricts properties included in response
- [x] `sortby` to sort output by a property
  - `sortby=name`, `sortby=+name`, `sortby=-name`
- [x] filtering by property value ( `name=value`, as per [spec sec. 7.15.5](http://docs.opengeospatial.org/is/17-069r3/17-069r3.html#_parameters_for_filtering_on_feature_properties) )
- [x] `filter` with CQL expressions (see below)
- [ ] `filter-lang` (only CQL-Text is supported)
- [ ] `filter-crs=srid`

### Query parameters - Extension
- [x] `precision` to set output precision of GeoJSON coordinates
  - `precision=n`
- [x] `transform` to specify geometry transformations
  - `transform=fn,arg,arg|fn,arg`
- [ ] convert transform function names to `ST_` equivalents
- [x] `groupBy=colname` to group by column (used with a `transform` spatial aggregate function)
- [ ] `f` parameter for formats?  (e.g. `f=json`, `f=html`)

### Query parameters - Functions
- [x] function arguments
  - `name=value`

### CQL expressions
- [x] property names
- [x] character literals
- [x] numeric literals
- [x] arithemetic expressions
  - `+`,`-`,`*`,`/`, `%`, parentheses
- [x] binary comparisons
  - `<`,`<=`,`>`,`>=`,`=`,`<>`
- [x] `property [NOT] BETWEEN a AND B`
- [x] `property [NOT] IN ( value-list )`
- [x] `property [NOT] (LIKE | ILIKE) pattern`
  - `pattern` can include `%` wildcards
- [x] `property [NOT] IS NULL`
- [x] boolean combinations (`AND`,`OR`,`NOT`, parentheses)
- [x] geometry literals
  - `POINT`,`LINESTRING`,`POLYGON`,`MULTIPOINT`,`MULTILINESTRING`,`MULTIPOLYGON`,`GEOMETRYCOLLECTION`,`ENVELOPE`
- [x] spatial predicates
  - `INTERSECTS`,`DISJOINT`,`CONTAINS`,`WITHIN`,`EQUALS`,`CROSSES`,`OVERLAPS`,`TOUCHES`
- [x] distance predicate
  - `DWITHIN`
- [x] temporal literals
  - `1999-01-01`, `2001-12-25T10:01:02`
- [x] temporal predicates
- [ ] functions

### Output formats
- [x] GeoJSON
- [ ] GML
- [x] JSON for metadata
- [x] JSON for non-geometry functions
- [ ] `next` link
- [ ] `prev` link

### Input formats
- [x] GeoJSON
- [ ] GML

### Transactions
- [X] Support POST, PUT, DELETE on tables with primary key
- [ ] Support PATCH...  TBD
- [ ] Support Optimistic locking

## User Interface (HTML)
- [x] `/home.html` landing page
- [x] metadata for collection
- [x] map display of collection
- [x] attribute display of collection
- [x] map display of single feature
- [x] attribute display of single feature
- [x] function metadata
- [ ] text display for non-geometry functions (`items` page)

### Map UI
- [x] map display for features (`items` page)
- [x] map display for geometry functions (`items` page)
- [x] map panel showing features attributes
- [x] control for `limit` parameter
- [x] control for `offset` parameter
- [x] control for `bbox` parameter
- [x] control for setting function parameter values
- [x] configurable base map URL

## Database

- [x] PostGIS 2.x `AsGeoJSON` geometry with attribute encoding
- [ ] PostGIS 3.0 `ST_AsGeoJSON` record

### Data Types
- [x] common scalar types: text, int, float, numeric
- [x] Arrays of text, int, float, numeric
- [x] JSON
- [x] Other types converted to text representation

### Tables / Views
- [x] table column schema
- [x] support tables with no primary key
- [x] support views (with PK as `fid` or missing)
- [x] support materialized views
- [x] read property descriptions from table/view column comments
- [ ] read table estimated and actual extents lazily
- [X] include/exclude published schemas and tables via configuration

### Functions
- [x] support functions returning geometry
- [x] support functions returning attribute-only data
- [x] support geometry functions with `id` output field
- [x] support geometry functions with no `id` output field
- [x] LIMIT/OFFSET for function output
- [x] BBOX filter for function output
- [ ] pass LIMIT as function parameter
- [ ] pass BBOX as function parameter

## Operational

- [x] graceful shutdown
- [x] enforce request timeouts

### Configuration
- [x] read config from file
- [ ] log levels
- [x] DB pool parameters
- [x] database connection string
- [x] whitelist for transformation functions (default: none)

### Security
- [ ] Authentication - TBD
- [ ] OpenID - TBD
- [x] Authorization via database role & grants
