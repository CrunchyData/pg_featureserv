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
	"fmt"

	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"
)

// FunctionIDColumnName is the name for a function-supplied ID
const FunctionIDColumnName = "id"

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
	cat.functions, cat.functionMap = readFunctionDefs(cat.dbconn)
}

func readFunctionDefs(db *pgxpool.Pool) ([]*Function, map[string]*Function) {
	log.Debug(sqlFunctions)
	rows, err := db.Query(context.Background(), sqlFunctions)
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
	inDefaults := toArray(inDefaultsTA)
	outNames := toArray(outNamesTA)
	outTypes := toArray(outTypesTA)
	outJSONTypes := toJSONTypeFromPGArray(outTypes)

	datatypes := make(map[string]string)
	addTypes(datatypes, inNames, inTypes)
	addTypes(datatypes, outNames, outTypes)

	// synthesize a description if none provided
	if description == "" {
		description = fmt.Sprintf("The function %v", id)
	}

	geomCol := geometryColumn(outNames, datatypes)

	return &Function{
		ID:             id,
		Schema:         schema,
		Name:           name,
		Description:    description,
		InNames:        inNames,
		InTypes:        inTypes,
		InDefaults:     inDefaults,
		OutNames:       outNames,
		OutTypes:       outTypes,
		OutJSONTypes:   outJSONTypes,
		Types:          datatypes,
		GeometryColumn: geomCol,
	}
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

func (cat *catalogDB) FunctionFeatures(name string, param QueryParam) ([]string, error) {
	fn, err := cat.FunctionByName(name)
	if err != nil || fn == nil {
		return nil, err
	}
	sqlNamedArgs := inputArgs(fn.InNames, param.Values)
	propCols := removeNames(param.Columns, fn.GeometryColumn, "")
	idColIndex := indexOfName(propCols, FunctionIDColumnName)
	sql, argValues := sqlGeomFunction(fn, sqlNamedArgs, propCols, param)
	log.Debugf("%v -- Args: %v", sql, argValues)
	features, err := readFeaturesWithArgs(cat.dbconn, sql, argValues, idColIndex, propCols)
	return features, err
}

func (cat *catalogDB) FunctionData(name string, param QueryParam) ([]map[string]interface{}, error) {
	fn, err := cat.FunctionByName(name)
	if err != nil || fn == nil {
		return nil, err
	}
	sqlNamedArgs := inputArgs(fn.InNames, param.Values)
	propCols := param.Columns
	sql, argValues := sqlFunction(fn, sqlNamedArgs, propCols, param)
	log.Debugf("%v -- Args: %v", sql, argValues)
	data, err := readDataWithArgs(cat.dbconn, propCols, sql, argValues)
	return data, err
}

// inputArgs extracts function arguments from any provided in the query parameters
// Arg values are stored as strings, and relies on Postgres to convert
// to the actual types required by the function
func inputArgs(params []string, queryArgs map[string]string) map[string]string {
	sqlArgs := make(map[string]string)
	for _, param := range params {
		if val, ok := queryArgs[param]; ok {
			sqlArgs[param] = val
		}
	}
	return sqlArgs
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

func readDataWithArgs(db *pgxpool.Pool, propCols []string, sql string, args []interface{}) ([]map[string]interface{}, error) {
	rows, err := db.Query(context.Background(), sql, args...)
	if err != nil {
		log.Warnf("Error running Data query: %v", err)
		return nil, err
	}
	return scanData(rows, propCols), nil
}

func scanData(rows pgx.Rows, propCols []string) []map[string]interface{} {
	var data []map[string]interface{}
	for rows.Next() {
		obj := scanDataRow(rows, true, propCols)
		//log.Println(feature)
		data = append(data, obj)
	}
	// Check for errors from iterating over rows.
	if err := rows.Err(); err != nil {
		log.Warnf("Error reading Data rows: %v", err)
	}
	rows.Close()
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
