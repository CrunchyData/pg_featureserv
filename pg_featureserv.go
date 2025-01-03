package main

/*
 Copyright 2019 - 2025 Crunchy Data Solutions, Inc.
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
	"fmt"
	"os"

	"github.com/CrunchyData/pg_featureserv/internal/conf"
	"github.com/CrunchyData/pg_featureserv/internal/data"
	"github.com/CrunchyData/pg_featureserv/internal/service"
	"github.com/CrunchyData/pg_featureserv/internal/ui"

	"github.com/pborman/getopt/v2"
	log "github.com/sirupsen/logrus"
)

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

	conf.InitConfig(flagConfigFilename, flagDebugOn)

	if flagTestModeOn || flagDevModeOn {
		ui.HTMLDynamicLoad = true
		log.Info("Running in development mode")
	}
	// Commandline over-rides config file for debugging
	if flagDebugOn || conf.Configuration.Server.Debug {
		log.SetLevel(log.TraceLevel)
		log.Debugf("Log level = DEBUG\n")
	}
	conf.DumpConfig()

	//-- Initialize catalog (with DB conn if used)
	var catalog data.Catalog
	if flagTestModeOn {
		catalog = data.CatMockInstance()
	} else {
		catalog = data.CatDBInstance()
	}
	includes := conf.Configuration.Database.TableIncludes
	excludes := conf.Configuration.Database.TableExcludes
	catalog.SetIncludeExclude(includes, excludes)

	//-- Start up service
	service.Initialize()
	service.Serve(catalog)
}
