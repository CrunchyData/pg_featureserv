---
title: "User Interface"
date:
draft: false
weight: 500
---

## Home page

The home page shows the service title and description,
and provides links to the listings of collections and functions,
the OpenAPI definition, and the conformance metadata.

```
http://localhosst:9000/home.html
```

## API User Interface

A user interface for the service API is available at the path `/api.html`.

## List Feature Collections

The path `/collections.html` shows a list of the feature collections exposed by the service.

## Show Feature Collection metadata

The path  `/collections/{collid}.html` shows metadata about the specified feature collection.

## View Feature Collection data on a map

The path `/collections/{collid}/items.html` shows the features returned by a basic query in a web map interface.
The map interface provides a simple UI to allow setting some basic query parameters.

All applicable query parameters may also be appended to the URL.

## View Feature on a map

The path `/collections/{collid}/items/{fid}` shows the feature requested by the query in a web map interface..

All relevant query parameters may also be used.

## List Functions

The path `/functions.html` shows a list of the functions exposed by the service.

## Show Function metadata

The path `/functions/{funid}.html` shows metadata about the specified function.

## View Function result data on a map

The path `/functions/{funid}/items.html` shows the features returned
by a basic function query in a web map interface.
The map interface provides a simple UI to allow specifying function arguments
and setting some basic query parameters.

All relevant query parameters may also be used.
