# Summary of API

## References

1. [OGC API for Features version 1.0](http://docs.opengeospatial.org/is/17-069r3/17-069r3.html)
1. [OpenAPI Specifcation version 3.0.2](https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.2.md)

## Notes

* The request response format can be indicated by suffixing URLs with `.json` or `.html`
* Paths are given relative to service root

## Root

Path: `/`, `index`

### Links
* self - `/index.json` - This document as JSON
* alternate - `/index.html` - This document as HTML
* conformance
* data - `/collections` - collections
* functions - `/functions` - functions

## Feature collections

Path: `/collections`

### Links
* self - `/collections.json|html` - This document as JSON | HTML
* alternate - `/collections.html|json` - This document as HTML | JSON

## Feature collection

Path: `/collections/{cid}`

Metadata about the collection

### Links
* self - `/collections/{cid}.json` - This document as JSON
* alternate - `/collections/{cid}.html` - This document as HTML
* items - `/collections/{cid}/items.json` - Features as GeoJSON
* items - `/collections/{cid}/items.html` - Features as HTML

## Features

Produces a dataset of items from the collection (as GeoJSON)

Path: `/collections/{cid}/items`

### Parameters
* `limit=N` - limits the number of features in the response
* `offset=N` - starts the response at the given offset
* `orderBy=PROP[:A | :D]` - order the response items by the given property (ascending (default) or descending)
* `properties=PROP-LIST`- return only the given properties (comma-separated)
* `transform` - transform the feature geometry by the given geometry function pipeline

### Response

#### Links
* self - `/collections/{cid}/items.json` - This document as JSON
* alternate - `/collections/{cid}/items.html` - This document as HTML
* collection - `/collections/{cid}` - The collection document
* next - TBD
* prev - TBD

## Feature

### Request
Path: `/collections/{cid}/items/{fid}`

#### Parameters
* `properties=PROP-LIST`- return only the given properties (comma-separated)
* `transform` - transform the feature geometry by the given geometry function pipeline

### Response

#### Links
* self - `/collections/{cid}/items/{fid}.json` - This document as JSON
* alternate - `/collections/{cid}/items/{fid}.html` - This document as HTML
* collection - `/collections/{cid}` - The collection document

## Functions

Path: `/functions`

### Response

#### Links
* self - `/functions.json` - This document as JSON
* alternate - `/functions.html` - This document as HTML

## Function

Path: `/functions/{fnid}`

Metadata about the function

### Response

#### Links
* self - `/functions/{fnid}.json` - This document as JSON
* alternate - `/functions/{fnid}.html` - This document as HTML
* items - `/functions/{fnid}/items.json` - Features as GeoJSON | JSON
* items - `/functions/{fnid}/items.html` - Features as HTML

## Function Result

Path: `/functions/{fid}/items`

Produces a dataset of items from the function (either GeoJSON or JSON)

#### Parameters
* `limit=N` - limits the number of features in the response
* `offset=N` - starts the response at the given offset
* `orderBy=PROP[:A | :D]` - order the response items by the given property (ascending (default) or descending)
* `properties=PROP-LIST`- return only the given properties (comma-separated)
* `transform` - transform the feature geometry by the given geometry function pipeline

### Response

#### Links
* self - `/functions/{fid}/items.json` - This document as JSON
* alternate - `/functions/{fid}/items.html` - This document as HTML
* function - `/functions/{fid}` - The function document
* next - TBD
* prev - TBD
