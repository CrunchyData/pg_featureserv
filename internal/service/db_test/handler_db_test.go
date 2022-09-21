package db_test

import (
	"context"
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
	util.Equals(t, 4, len(tables), "# tables in DB")
}

func TestPropertiesAllFromDb(t *testing.T) {
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

func TestCreateFeatureWithBadGeojsonInputDb(t *testing.T) {
	var header = make(http.Header)
	header.Add("Content-Type", "application/geo+json")

	jsonStr := `[{
		"id": 101,
		"name": "Test",
		"email": "test@test.com"
	      }, {
		"id": 102,
		"name": "Sample",
		"email": "sample@test.com"
	    }]`

	rr := hTest.DoRequestMethodStatus(t, "POST", "/collections/mock_a/items", []byte(jsonStr), header, http.StatusInternalServerError)

	util.Equals(t, http.StatusInternalServerError, rr.Code, "Should have failed")
	util.Assert(t, strings.Index(rr.Body.String(), fmt.Sprintf(api.ErrMsgCreateFeatureNotConform+"\n", "mock_a")) == 0, "Should have failed with not conform")
}

func TestCreateFeatureDb(t *testing.T) {
	var header = make(http.Header)
	header.Add("Content-Type", "application/geo+json")

	//--- retrieve max feature id before insert
	var features []*api.GeojsonFeatureData
	params := data.QueryParam{Limit: 100000, Offset: 0, Crs: 4326}
	features, _ = cat.TableFeatures(context.Background(), "mock_a", &params)
	maxIdBefore := len(features)

	//--- generate json from new object
	tables, _ := cat.Tables()
	var cols []string
	for _, t := range tables {
		if t.ID == "public.mock_a" {
			for _, c := range t.Columns {
				if c != "id" {
					cols = append(cols, c)
				}
			}
			break
		}
	}
	jsonStr := data.MakeFeatureMockPointAsJSON(99, 12, 34, cols)
	fmt.Println(jsonStr)

	// -- do the request call but we have to force the catalogInstance to db during this operation
	rr := hTest.DoPostRequest(t, "/collections/mock_a/items", []byte(jsonStr), header)

	loc := rr.Header().Get("Location")

	//--- retrieve max feature id after insert
	features, _ = cat.TableFeatures(context.Background(), "mock_a", &params)
	maxIdAfter := len(features)

	util.Assert(t, maxIdAfter > maxIdBefore, "# feature in db")
	util.Assert(t, len(loc) > 1, "Header location must not be empty")
	util.Equals(t, fmt.Sprintf("http://test/collections/mock_a/items/%d", maxIdAfter), loc,
		"Header location must contain valid data")

	// check if it can be read
	checkItem(t, maxIdAfter)
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
			"prop_c": "propC...",
			"prop_d": 1
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
	util.Equals(t, 1, int(props["prop_d"].(float64)), "feature value d")
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
	feature := checkItem(t, 2)
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
