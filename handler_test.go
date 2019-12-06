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
	"testing"

	"github.com/CrunchyData/pg_featureserv/api"
)

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

func TestCollectionMissing(t *testing.T) {
	execHTTPRequestStatus(t, "/collections/missing", handleCollection, http.StatusNotFound)

	fmt.Println("Response ==>")

	/*
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	*/
}

func execHTTPRequest(t *testing.T, url string,
	handler http.HandlerFunc) *httptest.ResponseRecorder {
	return execHTTPRequestStatus(t, url, handler, http.StatusOK)
}

func execHTTPRequestStatus(t *testing.T, url string,
	handler http.HandlerFunc,
	statusExpected int) *httptest.ResponseRecorder {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	h := http.HandlerFunc(handler)
	h.ServeHTTP(rr, req)

	// Check the status code
	fmt.Println("Status:", rr.Code)
	if status := rr.Code; status != statusExpected {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, statusExpected)
	}
	return rr
}