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
	"testing"

	"github.com/CrunchyData/pg_featureserv/internal/data"
	"github.com/CrunchyData/pg_featureserv/util"
	"github.com/getkin/kin-openapi/openapi3"
)

// checks swagger api contains delete feature operation from collection
func TestApiContainsDeleteFeature(t *testing.T) {
	resp := hTest.DoRequest(t, "/api")
	body, _ := ioutil.ReadAll(resp.Body)

	var v openapi3.Swagger
	errUnMarsh := json.Unmarshal(body, &v)
	util.Assert(t, errUnMarsh == nil, fmt.Sprintf("%v", errUnMarsh))

	path := v.Paths.Find("/collections/{collectionId}/items/{featureId}")
	util.Assert(t, path != nil, "Delete feature path does not exist")
	util.Equals(t, "deleteCollectionFeature", path.Delete.OperationID, "Delete path not present")
	util.Equals(t, 2, len(path.Delete.Parameters), "2 parameters expected")
	util.Assert(t, path.Delete.Parameters.GetByInAndName("path", "collectionId") != nil, "collectionId path parameter not found")
	util.Assert(t, path.Delete.Parameters.GetByInAndName("path", "featureId") != nil, "featureId path parameter not found")
}

func TestDeleteExistingFeature(t *testing.T) {

	params := data.QueryParam{Limit: 100, Offset: 0}
	features, _ := catalogMock.TableFeatures(context.Background(), "mock_b", &params)
	nb_feature_before := len(features)

	hTest.DoDeleteRequestStatus(t, "/collections/mock_b/items/1", http.StatusNoContent)

	featuresafter, _ := catalogMock.TableFeatures(context.Background(), "mock_b", &params)

	util.Assert(t, nb_feature_before-1 == len(featuresafter), "Number of features should have decreased")

	// Ensuring that the feature is no more present
	hTest.DoRequestStatus(t, "/collections/mock_b/items/1", http.StatusNotFound)
}

func TestDeleteFeatureErrorMalformedFeatureId(t *testing.T) {

	params := data.QueryParam{Limit: 100, Offset: 0}
	features, _ := catalogMock.TableFeatures(context.Background(), "mock_b", &params)
	nb_feature_before := len(features)

	hTest.DoDeleteRequestStatus(t, "/collections/mock_b/items/a1", http.StatusBadRequest)

	featuresafter, _ := catalogMock.TableFeatures(context.Background(), "mock_b", &params)

	util.Assert(t, nb_feature_before == len(featuresafter), "Number of features should be the same after request")

}

func TestDeleteFeatureErrorUnusedQueryParameters(t *testing.T) {

	hTest.DoDeleteRequestStatus(t, "/collections/mockette/items/1?param1=value1", http.StatusBadRequest)
	hTest.DoDeleteRequestStatus(t, "/collections/mockette/items/1?param2=value2&param3=value3", http.StatusBadRequest)

}

func TestDeleteFeatureErrorUnknownCollection(t *testing.T) {

	hTest.DoDeleteRequestStatus(t, "/collections/mockette/items/1", http.StatusNotFound)

}

func TestDeleteUnknownFeature(t *testing.T) {

	status := hTest.DoDeleteRequestStatus(t, "/collections/mock_b/items/123", http.StatusNotFound)
	util.Equals(t, http.StatusNotFound, status.Code, "Should have failed")

}
