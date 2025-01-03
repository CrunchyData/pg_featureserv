package data

/*
 Copyright 2019 - 2025 Crunchy Data Solutions, Inc.
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
	"fmt"
	"time"

	"github.com/CrunchyData/pg_featureserv/internal/conf"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"
)

// FunctionIDColumnName is the name for a function-supplied ID
const FunctionIDColumnName = "id"

const SchemaPostGISFTW = "postgisftw"

func (cat *catalogDB) Functions() ([]*Function, error) {
	cat.refreshFunctions(true)
	return cat.functions, nil
}

func (cat *catalogDB) FunctionByName(name string) (*Function, error) {
	cat.refreshFunctions(false)
	fn, ok := cat.functionMap[name]
	if !ok {
		return nil, nil
	}
	return fn, nil
}

func (cat *catalogDB) refreshFunctions(force bool) {
	// TODO: refresh on timed basis?
	if force || !isFunctionsLoaded {
		cat.loadFunctions()
	}
	isFunctionsLoaded = true
}

func (cat *catalogDB) loadFunctions() {
	cat.functions, cat.functionMap = readFunctionDefs(cat.dbconn, conf.Configuration.Database.FunctionIncludes)
}

func readFunctionDefs(db *pgxpool.Pool, funSchemas []string) ([]*Function, map[string]*Function) {
	sql := sqlFunctions(funSchemas)
	log.Debugf("Load function catalog:\n%v", sql)
	rows, err := db.Query(context.Background(), sql)
	if err != nil {
		log.Fatal(err)
	}
	var functions []*Function
	functionMap := make(map[string]*Function)
	for rows.Next() {
		fn := scanFunctionDef(rows)
		// TODO: for now only show geometry functions
		//if fn.IsGeometryFunction() {
		functions = append(functions, fn)
		functionMap[fn.ID] = fn
		//}
	}
	// Check for errors from iterating over rows.
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	rows.Close()
	return functions, functionMap
}

func scanFunctionDef(rows pgx.Rows) *Function {
	var (
		id, schema, name, description                              string
		inNamesTA, inTypesTA, inDefaultsTA, outNamesTA, outTypesTA pgtype.TextArray
	)

	err := rows.Scan(&id, &schema, &name, &description,
		&inNamesTA, &inTypesTA, &inDefaultsTA, &outNamesTA, &outTypesTA)
	if err != nil {
		log.Fatalf("Error reading function catalog: %v", err)
	}

	inNames := toArray(inNamesTA)
	inTypes := toArray(inTypesTA)
	inDefaultsDB := toArray(inDefaultsTA)
	numNoDefault := len(inNames) - len(inDefaultsDB)
	inDefaults := extendLeft(inDefaultsDB, len(inNames))
	outNames := toArray(outNamesTA)
	outTypes := toArray(outTypesTA)
	outJSONTypes := toJSONTypeFromPGArray(outTypes)

	inTypeMap := make(map[string]string)
	addTypes(inTypeMap, inNames, inTypes)

	datatypes := make(map[string]string)
	addTypes(datatypes, inNames, inTypes)
	addTypes(datatypes, outNames, outTypes)

	// synthesize a description if none provided
	if description == "" {
		description = fmt.Sprintf("The function %v", id)
	}

	geomCol := geometryColumn(outNames, datatypes)

	funDef := Function{
		ID:             id,
		Schema:         schema,
		Name:           name,
		Description:    description,
		InNames:        inNames,
		InDbTypes:      inTypes,
		InTypeMap:      inTypeMap,
		InDefaults:     inDefaults,
		NumNoDefault:   numNoDefault,
		OutNames:       outNames,
		OutDbTypes:     outTypes,
		OutJSONTypes:   outJSONTypes,
		Types:          datatypes,
		GeometryColumn: geomCol,
	}
	//fmt.Printf("DEBUG: Function definitions: %v\n", funDef)
	return &funDef
}
func addTypes(typeMap map[string]string, names []string, types []string) {
	for i, name := range names {
		typeMap[name] = types[i]
	}
}
func geometryColumn(names []string, types map[string]string) string {
	// TODO: extract from outNames, outTypes
	for _, name := range names {
		if types[name] == PGTypeGeometry {
			return name
		}
	}
	return ""
}
func toArray(ta pgtype.TextArray) []string {
	arrLen := 0
	arrStart := 0
	if len(ta.Dimensions) > 0 {
		arrLen = int(ta.Dimensions[0].Length)
		arrStart = int(ta.Dimensions[0].LowerBound - 1)
	}

	arr := make([]string, arrLen)

	for i := arrStart; i < arrLen; i++ {
		val := ta.Elements[i].String
		arr[i] = val
	}
	return arr
}

// extendLeft extends an array to have given size, with original contents right-aligned
func extendLeft(arr []string, size int) []string {
	if size <= len(arr) {
		return arr
	}
	// create array of requested size and right-justify input
	arr2 := make([]string, size)
	offset := size - len(arr)
	for i := 0; i < len(arr); i++ {
		arr2[i+offset] = arr[i]
	}
	return arr2
}

func (cat *catalogDB) FunctionFeatures(ctx context.Context, name string, args map[string]string, param *QueryParam) ([]string, error) {
	fn, err := cat.FunctionByName(name)
	if err != nil || fn == nil {
		return nil, err
	}
	errArg := checkArgsValid(fn, args)
	if errArg != nil {
		log.Debug("ERROR: " + errArg.Error())
		return nil, errArg
	}
	propCols := removeNames(param.Columns, fn.GeometryColumn, "")
	idColIndex := indexOfName(propCols, FunctionIDColumnName)
	sql, argValues := sqlGeomFunction(fn, args, propCols, param)
	log.Debugf("Function features query: %v", sql)
	log.Debugf("Function %v Args: %v", name, argValues)
	features, err := readFeaturesWithArgs(ctx, cat.dbconn, sql, argValues, idColIndex, propCols)
	return features, err
}

func (cat *catalogDB) FunctionData(ctx context.Context, name string, args map[string]string, param *QueryParam) ([]map[string]interface{}, error) {
	fn, err := cat.FunctionByName(name)
	if err != nil || fn == nil {
		return nil, err
	}
	errArg := checkArgsValid(fn, args)
	if errArg != nil {
		return nil, errArg
	}
	propCols := param.Columns
	sql, argValues := sqlFunction(fn, args, propCols, param)
	log.Debugf("Function data query: %v", sql)
	log.Debugf("Function %v Args: %v", name, argValues)
	data, err := readDataWithArgs(ctx, cat.dbconn, propCols, sql, argValues)
	return data, err
}

// inputArgs extracts function arguments from any provided in the query parameters
// Arg values are stored as strings, and relies on Postgres to convert
// to the actual types required by the function
func checkArgsValid(fn *Function, args map[string]string) error {
	for argName := range args {
		if _, ok := fn.InTypeMap[argName]; !ok {
			return fmt.Errorf("'%v' is not a parameter of function '%v'", argName, fn.ID)
		}
	}
	return nil
}

func removeNames(names []string, ex1 string, ex2 string) []string {
	var newNames []string
	for _, name := range names {
		if name != ex1 && name != ex2 {
			newNames = append(newNames, name)
		}
	}
	return newNames
}

func readDataWithArgs(ctx context.Context, db *pgxpool.Pool, propCols []string, sql string, args []interface{}) ([]map[string]interface{}, error) {
	start := time.Now()
	rows, err := db.Query(context.Background(), sql, args...)
	if err != nil {
		log.Warnf("Error running Data query: %v", err)
		return nil, err
	}
	defer rows.Close()
	data := scanData(ctx, rows, propCols)
	log.Debugf(fmtQueryStats, len(data), time.Since(start))
	return data, nil
}

func scanData(ctx context.Context, rows pgx.Rows, propCols []string) []map[string]interface{} {
	// init data array to empty (not nil)
	var data []map[string]interface{} = []map[string]interface{}{}
	for rows.Next() {
		obj := scanDataRow(rows, true, propCols)
		//log.Println(feature)
		data = append(data, obj)
	}
	// context check done outside rows loop,
	// because a long-running function might not produce any rows before timeout
	if err := ctx.Err(); err != nil {
		//log.Debugf("Context error scanning Features: %v", err)
		return data
	}
	// Check for errors from iterating over rows.
	if err := rows.Err(); err != nil {
		log.Warnf("Error scanning Data rows: %v", err)
		// TODO: return nil here ?
	}
	return data
}

func scanDataRow(rows pgx.Rows, hasID bool, propNames []string) map[string]interface{} {
	vals, err := rows.Values()
	if err != nil {
		log.Warnf("Error getting Data row values: %v", err)
		return nil
	}
	//fmt.Println(vals)

	//fmt.Println(geom)
	props := extractProperties(vals, 0, propNames)
	return props
}
