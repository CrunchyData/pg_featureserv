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
 Authors  : Jean-philippe Bazonnais (jean-philippe dot bazonnais at ign dot fr)
*/

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/CrunchyData/pg_featureserv/internal/api"
	util "github.com/CrunchyData/pg_featureserv/internal/utiltest"
	"github.com/getkin/kin-openapi/openapi3"
)

func (t *DbTests) TestSimpleReplaceFeatureSuccessDb() {
	t.Test.Run("TestReplaceFeatureSuccessDb", func(t *testing.T) {
		path := "/collections/mock_a/items/9"
		var header = make(http.Header)
		header.Add("Accept", api.ContentTypeSchemaPatchJSON)

		jsonStr := `{
		"type": "Feature",
		"id": "9",
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
			"prop_c": "propC..."
		}
	}`

		hTest.DoRequestMethodStatus(t, "PUT", path, []byte(jsonStr), header, http.StatusNoContent)

		resp := hTest.DoRequestMethodStatus(t, "GET", path, []byte(""), header, http.StatusOK)
		body, _ := ioutil.ReadAll(resp.Body)

		// fmt.Println(string(body))

		var jsonData map[string]interface{}
		err := json.Unmarshal(body, &jsonData)
		util.Assert(t, err == nil, fmt.Sprintf("%v", err))

		util.Equals(t, "9", jsonData["id"].(string), "feature ID")
		util.Equals(t, "Feature", jsonData["type"].(string), "feature Type")
		props := jsonData["properties"].(map[string]interface{})
		util.Equals(t, "propA...", props["prop_a"].(string), "feature value a")
		util.Equals(t, 1, int(props["prop_b"].(float64)), "feature value b")
		util.Equals(t, "propC...", props["prop_c"].(string), "feature value c")
		util.Equals(t, nil, props["prop_d"], "feature value d")
		geom := jsonData["geometry"].(map[string]interface{})
		util.Equals(t, "Point", geom["type"].(string), "feature Type")
		coordinate := geom["coordinates"].([]interface{})
		util.Equals(t, -120, int(coordinate[0].(float64)), "feature latitude")
		util.Equals(t, 40, int(coordinate[1].(float64)), "feature longitude")
	})
}

func (t *DbTests) TestGetComplexCollectionReplaceSchema() {
	t.Test.Run("TestGetComplexCollectionReplaceSchema", func(t *testing.T) {
		path := "/collections/complex.mock_multi/schema?type=replace"
		var header = make(http.Header)
		header.Add("Accept", api.ContentTypeSchemaJSON)

		resp := hTest.DoRequestMethodStatus(t, "GET", path, nil, header, http.StatusOK)
		body, _ := ioutil.ReadAll(resp.Body)

		var fis openapi3.Schema
		errUnMarsh := json.Unmarshal(body, &fis)
		util.Assert(t, errUnMarsh == nil, fmt.Sprintf("%v", errUnMarsh))

		util.Equals(t, "Data for table complex.mock_multi", fis.Description, "feature description")
		util.Equals(t, "GeoJSON Point", fis.Properties["geometry"].Value.Title, "feature geometry")

		util.Equals(t, "Feature", fis.Properties["type"].Value.Default, "feature type is feature")

		val := fis.Properties["properties"].Value
		util.Equals(t, "prop_b", val.Required[0], "feature required bool")
		util.Equals(t, "array", val.Properties["prop_b"].Value.Type, "feature type bool")
		util.Equals(t, "boolean", val.Properties["prop_b"].Value.Items.Value.Type, "feature array type bool")

		util.Equals(t, "prop_d", val.Required[1], "feature required date")
		util.Equals(t, "string", val.Properties["prop_d"].Value.Type, "feature type date")

		util.Equals(t, "prop_f", val.Required[2], "feature required float64")
		util.Equals(t, "number", val.Properties["prop_f"].Value.Type, "feature type float64")

		util.Equals(t, "prop_i", val.Required[3], "feature required int")
		util.Equals(t, "integer", val.Properties["prop_i"].Value.Type, "feature type int")

		util.Equals(t, "prop_j", val.Required[4], "feature required json")
		util.Equals(t, "object", val.Properties["prop_j"].Value.Type, "feature type json")

		util.Equals(t, "prop_l", val.Required[5], "feature required long")
		util.Equals(t, "integer", val.Properties["prop_l"].Value.Type, "feature type long")

		util.Equals(t, "prop_r", val.Required[6], "feature required real")
		util.Equals(t, "number", val.Properties["prop_r"].Value.Type, "feature type real")

		util.Equals(t, "prop_t", val.Required[7], "feature required text")
		util.Equals(t, "string", val.Properties["prop_t"].Value.Type, "feature type text")

		util.Equals(t, "prop_v", val.Required[8], "feature required varchar")
		util.Equals(t, "string", val.Properties["prop_v"].Value.Type, "feature type varchar")
	})
}

func (t *DbTests) TestReplaceComplexFeatureDb() {
	t.Test.Run("TestReplaceComplexFeatureDb", func(t *testing.T) {
		path := "/collections/complex.mock_multi/items/100"
		var header = make(http.Header)
		header.Add("Content-Type", api.ContentTypeSchemaPatchJSON)

		feat := util.MakeGeojsonFeatureMockPoint(100, -50, 35)
		jsonObj, err := json.Marshal(feat)
		util.Assert(t, err == nil, fmt.Sprintf("Error marshalling feature into JSON: %v", err))
		jsonStr := string(jsonObj)

		_ = hTest.DoRequestMethodStatus(t, "PUT", path, []byte(jsonStr), header, http.StatusNoContent)

		// check if it can be read
		checkItem(t, "complex.mock_multi", 100)
	})
}
