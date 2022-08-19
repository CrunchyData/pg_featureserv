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
	"github.com/getkin/kin-openapi/openapi3"
)

// checks swagger api contains get operation from collection schema
func TestApiContainsCollectionSchemas(t *testing.T) {
	resp := doRequest(t, "/api")
	body, _ := ioutil.ReadAll(resp.Body)

	var v openapi3.Swagger
	json.Unmarshal(body, &v)

	equals(t, 11, len(v.Paths), "# api paths")
	path := v.Paths.Find("/collections/{collectionId}/schema")
	assert(t, path != nil, "schema path exists")
	equals(t, "Provides access to data representation (schema) for any features in specified collection", path.Description, "schema path present")
	equals(t, "getCollectionSchema", path.Get.OperationID, "schema path get present")
	equals(t, 2, len(path.Get.Parameters), "schema path get present")
	assert(t, path.Get.Parameters.GetByInAndName("path", "collectionId") != nil, "collectionId path parameter exists")
	assert(t, path.Get.Parameters.GetByInAndName("query", "type") != nil, "type query parameter exists")
}

// checks collection schema contains valid data description
func TestGetCollectionCreateSchema(t *testing.T) {
	path := "/collections/mock_a/schema?type=create"
	var header = make(http.Header)
	header.Add("Accept", api.ContentTypeSchemaJSON)

	resp := doRequestMethodStatus(t, "GET", path, nil, header, http.StatusOK)
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))

	var fis openapi3.Schema
	err := fis.UnmarshalJSON(body)
	if err != nil {
		t.Fatal(err)
	}

	equals(t, "This dataset contains mock data about A (9 points)", fis.Description, "feature description")
	equals(t, "https://geojson.org/schema/Point.json", fis.Properties["geometry"].Value.Items.Ref, "feature geometry")
	equals(t, "prop_a", fis.Properties["properties"].Value.Required[0], "feature required a")
	equals(t, "prop_b", fis.Properties["properties"].Value.Required[1], "feature required b")
	equals(t, "Feature", fis.Properties["type"].Value.Default, "feature required b")
}
