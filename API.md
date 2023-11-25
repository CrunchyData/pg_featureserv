# Summary of API

## References

* [OGC API for Features version 1.0](http://docs.opengeospatial.org/is/17-069r3/17-069r3.html)
* [OGC API - Features - Part 3: Filtering and the Common Query Language (CQL)](https://portal.ogc.org/files/96288)
* [OpenAPI Specifcation version 3.0.2](https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.2.md)
* [OGC API - Features - Part 4: Create, Replace, Update and Delete](https://docs.ogc.org/DRAFTS/20-002.html)

## Notes

* The request response format can be indicated by suffixing URLs with `.json` or `.html`
* Paths are given relative to service root path

## Root

Landing page for the service.

### Request
Path: `/` or `index`

### Response

A JSON document some basic service information
and links to more detailed resources.

#### Links
* self - `/index.json` - This document as JSON
* alternate - `/index.html` - This document as HTML
* conformance
* data - `/collections` - collections
* functions - `/functions` - functions

## Feature collections

List collections provided by the service.

### Request
Path: `/collections`

### Response

JSON document listing service collections.

#### Links
* self - `/collections.json` - This document as JSON
* alternate - `/collections.html` - This document as HTML

## Feature collection

Provides metadata about a feature collection.

### Request
Path: `/collections/{cid}`

### Response

JSON document containing feature collection metadata.

#### Links
* self - `/collections/{cid}.json` - This document as JSON
* alternate - `/collections/{cid}.html` - This document as HTML
* items - `/collections/{cid}/items.json` - Features as GeoJSON
* items - `/collections/{cid}/items.html` - Features as HTML

## Features
Access to features in a collection.

### GET
Produces a dataset of items from the collection (as GeoJSON)

#### Request
Path: `/collections/{cid}/items`

### Parameters
* `bbox=minx,miny,maxx,maxy` - filter features in response to ones intersecting a bounding box (in lon/lat or specified CRS).
* `bbox-crs=SRID` - specify CRS for the `bbox` coordinates
* `<propname>=val` - filter features for a property having a value.
  Multiple property filters are ANDed together.
* `filter=cql-expr` - filters features via a CQL expression
* `filter-crs=SRID` - specifies the CRS for geometry values in the CQL filter
* `transform=fun1[,args][|fun2,args...]` - transform the feature geometry by a geometry function pipeline.
* `groupby=PROP-NAME` - group results on a property.
Usually used with an aggregate `transform` function.
* `properties=PROP-LIST`- return only specific properties (comma-separated).
  If PROP-LIST is empty, no properties are returned.
  If not present, all properties are returned.
* `crs=SRID` - specifies the CRS for the output feature geometry
* `precision=N` - set precision of GeoJSON ordinates to use N decimal places
* `sortby=[+|-]PROP` - sort the response items by a property (ascending (default) or descending).
* `limit=N` - limits the number of features in the response.
* `offset=N` - starts the response at an offset.

#### Response

GeoJSON document containing the features resulting from the request query.

#### Links
* self - `/collections/{cid}/items.json` - This document as JSON
* alternate - `/collections/{cid}/items.html` - This document as HTML
* collection - `/collections/{cid}` - The collection document
* next - TBD
* prev - TBD

### POST
Create a feature in collection.

#### Request
Path: `/collections/{cid}/items`
Content: JSON document representing a geojson feature.

#### Response
Empty response with 201 HTTP Status Code.

## Feature
Provides access to one collection feature.

### GET
Get one collection feature.

#### Request
Path: `/collections/{cid}/items/{fid}`

##### Parameters
* `properties=PROP-LIST`- return only the given properties (comma-separated)
* `transform` - transform the feature geometry by the given geometry function pipeline

#### Response

##### Links
* self - `/collections/{cid}/items/{fid}.json` - This document as JSON
* alternate - `/collections/{cid}/items/{fid}.html` - This document as HTML
* collection - `/collections/{cid}` - The collection document

### PUT
Replace one collection feature.
#### Request
Path: `/collections/{cid}/items/{fid}`
Content: JSON document representing a geojson feature.

#### Response
Empty response with 200 HTTP Status Code.

### DELETE
Delete one collection feature.

#### Request
Path: `/collections/{cid}/items/{fid}`

#### Response
Empty response with 200 HTTP Status Code.

## Functions

Lists the functions provided by the service.

### Request
Path: `/functions`

### Response

JSON document listing available functions.

#### Links
* self - `/functions.json` - This document as JSON
* alternate - `/functions.html` - This document as HTML

## Function

Provides metadata about a function

### Request
Path: `/functions/{fnid}`

### Response

#### Links
* self - `/functions/{fnid}.json` - This document as JSON
* alternate - `/functions/{fnid}.html` - This document as HTML
* items - `/functions/{fnid}/items.json` - Features as GeoJSON | JSON
* items - `/functions/{fnid}/items.html` - Features as HTML

## Function Result

Calls a function and produces a result dataset of items (either GeoJSON or JSON).

### Request
Path: `/functions/{fid}/items`

#### Parameters
* `name=value` - Supplies a value for the named function parameter
* `limit=N` - limits the number of features in the response
* `offset=N` - starts the response at the given offset
* `sortby=[+|-]PROP` - sort the response items by a property (ascending (default) or descending).
* `bbox=minx,miny,maxx,maxy` - filter features in response to ones intersecting given bounding box (in lon/lat, for now)
* `properties=PROP-LIST`- return only the given properties (comma-separated)
* `precision=N` - set precision of GeoJSON ordinates to use N decimal places
* `transform` - transform the feature geometry by the given geometry function pipeline

### Response

A GeoJSON or JSON dataset containing function call results

#### Links
* self - `/functions/{fid}/items.json` - This document as JSON
* alternate - `/functions/{fid}/items.html` - This document as HTML
* function - `/functions/{fid}` - The function document
* next - TBD
* prev - TBD
