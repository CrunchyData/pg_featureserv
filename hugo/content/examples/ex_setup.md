---
title: "Setup"
date:
draft: false
weight: 100
---

This section describes how to create and populate a database that will be used
in the following examples.

## Database preparation

The following terminal commands create a PostGIS database named `naturalearth`
(assuming that your user account has the create database privilege):
```
createdb naturalearth
```

Load the PostGIS extension as superuser (`postgres`):
```
psql -U postgres -d naturalearth -c 'CREATE EXTENSION postgis'
```

We're going to be tidy and load the data into a schema `ne`.
To create the schema run the command:
```
psql -U postgres -d naturalearth -c 'CREATE SCHEMA ne'
```

## Import data

The data used in the examples are loaded from
[Natural Earth](https://www.naturalearthdata.com/downloads/50m-cultural-vectors/).
Download the *Admin 0 - Countries* ZIP and extract to a directory on your computer.

In that directory, run the following terminal command to load the
data into the `ne` schema in the `naturalearth` database.
This creates a new table `ne.admin_0_countries`, with the application user as the owner.

```
shp2pgsql -D -s 4326 ne_50m_admin_0_countries.shp ne.admin_0_countries.shp | psql -U <username> -d naturalearth
```

You should see the `ne.admin_0_countries` table using the `\dt ne.*` command in the `psql` SQL shell.

For more information about publishing spatial tables in `pg_featureserv`
refer to the [Feature collections]({{< relref "collections" >}})
and [Security]({{< relref "security" >}}) sections.

## Configuring the service

Make sure that the service database connection specifies the `naturalearth` database.
As described in the [Configuration]({{< relref "configuration" >}}) section,
this can be provided either by an environment variable:
```
DATABASE_URL=postgres://username:password@localhost/naturalearth
```

or by a configuration file parameter:
```
DbConnection = "postgresql://username:password@localhost/naturalearth"
```

With the service running, you should see the layer listed on the web user interface at
http://localhost:9000/collections.html.
The layer metadata is viewable at http://localhost:9001/collections/ne.admin_0_countries.html.


![pg_tileserv web interface preview](/example-web-preview.png)
