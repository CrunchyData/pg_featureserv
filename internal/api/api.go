package api

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
	"encoding/json"
	"fmt"
	"time"

	"github.com/CrunchyData/pg_featureserv/internal/conf"
	"github.com/CrunchyData/pg_featureserv/internal/data"
	"github.com/getkin/kin-openapi/openapi3"
)

const (
	RootPageName   = "index"
	TagCollections = "collections"
	TagItems       = "items"
	TagConformance = "conformance"
	TagAPI         = "api"

	TagFunctions = "functions"

	ParamCrs        = "crs"
	ParamLimit      = "limit"
	ParamOffset     = "offset"
	ParamBbox       = "bbox"
	ParamBboxCrs    = "bbox-crs"
	ParamFilter     = "filter"
	ParamFilterCrs  = "filter-crs"
	ParamGroupBy    = "groupby"
	ParamOrderBy    = "orderby"
	ParamPrecision  = "precision"
	ParamProperties = "properties"
	ParamSortBy     = "sortby"
	ParamTransform  = "transform"

	OrderByDirSep = ":"
	OrderByDirD   = "d"
	OrderByDirA   = "a"

	RelSelf        = "self"
	RelAlt         = "alternate"
	RelServiceDesc = "service-desc"
	RelServiceDoc  = "service-doc"
	RelConformance = "conformance"
	RelData        = "data"
	RelFunctions   = "functions"
	RelItems       = "items"

	TitleFeatuuresGeoJSON = "Features as GeoJSON"
	TitleDataJSON         = "Data as JSON"
	TitleMetadata         = "Metadata"
	TitleDocument         = "This document"
	TitleAsJSON           = " as JSON"
	TitleAsHTML           = " as HTML"

	GeoJSONFeatureCollection = "FeatureCollection"
)

const (
	ErrMsgEncoding              = "Error encoding response"
	ErrMsgLoadCollections       = "Unable to access Collections"
	ErrMsgCollectionNotFound    = "Collection not found: %v"
	ErrMsgCollectionAccess      = "Unable to access Collection: %v"
	ErrMsgFeatureNotFound       = "Feature not found: %v"
	ErrMsgLoadFunctions         = "Unable to access Functions"
	ErrMsgFunctionNotFound      = "Function not found: %v"
	ErrMsgFunctionAccess        = "Unable to access Function: %v"
	ErrMsgInvalidParameterValue = "Invalid value for parameter %v: %v"
	ErrMsgInvalidQuery          = "Invalid query parameters"
	ErrMsgDataReadError         = "Unable to read data from: %v"
	ErrMsgDataWriteError        = "Unable to write data to: %v"
	ErrMsgNoDataRead            = "No data read from: %v"
	ErrMsgRequestTimeout        = "Maximum time exceeded.  Request cancelled."
)

const (
	ErrCodeCollectionNotFound = "CollectionNotFound"
	ErrCodeFeatureNotFound    = "FeatureNotFound"
)

var ParamReservedNames = []string{
	ParamCrs,
	ParamLimit,
	ParamOffset,
	ParamBbox,
	ParamBboxCrs,
	ParamFilter,
	ParamGroupBy,
	ParamOrderBy,
	ParamPrecision,
	ParamProperties,
	ParamSortBy,
	ParamTransform,
}

var ParamReservedNamesMap = makeSet(ParamReservedNames)

func makeSet(names []string) map[string]string {
	outMap := make(map[string]string)
	for _, s := range names {
		outMap[s] = s
	}
	return outMap
}

func IsParameterReservedName(name string) bool {
	_, ok := ParamReservedNamesMap[name]
	return ok
}

// RootInfo content at root
type RootInfo struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Links       []*Link `json:"links"`
}

var RootInfoSchema openapi3.Schema = openapi3.Schema{
	Type:     "object",
	Required: []string{"links"},
	Properties: map[string]*openapi3.SchemaRef{
		"title": {Value: &openapi3.Schema{
			Type:        "string",
			Description: "Title of this feature service",
		}},
		"description": {Value: &openapi3.Schema{
			Type:        "string",
			Description: "Description of this feature service",
		}},
		"links": {
			Value: &openapi3.Schema{
				Type:  "array",
				Items: &openapi3.SchemaRef{Value: &LinkSchema},
			},
		},
	},
}

// Link for links
type Link struct {
	Href  string `json:"href"`
	Rel   string `json:"rel"`
	Type  string `json:"type"`
	Title string `json:"title"`
}

var LinkSchema openapi3.Schema = openapi3.Schema{
	Description: "Describes links to other resources",
	Type:        "object",
	Required:    []string{"href"},
	Properties: map[string]*openapi3.SchemaRef{
		"href":     {Value: &openapi3.Schema{Type: "string", Description: "URL for the link"}},
		"rel":      {Value: &openapi3.Schema{Type: "string"}},
		"type":     {Value: &openapi3.Schema{Type: "string"}},
		"hreflang": {Value: &openapi3.Schema{Type: "string"}},
		"title":    {Value: &openapi3.Schema{Type: "string"}},
	},
}

var PropertySchema openapi3.Schema = openapi3.Schema{
	Description: "A data property of a collection or function result",
	Type:        "object",
	Required:    []string{"name", "type"},
	Properties: map[string]*openapi3.SchemaRef{
		"name":        {Value: &openapi3.Schema{Type: "string"}},
		"type":        {Value: &openapi3.Schema{Type: "string"}},
		"description": {Value: &openapi3.Schema{Type: "string"}},
	},
}

var ParameterSchema openapi3.Schema = openapi3.Schema{
	Description: "A parameter of a function",
	Type:        "object",
	Required:    []string{"name", "type"},
	Properties: map[string]*openapi3.SchemaRef{
		"name":    {Value: &openapi3.Schema{Type: "string"}},
		"type":    {Value: &openapi3.Schema{Type: "string"}},
		"default": {Value: &openapi3.Schema{Type: "string"}},
	},
}

// Bbox for extent
type Bbox struct {
	Crs    string    `json:"crs"`
	Extent []float64 `json:"bbox"`
}

// Extent OAPIF Extent structure (partial)
type Extent struct {
	Spatial *Bbox `json:"spatial"`
}

// --- @See https://raw.githubusercontent.com/opengeospatial/WFS_FES/master/core/openapi/schemas/bbox.yaml
//	for bbox schema

var BboxSchema openapi3.Schema = openapi3.Schema{
	Type:     "object",
	Required: []string{"bbox"},
	Properties: map[string]*openapi3.SchemaRef{
		"crs": {
			// TODO: This is supposed to have an enum & default based on: http://www.opengis.net/def/crs/OGC/1.3/CRS84
			Value: openapi3.NewStringSchema(),
		},
		"bbox": {
			Value: &openapi3.Schema{
				Type:     "array",
				MinItems: 4,
				MaxItems: openapi3.Uint64Ptr(4),
				Items:    openapi3.NewSchemaRef("", openapi3.NewFloat64Schema().WithMin(-180).WithMax(180)),
			},
		},
	},
}

var ExtentSchema openapi3.Schema = openapi3.Schema{
	Type:     "object",
	Required: []string{"extent"},
	Properties: map[string]*openapi3.SchemaRef{
		"spatial": {Value: &BboxSchema},
	},
}

type NameValMap map[string]string

// RequestParam holds the parameters for a request
type RequestParam struct {
	Crs           int
	Limit         int
	Offset        int
	Bbox          *data.Extent
	BboxCrs       int
	Properties    []string
	Filter        string
	FilterCrs     int
	GroupBy       []string
	SortBy        []data.Sorting
	Precision     int
	TransformFuns []data.TransformFunction
	Values        NameValMap
}

// CollectionsInfo for all collections
type CollectionsInfo struct {
	Links       []*Link           `json:"links"`
	Collections []*CollectionInfo `json:"collections"`
}

var CollectionsInfoSchema openapi3.Schema = openapi3.Schema{
	Type:     "object",
	Required: []string{"links", "collections"},
	Properties: map[string]*openapi3.SchemaRef{
		"links": {
			Value: &openapi3.Schema{
				Type: "array",
				Items: &openapi3.SchemaRef{
					Value: &LinkSchema,
				},
			},
		},
		"collections": {
			Value: &openapi3.Schema{
				Type: "array",
				Items: &openapi3.SchemaRef{
					Value: &CollectionInfoSchema,
				},
			},
		},
	},
}

type Parameter struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Default     string `json:"default,omitempty"`
}

type Property struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

// CollectionInfo for a collection
type CollectionInfo struct {
	Name         string   `json:"id"`
	Title        string   `json:"title,omitempty"`
	Description  string   `json:"description,omitempty"`
	Extent       *Extent  `json:"extent,omitempty"`
	Crs          []string `json:"crs,omitempty"`
	GeometryType *string  `json:"geometrytype,omitempty"`

	// these are omitempty so they don't show in summary metadata
	Properties []*Property `json:"properties,omitempty"`

	Links []*Link `json:"links"`
	// used for HTML response only
	URLMetadataHTML string `json:"-"`
	URLMetadataJSON string `json:"-"`
	URLItemsHTML    string `json:"-"`
	URLItemsJSON    string `json:"-"`
}

var CollectionInfoSchema openapi3.Schema = openapi3.Schema{
	Type:     "object",
	Required: []string{"id", "links"},
	Properties: map[string]*openapi3.SchemaRef{
		"id":          {Value: &openapi3.Schema{Type: "string"}},
		"title":       {Value: &openapi3.Schema{Type: "string"}},
		"description": {Value: &openapi3.Schema{Type: "string"}},
		"extent":      {Value: &ExtentSchema},
		"crs": {Value: &openapi3.Schema{
			Type: "array",
			Items: &openapi3.SchemaRef{
				Value: &openapi3.Schema{Type: "string"},
			},
		},
		},
		"geometrytype": {Value: &openapi3.Schema{Type: "string"}},
		"properties": {Value: &openapi3.Schema{
			Type:  "array",
			Items: &openapi3.SchemaRef{Value: &PropertySchema},
		},
		},
		"links": {Value: &openapi3.Schema{
			Type:  "array",
			Items: &openapi3.SchemaRef{Value: &LinkSchema},
		},
		},
	},
}

// FeatureCollection info
type FeatureCollectionRaw struct {
	Type           string             `json:"type"`
	Features       []*json.RawMessage `json:"features"`
	NumberMatched  uint               `json:"numberMatched,omitempty"`
	NumberReturned uint               `json:"numberReturned"`
	TimeStamp      string             `json:"timeStamp,omitempty"`
	Links          []*Link            `json:"links"`
}

// FunctionsInfo is the API metadata for all functions
type FunctionsInfo struct {
	Links     []*Link            `json:"links"`
	Functions []*FunctionSummary `json:"functions"`
}

var FunctionsInfoSchema openapi3.Schema = openapi3.Schema{
	Type:     "object",
	Required: []string{"links", "functions"},
	Properties: map[string]*openapi3.SchemaRef{
		"links": {
			Value: &openapi3.Schema{
				Type:  "array",
				Items: &openapi3.SchemaRef{Value: &LinkSchema},
			},
		},
		"functions": {
			Value: &openapi3.Schema{
				Type:  "array",
				Items: &openapi3.SchemaRef{Value: &FunctionSummarySchema},
			},
		},
	},
}

// FunctionSummary contains a restricted set of function metadata for use in list display and JSON
// This allows not including parameters and properties in list metadata,
// but ensuring those keys are always present in full metadata JSON.
// Note: Collections do not follow same pattern because their list JSON metadata
// is supposed to contain all properties, and they are always expected to have attribute properties
type FunctionSummary struct {
	Name        string  `json:"id"`
	Description string  `json:"description,omitempty"`
	Links       []*Link `json:"links"`

	//--- additional data used during processing
	Function *data.Function `json:"-"`
	// used for HTML response only
	URLMetadataHTML string `json:"-"`
	URLMetadataJSON string `json:"-"`
	URLItemsHTML    string `json:"-"`
	URLItemsJSON    string `json:"-"`
}

var FunctionSummarySchema openapi3.Schema = openapi3.Schema{
	Type:     "object",
	Required: []string{"id", "links"},
	Properties: map[string]*openapi3.SchemaRef{
		"id":          {Value: &openapi3.Schema{Type: "string"}},
		"description": {Value: &openapi3.Schema{Type: "string"}},
		"links": {Value: &openapi3.Schema{
			Type:  "array",
			Items: &openapi3.SchemaRef{Value: &LinkSchema},
		},
		},
	},
}

// FunctionInfo is the API metadata for a function
type FunctionInfo struct {
	Name        string `json:"id"`
	Description string `json:"description,omitempty"`

	// these properties are always present but may be empty arrays
	Parameters []*Parameter `json:"parameters"`
	Properties []*Property  `json:"properties"`

	Links []*Link `json:"links"`

	//--- additional data used during processing
	Function *data.Function `json:"-"`
}

var FunctionInfoSchema openapi3.Schema = openapi3.Schema{
	Type:     "object",
	Required: []string{"id", "links"},
	Properties: map[string]*openapi3.SchemaRef{
		"id":          {Value: &openapi3.Schema{Type: "string"}},
		"description": {Value: &openapi3.Schema{Type: "string"}},
		"parameters": {Value: &openapi3.Schema{
			Type:  "array",
			Items: &openapi3.SchemaRef{Value: &ParameterSchema},
		},
		},
		"properties": {Value: &openapi3.Schema{
			Type:  "array",
			Items: &openapi3.SchemaRef{Value: &PropertySchema},
		},
		},
		"links": {Value: &openapi3.Schema{
			Type:  "array",
			Items: &openapi3.SchemaRef{Value: &LinkSchema},
		},
		},
	},
}

type Conformance struct {
	ConformsTo []string `json:"conformsTo"`
}

var ConformanceSchema openapi3.Schema = openapi3.Schema{
	Type:     "object",
	Required: []string{"conformsTo"},
	Properties: map[string]*openapi3.SchemaRef{
		"conformsTo": {
			Value: &openapi3.Schema{
				Type: "array",
				Items: &openapi3.SchemaRef{
					Value: &openapi3.Schema{Type: "string"},
				},
			},
		},
	},
}

var conformance = Conformance{
	ConformsTo: []string{
		"http://www.opengis.net/spec/ogcapi-features-1/1.0/conf/core",
		"http://www.opengis.net/spec/ogcapi-features-1/1.0/conf/oas3",
		"http://www.opengis.net/spec/ogcapi-features-1/1.0/conf/geojson",
		"http://www.opengis.net/spec/ogcapi-features-1/1.0/conf/html",
		"http://www.opengis.net/spec/ogcapi-common-1/1.0/conf/core",
		"http://www.opengis.net/spec/ogcapi-common-1/1.0/conf/landing-page",
		"http://www.opengis.net/spec/ogcapi-common-1/1.0/conf/json",
		"http://www.opengis.net/spec/ogcapi-common-1/1.0/conf/html",
		"http://www.opengis.net/spec/ogcapi-common-1/1.0/conf/oas30",
		"http://www.opengis.net/spec/ogcapi-common-2/1.0/conf/collections",
		"http://www.opengis.net/spec/ogcapi-common-2/1.0/conf/simple-query",
		"http://www.opengis.net/spec/ogcapi-features-4/1.0/conf/create-replace-delete",
		"http://www.opengis.net/spec/ogcapi-features-4/1.0/conf/features",
	},
}

func toBbox(cc *data.Table) *Bbox {
	// extent bbox is always in 4326 for now
	crs := "http://www.opengis.net/def/crs/EPSG/0/4326"
	return &Bbox{
		Crs:    crs,
		Extent: []float64{cc.Extent.Minx, cc.Extent.Miny, cc.Extent.Maxx, cc.Extent.Maxy},
	}
}

func NewLink(href string, rel string, conType string, title string) *Link {
	return &Link{
		Href:  href,
		Rel:   rel,
		Type:  conType,
		Title: title,
	}
}

func NewRootInfo(conf *conf.Config) *RootInfo {
	root := &RootInfo{}
	root.Title = conf.Metadata.Title
	root.Description = conf.Metadata.Description
	return root
}

func NewCollectionsInfo(tables []*data.Table) *CollectionsInfo {
	csDoc := CollectionsInfo{Links: []*Link{}, Collections: []*CollectionInfo{}}
	for _, lyr := range tables {
		collDoc := NewCollectionInfo(lyr)
		csDoc.Collections = append(csDoc.Collections, collDoc)
	}
	return &csDoc
}

func NewCollectionInfo(tbl *data.Table) *CollectionInfo {
	doc := CollectionInfo{
		Name:        tbl.ID,
		Title:       tbl.Title,
		Description: tbl.Description,
		Extent: &Extent{
			Spatial: toBbox(tbl),
		},
	}
	return &doc
}

func TableProperties(tbl *data.Table) []*Property {
	props := make([]*Property, len(tbl.Columns))
	for i, name := range tbl.Columns {
		props[i] = &Property{
			Name:        name,
			Type:        tbl.JSONTypes[i],
			Description: tbl.ColDesc[i],
		}
	}
	return props
}

func NewFeatureCollectionInfo(featureJSON []string) *FeatureCollectionRaw {
	ts := time.Now().Format(time.RFC3339)
	doc := FeatureCollectionRaw{
		Type:           GeoJSONFeatureCollection,
		Features:       toRaw(featureJSON),
		NumberMatched:  0,
		NumberReturned: uint(len(featureJSON)),
		TimeStamp:      ts,
	}
	return &doc
}

func NewFunctionsInfo(fns []*data.Function) *FunctionsInfo {
	fnsDoc := FunctionsInfo{Links: []*Link{}, Functions: []*FunctionSummary{}}
	for _, fn := range fns {
		fnDoc := NewFunctionSummary(fn)
		fnsDoc.Functions = append(fnsDoc.Functions, fnDoc)
	}
	return &fnsDoc
}

func NewFunctionSummary(fn *data.Function) *FunctionSummary {
	info := FunctionSummary{
		Name:        fn.ID,
		Description: fn.Description,
		Function:    fn,
	}
	return &info
}

func NewFunctionInfo(fn *data.Function) *FunctionInfo {
	info := FunctionInfo{
		Name:        fn.ID,
		Description: fn.Description,
		Function:    fn,
	}
	return &info
}

func FunctionParameters(fn *data.Function) []*Parameter {
	params := make([]*Parameter, len(fn.InNames))
	for i, name := range fn.InNames {
		params[i] = &Parameter{
			Name: name,
			Type: fn.InDbTypes[i],
			// no description available from db catalog
			Default: fn.InDefaults[i],
		}
	}
	return params
}

func FunctionProperties(fn *data.Function) []*Property {
	props := make([]*Property, len(fn.OutNames))
	for i, name := range fn.OutNames {
		props[i] = &Property{
			Name: name,
			Type: fn.OutJSONTypes[i],
			// no description available from db catalog
		}
	}
	return props
}

func GetConformance() *Conformance {
	return &conformance
}

func toRaw(jsonStr []string) []*json.RawMessage {
	raw := make([]*json.RawMessage, len(jsonStr))
	for i, f := range jsonStr {
		fRaw := json.RawMessage(f)
		raw[i] = &fRaw
	}
	return raw
}

func PathCollection(name string) string {
	return fmt.Sprintf("%v/%v", TagCollections, name)
}

func PathCollectionItems(name string) string {
	return fmt.Sprintf("%v/%v/%v", TagCollections, name, TagItems)
}

func PathFunction(name string) string {
	return fmt.Sprintf("%v/%v", TagFunctions, name)
}

func PathFunctionItems(name string) string {
	return fmt.Sprintf("%v/%v/%v", TagFunctions, name, TagItems)
}

func PathItem(name string, fid string) string {
	return fmt.Sprintf("%v/%v/%v/%v", TagCollections, name, TagItems, fid)
}
