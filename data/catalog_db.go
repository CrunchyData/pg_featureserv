package data

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
	"context"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
	"text/template"
	"time"

	"github.com/CrunchyData/pg_featureserv/config"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/log/logrusadapter"
	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"
)

const (
	JSONTypeString  = "string"
	JSONTypeNumber  = "number"
	JSONTypeBoolean = "boolean"

	PGTypeBool     = "bool"
	PGTypeNumeric  = "numeric"
	PGTypeGeometry = "geometry"
)

type catalogDB struct {
	dbconn      *pgxpool.Pool
	tables      []*Table
	tableMap    map[string]*Table
	functions   []*Function
	functionMap map[string]*Function
}

var isStartup bool
var isFunctionsLoaded bool
var instanceDB catalogDB
var templateFeature *template.Template

func init() {
	isStartup = true
}

// CatDBInstance tbd
func CatDBInstance() Catalog {
	// TODO: make a singleton
	instanceDB = newCatalogDB()
	return &instanceDB
}

func newCatalogDB() catalogDB {
	conn := dbConnect()
	cat := catalogDB{
		dbconn: conn,
	}
	return cat
}

func dbConnect() *pgxpool.Pool {
	dbconfig := dbConfig()

	db, err := pgxpool.ConnectConfig(context.Background(), dbconfig)
	if err != nil {
		log.Fatal(err)
	}
	dbName := dbconfig.ConnConfig.Config.Database
	dbUser := dbconfig.ConnConfig.Config.User
	dbHost := dbconfig.ConnConfig.Config.Host
	log.Infof("Connected as '%s' to '%s' @ '%s'", dbUser, dbName, dbHost)
	return db
}

func dbConfig() *pgxpool.Config {
	dbconfig, err := pgxpool.ParseConfig(os.Getenv(config.AppConfig.EnvDBURL))
	if err != nil {
		log.Fatal(err)
	}
	// Read and parse connection lifetime
	dbPoolMaxLifeTime, errt := time.ParseDuration(config.Configuration.Database.DbPoolMaxConnLifeTime)
	if errt != nil {
		log.Fatal(errt)
	}
	dbconfig.MaxConnLifetime = dbPoolMaxLifeTime

	// Read and parse max connections
	dbPoolMaxConns := config.Configuration.Database.DbPoolMaxConns
	if dbPoolMaxConns > 0 {
		dbconfig.MaxConns = int32(dbPoolMaxConns)
	}

	// Read current log level and use one less-fine level
	// below that
	dbconfig.ConnConfig.Logger = logrusadapter.NewLogger(log.New())
	levelString, _ := (log.GetLevel() - 1).MarshalText()
	pgxLevel, _ := pgx.LogLevelFromString(string(levelString))
	dbconfig.ConnConfig.LogLevel = pgxLevel

	return dbconfig
}

func (cat *catalogDB) Close() {
	cat.dbconn.Close()
}

func (cat *catalogDB) Tables() ([]*Table, error) {
	cat.refreshTables(true)
	return cat.tables, nil
}

func (cat *catalogDB) TableByName(name string) (*Table, error) {
	cat.refreshTables(false)
	tbl, ok := cat.tableMap[name]
	if !ok {
		return nil, nil
	}
	return tbl, nil
}

func (cat *catalogDB) TableFeatures(name string, param *QueryParam) ([]string, error) {
	tbl, err := cat.TableByName(name)
	if err != nil || tbl == nil {
		return nil, err
	}
	cols := param.Columns
	sql := sqlFeatures(tbl, param)
	log.Debug("TableFeatures: " + sql)
	idColIndex := indexOfName(cols, tbl.IDColumn)

	features, err := readFeatures(cat.dbconn, sql, idColIndex, cols)
	return features, err
}

func (cat *catalogDB) TableFeature(name string, id string, param *QueryParam) (string, error) {
	tbl, err := cat.TableByName(name)
	if err != nil {
		return "", err
	}
	cols := param.Columns
	sql := sqlFeature(tbl, param)
	log.Debug(sql)
	idColIndex := indexOfName(cols, tbl.IDColumn)

	//--- Add a SQL arg for the feature ID
	argValues := make([]interface{}, 0)
	argValues = append(argValues, id)
	features, err := readFeaturesWithArgs(cat.dbconn, sql, argValues, idColIndex, cols)

	if len(features) == 0 {
		return "", err
	}
	return features[0], nil
}

func (cat *catalogDB) refreshTables(force bool) {
	// TODO: refresh on timed basis?
	if force || isStartup {
		cat.loadTables()
		isStartup = false
	}
}

func (cat *catalogDB) loadTables() {
	cat.tableMap = readTables(cat.dbconn)
	cat.tables = tablesSorted(cat.tableMap)
}

func tablesSorted(tableMap map[string]*Table) []*Table {
	// TODO: use database order instead of sorting here
	var lsort []*Table
	for key := range tableMap {
		lsort = append(lsort, tableMap[key])
	}
	sort.SliceStable(lsort, func(i, j int) bool {
		return lsort[i].Title < lsort[j].Title
	})
	return lsort
}

func readTables(db *pgxpool.Pool) map[string]*Table {
	log.Debug(sqlTables)
	rows, err := db.Query(context.Background(), sqlTables)
	if err != nil {
		log.Fatal(err)
	}
	tables := make(map[string]*Table)
	for rows.Next() {
		tbl := scanTable(rows)
		tables[tbl.ID] = tbl
	}
	// Check for errors from iterating over rows.
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	rows.Close()
	return tables
}

func scanTable(rows pgx.Rows) *Table {
	var (
		id, schema, table, description, geometryCol string
		srid                                        int
		geometryType, idColumn                      string
		props                                       pgtype.TextArray
	)

	err := rows.Scan(&id, &schema, &table, &description, &geometryCol,
		&srid, &geometryType, &idColumn, &props)
	if err != nil {
		log.Fatal(err)
	}

	// Use https://godoc.org/github.com/jackc/pgtype#TextArray
	// here to scan the text[][] map of attribute name/type
	// created in the query. It gets a little ugly demapping the
	// pgx TextArray type, but it is at least native handling of
	// the array. It's complex because of PgSQL ARRAY generality
	// really, no fault of pgx

	arrLen := 0
	arrStart := 0
	elmLen := 0
	if props.Status != pgtype.Null {
		arrLen = int(props.Dimensions[0].Length)
		arrStart = int(props.Dimensions[0].LowerBound - 1)
		elmLen = int(props.Dimensions[1].Length)
	}

	// TODO: query columns in table-defined order

	// Since Go map order is random, list columns in array
	columns := make([]string, arrLen)
	jsontypes := make([]string, arrLen)
	datatypes := make(map[string]string)
	colDesc := make([]string, arrLen)

	for i := arrStart; i < arrLen; i++ {
		elmPos := i * elmLen
		name := props.Elements[elmPos].String
		datatype := props.Elements[elmPos+1].String
		columns[i] = name
		datatypes[name] = datatype
		jsontypes[i] = toJSONTypeFromPG(datatype)
		colDesc[i] = props.Elements[elmPos+2].String
	}

	// Synthesize a title for now
	title := id
	// synthesize a description if none provided
	if description == "" {
		description = fmt.Sprintf("Data for table %v", id)
	}

	return &Table{
		ID:             id,
		Schema:         schema,
		Table:          table,
		Title:          title,
		Description:    description,
		GeometryColumn: geometryCol,
		Srid:           srid,
		GeometryType:   geometryType,
		IDColumn:       idColumn,
		Columns:        columns,
		DbTypes:        datatypes,
		JSONTypes:      jsontypes,
		ColDesc:        colDesc,
	}
}

//=================================================

func readFeatures(db *pgxpool.Pool, sql string, idColIndex int, propCols []string) ([]string, error) {
	return readFeaturesWithArgs(db, sql, nil, idColIndex, propCols)
}

func readFeaturesWithArgs(db *pgxpool.Pool, sql string, args []interface{}, idColIndex int, propCols []string) ([]string, error) {
	rows, err := db.Query(context.Background(), sql, args...)
	if err != nil {
		log.Warnf("Error running Features query: %v", err)
		return nil, err
	}
	return scanFeatures(rows, idColIndex, propCols), nil
}

func scanFeatures(rows pgx.Rows, idColIndex int, propCols []string) []string {
	var features []string
	for rows.Next() {
		feature := scanFeature(rows, idColIndex, propCols)
		//log.Println(feature)
		features = append(features, feature)
	}
	// Check for errors from iterating over rows.
	if err := rows.Err(); err != nil {
		log.Warnf("Error scanning rows for Features: %v", err)
	}
	rows.Close()
	return features
}

func scanFeature(rows pgx.Rows, idColIndex int, propNames []string) string {
	var id, geom string
	vals, err := rows.Values()
	if err != nil {
		log.Warnf("Error scanning row for Feature: %v", err)
		return ""
	}
	//fmt.Println(vals)
	//--- geom value is expected to be a GeoJSON string
	geom = vals[0].(string)
	propOffset := 1
	if idColIndex >= 0 {
		id = fmt.Sprintf("%v", vals[idColIndex+propOffset])
	}

	//fmt.Println(geom)
	props := extractProperties(vals, propOffset, propNames)
	return makeFeatureJSON(id, geom, props)
}

func extractProperties(vals []interface{}, propOffset int, propNames []string) map[string]interface{} {
	props := make(map[string]interface{})
	for i, name := range propNames {
		// offset vals index by 2 to skip geom, id
		props[name] = toJSONValue(vals[i+propOffset])
		//fmt.Printf("%v: %v\n", name, val)
	}
	return props
}

// toJSONValue convert PG types to JSON values if needed
func toJSONValue(value interface{}) interface{} {
	switch v := value.(type) {
	case *pgtype.Numeric:
		var num float64
		// TODO: handle error
		v.AssignTo(&num)
		return num
		// TODO: handle other conversions?
	}
	// for now all other values are returned  as is
	return value
}

func toJSONTypeFromPGArray(pgTypes []string) []string {
	jsonTypes := make([]string, len(pgTypes))
	for i, pgType := range pgTypes {
		jsonTypes[i] = toJSONTypeFromPG(pgType)
	}
	return jsonTypes
}

func toJSONTypeFromPG(pgType string) string {
	if strings.HasPrefix(pgType, "int") {
		return JSONTypeNumber
	}
	if strings.HasPrefix(pgType, "float") {
		return JSONTypeNumber
	}
	switch pgType {
	case PGTypeNumeric:
		return JSONTypeNumber
	case PGTypeBool:
		return JSONTypeBoolean
	// hack to allow displaying geometry type
	case PGTypeGeometry:
		return PGTypeGeometry
	}
	return JSONTypeString
}

type featureData struct {
	Type  string                 `json:"type"`
	ID    string                 `json:"id,omitempty"`
	Geom  *json.RawMessage       `json:"geometry"`
	Props map[string]interface{} `json:"properties"`
}

func makeFeatureJSON(id string, geom string, props map[string]interface{}) string {
	geomRaw := json.RawMessage(geom)
	featData := featureData{
		Type:  "Feature",
		ID:    id,
		Geom:  &geomRaw,
		Props: props,
	}
	json, err := json.Marshal(featData)
	if err != nil {
		log.Errorf("Error marshalling feature into JSON: %v", err)
		return ""
	}
	jsonStr := string(json)
	//fmt.Println(jsonStr)
	return jsonStr
}

// indexOfName finds the index of a name in an array of names
// It returns the index or -1 if not found
func indexOfName(names []string, name string) int {
	for i, nm := range names {
		if nm == name {
			return i
		}
	}
	return -1
}
