package db_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"testing"

	"github.com/CrunchyData/pg_featureserv/internal/api"
	util "github.com/CrunchyData/pg_featureserv/internal/utiltest"
	"github.com/go-http-utils/headers"
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

 Date     : November 2022
 Authors  : Nicolas Revelant (nicolas dot revelant at ign dot fr)

*/

func (t *DbTests) TestGetFeatureIfNoneMatchStaValueWithEtagPresentInCacheDb() {
	t.Test.Run("TestGetFeatureIfNoneMatchStaValueWithEtagPresentInCacheDb", func(t *testing.T) {

		getFeatureEtag(t, "mock_a", 1) // Prefetching an etag in cache

		var header = make(http.Header)
		header.Add(headers.IfNoneMatch, "*")
		hTest.DoRequestMethodStatus(t, "GET", "/collections/mock_a/items/1", []byte(""), header, http.StatusNotModified)
	})
}

func (t *DbTests) TestGetFeatureIfNoneMatchStaValueWithNoEtagDb() {
	t.Test.Run("TestGetFeatureIfNoneMatchStaValueWithNoEtagDb", func(t *testing.T) {

		purgeCacheFromAllEtags(t)

		var header = make(http.Header)
		header.Add(headers.IfNoneMatch, "*")
		hTest.DoRequestMethodStatus(t, "GET", "/collections/mock_a/items/1", []byte(""), header, http.StatusOK)
	})
}

func (t *DbTests) TestGetFeatureIfNonMatchAfterReplaceDb() {
	t.Test.Run("TestGetFeatureIfNonMatchAfterReplaceDb", func(t *testing.T) {

		strongEtagFromServer := getFeatureEtag(t, "mock_a", 1)

		// If-None-Match before replace
		var header = make(http.Header)
		header.Add(headers.IfNoneMatch, strongEtagFromServer)
		hTest.DoRequestMethodStatus(t, "GET", "/collections/mock_a/items/1", []byte(""), header, http.StatusNotModified)

		// Replace
		var headerPut = make(http.Header)
		headerPut.Add(headers.Accept, api.ContentTypeSchemaPatchJSON)
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
		hTest.DoRequestMethodStatus(t, "PUT", "/collections/mock_a/items/1", []byte(jsonStr), headerPut, http.StatusNoContent)

		// If-None-Match after replace
		strongEtagAfterPut := getFeatureEtag(t, "mock_a", 1)
		util.Assert(t, strongEtagFromServer != strongEtagAfterPut, "strong etag value is still the same after replace!")
	})
}

func (t *DbTests) TestUpdateFeatureIfNoneMatchWithEtagPresentInCacheDb() {
	t.Test.Run("TestUpdateFeatureIfNoneMatchWithEtagPresentInCache", func(t *testing.T) {
		// If-None-Match header + PATCH with an etag detected in cache
		// -> 412 Precondition failed

		encodedEtagFromServer := getFeatureEtag(t, "mock_a", 1)

		var header = make(http.Header)
		header.Add(headers.Accept, api.ContentTypeSchemaPatchJSON)
		header.Add(headers.IfNoneMatch, encodedEtagFromServer)

		path := "/collections/mock_a/items/1"
		// checkFeatureModification(t , requestedCoords , method, path, header, expectedStatus, expectedCoords)
		checkFeatureModification(t, [2]int{-500, 800}, "PATCH", path, header, http.StatusPreconditionFailed, [2]int{-120, 40})

	})
}

func (t *DbTests) TestUpdateFeatureIfNoneMatchWithEtagNotDetectedDb() {
	t.Test.Run("TestUpdateFeatureIfNoneMatchWithEtagNotDetected", func(t *testing.T) {
		// If-None-Match header + PATCH with Etag not present in Cache
		// -> 204 "No Content"

		encodedEtagFromServer := getFeatureEtag(t, "mock_a", 1)

		purgeCacheFromAllEtags(t)

		var header = make(http.Header)
		header.Add(headers.Accept, api.ContentTypeSchemaPatchJSON)
		header.Add(headers.IfNoneMatch, encodedEtagFromServer)

		path := "/collections/mock_a/items/1"
		// checkFeatureModification(t , requestedCoords , method, path, header, expectedStatus, expectedCoords)
		checkFeatureModification(t, [2]int{-500, 800}, "PATCH", path, header, http.StatusNoContent, [2]int{-500, 800})
	})
}

func (t *DbTests) TestUpdateFeatureIfNoneMatchStarValueWithExistingRepresentationInCacheDb() {
	t.Test.Run("TestUpdateFeatureIfNoneMatchStarValueWithExistingRepresentationInCache", func(t *testing.T) {
		// If-None-Match header + PATCH with an etag detected in cache
		// -> 412 Precondition failed

		getFeatureEtag(t, "mock_a", 1) // Prefetching an etag in cache

		var header = make(http.Header)
		header.Add(headers.Accept, api.ContentTypeSchemaPatchJSON)
		header.Add(headers.IfNoneMatch, "*")

		path := "/collections/mock_a/items/1"
		// checkFeatureModification(t , requestedCoords , method, path, header, expectedStatus, expectedCoords)
		checkFeatureModification(t, [2]int{-9000, 8000}, "PATCH", path, header, http.StatusPreconditionFailed, [2]int{-500, 800})
	})
}

func (t *DbTests) TestUpdateFeatureIfNoneMatchStarValueWithNoRepresentationInCacheDb() {
	t.Test.Run("TestUpdateFeatureIfNoneMatchStarValueWithNoRepresentationInCache", func(t *testing.T) {
		// IfNoneMatch "*" PATCH without any Etag present in cache
		// -> 204 No content

		purgeCacheFromAllEtags(t)

		var header = make(http.Header)
		header.Add(headers.Accept, api.ContentTypeSchemaPatchJSON)
		header.Add(headers.IfNoneMatch, "*")

		path := "/collections/mock_a/items/1"
		// checkFeatureModification(t , requestedCoords , method, path, header, expectedStatus, expectedCoords)
		checkFeatureModification(t, [2]int{-9000, 8000}, "PATCH", path, header, http.StatusNoContent, [2]int{-9000, 8000})
	})
}

func (t *DbTests) TestReplaceFeatureIfNoneMatchWithEtagPresentInCacheDb() {
	t.Test.Run("TestReplaceFeatureIfNoneMatchWithEtagPresentInCache", func(t *testing.T) {
		// If-None-Match header + PUT with an etag detected in cache
		// -> 412 Precondition failed

		encodedEtagFromServer := getFeatureEtag(t, "mock_a", 1)

		var header = make(http.Header)
		header.Add(headers.Accept, api.ContentTypeSchemaPatchJSON)
		header.Add(headers.IfNoneMatch, encodedEtagFromServer)

		path := "/collections/mock_a/items/1"
		// checkFeatureModification(t , requestedCoords , method, path, header, expectedStatus, expectedCoords)
		checkFeatureModification(t, [2]int{-9002, 8002}, "PUT", path, header, http.StatusPreconditionFailed, [2]int{-9000, 8000})
	})
}

func (t *DbTests) TestReplaceFeatureIfNoneMatchWithEtagNotDetectedDb() {
	t.Test.Run("TestReplaceFeatureIfNoneMatchWithEtagNotDetected", func(t *testing.T) {
		// If-None-Match header + PUT with Etag not present in Cache
		// -> 204 "No Content"

		encodedEtagFromServer := getFeatureEtag(t, "mock_a", 1)

		purgeCacheFromAllEtags(t)

		var header = make(http.Header)
		header.Add(headers.Accept, api.ContentTypeSchemaPatchJSON)
		header.Add(headers.IfNoneMatch, encodedEtagFromServer)

		path := "/collections/mock_a/items/1"
		// checkFeatureModification(t , requestedCoords , method, path, header, expectedStatus, expectedCoords)
		checkFeatureModification(t, [2]int{-501, 801}, "PUT", path, header, http.StatusNoContent, [2]int{-501, 801})
	})

}

func (t *DbTests) TestReplaceFeatureIfNoneMatchStarValueWithNoRepresentationInCacheDb() {
	t.Test.Run("TestReplaceFeatureIfNoneMatchStarValueWithNoRepresentationInCache", func(t *testing.T) {
		// IfNoneMatch "*" PUT without any Etag present in cache
		// -> 204 No content

		purgeCacheFromAllEtags(t)

		var header = make(http.Header)
		header.Add(headers.Accept, api.ContentTypeSchemaPatchJSON)
		header.Add(headers.IfNoneMatch, "*")

		path := "/collections/mock_a/items/1"
		// checkFeatureModification(t , requestedCoords , method, path, header, expectedStatus, expectedCoords)
		checkFeatureModification(t, [2]int{-9009, 8009}, "PUT", path, header, http.StatusNoContent, [2]int{-9009, 8009})
	})
}

func (t *DbTests) TestReplaceFeatureIfNoneMatchStarValueWithExistingRepresentationInCacheDb() {
	t.Test.Run("TestReplaceFeatureIfNoneMatchStarValueWithExistingRepresentationInCache", func(t *testing.T) {
		// If-None-Match header + PUT with an etag detected in cache
		// -> 412 Precondition failed

		getFeatureEtag(t, "mock_a", 1) // Prefetching an etag in cache

		var header = make(http.Header)
		header.Add(headers.Accept, api.ContentTypeSchemaPatchJSON)
		header.Add(headers.IfNoneMatch, "*")

		path := "/collections/mock_a/items/1"
		// checkFeatureModification(t , requestedCoords , method, path, header, expectedStatus, expectedCoords)
		checkFeatureModification(t, [2]int{-9000, 8000}, "PUT", path, header, http.StatusPreconditionFailed, [2]int{-9009, 8009})

	})
}

// cleans the catalog cache by instantiating a new cache structure with a void map entries
func purgeCacheFromAllEtags(t *testing.T) {
	purgePath := "/etags/purge"
	hTest.DoRequestStatus(t, purgePath, http.StatusOK)
}

// send a GET request on the wanted feature /collections/{collectionName}/items/{feature_id}
// and returns the etag obtained into the response
func getFeatureEtag(t *testing.T, collectionName string, featureId int) string {
	path := "/collections/" + collectionName + "/items/" + strconv.Itoa(featureId)
	resp := hTest.DoRequestMethodStatus(t, "GET", path, []byte(""), nil, http.StatusOK)
	return resp.Header().Get("Etag")
}

// send a GET request on the specified feature id and collection
// and returns a JSON formatted version of the feature data
func getFeatureAsJson(t *testing.T, collectionName string, featureId int) map[string]interface{} {
	feature := checkItem(t, "mock_a", 1)
	var jsonData map[string]interface{}
	errUnMarsh := json.Unmarshal(feature, &jsonData)
	util.Assert(t, errUnMarsh == nil, fmt.Sprintf("%v", errUnMarsh))
	return jsonData
}

func checkFeatureModification(t *testing.T, requestedCoords [2]int, method string, path string, header http.Header, expectedStatus int, expectedCoords [2]int) {

	jsonStr := fmt.Sprintf(`{
		"type": "Feature",
		"id": "1",
		"geometry": {
			"type": "Point",
			"coordinates": [
			%d,
			%d
			]
		},
		"properties": {
			"prop_a": "propA...",
			"prop_b": 1,
			"prop_c": "propC...",
			"prop_d": 1
		}
	}`,
		requestedCoords[0],
		requestedCoords[1],
	)

	_ = hTest.DoRequestMethodStatus(t, method, path, []byte(jsonStr), header, expectedStatus)

	// checking the feature data are still the same
	jsonData := getFeatureAsJson(t, "mock_a", 1)

	geom := jsonData["geometry"].(map[string]interface{})
	util.Equals(t, "Point", geom["type"].(string), "feature Type")
	coordinate := geom["coordinates"].([]interface{})
	if requestedCoords != expectedCoords {
		util.Equals(t, expectedCoords[0], int(coordinate[0].(float64)), "feature latitude not updated!")
		util.Equals(t, expectedCoords[1], int(coordinate[1].(float64)), "feature longitude not updated!")
	} else {
		util.Equals(t, expectedCoords[0], int(coordinate[0].(float64)), "feature latitude has been modified!")
		util.Equals(t, expectedCoords[1], int(coordinate[1].(float64)), "feature longitude has been modified!")
	}

}
