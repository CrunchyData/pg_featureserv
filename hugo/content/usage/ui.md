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
http://localhost:9000/index.html
```

![pg_featureserv UI home page](/ui-home-page.png)

## API user interface

A user interface for the service API is available at the path `/api.html`.

## List feature collections

The path `/collections.html` shows a list of the feature collections published by the service.

## Show feature collection metadata

The path  `/collections/{collid}.html` shows metadata about the specified feature collection.

## View features on a map

The path `/collections/{collid}/items.html` shows the features returned by a basic query in a web map interface.
The map interface provides a simple UI that allows setting some basic [query parameters](./query_data/).

Any applicable query parameters may be appended to the URL.

## View a feature on a map

The path `/collections/{collid}/items/{fid}` shows the feature requested by the query in a web map interface.

Any applicable query parameters may be appended to the URL.

## List functions

The path `/functions.html` shows a list of the functions published by the service.

## Show function metadata

The path `/functions/{funid}.html` shows metadata about the specified function.

## View function result data on a map

The path `/functions/{funid}/items.html` shows the features returned
by a basic function query in a web map interface.
The map interface provides a simple UI that allows specifying function arguments
and setting some basic [query parameters](./query_function/).

Note that only functions with spatial results can be viewed on a map.

Any applicable query parameters may be appended to the URL.
