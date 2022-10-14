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

func (t *MockTests) TestFunctionsJSON() {
	t.Test.Run("TestFunctionsJSON", func(t *testing.T) {
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
	})
}

func (t *MockTests) TestFunctionJSON() {
	t.Test.Run("TestFunctionJSON", func(t *testing.T) {
		for _, fun := range catalogMock.FunctionDefs {
			//fun := catalogMock.FunctionDefs[1]
			checkFunction(t, fun)
		}
	})
}

func (t *MockTests) TestFunctionNotFound() {
	t.Test.Run("TestFunctionNotFound", func(t *testing.T) {
		hTest.DoRequestStatus(t, "/functions/missing", http.StatusNotFound)
	})
}

func (t *MockTests) TestFunctionMissingItemsNotFound() {
	t.Test.Run("TestFunctionMissingItemsNotFound", func(t *testing.T) {
		hTest.DoRequestStatus(t, "/functions/missing/items", http.StatusNotFound)
	})
}
