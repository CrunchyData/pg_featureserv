package data

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
	"context"
	"encoding/json"
	"fmt"

	"github.com/CrunchyData/pg_featureserv/internal/api"
	"github.com/go-redis/redis/v8"
)

type CacheRedis struct {
	client *redis.Client
	ctx    context.Context
}

func (cache *CacheRedis) Init(addr string, password string) error {
	cache.client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       1, //TODO add redis db management
	})
	cache.ctx = context.Background()
	_, err := cache.client.Ping(cache.ctx).Result()
	if err != nil {
		return fmt.Errorf("redis connection error: %s", err.Error())
	}

	return nil
}

// returns the object if the weak etag (etag is string - strong or weak etag - or *api.WeakEtagData) is referenced into the cache
// returns nil otherwise
// an error will be returned if a malformed etag is detected
func (cache CacheRedis) GetWeakEtag(etag interface{}) (*api.WeakEtagData, error) {
	weakEtagValue, err := anyToEtag(cache, etag)
	if err != nil {
		return nil, err
	}

	var etagStr string
	etagStr, err = cache.client.Get(cache.ctx, weakEtagValue.CacheKey()).Result()
	// redis.Nil error means that the value is not available : https://redis.uptrace.dev/guide/go-redis.html#redis-nil
	if err == redis.Nil {
		return nil, nil
	} else {
		var out *api.WeakEtagData
		err = json.Unmarshal([]byte(etagStr), out)
		return out, err
	}
}

func (cache CacheRedis) ContainsEtag(etag interface{}) (bool, error) {
	weakEtagValue, err := anyToEtag(cache, etag)
	if err != nil {
		return false, err
	}

	_, err = cache.client.Get(cache.ctx, weakEtagValue.CacheKey()).Result()
	// redis.Nil error means that the value is not available : https://redis.uptrace.dev/guide/go-redis.html#redis-nil
	if err == nil {
		return true, nil
	} else if err == redis.Nil {
		return false, nil
	} else {
		return false, err
	}
}

func (cache CacheRedis) AddWeakEtag(etagKey string, etag *api.WeakEtagData) (bool, error) {
	err := cache.client.Set(cache.ctx, etagKey, *etag, 0).Err()
	return err == nil, err
}

func (cache CacheRedis) RemoveWeakEtag(etagKey string) (bool, error) {
	nb_deleted, err := cache.client.Del(cache.ctx, etagKey).Result()
	return nb_deleted > 0, err
}

func (cache CacheRedis) String() string {
	return fmt.Sprintf("Redis Cache running on %s", cache.client.Options().Addr)
}

func (cache CacheRedis) Type() string {
	return "CacheRedis"
}

func (cache CacheRedis) Size() int {
	size, err := cache.client.DBSize(cache.ctx).Result()
	if err == nil {
		return int(size)
	} else {
		return -1
	}
}

func (cache CacheRedis) Reset() (bool, error) {
	_, err := cache.client.FlushDB(cache.ctx).Result()
	if err == nil {
		return true, nil
	}
	return false, err
}
