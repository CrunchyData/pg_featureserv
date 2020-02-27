---
title: "Installing"
date:
draft: false
weight: 200
---

## Installation

To install `pg_featureserv`, download the binary file. Alternatively, you may run a container. These first two options will suit most use cases; needing to build the executable from source is rare.

### A. Download binaries

Builds of the latest code:

* [Linux](https://postgisftw.s3.amazonaws.com/pg_featureserv_latest_linux.zip)
* [Windows](https://postgisftw.s3.amazonaws.com/pg_featureserv_latest_windows.zip)
* [OSX](https://postgisftw.s3.amazonaws.com/pg_featureserv_latest_osx.zip)

Unzip the file, copy the `pg_featureserv` binary wherever you wish, or use it in place. If you move the binary, remember to move the `assets/` directory to the same location, or start the server using the `AssetsDir` configuration option.

### B. Run container

A Docker image is available on DockerHub:

* [Docker](https://hub.docker.com/r/pramsey/pg_featureserv/)

When you run the container, provide the database connection information in the `DATABASE_URL` environment variable and map the default service port (9000).

```sh
docker run -e DATABASE_URL=postgres://user:pass@host/dbname -p 9000:9000 pramsey/pg_featureserv
```

### C. Build from source

If not already installed, install the [Go software development environment](https://golang.org/doc/install). Make sure that the [`GOPATH` environment variable](https://github.com/golang/go/wiki/SettingGOPATH) is also set.

The application can downloaded and built with the following commands:

```sh
mkdir -p $GOPATH/src/github.com/CrunchyData
cd $GOPATH/src/github.com/CrunchyData
git clone git@github.com:CrunchyData/pg_featureserv.git
cd pg_featureserv
go build
```

To run the build to verify it, set the `DATABASE_URL` environment variable to the database you want to connect to, and run the binary.

```sh
export DATABASE_URL=postgres://user:pass@host/dbname
$GOPATH/bin/pg_featureserv
```
