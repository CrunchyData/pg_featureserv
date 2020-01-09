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
	"os"

	"github.com/CrunchyData/pg_featureserv/ui"

	"github.com/CrunchyData/pg_featureserv/config"
	"github.com/CrunchyData/pg_featureserv/data"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/pborman/getopt/v2"
	log "github.com/sirupsen/logrus"
)

var catalogInstance data.Catalog
var router *mux.Router
var flagTestMode bool
var flagDebugOn bool
var flagHelp bool
var flagConfigFilename string

func init() {
	initCommnandOptions()
}

func initCommnandOptions() {
	getopt.FlagLong(&flagHelp, "help", '?', "Show command usage")
	getopt.FlagLong(&flagDebugOn, "debug", 'd', "Set logging level to TRACE")
	getopt.FlagLong(&flagTestMode, "test", 't', "Serve mock data for testing")
	getopt.FlagLong(&flagConfigFilename, "config", 'c', "", "config file name")
}

func main() {
	getopt.Parse()

	if flagHelp {
		getopt.Usage()
		os.Exit(1)
	}

	log.Infof("----  %s - Version %s ----------\n", config.AppConfig.Name, config.AppConfig.Version)

	config.InitConfig(flagConfigFilename)

	log.Infof("%s\n", config.Configuration.Metadata.Title)

	if flagTestMode {
		catalogInstance = data.CatMockInstance()
		ui.HTMLDynamicLoad = true
	} else {
		catalogInstance = data.CatDBInstance()
	}
	// Commandline over-rides config file for debugging
	if flagDebugOn || config.Configuration.Server.Debug {
		log.SetLevel(log.TraceLevel)
		log.Debugf("Log level = DEBUG\n")
	}
	serve()
}

func serve() {

	confServ := config.Configuration.Server

	bindAddress := fmt.Sprintf("%v:%v", confServ.BindHost, confServ.BindPort)
	log.Infof("Serving at %v\n", bindAddress)
	log.Infof("CORS Allowed Origins: %v\n", config.Configuration.Server.CORSOrigins)

	router = initRouter()

	// set CORS handling to allow all access
	// TODO: make this runtime configurable?
	corsOpt := handlers.AllowedOrigins([]string{config.Configuration.Server.CORSOrigins})

	log.Fatal(http.ListenAndServe(bindAddress, handlers.CORS(corsOpt)(router)))
}
