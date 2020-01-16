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
	Layers() ([]*Layer, error)

	// LayerByName returns the layer with given name.
	// It returns nil if the layer does not exist
	LayerByName(name string) (*Layer, error)

	// LayerFeatures returns an array of the JSON for the features in a layer
	// It returns nil if the layer does not exist
	LayerFeatures(name string, param QueryParam) ([]string, error)

	// LayerFeature returns the JSON text for a layer feature with given id
	// It returns an empty string if the layer or feature does not exist
	LayerFeature(name string, id string, param QueryParam) (string, error)

	Functions() ([]*Function, error)

	// FunctionByName returns the function with given name.
	// It returns nil if the function does not exist
	FunctionByName(name string) (*Function, error)

	FunctionFeatures(name string, param QueryParam) ([]string, error)
}

// TransformFunction denotes a geometry function with arguments
type TransformFunction struct {
	Name string
	Arg  []string
}

// QueryParam holds the optional parameters for an items query
type QueryParam struct {
	Limit         int
	Bbox          *Extent
	Precision     int
	TransformFuns []TransformFunction
}

// Layer tbd
type Layer struct {
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

// Extent of a layer
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
	errMsgLayerNotFound   = "Layer not found: %v"
	errMsgFeatureNotFound = "Feature not found: %v"
)

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
