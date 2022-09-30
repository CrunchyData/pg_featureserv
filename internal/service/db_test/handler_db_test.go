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
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/CrunchyData/pg_featureserv/internal/api"
	"github.com/CrunchyData/pg_featureserv/internal/conf"
	"github.com/CrunchyData/pg_featureserv/internal/data"
	"github.com/CrunchyData/pg_featureserv/internal/service"
	util "github.com/CrunchyData/pg_featureserv/internal/utiltest"
	"github.com/jackc/pgx/v4/pgxpool"
)

var hTest util.HttpTesting
var db *pgxpool.Pool
var cat data.Catalog

func TestMain(m *testing.M) {
	conf.Configuration.Database.AllowWrite = true

	db = util.CreateTestDb()
	defer util.CloseTestDb(db)

	cat = data.CatDBInstance()
	service.SetCatalogInstance(cat)

	hTest = util.MakeHttpTesting("http://test", "/pg_featureserv", service.InitRouter("/pg_featureserv"))
	service.Initialize()

	os.Exit(m.Run())
}

func TestProperDbInit(t *testing.T) {
	tables, _ := cat.Tables()
	util.Equals(t, 5, len(tables), "# tables in DB")
}

func TestPropertiesAllFromDbSimpleTable(t *testing.T) {
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
}

func TestPropertiesAllFromDbComplexTable(t *testing.T) {
	rr := hTest.DoRequest(t, "/collections/mock_multi/items?limit=5")

	var v api.FeatureCollection
	errUnMarsh := json.Unmarshal(hTest.ReadBody(rr), &v)
	util.Assert(t, errUnMarsh == nil, fmt.Sprintf("%v", errUnMarsh))

	// Note that JSON numbers are read as float64
	util.Equals(t, 5, len(v.Features), "# features")
	util.Equals(t, 8, len(v.Features[0].Props), "feature 1 # properties")

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

	expectJson := map[string]interface{}{
		"Name":   features[0].Props["prop_t"].(string),
		"IsDesc": features[0].Props["prop_i"].(int32)%2 == 1}
	util.Equals(t, expectJson, features[0].Props["prop_j"], "feature 1 # property json")
}

func TestReplaceFeatureSuccessDb(t *testing.T) {
	path := "/collections/mock_a/items/1"
	var header = make(http.Header)
	header.Add("Accept", api.ContentTypeSchemaPatchJSON)

	jsonStr := `{
		"type": "Feature",
		"id": "1",
		"geometry": {
			"type": "Point",
			"coordinates": [
			-120,
			40
			]
		},
		"properties": {
			"prop_a": "propA...",
			"prop_b": 1,
			"prop_c": "propC..."
		}
	}`

	hTest.DoRequestMethodStatus(t, "PUT", path, []byte(jsonStr), header, http.StatusNoContent)

	resp := hTest.DoRequestMethodStatus(t, "GET", path, []byte(""), header, http.StatusOK)
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))

	var jsonData map[string]interface{}
	err := json.Unmarshal(body, &jsonData)
	util.Assert(t, err == nil, fmt.Sprintf("%v", err))

	util.Equals(t, "1", jsonData["id"].(string), "feature ID")
	util.Equals(t, "Feature", jsonData["type"].(string), "feature Type")
	props := jsonData["properties"].(map[string]interface{})
	util.Equals(t, "propA...", props["prop_a"].(string), "feature value a")
	util.Equals(t, 1, int(props["prop_b"].(float64)), "feature value b")
	util.Equals(t, "propC...", props["prop_c"].(string), "feature value c")
	util.Equals(t, nil, props["prop_d"], "feature value d")
	geom := jsonData["geometry"].(map[string]interface{})
	util.Equals(t, "Point", geom["type"].(string), "feature Type")
	coordinate := geom["coordinates"].([]interface{})
	util.Equals(t, -120, int(coordinate[0].(float64)), "feature latitude")
	util.Equals(t, 40, int(coordinate[1].(float64)), "feature longitude")
}

func TestPartialUpdateFeatureDb(t *testing.T) {
	path := "/collections/mock_a/items/2"
	var header = make(http.Header)
	header.Add("Content-Type", api.ContentTypeSchemaPatchJSON)

	jsonStr := `{
		"type": "Feature",
		"id": "2",
		"geometry": {
			"type": "Point",
			"coordinates": [
			-120,
			40
			]
		},
		"properties": {
			"prop_a": "propA...",
			"prop_b": 2
		}
	}`

	resp := hTest.DoRequestMethodStatus(t, "PATCH", path, []byte(jsonStr), header, http.StatusNoContent)
	loc := resp.Header().Get("Location")

	util.Assert(t, len(loc) > 1, "Header location must not be empty")
	util.Equals(t, fmt.Sprintf("http://test/collections/mock_a/items/%d", 2), loc,
		"Header location must contain valid data")

	// check if it can be read
	feature := checkItem(t, "mock_a", 2)
	var jsonData map[string]interface{}
	errUnMarsh := json.Unmarshal(feature, &jsonData)
	util.Assert(t, errUnMarsh == nil, fmt.Sprintf("%v", errUnMarsh))

	util.Equals(t, "2", jsonData["id"].(string), "feature ID")
	util.Equals(t, "Feature", jsonData["type"].(string), "feature Type")
	props := jsonData["properties"].(map[string]interface{})
	util.Equals(t, "propA...", props["prop_a"].(string), "feature value a")
	util.Equals(t, 2, int(props["prop_b"].(float64)), "feature value b")
	util.Equals(t, "propC", props["prop_c"].(string), "feature value c")
	util.Equals(t, 2, int(props["prop_d"].(float64)), "feature value d")
	geom := jsonData["geometry"].(map[string]interface{})
	util.Equals(t, "Point", geom["type"].(string), "feature Type")
	coordinate := geom["coordinates"].([]interface{})
	util.Equals(t, -120, int(coordinate[0].(float64)), "feature latitude")
	util.Equals(t, 40, int(coordinate[1].(float64)), "feature longitude")

}

func TestDeleteFeatureDb(t *testing.T) {

	//--- retrieve max feature id before delete
	var features []*api.GeojsonFeatureData
	params := data.QueryParam{Limit: 100000, Offset: 0, Crs: 4326}
	features, _ = cat.TableFeatures(context.Background(), "mock_b", &params)

	featuresNumbersBefore := len(features)

	// -- do the request call but we have to force the catalogInstance to db during this operation
	hTest.DoDeleteRequestStatus(t, "/collections/mock_b/items/1", http.StatusNoContent)

	//--- retrieve max feature id after delete
	features, _ = cat.TableFeatures(context.Background(), "mock_b", &params)
	featuresNumbersAfter := len(features)

	util.Assert(t, featuresNumbersBefore-1 == featuresNumbersAfter, "# feature still in db/not deleted")

}

// check if item is available and is not empty
// copy from service/handler_test.go
func checkItem(t *testing.T, tableName string, id int) []byte {
	path := fmt.Sprintf("/collections/%v/items/%d", tableName, id)
	resp := hTest.DoRequest(t, path)
	body, _ := ioutil.ReadAll(resp.Body)

	var v api.GeojsonFeatureData
	errUnMarsh := json.Unmarshal(body, &v)
	util.Assert(t, errUnMarsh == nil, fmt.Sprintf("%v", errUnMarsh))

	util.Equals(t, "Feature", v.Type, "feature type")

	actId, _ := strconv.Atoi(v.ID)
	util.Equals(t, id, actId, "feature id")

	tbl, _ := service.CatalogInstance().TableByName(tableName)
	util.Equals(t, len(tbl.Columns)-1, len(v.Props), "# feature props")

	return body
}
