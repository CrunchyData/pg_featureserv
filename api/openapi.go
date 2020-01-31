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
	"github.com/CrunchyData/pg_featureserv/config"
	"github.com/getkin/kin-openapi/openapi3"
)

func GetAPIContent() *openapi3.Swagger {
	return &openapi3.Swagger{
		OpenAPI: "3.0.0",
		Info: openapi3.Info{
			Title:       config.Configuration.Metadata.Title,
			Description: config.Configuration.Metadata.Description,
			Version:     "0.0.1",
			License: &openapi3.License{
				Name: "Apache 2.0",
				URL:  "http://www.apache.org/licenses/LICENSE-2.0",
			},
		},
		Paths: openapi3.Paths{
			"/": &openapi3.PathItem{
				Summary:     "top-level endpoints available",
				Description: "Root of API, all metadata & services are beneath these links",
				Get: &openapi3.Operation{
					OperationID: "getRoot",
					Parameters:  openapi3.Parameters{},
					Responses: openapi3.Responses{
						"200": &openapi3.ResponseRef{
							Ref: "",
							Value: &openapi3.Response{
								Content: openapi3.NewContentWithJSONSchema(&RootInfoSchema),
							},
						},
					},
				},
			},
			"/api": &openapi3.PathItem{
				Summary:     "api definition",
				Description: "OpenAPI 3.0 definition of this service",
				Get: &openapi3.Operation{
					OperationID: "getAPI",
					Parameters:  openapi3.Parameters{},
					Responses: openapi3.Responses{
						"200": &openapi3.ResponseRef{
							// TODO: Find better OpenAPI schema ref?
							Ref: "http://json-schema.org/draft-07/schema",
						},
					},
				},
			},
			"/conformance": &openapi3.PathItem{
				Summary:     "Conformance classes",
				Description: "Functionality requirements this api conforms to.",
				Get: &openapi3.Operation{
					OperationID: "getConformance",
					Parameters:  openapi3.Parameters{},
					Responses: openapi3.Responses{
						"200": &openapi3.ResponseRef{
							Value: &openapi3.Response{
								Content: openapi3.NewContentWithJSONSchema(&ConformanceSchema),
							},
						},
					},
				},
			},
			"/collections": &openapi3.PathItem{
				Summary:     "Feature collection metadata",
				Description: "Provides details about feature collections served",
				Get: &openapi3.Operation{
					OperationID: "getCollectionsMetaData",
					Responses: openapi3.Responses{
						"200": &openapi3.ResponseRef{
							Value: &openapi3.Response{
								Content: openapi3.NewContentWithJSONSchemaRef(
									&openapi3.SchemaRef{Value: &CollectionsInfoSchema}),
							},
						},
					},
				},
			},
			"/collections/{name}": &openapi3.PathItem{
				Summary:     "Feature collection metadata",
				Description: "Provides details about the feature collection named",
				Get: &openapi3.Operation{
					OperationID: "getCollectionMetaData",
					Parameters: openapi3.Parameters{
						&openapi3.ParameterRef{
							Value: &openapi3.Parameter{
								Description:     "Name of collection to retrieve metadata for.",
								Name:            "name",
								In:              "path",
								Required:        true,
								Schema:          &openapi3.SchemaRef{Value: openapi3.NewStringSchema()},
								AllowEmptyValue: false,
							},
						},
					},
					Responses: openapi3.Responses{
						"200": &openapi3.ResponseRef{
							Value: &openapi3.Response{
								Content: openapi3.NewContentWithJSONSchemaRef(
									&openapi3.SchemaRef{Value: &CollectionInfoSchema}),
							},
						},
					},
				},
			},
			"/collections/{name}/items": &openapi3.PathItem{
				Summary:     "Feature data for collection",
				Description: "Provides paged access to data for all features in collection",
				Get: &openapi3.Operation{
					OperationID: "getCollectionFeatures",
					Parameters: openapi3.Parameters{
						&openapi3.ParameterRef{
							Value: &openapi3.Parameter{
								Name:            "name",
								Description:     "Name of collection to retrieve data for.",
								In:              "path",
								Required:        true,
								Schema:          &openapi3.SchemaRef{Value: openapi3.NewStringSchema()},
								AllowEmptyValue: false,
							},
						},
						&openapi3.ParameterRef{
							Value: &openapi3.Parameter{
								Name:        "limit",
								Description: "Maximum number of results to return.",
								In:          "query",
								Required:    false,
								Schema: &openapi3.SchemaRef{
									Value: &openapi3.Schema{
										Type:    "integer",
										Min:     openapi3.Float64Ptr(1),
										Max:     openapi3.Float64Ptr(float64(config.Configuration.Paging.LimitMax)),
										Default: config.Configuration.Paging.LimitDefault,
									},
								},
								AllowEmptyValue: false,
							},
						},
						&openapi3.ParameterRef{
							Value: &openapi3.Parameter{
								Name:            "bbox",
								Description:     "Bounding box to limit results.",
								In:              "query",
								Required:        false,
								Schema:          &openapi3.SchemaRef{Value: &BboxSchema},
								AllowEmptyValue: false,
							},
						},
						&openapi3.ParameterRef{
							Value: &openapi3.Parameter{
								Name:            "<other>",
								Description:     "Any feature property name may be filtered on by including it as a query parameter",
								In:              "query",
								Required:        false,
								Schema:          &openapi3.SchemaRef{Value: openapi3.NewStringSchema()},
								AllowEmptyValue: false,
							},
						},
					},
					Responses: openapi3.Responses{
						"200": &openapi3.ResponseRef{
							Value: &openapi3.Response{
								Content: openapi3.NewContentWithJSONSchemaRef(
									&openapi3.SchemaRef{
										Ref: "http://geojson.org/schema/FeatureCollection.json",
									},
								),
							},
						},
					},
				},
			},
			"/collections/{name}/items/{feature_id}": &openapi3.PathItem{
				Summary:     "Single feature data from collection",
				Description: "Provides access to a single feature identitfied by {feature_id} from the named collection",
				Get: &openapi3.Operation{
					OperationID: "getCollectionFeature",
					Parameters: openapi3.Parameters{
						&openapi3.ParameterRef{
							Value: &openapi3.Parameter{
								Name:            "name",
								Description:     "Name of collection to retrieve data for.",
								In:              "path",
								Required:        true,
								Schema:          &openapi3.SchemaRef{Value: openapi3.NewStringSchema()},
								AllowEmptyValue: false,
							},
						},
						&openapi3.ParameterRef{
							Value: &openapi3.Parameter{
								Name:            "feature_id",
								Description:     "Id of feature in collection to retrieve data for.",
								In:              "path",
								Required:        true,
								Schema:          &openapi3.SchemaRef{Value: openapi3.NewStringSchema()},
								AllowEmptyValue: false,
							},
						},
					},
					Responses: openapi3.Responses{
						"200": &openapi3.ResponseRef{
							Value: &openapi3.Response{
								Content: openapi3.NewContentWithJSONSchemaRef(
									&openapi3.SchemaRef{
										Ref: "http://geojson.org/schema/Feature.json",
									},
								),
							},
						},
					},
				},
			},
			"/functions/{name}": &openapi3.PathItem{
				Summary:     "Function metadata",
				Description: "Provides details about the function named",
				Get: &openapi3.Operation{
					OperationID: "getFunctionMetaData",
					Parameters: openapi3.Parameters{
						&openapi3.ParameterRef{
							Value: &openapi3.Parameter{
								Description:     "Name of function to retrieve metadata for.",
								Name:            "name",
								In:              "path",
								Required:        true,
								Schema:          &openapi3.SchemaRef{Value: openapi3.NewStringSchema()},
								AllowEmptyValue: false,
							},
						},
					},
					Responses: openapi3.Responses{
						"200": &openapi3.ResponseRef{
							Value: &openapi3.Response{
								Content: openapi3.NewContentWithJSONSchemaRef(
									&openapi3.SchemaRef{
										Value: &CollectionInfoSchema,
									}),
							},
						},
					},
				},
			},
		},
	}
}
