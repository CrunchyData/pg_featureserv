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
	"github.com/CrunchyData/pg_featureserv/internal/data"
	"github.com/CrunchyData/pg_featureserv/internal/service"
	"github.com/CrunchyData/pg_featureserv/util"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/paulmach/orb/geojson"
)

var hTest util.HttpTesting
var db *pgxpool.Pool
var cat data.Catalog

// extracted from catalog_db.go
// TODO should be imported from catalog.go
type geojsonFeatureData struct {
	Type  string                 `json:"type"`
	ID    string                 `json:"id,omitempty"`
	Geom  *geojson.Geometry      `json:"geometry"`
	Props map[string]interface{} `json:"properties"`
}

func TestMain(m *testing.M) {
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
	util.Equals(t, 2, len(tables), "# table in DB")
}

func TestTestPropertiesAllFromDb(t *testing.T) {
	/*rr := hTest.DoRequest(t, "/collections/mock_a/items?limit=2")

	var v FeatureCollection
	errUnMarsh := json.Unmarshal(hTest.ReadBody(rr), &v)
	util.Assert(t, errUnMarsh == nil, fmt.Sprintf("%v", errUnMarsh))

	// Note that JSON numbers are read as float64
	util.Equals(t, 2, len(v.Features), "# features")
	util.Equals(t, 4, len(v.Features[0].Props), "feature 1 # properties")

	util.Equals(t, "propA", v.Features[0].Props["prop_a"], "feature 1 # property A")
	util.Equals(t, 1.0, v.Features[0].Props["prop_b"], "feature 1 # property B")
	util.Equals(t, "propC", v.Features[0].Props["prop_c"], "feature 1 # property C")
	util.Equals(t, 1.0, v.Features[0].Props["prop_d"], "feature 1 # property D")*/
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
	var features []string
	params := data.QueryParam{Limit: 100, Offset: 0, Crs: 4326}
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

// check if item is available and is not empty
// copy from service/handler_test.go
func checkItem(t *testing.T, id int) {
	path := fmt.Sprintf("/collections/mock_a/items/%d", id)
	resp := hTest.DoRequest(t, path)
	body, _ := ioutil.ReadAll(resp.Body)

	var v geojsonFeatureData
	errUnMarsh := json.Unmarshal(body, &v)
	util.Assert(t, errUnMarsh == nil, fmt.Sprintf("%v", errUnMarsh))

	util.Equals(t, "Feature", v.Type, "feature type")
	actId, _ := strconv.Atoi(v.ID)
	util.Equals(t, id, actId, "feature id")
	util.Equals(t, 4, len(v.Props), "# feature props")
}
