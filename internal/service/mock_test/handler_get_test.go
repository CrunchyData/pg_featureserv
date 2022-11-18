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
 Authors  : Benoit De Mezzo (benoit dot de dot mezzo at oslandia dot com)
 			Jean-philippe Bazonnais (jean-philippe dot bazonnais at ign dot fr)
 			Nicolas Revelant (nicolas dot revelant at ign dot fr)
*/

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"testing"

	"github.com/CrunchyData/pg_featureserv/internal/api"
	util "github.com/CrunchyData/pg_featureserv/internal/utiltest"
)

func (t *MockTests) TestRoot() {
	t.Test.Run("TestRoot", func(t *testing.T) {
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
	})
}

func (t *MockTests) TestGetFormatHandlingWithAcceptHeader() {
	t.Test.Run("TestGetFormatHandlingWithAcceptHeader", func(t *testing.T) {
		// This test targets the RequestedFormat() function from the net.go file

		// route / + "Accept: application/json"
		jsonBody := checkRouteWithAcceptHeader(t, "/", "application/json", http.StatusOK, api.ContentTypeJSON)
		jsonMap := new(map[string]interface{})
		errUnMarsh := json.Unmarshal(jsonBody, &jsonMap)
		util.Assert(t, errUnMarsh == nil, fmt.Sprintf("%v", errUnMarsh))

		checkRouteWithAcceptHeader(t, "/", "text/html", http.StatusOK, api.ContentTypeHTML)

		checkRouteWithAcceptHeader(t, "/api", "*/*", http.StatusOK, api.ContentTypeJSON)

		// Browser tests
		// -------------
		// route /api + default Accept header from Firefox 92
		firefoxAcceptHdr := "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8"
		checkRouteWithAcceptHeader(t, "/api", firefoxAcceptHdr, http.StatusOK, "")

		// route /api + default Accept header from Safari/Chrome
		chromeAcceptHdr := "text/html, application/xhtml+xml, image/jxr, */*"
		checkRouteWithAcceptHeader(t, "/api", chromeAcceptHdr, http.StatusOK, "")

		// route /api + default Accept header from Opera
		operaAcceptHdr := "text/html, application/xml;q=0.9, application/xhtml+xml, image/png, image/webp, image/jpeg, image/gif, image/x-xbitmap, */*;q=0.1"
		checkRouteWithAcceptHeader(t, "/api", operaAcceptHdr, http.StatusOK, "")

		// route /api + Accept header with a supported format present in the middle of the value
		messyHtmlAcceptHdr := "application/xhtml+xml,application/xml;q=0.9,image/avif,text/html,image/webp,*/*;q=0.8"
		checkRouteWithAcceptHeader(t, "/api", messyHtmlAcceptHdr, http.StatusOK, "")
	})
}

func (t *MockTests) TestGetFormatHandlingSuffix() {
	t.Test.Run("TestGetFormatHandlingSuffix", func(t *testing.T) {

		// checking supported suffixes HTML and JSON, and missing suffix
		checkRouteResponseFormat(t, "/home", api.ContentTypeJSON)
		checkRouteResponseFormat(t, "/home.html", api.ContentTypeHTML)
		checkRouteResponseFormat(t, "/home.json", api.ContentTypeJSON)
		checkRouteResponseFormat(t, "/index", api.ContentTypeJSON)
		checkRouteResponseFormat(t, "/index.html", api.ContentTypeHTML)
		checkRouteResponseFormat(t, "/index.json", api.ContentTypeJSON)
		checkRouteResponseFormat(t, "/api", api.ContentTypeJSON)
		checkRouteResponseFormat(t, "/api.html", api.ContentTypeHTML)
		checkRouteResponseFormat(t, "/api.json", api.ContentTypeJSON)
		checkRouteResponseFormat(t, "/collections", api.ContentTypeJSON)
		checkRouteResponseFormat(t, "/collections/mock_a", api.ContentTypeJSON)
		checkRouteResponseFormat(t, "/collections/mock_a.html", api.ContentTypeHTML)
		checkRouteResponseFormat(t, "/collections/mock_a.json", api.ContentTypeJSON)
		checkRouteResponseFormat(t, "/collections/mock_a/items", api.ContentTypeGeoJSON)
		checkRouteResponseFormat(t, "/collections/mock_a/items.html", api.ContentTypeHTML)
		checkRouteResponseFormat(t, "/collections/mock_a/items.json", api.ContentTypeGeoJSON)
		checkRouteResponseFormat(t, "/collections/mock_a/items/1", api.ContentTypeGeoJSON)
		checkRouteResponseFormat(t, "/collections/mock_a/items/1.html", api.ContentTypeHTML)
		checkRouteResponseFormat(t, "/collections/mock_a/items/1.json", api.ContentTypeGeoJSON)
		checkRouteResponseFormat(t, "/functions", api.ContentTypeJSON)
		checkRouteResponseFormat(t, "/functions.html", api.ContentTypeHTML)
		checkRouteResponseFormat(t, "/functions.json", api.ContentTypeJSON)
		checkRouteResponseFormat(t, "/functions/fun_a", api.ContentTypeJSON)
		checkRouteResponseFormat(t, "/functions/fun_a.html", api.ContentTypeHTML)
		checkRouteResponseFormat(t, "/functions/fun_a.json", api.ContentTypeJSON)
		// TODO : /functions/{id}/items
		checkRouteResponseFormat(t, "/conformance", api.ContentTypeJSON)
		checkRouteResponseFormat(t, "/conformance.html", api.ContentTypeHTML)
		checkRouteResponseFormat(t, "/conformance.json", api.ContentTypeJSON)

	})
}

func (t *MockTests) TestGetFormatHeaderAcceptUnsupportedMimeType() {
	t.Test.Run("TestGetFormatHeaderAcceptUnsupportedMimeType", func(t *testing.T) {

		gifAccept := "image/gif"
		xmlAccept := "application/xml"
		var dummyAcceptHdr = make(http.Header)
		dummyAcceptHdr.Add("Accept", "dummy/format")
		// Root
		checkRouteWithAcceptHeader(t, "/", gifAccept, http.StatusOK, api.ContentTypeJSON)
		checkRouteWithAcceptHeader(t, "/", xmlAccept, http.StatusOK, api.ContentTypeJSON)

		// Api
		checkRouteWithAcceptHeader(t, "/api", gifAccept, http.StatusOK, api.ContentTypeJSON)

		// Collections
		checkRouteWithAcceptHeader(t, "/collections", gifAccept, http.StatusOK, api.ContentTypeJSON)
		checkRouteWithAcceptHeader(t, "/collections/mock_a", gifAccept, http.StatusOK, api.ContentTypeJSON)

		// GET item(s)
		hTest.DoRequestMethodStatus(t, "GET", "/collections/mock_a/items", nil, dummyAcceptHdr, http.StatusNotAcceptable)
		hTest.DoRequestMethodStatus(t, "GET", "/collections/mock_a/items/1", nil, dummyAcceptHdr, http.StatusNotAcceptable)

	})
}

func (t *MockTests) TestGetFormatSuffixSupersedesAcceptHeader() {
	t.Test.Run("TestGetFormatSuffixSupersedesAcceptHeader", func(t *testing.T) {

		htmlAccept := "text/html"
		checkRouteWithAcceptHeader(t, "/api", htmlAccept, http.StatusOK, api.ContentTypeHTML)
		checkRouteWithAcceptHeader(t, "/api.json", htmlAccept, http.StatusOK, api.ContentTypeJSON)
		checkRouteWithAcceptHeader(t, "/api.html", htmlAccept, http.StatusOK, api.ContentTypeHTML)

	})
}

func (t *MockTests) TestFeatureFormats() {
	t.Test.Run("TestFeatureFormats", func(t *testing.T) {

		hTest.DoRequestStatus(t, "/collections/mock_a/items/1.dummyformat", http.StatusNotAcceptable)

		path := "/collections/mock_a/items/1"

		// From header Accept
		var header = make(http.Header)
		header.Add("Accept", "json")
		resp := hTest.DoRequestMethodStatus(t, "GET", path, nil, header, http.StatusOK)
		var geoJsonStruct api.GeojsonFeatureData
		errUnMarsh := json.Unmarshal(hTest.ReadBody(resp), &geoJsonStruct)
		util.Assert(t, errUnMarsh == nil, fmt.Sprintf("%v", errUnMarsh))

		// TODO HTML
		// TODO SVG
		// TODO TEXT

		// From URL extension
		resp2 := hTest.DoRequestStatus(t, "/collections/mock_a/items/1.json", http.StatusOK)
		errUnMarsh2 := json.Unmarshal(hTest.ReadBody(resp2), &geoJsonStruct)
		util.Assert(t, errUnMarsh == nil, fmt.Sprintf("%v", errUnMarsh2))

	})
}

func (t *MockTests) TestCollectionsResponse() {
	t.Test.Run("TestCollectionsResponse", func(t *testing.T) {
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
	})
}

func (t *MockTests) TestCollectionResponse() {
	t.Test.Run("TestCollectionResponse", func(t *testing.T) {
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
			util.Equals(t, string(tbl.JSONTypes[i]), v.Properties[i].Type, "Properties[].Type")
			util.Equals(t, tbl.ColDesc[i], v.Properties[i].Description, "Properties[].Description")
		}

		checkLink(t, v.Links[0], api.RelSelf, api.ContentTypeJSON, hTest.UrlBase+path)
		checkLink(t, v.Links[1], api.RelAlt, api.ContentTypeHTML, hTest.UrlBase+path+".html")
		checkLink(t, v.Links[2], api.RelItems, api.ContentTypeGeoJSON, hTest.UrlBase+path+"/items")
	})
}

func (t *MockTests) TestCollectionItemsResponse() {
	t.Test.Run("TestCollectionItemsResponse", func(t *testing.T) {
		path := "/collections/mock_a/items"
		resp := hTest.DoRequest(t, path)
		body, _ := ioutil.ReadAll(resp.Body)

		var v api.FeatureCollection
		errUnMarsh := json.Unmarshal(body, &v)
		util.Assert(t, errUnMarsh == nil, fmt.Sprintf("%v", errUnMarsh))

		util.Equals(t, int(catalogMock.TableSize("mock_a")), len(v.Features), "# features")
		checkLink(t, v.Links[0], api.RelSelf, api.ContentTypeJSON, hTest.UrlBase+path)
		checkLink(t, v.Links[1], api.RelAlt, api.ContentTypeHTML, hTest.UrlBase+path+".html")
	})
}

// check if item is available and is not empty
func (t *MockTests) TestCollectionItem() {
	t.Test.Run("TestCollectionItem", func(t *testing.T) {
		checkItem(t, 1)
	})
}

func (t *MockTests) TestCollectionItemPropertiesEmpty() {
	t.Test.Run("TestCollectionItemPropertiesEmpty", func(t *testing.T) {
		rr := hTest.DoRequest(t, "/collections/mock_a/items/1?properties=")

		var v api.GeojsonFeatureData
		errUnMarsh := json.Unmarshal(hTest.ReadBody(rr), &v)
		util.Assert(t, errUnMarsh == nil, fmt.Sprintf("%v", errUnMarsh))

		util.Equals(t, "Feature", v.Type, "feature type")
		util.Equals(t, 0, len(v.Props), "feature # properties")
	})
}

func (t *MockTests) TestCollectionNotFound() {
	t.Test.Run("TestCollectionNotFound", func(t *testing.T) {
		hTest.DoRequestStatus(t, "/collections/missing", http.StatusNotFound)
	})
}

func (t *MockTests) TestCollectionMissingItemsNotFound() {
	t.Test.Run("TestCollectionMissingItemsNotFound", func(t *testing.T) {
		hTest.DoRequestStatus(t, "/collections/missing/items", http.StatusNotFound)
	})
}

func (t *MockTests) TestFeatureNotFound() {
	t.Test.Run("TestFeatureNotFound", func(t *testing.T) {
		hTest.DoRequestStatus(t, "/collections/mock_a/items/999", http.StatusNotFound)
	})
}
