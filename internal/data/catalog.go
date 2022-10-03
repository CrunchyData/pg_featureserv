package data

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

 Date     : October 2022
 Authors  : Benoit De Mezzo (benoit dot de dot mezzo at oslandia dot com)
			Nicolas Revelant (nicolas dot revelant at ign dot fr)
*/

import (
	"context"

	"github.com/CrunchyData/pg_featureserv/internal/api"
)

const (
	//errMsgCollectionNotFound = "Collection not found: %v"
	//errMsgFeatureNotFound    = "Feature not found: %v"
	SRID_4326    = 4326
	SRID_UNKNOWN = -1
)

// Catalog tbd
type Catalog interface {
	Initialize(includeList []string, excludeList []string)

	Tables() ([]*api.Table, error)

	// TableByName returns the table with given name.
	// It returns nil if the table does not exist
	TableByName(name string) (*api.Table, error)

	// TableReload reloads volatile table data
	TableReload(name string)

	// TableFeatures returns an array of the JSON for the features in a table
	// It returns nil if the table does not exist
	TableFeatures(ctx context.Context, name string, param *QueryParam) ([]*api.GeojsonFeatureData, error)

	// TableFeature returns the JSON text for a table feature with given id, along with its weak etag value
	// It returns an empty string if the table or feature does not exist
	TableFeature(ctx context.Context, name string, id string, param *QueryParam) (*api.GeojsonFeatureData, error)

	// AddTableFeature returns the id of the new feature created in the table tableName
	// using the JSON data to create the feature
	AddTableFeature(ctx context.Context, tableName string, jsonData []byte) (int64, error)

	// PartialUpdateTableFeature updates a table feature with given id with the JSON data
	PartialUpdateTableFeature(ctx context.Context, tableName string, id string, jsonData []byte) error

	// ReplaceTableFeature replaces a table feature with given id with the new jsonData
	ReplaceTableFeature(ctx context.Context, tableName string, id string, jsonData []byte) error

	// DeleteTableFeature returns the status code from the delete operation on the feature which ID is provided
	DeleteTableFeature(ctx context.Context, tableName string, id string) error

	Functions() ([]*api.Function, error)

	// FunctionByName returns the function with given name.
	// It returns nil if the function does not exist
	FunctionByName(name string) (*api.Function, error)

	FunctionFeatures(ctx context.Context, name string, args map[string]string, param *QueryParam) ([]*api.GeojsonFeatureData, error)

	FunctionData(ctx context.Context, name string, args map[string]string, param *QueryParam) ([]map[string]interface{}, error)

	// CheckStrongEtags checks if at least one of the etags provided is present into the cache
	// Returns true at the first etag detected as present, false otherwise
	// -> error != nil if a malformed etag is detected (wrong encoding, bad format.)
	// -> The provided etags have to be in their strong form and Base64 encoded
	CheckStrongEtags(etagsList []string) (bool, error)
	// GetCache returns a copy of the cache
	// GetCache() map[string]interface{}

	Close()
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
	Bbox      *api.Extent
	BboxCrs   int
	FilterSql string
	Filter    []*PropertyFilter
	// Columns is the list of columns to return
	Columns       []string
	GroupBy       []string
	SortBy        []api.Sorting
	Precision     int
	TransformFuns []api.TransformFunction
}
