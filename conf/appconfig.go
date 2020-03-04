package conf

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

// AppConfiguration is the set of global application configuration constants.
type AppConfiguration struct {
	// AppName name of the software
	Name string
	// AppVersion version number of the software
	Version  string
	EnvDBURL string
}

var AppConfig = AppConfiguration{
	Name:     "pg_featureserv",
	Version:  "0.1",
	EnvDBURL: "DATABASE_URL",
}
