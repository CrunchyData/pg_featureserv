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
	"net/http"

	"testing"

	"github.com/CrunchyData/pg_featureserv/internal/api"
	util "github.com/CrunchyData/pg_featureserv/internal/utiltest"
)

func (t *MockTests) TestFilterB() {
	t.Test.Run("TestFilterB", func(t *testing.T) {
		rr := hTest.DoRequest(t, "/collections/mock_a/items?prop_b=1")

		var v api.FeatureCollection
		errUnMarsh := json.Unmarshal(hTest.ReadBody(rr), &v)
		util.Assert(t, errUnMarsh == nil, fmt.Sprintf("%v", errUnMarsh))

		util.Equals(t, 1, len(v.Features), "# features")
	})
}

func (t *MockTests) TestFilterD() {
	t.Test.Run("TestFilterD", func(t *testing.T) {
		rr := hTest.DoRequest(t, "/collections/mock_c/items?prop_d=1")

		var v api.FeatureCollection
		errUnMarsh := json.Unmarshal(hTest.ReadBody(rr), &v)
		util.Assert(t, errUnMarsh == nil, fmt.Sprintf("%v", errUnMarsh))

		util.Equals(t, 10, len(v.Features), "# features")
	})
}

func (t *MockTests) TestFilterBD() {
	t.Test.Run("TestFilterBD", func(t *testing.T) {
		rr := hTest.DoRequest(t, "/collections/mock_c/items?prop_b=2&prop_d=2")

		var v api.FeatureCollection
		errUnMarsh := json.Unmarshal(hTest.ReadBody(rr), &v)
		util.Assert(t, errUnMarsh == nil, fmt.Sprintf("%v", errUnMarsh))

		util.Equals(t, 1, len(v.Features), "# features")
	})
}

func (t *MockTests) TestFilterBDNone() {
	t.Test.Run("TestFilterBDNone", func(t *testing.T) {
		rr := hTest.DoRequest(t, "/collections/mock_c/items?prop_b=1&prop_d=2")

		var v api.FeatureCollection
		errUnMarsh := json.Unmarshal(hTest.ReadBody(rr), &v)
		util.Assert(t, errUnMarsh == nil, fmt.Sprintf("%v", errUnMarsh))

		util.Equals(t, 0, len(v.Features), "# features")
	})
}

func (t *MockTests) TestSortBy() {
	t.Test.Run("TestSortBy", func(t *testing.T) {
		rr := hTest.DoRequest(t, "/collections/mock_a/items?sortby=prop_b")

		var v api.FeatureCollection
		errUnMarsh := json.Unmarshal(hTest.ReadBody(rr), &v)
		util.Assert(t, errUnMarsh == nil, fmt.Sprintf("%v", errUnMarsh))

		util.Equals(t, int(catalogMock.TableSize("mock_a")), len(v.Features), "# features")
	})
}

func (t *MockTests) TestSortByDesc() {
	t.Test.Run("TestSortByDesc", func(t *testing.T) {
		rr := hTest.DoRequest(t, "/collections/mock_a/items?sortby=-prop_b")

		var v api.FeatureCollection
		errUnMarsh := json.Unmarshal(hTest.ReadBody(rr), &v)
		util.Assert(t, errUnMarsh == nil, fmt.Sprintf("%v", errUnMarsh))

		util.Equals(t, int(catalogMock.TableSize("mock_a")), len(v.Features), "# features")
	})
}

func (t *MockTests) TestSortByAsc() {
	t.Test.Run("TestSortByAsc", func(t *testing.T) {
		rr := hTest.DoRequest(t, "/collections/mock_a/items?sortby=+prop_b")

		var v api.FeatureCollection
		errUnMarsh := json.Unmarshal(hTest.ReadBody(rr), &v)
		util.Assert(t, errUnMarsh == nil, fmt.Sprintf("%v", errUnMarsh))

		util.Equals(t, int(catalogMock.TableSize("mock_a")), len(v.Features), "# features")
	})
}

func (t *MockTests) TestLimit() {
	t.Test.Run("TestLimit", func(t *testing.T) {
		rr := hTest.DoRequest(t, "/collections/mock_a/items?limit=3")

		var v api.FeatureCollection
		errUnMarsh := json.Unmarshal(hTest.ReadBody(rr), &v)
		util.Assert(t, errUnMarsh == nil, fmt.Sprintf("%v", errUnMarsh))

		util.Equals(t, 3, len(v.Features), "# features")
		util.Equals(t, "1", v.Features[0].ID, "feature 1 id")
		util.Equals(t, "2", v.Features[1].ID, "feature 2 id")
		util.Equals(t, "3", v.Features[2].ID, "feature 3 id")
	})
}

func (t *MockTests) TestLimitZero() {
	t.Test.Run("TestLimitZero", func(t *testing.T) {
		rr := hTest.DoRequest(t, "/collections/mock_a/items?limit=0")

		var v api.FeatureCollection
		errUnMarsh := json.Unmarshal(hTest.ReadBody(rr), &v)
		util.Assert(t, errUnMarsh == nil, fmt.Sprintf("%v", errUnMarsh))

		util.Equals(t, "FeatureCollection", v.Type, "type FeatureCollection")
		util.Equals(t, 0, len(v.Features), "# features")
	})
}

func (t *MockTests) TestLimitInvalid() {
	t.Test.Run("TestLimitInvalid", func(t *testing.T) {
		hTest.DoRequestStatus(t, "/collections/mock_a/items?limit=x", http.StatusBadRequest)
	})
}

func (t *MockTests) TestQueryParamCase() {
	t.Test.Run("TestQueryParamCase", func(t *testing.T) {
		rr := hTest.DoRequest(t, "/collections/mock_a/items?LIMIT=2&Offset=4")

		var v api.FeatureCollection
		errUnMarsh := json.Unmarshal(hTest.ReadBody(rr), &v)
		util.Assert(t, errUnMarsh == nil, fmt.Sprintf("%v", errUnMarsh))

		util.Equals(t, 2, len(v.Features), "# features")
		util.Equals(t, "5", v.Features[0].ID, "feature 5 id")
		util.Equals(t, "6", v.Features[1].ID, "feature 6 id")
	})
}

func (t *MockTests) TestOffset() {
	t.Test.Run("TestOffset", func(t *testing.T) {
		rr := hTest.DoRequest(t, "/collections/mock_a/items?limit=2&offset=4")

		var v api.FeatureCollection
		errUnMarsh := json.Unmarshal(hTest.ReadBody(rr), &v)
		util.Assert(t, errUnMarsh == nil, fmt.Sprintf("%v", errUnMarsh))

		util.Equals(t, 2, len(v.Features), "# features")
		util.Equals(t, "5", v.Features[0].ID, "feature 5 id")
		util.Equals(t, "6", v.Features[1].ID, "feature 6 id")
	})
}

func (t *MockTests) TestOffsetInvalid() {
	t.Test.Run("TestOffsetInvalid", func(t *testing.T) {
		hTest.DoRequestStatus(t, "/collections/mock_a/items?offset=x", http.StatusBadRequest)
	})
}

func (t *MockTests) TestTransformValid() {
	t.Test.Run("TestTransformValid", func(t *testing.T) {
		hTest.DoRequest(t, "/collections/mock_a/items?transform=centroid")
		hTest.DoRequest(t, "/collections/mock_a/items?transform=ST_centroid")
		hTest.DoRequest(t, "/collections/mock_a/items?transform=st_centroid")
		hTest.DoRequest(t, "/collections/mock_a/items?transform=pointonsurface")
		hTest.DoRequest(t, "/collections/mock_a/items?transform=pointonsurface|centroid")
	})
}

func (t *MockTests) TestTransformInvalid() {
	t.Test.Run("TestTransformInvalid", func(t *testing.T) {
		// envelope is not defined as a transform function
		hTest.DoRequestStatus(t, "/collections/mock_a/items?transform=envelope", http.StatusBadRequest)
		hTest.DoRequestStatus(t, "/collections/mock_a/items?transform=centroid|envelope", http.StatusBadRequest)
	})
}

func (t *MockTests) TestBBox() {
	t.Test.Run("TestBBox", func(t *testing.T) {
		hTest.DoRequest(t, "/collections/mock_a/items?bbox=1,2,3,4")
		// TODO: add some tests
	})
}

func (t *MockTests) TestBBoxInvalid() {
	t.Test.Run("TestBBoxInvalid", func(t *testing.T) {
		hTest.DoRequestStatus(t, "/collections/mock_a/items?bbox=1,2,3,x", http.StatusBadRequest)
	})
}

func (t *MockTests) TestProperties() {
	t.Test.Run("TestProperties", func(t *testing.T) {
		// Tests:
		// - property names are non-case-sensitive
		// - names are made unique (properties only include once)
		// - non-existing names are ignored
		rr := hTest.DoRequest(t, "/collections/mock_a/items?limit=2&properties=PROP_A,prop_C,prop_a,not_prop")

		var v api.FeatureCollection
		errUnMarsh := json.Unmarshal(hTest.ReadBody(rr), &v)
		util.Assert(t, errUnMarsh == nil, fmt.Sprintf("%v", errUnMarsh))

		util.Equals(t, 2, len(v.Features), "# features")
		util.Equals(t, 2, len(v.Features[0].Props), "feature 1 # properties")
		util.Equals(t, "propA", v.Features[0].Props["prop_a"], "feature 1 # property A")
		util.Equals(t, "propC", v.Features[0].Props["prop_c"], "feature 1 # property C")
	})
}

func (t *MockTests) TestPropertiesEmpty() {
	t.Test.Run("TestPropertiesEmpty", func(t *testing.T) {
		rr := hTest.DoRequest(t, "/collections/mock_a/items?limit=2&properties=")

		var v api.FeatureCollection
		errUnMarsh := json.Unmarshal(hTest.ReadBody(rr), &v)
		util.Assert(t, errUnMarsh == nil, fmt.Sprintf("%v", errUnMarsh))

		util.Equals(t, 2, len(v.Features), "# features")
		util.Equals(t, 0, len(v.Features[0].Props), "feature 1 # properties")
	})
}

// TestPropertiesAll tests that no properties parameter returns all props
func (t *MockTests) TestPropertiesAll() {
	t.Test.Run("TestPropertiesAll", func(t *testing.T) {
		rr := hTest.DoRequest(t, "/collections/mock_a/items?limit=2")

		var v api.FeatureCollection
		errUnMarsh := json.Unmarshal(hTest.ReadBody(rr), &v)
		util.Assert(t, errUnMarsh == nil, fmt.Sprintf("%v", errUnMarsh))

		// Note that JSON numbers are read as float64
		util.Equals(t, 2, len(v.Features), "# features")
		util.Equals(t, 4, len(v.Features[0].Props), "feature 1 # properties")

		util.Equals(t, "propA", v.Features[0].Props["prop_a"], "feature 1 # property A")
		util.Equals(t, 1.0, v.Features[0].Props["prop_b"], "feature 1 # property B")
		util.Equals(t, "propC", v.Features[0].Props["prop_c"], "feature 1 # property C")
		util.Equals(t, 1.0, v.Features[0].Props["prop_d"], "feature 1 # property D")
	})
}
