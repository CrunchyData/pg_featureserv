package conf

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

import (
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Configuration for system
var Configuration Config

func setDefaultConfig() {
	viper.SetDefault("Server.HttpHost", "0.0.0.0")
	viper.SetDefault("Server.HttpPort", 9000)
	viper.SetDefault("Server.HttpsPort", 9001)
	viper.SetDefault("Server.TlsServerCertificateFile", "")
	viper.SetDefault("Server.TlsServerPrivateKeyFile", "")
	viper.SetDefault("Server.UrlBase", "")
	viper.SetDefault("Server.BasePath", "")
	viper.SetDefault("Server.CORSOrigins", "*")
	viper.SetDefault("Server.Debug", false)
	viper.SetDefault("Server.AssetsPath", "./assets")
	viper.SetDefault("Server.ReadTimeoutSec", 5)
	viper.SetDefault("Server.WriteTimeoutSec", 30)

	viper.SetDefault("Database.DbPoolMaxConnLifeTime", "1h")
	viper.SetDefault("Database.DbPoolMaxConns", 4)
	viper.SetDefault("Database.TableIncludes", []string{})
	viper.SetDefault("Database.TableExcludes", []string{})
	viper.SetDefault("Database.FunctionIncludes", []string{"postgisftw"})

	viper.SetDefault("Paging.LimitDefault", 10)
	viper.SetDefault("Paging.LimitMax", 1000)

	viper.SetDefault("Metadata.Title", "pg-featureserv")
	viper.SetDefault("Metadata.Description", "Crunchy Data Feature Server for PostGIS")

	viper.SetDefault("Website.BasemapUrl", "")
}

// Config for system
type Config struct {
	Server   Server
	Paging   Paging
	Metadata Metadata
	Database Database
	Website  Website
}

// Server config
type Server struct {
	HttpHost                 string
	HttpPort                 int
	HttpsPort                int
	TlsServerCertificateFile string
	TlsServerPrivateKeyFile  string
	UrlBase                  string
	BasePath                 string
	CORSOrigins              string
	Debug                    bool
	AssetsPath               string
	ReadTimeoutSec           int
	WriteTimeoutSec          int
	TransformFunctions       []string
}

// Paging config
type Paging struct {
	LimitDefault int
	LimitMax     int
}

// Database config
type Database struct {
	DbConnection          string
	DbPoolMaxConnLifeTime string
	DbPoolMaxConns        int
	TableIncludes         []string
	TableExcludes         []string
	FunctionIncludes      []string
}

// Metadata config
type Metadata struct {
	Title       string //`mapstructure:"METADATA_TITLE"`
	Description string
}

type Website struct {
	BasemapUrl string
}

// IsHTTPSEnabled tests whether HTTPS is enabled
func (conf *Config) IsTLSEnabled() bool {
	return conf.Server.TlsServerCertificateFile != "" && conf.Server.TlsServerPrivateKeyFile != ""
}

// InitConfig initializes the configuration from the config file
func InitConfig(configFilename string, isDebug bool) {
	// --- defaults
	setDefaultConfig()

	if isDebug {
		viper.Set("Debug", true)
	}

	viper.SetEnvPrefix(AppConfig.EnvPrefix)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	isExplictConfigFile := configFilename != ""
	confFile := AppConfig.Name + ".toml"
	if configFilename != "" {
		viper.SetConfigFile(configFilename)
	} else {
		viper.SetConfigName(confFile)
		viper.SetConfigType("toml")
		viper.AddConfigPath("./config")
		viper.AddConfigPath("/config")
		viper.AddConfigPath("/etc")
	}
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {
		_, isConfigFileNotFound := err.(viper.ConfigFileNotFoundError)
		errrConfRead := fmt.Errorf("fatal error reading config file: %s", err)
		isUseDefaultConfig := isConfigFileNotFound && !isExplictConfigFile
		if isUseDefaultConfig {
			log.Debug(errrConfRead)
		} else {
			log.Fatal(errrConfRead)
		}
	}

	log.Infof("Using config file: %s", viper.ConfigFileUsed())
	errUnM := viper.Unmarshal(&Configuration)
	if errUnM != nil {
		log.Fatal(fmt.Errorf("fatal error decoding config file: %v", errUnM))
	}

	// Read environment variable database configuration
	// It takes precedence over config file (if any)
	// A blank value is ignored
	dbconnSrc := "config file"
	if dbURL := os.Getenv(AppConfig.EnvDBURL); dbURL != "" {
		Configuration.Database.DbConnection = dbURL
		dbconnSrc = "environment variable " + AppConfig.EnvDBURL
	}
	log.Infof("Using database connection info from %v", dbconnSrc)

	// sanitize the configuration
	Configuration.Server.BasePath = strings.TrimRight(Configuration.Server.BasePath, "/")
}

func DumpConfig() {
	log.Debugf("--- Configuration ---")
	//fmt.Printf("Viper: %v\n", viper.AllSettings())
	//fmt.Printf("Config: %v\n", Configuration)
	var basemapURL = Configuration.Website.BasemapUrl
	if basemapURL == "" {
		basemapURL = "*** NO URL PROVIDED ***"
	}
	log.Debugf("  BasemapUrl = %v", basemapURL)
	log.Debugf("  TableIncludes = %v", Configuration.Database.TableIncludes)
	log.Debugf("  TableExcludes = %v", Configuration.Database.TableExcludes)
	log.Debugf("  FunctionIncludes = %v", Configuration.Database.FunctionIncludes)
	log.Debugf("  TransformFunctions = %v", Configuration.Server.TransformFunctions)
}
