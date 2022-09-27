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
	"strings"
	"time"

	"github.com/CrunchyData/pg_featureserv/internal/conf"
	"github.com/paulmach/orb/geojson"
	log "github.com/sirupsen/logrus"
)

const (
	RootPageName   = "index"
	TagCollections = "collections"
	TagItems       = "items"
	TagConformance = "conformance"
	TagAPI         = "api"
	TagFunctions   = "functions"

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

	TitleFeaturesGeoJSON = "Features as GeoJSON"
	TitleDataJSON        = "Data as JSON"
	TitleMetadata        = "Metadata"
	TitleDocument        = "This document"
	TitleAsJSON          = " as JSON"
	TitleAsHTML          = " as HTML"

	GeoJSONFeatureCollection = "FeatureCollection"
)

const (
	ErrMsgEncoding                       = "Error encoding response"
	ErrMsgLoadCollections                = "Unable to access Collections"
	ErrMsgCollectionNotFound             = "Collection not found: %v"
	ErrMsgCollectionAccess               = "Unable to access Collection: %v"
	ErrMsgCollectionRequestBodyRead      = "Unable to read request body for Collection: %v"
	ErrMsgFeatureNotFound                = "Feature not found: %v"
	ErrMsgCreateFeatureNotConform        = "Unable to create new feature in Collection - data does not respect schema: %v"
	ErrMsgCreateFeatureInCatalog         = "Unable to create new feature in Collection - catalog error: %v"
	ErrMsgPartialUpdateFeatureNotConform = "Unable to patch feature in Collection - data does not respect schema: %v"
	ErrMsgLoadFunctions                  = "Unable to access Functions"
	ErrMsgPartialUpdateFeature           = "Unable to update feature in Collection: %v"
	ErrMsgFunctionNotFound               = "Function not found: %v"
	ErrMsgFunctionAccess                 = "Unable to access Function: %v"
	ErrMsgInvalidParameterValue          = "Invalid value for parameter %v: %v"
	ErrMsgInvalidQuery                   = "Invalid query parameters"
	ErrMsgDataReadError                  = "Unable to read data from: %v"
	ErrMsgDataWriteError                 = "Unable to write data to: %v"
	ErrMsgNoDataRead                     = "No data read from: %v"
	ErrMsgRequestTimeout                 = "Maximum time exceeded.  Request cancelled."
	ErrMsgReplaceFeature                 = "Unable to replace feature in Collection: %v"
	ErrMsgReplaceFeatureNotConform       = "Unable to replace feature in Collection - data does not respect schema"
	ErrMsgMarshallingJSON                = "Error marshalling into JSON (table: %v, id: %v)"
	ErrMsgMarshallingJSONEtag            = "Error marshalling into JSON: %v"
	ErrMsgNoParameters                   = "No parameter allowed"
)

// ==================================================
// ================== ParamReserved ==================

const (
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
	ParamType       = "type"
)

// known query parameter name
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

// =======================================================
// =================== RootInfo ==========================

// RootInfo content at root
type RootInfo struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Links       []*Link `json:"links"`
}

func NewRootInfo(conf *conf.Config) *RootInfo {
	root := &RootInfo{}
	root.Title = conf.Metadata.Title
	root.Description = conf.Metadata.Description
	return root
}

// =======================================================
// =================== Link ==============================

// Link for links
type Link struct {
	Href  string `json:"href"`
	Rel   string `json:"rel"`
	Type  string `json:"type"`
	Title string `json:"title"`
}

func NewLink(href string, rel string, conType string, title string) *Link {
	return &Link{
		Href:  href,
		Rel:   rel,
		Type:  conType,
		Title: title,
	}
}

// =======================================================
// =======================================================

// Extent of a table
type Extent struct {
	Minx, Miny, Maxx, Maxy float64
}

type Sorting struct {
	Name   string
	IsDesc bool // false = ASC (default), true = DESC
}

// =======================================================
// ================== TransformFunction ==================

// TransformFunction denotes a geometry function with arguments
type TransformFunction struct {
	Name string
	Arg  []string
}

func (fun *TransformFunction) Apply(expr string) string {
	if fun.Name == "" {
		return expr
	}
	if len(fun.Arg) == 0 {
		return fmt.Sprintf("%v( %v )", fun.Name, expr)
	}
	args := strings.Join(fun.Arg, ",")
	return fmt.Sprintf("%v( %v, %v )", fun.Name, expr, args)
}

// ===============================================
// ================== Collection ==================

// OAPIF Bbox for extent
type Bbox struct {
	Crs    string    `json:"crs"`
	Extent []float64 `json:"bbox"`
}

// Extent OAPIF Extent structure (partial)
type CollectionExtent struct {
	Spatial *Bbox `json:"spatial"`
}

// CollectionsInfo for all collections
type CollectionsInfo struct {
	Links       []*Link           `json:"links"`
	Collections []*CollectionInfo `json:"collections"`
}

type Property struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

// CollectionInfo for a collection
type CollectionInfo struct {
	Name         string            `json:"id"`
	Title        string            `json:"title,omitempty"`
	Description  string            `json:"description,omitempty"`
	Extent       *CollectionExtent `json:"extent,omitempty"`
	Crs          []string          `json:"crs,omitempty"`
	GeometryType *string           `json:"geometrytype,omitempty"`

	// these are omitempty so they don't show in summary metadata
	Properties []*Property `json:"properties,omitempty"`

	Links []*Link `json:"links"`
	// used for HTML response only
	URLMetadataHTML string `json:"-"`
	URLMetadataJSON string `json:"-"`
	URLItemsHTML    string `json:"-"`
	URLItemsJSON    string `json:"-"`
}

func NewCollectionsInfo(tables []*Table) *CollectionsInfo {
	csDoc := CollectionsInfo{Links: []*Link{}, Collections: []*CollectionInfo{}}
	for _, tbl := range tables {
		collDoc := tbl.NewCollectionInfo()
		csDoc.Collections = append(csDoc.Collections, collDoc)
	}
	return &csDoc
}

// =============================================
// ================== Feature ==================

// Generic representation of Db data
type GeojsonFeatureData struct {
	Type             string                 `json:"type"`
	ID               string                 `json:"id,omitempty"`
	Geom             *geojson.Geometry      `json:"geometry"`
	Props            map[string]interface{} `json:"properties"`
	WeakEtag         string                 `json:"-"`
	LastModifiedDate string                 `json:"-"`
}

// Define a FeatureCollection structure for parsing test data
type FeatureCollection struct {
	Type           string                `json:"type"`
	Features       []*GeojsonFeatureData `json:"features"`
	NumberMatched  uint                  `json:"numberMatched,omitempty"`
	NumberReturned uint                  `json:"numberReturned"`
	TimeStamp      string                `json:"timeStamp,omitempty"`
	Links          []*Link               `json:"links"`
}

func MakeGeojsonFeatureJSON(id string, geom geojson.Geometry, props map[string]interface{}, weakEtag string, lastModifiedDate string) string {

	featData := MakeGeojsonFeature(id, geom, props, weakEtag, lastModifiedDate)
	json, err := json.Marshal(featData)
	if err != nil {
		log.Errorf("Error marshalling feature into JSON: %v", err)
		return ""
	}
	jsonStr := string(json)
	return jsonStr
}

func MakeGeojsonFeature(id string, geom geojson.Geometry, props map[string]interface{}, weakEtag string, lastModifiedHttpDate string) *GeojsonFeatureData {

	featData := GeojsonFeatureData{
		Type:             "Feature",
		ID:               id,
		Geom:             &geom,
		Props:            props,
		WeakEtag:         weakEtag,
		LastModifiedDate: lastModifiedHttpDate,
	}
	return &featData
}

func NewFeatureCollectionInfo(features []*GeojsonFeatureData) *FeatureCollection {
	ts := time.Now().Format(time.RFC3339)
	doc := FeatureCollection{
		Type:           GeoJSONFeatureCollection,
		Features:       features,
		NumberMatched:  0,
		NumberReturned: uint(len(features)),
		TimeStamp:      ts,
	}
	return &doc
}

// =================================================
// ================== Conformance ==================

type Conformance struct {
	ConformsTo []string `json:"conformsTo"`
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
		"http://www.opengis.net/spec/ogcapi-features-2/1.0/conf/crs",
		"http://www.opengis.net/spec/ogcapi-features-3/1.0/conf/filter",
		"http://www.opengis.net/spec/ogcapi-features-3/1.0/conf/features-filter",
		"http://www.opengis.net/spec/cql2/1.0/conf/basic-cql2",
		"http://www.opengis.net/spec/cql2/1.0/conf/basic-spatial-operators",
		"http://www.opengis.net/spec/cql2/1.0/conf/spatial-operators",
		"http://www.opengis.net/spec/cql2/1.0/conf/temporal-operators",
		"http://www.opengis.net/spec/cql2/1.0/conf/arithmetic",
		"http://www.opengis.net/spec/ogcapi-features-4/1.0/conf/create-replace-delete",
		"http://www.opengis.net/spec/ogcapi-features-4/1.0/conf/update",
	},
}

func GetConformance() *Conformance {
	return &conformance
}

// =================================================
// ================== Path helper ==================

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
