---
title: "Web Service API"
date:
draft: false
weight: 50
---

`pg_featureserv` provides a HTTP-based RESTful web service API
to access metadata about as well as data from
the PostGIS objects it publishes.
This section discusses general aspects of the API.

## OpenAPI

The service API is described by an
[OpenAPI](https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.2.md) specification.
This is available as a JSON document at the path `/api`.

The service provides an interactive user interface for
the API at `/api.html`. On this page, you can view the service paths and parameters, and the schemas for the responses. It allows you to try out the API as well.

## OGC API - Features

The service implements a broad subset of the
[OGC API - Features](http://docs.opengeospatial.org/is/17-069r3/17-069r3.html) standard. It implements the following paths defined by the standard:

* `/` - landing page
* `/index,html` - Landing page (HTML)
* `/index.json` - Landing page (JSON)
* `/conformance` - links to conformance resources
* `/api` - API specification OpenAPI document
* `/collections` - list of feature collections
* `/collections.html` - Collections UI
* `/collections/{id}` - metadata for a feature collection
* `/collections/{id}/items` - data set of features from a feature collection
* `/collections/{id}/items.html` - Features from a single feature collection (Map UI)
* `/collections/{id}/items/{fid}` - data for a specific feature
* `/functions` - Functions (JSON)
* `/functions.html` - Functions UI
* `/functions/{name}` - Function metadata
* `/functions/{name}.html` - Function UI
* `/functions/{name}/items` - Features from a function (JSON)
* `/functions/{name}/items.html` - Features from a function (Map UI)

The standard defines various query parameters for certain paths.
Many of these are provided by the service, although some are not yet implemented.

The service extends the standard API to provide richer access to the
capabilities of PostGIS.
Extensions include the `/functions` paths, and additional query parameters. See the other Usage sections for more details.

## Linked data

The **OCG API - Features** standard promotes the concept of [Linked Data](https://www.w3.org/TR/sdw-bp/#linked-data).
This makes web data more usable by providing stable links between related resources.
To enable this the standard, we make sure that response documents
include structured links to other resources. Like most service resources, `pg_featureserv` API response includes a `links` property containing an array of links to related resources.

A structured link includes the following properties:

* `rel` - the name describing the relationship of the current resource to the linked resource
* `href` - the URI for the link
* `type` - the format of the linked resource
* `title` - a title for the linked resource

## CORS

The server supports [Cross-origin Resource Sharing](https://en.wikipedia.org/wiki/Cross-origin_resource_sharing) (CORS) to allow service resources to be
requested by web pages which originate from another domain.
The `Access-Control-Allow-Origin` header required by CORS-compatible responses
can be set via the `CORSOrigins` configuration parameter.

## Request headers

The service behaviour can be influenced by some request headers.
These include:

* `Forwarded` allows a proxy to specify host and protocol for the service Base URL.
* `X-Forwarded-Host` allows a proxy to specify host for Base URL.
* `X-Forwarded-Proto` allows a sproxy to specify protocol for Base URL.
* `Accept` allows a client to indicate what response format(s) it can accept.  Supported values are:
  * `text/html`: indicates HTML
  * `application/json`: indicates JSON
  * `application/geo+json`: indicates GeoJSON

## Request methods

Currently the service provides only Read-Only access to resources.
The only HTTP method supported is `GET`.

## Response formats

The service returns responses in several different formats,
depending on the nature of the request.
Formats include:

* [JSON](https://www.w3.org/TR/sdw-bp/#bib-RFC7159)-formatted text, for non-spatial data
* [GeoJSON](https://tools.ietf.org/rfc/rfc7946.txt) for feature collections and features
* HTML documents for user interface pages

For some requests, there may be more than one format that could be returned.
In particular, many paths provide both a data document (JSON or GeoJSON)
and an HTML view of the data.
The actual format returned is determined in one of the following ways (in descending order of precedence):

* The path extension. Values allowed are:
  * `.json`, which indicates JSON or GeoJSON (the resource itself determines which)
  * `.html`, which indicates an HTML page should be returned, if available
* The `Accept` request header value (see above for supported values).
* If the path extension or `Accept` request header is not specified, the default is to return a data document (JSON or GeoJSON).


When using a web browser to query the service,
the browser generally provides an `Accept` header of `text/html`.
So you may need to explicitly specify the `.json` extension
to retrieve a data document instead of an HTML page.

## Status codes and messages

The HTTP protocol defines a standard set of status codes returned by responses.
`pg_featureserv` can return the following codes:

|  Code  |  Meaning  |
|-------------|-----------|
| `200 OK` | The request has succeeded. |
| `400 Bad Request` | The server could not understand the request due to invalid syntax. |
| `404 Not Found` | The server can not find the requested resource. |
| `500 Internal Server Error` | The server has encountered a situation it is unable to handle. |
| `503 Service Unavailable` | The server is unable to handle the request. Can indicate a timeout caused by a long-running query or very large response. |
