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

 Date     : September 2022
 Authors  : Benoit De Mezzo (benoit dot de dot mezzo at oslandia dot com)
*/

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/CrunchyData/pg_featureserv/internal/api"
	"github.com/CrunchyData/pg_featureserv/internal/data"
	util "github.com/CrunchyData/pg_featureserv/internal/utiltest"
	"github.com/getkin/kin-openapi/openapi3"
)

// checks swagger api contains get operation from collection schema
func TestApiContainsCollectionSchemas(t *testing.T) {
	resp := hTest.DoRequest(t, "/api")
	body, _ := ioutil.ReadAll(resp.Body)

	var v openapi3.Swagger
	errUnMarsh := json.Unmarshal(body, &v)
	util.Assert(t, errUnMarsh == nil, fmt.Sprintf("%v", errUnMarsh))

	util.Assert(t, len(v.Paths) > 10, "# api paths")
	path := v.Paths.Find("/collections/{collectionId}/schema")
	util.Assert(t, path != nil, "schema path exists")
	util.Equals(t, "Provides access to data representation (schema) for any features in specified collection", path.Description, "schema path present")
	util.Equals(t, "getCollectionSchema", path.Get.OperationID, "schema path get present")
	util.Equals(t, 2, len(path.Get.Parameters), "# path")
	util.Assert(t, path.Get.Parameters.GetByInAndName("path", "collectionId") != nil, "collectionId path parameter exists")
	util.Assert(t, path.Get.Parameters.GetByInAndName("query", "type") != nil, "type query parameter exists")
}

// checks swagger api contains method PATCH for updating a feaure from a specified collection
func TestApiContainsMethodPostFeature(t *testing.T) {
	resp := hTest.DoRequest(t, "/api")
	body, _ := ioutil.ReadAll(resp.Body)

	var v openapi3.Swagger
	err := json.Unmarshal(body, &v)
	util.Assert(t, err == nil, fmt.Sprintf("%v", err))

	path := v.Paths.Find("/collections/{collectionId}/items")
	util.Assert(t, path != nil, "collection path exists")
	util.Equals(t, "createCollectionFeature", path.Post.OperationID, "method POST present")
}

// checks collection schema contains valid data description
func TestGetCollectionCreateSchema(t *testing.T) {
	path := "/collections/mock_a/schema?type=create"
	var header = make(http.Header)
	header.Add("Accept", api.ContentTypeSchemaJSON)

	resp := hTest.DoRequestMethodStatus(t, "GET", path, nil, header, http.StatusOK)
	body, _ := ioutil.ReadAll(resp.Body)

	var fis openapi3.Schema
	errUnMarsh := json.Unmarshal(body, &fis)
	util.Assert(t, errUnMarsh == nil, fmt.Sprintf("%v", errUnMarsh))

	util.Equals(t, "This dataset contains mock data about A (9 points)", fis.Description, "feature description")
	util.Equals(t, "https://geojson.org/schema/Point.json", fis.Properties["geometry"].Ref, "feature geometry")
	util.Equals(t, "prop_a", fis.Properties["properties"].Value.Required[0], "feature required a")
	util.Equals(t, "prop_b", fis.Properties["properties"].Value.Required[1], "feature required b")
	util.Equals(t, "Feature", fis.Properties["type"].Value.Default, "feature required b")
}

func TestCreateFeature(t *testing.T) {
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
		rr := hTest.DoRequestMethodStatus(t, "POST", "/collections/mock_a/items", []byte(jsonStr), header, http.StatusInternalServerError)
		util.Equals(t, http.StatusInternalServerError, rr.Code, "Should have failed")
		util.Assert(t, strings.Index(rr.Body.String(), fmt.Sprintf(api.ErrMsgCreateFeatureNotConform+"\n", "mock_a")) == 0, "Should have failed with not conform")
	}

	{
		var cols []string
		for _, t := range catalogMock.TableDefs {
			if t.ID == "mock_a" {
				cols = t.Columns
				break
			}
		}
		jsonStr := data.MakeFeatureMockPointAsJSON(0, 12, 34, cols)
		fmt.Println(jsonStr)
		rr := hTest.DoPostRequest(t, "/collections/mock_a/items", []byte(jsonStr), header)

		loc := rr.Header().Get("Location")

		//--- retrieve max feature id
		params := data.QueryParam{Limit: 100, Offset: 0}
		features, _ := catalogMock.TableFeatures(context.Background(), "mock_a", &params)
		maxId := len(features)
		util.Assert(t, len(loc) > 1, "Header location must not be empty")
		util.Assert(t, strings.Contains(loc, fmt.Sprintf("/collections/mock_a/items/%d", maxId)), "Header location must contain valid data")

		// check if it can be read
		checkItem(t, maxId)
	}
}
