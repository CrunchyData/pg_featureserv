package config

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

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	ConfigFileNameDefault = "config"
)

// Configuration for system
var Configuration Config

func setDefaultConfig() {
	viper.SetDefault("Server.BindHost", "localhost")
	viper.SetDefault("Server.BindPort", 9000)
	viper.SetDefault("Server.CORSOrigins", "*")
	viper.SetDefault("Server.Debug", false)

	viper.SetDefault("Paging.LimitDefault", 10)
	viper.SetDefault("Paging.LimitMax", 1000)

	viper.SetDefault("Metadata.Title", "pg-featureserv Demo")
	viper.SetDefault("Metadata.Description", "Demo of Crunchy Data PostGIS Feature Server")
}

// Config for system
type Config struct {
	Server   Server
	Paging   Paging
	Metadata Metadata
	//Database Database
}

// Server config
type Server struct {
	BindHost    string
	BindPort    int
	CORSOrigins string
	Debug       bool
}

// Paging config
type Paging struct {
	LimitDefault int
	LimitMax     int
}

// Database config
type Database struct {
	ConnectString string
}

// Metadata config
type Metadata struct {
	Title       string
	Description string
}

// InitConfig initializes the configuration from the config file
func InitConfig(configFilename string) {
	// --- defaults
	setDefaultConfig()

	isExplictConfigFile := configFilename != ""
	confFile := ConfigFileNameDefault
	if configFilename != "" {
		viper.SetConfigFile(configFilename)
		confFile = configFilename
	} else {
		viper.SetConfigName(confFile)
		viper.AddConfigPath(".")
	}
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {
		_, isConfigFileNotFound := err.(viper.ConfigFileNotFoundError)
		errrConfRead := fmt.Errorf("Fatal error reading config file: %s", err)
		isUseDefaultConfig := isConfigFileNotFound && !isExplictConfigFile
		if isUseDefaultConfig {
			confFile = "DEFAULT" // let user know config is defaulted
			log.Debug(errrConfRead)
		} else {
			log.Fatal(errrConfRead)
		}
	}
	log.Infof("Using config file: %s", confFile)
	viper.Unmarshal(&Configuration)

	//fmt.Printf("Viper: %v\n", viper.AllSettings())
	//fmt.Printf("Config: %v\n", Configuration)
}
