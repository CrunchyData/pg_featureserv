package db_test

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
	"encoding/base64"
	"net/http"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/CrunchyData/pg_featureserv/internal/api"
	"github.com/CrunchyData/pg_featureserv/internal/data"
	util "github.com/CrunchyData/pg_featureserv/internal/utiltest"
)

func (t *DbTests) TestCacheActivationDb() {
	t.Test.Run("TestCacheActivationDb", func(t *testing.T) {

		var cache data.Cacher

		// testing IsCacheActive() function

		cache = data.CacheDisabled{}
		util.Assert(t, reflect.TypeOf(cache).Name() == "CacheDisabled", "wrong type for disabled cache")
		util.Assert(t, cache.IsCacheActive() == false, "CacheDisabled should be not active")

		cache = data.CacheNaive{}
		util.Assert(t, reflect.TypeOf(cache).Name() == "CacheNaive", "active cache should have CacheNaive type")
		util.Assert(t, cache.IsCacheActive() == true, "CacheNaive should be active")

	})
}

func (t *DbTests) TestLastModifiedDb() {
	t.Test.Run("TestLastModifiedDb", func(t *testing.T) {
		path := "/collections/mock_a/items/1"

		resp := hTest.DoRequestMethodStatus(t, "GET", path, []byte(""), nil, http.StatusOK)
		lastModifiedDate := resp.Header().Get("Last-Modified")
		parsedDate, err := time.Parse(time.RFC3339, lastModifiedDate)

		util.Assert(t, err != nil, "unparseable date")
		util.Assert(t, reflect.TypeOf(parsedDate).String() == "time.Time", "not a http date")
	})
}

func (t *DbTests) TestEtagDb() {
	t.Test.Run("TestEtagDb", func(t *testing.T) {
		path := "/collections/mock_a/items/1"

		resp := hTest.DoRequestMethodStatus(t, "GET", path, []byte(""), nil, http.StatusOK)
		etagFromServer := resp.Header().Get("Etag")

		// Verifying the format of the strong etag received from server
		decodedStrongEtag, err := api.DecodeStrongEtag(etagFromServer)
		util.Assert(t, err == nil, "wrong strong etag form")

		// Verifying the content from the strong etag
		util.Assert(t, decodedStrongEtag.Collection == "mock_a", "wrong collection name")
		util.Assert(t, decodedStrongEtag.Srid == 4326, "wrong SRID value")
		util.Assert(t, decodedStrongEtag.Format == "json", "wrong format")
		util.Assert(t, reflect.TypeOf(decodedStrongEtag.WeakEtag).String() == "string", "weak etag is not a string value")
	})
}

func (t *DbTests) TestWeakEtagStableOnRequestsDb() {
	t.Test.Run("TestWeakEtagStableOnRequestsDb", func(t *testing.T) {
		path := "/collections/mock_b/items/1"
		var headerJson = make(http.Header)
		headerJson.Add("Accept", api.ContentTypeJSON)

		resp := hTest.DoRequestMethodStatus(t, "GET", path, []byte(""), headerJson, http.StatusOK)
		encodedStrongEtag1 := resp.Header().Get("Etag")

		strongEtag1, err := api.DecodeStrongEtag(encodedStrongEtag1)
		util.Assert(t, err == nil, "wrong strong etag form")

		resp2 := hTest.DoRequestMethodStatus(t, "GET", path, []byte(""), nil, http.StatusOK)
		encodeStrongEtag2 := resp2.Header().Get("Etag")

		strongEtag2, err2 := api.DecodeStrongEtag(encodeStrongEtag2)
		util.Assert(t, err2 == nil, "wrong strong etag form")

		util.Assert(t, strongEtag1.WeakEtag == strongEtag2.WeakEtag, "weak etag values are different for the same feature!")
	})
}

func (t *DbTests) TestEtagHeaderIfNoneMatchDb() {
	t.Test.Run("TestEtagHeaderIfNoneMatchDb", func(t *testing.T) {
		path := "/collections/mock_a/items/1"

		// first GET prefetches the etag into the server cache
		resp := hTest.DoRequestMethodStatus(t, "GET", path, []byte(""), nil, http.StatusOK)

		strongEtagFromServer := resp.Header().Get("ETag")

		var header = make(http.Header)
		header.Add("If-None-Match", strongEtagFromServer)
		hTest.DoRequestMethodStatus(t, "GET", path, []byte(""), header, http.StatusNotModified)
	})
}

func (t *DbTests) TestEtagHeaderIfNonMatchAfterReplaceDb() {
	t.Test.Run("TestEtagHeaderIfNonMatchAfterReplaceDb", func(t *testing.T) {
		path := "/collections/mock_a/items/1"
		resp := hTest.DoRequestMethodStatus(t, "GET", path, []byte(""), nil, http.StatusOK)
		strongEtagFromServer := resp.Header().Get("Etag")

		// If-None-Match before replace
		var header = make(http.Header)
		header.Add("If-None-Match", strongEtagFromServer)
		hTest.DoRequestMethodStatus(t, "GET", path, []byte(""), header, http.StatusNotModified)

		// Replace
		var headerPut = make(http.Header)
		headerPut.Add("Accept", api.ContentTypeSchemaPatchJSON)
		jsonStr := `{
			"type": "Feature",
			"id": "1",
			"geometry": {
				"type": "Point",
				"coordinates": [
				-120,
				40
				]
			},
			"properties": {
				"prop_a": "propA...",
				"prop_b": 1,
				"prop_c": "propC...",
				"prop_d": 1
			}
		}`
		hTest.DoRequestMethodStatus(t, "PUT", path, []byte(jsonStr), headerPut, http.StatusNoContent)

		// TODO : test will be OK once merge done with trigger

		// If-None-Match after replace
		// resp2 := hTest.DoRequestMethodStatus(t, "GET", path, []byte(""), header, http.StatusOK)
		// strongEtagAfterPut := resp2.Header().Get("Etag")

		// util.Assert(t, strongEtagFromServer != strongEtagAfterPut, "strong etag value is still the same after replace!")
	})
}

func (t *DbTests) TestEtagHeaderIfNonMatchMalformedEtagDb() {
	t.Test.Run("TestEtagHeaderIfNonMatchMalformedEtagDb", func(t *testing.T) {
		path := "/collections/mock_a/items/1"

		var header = make(http.Header)
		header.Add("If-None-Match", "\"unknown_etag\"")
		hTest.DoRequestMethodStatus(t, "GET", path, []byte(""), header, http.StatusBadRequest)

		var header2 = make(http.Header)
		header2.Add("If-None-Match", "\"mock_a-4326-json-812\"")
		hTest.DoRequestMethodStatus(t, "GET", path, []byte(""), header2, http.StatusBadRequest)

		var header3 = make(http.Header)
		header3.Add("If-None-Match", "mock_a-4326-json-812")
		hTest.DoRequestMethodStatus(t, "GET", path, []byte(""), header3, http.StatusBadRequest)
	})
}

func (t *DbTests) TestEtagHeaderIfNonMatchVariousEtagsDb() {
	t.Test.Run("TestEtagHeaderIfNonMatchVariousEtagsDb", func(t *testing.T) {
		path := "/collections/mock_a/items/1"
		resp := hTest.DoRequestMethodStatus(t, "GET", path, []byte(""), nil, http.StatusOK)
		encodedValidEtagFromServer := resp.Header().Get("Etag")

		wrongEtag := "mock_a-4326-json-99999"
		wrongEtag2 := "collection2-4326-html-99999"

		encodedWrongEtag1 := base64.StdEncoding.EncodeToString([]byte(wrongEtag))
		encodedWrongEtag2 := base64.StdEncoding.EncodeToString([]byte(wrongEtag2))

		// If-None-Match before replace
		var header = make(http.Header)
		etags := []string{encodedWrongEtag1, encodedWrongEtag2, encodedValidEtagFromServer}
		headerValue := strings.Join(etags, ",")
		header.Add("If-None-Match", headerValue)
		hTest.DoRequestMethodStatus(t, "GET", path, []byte(""), header, http.StatusNotModified)
	})
}

func (t *DbTests) TestEtagHeaderIfNonMatchWeakEtagDb() {
	t.Test.Run("TestEtagHeaderIfNonMatchWeakEtagDb", func(t *testing.T) {
		path := "/collections/mock_b/items/1"
		resp := hTest.DoRequestMethodStatus(t, "GET", path, nil, nil, http.StatusOK)

		strongEtagFromServer := resp.Header().Get("Etag")
		decodedStrongEtag, err := api.DecodeStrongEtag(strongEtagFromServer)
		util.Assert(t, err == nil, "wrong strong etag form")

		// strong-etag => "<collection>-<srid>-<format>-<weakEtag>"
		weakEtag := "W/" + "\"" + decodedStrongEtag.WeakEtag + "\""
		var header = make(http.Header)
		header.Add("If-None-Match", weakEtag)
		hTest.DoRequestMethodStatus(t, "GET", path, nil, header, http.StatusNotModified)
	})
}

func (t *DbTests) TestEtagHeaderIfMatchDb() {
	t.Test.Run("TestEtagHeaderIfMatchDb", func(t *testing.T) {
		// TODO
	})
}

func (t *DbTests) TestEtagReplaceFeatureDb() {
	t.Test.Run("TestEtagReplaceFeatureDb", func(t *testing.T) {
		path := "/collections/mock_b/items/1"
		var header = make(http.Header)
		header.Add("Accept", api.FormatJSON)

		jsonStr := `{
			"type": "Feature",
			"id": "1",
			"geometry": {
				"type": "Point",
				"coordinates": [
				-120,
				40
				]
			},
			"properties": {
				"prop_a": "propA...",
				"prop_b": 1,
				"prop_c": "propC...",
				"prop_d": 1
			}
		}`

		resp := hTest.DoRequestMethodStatus(t, "GET", path, []byte(""), header, http.StatusOK)
		encodedStrongEtagBeforePut := resp.Header().Get("Etag")

		strongEtagBeforePut, err := api.DecodeStrongEtag(encodedStrongEtagBeforePut)
		util.Assert(t, err == nil, "wrong strong etag form")

		// Replace
		var header2 = make(http.Header)
		header.Add("Accept", api.ContentTypeSchemaPatchJSON)
		hTest.DoRequestMethodStatus(t, "PUT", path, []byte(jsonStr), header2, http.StatusNoContent)

		resp2 := hTest.DoRequestMethodStatus(t, "GET", path, []byte(""), header, http.StatusOK)
		encodedStrongEtagAfterPut := resp2.Header().Get("Etag")

		strongEtagAfterPut, err := api.DecodeStrongEtag(encodedStrongEtagAfterPut)
		util.Assert(t, err == nil, "wrong strong etag form")

		util.Assert(t, strongEtagBeforePut != strongEtagAfterPut, "strong etag value is still the same after replace!")
		util.Assert(t, strongEtagBeforePut != strongEtagAfterPut, "weak etag value is still the same after replace!")
	})
}
