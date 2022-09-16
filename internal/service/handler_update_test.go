package service

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
*/

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/CrunchyData/pg_featureserv/internal/api"
	"github.com/CrunchyData/pg_featureserv/util"
	"github.com/getkin/kin-openapi/openapi3"
)

// checks swagger api contains method PATCH for updating a feaure from a specified collection
func TestApiContainsMethodPatchFeature(t *testing.T) {
	resp := hTest.DoRequest(t, "/api")
	body, _ := ioutil.ReadAll(resp.Body)

	var v openapi3.Swagger
	err := json.Unmarshal(body, &v)
	util.Assert(t, err == nil, fmt.Sprintf("%v", err))

	util.Equals(t, "Provides access to a single feature identitfied by {featureId} from the specified collection", v.Paths.Find("/collections/{collectionId}/items/{featureId}").Description, "path present")
	util.Equals(t, "updateCollectionFeature", v.Paths.Find("/collections/{collectionId}/items/{featureId}").Patch.OperationID, "method PATCH present")
}

func TestGetCollectionUpdateSchema(t *testing.T) {
	path := "/collections/mock_a/schema?type=update"
	var header = make(http.Header)
	header.Add("Accept", api.ContentTypeSchemaJSON)

	resp := hTest.DoRequestMethodStatus(t, "GET", path, nil, header, http.StatusOK)
	body, _ := ioutil.ReadAll(resp.Body)

	var fis openapi3.Schema
	errUnMarsh := json.Unmarshal(body, &fis)
	util.Assert(t, errUnMarsh == nil, fmt.Sprintf("%v", errUnMarsh))

	util.Equals(t, "This dataset contains mock data about A (9 points)", fis.Description, "feature description")
	util.Equals(t, "https://geojson.org/schema/Point.json", fis.Properties["geometry"].Value.Items.Ref, "feature geometry")
	util.Equals(t, 0, len(fis.Required), "no required field")
	util.Equals(t, 2, len(fis.Properties["properties"].Value.Properties["prop_a"].Value.OneOf), "properties have 2 possible values, one is nil")
}

func TestSuccessAllUpdateFeature(t *testing.T) {
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

	resp := hTest.DoRequestMethodStatus(t, "PATCH", path, []byte(jsonStr), header, http.StatusNoContent)

	loc := resp.Header().Get("Location")

	util.Assert(t, len(loc) > 1, "Header location must not be empty")
	util.Equals(t, fmt.Sprintf("http://test/collections/mock_a/items/%d", 1), loc,
		"Header location must contain valid data")

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

}

func TestSuccessPartialUpdateFeature(t *testing.T) {
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

	resp := hTest.DoRequestMethodStatus(t, "PATCH", path, []byte(jsonStr), header, http.StatusNoContent)
	loc := resp.Header().Get("Location")

	util.Assert(t, len(loc) > 1, "Header location must not be empty")
	util.Equals(t, fmt.Sprintf("http://test/collections/mock_a/items/%d", 2), loc,
		"Header location must contain valid data")

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

}

func TestSuccessdOnlyPropUpdateFeature(t *testing.T) {
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

	resp := hTest.DoRequestMethodStatus(t, "PATCH", path, []byte(jsonStr), header, http.StatusNoContent)
	loc := resp.Header().Get("Location")

	util.Assert(t, len(loc) > 1, "Header location must not be empty")
	util.Equals(t, fmt.Sprintf("http://test/collections/mock_a/items/%d", 3), loc,
		"Header location must contain valid data")

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
}

func TestSuccessdOnlyGeomUpdateFeature(t *testing.T) {
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

	resp := hTest.DoRequestMethodStatus(t, "PATCH", path, []byte(jsonStr), header, http.StatusNoContent)
	loc := resp.Header().Get("Location")

	util.Assert(t, len(loc) > 1, "Header location must not be empty")
	util.Equals(t, fmt.Sprintf("http://test/collections/mock_a/items/%d", 4), loc,
		"Header location must contain valid data")

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
}

func TestFailedPartialGeomUpdateFeature(t *testing.T) {
	path := "/collections/mock_a/items/4"
	var header = make(http.Header)
	header.Add("Accept", api.ContentTypeSchemaPatchJSON)

	jsonStr := `{
		"type": "Feature",
		"id": "4",
		"geometry": {
		}
	}`

	resp := hTest.DoRequestMethodStatus(t, "PATCH", path, []byte(jsonStr), header, http.StatusInternalServerError)
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))

	util.Equals(t, "Unable to update feature in Collection: mock_a\n\tCaused by: geojson: invalid geometry\n", string(body), "feature Error with geometry")
}
