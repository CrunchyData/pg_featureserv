package main

/*
 Copyright 2019 Crunchy Data Solutions, Inc.
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
	"net/http/httptest"
	"os"
	"testing"

	"github.com/CrunchyData/pg_featureserv/api"
	"github.com/CrunchyData/pg_featureserv/data"
)

func TestMain(m *testing.M) {
	catalogInstance = data.CatMockInstance()
	os.Exit(m.Run())
}

func TestRootHandler(t *testing.T) {
	resp := execHTTPRequest(t, "/", handleRootJSON)
	body, _ := ioutil.ReadAll(resp.Body)

	// Check the response body
	//var v map[string]interface{}
	var v api.RootInfo
	json.Unmarshal(body, &v)

	fmt.Println("Response ==>")
	fmt.Println(v.Title)
	fmt.Println(v.Description)
	/*
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	*/
}

func TestCollectionNoFound(t *testing.T) {
	execHTTPRequestStatus(t, "/collections/missing", handleCollection, http.StatusNotFound)
}

func TestCollectionItemsNoFound(t *testing.T) {
	execHTTPRequestStatus(t, "/collections/missing/items", handleCollectionItems, http.StatusNotFound)
}

func TestFeatureNotFound(t *testing.T) {
	execHTTPRequestStatus(t, "/collections/mock_a/items/999", handleItem, http.StatusNotFound)
}

func execHTTPRequest(t *testing.T, url string, handler appHandler) *httptest.ResponseRecorder {
	return execHTTPRequestStatus(t, url, handler, http.StatusOK)
}

func execHTTPRequestStatus(t *testing.T, url string,
	handler appHandler,
	statusExpected int) *httptest.ResponseRecorder {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	// Check the status code
	fmt.Println("Status:", rr.Code)
	if status := rr.Code; status != statusExpected {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, statusExpected)
	}
	return rr
}
