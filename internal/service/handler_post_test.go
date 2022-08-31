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
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"testing"

	"github.com/CrunchyData/pg_featureserv/internal/api"
	"github.com/CrunchyData/pg_featureserv/internal/data"
	"github.com/getkin/kin-openapi/openapi3"
)

// checks swagger api contains get operation from collection schema
func TestApiContainsCollectionSchemas(t *testing.T) {
	resp := doRequest(t, "/api")
	body, _ := ioutil.ReadAll(resp.Body)

	var v openapi3.Swagger
	errUnMarsh := json.Unmarshal(body, &v)
	assert(t, errUnMarsh == nil, fmt.Sprintf("%v", errUnMarsh))

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

	var fis openapi3.Schema
	errUnMarsh := json.Unmarshal(body, &fis)
	assert(t, errUnMarsh == nil, fmt.Sprintf("%v", errUnMarsh))

	equals(t, "This dataset contains mock data about A (9 points)", fis.Description, "feature description")
	equals(t, "https://geojson.org/schema/Point.json", fis.Properties["geometry"].Value.Items.Ref, "feature geometry")
	equals(t, "prop_a", fis.Properties["properties"].Value.Required[0], "feature required a")
	equals(t, "prop_b", fis.Properties["properties"].Value.Required[1], "feature required b")
	equals(t, "Feature", fis.Properties["type"].Value.Default, "feature required b")
}

func TestCreateFeature(t *testing.T) {
	//--- retrieve max feature id
	params := data.QueryParam{
		Limit:  100,
		Offset: 0,
	}
	features, _ := catalogMock.TableFeatures(context.Background(), "mock_a", &params)
	maxId := len(features)

	var header = make(http.Header)
	header.Add("Content-Type", "application/geo+json")
	{
		jsonStr := `[{
		"id": 101,
		"name": "Test",
		"email": "test@test.com"
	      }, {
		"id": 102,
		"name": "Sample",
		"email": "sample@test.com"
	    }]`
		rr := doRequestMethodStatus(t, "POST", "/collections/mock_a/items", []byte(jsonStr), header, http.StatusInternalServerError)
		equals(t, http.StatusInternalServerError, rr.Code, "Should have failed")
		assert(t, strings.Index(rr.Body.String(), fmt.Sprintf(api.ErrMsgCreateFeatureNotConform+"\n", "mock_a")) == 0, "Should have failed with not conform")
	}

	{
		jsonStr := catalogMock.MakeFeatureMockPointAsJSON(maxId, 12, 34)
		fmt.Println(jsonStr)
		rr := doPostRequest(t, "/collections/mock_a/items", []byte(jsonStr), header)

		loc := rr.Header().Get("Location")

		assert(t, len(loc) > 1, "Header location must not be empty")
		assert(t, strings.Contains(loc, "/collections/mock_a/items/"), "Header location must contain valid data")

		// retrieve new object id from location header
		parts := strings.Split(loc, "/")
		actId, err := strconv.Atoi(parts[len(parts)-1])
		if err != nil {
			t.Fatal(err)
		}

		assert(t, actId > maxId, fmt.Sprintf("Returned id must be > actual feature number: %d > %d", actId, maxId))

		// check if it can be read
		checkItem(t, actId)
	}
}
