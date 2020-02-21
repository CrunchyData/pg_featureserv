---
title: "Installing"
date:
draft: false
weight: 225
---

## Installation

To install `pg_featureserv`, download the binary file. Alternatively, you may build the executable from source.

### A. Download Binaries

Builds of the latest code:

* [Linux](https://postgisftw.s3.amazonaws.com/pg_featureserv_latest_linux.zip)
* [Windows](https://postgisftw.s3.amazonaws.com/pg_featureserv_latest_windows.zip)
* [OSX](https://postgisftw.s3.amazonaws.com/pg_featureserv_latest_osx.zip)

Unzip the file, copy the `pg_featureserv` binary wherever you wish, or use it in place. If you move the binary, remember to move the `assets/` directory to the same location, or start the server using the `AssetsDir` configuration option.

### B. Build From Source

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
