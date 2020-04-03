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

	"github.com/CrunchyData/pg_featureserv/api"
	"github.com/CrunchyData/pg_featureserv/conf"
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
		fmt.Printf("%s %s\n", conf.AppConfig.Name, conf.AppConfig.Version)
		os.Exit(1)
	}

	log.Infof("----  %s - Version %s ----------\n", conf.AppConfig.Name, conf.AppConfig.Version)

	conf.InitConfig(flagConfigFilename)
	initTransforms(conf.Configuration.Server.TransformFunctions)

	log.Infof("%s\n", conf.Configuration.Metadata.Title)

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
	if flagDebugOn || conf.Configuration.Server.Debug {
		log.SetLevel(log.TraceLevel)
		log.Debugf("Log level = DEBUG\n")
	}
	serve()
}

func createServer() *http.Server {
	confServ := conf.Configuration.Server

	bindAddress := fmt.Sprintf("%v:%v", confServ.HttpHost, confServ.HttpPort)
	log.Infof("Serving at %v\n", bindAddress)
	log.Infof("CORS Allowed Origins: %v\n", conf.Configuration.Server.CORSOrigins)

	router = initRouter()

	// writeTimeout is slighlty longer than request timeout to allow writing error response
	timeoutSecRequest := conf.Configuration.Server.WriteTimeoutSec
	timeoutSecWrite := timeoutSecRequest + 1

	// ----  Handler chain  --------
	// set CORS handling according to config
	corsOpt := handlers.AllowedOrigins([]string{conf.Configuration.Server.CORSOrigins})
	corsHandler := handlers.CORS(corsOpt)(router)
	compressHandler := handlers.CompressHandler(corsHandler)

	// Use a TimeoutHandler to ensure a request does not run past the WriteTimeout duration.
	// This provides a context that allows cancellation to be propagated
	// down to the database driver.
	//(Unfortunately this does not propagate to the database itself.
	// That will require another mechanism such as session config statement_timeout)
	// If timeout expires, service returns 503 and a text message
	timeoutHandler := http.TimeoutHandler(compressHandler,
		time.Duration(timeoutSecRequest)*time.Second,
		api.ErrMsgRequestTimeout)

	// more "production friendly" timeouts
	// https://blog.simon-frey.eu/go-as-in-golang-standard-net-http-config-will-break-your-production/#You_should_at_least_do_this_The_easy_path
	server := &http.Server{
		ReadTimeout:  time.Duration(conf.Configuration.Server.ReadTimeoutSec) * time.Second,
		WriteTimeout: time.Duration(timeoutSecWrite) * time.Second,
		Addr:         bindAddress,
		Handler:      timeoutHandler,
	}
	return server
}

func serve() {

	server := createServer()

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

	// abort after waiting long enough for service to shutdown gracefully
	// this terminates long-running DB queries, which otherwise block shutdown
	abortTimeoutSec := conf.Configuration.Server.WriteTimeoutSec + 10
	chanCancelFatal := FatalAfter(abortTimeoutSec, "Timeout on shutdown - aborting.")

	log.Debugln("Closing DB connections")
	catalogInstance.Close()

	log.Infoln("Server stopped.")
	// cancel the abort since it is not needed
	close(chanCancelFatal)
}
