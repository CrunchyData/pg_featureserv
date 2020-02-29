---
title: "pg_featureserv"
date:
draft: false
---

# pg_featureserv
`pg_featureserv` is a [PostGIS](https://postgis.net/)-based feature server written in [Go](https://golang.org/).
It is a lightweight, low-configuration RESTful web service that provides
access to spatial data stored in PostGIS tables, as well as spatial processing capability
based on PostGIS spatial functions.
The service API follows the [OGC API for Features Version 1.0](http://docs.opengeospatial.org/is/17-069r3/17-069r3.html) standard.
It extends the API to expose more of the powerful spatial capabilities of PostGIS.

As noted in the W3C/OGC [Spatial Data on the Web Best Practices](https://www.w3.org/TR/sdw-bp/), exposing spatial data using modern web standards improves spatial data discoverability, accessibility and interoperability.
`pg_featureserv` supports a wide variety of situations where web access
to spatial data enables richer functionality.  Use cases include:

* display features at a point or in an area of interest
* query features using spatial and/or attribute filters
* retrieve features for use in a web application (e.g. for tabular or map display)
* download spatial data for use in applications


This guide walks you through how to install and use `pg_featureserv` for your spatial applications.
The [Usage](/usage/) section goes in-depth on how the service works.
We'll soon be adding more [basic examples](/examples/) of web map applications that source feature data from `pg_featureserv`.
