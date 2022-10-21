package mock_test

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
