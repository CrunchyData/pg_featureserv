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
	"strings"
	"testing"

	"github.com/CrunchyData/pg_featureserv/internal/api"
	"github.com/CrunchyData/pg_featureserv/internal/data"
	"github.com/CrunchyData/pg_featureserv/util"
	"github.com/getkin/kin-openapi/openapi3"
)

// checks swagger api contains put operation from collection schema
func TestApiContainsMethodPut(t *testing.T) {
	resp := hTest.DoRequest(t, "/api")
	body, _ := ioutil.ReadAll(resp.Body)

	var v openapi3.Swagger
	err := json.Unmarshal(body, &v)
	util.Assert(t, err == nil, fmt.Sprintf("%v", err))

	util.Equals(t, 11, len(v.Paths), "# api paths")
	util.Equals(t, "Provides access to a single feature identitfied by {featureId} from the specified collection", v.Paths.Find("/collections/{collectionId}/items/{featureId}").Description, "feature path present")
	util.Equals(t, "replaceCollectionFeature", v.Paths.Find("/collections/{collectionId}/items/{featureId}").Put.OperationID, "method PUT present")
}

func TestReplaceFeatureSuccess(t *testing.T) {
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

	resp := hTest.DoRequestMethodStatus(t, "PUT", path, []byte(jsonStr), header, http.StatusOK)
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))

	var jsonData map[string]interface{}
	err := json.Unmarshal(body, &jsonData)
	util.Assert(t, err == nil, fmt.Sprintf("%v", err))

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

func TestReplaceFeatureRequiredPropertiesSuccess(t *testing.T) {
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
			"prop_b": 2
		}
	}`

	resp := hTest.DoRequestMethodStatus(t, "PUT", path, []byte(jsonStr), header, http.StatusOK)
	body, _ := ioutil.ReadAll(resp.Body)

	var jsonData map[string]interface{}
	err := json.Unmarshal(body, &jsonData)
	util.Assert(t, err == nil, fmt.Sprintf("%v", err))

	util.Equals(t, "1", jsonData["id"].(string), "feature ID")
	util.Equals(t, "Feature", jsonData["type"].(string), "feature Type")
	props := jsonData["properties"].(map[string]interface{})

	util.Equals(t, "propA...", props["prop_a"].(string), "feature value a")
	util.Equals(t, 2, int(props["prop_b"].(float64)), "feature value b")
	// TODO: Quelle est la valeur par défaut des champs non requis ???
	util.Equals(t, "", props["prop_c"].(string), "feature value c")
	// TODO: Quelle est la valeur par défaut des champs non requis ???
	util.Equals(t, 0, int(props["prop_d"].(float64)), "feature value d")
	geom := jsonData["geometry"].(map[string]interface{})
	util.Equals(t, "Point", geom["type"].(string), "feature Type")
	coordinate := geom["coordinates"].([]interface{})
	util.Equals(t, -120, int(coordinate[0].(float64)), "feature latitude")
	util.Equals(t, 40, int(coordinate[1].(float64)), "feature longitude")
}

func TestReplaceFeatureMissingRequiredPropertiesFailure(t *testing.T) {
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
			"prop_d": 2
		}
	}`

	resp := hTest.DoRequestMethodStatus(t, "PUT", path, []byte(jsonStr), header, http.StatusBadRequest)
	body, _ := ioutil.ReadAll(resp.Body)

	util.Equals(t, http.StatusBadRequest, resp.Code, "Should have failed")
	util.Assert(t, strings.Index(string(body), api.ErrMsgReplaceFeatureNotConform) == 0, "Should have failed with not conform")
}

func TestReplaceFeatureOnlyPropFailure(t *testing.T) {
	path := "/collections/mock_a/items/1"
	var header = make(http.Header)
	header.Add("Accept", api.ContentTypeSchemaPatchJSON)

	jsonStr := `{
		"type": "Feature",
		"id": "1",
		"properties": {
			"prop_a": "propA...",
			"prop_b": 1,
			"prop_c": "propC...",
			"prop_d": 1
		}
	}`

	resp := hTest.DoRequestMethodStatus(t, "PUT", path, []byte(jsonStr), header, http.StatusBadRequest)
	body, _ := ioutil.ReadAll(resp.Body)

	util.Equals(t, http.StatusBadRequest, resp.Code, "Should have failed")
	util.Assert(t, strings.Index(string(body), api.ErrMsgReplaceFeatureNotConform) == 0, "Should have failed with not conform")
}

func TestReplaceFeatureOnlyGeomFailure(t *testing.T) {
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
		}
	}`

	resp := hTest.DoRequestMethodStatus(t, "PUT", path, []byte(jsonStr), header, http.StatusBadRequest)
	body, _ := ioutil.ReadAll(resp.Body)

	util.Equals(t, http.StatusBadRequest, resp.Code, "Should have failed")
	util.Assert(t, strings.Index(string(body), api.ErrMsgReplaceFeatureNotConform) == 0, "Should have failed with not conform")
}

func TestReplaceAfterAll(t *testing.T) {
	// Reset the mock catalog state after all replace tests
	catalogMock = data.CatMockInstance()
	catalogInstance = catalogMock

	hTest = util.MakeHttpTesting("http://test", "/pg_featureserv", InitRouter("/pg_featureserv"))
	Initialize()
}
