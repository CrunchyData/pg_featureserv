package api

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
*/

import (
	"time"

	"github.com/getkin/kin-openapi/openapi3"
)

type PGType string
type JSONType string

// Constants
const (
	JSONTypeString       JSONType = "string"
	JSONTypeNumber       JSONType = "number"
	JSONTypeBoolean      JSONType = "boolean"
	JSONTypeJSON         JSONType = "json"
	JSONTypeDate         JSONType = "date"
	JSONTypeGeometry     JSONType = "geometry"
	JSONTypeBooleanArray JSONType = "boolean[]"
	JSONTypeStringArray  JSONType = "string[]"
	JSONTypeNumberArray  JSONType = "number[]"
)

// Constants
const (
	PGTypeBool         PGType = "bool"
	PGTypeBoolArray    PGType = "_bool"
	PGTypeInt          PGType = "int"
	PGTypeIntArray     PGType = "_int"
	PGTypeInt4         PGType = "int4"
	PGTypeInt4Array    PGType = "_int4"
	PGTypeInt8         PGType = "int8"
	PGTypeInt8Array    PGType = "_int8"
	PGTypeBigInt       PGType = "bigint"
	PGTypeBigIntArray  PGType = "_bigint"
	PGTypeFloat4       PGType = "float4"
	PGTypeFloat4Array  PGType = "_float4"
	PGTypeFloat8       PGType = "float8"
	PGTypeFloat8Array  PGType = "_float8"
	PGTypeNumeric      PGType = "numeric"
	PGTypeDate         PGType = "date"
	PGTypeJSON         PGType = "json"
	PGTypeGeometry     PGType = "geometry"
	PGTypeText         PGType = "text"
	PGTypeTextArray    PGType = "_text"
	PGTypeTSVECTOR     PGType = "tsvector"
	PGTypeVarChar      PGType = "varchar"
	PGTypeVarCharArray PGType = "_varchar"
)

// returns JSONType matching PGType
func (dbType PGType) ToJSONType() JSONType {
	//fmt.Printf("ToJSONType: %v\n", pgType)
	switch dbType {
	case PGTypeNumeric, PGTypeInt, PGTypeInt4, PGTypeInt8, PGTypeBigInt, PGTypeFloat4, PGTypeFloat8:
		return JSONTypeNumber
	case PGTypeIntArray, PGTypeInt4Array, PGTypeInt8Array, PGTypeBigIntArray, PGTypeFloat4Array, PGTypeFloat8Array:
		return JSONTypeNumberArray
	case PGTypeBool:
		return JSONTypeBoolean
	case PGTypeBoolArray:
		return JSONTypeBooleanArray
	case PGTypeJSON:
		return JSONTypeJSON
	case PGTypeTextArray, PGTypeVarCharArray:
		return JSONTypeStringArray
	case PGTypeDate:
		return JSONTypeDate
		// hack to allow displaying geometry type
	case PGTypeGeometry:
		return JSONTypeGeometry
		// default is string
		// this forces conversion to text in SQL query
	default:
		return JSONTypeString
	}
}

// creates openapi schema type according to PGType
func (dbType PGType) ToOpenApiSchema() *openapi3.Schema {
	//fmt.Printf("ToOpenApiType: %v\n", pgType)
	switch dbType {
	case PGTypeBool:
		return &openapi3.Schema{Type: "boolean"}

	case PGTypeInt, PGTypeInt4, PGTypeInt8, PGTypeBigInt:
		return &openapi3.Schema{Type: "integer"}

	case PGTypeFloat4, PGTypeFloat8:
		return &openapi3.Schema{Type: "number"}

	case PGTypeText, PGTypeVarChar:
		return &openapi3.Schema{Type: "string"}

	case PGTypeDate:
		return &openapi3.Schema{Type: "string"}

	case PGTypeGeometry, PGTypeJSON:
		return &openapi3.Schema{Type: "object"}

	case PGTypeIntArray, PGTypeInt4Array, PGTypeInt8Array, PGTypeBigIntArray, PGTypeFloat4Array, PGTypeFloat8Array, PGTypeTextArray, PGTypeVarCharArray, PGTypeBoolArray:
		var subPropType string

		switch dbType {
		case PGTypeBoolArray:
			subPropType = "boolean"
		case PGTypeIntArray, PGTypeInt4Array, PGTypeInt8Array, PGTypeBigIntArray:
			subPropType = "integer"
		case PGTypeFloat4Array, PGTypeFloat8Array:
			subPropType = "number"
		case PGTypeTextArray, PGTypeVarChar:
			subPropType = "string"
		default:
			subPropType = string(dbType)
		}

		return &openapi3.Schema{
			Type: "array",
			Items: &openapi3.SchemaRef{
				Value: &openapi3.Schema{
					Type: subPropType,
				},
			},
		}

	default:
		return &openapi3.Schema{Type: string(dbType)}
	}
}

// converts json marshalled interface to valid object according to PGType
func (dbType PGType) ParseJSONInterface(val interface{}) (interface{}, error) {
	var convVal interface{}

	switch dbType {
	case PGTypeInt4:
		convVal = int32(val.(float64))

	case PGTypeInt8, PGTypeBigInt:
		convVal = int64(val.(float64))

	case PGTypeFloat4:
		convVal = float32(val.(float64))

	case PGTypeFloat8, PGTypeNumeric:
		convVal = val.(float64)

	case PGTypeText, PGTypeVarChar, PGTypeTSVECTOR:
		convVal = val.(string)

	case PGTypeDate:
		var err error
		convVal, err = time.Parse(time.RFC3339, val.(string))
		if err != nil {
			return -9999, err
		}

	// TODO: find a solution to avoid code duplicates
	case PGTypeBoolArray:
		arrI := val.([]interface{})
		convArr := make([]bool, len(arrI))
		for i, v := range arrI {
			convArr[i] = v.(bool)
		}
		convVal = convArr

	case PGTypeInt4Array:
		arrI := val.([]interface{})
		convArr := make([]int32, len(arrI))
		for i, v := range arrI {
			convArr[i] = int32(v.(float64))
		}
		convVal = convArr

	case PGTypeInt8Array, PGTypeBigIntArray:
		arrI := val.([]interface{})
		convArr := make([]int64, len(arrI))
		for i, v := range arrI {
			convArr[i] = int64(v.(float64))
		}
		convVal = convArr

	case PGTypeFloat4Array:
		arrI := val.([]interface{})
		convArr := make([]float32, len(arrI))
		for i, v := range arrI {
			convArr[i] = float32(v.(float64))
		}
		convVal = convArr

	case PGTypeFloat8Array:
		arrI := val.([]interface{})
		convArr := make([]float64, len(arrI))
		for i, v := range arrI {
			convArr[i] = v.(float64)
		}
		convVal = convArr

	default:
		convVal = val
	}

	return convVal, nil
}

func ToJSONTypeFromPGArray(pgTypes []string) []JSONType {
	jsonTypes := make([]JSONType, len(pgTypes))
	for i, pgType := range pgTypes {
		jsonTypes[i] = PGType(pgType).ToJSONType()
	}
	return jsonTypes
}
