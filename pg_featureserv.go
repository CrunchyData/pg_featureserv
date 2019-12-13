package main

/*
 Copyright 2019 Crunchy Data Solutions, Inc.
 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at
      http://www.apache.org/licenses/LICENSE-2.0
 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

/*
# Running
Usage: ./pg_featureserv [ -test ]

Browser: e.g. http://localhost:9000/home.html

# Configuration
Database URL in env var `DATABASE_URL`
Example: `export DATABASE_URL="host=localhost"`

# Logging
Logging to stdout
*/

import (
	"fmt"
	"net/http"

	"github.com/CrunchyData/pg_featureserv/config"
	"github.com/CrunchyData/pg_featureserv/data"
	"github.com/gorilla/handlers"
	"github.com/pborman/getopt/v2"
	log "github.com/sirupsen/logrus"
)

var catalogInstance data.Catalog
var flagTestMode bool

func init() {
	initFlags()
}

func initFlags() {
	getopt.FlagLong(&flagTestMode, "test", 't', "Serve test data")
}

func main() {
	getopt.Parse()
	log.Infof("%s %s\n", config.AppConfig.Name, config.AppConfig.Version)

	config.InitConfig("config")

	if flagTestMode {
		catalogInstance = data.CatMockInstance()
	} else {
		catalogInstance = data.CatDBInstance()
	}

	serve()
}

func serve() {

	confServ := config.Configuration.Server
	log.Infof("%s\n", config.Configuration.Metadata.Title)

	bindAddress := fmt.Sprintf("%v:%v", confServ.BindHost, confServ.BindPort)
	log.Infof("Serving at %v\n", bindAddress)
	log.Infof("CORS Allowed Origins: %v\n", config.Configuration.Server.CORSOrigins)

	router := initRouter()

	// set CORS handling to allow all access
	// TODO: make this runtime configurable?
	corsOpt := handlers.AllowedOrigins([]string{config.Configuration.Server.CORSOrigins})

	log.Fatal(http.ListenAndServe(bindAddress, handlers.CORS(corsOpt)(router)))
}
