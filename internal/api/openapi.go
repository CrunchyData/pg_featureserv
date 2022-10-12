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
	"net/url"

	"github.com/CrunchyData/pg_featureserv/internal/conf"
	"github.com/getkin/kin-openapi/openapi3"
	log "github.com/sirupsen/logrus"
)

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

func getFeatureExample() map[string]interface{} {
	var result map[string]interface{}
	var jsonStr = `{"type":"Feature","geometry":{"type":"Point","coordinates":[-70.88461956597838,47.807897059236495]},"properties":{"prop_a":"propA","prop_b":1,"prop_c":"propC","prop_d":1}}`
	err := json.Unmarshal([]byte(jsonStr), &result)
	if err != nil {
		return nil
	}
	return result
}

var FeatureSchema openapi3.Schema = openapi3.Schema{
	Type:     "object",
	Required: []string{},
	Properties: map[string]*openapi3.SchemaRef{
		"id": {
			Value: &openapi3.Schema{
				OneOf: []*openapi3.SchemaRef{
					openapi3.NewSchemaRef("", &openapi3.Schema{
						Type: "number", Format: "long",
					}),
					openapi3.NewSchemaRef("", &openapi3.Schema{
						Type: "string",
					}),
				},
			},
		},
		"type": {
			Value: &openapi3.Schema{
				Type:    "string",
				Default: "Feature",
			},
		},
		"geometry": {
			Value: &openapi3.Schema{
				OneOf: []*openapi3.SchemaRef{
					openapi3.NewSchemaRef("https://geojson.org/schema/Point.json", &openapi3.Schema{}),
					openapi3.NewSchemaRef("https://geojson.org/schema/LineString.json", &openapi3.Schema{}),
					openapi3.NewSchemaRef("https://geojson.org/schema/Polygon.json", &openapi3.Schema{}),
					openapi3.NewSchemaRef("https://geojson.org/schema/MultiPoint.json", &openapi3.Schema{}),
					openapi3.NewSchemaRef("https://geojson.org/schema/MultiLineString.json", &openapi3.Schema{}),
					openapi3.NewSchemaRef("https://geojson.org/schema/MultiPolygon.json", &openapi3.Schema{}),
				},
			},
		},
		"properties": {
			Value: &openapi3.Schema{
				Type: "object",
			},
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
	Example: getFeatureExample(),
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

// GetOpenAPIContent returns a Swagger OpenAPI structure
func GetOpenAPIContent(urlBase string) *openapi3.Swagger {

	apiBase := "/"
	u, err := url.Parse(urlBase)
	if err == nil {
		apiBase = u.Path
	}
	log.Debugf("API base path = %v", apiBase)

	servername := u.Scheme + "://" + u.Host

	paramCollectionID := openapi3.ParameterRef{
		Value: &openapi3.Parameter{
			Description:     "ID of collection.",
			Name:            "collectionId",
			In:              "path",
			Required:        true,
			Schema:          &openapi3.SchemaRef{Value: openapi3.NewStringSchema()},
			AllowEmptyValue: false,
		},
	}
	paramFunctionID := openapi3.ParameterRef{
		Value: &openapi3.Parameter{
			Description:     "ID of function.",
			Name:            "functionId",
			In:              "path",
			Required:        true,
			Schema:          &openapi3.SchemaRef{Value: openapi3.NewStringSchema()},
			AllowEmptyValue: false,
		},
	}
	paramFeatureID := openapi3.ParameterRef{
		Value: &openapi3.Parameter{
			Name:            "featureId",
			Description:     "Id of feature in collection to retrieve data for.",
			In:              "path",
			Required:        true,
			Schema:          &openapi3.SchemaRef{Value: openapi3.NewStringSchema()},
			AllowEmptyValue: false,
		},
	}
	paramBbox := openapi3.ParameterRef{
		Value: &openapi3.Parameter{
			Name:        "bbox",
			Description: "Bounding box to restrict results to given extent (as minLon,minLat,maxLon,maxLat).",
			In:          "query",
			Required:    false,
			Explode:     openapi3.BoolPtr(false),
			Example:     "-120,30,-100,49",
			Schema: &openapi3.SchemaRef{
				Value: &openapi3.Schema{
					Type:     "array",
					MinItems: 4,
					MaxItems: openapi3.Uint64Ptr(4),
					Items:    openapi3.NewSchemaRef("", openapi3.NewFloat64Schema().WithMin(-180).WithMax(180)),
				},
			},
			AllowEmptyValue: false,
		},
	}
	paramBboxCrs := openapi3.ParameterRef{
		Value: &openapi3.Parameter{
			Name:        "bbox-crs",
			Description: "SRID for coordinate reference system of bbox parameter.",
			In:          "query",
			Required:    false,
			Schema: &openapi3.SchemaRef{
				Value: &openapi3.Schema{
					Type:    "integer",
					Min:     openapi3.Float64Ptr(1),
					Default: 4326,
				},
			},
			AllowEmptyValue: false,
		},
	}
	paramFilter := openapi3.ParameterRef{
		Value: &openapi3.Parameter{
			Name:            "filter",
			Description:     "CQL filter to apply.",
			In:              "query",
			Required:        false,
			Schema:          &openapi3.SchemaRef{Value: openapi3.NewStringSchema()},
			AllowEmptyValue: false,
		},
	}
	paramFilterCrs := openapi3.ParameterRef{
		Value: &openapi3.Parameter{
			Name:        "filter-crs",
			Description: "SRID for filter geometry literals.",
			In:          "query",
			Required:    false,
			Schema: &openapi3.SchemaRef{
				Value: &openapi3.Schema{
					Type:    "integer",
					Min:     openapi3.Float64Ptr(1),
					Default: 4326,
				},
			},
			AllowEmptyValue: false,
		},
	}
	paramProperties := openapi3.ParameterRef{
		Value: &openapi3.Parameter{
			Name:        "properties",
			Description: "List of properties to return in response objects",
			In:          "query",
			Required:    false,
			Explode:     openapi3.BoolPtr(false),
			Example:     "a,b,c",
			Schema: &openapi3.SchemaRef{
				Value: &openapi3.Schema{
					Type:     "array",
					MinItems: 0,
					Items:    &openapi3.SchemaRef{Value: openapi3.NewStringSchema()},
				},
			},
			AllowEmptyValue: false,
		},
	}
	paramTransform := openapi3.ParameterRef{
		Value: &openapi3.Parameter{
			Name:        "transform",
			Description: "Geometry transformation function pipeline to apply",
			In:          "query",
			Required:    false,
			Explode:     openapi3.BoolPtr(false),
			Style:       "pipeDelimited",
			Example:     "Centroid|Buffer,1",
			Schema: &openapi3.SchemaRef{
				Value: &openapi3.Schema{
					Type:     "array",
					MinItems: 0,
					Items:    &openapi3.SchemaRef{Value: openapi3.NewStringSchema()},
				},
			},
			AllowEmptyValue: false,
		},
	}
	paramCrs := openapi3.ParameterRef{
		Value: &openapi3.Parameter{
			Name:        "crs",
			Description: "SRID for coordinate reference system of output features.",
			In:          "query",
			Required:    false,
			Schema: &openapi3.SchemaRef{
				Value: &openapi3.Schema{
					Type:    "integer",
					Min:     openapi3.Float64Ptr(1),
					Default: 4326,
				},
			},
			AllowEmptyValue: false,
		},
	}
	paramLimit := openapi3.ParameterRef{
		Value: &openapi3.Parameter{
			Name:        "limit",
			Description: "Maximum number of results to return.",
			In:          "query",
			Required:    false,
			Schema: &openapi3.SchemaRef{
				Value: &openapi3.Schema{
					Type:    "integer",
					Min:     openapi3.Float64Ptr(0),
					Max:     openapi3.Float64Ptr(float64(conf.Configuration.Paging.LimitMax)),
					Default: conf.Configuration.Paging.LimitDefault,
				},
			},
			AllowEmptyValue: false,
		},
	}
	paramOffset := openapi3.ParameterRef{
		Value: &openapi3.Parameter{
			Name:        "offset",
			Description: "Offset of start of returned results.",
			In:          "query",
			Required:    false,
			Schema: &openapi3.SchemaRef{
				Value: &openapi3.Schema{
					Type: "integer",
					Min:  openapi3.Float64Ptr(0),
					//Max:     openapi3.Float64Ptr(float64(conf.Configuration.Paging.LimitMax)),
					Default: 0,
				},
			},
			AllowEmptyValue: false,
		},
	}
	paramSortBy := openapi3.ParameterRef{
		Value: &openapi3.Parameter{
			Name:            "sortby",
			Description:     "Column to sort by.",
			In:              "query",
			Required:        false,
			Schema:          &openapi3.SchemaRef{Value: openapi3.NewStringSchema()},
			AllowEmptyValue: false,
		},
	}
	paramType := openapi3.ParameterRef{
		Value: &openapi3.Parameter{
			Name:            "type",
			Description:     "Data schema type (create, update, etc.).",
			In:              "query",
			Required:        false,
			Schema:          &openapi3.SchemaRef{Value: openapi3.NewStringSchema()},
			AllowEmptyValue: false,
		},
	}
	return &openapi3.Swagger{
		OpenAPI: "3.0.0",
		Info: openapi3.Info{
			Title:       conf.Configuration.Metadata.Title,
			Description: conf.Configuration.Metadata.Description,
			Version:     conf.AppConfig.Version,
			License: &openapi3.License{
				Name: "Apache 2.0",
				URL:  "http://www.apache.org/licenses/LICENSE-2.0",
			},
		},
		Servers: openapi3.Servers{
			&openapi3.Server{
				URL: servername,
			},
		},
		Paths: openapi3.Paths{
			apiBase: &openapi3.PathItem{
				Summary:     "top-level endpoints available",
				Description: "Root of API, all metadata & services are beneath these links",
				Get: &openapi3.Operation{
					OperationID: "getRoot",
					Parameters:  openapi3.Parameters{},
					Responses: openapi3.Responses{
						"200": &openapi3.ResponseRef{
							Ref: "",
							Value: &openapi3.Response{
								Content:     openapi3.NewContentWithJSONSchema(&RootInfoSchema),
								Description: "Results for root of API",
							},
						},
					},
				},
			},
			apiBase + "api": &openapi3.PathItem{
				Summary:     "API definition",
				Description: "OpenAPI 3.0 definition of this service",
				Get: &openapi3.Operation{
					OperationID: "getAPI",
					Parameters:  openapi3.Parameters{},
					Responses: openapi3.Responses{
						"200": &openapi3.ResponseRef{
							// TODO: Find better OpenAPI schema ref?
							Ref: "https://json-schema.org/draft-07/schema",
						},
					},
				},
			},
			apiBase + "conformance": &openapi3.PathItem{
				Summary:     "Conformance classes",
				Description: "Functionality requirements this api conforms to.",
				Get: &openapi3.Operation{
					OperationID: "getConformance",
					Parameters:  openapi3.Parameters{},
					Responses: openapi3.Responses{
						"200": &openapi3.ResponseRef{
							Value: &openapi3.Response{
								Content:     openapi3.NewContentWithJSONSchema(&ConformanceSchema),
								Description: "Results for conformance classes",
							},
						},
					},
				},
			},
			apiBase + "collections": &openapi3.PathItem{
				Summary:     "Feature collections metadata",
				Description: "Provides details about feature collections served",
				Get: &openapi3.Operation{
					OperationID: "getCollectionsMetaData",
					Responses: openapi3.Responses{
						"200": &openapi3.ResponseRef{
							Value: &openapi3.Response{
								Content: openapi3.NewContentWithJSONSchemaRef(
									&openapi3.SchemaRef{Value: &CollectionsInfoSchema}),
								Description: "Results for details about the specified feature collection",
							},
						},
					},
				},
			},
			apiBase + "collections/{collectionId}": &openapi3.PathItem{
				Summary:     "Feature collection metadata",
				Description: "Provides details about the specified feature collection",
				Get: &openapi3.Operation{
					OperationID: "getCollectionMetaData",
					Parameters: openapi3.Parameters{
						&paramCollectionID},
					Responses: openapi3.Responses{
						"200": &openapi3.ResponseRef{
							Value: &openapi3.Response{
								Content: openapi3.NewContentWithJSONSchemaRef(
									&openapi3.SchemaRef{Value: &CollectionInfoSchema}),
								Description: "Results for details about the specified feature collection",
							},
						},
					},
				},
			},
			apiBase + "collections/{collectionId}/items": &openapi3.PathItem{
				Summary:     "Feature data for collection",
				Description: "Provides paged access to data for all features in specified collection",
				Get: &openapi3.Operation{
					OperationID: "getCollectionFeatures",
					Parameters: openapi3.Parameters{
						&paramCollectionID,
						&paramBbox,
						&paramBboxCrs,
						&paramFilter,
						&paramFilterCrs,
						&paramTransform,
						&paramProperties,
						&paramSortBy,
						&paramCrs,
						&paramLimit,
						&paramOffset,
						/* TODO
						&openapi3.ParameterRef{
							Value: &openapi3.Parameter{
								Name:            "<prop-name>",
								Description:     "Any feature property name may be filtered on by including it as a query parameter",
								In:              "query",
								Required:        false,
								Schema:          &openapi3.SchemaRef{Value: openapi3.NewStringSchema()},
								AllowEmptyValue: false,
							},
						},
						*/
					},
					Responses: openapi3.Responses{
						"200": &openapi3.ResponseRef{
							Value: &openapi3.Response{
								Description: "GeoJSON Feature Collection document containing data for features",
								/*
									// TODO: create schema for result?
									Content: openapi3.NewContentWithJSONSchemaRef(
										&openapi3.SchemaRef{
											Ref: "http://geojson.org/schema/FeatureCollection.json",
										},
									),
								*/
							},
						},
					},
				},
				Post: &openapi3.Operation{
					OperationID: "createCollectionFeature",
					Parameters: openapi3.Parameters{
						&paramCollectionID,
						// TODO keep it for the next evolution
						// &paramCrs,
					},
					RequestBody: &openapi3.RequestBodyRef{
						Value: &openapi3.RequestBody{
							Description: "Add a new feature",
							Required:    true,
							Content:     openapi3.NewContentWithJSONSchema(&FeatureSchema),
						},
					},
					Responses: openapi3.Responses{
						"201": &openapi3.ResponseRef{
							Value: &openapi3.Response{
								Description: "Empty body with location header",
								Headers: map[string]*openapi3.HeaderRef{
									"location": {
										Value: &openapi3.Header{
											Description: "Contains a link to access to the new feature data",
										},
									},
								},
							},
						},
					},
				},
			},
			apiBase + "collections/{collectionId}/schema": &openapi3.PathItem{
				Summary:     "Feature schema for collection",
				Description: "Provides access to data representation (schema) for any features in specified collection",
				Get: &openapi3.Operation{
					OperationID: "getCollectionSchema",
					Parameters: openapi3.Parameters{
						&paramCollectionID,
						&paramType,
					},
					Responses: openapi3.Responses{
						"200": &openapi3.ResponseRef{
							Value: &openapi3.Response{
								Description: "GeoJSON Feature Collection document containing data schema for specific type",
								/*
									// TODO: create schema for result?
									Content: openapi3.NewContentWithJSONSchemaRef(
										&openapi3.SchemaRef{
											Ref: "http://geojson.org/schema/FeatureSchema.json",
										},
									),
								*/
							},
						},
					},
				},
			},
			apiBase + "collections/{collectionId}/items/{featureId}": &openapi3.PathItem{
				Summary:     "Single feature data from collection",
				Description: "Provides access to a single feature identitfied by {featureId} from the specified collection",
				Get: &openapi3.Operation{
					OperationID: "getCollectionFeature",
					Parameters: openapi3.Parameters{
						&paramCollectionID,
						&paramFeatureID,
						&paramProperties,
						&paramTransform,
						&paramCrs,
					},
					Responses: openapi3.Responses{
						"200": &openapi3.ResponseRef{
							Value: &openapi3.Response{
								Description: "GeoJSON Feature document containing feature data",
								/*
									// TODO: create schema for result?
									Content: openapi3.NewContentWithJSONSchemaRef(
										&openapi3.SchemaRef{
											Ref: "http://geojson.org/schema/Feature.json",
										},
									),
								*/
							},
						},
					},
				},
				Patch: &openapi3.Operation{
					OperationID: "updateCollectionFeature",
					Parameters: openapi3.Parameters{
						&paramCollectionID,
						&paramFeatureID,
						// TODO keep it for the next evolution
						// &paramCrs,
					},
					RequestBody: &openapi3.RequestBodyRef{
						Value: &openapi3.RequestBody{
							Description: "Add a partial feature",
							Required:    true,
							Content:     openapi3.NewContentWithJSONSchema(&FeatureSchema),
						},
					},
					Responses: openapi3.Responses{
						"204": &openapi3.ResponseRef{
							Value: &openapi3.Response{
								Description: "No Content : feature updated",
							},
						},
						"404": &openapi3.ResponseRef{
							Value: &openapi3.Response{
								Description: "Resource not found",
							},
						},
					},
				},
				Put: &openapi3.Operation{
					OperationID: "replaceCollectionFeature",
					Parameters: openapi3.Parameters{
						&paramCollectionID,
						&paramFeatureID,
						&paramProperties,
						&paramTransform,
						&paramCrs,
					},
					RequestBody: &openapi3.RequestBodyRef{
						Value: &openapi3.RequestBody{
							Description: "Replace a feature",
							Required:    true,
							Content:     openapi3.NewContentWithJSONSchema(&FeatureSchema),
						},
					},
					Responses: openapi3.Responses{
						"204": &openapi3.ResponseRef{
							Value: &openapi3.Response{
								Description: "No Content: feature replaced succesfully",
							},
						},
						"400": &openapi3.ResponseRef{
							Value: &openapi3.Response{
								Description: "Malformed feature ID or unsuitable query parameters",
							},
						},
						"404": &openapi3.ResponseRef{
							Value: &openapi3.Response{
								Description: "Target resource not found",
							},
						},
					},
				},
				Delete: &openapi3.Operation{
					OperationID: "deleteCollectionFeature",
					Parameters: openapi3.Parameters{
						&paramCollectionID,
						&paramFeatureID,
					},
					Responses: openapi3.Responses{
						"204": &openapi3.ResponseRef{
							Value: &openapi3.Response{
								Description: "No Content",
							},
						},
						"400": &openapi3.ResponseRef{
							Value: &openapi3.Response{
								Description: "Malformed feature ID or unsuitable query parameters",
							},
						},
						"404": &openapi3.ResponseRef{
							Value: &openapi3.Response{
								Description: "Resource ID not found, collection not found, catalog error.",
							},
						},
					},
				},
			},
			apiBase + "functions": &openapi3.PathItem{
				Summary:     "Functions metadata",
				Description: "Provides details about functions served",
				Get: &openapi3.Operation{
					OperationID: "getFunctionsMetaData",
					Responses: openapi3.Responses{
						"200": &openapi3.ResponseRef{
							Value: &openapi3.Response{
								Content: openapi3.NewContentWithJSONSchemaRef(
									&openapi3.SchemaRef{Value: &FunctionsInfoSchema}),
								Description: "Results for details about functions served",
							},
						},
					},
				},
			},
			apiBase + "functions/{functionId}": &openapi3.PathItem{
				Summary:     "Function metadata",
				Description: "Provides details about the specified function",
				Get: &openapi3.Operation{
					OperationID: "getFunctionMetaData",
					Parameters: openapi3.Parameters{
						&paramFunctionID,
					},
					Responses: openapi3.Responses{
						"200": &openapi3.ResponseRef{
							Value: &openapi3.Response{
								Content: openapi3.NewContentWithJSONSchemaRef(
									&openapi3.SchemaRef{
										Value: &FunctionInfoSchema,
									}),
								Description: "Results for details about the specified function",
							},
						},
					},
				},
			},
			apiBase + "functions/{functionId}/items": &openapi3.PathItem{
				Summary:     "Features or data for a function result",
				Description: "Provides paged access to data in specified function result",
				Get: &openapi3.Operation{
					OperationID: "getFunctionFeatures",
					Parameters: openapi3.Parameters{
						&paramFunctionID,
						&paramBbox,
						&paramBboxCrs,
						&paramFilter,
						&paramFilterCrs,
						&paramTransform,
						&paramProperties,
						&paramSortBy,
						&paramCrs,
						&paramLimit,
						&paramOffset,

						/* TODO
						&openapi3.ParameterRef{
							Value: &openapi3.Parameter{
								Name:            "<prop-name>",
								Description:     "Any feature property name may be filtered on by including it as a query parameter",
								In:              "query",
								Required:        false,
								Schema:          &openapi3.SchemaRef{Value: openapi3.NewStringSchema()},
								AllowEmptyValue: false,
							},
						},
						*/
					},
					Responses: openapi3.Responses{
						"200": &openapi3.ResponseRef{
							Value: &openapi3.Response{
								Description: "GeoJSON or JSON document containing function results",
								/*
									Content: openapi3.NewContentWithJSONSchemaRef(
										&openapi3.SchemaRef{
											Ref: "http://geojson.org/schema/FeatureCollection.json",
										},
									),
								*/
							},
						},
					},
				},
			},
		},
	}
}
