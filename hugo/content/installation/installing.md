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
* [Mac OSX](https://postgisftw.s3.amazonaws.com/pg_featureserv_latest_macos.zip)
* [Docker](https://hub.docker.com/r/pramsey/pg_featureserv)

Unzip the file, copy the `pg_featureserv` binary wherever you wish, or use it in place. If you move the binary, remember to move the `assets/` directory to the same location, or start the server using the `AssetsDir` configuration option.

To run the build to verify it:

* Set the `DATABASE_URL` environment variable to the database you want to connect to, and run the binary.
  `export DATABASE_URL=postgres://username:password@host/dbname`
* Start the server:
  `./pg_featureserv`
* Open the service home page in a browser:
  `http:/localhost:9000/home.html`


### B. Run container

A Docker image is available on DockerHub:

* [Docker](https://hub.docker.com/r/pramsey/pg_featureserv/)

When you run the container, provide the database connection information in the `DATABASE_URL` environment variable and map the default service port (9000).

```sh
docker run -e DATABASE_URL=postgres://username:password@host/dbname -p 9000:9000 pramsey/pg_featureserv
```

Then open the service home page in a browser: `http:/localhost:9000/home.html`


### C. Build from source

`pg_featureserv` is developed under Go 1.13.  It may also work with earlier versions.

In the following, replace version `<VERSION>` with the `pg_featureserv` version are building against.

### Without a Go environment

Without `go` installed, you can build `pg_featureserv` in a docker image:

* Download or clone this repository into `$GOPATH/src/github.com/CrunchyData/pg_featureserv`
* Run the following command in `pg_featureserv/`:

  ```bash
  make APPVERSION=<VERSION> clean build-in-docker
  ```

To run the build to verify it:

* Set the `DATABASE_URL` environment variable to the database you want to connect to, and run the binary.
  `export DATABASE_URL=postgres://username:password@host/dbname`
* Start the server:
  `./pg_featureserv`
* Open the service home page in a browser:
  `http:/localhost:9000/home.html`

### In Go environment

* Ensure the Go compiler is installed. If not already installed, install the [Go software development environment](https://golang.org/doc/install). Make sure that the [`GOPATH` environment variable](https://github.com/golang/go/wiki/SettingGOPATH) is also set.
* Download or clone this repository into `$GOPATH/src/github.com/CrunchyData/pg_featureserv`
* To build the executable, run the following commands:

  ```bash
  cd $GOPATH/src/github.com/CrunchyData/pg_featureserv/
  go build
  ```

* This creates a `pg_featureserv` executable in the application directory

* (Optional) To run the unit tests, use the following command :

  ```bash
  # all tests
  go test ./...
  # OR
  make test
  ```

  Or, to run only CRUD-like tests in mock mode :

  ```bash
  # all CRUD tests (mock mode)
  go test -run ^TestRunnerHandlerMock$ github.com/CrunchyData/pg_featureserv/internal/service/mock_test
  ```

  ```bash
  # a group of CRUD tests (mock mode)
  go test -run ^TestRunnerHandlerMock/DELETE$ github.com/CrunchyData/pg_featureserv/internal/service/mock_test
  ```

  ```bash
  # an unique CRUD test (mock mode)
  go test -run ^TestRunnerHandlerMock/DELETE/TestDeleteExistingFeature$ github.com/CrunchyData/pg_featureserv/internal/service/mock_test
  ```

  It's also possible to run the tests on a database :

  ```bash
  # all CRUD tests (db mode)
  go test -run ^TestRunnerHandlerDb$ github.com/CrunchyData/pg_featureserv/internal/service/db_test
  ```
* To run the build to verify it:

  * Set the `DATABASE_URL` environment variable to the database you want to connect to, and run the binary.
    `export DATABASE_URL=postgres://username:password@host/dbname`
  * Start the server:
    `./pg_featureserv`
  * Open the service home page in a browser:
    `http:/localhost:9000/home.html`


### D. Build the docker image of `pg_featureserv`

```bash
make APPVERSION=<VERSION> clean docker
```
