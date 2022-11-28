package cache_test

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
	"os"
	"testing"

	log "github.com/sirupsen/logrus"
)

// ...
func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

// ...
type CacheTests struct {
	Test     *testing.T
	RedisUrl string
}

// ...
func TestRunnerCache(t *testing.T) {
	// initialisation avant l'execution des tests
	beforeRun()

	redisCacheUrl := os.Getenv("PGFS_CACHE_REDIS_URL")

	t.Run("REDIS", func(t *testing.T) {
		beforeEachRun()
		m := CacheTests{Test: t}
		m.RedisUrl = redisCacheUrl
		m.TestRedisInvalidAddress()
		m.TestRedisValidAddress()
		m.TestRedisContainsEtag()
		m.TestRedisAddWeakEtag()
		m.TestRedisRemoveWeakEtag()
		afterEachRun()
	})

	// nettoyage après execution des tests
	afterRun()
}

// Run before all tests
func beforeRun() {
	log.Debug("beforeRun")
	// some stuff...
}

// Run after all tests
func afterRun() {
	log.Debug("afterRun")
	// some stuff...
}

// Run before each test
func beforeEachRun() {
	log.Debug("beforeEachRun")
}

// Run after each test
func afterEachRun() {
	log.Debug("afterEachRun")
	// some stuff...
}
