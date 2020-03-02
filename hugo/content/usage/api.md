---
title: "Web Service API"
date:
draft: false
weight: 50
---

`pg_featureserv` provides a HTTP-based RESTful web service API
to access metadata about and data from
the PostGIS objects it exposes.
This section discusses general aspects of the API.

## OGC API - Features

The service implements a broad subset of the
[OGC API - Features](http://docs.opengeospatial.org/is/17-069r3/17-069r3.html) standard.
it implements the following paths defined by the standard:

* `/` - landing page
* `/conformance` - links to conformance resources
* `/api` - API specification OpenAPI document
* `/collections` - list of feature scollections
* `/collections/{id}` - metadata for a feature collection
* `/collections/{id}/items` - dataset of features from a feature collection
* `/collections/{id}/items/{fid}` - data for a specific feature

The standard defines various query parameters for certain paths.
Many these are provided.

The service extends the standard API to provide richer access to the
capabilities of PostGIS.
Extensions include the `/functions` paths, and additional query parameters.

See the subsequent sections for details.

## Linked Data

The **OCG API - Features** standard promotes the concept of [Linked Data](https://www.w3.org/TR/sdw-bp/#linked-data).
This makes web data more usable by providing stable links between related resources.
To enable this the standard mandates that response documents
include structured links to other resources.

A structured link includes the following properties:

* `rel` - the name describing the relationship of the current resource to the linked resource
* `href` - the URI for the link
* `type` - the format of the linked resource
* `title` - a title for the linked resource

Most service resources include a `links` property containing an array of links
to related resources.

## OpenAPI

The service API is described by an
[OpenAPI](https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.2.md) specification.
This is available as a JSON document at the path `/api`.
The service provides an interactive user interface for
the API at `/api.html`.

## CORS

The server supports [Cross-origin Resource Sharing](https://en.wikipedia.org/wiki/Cross-origin_resource_sharing) (CORS) to allow service resources to be
requested by web pages which originate from another domain.
The `Access-Control-Allow-Origin` header required by CORS-compatible responses
can be set via the `CORSOrigins` configuration parameter.

## Request Headers

The service behaviour can be influenced by some request headers.
These include:

* `Forwarded` allows a proxy to specify host and protocol for service base URL.
* `X-Forwarded-Host` allows a proxy to specify host for service base URL.
* `X-Forwarded-Proto` allows a sproxy to specify protocol for service base URL.
* `Accept` allows a client to indicate what response format(s) it can accept.  Supported values are:
  * `text/html` indicates HTML
  * `application/json` indicates JSON
  * `application/geo+json` indicates GeoJSON

## Request Methods

Currently the service provides only Read-Only access to resources.
Consequently the only HTTP method supported is `GET`.

## Response Formats

The service returns responses in several different formats,
depending on the nature of the request.
Formats include:

* [JSON](https://www.w3.org/TR/sdw-bp/#bib-RFC7159)-formatted text, for non-spatial data
* [GeoJSON](https://tools.ietf.org/rfc/rfc7946.txt) feature collections and features
* HTML documents for user interface pages

For some requests there may be more than one format which could be returned.
In particular, many paths provide both a data document (JSON or GeoJSON)
and an HTML view of the data.
The actual format returned is determined in one of the following ways:

* the default is to return a data document (JSON or GeoJSON)
* the format can be determined by the request `Accept`
* the format can be specified by the path extension. Values are:
  * `.json` indicates JSON or GeoJSON (the resource determines which)
  * `.html` indicates an HTML page should be returned, if available

A web browser generally provides an `Accept` header of `text/html`,
so you may need to explicitly specify the `.json` extension
to retrieve a data document.


## Status Codes and Messages

The HTTP protocol defines a standard set of status codes returned by responses.
The codes which may be returned by the service are listed below.

|  Code  |  Meaning  |
|-------------|-----------|
| `200 OK` | The request has succeeded. |
| `400 Bad Request` | The server could not understand the request due to invalid syntax. |
| `404 Not Found` | The server can not find the requested resource. |
| `500 Internal Server Error` | The server has encountered a situation it is unable to handle. |
