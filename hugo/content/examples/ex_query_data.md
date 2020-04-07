---
title: "Example: Querying Features"
date:
draft: false
weight: 200
---

This example shows how to use the `pg_featureserv` API to query the
`ne.countries` feature collection
created in the [Quick Start](/quickstart/) section.

For more information about querying feature collections,
see the [Usage](/usage/) section.

## Basic Query

The most basic query against a feature collection is to
retrieve an unfiltered list of the features in a collection.
The number of features returned is limited by the service
configuration for the default feature limit.

The following query returns a partial list of
the countries in the `ne.countries` collection, as a GeoJSON FeatureCollection:
```
http://localhost:9000/collections/ne.countries/items.json
```

The query can also be returned as a map view in the web UI:
```
http://localhost:9000/collections/ne.countries/items.html
```
which should display a page like this:

![Map view of basic query](/ex-query-data-countries-basic.png)

## Query using a Bounding Box Filter and Limit

The extent and number of features returned can be controlled
by the `bbox` and `limit` query parameters.
To query the countries in the Caribbean and ensure that all of
them are returned this query could be used:
```
http://localhost:9000/collections/ne.countries/items.html
?bbox=-93.0688,9.3746,-54.0296,25.9053&limit=100
```
![Map view of query with bbox and limit](/ex-query-data-countries-bbox-limit.png)

## Query with a Property filter and Properties list

Another way to limit the features returned is via a property filter query parameter.
For instance, the countries in Europe can be returned via using the query parameter `continent=Europe`.
To make it easy to verify the result, the returned properties have been
restricted using the `properties` query parameter.
And as before a higher `limit` value is used to ensure all features are returned.
```
http://localhost:9000/collections/ne.countries/items.html
?continent=Europe&properties=gid,name,continent&limit=100
```
![Map view of query with property filter](/ex-query-data-countries-prop-filter.png)

## Query a Feature by ID

A single feature can queried by providing the feature ID
as part of the resource path.
Fitering parameters do not apply to feature queries,
but the response properties can be specified
using `properties`.

```
http://localhost:9001/collections/ne.countries/items/55.html
?properties=gid,name,continent
```
![Map view of query for feature by ID](/ex-query-data-countries-feature.png)
