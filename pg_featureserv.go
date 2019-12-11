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
Usage: ./pg_featureserv

# Configuration
Database URL in env var `DATABASE_URL`
Example: `export DATABASE_URL="host=localhost"`

# Logging
Logging to stdout
*/

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/CrunchyData/pg_featureserv/config"
	"github.com/CrunchyData/pg_featureserv/data"
	"github.com/gorilla/handlers"
	log "github.com/sirupsen/logrus"
)

// CatalogInstance mock
var catalogInstance data.Catalog
var stateTest bool

func init() {
	initFlags()
}

func initFlags() {
	flag.BoolVar(&stateTest, "test", false, "Run server with test data")
}

func main() {
	flag.Parse()
	log.Printf("%s %s\n", config.AppConfig.Name, config.AppConfig.Version)

	if stateTest {
		catalogInstance = data.CatMockInstance()
	} else {
		catalogInstance = data.CatDBInstance()
	}

	serve()
}

func serve() {

	confServ := config.Configuration.Server
	bindAddress := fmt.Sprintf("%v:%v", confServ.BindHost, confServ.BindPort)
	log.Printf("Serving at: %v\n", bindAddress)

	router := initRouter()
	// set CORS handling to allow all access
	corsOpt := handlers.AllowedOrigins([]string{"*"})
	log.Fatal(http.ListenAndServe(bindAddress, handlers.CORS(corsOpt)(router)))
}
