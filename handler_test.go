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
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"github.com/CrunchyData/pg_featureserv/api"
	"github.com/CrunchyData/pg_featureserv/config"
	"github.com/CrunchyData/pg_featureserv/data"
)

// Define a FeatureCollection structure for parsing test data

type Feature struct {
	Type  string                 `json:"type"`
	ID    string                 `json:"id,omitempty"`
	Geom  *json.RawMessage       `json:"geometry"`
	Props map[string]interface{} `json:"properties"`
}

type FeatureCollection struct {
	Type           string      `json:"type"`
	Features       []*Feature  `json:"features"`
	NumberMatched  uint        `json:"numberMatched,omitempty"`
	NumberReturned uint        `json:"numberReturned"`
	TimeStamp      string      `json:"timeStamp,omitempty"`
	Links          []*api.Link `json:"links"`
}

// testConfir is a config spec for using in running tests
var testConfig config.Config = config.Config{
	Server: config.Server{
		HttpHost:   "localhost",
		HttpPort:   9000,
		AssetsPath: "./assets",
	},
	Paging: config.Paging{
		LimitDefault: 10,
		LimitMax:     1000,
	},
	Metadata: config.Metadata{
		Title:       "test",
		Description: "test",
	},
}

func TestMain(m *testing.M) {
	catalogInstance = data.CatMockInstance()
	router = initRouter()
	config.Configuration = testConfig
	os.Exit(m.Run())
}

func TestRoot(t *testing.T) {
	resp := doRequest(t, "/")
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

func TestCollectionResponse(t *testing.T) {
	resp := doRequest(t, "/collections/mock_a/items")
	body, _ := ioutil.ReadAll(resp.Body)

	// Check the response body
	//var v map[string]interface{}
	var v api.FeatureCollectionRaw
	json.Unmarshal(body, &v)

	equals(t, 9, len(v.Features), "# features")
}

func TestLimit(t *testing.T) {
	rr := doRequest(t, "/collections/mock_a/items?limit=3")

	var v FeatureCollection
	json.Unmarshal(readBody(rr), &v)

	equals(t, 3, len(v.Features), "# features")
	equals(t, "1", v.Features[0].ID, "feature 1 id")
	equals(t, "2", v.Features[1].ID, "feature 2 id")
	equals(t, "3", v.Features[2].ID, "feature 3 id")
}

func TestLimitInvalid(t *testing.T) {
	doRequestStatus(t, "/collections/mock_a/items?limit=x", http.StatusBadRequest)
}

func TestParamCase(t *testing.T) {
	rr := doRequest(t, "/collections/mock_a/items?LIMIT=2&Offset=4")

	var v FeatureCollection
	json.Unmarshal(readBody(rr), &v)

	equals(t, 2, len(v.Features), "# features")
	equals(t, "5", v.Features[0].ID, "feature 5 id")
	equals(t, "6", v.Features[1].ID, "feature 6 id")
}

func TestOffset(t *testing.T) {
	rr := doRequest(t, "/collections/mock_a/items?limit=2&offset=4")

	var v FeatureCollection
	json.Unmarshal(readBody(rr), &v)

	equals(t, 2, len(v.Features), "# features")
	equals(t, "5", v.Features[0].ID, "feature 5 id")
	equals(t, "6", v.Features[1].ID, "feature 6 id")
}

func TestOffsetInvalid(t *testing.T) {
	doRequestStatus(t, "/collections/mock_a/items?offset=x", http.StatusBadRequest)
}

func TestBBox(t *testing.T) {
	doRequest(t, "/collections/mock_a/items?bbox=1,2,3,4")
	// TODO: add some tests
}

func TestBBoxInvalid(t *testing.T) {
	doRequestStatus(t, "/collections/mock_a/items?bbox=1,2,3,x", http.StatusBadRequest)
}

func TestCollectionNoFound(t *testing.T) {
	doRequestStatus(t, "/collections/missing", http.StatusNotFound)
}

func TestCollectionItemsNoFound(t *testing.T) {
	doRequestStatus(t, "/collections/missing/items", http.StatusNotFound)
}

func TestFeatureNotFound(t *testing.T) {
	doRequestStatus(t, "/collections/mock_a/items/999", http.StatusNotFound)
}

//--------  Test HTML generation
func TestHTMLHome(t *testing.T) {
	doRequest(t, "/home.html")
}
func TestHTMLConformance(t *testing.T) {
	doRequest(t, "/conformance.html")
}
func TestHTMLCollections(t *testing.T) {
	doRequest(t, "/collections.html")
}
func TestHTMLCollection(t *testing.T) {
	doRequest(t, "/collections/mock_a.html")
}
func TestHTMLItems(t *testing.T) {
	doRequest(t, "/collections/mock_a/items.html")
}
func TestHTMLItem(t *testing.T) {
	doRequest(t, "/collections/mock_a/items/1.html")
}

//===================================================

func readBody(resp *httptest.ResponseRecorder) []byte {
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}

func doRequest(t *testing.T, url string) *httptest.ResponseRecorder {
	return doRequestStatus(t, url, http.StatusOK)
}

func doRequestStatus(t *testing.T, url string,
	statusExpected int) *httptest.ResponseRecorder {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// Check the status code
	//fmt.Println("Status:", rr.Code)
	if status := rr.Code; status != statusExpected {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, statusExpected)
	}
	return rr
}

//---- testing utilities from https://github.com/benbjohnson/testing

// assert fails the test if the condition is false.
func assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: "+msg+"\033[39m\n\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		tb.FailNow()
	}
}

// equals fails the test if exp is not equal to act.
func equals(tb testing.TB, exp, act interface{}, msg string) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d: %s - expected: %#v; got: %#v\n", filepath.Base(file), line, msg, exp, act)
		tb.FailNow()
	}
}
