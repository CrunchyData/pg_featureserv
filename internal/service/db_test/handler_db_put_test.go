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

 Date     : Octobre 2022
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
)

func (t *DbTests) TestReplaceFeatureSuccessDb() {
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
