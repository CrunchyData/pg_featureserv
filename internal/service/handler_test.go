package service

/*
 Copyright 2019 - 2025 Crunchy Data Solutions, Inc.
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
	"strings"
	"testing"

	"github.com/CrunchyData/pg_featureserv/internal/api"
	"github.com/CrunchyData/pg_featureserv/internal/conf"
	"github.com/CrunchyData/pg_featureserv/internal/data"
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

var basePath = "/pg_featureserv"

var catalogMock *data.CatalogMock

func TestMain(m *testing.M) {
	catalogMock = data.CatMockInstance()
	catalogInstance = catalogMock
	setup(basePath)
	Initialize()
	os.Exit(m.Run())
}

func setup(path string) {
	router = initRouter(path)
	conf.Configuration = conf.Config{
		Server: conf.Server{
			HttpHost:   "0.0.0.0",
			HttpPort:   9000,
			UrlBase:    urlBase,
			BasePath:   path,
			AssetsPath: "../../assets",
			TransformFunctions: []string{
				"ST_Centroid",
				"ST_PointOnSurface",
			},
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
}

func TestRoot(t *testing.T) {
	resp := doRequest(t, "/")
	body, _ := ioutil.ReadAll(resp.Body)

	var v api.RootInfo
	errUnMarsh := json.Unmarshal(body, &v)
	assert(t, errUnMarsh == nil, fmt.Sprintf("%v", errUnMarsh))

	checkLink(t, v.Links[0], api.RelSelf, api.ContentTypeJSON, urlBase+"/"+api.RootPageName)
	checkLink(t, v.Links[1], api.RelAlt, api.ContentTypeHTML, urlBase+"/"+api.RootPageName+".html")
	checkLink(t, v.Links[2], api.RelServiceDesc, api.ContentTypeOpenAPI, urlBase+"/api")
	checkLink(t, v.Links[3], api.RelConformance, api.ContentTypeJSON, urlBase+"/conformance")
	checkLink(t, v.Links[4], api.RelData, api.ContentTypeJSON, urlBase+"/collections")
	checkLink(t, v.Links[5], api.RelFunctions, api.ContentTypeJSON, urlBase+"/functions")

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
	errUnMarsh := json.Unmarshal(body, &v)
	assert(t, errUnMarsh == nil, fmt.Sprintf("%v", errUnMarsh))

	checkLink(t, v.Links[0], api.RelSelf, api.ContentTypeJSON, urlBase+path)
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
	errUnMarsh := json.Unmarshal(body, &v)
	assert(t, errUnMarsh == nil, fmt.Sprintf("%v", errUnMarsh))

	// use mock data as expected
	tbl := catalogMock.TableDefs[0]

	equals(t, tbl.ID, v.Name, "Name")
	equals(t, tbl.Title, v.Title, "Title")
	equals(t, tbl.Description, v.Description, "Description")
	equals(t, tbl.Extent.Minx, v.Extent.Spatial.Extent[0][0], "Extent Minxx")
	equals(t, tbl.Extent.Miny, v.Extent.Spatial.Extent[0][1], "Extent Minxy")
	equals(t, tbl.Extent.Maxx, v.Extent.Spatial.Extent[0][2], "Extent Maxx")
	equals(t, tbl.Extent.Maxy, v.Extent.Spatial.Extent[0][3], "Extent Maxy")
	equals(t, "http://www.opengis.net/def/crs/EPSG/0/4326", v.Extent.Spatial.Crs, "Extent Crs")
	// check properties
	equals(t, len(tbl.Columns), len(v.Properties), "Properties len")
	for i := 0; i < len(v.Properties); i++ {
		equals(t, tbl.Columns[i], v.Properties[i].Name, "Properties[].Name")
		equals(t, tbl.JSONTypes[i], v.Properties[i].Type, "Properties[].Type")
		equals(t, tbl.ColDesc[i], v.Properties[i].Description, "Properties[].Description")
	}

	checkLink(t, v.Links[0], api.RelSelf, api.ContentTypeJSON, urlBase+path)
	checkLink(t, v.Links[1], api.RelAlt, api.ContentTypeHTML, urlBase+path+".html")
	checkLink(t, v.Links[2], api.RelItems, api.ContentTypeGeoJSON, urlBase+path+"/items")
}

func TestCollectionItemsResponse(t *testing.T) {
	path := "/collections/mock_a/items"
	resp := doRequest(t, path)
	body, _ := ioutil.ReadAll(resp.Body)

	var v api.FeatureCollectionRaw
	errUnMarsh := json.Unmarshal(body, &v)
	assert(t, errUnMarsh == nil, fmt.Sprintf("%v", errUnMarsh))

	equals(t, 9, len(v.Features), "# features")
	checkLink(t, v.Links[0], api.RelSelf, api.ContentTypeJSON, urlBase+path)
	checkLink(t, v.Links[1], api.RelAlt, api.ContentTypeHTML, urlBase+path+".html")
}

func TestFilterB(t *testing.T) {
	rr := doRequest(t, "/collections/mock_a/items?prop_b=1")

	var v FeatureCollection
	errUnMarsh := json.Unmarshal(readBody(rr), &v)
	assert(t, errUnMarsh == nil, fmt.Sprintf("%v", errUnMarsh))

	equals(t, 1, len(v.Features), "# features")
}

func TestFilterD(t *testing.T) {
	rr := doRequest(t, "/collections/mock_c/items?prop_d=1")

	var v FeatureCollection
	errUnMarsh := json.Unmarshal(readBody(rr), &v)
	assert(t, errUnMarsh == nil, fmt.Sprintf("%v", errUnMarsh))

	equals(t, 10, len(v.Features), "# features")
}

func TestFilterBD(t *testing.T) {
	rr := doRequest(t, "/collections/mock_c/items?prop_b=2&prop_d=2")

	var v FeatureCollection
	errUnMarsh := json.Unmarshal(readBody(rr), &v)
	assert(t, errUnMarsh == nil, fmt.Sprintf("%v", errUnMarsh))

	equals(t, 1, len(v.Features), "# features")
}

func TestFilterBDNone(t *testing.T) {
	rr := doRequest(t, "/collections/mock_c/items?prop_b=1&prop_d=2")

	var v FeatureCollection
	errUnMarsh := json.Unmarshal(readBody(rr), &v)
	assert(t, errUnMarsh == nil, fmt.Sprintf("%v", errUnMarsh))

	equals(t, 0, len(v.Features), "# features")
}

func TestSortBy(t *testing.T) {
	rr := doRequest(t, "/collections/mock_a/items?sortby=prop_b")

	var v FeatureCollection
	errUnMarsh := json.Unmarshal(readBody(rr), &v)
	assert(t, errUnMarsh == nil, fmt.Sprintf("%v", errUnMarsh))

	equals(t, 9, len(v.Features), "# features")
}

func TestSortByDesc(t *testing.T) {
	rr := doRequest(t, "/collections/mock_a/items?sortby=-prop_b")

	var v FeatureCollection
	errUnMarsh := json.Unmarshal(readBody(rr), &v)
	assert(t, errUnMarsh == nil, fmt.Sprintf("%v", errUnMarsh))

	equals(t, 9, len(v.Features), "# features")
}

func TestSortByAsc(t *testing.T) {
	rr := doRequest(t, "/collections/mock_a/items?sortby=+prop_b")

	var v FeatureCollection
	errUnMarsh := json.Unmarshal(readBody(rr), &v)
	assert(t, errUnMarsh == nil, fmt.Sprintf("%v", errUnMarsh))

	equals(t, 9, len(v.Features), "# features")
}

func TestLimit(t *testing.T) {
	rr := doRequest(t, "/collections/mock_a/items?limit=3")

	var v FeatureCollection
	errUnMarsh := json.Unmarshal(readBody(rr), &v)
	assert(t, errUnMarsh == nil, fmt.Sprintf("%v", errUnMarsh))

	equals(t, 3, len(v.Features), "# features")
	equals(t, "1", v.Features[0].ID, "feature 1 id")
	equals(t, "2", v.Features[1].ID, "feature 2 id")
	equals(t, "3", v.Features[2].ID, "feature 3 id")
}

func TestLimitZero(t *testing.T) {
	rr := doRequest(t, "/collections/mock_a/items?limit=0")

	var v FeatureCollection
	errUnMarsh := json.Unmarshal(readBody(rr), &v)
	assert(t, errUnMarsh == nil, fmt.Sprintf("%v", errUnMarsh))

	equals(t, "FeatureCollection", v.Type, "type FeatureCollection")
	equals(t, 0, len(v.Features), "# features")
}

func TestLimitInvalid(t *testing.T) {
	doRequestStatus(t, "/collections/mock_a/items?limit=x", http.StatusBadRequest)
}

func TestQueryParamCase(t *testing.T) {
	rr := doRequest(t, "/collections/mock_a/items?LIMIT=2&Offset=4")

	var v FeatureCollection
	errUnMarsh := json.Unmarshal(readBody(rr), &v)
	assert(t, errUnMarsh == nil, fmt.Sprintf("%v", errUnMarsh))

	equals(t, 2, len(v.Features), "# features")
	equals(t, "5", v.Features[0].ID, "feature 5 id")
	equals(t, "6", v.Features[1].ID, "feature 6 id")
}

func TestOffset(t *testing.T) {
	rr := doRequest(t, "/collections/mock_a/items?limit=2&offset=4")

	var v FeatureCollection
	errUnMarsh := json.Unmarshal(readBody(rr), &v)
	assert(t, errUnMarsh == nil, fmt.Sprintf("%v", errUnMarsh))

	equals(t, 2, len(v.Features), "# features")
	equals(t, "5", v.Features[0].ID, "feature 5 id")
	equals(t, "6", v.Features[1].ID, "feature 6 id")
}

func TestOffsetInvalid(t *testing.T) {
	doRequestStatus(t, "/collections/mock_a/items?offset=x", http.StatusBadRequest)
}

func TestTransformValid(t *testing.T) {
	doRequest(t, "/collections/mock_a/items?transform=centroid")
	doRequest(t, "/collections/mock_a/items?transform=ST_centroid")
	doRequest(t, "/collections/mock_a/items?transform=st_centroid")
	doRequest(t, "/collections/mock_a/items?transform=pointonsurface")
	doRequest(t, "/collections/mock_a/items?transform=pointonsurface|centroid")
}

func TestTransformInvalid(t *testing.T) {
	// envelope is not defined as a transform function
	doRequestStatus(t, "/collections/mock_a/items?transform=envelope", http.StatusBadRequest)
	doRequestStatus(t, "/collections/mock_a/items?transform=centroid|envelope", http.StatusBadRequest)
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
	// - names are made unique (properties only include once)
	// - non-existing names are ignored
	rr := doRequest(t, "/collections/mock_a/items?limit=2&properties=PROP_A,prop_c,prop_a,not_prop")

	var v FeatureCollection
	errUnMarsh := json.Unmarshal(readBody(rr), &v)
	assert(t, errUnMarsh == nil, fmt.Sprintf("%v", errUnMarsh))

	equals(t, 2, len(v.Features), "# features")
	equals(t, 2, len(v.Features[0].Props), "feature 1 # properties")
	equals(t, "propA", v.Features[0].Props["prop_a"], "feature 1 # property A")
	equals(t, "propC", v.Features[0].Props["prop_c"], "feature 1 # property C")
}

// TestPropertiesAll tests that no properties parameter returns all props
func TestPropertiesAll(t *testing.T) {
	rr := doRequest(t, "/collections/mock_a/items?limit=2")

	var v FeatureCollection
	errUnMarsh := json.Unmarshal(readBody(rr), &v)
	assert(t, errUnMarsh == nil, fmt.Sprintf("%v", errUnMarsh))

	// Note that JSON numbers are read as float64
	equals(t, 2, len(v.Features), "# features")
	equals(t, 4, len(v.Features[0].Props), "feature 1 # properties")

	equals(t, "propA", v.Features[0].Props["prop_a"], "feature 1 # property A")
	equals(t, 1.0, v.Features[0].Props["prop_b"], "feature 1 # property B")
	equals(t, "propC", v.Features[0].Props["prop_c"], "feature 1 # property C")
	equals(t, 1.0, v.Features[0].Props["prop_d"], "feature 1 # property D")
}

func TestCollectionNotFound(t *testing.T) {
	doRequestStatus(t, "/collections/missing", http.StatusNotFound)
}

func TestCollectionMissingItemsNotFound(t *testing.T) {
	doRequestStatus(t, "/collections/missing/items", http.StatusNotFound)
}

func TestFeatureNotFound(t *testing.T) {
	doRequestStatus(t, "/collections/mock_a/items/999", http.StatusNotFound)
}

//=============  Test functions

func TestFunctionsJSON(t *testing.T) {
	path := "/functions"
	resp := doRequest(t, path)
	body, _ := ioutil.ReadAll(resp.Body)

	var v api.FunctionsInfo
	errUnMarsh := json.Unmarshal(body, &v)
	assert(t, errUnMarsh == nil, fmt.Sprintf("%v", errUnMarsh))

	checkLink(t, v.Links[0], api.RelSelf, api.ContentTypeJSON, urlBase+path)
	checkLink(t, v.Links[1], api.RelAlt, api.ContentTypeHTML, urlBase+path+".html")

	for i, fun := range catalogMock.FunctionDefs {
		checkFunctionSummary(t, v.Functions[i], fun)
	}
}

func TestFunctionJSON(t *testing.T) {
	for _, fun := range catalogMock.FunctionDefs {
		//fun := catalogMock.FunctionDefs[1]
		checkFunction(t, fun)
	}
}

func TestFunctionNotFound(t *testing.T) {
	doRequestStatus(t, "/functions/missing", http.StatusNotFound)
}

func TestFunctionMissingItemsNotFound(t *testing.T) {
	doRequestStatus(t, "/functions/missing/items", http.StatusNotFound)
}

// ============  Test HTML generation
// For now these just test that the template executes correctly
// correctness/completess of HTML is not tested
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
func TestHTMLFunctions(t *testing.T) {
	rr := doRequest(t, "/functions.html")
	for _, fun := range catalogMock.FunctionDefs {
		if !strings.Contains(rr.Body.String(), "http://test/functions/"+fun.Name+".json") {
			t.Errorf("Functions response should contain reference to " + fun.Name + ".json")
		}
	}
}
func TestHTMLFunction(t *testing.T) {
	doRequest(t, "/functions/fun_a.html")
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
	req, err := http.NewRequest("GET", basePath+url, nil)
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
	checkLink(tb, coll.Links[0], api.RelSelf, api.ContentTypeJSON, urlBase+path)
	checkLink(tb, coll.Links[1], api.RelAlt, api.ContentTypeHTML, urlBase+path+".html")

	pathItems := path + "/items"
	checkLink(tb, coll.Links[2], api.RelItems, api.ContentTypeGeoJSON, urlBase+pathItems)
}
func checkLink(tb testing.TB, link *api.Link, rel string, conType string, href string) {
	equals(tb, rel, link.Rel, "Link rel")
	equals(tb, conType, link.Type, "Link type")
	equals(tb, href, link.Href, "Link href")
}

func checkFunctionSummary(tb testing.TB, v *api.FunctionSummary, fun *data.Function) {
	equals(tb, fun.Name, v.Name, "Function name")
	equals(tb, fun.Description, v.Description, "Function description")

	path := "/functions/" + fun.Name
	checkLink(tb, v.Links[0], api.RelSelf, api.ContentTypeJSON, urlBase+path)
	checkLink(tb, v.Links[1], api.RelAlt, api.ContentTypeHTML, urlBase+path+".html")

	pathItems := path + "/items"
	itemsType := api.ContentTypeJSON
	if fun.IsGeometryFunction() {
		itemsType = api.ContentTypeGeoJSON
	}
	checkLink(tb, v.Links[2], api.RelItems, itemsType, urlBase+pathItems)
}
func checkFunction(t *testing.T, fun *data.Function) {
	path := "/functions/" + fun.ID
	resp := doRequest(t, path)
	body, _ := ioutil.ReadAll(resp.Body)

	var v api.FunctionInfo
	errUnMarsh := json.Unmarshal(body, &v)
	assert(t, errUnMarsh == nil, fmt.Sprintf("%v", errUnMarsh))

	equals(t, fun.ID, v.Name, "Name")
	equals(t, fun.Description, v.Description, "Description")

	//--- check parameters
	assert(t, v.Parameters != nil, "Parameters property must be present")
	equals(t, len(fun.InNames), len(v.Parameters), "Parameters len")
	for i := 0; i < len(v.Parameters); i++ {
		equals(t, fun.InNames[i], v.Parameters[i].Name, "Parameters[].Name")
		equals(t, fun.InDbTypes[i], v.Parameters[i].Type, "Parameters[].Type")
	}

	//--- check properties
	assert(t, v.Properties != nil, "Properties property must be present")
	equals(t, len(fun.OutNames), len(v.Properties), "Properties len")
	for i := 0; i < len(v.Properties); i++ {
		equals(t, fun.OutNames[i], v.Properties[i].Name, "Properties[].Name")
		equals(t, fun.OutJSONTypes[i], v.Properties[i].Type, "Properties[].Type")
	}

	//--- check links
	checkLink(t, v.Links[0], api.RelSelf, api.ContentTypeJSON, urlBase+path)
	checkLink(t, v.Links[1], api.RelAlt, api.ContentTypeHTML, urlBase+path+".html")
	itemsType := api.ContentTypeJSON
	if fun.IsGeometryFunction() {
		itemsType = api.ContentTypeGeoJSON
	}
	checkLink(t, v.Links[2], api.RelItems, itemsType, urlBase+path+"/items")
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
