package mock_test

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
 Date     : October 2022
 Authors  : Jean-philippe Bazonnais (jean-philippe dot bazonnais at ign dot fr)
			Nicolas Revelant (nicolas dot revelant at ign dot fr)
*/

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/CrunchyData/pg_featureserv/internal/api"
	"github.com/CrunchyData/pg_featureserv/internal/conf"
	"github.com/CrunchyData/pg_featureserv/internal/data"
	"github.com/CrunchyData/pg_featureserv/internal/service"
	util "github.com/CrunchyData/pg_featureserv/internal/utiltest"
	log "github.com/sirupsen/logrus"
)

var hTest util.HttpTesting
var catalogMock *data.CatalogMock

// ...
func TestMain(m *testing.M) {
	conf.InitConfig("", false) // getting default configuration
	initCatMock()
	os.Exit(m.Run())
}

// ...
type MockTests struct {
	Test *testing.T
}

// ...
func TestRunnerHandlerMock(t *testing.T) {
	// t.Skip("skipping test.")

	// initialisation avant l'execution des tests
	beforeRun()

	t.Run("GET", func(t *testing.T) {
		m := MockTests{Test: t}
		m.TestRoot()
		m.TestGetFormatHandlingWithAcceptHeader()
		m.TestGetFormatHandlingSuffix()
		m.TestGetFormatHeaderAcceptUnsupportedMimeType()
		m.TestGetFormatSuffixSupersedesAcceptHeader()
		m.TestFeatureFormats()
		m.TestCollectionItem()
		m.TestCollectionItemsResponse()
		m.TestCollectionMissingItemsNotFound()
		m.TestCollectionItemPropertiesEmpty()
		m.TestCollectionNotFound()
		m.TestCollectionResponse()
		m.TestCollectionsResponse()
		m.TestFeatureNotFound()
	})
	t.Run("CACHE AND ETAGS", func(t *testing.T) {
		m := MockTests{Test: t}
		m.TestApiDecodeStrongEtag()
		m.TestLastModifiedMock()
		m.TestGetFeatureNoHeaderCheckEtag()
		m.TestGetFeatureHeaderIfNoneMatchWeakEtag()
		m.TestGetFeatureHeaderIfNoneMatchMalformedEtag()
		m.TestGetFeatureHeaderIfNoneMatchWithETagInCache()
		m.TestPutFeatureEtag()
		m.TestPatchFeatureEtag()
	})
	t.Run("GET - Params", func(t *testing.T) {
		m := MockTests{Test: t}
		m.TestRoot()
		m.TestBBox()
		m.TestBBoxInvalid()
		m.TestFilterB()
		m.TestFilterBD()
		m.TestFilterBDNone()
		m.TestFilterD()
		m.TestLimit()
		m.TestLimitInvalid()
		m.TestLimitZero()
		m.TestOffset()
		m.TestOffsetInvalid()
		m.TestProperties()
		m.TestPropertiesEmpty()
		m.TestPropertiesAll()
		m.TestQueryParamCase()
		m.TestSortBy()
		m.TestSortByAsc()
		m.TestSortByDesc()
		m.TestTransformInvalid()
		m.TestTransformValid()
	})
	t.Run("GET - Html", func(t *testing.T) {
		m := MockTests{Test: t}
		m.TestHTMLCollection()
		m.TestHTMLCollections()
		m.TestHTMLConformance()
		m.TestHTMLFunction()
		m.TestHTMLFunctions()
		m.TestHTMLItem()
		m.TestHTMLItems()
		m.TestHTMLRoot()
	})
	t.Run("GET - functions", func(t *testing.T) {
		m := MockTests{Test: t}
		m.TestFunctionJSON()
		m.TestFunctionMissingItemsNotFound()
		m.TestFunctionNotFound()
		m.TestFunctionsJSON()
	})
	// liste de tests sur la suppression des features
	t.Run("DELETE", func(t *testing.T) {
		beforeEachRun()
		m := MockTests{Test: t}
		m.TestApiContainsDeleteFeature()
		m.TestDeleteExistingFeature()
		m.TestDeleteFeatureErrorMalformedFeatureId()
		m.TestDeleteFeatureErrorUnknownCollection()
		m.TestDeleteFeatureErrorUnusedQueryParameters()
		m.TestDeleteUnknownFeature()
		afterEachRun()
	})
	t.Run("PUT", func(t *testing.T) {
		beforeEachRun()
		m := MockTests{Test: t}
		m.TestApiContainsMethodPut()
		m.TestGetCollectionReplaceSchema()
		m.TestReplaceFeatureMissingRequiredPropertiesFailure()
		m.TestReplaceFeatureOnlyGeomFailure()
		m.TestReplaceFeatureOnlyPropFailure()
		m.TestReplaceFeatureRequiredPropertiesSuccess()
		m.TestReplaceFeatureSuccess()
		afterEachRun()
	})
	t.Run("POST", func(t *testing.T) {
		beforeEachRun()
		m := MockTests{Test: t}
		m.TestApiContainsCollectionSchemas()
		m.TestApiContainsMethodPostFeature()
		m.TestGetCollectionCreateSchema()
		m.TestCreateFeature()
		afterEachRun()
	})
	t.Run("UPDATE", func(t *testing.T) {
		beforeEachRun()
		m := MockTests{Test: t}
		m.TestApiContainsMethodPatchFeature()
		m.TestGetCollectionUpdateSchema()
		m.TestUpdateFeatureSuccess()
		m.TestUpdateFeaturePartialSuccess()
		m.TestUpdateFeatureOnlyGeomSuccess()
		m.TestUpdateFeatureOnlyPropSuccess()
		m.TestUpdateFeaturePartialGeomFailure()
		afterEachRun()
	})

	// nettoyage après execution des tests
	afterRun()
}

// Run before all tests
func beforeRun() {
	log.Debug("beforeRun")
	// some stuff...
}

// Run after all tests
func afterRun() {
	log.Debug("afterRun")
	// some stuff...
}

// Run before each test
func beforeEachRun() {
	log.Debug("beforeEachRun")
	// re init catalog mock...
	initCatMock()
}

// Run after each test
func afterEachRun() {
	log.Debug("afterEachRun")
	// some stuff...
}

// ...
func initCatMock() {
	conf.Configuration.Database.AllowWrite = true
	catalogMock = data.CatMockInstance()
	service.SetCatalogInstance(catalogMock)

	hTest = util.MakeHttpTesting("http://test", "/pg_featureserv", "../../../assets", service.InitRouter("/pg_featureserv"))
	service.Initialize()
}

// ...
func checkCollection(tb testing.TB, coll *api.CollectionInfo, name string, title string) {
	util.Equals(tb, name, coll.Name, "Collection name")
	util.Equals(tb, title, coll.Title, "Collection title")

	path := "/collections/" + name
	checkLink(tb, coll.Links[0], api.RelSelf, api.ContentTypeJSON, hTest.UrlBase+path)
	checkLink(tb, coll.Links[1], api.RelAlt, api.ContentTypeHTML, hTest.UrlBase+path+".html")

	pathItems := path + "/items"
	checkLink(tb, coll.Links[2], api.RelItems, api.ContentTypeGeoJSON, hTest.UrlBase+pathItems)
}

// ...
func checkLink(tb testing.TB, link *api.Link, rel string, conType string, href string) {
	util.Equals(tb, rel, link.Rel, "Link rel")
	util.Equals(tb, conType, link.Type, "Link type")
	util.Equals(tb, href, link.Href, "Link href")
}

// ...
func checkFunctionSummary(tb testing.TB, v *api.FunctionSummary, fun *api.Function) {
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

// ...
func checkFunction(t *testing.T, fun *api.Function) {
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
		util.Equals(t, string(fun.OutJSONTypes[i]), v.Properties[i].Type, "Properties[].Type")
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
func checkItem(t *testing.T, id int) []byte {
	path := fmt.Sprintf("/collections/mock_a/items/%d", id)
	resp := hTest.DoRequest(t, path)
	body, _ := ioutil.ReadAll(resp.Body)

	var v api.GeojsonFeatureData
	errUnMarsh := json.Unmarshal(body, &v)
	util.Assert(t, errUnMarsh == nil, fmt.Sprintf("%v", errUnMarsh))

	util.Equals(t, "Feature", v.Type, "feature type")
	actId, _ := strconv.Atoi(v.ID)
	util.Equals(t, id, actId, "feature id")
	util.Equals(t, 4, len(v.Props), "# feature props")

	return body
}

// sends a GET request and checks the expected format (Content-Type header) from the response
func checkRouteResponseFormat(t *testing.T, url string, expectedContentType string) {

	resp := hTest.DoRequestStatus(t, url, http.StatusOK)
	respContentType := resp.Result().Header["Content-Type"][0]
	util.Assert(t, respContentType == expectedContentType, fmt.Sprintf("wrong Content-Type: %s", respContentType))
}

// sends a GET request with the specific Accept header provided, and checks the response received for :
// - the expected status
// - the expected format provided, or according to the initial Accept header (Content-Type header)
func checkRouteWithAcceptHeader(t *testing.T, url string, acceptValue string, expectedStatus int, expectedFormat string) []byte {

	var acceptHeader = make(http.Header)
	acceptHeader.Add("Accept", acceptValue)

	resp := hTest.DoRequestMethodStatus(t, "GET", url, nil, acceptHeader, expectedStatus)
	contentType := resp.Result().Header["Content-Type"][0]

	if expectedFormat != "" {
		util.Equals(t, expectedFormat, contentType, fmt.Sprintf("Content-Type: %s", contentType))
	} else {
		util.Assert(t, strings.Contains(acceptValue, contentType), fmt.Sprintf("Content-Type: %s", contentType))
	}

	body, _ := ioutil.ReadAll(resp.Body)
	return body
}
