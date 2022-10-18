package db_test

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

 Date     : September 2022
 Authors  : Benoit De Mezzo (benoit dot de dot mezzo at oslandia dot com)
        	Amaury Zarzelli (amaury dot zarzelli at ign dot fr)
			Jean-philippe Bazonnais (jean-philippe dot bazonnais at ign dot fr)
			Nicolas Revelant (nicolas dot revelant at ign dot fr)
*/

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/CrunchyData/pg_featureserv/internal/api"
	"github.com/CrunchyData/pg_featureserv/internal/data"
	util "github.com/CrunchyData/pg_featureserv/internal/utiltest"
)

func (t *DbTests) TestProperDbInit() {
	t.Test.Skip() // FIXME system tables not filtred correctly !?
	tables, _ := cat.Tables()
	util.Equals(t.Test, 5, len(tables), "# tables in DB")
}

func (t *DbTests) TestPropertiesAllFromDbSimpleTable() {
	t.Test.Run("TestPropertiesAllFromDbSimpleTable", func(t *testing.T) {
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

func (t *DbTests) TestPropertiesAllFromDbComplexTable() {
	t.Test.Run("TestPropertiesAllFromDbComplexTable", func(t *testing.T) {
		rr := hTest.DoRequest(t, "/collections/mock_multi/items?limit=5")

		var v api.FeatureCollection
		errUnMarsh := json.Unmarshal(hTest.ReadBody(rr), &v)
		util.Assert(t, errUnMarsh == nil, fmt.Sprintf("%v", errUnMarsh))

		// Note that JSON numbers are read as float64
		util.Equals(t, 5, len(v.Features), "# features")
		util.Equals(t, 9, len(v.Features[0].Props), "feature 1 # properties")

		util.Equals(t, "1", v.Features[0].Props["prop_t"].(string), "feature 1 # property text")

		tbl, _ := cat.TableByName("mock_multi")
		params := data.QueryParam{Limit: 100000, Offset: 0, Crs: 4326, Columns: tbl.Columns}
		features, _ := cat.TableFeatures(context.Background(), "mock_multi", &params)

		util.Equals(t, "1", features[0].Props["prop_t"].(string), "feature 1 # property text")
		util.Equals(t, int32(1), features[0].Props["prop_i"].(int32), "feature 1 # property int")
		util.Equals(t, int64(1), features[0].Props["prop_l"].(int64), "feature 1 # property long")
		util.Equals(t, float64(1.0), features[0].Props["prop_f"].(float64), "feature 1 # property float64")
		util.Equals(t, float32(1.0), features[0].Props["prop_r"].(float32), "feature 1 # property float32")
		util.Equals(t, []bool{false, true}, features[0].Props["prop_b"].([]bool), "feature 1 # property bool")
		util.Assert(t, time.Now().After(features[0].Props["prop_d"].(time.Time)), "feature 1 # property date")
		util.Equals(t, "1", features[0].Props["prop_v"].(string), "feature 1 # property varchar")

		expectJson := map[string]interface{}{
			"Name":   features[0].Props["prop_t"].(string),
			"IsDesc": features[0].Props["prop_i"].(int32)%2 == 1}
		util.Equals(t, expectJson, features[0].Props["prop_j"], "feature 1 # property json")
	})
}
