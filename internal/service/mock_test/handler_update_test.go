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
 Authors  : Jean-philippe Bazonnais (jean-philippe dot bazonnais at ign dot fr)
*/

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/CrunchyData/pg_featureserv/internal/api"
	util "github.com/CrunchyData/pg_featureserv/internal/utiltest"
	"github.com/getkin/kin-openapi/openapi3"
)

// checks swagger api contains method PATCH for updating a feaure from a specified collection
func (t *MockTests) TestApiContainsMethodPatchFeature() {
	t.Test.Run("TestApiContainsMethodPatchFeature", func(t *testing.T) {
		resp := hTest.DoRequest(t, "/api")
		body, _ := ioutil.ReadAll(resp.Body)

		var v openapi3.T
		err := json.Unmarshal(body, &v)
		util.Assert(t, err == nil, fmt.Sprintf("%v", err))

		util.Equals(t, "Provides access to a single feature identitfied by {featureId} from the specified collection", v.Paths.Find("/collections/{collectionId}/items/{featureId}").Description, "path present")
		util.Equals(t, "updateCollectionFeature", v.Paths.Find("/collections/{collectionId}/items/{featureId}").Patch.OperationID, "method PATCH present")
	})
}

func (t *MockTests) TestGetCollectionUpdateSchema() {
	t.Test.Run("TestGetCollectionUpdateSchema", func(t *testing.T) {
		path := "/collections/mock_a/schema?type=update"
		var header = make(http.Header)
		header.Add("Accept", api.ContentTypeSchemaJSON)

		resp := hTest.DoRequestMethodStatus(t, "GET", path, nil, header, http.StatusOK)
		body, _ := ioutil.ReadAll(resp.Body)

		var fis openapi3.Schema
		errUnMarsh := json.Unmarshal(body, &fis)
		util.Assert(t, errUnMarsh == nil, fmt.Sprintf("%v", errUnMarsh))

		util.Equals(t, "This dataset contains mock data about A (9 points)", fis.Description, "feature description")
		util.Equals(t, "GeoJSON Point", fis.Properties["geometry"].Value.Title, "feature geometry")
		util.Equals(t, 0, len(fis.Required), "no required field")
	})
}

func (t *MockTests) TestUpdateFeatureSuccess() {
	t.Test.Run("TestUpdateFeatureSuccess", func(t *testing.T) {
		path := "/collections/mock_a/items/1"
		var header = make(http.Header)
		header.Add("Accept", api.ContentTypeSchemaPatchJSON)

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

		_ = hTest.DoRequestMethodStatus(t, "PATCH", path, []byte(jsonStr), header, http.StatusNoContent)

		// check if it can be read
		feature := checkItem(t, 1)
		var jsonData map[string]interface{}
		errUnMarsh := json.Unmarshal(feature, &jsonData)
		util.Assert(t, errUnMarsh == nil, fmt.Sprintf("%v", errUnMarsh))

		util.Equals(t, "1", jsonData["id"].(string), "feature ID")
		util.Equals(t, "Feature", jsonData["type"].(string), "feature Type")
		props := jsonData["properties"].(map[string]interface{})
		util.Equals(t, "propA...", props["prop_a"].(string), "feature value a")
		util.Equals(t, 1, int(props["prop_b"].(float64)), "feature value b")
		util.Equals(t, "propC...", props["prop_c"].(string), "feature value c")
		util.Equals(t, 1, int(props["prop_d"].(float64)), "feature value d")
		geom := jsonData["geometry"].(map[string]interface{})
		util.Equals(t, "Point", geom["type"].(string), "feature Type")
		coordinate := geom["coordinates"].([]interface{})
		util.Equals(t, -120, int(coordinate[0].(float64)), "feature latitude")
		util.Equals(t, 40, int(coordinate[1].(float64)), "feature longitude")

	})
}

func (t *MockTests) TestUpdateFeaturePartialSuccess() {
	t.Test.Run("TestUpdateFeaturePartialSuccess", func(t *testing.T) {
		path := "/collections/mock_a/items/2"
		var header = make(http.Header)
		header.Add("Accept", api.ContentTypeSchemaPatchJSON)

		jsonStr := `{
			"type": "Feature",
			"id": "2",
			"geometry": {
				"type": "Point",
				"coordinates": [
				-120,
				40
				]
			},
			"properties": {
				"prop_a": "propA...",
				"prop_b": 2
			}
		}`

		_ = hTest.DoRequestMethodStatus(t, "PATCH", path, []byte(jsonStr), header, http.StatusNoContent)

		// check if it can be read
		feature := checkItem(t, 2)
		var jsonData map[string]interface{}
		errUnMarsh := json.Unmarshal(feature, &jsonData)
		util.Assert(t, errUnMarsh == nil, fmt.Sprintf("%v", errUnMarsh))

		util.Equals(t, "2", jsonData["id"].(string), "feature ID")
		util.Equals(t, "Feature", jsonData["type"].(string), "feature Type")
		props := jsonData["properties"].(map[string]interface{})
		util.Equals(t, "propA...", props["prop_a"].(string), "feature value a")
		util.Equals(t, 2, int(props["prop_b"].(float64)), "feature value b")
		util.Equals(t, "propC", props["prop_c"].(string), "feature value c")
		util.Equals(t, 2, int(props["prop_d"].(float64)), "feature value d")
		geom := jsonData["geometry"].(map[string]interface{})
		util.Equals(t, "Point", geom["type"].(string), "feature Type")
		coordinate := geom["coordinates"].([]interface{})
		util.Equals(t, -120, int(coordinate[0].(float64)), "feature latitude")
		util.Equals(t, 40, int(coordinate[1].(float64)), "feature longitude")

	})
}

func (t *MockTests) TestUpdateFeatureOnlyPropSuccess() {
	t.Test.Run("TestUpdateFeatureOnlyPropSuccess", func(t *testing.T) {
		path := "/collections/mock_a/items/3"
		var header = make(http.Header)
		header.Add("Accept", api.ContentTypeSchemaPatchJSON)

		jsonStr := `{
			"type": "Feature",
			"id": "3",
			"properties": {
				"prop_a": "propA..."
			}
		}`

		_ = hTest.DoRequestMethodStatus(t, "PATCH", path, []byte(jsonStr), header, http.StatusNoContent)

		// check if it can be read
		feature := checkItem(t, 3)
		var jsonData map[string]interface{}
		errUnMarsh := json.Unmarshal(feature, &jsonData)
		util.Assert(t, errUnMarsh == nil, fmt.Sprintf("%v", errUnMarsh))

		util.Equals(t, "3", jsonData["id"].(string), "feature ID")
		util.Equals(t, "Feature", jsonData["type"].(string), "feature Type")
		props := jsonData["properties"].(map[string]interface{})
		util.Equals(t, "propA...", props["prop_a"].(string), "feature value a")
		util.Equals(t, 3, int(props["prop_b"].(float64)), "feature value b")
		util.Equals(t, "propC", props["prop_c"].(string), "feature value c")
		util.Equals(t, 3, int(props["prop_d"].(float64)), "feature value d")
	})
}

func (t *MockTests) TestUpdateFeatureOnlyGeomSuccess() {
	t.Test.Run("TestUpdateFeatureOnlyGeomSuccess", func(t *testing.T) {
		path := "/collections/mock_a/items/4"
		var header = make(http.Header)
		header.Add("Accept", api.ContentTypeSchemaPatchJSON)

		jsonStr := `{
			"type": "Feature",
			"id": "4",
			"geometry": {
				"type": "Point",
				"coordinates": [
				-120,
				40
				]
			}
		}`

		_ = hTest.DoRequestMethodStatus(t, "PATCH", path, []byte(jsonStr), header, http.StatusNoContent)

		// check if it can be read
		feature := checkItem(t, 4)
		var jsonData map[string]interface{}
		errUnMarsh := json.Unmarshal(feature, &jsonData)
		util.Assert(t, errUnMarsh == nil, fmt.Sprintf("%v", errUnMarsh))

		util.Equals(t, "4", jsonData["id"].(string), "feature ID")
		util.Equals(t, "Feature", jsonData["type"].(string), "feature Type")
		props := jsonData["properties"].(map[string]interface{})
		util.Equals(t, "propA", props["prop_a"].(string), "feature value a")
		util.Equals(t, 4, int(props["prop_b"].(float64)), "feature value b")
		util.Equals(t, "propC", props["prop_c"].(string), "feature value c")
		util.Equals(t, 4, int(props["prop_d"].(float64)), "feature value d")
		geom := jsonData["geometry"].(map[string]interface{})
		util.Equals(t, "Point", geom["type"].(string), "feature Type")
		coordinate := geom["coordinates"].([]interface{})
		util.Equals(t, -120, int(coordinate[0].(float64)), "feature latitude")
		util.Equals(t, 40, int(coordinate[1].(float64)), "feature longitude")
	})
}

func (t *MockTests) TestUpdateFeaturePartialGeomFailure() {
	t.Test.Run("TestUpdateFeaturePartialGeomFailure", func(t *testing.T) {
		path := "/collections/mock_a/items/4"
		var header = make(http.Header)
		header.Add("Accept", api.ContentTypeSchemaPatchJSON)

		jsonStr := `{
			"type": "Feature",
			"id": "4",
			"geometry": {
			}
		}`

		resp := hTest.DoRequestMethodStatus(t, "PATCH", path, []byte(jsonStr), header, http.StatusBadRequest)
		body, _ := ioutil.ReadAll(resp.Body)

		// fmt.Println(string(body))

		util.Assert(t, strings.HasPrefix(string(body),
			"Unable to patch feature in Collection - data does not respect schema: mock_a\n\tCaused by: Error at \"/geometry/type\": property \"type\" is missing\nSchema:\n"),
			"feature Error with geometry")
	})
}
