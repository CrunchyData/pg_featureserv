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
	"github.com/CrunchyData/pg_featureserv/conf"
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

const urlBase = "http://test"

// testConfir is a config spec for using in running tests
var testConfig conf.Config = conf.Config{
	Server: conf.Server{
		HttpHost:   "0.0.0.0",
		HttpPort:   9000,
		UrlBase:    urlBase,
		AssetsPath: "./assets",
	},
	Paging: conf.Paging{
		LimitDefault: 10,
		LimitMax:     1000,
	},
	Metadata: conf.Metadata{
		Title:       "test",
		Description: "test",
	},
}

var catalogMock *data.CatalogMock

func TestMain(m *testing.M) {
	catalogMock = data.CatMockInstance()
	catalogInstance = catalogMock
	router = initRouter()
	conf.Configuration = testConfig
	os.Exit(m.Run())
}

func TestRoot(t *testing.T) {
	resp := doRequest(t, "/")
	body, _ := ioutil.ReadAll(resp.Body)

	var v api.RootInfo
	json.Unmarshal(body, &v)

	checkLink(t, v.Links[0], api.RelSelf, api.ContentTypeJSON, urlBase+"/")
	checkLink(t, v.Links[1], api.RelAlt, api.ContentTypeHTML, urlBase+"/"+api.RootPageName+".html")
	checkLink(t, v.Links[2], api.RelData, api.ContentTypeJSON, urlBase+"/collections.json")
	checkLink(t, v.Links[3], api.RelFunctions, api.ContentTypeJSON, urlBase+"/functions.json")

	/*
		fmt.Println("Response ==>")
		fmt.Println(v.Title)
		fmt.Println(v.Description)
	*/
}

func TestCollectionsResponse(t *testing.T) {
	path := "/collections"
	resp := doRequest(t, path)
	body, _ := ioutil.ReadAll(resp.Body)

	var v api.CollectionsInfo
	json.Unmarshal(body, &v)

	checkLink(t, v.Links[0], api.RelSelf, api.ContentTypeJSON, urlBase+path+".json")
	checkLink(t, v.Links[1], api.RelAlt, api.ContentTypeHTML, urlBase+path+".html")

	checkCollection(t, v.Collections[0], "mock_a", "Mock A")
	checkCollection(t, v.Collections[1], "mock_b", "Mock B")
	checkCollection(t, v.Collections[2], "mock_c", "Mock C")
}

func TestCollectionResponse(t *testing.T) {
	path := "/collections/mock_a"
	resp := doRequest(t, path)
	body, _ := ioutil.ReadAll(resp.Body)

	var v api.CollectionInfo
	json.Unmarshal(body, &v)

	// use mock data as expected
	tbl := catalogMock.TableDefs[0]

	equals(t, tbl.ID, v.Name, "Name")
	equals(t, tbl.Title, v.Title, "Title")
	equals(t, tbl.Description, v.Description, "Description")
	// check properties
	equals(t, len(tbl.Columns), len(v.Properties), "Properties len")
	for i := 0; i < len(v.Properties); i++ {
		equals(t, tbl.Columns[i], v.Properties[i].Name, "Properties[].Name")
		equals(t, tbl.JSONTypes[i], v.Properties[i].Type, "Properties[].Type")
		equals(t, tbl.ColDesc[i], v.Properties[i].Description, "Properties[].Description")
	}

	checkLink(t, v.Links[0], api.RelSelf, api.ContentTypeJSON, urlBase+path+".json")
	checkLink(t, v.Links[1], api.RelAlt, api.ContentTypeHTML, urlBase+path+".html")
	checkLink(t, v.Links[2], api.RelItems, api.ContentTypeGeoJSON, urlBase+path+"/items.json")
}

func TestCollectionItemsResponse(t *testing.T) {
	path := "/collections/mock_a/items"
	resp := doRequest(t, path)
	body, _ := ioutil.ReadAll(resp.Body)

	var v api.FeatureCollectionRaw
	json.Unmarshal(body, &v)

	equals(t, 9, len(v.Features), "# features")
	checkLink(t, v.Links[0], api.RelSelf, api.ContentTypeJSON, urlBase+path+".json")
	checkLink(t, v.Links[1], api.RelAlt, api.ContentTypeHTML, urlBase+path+".html")
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

func TestQueryParamCase(t *testing.T) {
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

func TestProperties(t *testing.T) {
	// Tests:
	// - property names are non-case-sensitive
	// - names are made unique (properties only include once)
	// - non-existing names are ignored
	rr := doRequest(t, "/collections/mock_a/items?limit=2&properties=PROP_A,prop_C,prop_a,not_prop")

	var v FeatureCollection
	json.Unmarshal(readBody(rr), &v)

	equals(t, 2, len(v.Features), "# features")
	equals(t, 2, len(v.Features[0].Props), "feature 1 # properties")
	equals(t, "propA", v.Features[0].Props["prop_a"], "feature 1 # property A")
	equals(t, "propC", v.Features[0].Props["prop_c"], "feature 1 # property C")
}

// TestPropertiesAll tests that no properties parameter returns all props
func TestPropertiesAll(t *testing.T) {
	rr := doRequest(t, "/collections/mock_a/items?limit=2")

	var v FeatureCollection
	json.Unmarshal(readBody(rr), &v)

	// Note that JSON numbers are read as float64
	equals(t, 2, len(v.Features), "# features")
	equals(t, 4, len(v.Features[0].Props), "feature 1 # properties")
	equals(t, "propA", v.Features[0].Props["prop_a"], "feature 1 # property A")
	equals(t, 1.0, v.Features[0].Props["prop_b"], "feature 1 # property B")
	equals(t, "propC", v.Features[0].Props["prop_c"], "feature 1 # property C")
	equals(t, 999.0, v.Features[0].Props["prop_d"], "feature 1 # property D")
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
func TestHTMLRoot(t *testing.T) {
	doRequest(t, "/index.html")
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

func checkCollection(tb testing.TB, coll *api.CollectionInfo, name string, title string) {
	equals(tb, name, coll.Name, "Collection name")
	equals(tb, title, coll.Title, "Collection title")

	path := "/collections/" + name
	checkLink(tb, coll.Links[0], api.RelSelf, api.ContentTypeJSON, urlBase+path+".json")
	checkLink(tb, coll.Links[1], api.RelAlt, api.ContentTypeHTML, urlBase+path+".html")

	pathItems := path + "/items"
	checkLink(tb, coll.Links[2], api.RelItems, api.ContentTypeGeoJSON, urlBase+pathItems+".json")
}
func checkLink(tb testing.TB, link *api.Link, rel string, conType string, href string) {
	equals(tb, rel, link.Rel, "Link rel")
	equals(tb, conType, link.Type, "Link type")
	equals(tb, href, link.Href, "Link href")
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
