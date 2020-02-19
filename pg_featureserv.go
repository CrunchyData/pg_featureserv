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

Browser: e.g. http://localhost:9000/index.html

# Configuration
Database URL in env var `DATABASE_URL`
Example: `export DATABASE_URL="host=localhost"`

# Logging
Logging to stdout
*/

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

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
var flagTestModeOn bool
var flagDebugOn bool
var flagDevModeOn bool
var flagHelp bool
var flagVersion bool
var flagConfigFilename string

func init() {
	initCommnandOptions()
}

func initCommnandOptions() {
	getopt.FlagLong(&flagHelp, "help", '?', "Show command usage")
	getopt.FlagLong(&flagConfigFilename, "config", 'c', "", "config file name")
	getopt.FlagLong(&flagDebugOn, "debug", 'd', "Set logging level to TRACE")
	getopt.FlagLong(&flagDevModeOn, "devel", 0, "Run in development mode")
	getopt.FlagLong(&flagTestModeOn, "test", 't', "Serve mock data for testing")
	getopt.FlagLong(&flagVersion, "version", 'v', "Output the version information")
}

func main() {
	getopt.Parse()

	if flagHelp {
		getopt.Usage()
		os.Exit(1)
	}

	if flagVersion {
		fmt.Printf("%s %s\n", config.AppConfig.Name, config.AppConfig.Version)
		os.Exit(1)
	}

	log.Infof("----  %s - Version %s ----------\n", config.AppConfig.Name, config.AppConfig.Version)

	config.InitConfig(flagConfigFilename)

	log.Infof("%s\n", config.Configuration.Metadata.Title)

	if flagTestModeOn {
		catalogInstance = data.CatMockInstance()
	} else {
		catalogInstance = data.CatDBInstance()
	}
	if flagTestModeOn || flagDevModeOn {
		ui.HTMLDynamicLoad = true
		log.Info("Running in development mode")
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

	bindAddress := fmt.Sprintf("%v:%v", confServ.HttpHost, confServ.HttpPort)
	log.Infof("Serving at %v\n", bindAddress)
	log.Infof("CORS Allowed Origins: %v\n", config.Configuration.Server.CORSOrigins)

	router = initRouter()

	// set CORS handling according to config
	corsOpt := handlers.AllowedOrigins([]string{config.Configuration.Server.CORSOrigins})

	// more "production friendly" timeouts
	// https://blog.simon-frey.eu/go-as-in-golang-standard-net-http-config-will-break-your-production/#You_should_at_least_do_this_The_easy_path
	server := &http.Server{
		ReadTimeout:  time.Duration(config.Configuration.Server.ReadTimeoutSec) * time.Second,
		WriteTimeout: time.Duration(config.Configuration.Server.WriteTimeoutSec) * time.Second,
		Addr:         bindAddress,
		Handler:      handlers.CompressHandler(handlers.CORS(corsOpt)(router)),
	}

	// start http service
	go func() {
		// ListenAndServe returns http.ErrServerClosed when the server receives
		// a call to Shutdown(). Other errors are unexpected.
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	// wait here for interrupt signal (^C)
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig

	// Interrupt signal received:  Start shutting down
	log.Infoln("Shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	server.Shutdown(ctx)

	log.Debugln("Closing DB connections")
	catalogInstance.Close()

	log.Infoln("Server stopped.")
	//log.Fatal(server.ListenAndServe())
}
