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
	"fmt"
	"testing"

	"github.com/CrunchyData/pg_featureserv/internal/data"
	util "github.com/CrunchyData/pg_featureserv/internal/utiltest"
)

func (t *CacheTests) TestRedisInvalidAddress() {
	t.Test.Run("TestRedisInvalidAddress", func(t *testing.T) {

		cache := data.CacheRedis{}
		err := cache.Init("invalid_addr", "")
		util.Assert(t, err != nil, "Error in CacheRedis initialization expected")
	})
}

func (t *CacheTests) TestRedisValidAddress() {
	url := t.RedisUrl
	t.Test.Run("TestRedisValidAddress", func(t *testing.T) {

		cache := data.CacheRedis{}
		err := cache.Init(url, "")
		util.Assert(t, err == nil, "No error in CacheRedis initialization expected")

		util.Assert(t, cache.String() == "Redis Cache running on "+url, "Invalid CacheRedis string")
	})
}

func (t *CacheTests) TestRedisContainsWeakEtag() {
	url := t.RedisUrl
	t.Test.Run("TestRedisContainsWeakEtag", func(t *testing.T) {

		cache := data.CacheRedis{}
		err := cache.Init(url, "")
		util.Assert(t, err == nil, "No error in CacheRedis initialization expected")

		// Test invalid etag use
		_, err = cache.ContainsWeakEtag("invalid_val")
		util.Assert(t, err != nil, "Invalid etag used for RedisCache, expecting failure")

		validWeakEtag := "collection"
		wValidWeakEtag := "W/" + validWeakEtag

		// Test valid etag but not available
		res, err := cache.ContainsWeakEtag(wValidWeakEtag)
		util.Assert(t, err == nil, "No error expected with valid Etag use")
		util.Assert(t, res == false, wValidWeakEtag+" Etag should not be available in Redis cache")

		// Test contains after add
		res, err = cache.AddWeakEtag(validWeakEtag, "value")
		util.Assert(t, err == nil, "No error expected with valid Etag use")
		util.Assert(t, res, "Result should be true when adding a valid weak Etag")

		res, err = cache.ContainsWeakEtag(wValidWeakEtag)
		util.Assert(t, err == nil, "No error expected with valid Etag use")
		util.Assert(t, res, wValidWeakEtag+" Etag should be available in Redis cache")

		// Test etag not available after remove
		res, err = cache.RemoveWeakEtag(validWeakEtag)
		util.Assert(t, err == nil, "No error expected with valid Etag use")
		util.Assert(t, res, validWeakEtag+" Result should be true when removing an available weak etag")

		res, err = cache.ContainsWeakEtag(wValidWeakEtag)
		util.Assert(t, err == nil, "No error expected with valid Etag use")
		util.Assert(t, res == false, wValidWeakEtag+" Etag should not be available in Redis cache")
	})
}

func (t *CacheTests) TestRedisAddWeakEtag() {
	url := t.RedisUrl
	t.Test.Run("TestRedisAddWeakEtag", func(t *testing.T) {

		cache := data.CacheRedis{}
		err := cache.Init(url, "")
		util.Assert(t, err == nil, "No error in CacheRedis initialization expected")

		validWeakEtag := "collection"

		initialSize := cache.Size()

		//Test add a valid etag and check size update
		res, err := cache.AddWeakEtag(validWeakEtag, "value")
		util.Assert(t, err == nil, "No error expected with valid Etag use")
		util.Assert(t, res, "Result should be true when adding a valid weak Etag")

		new_size := cache.Size()
		util.Assert(t, new_size == initialSize+1, fmt.Sprintf("Invalid Redis cache size %d / expecting %d", new_size, initialSize+1))

		//Test remove a valid etag and check size update
		res, err = cache.RemoveWeakEtag(validWeakEtag)
		util.Assert(t, err == nil, "No error expected with valid Etag use")
		util.Assert(t, res, validWeakEtag+" Result should be true when removing an available weak etag")

		new_size = cache.Size()
		util.Assert(t, new_size == initialSize, fmt.Sprintf("Invalid Redis cache size %d / expecting %d", new_size, initialSize))
	})
}

func (t *CacheTests) TestRedisRemoveWeakEtag() {
	url := t.RedisUrl
	t.Test.Run("TestRedisRemoveWeakEtag", func(t *testing.T) {

		cache := data.CacheRedis{}
		err := cache.Init(url, "")
		util.Assert(t, err == nil, "No error in CacheRedis initialization expected")

		validWeakEtag := "collection"
		wValidWeakEtag := "W/" + validWeakEtag

		// Test remove a etag not available
		res, err := cache.RemoveWeakEtag(validWeakEtag)
		util.Assert(t, err == nil, "No error expected with valid Etag use")
		util.Assert(t, res == false, "Result should be false when removing a weak etag not available")

		// Add etag and test removal
		res, err = cache.AddWeakEtag(validWeakEtag, "value")
		util.Assert(t, err == nil, "No error expected with valid Etag use")
		util.Assert(t, res, "Result should be true when adding a valid weak Etag")

		res, err = cache.RemoveWeakEtag(validWeakEtag)
		util.Assert(t, err == nil, "No error expected with valid Etag use")
		util.Assert(t, res, validWeakEtag+" Result should be true when removing an available weak etag")

		res, err = cache.ContainsWeakEtag(wValidWeakEtag)
		util.Assert(t, err == nil, "No error expected with valid Etag use")
		util.Assert(t, res == false, wValidWeakEtag+" Etag should not be available in Redis cache")
	})
}
