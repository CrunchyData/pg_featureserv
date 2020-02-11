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

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/CrunchyData/pg_featureserv/config"
	"github.com/CrunchyData/pg_featureserv/data"
	"github.com/getkin/kin-openapi/openapi3"
)

const (
	RootPageName   = "index"
	TagCollections = "collections"
	TagItems       = "items"
	TagConformance = "conformance"
	TagAPI         = "api"

	TagFunctions = "functions"

	ParamLimit      = "limit"
	ParamOffset     = "offset"
	ParamBbox       = "bbox"
	ParamOrderBy    = "orderby"
	ParamPrecision  = "precision"
	ParamProperties = "properties"
	ParamTransform  = "transform"

	RelSelf      = "self"
	RelAlt       = "alternate"
	RelData      = "data"
	RelFunctions = "functions"
	RelItems     = "items"

	GeoJSONFeatureCollection = "FeatureCollection"
)

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
	Type:     "object",
	Required: []string{"href"},
	Properties: map[string]*openapi3.SchemaRef{
		"href":     {Value: &openapi3.Schema{Type: "string"}},
		"rel":      {Value: &openapi3.Schema{Type: "string"}},
		"type":     {Value: &openapi3.Schema{Type: "string"}},
		"hreflang": {Value: &openapi3.Schema{Type: "string"}},
		"title":    {Value: &openapi3.Schema{Type: "string"}},
	},
}

// Bbox for extent
type Bbox struct {
	Crs    string    `json:"crs"`
	Extent []float64 `json:"bbox"`
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

// CollectionInfo for a collection
type CollectionInfo struct {
	Name        string   `json:"id"`
	Title       string   `json:"title,omitempty"`
	Description string   `json:"description,omitempty"`
	Extent      *Bbox    `json:"extent,omitempty"`
	Crs         []string `json:"crs,omitempty"`
	Links       []*Link  `json:"links"`
}

var CollectionInfoSchema openapi3.Schema = openapi3.Schema{
	Type:     "object",
	Required: []string{"name", "links"},
	Properties: map[string]*openapi3.SchemaRef{
		"id":          {Value: &openapi3.Schema{Type: "string"}},
		"title":       {Value: &openapi3.Schema{Type: "string"}},
		"description": {Value: &openapi3.Schema{Type: "string"}},
		"links": {Value: &openapi3.Schema{
			Type:  "array",
			Items: &openapi3.SchemaRef{Value: &LinkSchema},
		},
		},
		"extent": {Value: &BboxSchema},
		"crs": {Value: &openapi3.Schema{
			Type: "array",
			Items: &openapi3.SchemaRef{
				Value: &openapi3.Schema{Type: "string"},
			},
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
	Links     []*Link         `json:"links"`
	Functions []*FunctionInfo `json:"functions"`
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
				Items: &openapi3.SchemaRef{Value: &FunctionInfoSchema},
			},
		},
	},
}

// FunctionInfo is the API metadata for a function
type FunctionInfo struct {
	Name        string  `json:"id"`
	Description string  `json:"description,omitempty"`
	Links       []*Link `json:"links"`
	Function    *data.Function
}

var FunctionInfoSchema openapi3.Schema = openapi3.Schema{
	Type:     "object",
	Required: []string{"name", "links"},
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
	ErrMsgDataRead              = "Unable to read data from: %v"
)

const (
	ErrCodeCollectionNotFound = "CollectionNotFound"
	ErrCodeFeatureNotFound    = "FeatureNotFound"
)

var conformance = Conformance{
	ConformsTo: []string{
		"http://www.opengis.net/spec/wfs-1/3.0/req/core",
		"http://www.opengis.net/spec/wfs-1/3.0/req/geojson",
		"http://www.opengis.net/spec/wfs-1/3.0/req/html",
	},
}

func toBbox(cc *data.Table) *Bbox {
	return &Bbox{
		Crs:    fmt.Sprintf("EPSG:%v", cc.Srid),
		Extent: []float64{cc.Extent.Minx, cc.Extent.Miny, cc.Extent.Maxx, cc.Extent.Maxy},
	}
}

func NewRootInfo(conf *config.Config) *RootInfo {
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
		Extent:      toBbox(tbl),
	}
	return &doc
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
	fnsDoc := FunctionsInfo{Links: []*Link{}, Functions: []*FunctionInfo{}}
	for _, fn := range fns {
		fnDoc := NewFunctionInfo(fn)
		fnsDoc.Functions = append(fnsDoc.Functions, fnDoc)
	}
	return &fnsDoc
}

func NewFunctionInfo(fn *data.Function) *FunctionInfo {
	info := FunctionInfo{
		Name:        fn.Name,
		Description: fn.Description,
		Function:    fn,
	}
	return &info
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

func PathFunction(name string) string {
	return fmt.Sprintf("%v/%v", TagFunctions, name)
}

func PathItems(tagType string, name string) string {
	return fmt.Sprintf("%v/%v/%v", tagType, name, TagItems)
}

func PathItem(name string, fid string) string {
	return fmt.Sprintf("%v/%v/%v/%v", TagCollections, name, TagItems, fid)
}
