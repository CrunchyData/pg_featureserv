package conf

/*
 Copyright 2022 Crunchy Data Solutions, Inc.
 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at
      http://www.apache.org/licenses/LICENSE-2.0
 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.

 Date     : November 2022
 Authors  : Jean-Marie KERLOCH (jean-marie dot kerloch at oslandia dot com)
*/

import (
	"fmt"
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
)

// Cache config
type Cache struct {
	Type  string
	Naive NaiveCacheConfig
	Redis RedisCacheConfig
}

// Init Cache configuration from environnement variables
func (cache *Cache) InitFromEnvVariables() {
	origin := "config file"
	cacheTypeInput := os.Getenv(AppConfig.EnvCacheType)
	if cacheTypeInput != "" {
		origin = "environment variable"
		cache.Type = cacheTypeInput
	}
	if cache.Type != "Disabled" && cache.Type != "Naive" && cache.Type != "Redis" {
		log.Fatal(fmt.Errorf("Invalid cache type: Disabled, Naive and Redis are supported. %v defined", cache.Type))
	}

	log.Infof("Using cache type %s set from %s", Configuration.Cache.Type, origin)

	if Configuration.Cache.Type == "Naive" {
		cache.Naive.InitFromEnvVariables()
	} else if Configuration.Cache.Type == "Redis" {
		cache.Redis.InitFromEnvVariables()
	}
}

func (cache *Cache) DumpConfig() {
	log.Debug("  --- Cache ---")
	log.Debugf("  Type = %v", Configuration.Cache.Type)
	if Configuration.Cache.Type == "Naive" {
		log.Debugf("  NaiveCache.MapSize = %v", Configuration.Cache.Naive.MapSize)
	}
	if Configuration.Cache.Type == "Redis" {
		log.Debugf("  RedisCache.Url = %v", Configuration.Cache.Redis.Url)
	}
}

// NaiveCache config
type NaiveCacheConfig struct {
	MapSize int
}

// Init Naive cache configuration from environnement variables
func (cache *NaiveCacheConfig) InitFromEnvVariables() {
	origin := "config file"
	if cacheSizeInput, err := strconv.Atoi(os.Getenv(AppConfig.EnvCacheNaiveSize)); cacheSizeInput != 0 {
		if err != nil {
			log.Fatal(fmt.Errorf("fatal error reading env variable: %v", err))
		}
		cache.MapSize = cacheSizeInput
		origin = "environment variable"
	}
	log.Infof("Using etag cache size set from %s (%d entries)", origin, cache.MapSize)
}

// RedisCache config
type RedisCacheConfig struct {
	Url      string
	Password string
}

// Init Redis cache configuration from environnement variables
func (cache *RedisCacheConfig) InitFromEnvVariables() {
	origin := "config file"
	urlInput := os.Getenv(AppConfig.EnvCacheRedisUrl)
	if urlInput != "" {
		origin = "environment variable"
		cache.Url = urlInput
	}
	log.Infof("Using Redis cache url %s set from %s", cache.Url, origin)

	origin = "config file"
	passwordInput := os.Getenv(AppConfig.EnvCacheRedisPassword)
	if passwordInput != "" {
		origin = "environment variable"
		cache.Password = passwordInput
	}
	log.Infof("Using Redis cache password set from %s", origin)
}
