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

 Date     : October 2022
 Authors  : Nicolas Revelant (nicolas dot revelant at ign dot fr)
*/

import (
	"bytes"
	"fmt"
	"sync"

	"github.com/CrunchyData/pg_featureserv/internal/api"
)

type CacheNaive struct {
	entries map[string]interface{}
}

var mutex = &sync.Mutex{} // allows concurrent accesses to the cache map

func (cache CacheNaive) GetWeakEtag(etag interface{}) (*api.WeakEtagData, error) {
	weakEtagValue, err := anyToEtag(cache, etag)
	if err != nil {
		return nil, err
	}

	mutex.Lock()
	var out interface{}
	out, present := cache.entries[weakEtagValue.CacheKey()]
	mutex.Unlock()
	if present {
		return out.(*api.WeakEtagData), nil
	}
	return nil, nil
}
func (cache CacheNaive) ContainsEtag(etag interface{}) (bool, error) {
	weakEtagValue, err := anyToEtag(cache, etag)
	if err != nil {
		return false, err
	}

	mutex.Lock()
	_, present := cache.entries[weakEtagValue.CacheKey()]
	mutex.Unlock()
	return present, nil
}

func (cache CacheNaive) AddWeakEtag(etagKey string, etagData *api.WeakEtagData) (bool, error) {
	mutex.Lock()
	cache.entries[etagKey] = etagData
	mutex.Unlock()
	return true, nil
}

func (cache CacheNaive) RemoveWeakEtag(etagKey string) (bool, error) {
	mutex.Lock()
	delete(cache.entries, etagKey)
	mutex.Unlock()
	return true, nil
}

func (cache CacheNaive) String() string {
	b := new(bytes.Buffer)
	for key, value := range cache.entries {
		fmt.Fprintf(b, "%s=\"%s\"\n", key, value)
	}
	return b.String()
}

func (cache CacheNaive) Type() string {
	return "CacheNaive"
}

func (cache CacheNaive) Size() int {
	return len(cache.entries)
}

func (cache CacheNaive) Reset() (bool, error) {
	mutex.Lock()
	for k := range cache.entries {
		delete(cache.entries, k)
	}
	mutex.Unlock()
	return true, nil
}
