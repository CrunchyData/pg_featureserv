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
)

const (
	TagCollections = "collections"
	TagItems       = "items"
	TagConformance = "conformance"
	TagAPI         = "api"

	TagFunctions = "functions"

	ParamLimit     = "limit"
	ParamOffset    = "offset"
	ParamBbox      = "bbox"
	ParamPrecision = "precision"
	ParamTransform = "transform"

	RelSelf = "self"
	RelAlt  = "alternate"
	RelData = "data"

	GeoJSONFeatureCollection = "FeatureCollection"
)

// RootInfo content at root
type RootInfo struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Links       []*Link `json:"links"`
}

// Link for links
type Link struct {
	Href  string `json:"href"`
	Rel   string `json:"rel"`
	Type  string `json:"type"`
	Title string `json:"title"`
}

// Bbox for extent
type Bbox struct {
	Crs    string    `json:"crs"`
	Extent []float64 `json:"bbox"`
}

// CollectionsInfo for all collections
type CollectionsInfo struct {
	Links       []*Link           `json:"links"`
	Collections []*CollectionInfo `json:"collections"`
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

// FeatureCollection info
type FeatureCollectionRaw struct {
	Type           string             `json:"type"`
	Features       []*json.RawMessage `json:"features"`
	NumberMatched  uint               `json:"numberMatched,omitempty"`
	NumberReturned uint               `json:"numberReturned"`
	TimeStamp      string             `json:"timeStamp,omitempty"`
	Links          []*Link            `json:"links"`
}

// FunctionsInfo for all functions
type FunctionsInfo struct {
	Links     []*Link         `json:"links"`
	Functions []*FunctionInfo `json:"functions"`
}

// FunctionInfo for a collection
type FunctionInfo struct {
	ID          string  `json:"id"`
	Description string  `json:"description,omitempty"`
	Links       []*Link `json:"links"`
	Function    *data.Function
}

type Conformance struct {
	ConformsTo []string `json:"conformsTo"`
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
		ID:          fn.ID,
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
