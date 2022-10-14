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
	"strings"
	"testing"
)

// For now these just test that the template executes correctly
// correctness/completess of HTML is not tested
func (t *MockTests) TestHTMLRoot() {
	t.Test.Run("TestHTMLRoot", func(t *testing.T) {
		hTest.DoRequest(t, "/index.html")
	})
}
func (t *MockTests) TestHTMLConformance() {
	t.Test.Run("TestHTMLConformance", func(t *testing.T) {
		hTest.DoRequest(t, "/conformance.html")
	})
}
func (t *MockTests) TestHTMLCollections() {
	t.Test.Run("TestHTMLCollections", func(t *testing.T) {
		hTest.DoRequest(t, "/collections.html")
	})
}
func (t *MockTests) TestHTMLCollection() {
	t.Test.Run("TestHTMLCollection", func(t *testing.T) {
		hTest.DoRequest(t, "/collections/mock_a.html")
	})
}
func (t *MockTests) TestHTMLItems() {
	t.Test.Run("TestHTMLItems", func(t *testing.T) {
		hTest.DoRequest(t, "/collections/mock_a/items.html")
	})
}
func (t *MockTests) TestHTMLItem() {
	t.Test.Run("TestHTMLItem", func(t *testing.T) {
		hTest.DoRequest(t, "/collections/mock_a/items/1.html")
	})
}
func (t *MockTests) TestHTMLFunctions() {
	t.Test.Run("TestHTMLFunctions", func(t *testing.T) {
		rr := hTest.DoRequest(t, "/functions.html")
		for _, fun := range catalogMock.FunctionDefs {
			if !strings.Contains(rr.Body.String(), "http://test/functions/"+fun.Name+".json") {
				t.Errorf("Functions response should contain reference to " + fun.Name + ".json")
			}
		}
	})
}
func (t *MockTests) TestHTMLFunction() {
	t.Test.Run("TestHTMLFunction", func(t *testing.T) {
		hTest.DoRequest(t, "/functions/fun_a.html")
	})
}
