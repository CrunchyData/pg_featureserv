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

 Date     : October 2022
 Authors  : Nicolas Revelant (nicolas dot revelant at ign dot fr)
*/

var setVersion string = "1.3"

// AppConfiguration is the set of global application configuration constants.
type AppConfiguration struct {
	// AppName name of the software
	Name      string
	EnvPrefix string
	// AppVersion version number of the software
	Version string
	// Database URL
	EnvDBURL string
	// Cache configuration
	EnvCacheType          string
	EnvCacheNaiveSize     string
	EnvCacheRedisUrl      string
	EnvCacheRedisPassword string
}

var AppConfig = AppConfiguration{
	Name:                  "pg_featureserv",
	Version:               setVersion,
	EnvPrefix:             "PGFS",
	EnvDBURL:              "DATABASE_URL",
	EnvCacheType:          "PGFS_CACHE_TYPE",
	EnvCacheNaiveSize:     "PGFS_CACHE_NAIVE_SIZE",
	EnvCacheRedisUrl:      "PGFS_CACHE_REDIS_URL",
	EnvCacheRedisPassword: "PGFS_CACHE_REDIS_PASSWORD",
}
