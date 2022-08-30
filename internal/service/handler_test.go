package service

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
	"os"
	"reflect"

	"strconv"
	"strings"
	"testing"

	"github.com/CrunchyData/pg_featureserv/internal/api"
	"github.com/CrunchyData/pg_featureserv/internal/conf"
	"github.com/CrunchyData/pg_featureserv/internal/data"
	"github.com/CrunchyData/pg_featureserv/util"
)

var hTest util.HttpTesting

var catalogMock *data.CatalogMock

func TestMain(m *testing.M) {
	conf.Configuration.Database.AllowWrite = true

	catalogMock = data.CatMockInstance()
	catalogInstance = catalogMock

	hTest = util.MakeHttpTesting("http://test", "/pg_featureserv", InitRouter("/pg_featureserv"))
	Initialize()

	os.Exit(m.Run())
}

func TestRoot(t *testing.T) {
	resp := hTest.DoRequest(t, "/")
	body, _ := ioutil.ReadAll(resp.Body)

	var v api.RootInfo
	errUnMarsh := json.Unmarshal(body, &v)
	util.Assert(t, errUnMarsh == nil, fmt.Sprintf("%v", errUnMarsh))

	checkLink(t, v.Links[0], api.RelSelf, api.ContentTypeJSON, hTest.UrlBase+"/"+api.RootPageName)
	checkLink(t, v.Links[1], api.RelAlt, api.ContentTypeHTML, hTest.UrlBase+"/"+api.RootPageName+".html")
	checkLink(t, v.Links[2], api.RelServiceDesc, api.ContentTypeOpenAPI, hTest.UrlBase+"/api")
	checkLink(t, v.Links[3], api.RelConformance, api.ContentTypeJSON, hTest.UrlBase+"/conformance")
	checkLink(t, v.Links[4], api.RelData, api.ContentTypeJSON, hTest.UrlBase+"/collections")
	checkLink(t, v.Links[5], api.RelFunctions, api.ContentTypeJSON, hTest.UrlBase+"/functions")

	/*
		fmt.Println("Response ==>")
		fmt.Println(v.Title)
		fmt.Println(v.Description)
	*/
}

func TestCollectionsResponse(t *testing.T) {
	path := "/collections"
	resp := hTest.DoRequest(t, path)
	body, _ := ioutil.ReadAll(resp.Body)

	var v api.CollectionsInfo
	errUnMarsh := json.Unmarshal(body, &v)
	util.Assert(t, errUnMarsh == nil, fmt.Sprintf("%v", errUnMarsh))

	checkLink(t, v.Links[0], api.RelSelf, api.ContentTypeJSON, hTest.UrlBase+path)
	checkLink(t, v.Links[1], api.RelAlt, api.ContentTypeHTML, hTest.UrlBase+path+".html")

	checkCollection(t, v.Collections[0], "mock_a", "Mock A")
	checkCollection(t, v.Collections[1], "mock_b", "Mock B")
	checkCollection(t, v.Collections[2], "mock_c", "Mock C")
}

func TestCollectionResponse(t *testing.T) {
	path := "/collections/mock_a"
	resp := hTest.DoRequest(t, path)
	body, _ := ioutil.ReadAll(resp.Body)

	var v api.CollectionInfo
	errUnMarsh := json.Unmarshal(body, &v)
	util.Assert(t, errUnMarsh == nil, fmt.Sprintf("%v", errUnMarsh))

	// use mock data as expected
	tbl := catalogMock.TableDefs[0]

	util.Equals(t, tbl.ID, v.Name, "Name")
	util.Equals(t, tbl.Title, v.Title, "Title")
	util.Equals(t, tbl.Description, v.Description, "Description")
	// check properties
	util.Equals(t, len(tbl.Columns), len(v.Properties), "Properties len")
	for i := 0; i < len(v.Properties); i++ {
		util.Equals(t, tbl.Columns[i], v.Properties[i].Name, "Properties[].Name")
		util.Equals(t, tbl.JSONTypes[i], v.Properties[i].Type, "Properties[].Type")
		util.Equals(t, tbl.ColDesc[i], v.Properties[i].Description, "Properties[].Description")
	}

	checkLink(t, v.Links[0], api.RelSelf, api.ContentTypeJSON, hTest.UrlBase+path)
	checkLink(t, v.Links[1], api.RelAlt, api.ContentTypeHTML, hTest.UrlBase+path+".html")
	checkLink(t, v.Links[2], api.RelItems, api.ContentTypeGeoJSON, hTest.UrlBase+path+"/items")
}

func TestCollectionItemsResponse(t *testing.T) {
	path := "/collections/mock_a/items"
	resp := hTest.DoRequest(t, path)
	body, _ := ioutil.ReadAll(resp.Body)

	var v api.FeatureCollectionRaw
	errUnMarsh := json.Unmarshal(body, &v)
	util.Assert(t, errUnMarsh == nil, fmt.Sprintf("%v", errUnMarsh))

	util.Equals(t, int(catalogMock.TableSize("mock_a")), len(v.Features), "# features")
	checkLink(t, v.Links[0], api.RelSelf, api.ContentTypeJSON, hTest.UrlBase+path)
	checkLink(t, v.Links[1], api.RelAlt, api.ContentTypeHTML, hTest.UrlBase+path+".html")
}

// check if item is available and is not empty
func TestCollectionItem(t *testing.T) {
	checkItem(t, 1)
}

func TestFilterB(t *testing.T) {
	rr := hTest.DoRequest(t, "/collections/mock_a/items?prop_b=1")

	var v util.FeatureCollection
	errUnMarsh := json.Unmarshal(hTest.ReadBody(rr), &v)
	util.Assert(t, errUnMarsh == nil, fmt.Sprintf("%v", errUnMarsh))

	util.Equals(t, 1, len(v.Features), "# features")
}

func TestFilterD(t *testing.T) {
	rr := hTest.DoRequest(t, "/collections/mock_c/items?prop_d=1")

	var v util.FeatureCollection
	errUnMarsh := json.Unmarshal(hTest.ReadBody(rr), &v)
	util.Assert(t, errUnMarsh == nil, fmt.Sprintf("%v", errUnMarsh))

	util.Equals(t, 10, len(v.Features), "# features")
}

func TestFilterBD(t *testing.T) {
	rr := hTest.DoRequest(t, "/collections/mock_c/items?prop_b=2&prop_d=2")

	var v util.FeatureCollection
	errUnMarsh := json.Unmarshal(hTest.ReadBody(rr), &v)
	util.Assert(t, errUnMarsh == nil, fmt.Sprintf("%v", errUnMarsh))

	util.Equals(t, 1, len(v.Features), "# features")
}

func TestFilterBDNone(t *testing.T) {
	rr := hTest.DoRequest(t, "/collections/mock_c/items?prop_b=1&prop_d=2")

	var v util.FeatureCollection
	errUnMarsh := json.Unmarshal(hTest.ReadBody(rr), &v)
	util.Assert(t, errUnMarsh == nil, fmt.Sprintf("%v", errUnMarsh))

	util.Equals(t, 0, len(v.Features), "# features")
}

func TestSortBy(t *testing.T) {
	rr := hTest.DoRequest(t, "/collections/mock_a/items?sortby=prop_b")

	var v util.FeatureCollection
	errUnMarsh := json.Unmarshal(hTest.ReadBody(rr), &v)
	util.Assert(t, errUnMarsh == nil, fmt.Sprintf("%v", errUnMarsh))

	util.Equals(t, int(catalogMock.TableSize("mock_a")), len(v.Features), "# features")
}

func TestSortByDesc(t *testing.T) {
	rr := hTest.DoRequest(t, "/collections/mock_a/items?sortby=-prop_b")

	var v util.FeatureCollection
	errUnMarsh := json.Unmarshal(hTest.ReadBody(rr), &v)
	util.Assert(t, errUnMarsh == nil, fmt.Sprintf("%v", errUnMarsh))

	util.Equals(t, int(catalogMock.TableSize("mock_a")), len(v.Features), "# features")
}

func TestSortByAsc(t *testing.T) {
	rr := hTest.DoRequest(t, "/collections/mock_a/items?sortby=+prop_b")

	var v util.FeatureCollection
	errUnMarsh := json.Unmarshal(hTest.ReadBody(rr), &v)
	util.Assert(t, errUnMarsh == nil, fmt.Sprintf("%v", errUnMarsh))

	util.Equals(t, int(catalogMock.TableSize("mock_a")), len(v.Features), "# features")
}

func TestLimit(t *testing.T) {
	rr := hTest.DoRequest(t, "/collections/mock_a/items?limit=3")

	var v util.FeatureCollection
	errUnMarsh := json.Unmarshal(hTest.ReadBody(rr), &v)
	util.Assert(t, errUnMarsh == nil, fmt.Sprintf("%v", errUnMarsh))

	util.Equals(t, 3, len(v.Features), "# features")
	util.Equals(t, "1", v.Features[0].ID, "feature 1 id")
	util.Equals(t, "2", v.Features[1].ID, "feature 2 id")
	util.Equals(t, "3", v.Features[2].ID, "feature 3 id")
}

func TestLimitZero(t *testing.T) {
	rr := hTest.DoRequest(t, "/collections/mock_a/items?limit=0")

	var v util.FeatureCollection
	errUnMarsh := json.Unmarshal(hTest.ReadBody(rr), &v)
	util.Assert(t, errUnMarsh == nil, fmt.Sprintf("%v", errUnMarsh))

	util.Equals(t, "FeatureCollection", v.Type, "type FeatureCollection")
	util.Equals(t, 0, len(v.Features), "# features")
}

func TestLimitInvalid(t *testing.T) {
	hTest.DoRequestStatus(t, "/collections/mock_a/items?limit=x", http.StatusBadRequest)
}

func TestQueryParamCase(t *testing.T) {
	rr := hTest.DoRequest(t, "/collections/mock_a/items?LIMIT=2&Offset=4")

	var v util.FeatureCollection
	errUnMarsh := json.Unmarshal(hTest.ReadBody(rr), &v)
	util.Assert(t, errUnMarsh == nil, fmt.Sprintf("%v", errUnMarsh))

	util.Equals(t, 2, len(v.Features), "# features")
	util.Equals(t, "5", v.Features[0].ID, "feature 5 id")
	util.Equals(t, "6", v.Features[1].ID, "feature 6 id")
}

func TestOffset(t *testing.T) {
	rr := hTest.DoRequest(t, "/collections/mock_a/items?limit=2&offset=4")

	var v util.FeatureCollection
	errUnMarsh := json.Unmarshal(hTest.ReadBody(rr), &v)
	util.Assert(t, errUnMarsh == nil, fmt.Sprintf("%v", errUnMarsh))

	util.Equals(t, 2, len(v.Features), "# features")
	util.Equals(t, "5", v.Features[0].ID, "feature 5 id")
	util.Equals(t, "6", v.Features[1].ID, "feature 6 id")
}

func TestOffsetInvalid(t *testing.T) {
	hTest.DoRequestStatus(t, "/collections/mock_a/items?offset=x", http.StatusBadRequest)
}

func TestTransformValid(t *testing.T) {
	hTest.DoRequest(t, "/collections/mock_a/items?transform=centroid")
	hTest.DoRequest(t, "/collections/mock_a/items?transform=ST_centroid")
	hTest.DoRequest(t, "/collections/mock_a/items?transform=st_centroid")
	hTest.DoRequest(t, "/collections/mock_a/items?transform=pointonsurface")
	hTest.DoRequest(t, "/collections/mock_a/items?transform=pointonsurface|centroid")
}

func TestTransformInvalid(t *testing.T) {
	// envelope is not defined as a transform function
	hTest.DoRequestStatus(t, "/collections/mock_a/items?transform=envelope", http.StatusBadRequest)
	hTest.DoRequestStatus(t, "/collections/mock_a/items?transform=centroid|envelope", http.StatusBadRequest)
}

func TestBBox(t *testing.T) {
	hTest.DoRequest(t, "/collections/mock_a/items?bbox=1,2,3,4")
	// TODO: add some tests
}

func TestBBoxInvalid(t *testing.T) {
	hTest.DoRequestStatus(t, "/collections/mock_a/items?bbox=1,2,3,x", http.StatusBadRequest)
}

func TestProperties(t *testing.T) {
	// Tests:
	// - property names are non-case-sensitive
	// - names are made unique (properties only include once)
	// - non-existing names are ignored
	rr := hTest.DoRequest(t, "/collections/mock_a/items?limit=2&properties=PROP_A,prop_C,prop_a,not_prop")

	var v util.FeatureCollection
	errUnMarsh := json.Unmarshal(hTest.ReadBody(rr), &v)
	util.Assert(t, errUnMarsh == nil, fmt.Sprintf("%v", errUnMarsh))

	util.Equals(t, 2, len(v.Features), "# features")
	util.Equals(t, 2, len(v.Features[0].Props), "feature 1 # properties")
	util.Equals(t, "propA", v.Features[0].Props["prop_a"], "feature 1 # property A")
	util.Equals(t, "propC", v.Features[0].Props["prop_c"], "feature 1 # property C")
}

// TestPropertiesAll tests that no properties parameter returns all props
func TestPropertiesAll(t *testing.T) {
	rr := hTest.DoRequest(t, "/collections/mock_a/items?limit=2")

	var v util.FeatureCollection
	errUnMarsh := json.Unmarshal(hTest.ReadBody(rr), &v)
	util.Assert(t, errUnMarsh == nil, fmt.Sprintf("%v", errUnMarsh))

	// Note that JSON numbers are read as float64
	util.Equals(t, 2, len(v.Features), "# features")
	util.Equals(t, 4, len(v.Features[0].Props), "feature 1 # properties")

	util.Equals(t, "propA", v.Features[0].Props["prop_a"], "feature 1 # property A")
	util.Equals(t, 1.0, v.Features[0].Props["prop_b"], "feature 1 # property B")
	util.Equals(t, "propC", v.Features[0].Props["prop_c"], "feature 1 # property C")
	util.Equals(t, 1.0, v.Features[0].Props["prop_d"], "feature 1 # property D")
}

func TestCollectionNotFound(t *testing.T) {
	hTest.DoRequestStatus(t, "/collections/missing", http.StatusNotFound)
}

func TestCollectionMissingItemsNotFound(t *testing.T) {
	hTest.DoRequestStatus(t, "/collections/missing/items", http.StatusNotFound)
}

func TestFeatureNotFound(t *testing.T) {
	hTest.DoRequestStatus(t, "/collections/mock_a/items/999", http.StatusNotFound)
}

//=============  Test functions

func TestFunctionsJSON(t *testing.T) {
	path := "/functions"
	resp := hTest.DoRequest(t, path)
	body, _ := ioutil.ReadAll(resp.Body)

	var v api.FunctionsInfo
	errUnMarsh := json.Unmarshal(body, &v)
	util.Assert(t, errUnMarsh == nil, fmt.Sprintf("%v", errUnMarsh))

	checkLink(t, v.Links[0], api.RelSelf, api.ContentTypeJSON, hTest.UrlBase+path)
	checkLink(t, v.Links[1], api.RelAlt, api.ContentTypeHTML, hTest.UrlBase+path+".html")

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
	hTest.DoRequestStatus(t, "/functions/missing", http.StatusNotFound)
}

func TestFunctionMissingItemsNotFound(t *testing.T) {
	hTest.DoRequestStatus(t, "/functions/missing/items", http.StatusNotFound)
}

// ============  Test HTML generation
// For now these just test that the template executes correctly
// correctness/completess of HTML is not tested
func TestHTMLRoot(t *testing.T) {
	hTest.DoRequest(t, "/index.html")
}
func TestHTMLConformance(t *testing.T) {
	hTest.DoRequest(t, "/conformance.html")
}
func TestHTMLCollections(t *testing.T) {
	hTest.DoRequest(t, "/collections.html")
}
func TestHTMLCollection(t *testing.T) {
	hTest.DoRequest(t, "/collections/mock_a.html")
}
func TestHTMLItems(t *testing.T) {
	hTest.DoRequest(t, "/collections/mock_a/items.html")
}
func TestHTMLItem(t *testing.T) {
	hTest.DoRequest(t, "/collections/mock_a/items/1.html")
}
func TestHTMLFunctions(t *testing.T) {
	rr := hTest.DoRequest(t, "/functions.html")
	for _, fun := range catalogMock.FunctionDefs {
		if !strings.Contains(rr.Body.String(), "http://test/functions/"+fun.Name+".json") {
			t.Errorf("Functions response should contain reference to " + fun.Name + ".json")
		}
	}
}
func TestHTMLFunction(t *testing.T) {
	hTest.DoRequest(t, "/functions/fun_a.html")
}

//===================================================

func checkCollection(tb testing.TB, coll *api.CollectionInfo, name string, title string) {
	util.Equals(tb, name, coll.Name, "Collection name")
	util.Equals(tb, title, coll.Title, "Collection title")

	path := "/collections/" + name
	checkLink(tb, coll.Links[0], api.RelSelf, api.ContentTypeJSON, hTest.UrlBase+path)
	checkLink(tb, coll.Links[1], api.RelAlt, api.ContentTypeHTML, hTest.UrlBase+path+".html")

	pathItems := path + "/items"
	checkLink(tb, coll.Links[2], api.RelItems, api.ContentTypeGeoJSON, hTest.UrlBase+pathItems)
}

func checkLink(tb testing.TB, link *api.Link, rel string, conType string, href string) {
	util.Equals(tb, rel, link.Rel, "Link rel")
	util.Equals(tb, conType, link.Type, "Link type")
	util.Equals(tb, href, link.Href, "Link href")
}

func checkFunctionSummary(tb testing.TB, v *api.FunctionSummary, fun *data.Function) {
	util.Equals(tb, fun.Name, v.Name, "Function name")
	util.Equals(tb, fun.Description, v.Description, "Function description")

	path := "/functions/" + fun.Name
	checkLink(tb, v.Links[0], api.RelSelf, api.ContentTypeJSON, hTest.UrlBase+path)
	checkLink(tb, v.Links[1], api.RelAlt, api.ContentTypeHTML, hTest.UrlBase+path+".html")

	pathItems := path + "/items"
	itemsType := api.ContentTypeJSON
	if fun.IsGeometryFunction() {
		itemsType = api.ContentTypeGeoJSON
	}
	checkLink(tb, v.Links[2], api.RelItems, itemsType, hTest.UrlBase+pathItems)
}
func checkFunction(t *testing.T, fun *data.Function) {
	path := "/functions/" + fun.ID
	resp := hTest.DoRequest(t, path)
	body, _ := ioutil.ReadAll(resp.Body)

	var v api.FunctionInfo
	errUnMarsh := json.Unmarshal(body, &v)
	util.Assert(t, errUnMarsh == nil, fmt.Sprintf("%v", errUnMarsh))

	util.Equals(t, fun.ID, v.Name, "Name")
	util.Equals(t, fun.Description, v.Description, "Description")

	//--- check parameters
	util.Assert(t, v.Parameters != nil, "Parameters property must be present")
	util.Equals(t, len(fun.InNames), len(v.Parameters), "Parameters len")
	for i := 0; i < len(v.Parameters); i++ {
		util.Equals(t, fun.InNames[i], v.Parameters[i].Name, "Parameters[].Name")
		util.Equals(t, fun.InDbTypes[i], v.Parameters[i].Type, "Parameters[].Type")
	}

	//--- check properties
	util.Assert(t, v.Properties != nil, "Properties property must be present")
	util.Equals(t, len(fun.OutNames), len(v.Properties), "Properties len")
	for i := 0; i < len(v.Properties); i++ {
		util.Equals(t, fun.OutNames[i], v.Properties[i].Name, "Properties[].Name")
		util.Equals(t, fun.OutJSONTypes[i], v.Properties[i].Type, "Properties[].Type")
	}

	//--- check links
	checkLink(t, v.Links[0], api.RelSelf, api.ContentTypeJSON, hTest.UrlBase+path)
	checkLink(t, v.Links[1], api.RelAlt, api.ContentTypeHTML, hTest.UrlBase+path+".html")
	itemsType := api.ContentTypeJSON
	if fun.IsGeometryFunction() {
		itemsType = api.ContentTypeGeoJSON
	}
	checkLink(t, v.Links[2], api.RelItems, itemsType, hTest.UrlBase+path+"/items")
}

// check if item is available and is not empty
func checkItem(t *testing.T, id int) {
	path := fmt.Sprintf("/collections/mock_a/items/%d", id)
	resp := hTest.DoRequest(t, path)
	body, _ := ioutil.ReadAll(resp.Body)

	// extracted from catalog_db.go
	type featureData struct {
		Type  string                 `json:"type"`
		ID    string                 `json:"id,omitempty"`
		Geom  *json.RawMessage       `json:"geometry"`
		Props map[string]interface{} `json:"properties"`
	}

	var v featureData
	errUnMarsh := json.Unmarshal(body, &v)
	util.Assert(t, errUnMarsh == nil, fmt.Sprintf("%v", errUnMarsh))

	util.Equals(t, "Feature", v.Type, "feature type")
	actId, _ := strconv.Atoi(v.ID)
	util.Equals(t, id, actId, "feature id")
	util.Equals(t, 4, len(v.Props), "# feature props")

}

// check if item is available and corresponds to json string
func checkItemEquals(t *testing.T, id int, jsonStr string) {
	path := fmt.Sprintf("/collections/mock_a/items/%d", id)
	resp := hTest.DoRequest(t, path)
	body, _ := ioutil.ReadAll(resp.Body)

	// extracted from catalog_db.go
	type featureData struct {
		Type  string                 `json:"type"`
		ID    string                 `json:"id,omitempty"`
		Geom  *json.RawMessage       `json:"geometry"`
		Props map[string]interface{} `json:"properties"`
	}

	var v featureData
	errUnMarsh := json.Unmarshal(body, &v)
	util.Assert(t, errUnMarsh == nil, fmt.Sprintf("%v", errUnMarsh))

	util.Equals(t, "Feature", v.Type, "feature type")
	actId, _ := strconv.Atoi(v.ID)
	util.Equals(t, id, actId, "feature id")
	util.Equals(t, 4, len(v.Props), "# feature props")

	var v_test featureData
	errUnMarshTest := json.Unmarshal(body, &v_test)
	util.Assert(t, errUnMarshTest == nil, fmt.Sprintf("%v", errUnMarshTest))
	util.Assert(t, reflect.DeepEqual(v, v_test), "Items are not equal")
}
