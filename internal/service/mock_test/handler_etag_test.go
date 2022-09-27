package mock_test

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

 Date     : September 2022
 Authors  : Nicolas Revelant (nicolas dot revelant at ign dot fr)
*/

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/CrunchyData/pg_featureserv/internal/api"
	util "github.com/CrunchyData/pg_featureserv/internal/utiltest"
)

func (t *MockTests) TestApiDecodeStrongEtag() {
	t.Test.Run("TestApiDecodeStrongEtag", func(t *testing.T) {
		// Valid Etag
		path := "/etags/decodestrong/Im1vY2tfYi00MzI2LWpzb24tMzk1NzI3NTc0NCI="
		resp := hTest.DoRequestStatus(t, path, http.StatusOK)

		strongEtag, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		// Check strong ETag validity
		// {"collection":"mock_b","srid":4326,"format":"json","weaketag":"3957275744"}
		var etagContent api.StrongEtagData
		util.Assert(t, json.Unmarshal([]byte(strongEtag), &etagContent) == nil, "the returned etag has to be in json format")
		util.Assert(t, etagContent.Collection == "mock_b", "missing/wrong collection inside decoded etag")
		util.Assert(t, etagContent.Srid == 4326, "missing/wrong srid inside decoded etag")
		util.Assert(t, etagContent.Format == "json", "missing/wrong format inside decoded etag")
		util.Assert(t, etagContent.WeakEtag == "3957275744", "missing/wrong weak value inside decoded etag")
	})
}

func (t *MockTests) TestLastModifiedMock() {
	t.Test.Run("TestLastModifiedMock", func(t *testing.T) {
		path := "/collections/mock_a/items/1"

		resp := hTest.DoRequestMethodStatus(t, "GET", path, []byte(""), nil, http.StatusOK)
		lastModifiedDate := resp.Header().Get("Last-Modified")
		parsedDate, err := time.Parse(time.RFC3339, lastModifiedDate)

		util.Assert(t, err != nil, "unparseable date")
		util.Assert(t, reflect.TypeOf(parsedDate).String() == "time.Time", "not a http date")
	})
}

func (t *MockTests) TestGetFeatureNoHeaderCheckEtag() {
	t.Test.Run("TestGetFeatureNoHeaderCheckEtag", func(t *testing.T) {
		// - GET
		path := "/collections/mock_b/items/1"
		resp := hTest.DoRequestStatus(t, path, http.StatusOK)

		// Check strong ETag validity
		strongEtag := resp.Result().Header["Etag"][0]
		decodedString, _ := base64.StdEncoding.DecodeString(strongEtag)
		decodedStrongEtag := string(decodedString)
		decodedStrongEtag = strings.Replace(decodedStrongEtag, "\"", "", -1)
		etagElements := strings.Split(decodedStrongEtag, "-")
		util.Equals(t, 4, len(etagElements), "strong ETag has to contain 4 values")
		util.Equals(t, "mock_b-4326-json-3957275744", decodedStrongEtag, "wrong strong ETag value")
		// - Extract weak eTag from Strong eTag
		weakEtag := etagElements[3]
		util.Equals(t, 10, len(weakEtag), "wrong weak ETag string size")
		util.Equals(t, "3957275744", weakEtag, "wrong weak ETag string")
	})
}

func (t *MockTests) TestGetFeatureHeaderIfNoneMatchWeakEtag() {
	t.Test.Run("TestGetFeatureHeaderIfNoneMatchWeakEtag", func(t *testing.T) {
		path := "/collections/mock_b/items/1"
		resp := hTest.DoRequestMethodStatus(t, "GET", path, nil, nil, http.StatusOK)

		strongEtag := resp.Result().Header["Etag"][0]
		decodedString, _ := base64.StdEncoding.DecodeString(strongEtag)
		decodedStrongEtag := string(decodedString)
		// strong-etag => "<collection>-<srid>-<format>-<weakEtag>"
		weakEtag := "W/" + "\"" + strings.Split(decodedStrongEtag, "-")[3] + "\""
		var header = make(http.Header)
		header.Add("If-None-Match", weakEtag)
		hTest.DoRequestMethodStatus(t, "GET", path, nil, header, http.StatusNotModified)
	})
}

func (t *MockTests) TestGetFeatureHeaderIfNoneMatchMalformedEtag() {
	t.Test.Run("TestGetFeatureHeaderIfNoneMatchMalformedEtag", func(t *testing.T) {
		path := "/collections/mock_b/items/1"
		var header = make(http.Header)
		header.Add("If-None-Match", "\"aa-mock_b-4326-json-999999999\"")
		hTest.DoRequestMethodStatus(t, "GET", path, nil, header, http.StatusBadRequest)
	})
}

func (t *MockTests) TestGetFeatureHeaderIfNoneMatchWithETagInCache() {
	t.Test.Run("TestGetFeatureHeaderIfNoneMatchWithETagInCache", func(t *testing.T) {
		path := "/collections/mock_b/items/1"
		// First GET to prefetch weak etag
		resp := hTest.DoRequestMethodStatus(t, "GET", path, nil, nil, http.StatusOK)
		strongEtagFromServer := resp.Result().Header["Etag"][0]

		var header = make(http.Header)

		header.Add("If-None-Match", strongEtagFromServer)
		hTest.DoRequestMethodStatus(t, "GET", path, nil, header, http.StatusNotModified)
	})
}

func (t *MockTests) TestPutFeatureEtag() {
	t.Test.Run("TestPutFeatureEtag", func(t *testing.T) {
		// Not yet implemented
	})
}

func (t *MockTests) TestPatchFeatureEtag() {
	t.Test.Run("TestPatchFeatureEtag", func(t *testing.T) {
		// Not yet implemented
	})
}
