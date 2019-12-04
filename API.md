## Summary of API

### References

1. [OGC API for Features version 1.0](http://docs.opengeospatial.org/is/17-069r3/17-069r3.html)
1. [OpenAPI Specifcation version 3.0.2](https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.2.md)

### General notes

* The request response format can be indicated by suffixing URLs with `.json` or `.html`

### Root

Path: `/` or `home`

#### Links
* self - `/` - This document as JSON | HTML
* alternate - `/` - This document as HTML | JSON
* conformance
* data - `/collections` - collections

### Feature collections

Path: `/collections`

#### Links
* self - `/collections.json|html` - This document as JSON | HTML
* alternate - `/collections.html|json` - This document as HTML | JSON

### Feature collection

Path: `/collections/{cid}`

#### Links
* self - `/collections/{cid}.json|html` - This document as JSON | HTML
* alternate - `/collections/{cid}.html|json` - This document as HTML | JSON
* items - `/collections/{cid}/items.json` - Features as GeoJSON
* items - `/collections/{cid}/items.html` - Features as HTML

### Features

Path: `/collections/{cid}/items`

### Parameters
TBD

#### Links
* self - `/collections/{cid}/items.json|html` - This document as JSON | HTML
* alternate - `/collections/{cid}/items.html|json` - This document as HTML | JSON
* collection - `/collections/{cid}` - The collection document
* next - TBD
* prev - TBD

### Feature

Path: `/collections/{cid}/items/{fid}`

#### Links
* self - `/collections/{cid}/items/{fid}.json|html` - This document as JSON | HTML
* alternate - `/collections/{cid}/items/{fid}.html|json` - This document as HTML | JSON
* collection - `/collections/{cid}` - The collection document
