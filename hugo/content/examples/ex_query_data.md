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

## Basic query

The most basic query against a feature collection is to
retrieve an unfiltered list of the features in a collection.
The number of features returned is limited by the [service
configuration](/installation/configuration/) for the default feature limit.

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

## Query using a bounding box filter and limit

You can control the extent as well as number of features returned with the `bbox` and `limit` query parameters.

For example, to query the countries in the Caribbean and ensure that all of
them are returned, you can use the query parameters like so:
```
http://localhost:9000/collections/ne.countries/items.html
?bbox=-93.0688,9.3746,-54.0296,25.9053&limit=100
```
![Map view of query with bbox and limit](/ex-query-data-countries-bbox-limit.png)

## Query with a property filter and properties list

Another way to limit the features returned is via a property filter query parameter.
For instance, the countries in Europe can be returned using the query parameter `continent=Europe`.

To make it easy to verify the result, the `properties` query parameter has been restricted to only three properties (including `continent` itself).
And as before, a higher `limit` value ensures that all features are returned.

```
http://localhost:9000/collections/ne.countries/items.html
?continent=Europe&properties=gid,name,continent&limit=100
```
![Map view of query with property filter](/ex-query-data-countries-prop-filter.png)

## Query a Feature by ID

You can also query a single feature by providing the feature ID
as part of the resource path.

Most query parameters do not apply to single feature queries. With that said, the `properties` parameter can be used to specify what response properties are included.

```
http://localhost:9001/collections/ne.countries/items/55.html
?properties=gid,name,continent
```
![Map view of query for feature by ID](/ex-query-data-countries-feature.png)
