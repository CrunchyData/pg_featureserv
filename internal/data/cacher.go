package data

import (
	"reflect"

	"github.com/CrunchyData/pg_featureserv/internal/api"
)

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

type Cacher interface {

	// returns true if the weak etag (etag is string - strong or weak etag - or *api.WeakEtagData) is referenced into the cache
	// returns false otherwise
	// an error will be returned if a malformed etag is detected
	ContainsEtag(etag interface{}) (bool, error)

	// returns the object if the weak etag (etag is string - strong or weak etag - or *api.WeakEtagData) is referenced into the cache
	// returns nil otherwise
	// an error will be returned if a malformed etag is detected
	GetWeakEtag(etag interface{}) (*api.WeakEtagData, error)

	// adds the weak etag string into the cache and returns true if successful
	// returns false if error occurs during the operation
	AddWeakEtag(etagKey string, etagData *api.WeakEtagData) (bool, error)

	// removes the weak etag string from the cache and returns true if successful
	// returns false if error occurs during the operation
	RemoveWeakEtag(weakEtag string) (bool, error)

	// Stringer: returns a string representation of the cache for dev purpose
	String() string

	// returns the cache name. May be a duplicate of `reflect.TypeOf(cache).Name()``
	Type() string

	// returns approx cache size
	Size() int

	// clean all cache content
	Reset() (bool, error)
}

// IsOneEtagInCache checks if the weak value of at least one of the etags provided is present into the cache
// Returns true at the first listed etag present into the cache, false otherwise
// -> error != nil if a malformed etag is detected (wrong encoding, bad format.)
// -> The provided etags have to be in their strong form and Base64 encoded
func IsOneEtagInCache(cache Cacher, etagsList []string) (bool, error) {
	for _, etag := range etagsList {
		found, err := cache.ContainsEtag(etag)
		if err != nil {
			return false, err
		}
		if found {
			return true, nil
		}
	}
	return false, nil
}

// extracts etag from string (weak or strong etag received by http client) or from WeakEtagData or StrongEtagData
func anyToEtag(cache Cacher, etag interface{}) (*api.WeakEtagData, error) {
	var weakEtagValue *api.WeakEtagData = nil
	var err error

	typeE := reflect.TypeOf(etag).String()
	if typeE == "string" { // from weak or strong etag string (from http client)
		etagStr := etag.(string)
		weakEtagValue, err = api.EtagStrToObject(etagStr)
		if err != nil {
			return nil, err
		}

	} else if typeE == "*api.WeakEtagData" {
		weakEtagValue = etag.(*api.WeakEtagData)

	} else if typeE == "*api.StrongEtagData" {
		strongEtagValue := etag.(*api.StrongEtagData)
		weakEtagValue = strongEtagValue.WeakEtagData
	}

	return weakEtagValue, nil
}
