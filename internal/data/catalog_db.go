package data

/*
 Copyright 2019 - 2023 Crunchy Data Solutions, Inc.
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
	"sort"
	"strings"
	"time"

	"github.com/CrunchyData/pg_featureserv/internal/conf"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/log/logrusadapter"
	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"
)

// Constants
const (
	JSONTypeString       = "string"
	JSONTypeNumber       = "number"
	JSONTypeBoolean      = "boolean"
	JSONTypeJSON         = "json"
	JSONTypeBooleanArray = "boolean[]"
	JSONTypeStringArray  = "string[]"
	JSONTypeNumberArray  = "number[]"

	PGTypeBool      = "bool"
	PGTypeNumeric   = "numeric"
	PGTypeJSON      = "json"
	PGTypeGeometry  = "geometry"
	PGTypeTextArray = "_text"
)

type catalogDB struct {
	dbconn        *pgxpool.Pool
	tableIncludes map[string]string
	tableExcludes map[string]string
	tables        []*Table
	tableMap      map[string]*Table
	functions     []*Function
	functionMap   map[string]*Function
}

var isStartup bool
var isFunctionsLoaded bool
var instanceDB catalogDB

const fmtQueryStats = "Database query result: %v rows in %v"

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
	log.Infof("Connected as %s to %s @ %s", dbUser, dbName, dbHost)
	return db
}

func dbConfig() *pgxpool.Config {
	dbconf := conf.Configuration.Database.DbConnection
	// disallow blank config for safety
	if dbconf == "" {
		log.Fatal("Blank DbConnection is disallowed for security reasons")
	}

	dbconfig, err := pgxpool.ParseConfig(conf.Configuration.Database.DbConnection)
	if err != nil {
		log.Fatal(err)
	}
	// Read and parse connection lifetime
	dbPoolMaxLifeTime, errt := time.ParseDuration(conf.Configuration.Database.DbPoolMaxConnLifeTime)
	if errt != nil {
		log.Fatal(errt)
	}
	dbconfig.MaxConnLifetime = dbPoolMaxLifeTime

	// Read and parse max connections
	dbPoolMaxConns := conf.Configuration.Database.DbPoolMaxConns
	if dbPoolMaxConns > 0 {
		dbconfig.MaxConns = int32(dbPoolMaxConns)
	}

	// Read current log level and use one less-fine level
	dbconfig.ConnConfig.Logger = logrusadapter.NewLogger(log.New())
	levelString, _ := (log.GetLevel() - 1).MarshalText()
	pgxLevel, _ := pgx.LogLevelFromString(string(levelString))
	dbconfig.ConnConfig.LogLevel = pgxLevel

	return dbconfig
}

func (cat *catalogDB) SetIncludeExclude(includeList []string, excludeList []string) {
	//-- include schemas / tables
	cat.tableIncludes = make(map[string]string)
	for _, name := range includeList {
		nameLow := strings.ToLower(name)
		cat.tableIncludes[nameLow] = nameLow
	}
	//-- excluded schemas / tables
	cat.tableExcludes = make(map[string]string)
	for _, name := range excludeList {
		nameLow := strings.ToLower(name)
		cat.tableExcludes[nameLow] = nameLow
	}
}

func (cat *catalogDB) Close() {
	cat.dbconn.Close()
}

func (cat *catalogDB) Tables() ([]*Table, error) {
	cat.refreshTables(true)
	return cat.tables, nil
}

func (cat *catalogDB) TableReload(name string) {
	tbl, ok := cat.tableMap[name]
	if !ok {
		return
	}
	// load extent (which may change over time
	sqlExtentEst := sqlExtentEstimated(tbl)
	isExtentLoaded := cat.loadExtent(sqlExtentEst, tbl)
	if !isExtentLoaded {
		log.Debugf("Can't get estimated extent for %s", name)
		sqlExtentExact := sqlExtentExact(tbl)
		cat.loadExtent(sqlExtentExact, tbl)
	}
}

func (cat *catalogDB) loadExtent(sql string, tbl *Table) bool {
	var (
		xmin pgtype.Float8
		xmax pgtype.Float8
		ymin pgtype.Float8
		ymax pgtype.Float8
	)
	log.Debug("Extent query: " + sql)
	err := cat.dbconn.QueryRow(context.Background(), sql).Scan(&xmin, &ymin, &xmax, &ymax)
	if err != nil {
		log.Debugf("Error querying Extent for %s: %v", tbl.ID, err)
	}
	// no extent was read (perhaps a view...)
	if xmin.Status == pgtype.Null {
		return false
	}
	tbl.Extent.Minx = xmin.Float
	tbl.Extent.Miny = ymin.Float
	tbl.Extent.Maxx = xmax.Float
	tbl.Extent.Maxy = ymax.Float
	return true
}

func (cat *catalogDB) TableByName(name string) (*Table, error) {
	cat.refreshTables(false)
	tbl, ok := cat.tableMap[name]
	if !ok {
		return nil, nil
	}
	return tbl, nil
}

func (cat *catalogDB) TableFeatures(ctx context.Context, name string, param *QueryParam) ([]string, error) {
	tbl, err := cat.TableByName(name)
	if err != nil || tbl == nil {
		return nil, err
	}
	cols := param.Columns
	sql, argValues := sqlFeatures(tbl, param)
	log.Debug("Features query: " + sql)
	idColIndex := indexOfName(cols, tbl.IDColumn)

	features, err := readFeaturesWithArgs(ctx, cat.dbconn, sql, argValues, idColIndex, cols)
	return features, err
}

func (cat *catalogDB) TableFeature(ctx context.Context, name string, id string, param *QueryParam) (string, error) {
	tbl, err := cat.TableByName(name)
	if err != nil {
		return "", err
	}
	cols := param.Columns
	sql := sqlFeature(tbl, param)
	log.Debug("Feature query: " + sql)
	idColIndex := indexOfName(cols, tbl.IDColumn)

	//--- Add a SQL arg for the feature ID
	argValues := make([]interface{}, 0)
	argValues = append(argValues, id)
	features, err := readFeaturesWithArgs(ctx, cat.dbconn, sql, argValues, idColIndex, cols)

	if len(features) == 0 {
		return "", err
	}
	return features[0], nil
}

func (cat *catalogDB) CreateTableFeature(ctx context.Context, name string, feature Feature) error {
	tbl, err := cat.TableByName(name)
	if err != nil {
		return err
	}
	sql, argValues, err := sqlCreateFeature(tbl, feature)
	log.Debug("Create feature query: " + sql)
	result, err := cat.dbconn.Exec(ctx, sql, argValues...)
	if err != nil {
		return err
	}
	rows := result.RowsAffected()
	if rows != 1 {
		return fmt.Errorf("expected to affect 1 row, affected %d", rows)
	}
	return nil
}

func (cat *catalogDB) ReplaceTableFeature(ctx context.Context, name string, id string, feature Feature) error {
	tbl, err := cat.TableByName(name)
	if err != nil {
		return err
	}
	sql, argValues, err := sqlReplaceFeature(tbl, id, feature)
	log.Debug("Replace feature query: " + sql)
	result, err := cat.dbconn.Exec(ctx, sql, argValues...)
	if err != nil {
		return err
	}
	rows := result.RowsAffected()
	if rows != 1 {
		return fmt.Errorf("expected to affect 1 row, affected %d", rows)
	}
	return nil
}

func (cat *catalogDB) DeleteTableFeature(ctx context.Context, name string, id string) error {
	tbl, err := cat.TableByName(name)
	if err != nil {
		return err
	}
	sql, argValues := sqlDeleteFeature(tbl, id)
	log.Debug("Delete feature query: " + sql)
	result, err := cat.dbconn.Exec(ctx, sql, argValues...)
	if err != nil {
		return err
	}
	rows := result.RowsAffected()
	if rows != 1 {
		return fmt.Errorf("expected to affect 1 row, affected %d", rows)
	}
	return nil
}

func (cat *catalogDB) refreshTables(force bool) {
	// TODO: refresh on timed basis?
	if force || isStartup {
		cat.loadTables()
		isStartup = false
	}
}

func (cat *catalogDB) loadTables() {
	cat.tableMap = cat.readTables(cat.dbconn)
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

func (cat *catalogDB) readTables(db *pgxpool.Pool) map[string]*Table {
	log.Debugf("Load table catalog:\n%v", sqlTables)
	rows, err := db.Query(context.Background(), sqlTables)
	if err != nil {
		log.Fatal(err)
	}
	tables := make(map[string]*Table)
	for rows.Next() {
		tbl := scanTable(rows)
		if cat.isIncluded(tbl) {
			tables[tbl.ID] = tbl
		}
	}
	// Check for errors from iterating over rows.
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	rows.Close()
	return tables
}

func (cat *catalogDB) isIncluded(tbl *Table) bool {
	//--- if no includes defined, always include
	isIncluded := true
	if len(cat.tableIncludes) > 0 {
		isIncluded = isMatchSchemaTable(tbl, cat.tableIncludes)
	}
	isExcluded := false
	if len(cat.tableExcludes) > 0 {
		isExcluded = isMatchSchemaTable(tbl, cat.tableExcludes)
	}
	return isIncluded && !isExcluded
}

func isMatchSchemaTable(tbl *Table, list map[string]string) bool {
	schemaLow := strings.ToLower(tbl.Schema)
	if _, ok := list[schemaLow]; ok {
		return true
	}
	idLow := strings.ToLower(tbl.ID)
	if _, ok := list[idLow]; ok {
		return true
	}
	return false
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

//nolint:unused
func readFeatures(ctx context.Context, db *pgxpool.Pool, sql string, idColIndex int, propCols []string) ([]string, error) {
	return readFeaturesWithArgs(ctx, db, sql, nil, idColIndex, propCols)
}

//nolint:unused
func readFeaturesWithArgs(ctx context.Context, db *pgxpool.Pool, sql string, args []interface{}, idColIndex int, propCols []string) ([]string, error) {
	start := time.Now()
	rows, err := db.Query(ctx, sql, args...)
	if err != nil {
		log.Warnf("Error running Features query: %v", err)
		return nil, err
	}
	defer rows.Close()

	data, err := scanFeatures(ctx, rows, idColIndex, propCols)
	if err != nil {
		return data, err
	}
	log.Debugf(fmtQueryStats, len(data), time.Since(start))
	return data, nil
}

func scanFeatures(ctx context.Context, rows pgx.Rows, idColIndex int, propCols []string) ([]string, error) {
	// init features array to empty (not nil)
	var features []string = []string{}
	for rows.Next() {
		feature := scanFeature(rows, idColIndex, propCols)
		//log.Println(feature)
		features = append(features, feature)
	}
	// context check done outside rows loop,
	// because a long-running function might not produce any rows before timeout
	if err := ctx.Err(); err != nil {
		//log.Debugf("Context error scanning Features: %v", err)
		return features, err
	}
	// Check for errors from scanning rows.
	if err := rows.Err(); err != nil {
		log.Warnf("Error scanning rows for Features: %v", err)
		// TODO: return nil here ?
		return features, err
	}
	return features, nil
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
	//--- convert NULL to an empty string
	if vals[0] != nil {
		geom = vals[0].(string)
	} else {
		geom = ""
	}

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

// toJSONValue convert PG types to JSON values
func toJSONValue(value interface{}) interface{} {
	//fmt.Printf("toJSONValue: %v\n", reflect.TypeOf(value))
	switch v := value.(type) {
	case *pgtype.Numeric:
		var num float64
		// TODO: handle error
		v.AssignTo(&num) //nolint:errcheck
		return num
	case *pgtype.JSON:
		var jsonval string
		v.AssignTo(&jsonval) //nolint:errcheck
		return json.RawMessage(jsonval)
	case *pgtype.TextArray:
		var strarr []string
		v.AssignTo(&strarr) //nolint:errcheck
		return strarr
	case *pgtype.BoolArray:
		var valarr []bool
		v.AssignTo(&valarr) //nolint:errcheck
		return valarr
	case *pgtype.Int2Array:
		var numarr []int16
		v.AssignTo(&numarr) //nolint:errcheck
		return numarr
	case *pgtype.Int4Array:
		var numarr []int32
		v.AssignTo(&numarr) //nolint:errcheck
		return numarr
	case *pgtype.Int8Array:
		var numarr []int64
		v.AssignTo(&numarr) //nolint:errcheck
		return numarr
	case *pgtype.Float4Array:
		var numarr []float64
		v.AssignTo(&numarr) //nolint:errcheck
		return numarr
	case *pgtype.Float8Array:
		var numarr []float64
		v.AssignTo(&numarr) //nolint:errcheck
		return numarr
	case *pgtype.NumericArray:
		var numarr []float64
		v.AssignTo(&numarr) //nolint:errcheck
		return numarr
		// TODO: handle other conversions?
	}
	// for now all other values are returned  as is
	// this is only safe if the values are text!
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
	//fmt.Printf("toJSONTypeFromPG: %v\n", pgType)
	if strings.HasPrefix(pgType, "int") || strings.HasPrefix(pgType, "float") {
		return JSONTypeNumber
	}
	if strings.HasPrefix(pgType, "_int") || strings.HasPrefix(pgType, "_float") {
		return JSONTypeNumberArray
	}
	if strings.HasPrefix(pgType, "_bool") {
		return JSONTypeBooleanArray
	}
	switch pgType {
	case PGTypeNumeric:
		return JSONTypeNumber
	case PGTypeBool:
		return JSONTypeBoolean
	case PGTypeJSON:
		return JSONTypeJSON
	case PGTypeTextArray:
		return JSONTypeStringArray
	// hack to allow displaying geometry type
	case PGTypeGeometry:
		return PGTypeGeometry
	}
	// default is string
	// this forces conversion to text in SQL query
	return JSONTypeString
}

type featureData struct {
	Type  string                 `json:"type"`
	ID    string                 `json:"id,omitempty"`
	Geom  *json.RawMessage       `json:"geometry"`
	Props map[string]interface{} `json:"properties"`
}

func makeFeatureJSON(id string, geom string, props map[string]interface{}) string {
	//--- convert empty geom string to JSON null
	var geomRaw json.RawMessage
	if geom != "" {
		geomRaw = json.RawMessage(geom)
	}

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
