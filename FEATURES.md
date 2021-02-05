# `pg-featureserv` Functionality

This is a list of current and planned functionality.  It also includes OGC API for Features core requirements (although some of them are not on the current development roadmap).

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
- [x] `bbox` (4 numbers)
- [ ] `bbox` (6 numbers)
- [ ] `bbox-crs`
- [ ] `datetime`
- [x] `properties` list (to restrict attributes in response)
- [ ] filtering by property value ( `name=value`, as per [spec sec. 7.15.5](http://docs.opengeospatial.org/is/17-069r3/17-069r3.html#_parameters_for_filtering_on_feature_properties) )

### Query parameters - Extension
- [x] `orderBy` to order output by a property
  - `orderBy=name`, `orderBy=name:A`, `orderBy=name:D`
- [x] `precision` to set output precision of GeoJSON coordinates
  - `precision=n`
- [x] `transform` to specify geometry transformations
  - `transform=fn,arg,arg|fn,arg`
- [ ] convert transform function names to `ST_` equivalents

- [ ] `f` parameter for formats?  (e.g. `f=json`, `f=html`)

### Query parameters - Functions
- [x] function arguments
  - `name=value`

### Output formats
- [x] GeoJSON
- [x] JSON for metadata
- [x] JSON for non-geometry functions
- [ ] `next` link
- [ ] `prev` link

### Transactions
- [ ] Support POST, PUT, PATCH, DELETE...  TBD

## User Interface (HTML)
- [x] `/home.html` landing page
- [x] metadata for collection
- [x] map display of collection
- [x] attribute display of collection
- [x] map display of single feature
- [x] attribute display of single feature
- [x] function metadata
- [x] map display for geometry functions (`items` page)
- [ ] text display for non-geometry functions (`items` page)
- [x] UI for `limit` parameter on `items` page
- [x] UI for `offset` parameter on `items` page
- [x] UI for `bbox` parameter on `items` page
- [x] UI for setting function parameters on `items` page

## Database

- [x] PostGIS 2.x `AsGeoJSON` geometry with attribute encoding
- [ ] PostGIS 3.0 `ST_AsGeoJSON` record

### Tables / Views
- [x] table column schema
- [x] support tables with no primary key
- [x] support views (with PK as `fid` or missing)
- [x] support materialized views
- [x] read property descriptions from table/view column comments
- [ ] read table estimated and actual extents lazily

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
