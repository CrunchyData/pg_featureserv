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

// Configuration for system
var Configuration Config

func init() {
	Configuration = Config{
		Server: Server{
			BindHost: "",
			BindPort: 10000,
		},
		Metadata: Metadata{
			Title:       "pg-featureserv Demo",
			Description: "Demo of Crunchy Data PostGIS Feature Server",
		},
		Database: Database{
			ConnectString: "",
		},
	}
}

// Config for system
type Config struct {
	Server   Server
	Metadata Metadata
	Database Database
}

// Server config
type Server struct {
	BindHost string `toml:"bind_host"`
	BindPort int    `toml:"bind_port"`
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
