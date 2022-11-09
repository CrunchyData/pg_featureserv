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

func (cache CacheNaive) ContainsWeakEtag(strongEtag string) (bool, error) {
	weakEtagValue, err := api.EtagToWeakEtag(strongEtag)
	if err != nil {
		return false, err
	}
	mutex.Lock()
	_, present := cache.entries[weakEtagValue]
	mutex.Unlock()
	return present, nil
}

func (cache CacheNaive) AddWeakEtag(weakEtag string, etag interface{}) (bool, error) {
	mutex.Lock()
	cache.entries[weakEtag] = etag
	mutex.Unlock()
	return true, nil
}

func (cache CacheNaive) RemoveWeakEtag(weakEtag string) (bool, error) {
	mutex.Lock()
	delete(cache.entries, weakEtag)
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
