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
	})
	cache.ctx = context.Background()
	_, err := cache.client.Ping(cache.ctx).Result()
	if err != nil {
		return fmt.Errorf("redis connection error: %s", err.Error())
	}

	return nil
}

func (cache CacheRedis) ContainsWeakEtag(strongEtag string) (bool, error) {
	weakEtagValue, err := api.EtagToWeakEtag(strongEtag)
	if err != nil {
		return false, err
	}

	_, err = cache.client.Get(cache.ctx, weakEtagValue).Result()
	// redis.Nil error means that the value is not available : https://redis.uptrace.dev/guide/go-redis.html#redis-nil
	if err == redis.Nil {
		return false, nil
	} else {
		return err == nil, err
	}
}

func (cache CacheRedis) AddWeakEtag(weakEtag string, etag interface{}) (bool, error) {
	err := cache.client.Set(cache.ctx, weakEtag, etag, 0).Err()
	return err == nil, err
}

func (cache CacheRedis) RemoveWeakEtag(weakEtag string) (bool, error) {
	nb_deleted, err := cache.client.Del(cache.ctx, weakEtag).Result()
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
