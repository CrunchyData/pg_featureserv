---
title: "pg_featureserv"
date:
draft: false
---

![Crunchy Spatial](/crunchy-spatial-logo.png)

# pg_featureserv

`pg_featureserv` is a [PostGIS](https://postgis.net/)-based feature server written in [Go](https://golang.org/).
It is a lightweight, low-configuration RESTful web service that provides
access to spatial data stored in PostGIS tables, as well as spatial processing capability
based on PostGIS spatial functions.

`pg_featureserv` supports a wide variety of situations where web access
to spatial data enables richer functionality.  Use cases include:

* Display features at a point or in an area of interest
* Query features using spatial and/or attribute filters
* Retrieve features for use in a web application (e.g. for tabular or map display)
* Download spatial data for use in applications

This guide walks you through how to install and use `pg_featureserv` for your spatial applications. See [Quick Start](/quickstart/) to learn how to get the service up and running with a spatial database. The [Usage](/usage/) section goes in-depth on how the service works.
We're continuing to add [basic examples](/examples/) of working with feature data from `pg_featureserv`.
