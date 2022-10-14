package db_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"testing"

	"github.com/CrunchyData/pg_featureserv/internal/api"
	"github.com/CrunchyData/pg_featureserv/internal/conf"
	"github.com/CrunchyData/pg_featureserv/internal/data"
	"github.com/CrunchyData/pg_featureserv/internal/service"
	util "github.com/CrunchyData/pg_featureserv/internal/utiltest"
	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"
)

// ...
var hTest util.HttpTesting
var db *pgxpool.Pool
var cat data.Catalog

// ...
type DbTests struct {
	Test *testing.T
}

// ...
func TestMain(m *testing.M) {
	conf.Configuration.Database.AllowWrite = true

	log.Debug("init : Db/Service")
	db = util.CreateTestDb()
	// defer util.CloseTestDb(db)

	cat = data.CatDBInstance()
	service.SetCatalogInstance(cat)

	hTest = util.MakeHttpTesting("http://test", "/pg_featureserv", "../../../assets", service.InitRouter("/pg_featureserv"))
	service.Initialize()

	os.Exit(m.Run())
}

// ...
func TestRunnerHandlerDb(t *testing.T) {

	// initialisation avant l'execution des tests
	beforeRun()

	t.Run("Init", func(t *testing.T) {
		beforeEachRun()
		test := DbTests{Test: t}
		test.TestProperDbInit()
		afterEachRun()
	})
	t.Run("GET", func(t *testing.T) {
		beforeEachRun()
		test := DbTests{Test: t}
		test.TestPropertiesAllFromDbSimpleTable()
		test.TestPropertiesAllFromDbComplexTable()
		afterEachRun()
	})
	// liste de tests sur la suppression des features
	t.Run("DELETE", func(t *testing.T) {
		beforeEachRun()
		test := DbTests{Test: t}
		test.TestDeleteFeatureDb()
		afterEachRun()
	})
	t.Run("PUT", func(t *testing.T) {
		beforeEachRun()
		test := DbTests{Test: t}
		test.TestSimpleReplaceFeatureSuccessDb()
		test.TestGetComplexCollectionReplaceSchema()
		test.TestReplaceComplexFeatureDb()
		afterEachRun()
	})
	t.Run("POST", func(t *testing.T) {
		beforeEachRun()
		test := DbTests{Test: t}
		test.TestCreateSimpleFeatureWithBadGeojsonInputDb()
		test.TestCreateSimpleFeatureDb()
		test.TestCreateComplexFeatureDb()
		test.TestGetComplexCollectionCreateSchema()
		afterEachRun()
	})
	t.Run("UPDATE", func(t *testing.T) {
		beforeEachRun()
		test := DbTests{Test: t}
		test.TestGetComplexCollectionUpdateSchema()
		test.TestUpdateComplexFeatureDb()
		test.TestUpdateSimpleFeatureDb()
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
	// close Db
	util.CloseTestDb(db)
}

// Run before each test
func beforeEachRun() {
	log.Debug("beforeEachRun")
	// drop and create table
	util.InsertSimpleDataset(db)
	util.InsertComplexDataset(db)
}

// Run after each test
func afterEachRun() {
	log.Debug("afterEachRun")
	// some stuff...
}

// Check if item is available and is not empty
// (copy from service/handler_test.go)
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
