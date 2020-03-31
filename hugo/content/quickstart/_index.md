---
title: "Quick Start"
date:
draft: false
weight: 5
---

This section describes how to set up `pg_featureserv` and connect the service to a spatial database.

The first half of this section walks through how to prepare a spatial database and import spatial data, using the terminal. If you already have a spatial database, you can go ahead and start with "Configuring the service."

## Database preparation

The following terminal command creates a new database named `naturalearth`
(assuming your user role has the create database privilege):
```
createdb naturalearth
```

Using the `psql` tool, load the PostGIS extension as superuser (we'll go with `postgres`):
```
psql -U postgres -d naturalearth -c 'CREATE EXTENSION postgis'
```

We're going to be tidy and load the data into a schema `ne`.
To create the schema, run the command:
```
psql -U postgres -d naturalearth -c 'CREATE SCHEMA ne'
```

When we get to the step below to connect `pg_featureserv` to the database, the user must have access to the new schema as well.

## Import data

The data used in the examples are loaded from
[Natural Earth](https://www.naturalearthdata.com/downloads/50m-cultural-vectors/).
Download the *Admin 0 - Countries* ZIP and extract to a directory on your computer.

In that directory, run the following terminal command to load the
data into the `ne` schema in the `naturalearth` database.
This creates a new table `countries`, with the application user as the owner.

```
shp2pgsql -D -s 4326 ne_50m_admin_0_countries.shp ne.countries | psql -U <username> -d naturalearth
```

You should see the `ne.countries` table using the `\dt ne.*` command in the `psql` SQL shell.

For more information about publishing spatial tables in `pg_featureserv`,
refer to the [Feature ollections](/usage/collections/)
and [Security](/usage/security/) sections.

## Configuring the service

Make sure that the service database connection specifies the `naturalearth` database.
As described in the [Configuration](/installation/configuration/) section,
this can be provided either by an environment variable:

Linux/OSX
```sh
export DATABASE_URL=postgresql://username:password@host/naturalearth
```

Windows
```
SET DATABASE_URL=postgresql://username:password@host/naturalearth
```

Or by a configuration file parameter:

```toml
DbConnection = "postgresql://username:password@localhost/naturalearth"
```

Download the build of the latest code:

* [Linux](https://postgisftw.s3.amazonaws.com/pg_featureserv_latest_linux.zip)
* [Windows](https://postgisftw.s3.amazonaws.com/pg_featureserv_latest_windows.zip)
* [OSX](https://postgisftw.s3.amazonaws.com/pg_featureserv_latest_osx.zip)

Unzip the file, copy the `pg_featureserv` binary wherever you wish, or use it in place. (If you move the binary, remember to move the `assets/` directory to the same location, or start the server using the `AssetsDir` configuration option.)

## Deploy `pg_featureserv`

In the directory where the `pg_featureserv` binary is located, run the service in the terminal:

Linux/OSX

```sh
./pg_featureserv
```

Windows

```
pg_featureserv.exe
```

With the service running, you should see the layer listed on the web user interface at
http://localhost:9000/collections.html.
The layer metadata is viewable at http://localhost:9000/collections/ne.countries.html.

![pg_featureserv web interface preview](/quickstart-web-preview.png)
