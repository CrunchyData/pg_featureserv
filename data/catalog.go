package data

import (
	"fmt"
	"strings"
)

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

// Catalog tbd
type Catalog interface {
	Tables() ([]*Table, error)

	// TableByName returns the table with given name.
	// It returns nil if the table does not exist
	TableByName(name string) (*Table, error)

	// TableFeatures returns an array of the JSON for the features in a table
	// It returns nil if the table does not exist
	TableFeatures(name string, param QueryParam) ([]string, error)

	// TableFeature returns the JSON text for a table feature with given id
	// It returns an empty string if the table or feature does not exist
	TableFeature(name string, id string, param QueryParam) (string, error)

	Functions() ([]*Function, error)

	// FunctionByName returns the function with given name.
	// It returns nil if the function does not exist
	FunctionByName(name string) (*Function, error)

	FunctionFeatures(name string, param QueryParam) ([]string, error)

	FunctionData(name string, param QueryParam) ([]map[string]interface{}, error)
}

// TransformFunction denotes a geometry function with arguments
type TransformFunction struct {
	Name string
	Arg  []string
}

type ParamNameVal map[string]string

// QueryParam holds the optional parameters for an items query
type QueryParam struct {
	Limit      int
	Offset     int
	Bbox       *Extent
	Properties []string
	// Columns is the clean list of columns to query
	Columns       []string
	Precision     int
	TransformFuns []TransformFunction
	Values        ParamNameVal
}

// Table holds metadata for table/view objects
type Table struct {
	ID             string
	Schema         string
	Table          string
	Title          string
	Description    string
	GeometryType   string
	GeometryColumn string
	IDColumn       string
	Srid           int
	Extent         Extent
	Columns        []string
	Types          map[string]string
	JSONTypes      []string
}

// Extent of a table
type Extent struct {
	Minx, Miny, Maxx, Maxy float64
}

// Function tbd
type Function struct {
	ID             string
	Schema         string
	Name           string
	Description    string
	InNames        []string
	InTypes        []string
	InDefaults     []string
	OutNames       []string
	OutTypes       []string
	Types          map[string]string
	JSONTypes      []string
	GeometryColumn string
	IDColumn       string
}

const (
	errMsgCollectionNotFound = "Collection not found: %v"
	errMsgFeatureNotFound    = "Feature not found: %v"
)

func (fun *Function) IsGeometryFunction() bool {
	for _, typ := range fun.OutTypes {
		if typ == "geometry" {
			return true
		}
	}
	return false
}

func (fun *TransformFunction) apply(expr string) string {
	if fun.Name == "" {
		return expr
	}
	if len(fun.Arg) == 0 {
		return fmt.Sprintf("%v( %v )", fun.Name, expr)
	}
	args := strings.Join(fun.Arg, ",")
	return fmt.Sprintf("%v( %v, %v )", fun.Name, expr, args)
}
