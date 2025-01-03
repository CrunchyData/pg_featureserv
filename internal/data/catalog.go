package data

import (
	"context"
	"fmt"
	"strings"
)

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

const (
	//errMsgCollectionNotFound = "Collection not found: %v"
	//errMsgFeatureNotFound    = "Feature not found: %v"
	SRID_4326    = 4326
	SRID_UNKNOWN = -1
)

// Catalog tbd
type Catalog interface {
	SetIncludeExclude(includeList []string, excludeList []string)

	Tables() ([]*Table, error)

	// TableByName returns the table with given name.
	// It returns nil if the table does not exist
	TableByName(name string) (*Table, error)

	// TableReload reloads volatile table data
	TableReload(name string)

	// TableFeatures returns an array of the JSON for the features in a table
	// It returns nil if the table does not exist
	TableFeatures(ctx context.Context, name string, param *QueryParam) ([]string, error)

	// TableFeature returns the JSON text for a table feature with given id
	// It returns an empty string if the table or feature does not exist
	TableFeature(ctx context.Context, name string, id string, param *QueryParam) (string, error)

	Functions() ([]*Function, error)

	// FunctionByName returns the function with given name.
	// It returns nil if the function does not exist
	FunctionByName(name string) (*Function, error)

	FunctionFeatures(ctx context.Context, name string, args map[string]string, param *QueryParam) ([]string, error)

	FunctionData(ctx context.Context, name string, args map[string]string, param *QueryParam) ([]map[string]interface{}, error)

	Close()
}

// TransformFunction denotes a geometry function with arguments
type TransformFunction struct {
	Name string
	Arg  []string
}

type Sorting struct {
	Name   string
	IsDesc bool // false = ASC (default), true = DESC
}

type PropertyFilter struct {
	Name  string
	Value string
}

// QueryParam holds the optional parameters for a data query
type QueryParam struct {
	Crs       int
	Limit     int
	Offset    int
	Bbox      *Extent
	BboxCrs   int
	FilterSql string
	Filter    []*PropertyFilter
	// Columns is the list of columns to return
	Columns       []string
	GroupBy       []string
	SortBy        []Sorting
	Precision     int
	TransformFuns []TransformFunction
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
	DbTypes        map[string]string
	JSONTypes      []string
	ColDesc        []string
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
	InDbTypes      []string
	InTypeMap      map[string]string
	InDefaults     []string
	NumNoDefault   int
	OutNames       []string
	OutDbTypes     []string
	OutJSONTypes   []string
	Types          map[string]string
	GeometryColumn string
	IDColumn       string
}

func (fun *Function) IsGeometryFunction() bool {
	for _, typ := range fun.OutDbTypes {
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

// Creates a fully qualified function id.
// adds default postgisftw schema if name arg has no schema
func FunctionQualifiedId(name string) string {
	if strings.Contains(name, ".") {
		return name
	}
	return SchemaPostGISFTW + "." + name
}
