package api

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
	PGTypeBool        PGType = "bool"
	PGTypeBoolArray   PGType = "_bool"
	PGTypeInt         PGType = "int"
	PGTypeIntArray    PGType = "_int"
	PGTypeInt4        PGType = "int4"
	PGTypeInt4Array   PGType = "_int4"
	PGTypeBigInt      PGType = "bigint"
	PGTypeBigIntArray PGType = "_bigint"
	PGTypeFloat4      PGType = "float4"
	PGTypeFloat4Array PGType = "_float4"
	PGTypeFloat8      PGType = "float8"
	PGTypeFloat8Array PGType = "_float8"
	PGTypeNumeric     PGType = "numeric"
	PGTypeDate        PGType = "date"
	PGTypeJSON        PGType = "json"
	PGTypeGeometry    PGType = "geometry"
	PGTypeText        PGType = "text"
	PGTypeTextArray   PGType = "_text"
	PGTypeTSVECTOR    PGType = "tsvector"
)

func (dbType PGType) ToJSONType() JSONType {
	//fmt.Printf("ToJSONType: %v\n", pgType)
	switch dbType {
	case PGTypeNumeric, PGTypeInt, PGTypeInt4, PGTypeBigInt, PGTypeFloat4, PGTypeFloat8:
		return JSONTypeNumber
	case PGTypeIntArray, PGTypeInt4Array, PGTypeBigIntArray, PGTypeFloat4Array, PGTypeFloat8Array:
		return JSONTypeNumberArray
	case PGTypeBool:
		return JSONTypeBoolean
	case PGTypeBoolArray:
		return JSONTypeBooleanArray
	case PGTypeJSON:
		return JSONTypeJSON
	case PGTypeTextArray:
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

func (dbType PGType) ToOpenApiType() string {
	//fmt.Printf("ToOpenApiType: %v\n", pgType)
	switch dbType {
	case PGTypeBool:
		return "boolean"
	case PGTypeInt:
		return "integer"
	case PGTypeInt4:
		return "integer"
	case PGTypeBigInt:
		return "int64"
	case PGTypeText:
		return "string"
	default:
		return ""
	}
}

func ToJSONTypeFromPGArray(pgTypes []string) []JSONType {
	jsonTypes := make([]JSONType, len(pgTypes))
	for i, pgType := range pgTypes {
		jsonTypes[i] = PGType(pgType).ToJSONType()
	}
	return jsonTypes
}
