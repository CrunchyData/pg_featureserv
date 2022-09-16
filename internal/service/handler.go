package service

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
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"

	"github.com/CrunchyData/pg_featureserv/internal/api"
	"github.com/CrunchyData/pg_featureserv/internal/conf"
	"github.com/CrunchyData/pg_featureserv/internal/data"
	"github.com/CrunchyData/pg_featureserv/internal/ui"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gorilla/mux"
)

const (
	routeVarID        = "id"
	routeVarFeatureID = "fid"
)

func InitRouter(basePath string) *mux.Router {
	router := mux.NewRouter().
		StrictSlash(true).
		PathPrefix("/" + strings.TrimRight(strings.TrimLeft(basePath, "/"), "/")).
		Subrouter()

	addRoute(router, "/", handleRoot)
	addRoute(router, "/home{.fmt}", handleRoot)
	// consistent with pg_tileserv
	addRoute(router, "/index{.fmt}", handleRoot)

	addRoute(router, "/api", handleAPI)
	addRoute(router, "/api.{fmt}", handleAPI)

	addRoute(router, "/conformance", handleConformance)
	addRoute(router, "/conformance.{fmt}", handleConformance)

	addRoute(router, "/collections", handleCollections)
	addRoute(router, "/collections.{fmt}", handleCollections)

	addRoute(router, "/collections/{id}", handleCollection)
	addRoute(router, "/collections/{id}.{fmt}", handleCollection)

	addRoute(router, "/collections/{id}/items", handleCollectionItems)
	addRoute(router, "/collections/{id}/items.{fmt}", handleCollectionItems)

	if conf.Configuration.Database.AllowWrite {
		addRouteWithMethod(router, "/collections/{id}/items", handleCreateCollectionItem, "POST")

		addRoute(router, "/collections/{id}/schema", handleCollectionSchemas)
	}

	addRoute(router, "/collections/{id}/items/{fid}", handleItem)
	addRoute(router, "/collections/{id}/items/{fid}.{fmt}", handleItem)

	addRouteWithMethod(router, "/collections/{id}/items/{fid}", handlePartialUpdateItem, "PATCH")
	addRouteWithMethod(router, "/collections/{id}/items/{fid}.{fmt}", handlePartialUpdateItem, "PATCH")

	addRouteWithMethod(router, "/collections/{id}/items/{fid}", handleReplaceItem, "PUT")
	addRouteWithMethod(router, "/collections/{id}/items/{fid}.{fmt}", handleReplaceItem, "PUT")

	addRoute(router, "/functions", handleFunctions)
	addRoute(router, "/functions.{fmt}", handleFunctions)

	addRoute(router, "/functions/{id}", handleFunction)
	addRoute(router, "/functions/{id}.{fmt}", handleFunction)

	addRoute(router, "/functions/{id}/items", handleFunctionItems)
	addRoute(router, "/functions/{id}/items.{fmt}", handleFunctionItems)

	return router
}

func addRoute(router *mux.Router, path string, handler func(http.ResponseWriter, *http.Request) *appError) {
	addRouteWithMethod(router, path, handler, "GET")
}

func addRouteWithMethod(router *mux.Router, path string, handler func(http.ResponseWriter, *http.Request) *appError, method string) {
	router.Handle(path, appHandler(handler)).Methods(method)
}

//nolint:unused
func handleRootJSON(w http.ResponseWriter, r *http.Request) *appError {
	return doRoot(w, r, api.FormatJSON)
}

func handleRoot(w http.ResponseWriter, r *http.Request) *appError {
	format := api.RequestedFormat(r)
	return doRoot(w, r, format)
}

func doRoot(w http.ResponseWriter, r *http.Request, format string) *appError {
	//log.Printf("Content-Type: %v  Accept: %v", r.Header.Get("Content-Type"), r.Header.Get("Accept"))
	urlBase := serveURLBase(r)

	// --- create content
	content := api.NewRootInfo(&conf.Configuration)

	switch format {
	case api.FormatHTML:
		context := ui.NewPageData()
		context.URLHome = urlPathFormat(urlBase, "", api.FormatHTML)
		context.URLJSON = urlPathFormat(urlBase, api.RootPageName, api.FormatJSON)

		return writeHTML(w, content, context, ui.PageHome())
	default:
		content.Links = linksRoot(urlBase)
		return writeJSON(w, api.ContentTypeJSON, content)
	}
}

func linksRoot(urlBase string) []*api.Link {
	var links []*api.Link
	links = append(links, linkSelf(urlBase, api.RootPageName, api.TitleDocument))
	links = append(links, linkAlt(urlBase, api.RootPageName, api.TitleDocument))

	links = append(links, &api.Link{
		Href: urlPath(urlBase, api.TagAPI),
		Rel:  api.RelServiceDesc, Type: api.ContentTypeOpenAPI, Title: "API definition"})
	links = append(links, &api.Link{
		Href: urlPath(urlBase, api.TagConformance),
		Rel:  api.RelConformance, Type: api.ContentTypeJSON, Title: "OGC API conformance classes implemented by this server"})
	links = append(links, &api.Link{
		Href: urlPath(urlBase, api.TagCollections),
		Rel:  api.RelData, Type: api.ContentTypeJSON, Title: "collections"})
	links = append(links, &api.Link{
		Href: urlPath(urlBase, api.TagFunctions),
		Rel:  api.RelFunctions, Type: api.ContentTypeJSON, Title: "functions"})

	return links
}

func linkSelf(urlBase string, path string, desc string) *api.Link {
	return &api.Link{
		Href:  urlPath(urlBase, path),
		Rel:   api.RelSelf,
		Type:  api.ContentTypeJSON,
		Title: desc + api.TitleAsJSON}
}

func linkAlt(urlBase string, path string, desc string) *api.Link {
	return &api.Link{
		Href:  urlPathFormat(urlBase, path, api.FormatHTML),
		Rel:   api.RelAlt,
		Type:  api.ContentTypeHTML,
		Title: desc + api.TitleAsHTML}
}

func handleCollections(w http.ResponseWriter, r *http.Request) *appError {
	format := api.RequestedFormat(r)
	urlBase := serveURLBase(r)

	colls, err := catalogInstance.Tables()
	if err != nil {
		return appErrorInternal(err, api.ErrMsgLoadCollections)
	}

	content := api.NewCollectionsInfo(colls)
	for _, coll := range content.Collections {
		switch format {
		case api.FormatHTML:
			addCollectionURLs(coll, urlBase)
		default:
			coll.Links = linksCollection(coll.Name, urlBase, true)
		}
	}

	switch format {
	case api.FormatHTML:
		context := ui.NewPageData()
		context.URLHome = urlPathFormat(urlBase, "", api.FormatHTML)
		context.URLJSON = urlPathFormat(urlBase, api.TagCollections, api.FormatJSON)

		return writeHTML(w, content, context, ui.PageCollections())
	default:
		content.Links = linksCollections(urlBase)
		return writeJSON(w, api.ContentTypeJSON, content)
	}
}

func linksCollections(urlBase string) []*api.Link {
	var links []*api.Link
	links = append(links, linkSelf(urlBase, api.TagCollections, api.TitleDocument))
	links = append(links, linkAlt(urlBase, api.TagCollections, api.TitleDocument))
	return links
}

func addCollectionURLs(coll *api.CollectionInfo, urlBase string) {
	name := coll.Name
	path := api.PathCollection(name)
	pathItems := api.PathCollectionItems(name)

	coll.URLMetadataJSON = urlPathFormat(urlBase, path, api.FormatJSON)
	coll.URLMetadataHTML = urlPathFormat(urlBase, path, api.FormatHTML)
	coll.URLItemsHTML = urlPathFormat(urlBase, pathItems, api.FormatHTML)
	coll.URLItemsJSON = urlPathFormat(urlBase, pathItems, api.FormatJSON)
}

func linksCollection(name string, urlBase string, isSummary bool) []*api.Link {
	path := api.PathCollection(name)
	pathItems := api.PathCollectionItems(name)

	titleDesc := api.TitleDocument
	if isSummary {
		titleDesc = api.TitleMetadata
	}

	var links []*api.Link
	links = append(links, linkSelf(urlBase, path, titleDesc))
	links = append(links, linkAlt(urlBase, path, titleDesc))

	linkItemsJSON := urlPath(urlBase, pathItems)
	links = append(links, &api.Link{
		Href:  linkItemsJSON,
		Rel:   api.RelItems,
		Type:  api.ContentTypeGeoJSON,
		Title: api.TitleFeatuuresGeoJSON})

	return links
}

func handleCollection(w http.ResponseWriter, r *http.Request) *appError {
	format := api.RequestedFormat(r)
	urlBase := serveURLBase(r)

	name := getRequestVar(routeVarID, r)

	tbl, err := catalogInstance.TableByName(name)
	if tbl == nil && err == nil {
		return appErrorNotFoundFmt(err, api.ErrMsgCollectionNotFound, name)
	}
	catalogInstance.TableReload(name)
	content := api.NewCollectionInfo(tbl)
	content.GeometryType = &tbl.GeometryType
	content.Properties = api.TableProperties(tbl)

	// --- encoding
	switch format {
	case api.FormatHTML:
		pathItems := api.PathCollectionItems(name)
		context := ui.NewPageData()
		context.URLHome = urlPathFormat(urlBase, "", api.FormatHTML)
		context.URLCollections = urlPathFormat(urlBase, api.TagCollections, api.FormatHTML)
		context.URLCollection = urlPathFormat(urlBase, api.PathCollection(name), api.FormatHTML)
		context.URLJSON = urlPathFormat(urlBase, api.PathCollection(name), api.FormatJSON)
		context.URLItems = urlPathFormat(urlBase, pathItems, api.FormatHTML)
		context.URLItemsJSON = urlPathFormat(urlBase, pathItems, api.FormatJSON)
		context.Title = tbl.Title
		context.Table = tbl
		context.IDColumn = tbl.IDColumn

		return writeHTML(w, content, context, ui.PageCollection())
	default:
		content.Links = linksCollection(name, urlBase, false)
		return writeJSON(w, api.ContentTypeJSON, content)
	}
}

func handleCollectionSchemas(w http.ResponseWriter, r *http.Request) *appError {
	// TODO: determine content from request header?
	format := api.RequestedFormat(r)

	//--- extract request parameters
	name := getRequestVar(routeVarID, r)
	tbl, err1 := catalogInstance.TableByName(name)
	if err1 != nil {
		return appErrorInternalFmt(err1, api.ErrMsgCollectionAccess, name)
	}
	if tbl == nil {
		return appErrorNotFoundFmt(err1, api.ErrMsgCollectionNotFound, name)
	}

	queryValues := r.URL.Query()
	paramValues := extractSingleArgs(queryValues)
	// --- type parameter
	schemaType := parseString(paramValues, api.ParamType)

	ctx := r.Context()
	switch format {
	case api.FormatSchemaJSON:
		{
			switch schemaType {
			case "create", "replace":
				// The "replace" schema is identical to the "create" schema.
				// http://docs.ogc.org/DRAFTS/20-002.html#feature-geojson
				return writeCreateItemSchemaJSON(ctx, w, tbl)
			case "update":
				return writeUpdateItemSchemaJSON(ctx, w, tbl)
			default:
				return appErrorBadRequest(nil, fmt.Sprintf("Asked schema type %s not implemented!", schemaType))
			}
		}
	default:
		{
			return appErrorBadRequest(nil, fmt.Sprintf("Format %s not implemented!", format))
		}
	}
}

func handleCreateCollectionItem(w http.ResponseWriter, r *http.Request) *appError {
	urlBase := serveURLBase(r)

	//--- extract request parameters
	name := getRequestVar(routeVarID, r)

	//--- check query parameters
	queryValues := r.URL.Query()
	paramValues := extractSingleArgs(queryValues)
	if len(paramValues) != 0 {
		return appErrorMsg(nil, "No parameter allowed", http.StatusBadRequest)
	}

	//--- check feature availability
	tbl, err1 := catalogInstance.TableByName(name)
	if err1 != nil {
		return appErrorInternalFmt(err1, api.ErrMsgCollectionAccess, name)
	}
	if tbl == nil {
		return appErrorNotFoundFmt(err1, api.ErrMsgCollectionNotFound, name)
	}

	//--- json body
	bodyContent, errBody := ioutil.ReadAll(r.Body)
	if errBody != nil || len(bodyContent) == 0 {
		return appErrorInternalFmt(errBody, "Unable to read request body for Collection: %v", name)
	}

	//--- check if body matches the schema
	createSchema, errGetSch := getCreateItemSchema(r.Context(), tbl)
	if errGetSch != nil {
		return appErrorInternalFmt(errGetSch, errGetSch.Error())
	}
	var val interface{}
	_ = json.Unmarshal(bodyContent, &val)
	errValSch := createSchema.VisitJSON(val)
	if errValSch != nil {
		return appErrorInternalFmt(errValSch, api.ErrMsgCreateFeatureNotConform, name)
	}

	newId, err2 := catalogInstance.AddTableFeature(r.Context(), name, bodyContent)
	if err2 != nil {
		return appErrorInternalFmt(err2, api.ErrMsgCreateFeatureInCatalog, name)
	}

	w.Header().Set("Location", fmt.Sprintf("%scollections/%s/items/%d", urlBase, name, newId))
	w.WriteHeader(http.StatusCreated)
	return nil
}

func handleCollectionItems(w http.ResponseWriter, r *http.Request) *appError {
	// TODO: determine content from request header?
	format := api.RequestedFormat(r)
	urlBase := serveURLBase(r)
	query := api.URLQuery(r.URL)

	//--- extract request parameters
	name := getRequestVar(routeVarID, r)
	reqParam, err := parseRequestParams(r)
	if err != nil {
		return appErrorMsg(err, err.Error(), http.StatusBadRequest)
	}

	tbl, err1 := catalogInstance.TableByName(name)
	if err1 != nil {
		return appErrorInternalFmt(err1, api.ErrMsgCollectionAccess, name)
	}
	if tbl == nil {
		return appErrorNotFoundFmt(err1, api.ErrMsgCollectionNotFound, name)
	}
	param, err := createQueryParams(&reqParam, tbl.Columns, tbl.Srid)
	if err != nil {
		return appErrorBadRequest(err, err.Error())
	}
	param.Filter = parseFilter(reqParam.Values, tbl.DbTypes)

	ctx := r.Context()
	switch format {
	case api.FormatJSON:
		return writeItemsJSON(ctx, w, name, param, urlBase)
	case api.FormatHTML:
		return writeItemsHTML(w, tbl, name, query, urlBase)
	}
	return nil
}

func writeCreateItemSchemaJSON(ctx context.Context, w http.ResponseWriter, table *data.Table) *appError {
	createSchema, err := getCreateItemSchema(ctx, table)
	if err != nil {
		return appErrorMsg(err, err.Error(), http.StatusInternalServerError)
	}
	return writeJSON(w, api.ContentTypeSchemaJSON, createSchema)
}

func writeUpdateItemSchemaJSON(ctx context.Context, w http.ResponseWriter, table *data.Table) *appError {
	updateSchema, err := getUpdateItemSchema(ctx, table)
	if err != nil {
		return appErrorMsg(err, err.Error(), http.StatusInternalServerError)
	}
	return writeJSON(w, api.ContentTypeSchemaPatchJSON, updateSchema)
}

func getCreateItemSchema(ctx context.Context, table *data.Table) (openapi3.Schema, error) {
	// Feature schema skeleton
	var featureInfoSchema openapi3.Schema = openapi3.Schema{
		Type:     "object",
		Required: []string{"type", "geometry", "properties"},
		Properties: map[string]*openapi3.SchemaRef{
			"id": {Value: &openapi3.Schema{Type: "string", Format: "uri"}},
			"type": {
				Value: &openapi3.Schema{
					Type:    "string",
					Default: "Feature",
				},
			},
			"geometry": {
				Value: &openapi3.Schema{
					Items: &openapi3.SchemaRef{
						Ref: fmt.Sprintf("https://geojson.org/schema/%v.json", table.GeometryType),
						Value: &openapi3.Schema{
							Type: "string", // mandatory to validate the schema
						},
					},
				},
			},
			"properties": {
				Value: &openapi3.Schema{},
			},
		},
	}
	featureInfoSchema.Description = table.Description

	props := featureInfoSchema.Properties["properties"].Value
	props.Type = "object"

	// update required properties
	requiredTypeKeys := make([]string, 0, len(table.DbTypes))

	for k := range table.DbTypes {
		if k != table.IDColumn {
			requiredTypeKeys = append(requiredTypeKeys, k)
		}
	}
	sort.Strings(requiredTypeKeys)

	var requiredTypes []string
	for _, k := range requiredTypeKeys {
		if table.DbTypes[k].IsRequired {
			requiredTypes = append(requiredTypes, k)
		}
	}

	props.Required = requiredTypes

	// update properties by their name and type
	props.Properties = make(map[string]*openapi3.SchemaRef)
	for k, v := range table.DbTypes {
		if k != table.IDColumn {
			propType := v.Type
			if api.Db2OpenapiFormatMap[v.Type] != "" {
				propType = api.Db2OpenapiFormatMap[v.Type]
			}
			props.Properties[k] = &openapi3.SchemaRef{
				Value: &openapi3.Schema{
					Type: propType,
				},
			}
		}
	}

	errVal := featureInfoSchema.Validate(ctx)
	if errVal != nil {
		encodedContent, _ := json.Marshal(featureInfoSchema)
		return featureInfoSchema, fmt.Errorf("schema not valid: %v\n\t%v", errVal, string(encodedContent))
	}

	return featureInfoSchema, nil
}

func getUpdateItemSchema(ctx context.Context, table *data.Table) (openapi3.Schema, error) {
	// Feature schema skeleton
	var featureInfoSchema openapi3.Schema = openapi3.Schema{
		Type: "object",
		Properties: map[string]*openapi3.SchemaRef{
			"type": {
				Value: &openapi3.Schema{
					Type:    "string",
					Default: "Feature",
				},
			},
			"geometry": {
				Value: &openapi3.Schema{
					Items: &openapi3.SchemaRef{
						Ref: fmt.Sprintf("https://geojson.org/schema/%v.json", table.GeometryType),
						Value: &openapi3.Schema{
							Type: "string", // mandatory to validate the schema
						},
					},
				},
			},
			"properties": {
				Value: &openapi3.Schema{},
			},
		},
	}
	featureInfoSchema.Description = table.Description

	props := featureInfoSchema.Properties["properties"].Value
	props.Type = "object"

	// update properties by their name and type
	props.Properties = make(map[string]*openapi3.SchemaRef)
	for k, v := range table.DbTypes {
		if k != table.IDColumn {
			propType := v.Type
			if api.Db2OpenapiFormatMap[v.Type] != "" {
				propType = api.Db2OpenapiFormatMap[v.Type]
			}
			props.Properties[k] = &openapi3.SchemaRef{
				Value: openapi3.NewOneOfSchema(
					&openapi3.Schema{
						Type: propType,
					},
					&openapi3.Schema{
						Type: "null",
					},
				),
			}
		}
	}

	errVal := featureInfoSchema.Validate(ctx)
	if errVal != nil {
		encodedContent, _ := json.Marshal(featureInfoSchema)
		return featureInfoSchema, fmt.Errorf("schema not valid: %v\n\t%v", errVal, string(encodedContent))
	}

	return featureInfoSchema, nil
}

func writeItemsHTML(w http.ResponseWriter, tbl *data.Table, name string, query string, urlBase string) *appError {

	pathItems := api.PathCollectionItems(name)
	// --- encoding
	context := ui.NewPageData()
	context.URLHome = urlPathFormat(urlBase, "", api.FormatHTML)
	context.URLCollections = urlPathFormat(urlBase, api.TagCollections, api.FormatHTML)
	context.URLCollection = urlPathFormat(urlBase, api.PathCollection(name), api.FormatHTML)
	context.URLItems = urlPathFormatQuery(urlBase, pathItems, api.FormatHTML, query)
	context.URLJSON = urlPathFormatQuery(urlBase, pathItems, api.FormatJSON, query)
	context.Group = "Collections"
	context.Title = tbl.Title
	context.IDColumn = tbl.IDColumn
	context.ShowFeatureLink = true

	// features are not needed for items page (page queries for them)
	return writeHTML(w, nil, context, ui.PageItems())
}

func writeItemsJSON(ctx context.Context, w http.ResponseWriter, name string, param *data.QueryParam, urlBase string) *appError {
	//--- query features data
	features, err := catalogInstance.TableFeatures(ctx, name, param)
	if err != nil {
		return appErrorInternalFmt(err, api.ErrMsgDataReadError, name)
	}
	if features == nil {
		return appErrorNotFoundFmt(err, api.ErrMsgCollectionNotFound, name)
	}

	//--- assemble resonse
	content := api.NewFeatureCollectionInfo(features)
	content.Links = linksItems(name, urlBase)

	return writeJSON(w, api.ContentTypeGeoJSON, content)
}

func linksItems(name string, urlBase string) []*api.Link {
	path := api.PathCollectionItems(name)

	var links []*api.Link
	links = append(links, linkSelf(urlBase, path, api.TitleDocument))
	links = append(links, linkAlt(urlBase, path, api.TitleDocument))

	return links
}

func handleItem(w http.ResponseWriter, r *http.Request) *appError {
	// TODO: determine content from request header?
	format := api.RequestedFormat(r)
	urlBase := serveURLBase(r)

	query := api.URLQuery(r.URL)
	//--- extract request parameters
	name := getRequestVar(routeVarID, r)
	fid := getRequestVar(routeVarFeatureID, r)
	reqParam, err := parseRequestParams(r)
	if err != nil {
		return appErrorMsg(err, err.Error(), http.StatusBadRequest)
	}

	tbl, err1 := catalogInstance.TableByName(name)
	if err1 != nil {
		return appErrorInternalFmt(err1, api.ErrMsgCollectionAccess, name)
	}
	if tbl == nil {
		return appErrorNotFoundFmt(err1, api.ErrMsgCollectionNotFound, name)
	}
	param, errQuery := createQueryParams(&reqParam, tbl.Columns, tbl.Srid)

	if errQuery == nil {
		ctx := r.Context()
		switch format {
		case api.FormatJSON:
			return writeItemJSON(ctx, w, name, fid, param, urlBase)
		case api.FormatHTML:
			return writeItemHTML(w, tbl, name, fid, query, urlBase)
		default:
			return nil
		}
	} else {
		return appErrorInternalFmt(errQuery, api.ErrMsgInvalidQuery)
	}
}

func handlePartialUpdateItem(w http.ResponseWriter, r *http.Request) *appError {
	urlBase := serveURLBase(r)

	// extract request parameters
	name := getRequestVar(routeVarID, r)
	fid := getRequestVar(routeVarFeatureID, r)

	// check query parameters
	queryValues := r.URL.Query()
	paramValues := extractSingleArgs(queryValues)
	if len(paramValues) != 0 {
		return appErrorMsg(nil, "No parameter allowed", http.StatusBadRequest)
	}

	// check that collection exists
	tbl, err1 := catalogInstance.TableByName(name)
	if err1 != nil {
		return appErrorInternalFmt(err1, api.ErrMsgCollectionAccess, name)
	}
	if tbl == nil {
		return appErrorNotFoundFmt(err1, api.ErrMsgCollectionNotFound, name)
	}

	// extract JSON from request body
	body, errBody := ioutil.ReadAll(r.Body)
	if errBody != nil || len(body) == 0 {
		return appErrorInternalFmt(errBody, "Unable to read request body for Collection: %v", name)
	}

	// check schema
	var val map[string]interface{}
	errUnMarsh := json.Unmarshal(body, &val)
	if errUnMarsh != nil {
		return appErrorInternalFmt(errUnMarsh, "Json not conform")
	}

	errSchema := api.FeatureSchema.VisitJSONObject(val)
	if errSchema != nil {
		return appErrorInternalFmt(errSchema, "Data not respect schema: %v", name)
	}

	check, errChck := tbl.CheckTableFields(val)
	if !check && errChck != nil {
		return appErrorInternalFmt(errChck, "validation error")
	}

	// perform update in database
	id, err := catalogInstance.PartialUpdateTableFeature(r.Context(), name, fid, body)
	if err != nil {
		return appErrorInternalFmt(err, api.ErrMsgPartialUpdateFeature, name)
	}

	w.Header().Set("Location", fmt.Sprintf("%scollections/%s/items/%d", urlBase, name, id))
	w.WriteHeader(http.StatusNoContent)

	return nil
}

func handleReplaceItem(w http.ResponseWriter, r *http.Request) *appError {
	// extract request parameters
	name := getRequestVar(routeVarID, r)
	fid := getRequestVar(routeVarFeatureID, r)

	// check query parameters
	queryValues := r.URL.Query()
	paramValues := extractSingleArgs(queryValues)
	if len(paramValues) != 0 {
		return appErrorMsg(nil, "No parameter allowed", http.StatusBadRequest)
	}

	// check that collection exists
	tbl, err1 := catalogInstance.TableByName(name)
	if err1 != nil {
		return appErrorInternalFmt(err1, api.ErrMsgCollectionAccess, name)
	}
	if tbl == nil {
		return appErrorNotFoundFmt(err1, api.ErrMsgCollectionNotFound, name)
	}

	// extract JSON from request body
	body, errBody := ioutil.ReadAll(r.Body)
	if errBody != nil || len(body) == 0 {
		return appErrorInternalFmt(errBody, "Unable to read request body for Collection: %v", name)
	}

	//--- check if body matches the schema
	// schema for replace is the same as in create
	createSchema, errGetSch := getCreateItemSchema(r.Context(), tbl)
	if errGetSch != nil {
		return appErrorInternalFmt(errGetSch, errGetSch.Error())
	}
	var val interface{}
	_ = json.Unmarshal(body, &val)
	errValSch := createSchema.VisitJSON(val)
	if errValSch != nil {
		return appErrorBadRequest(errValSch, api.ErrMsgReplaceFeatureNotConform)
	}

	// perform replace in database
	err2 := catalogInstance.ReplaceTableFeature(r.Context(), name, fid, body)
	if err2 != nil {
		return appErrorInternalFmt(err2, api.ErrMsgReplaceFeature, name)
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}

func writeItemHTML(w http.ResponseWriter, tbl *data.Table, name string, fid string, query string, urlBase string) *appError {
	//--- query data for request

	pathItems := api.PathCollectionItems(name)
	// --- encoding
	context := ui.NewPageData()
	context.URLHome = urlPathFormat(urlBase, "", api.FormatHTML)
	context.URLCollections = urlPathFormat(urlBase, api.TagCollections, api.FormatHTML)
	context.URLCollection = urlPathFormat(urlBase, api.PathCollection(name), api.FormatHTML)
	context.URLItems = urlPathFormat(urlBase, pathItems, api.FormatHTML)
	context.URLJSON = urlPathFormatQuery(urlBase, api.PathItem(name, fid), api.FormatJSON, query)
	context.Group = "Collections"
	context.Title = tbl.Title
	context.FeatureID = fid
	context.IDColumn = tbl.IDColumn

	// feature is not needed for item page (page queries for them)
	return writeHTML(w, nil, context, ui.PageItem())
}

func writeItemJSON(ctx context.Context, w http.ResponseWriter, name string, fid string, param *data.QueryParam, urlBase string) *appError {
	//--- query data for request
	feature, err := catalogInstance.TableFeature(ctx, name, fid, param)
	if err != nil {
		return appErrorInternalFmt(err, api.ErrMsgDataReadError, name)
	}
	if len(feature) == 0 {
		return appErrorNotFoundFmt(nil, api.ErrMsgFeatureNotFound, fid)
	}

	//--- assemble resonse
	//content := feature
	// for now can't add links to feature JSON
	//content.Links = linksItems(name, urlBase, api.FormatJSON)
	encodedContent := []byte(feature)
	writeResponse(w, api.ContentTypeGeoJSON, encodedContent)
	return nil
}

func handleConformance(w http.ResponseWriter, r *http.Request) *appError {
	// TODO: determine content from request header?
	format := api.RequestedFormat(r)
	urlBase := serveURLBase(r)

	content := api.GetConformance()

	switch format {
	case api.FormatHTML:
		context := ui.NewPageData()
		context.URLHome = urlPathFormat(urlBase, "", api.FormatHTML)
		context.URLJSON = urlPathFormat(urlBase, api.TagConformance, api.FormatJSON)

		return writeHTML(w, content, context, ui.PageConformance())
	default:
		return writeJSON(w, api.ContentTypeJSON, content)
	}
}

func handleAPI(w http.ResponseWriter, r *http.Request) *appError {
	// TODO: determine content from request header?
	format := api.RequestedFormat(r)
	urlBase := serveURLBase(r)

	content := api.GetOpenAPIContent(urlBase)

	switch format {
	case api.FormatHTML:
		context := ui.NewPageData()
		context.URLHome = urlPathFormat(urlBase, "", api.FormatHTML)
		context.URLJSON = urlPathFormat(urlBase, api.TagAPI, api.FormatJSON)

		return writeHTML(w, content, context, ui.PageAPI())
	default:
		return writeJSON(w, api.ContentTypeJSON, content)
	}
}

func handleFunctions(w http.ResponseWriter, r *http.Request) *appError {
	format := api.RequestedFormat(r)
	urlBase := serveURLBase(r)

	fns, err := catalogInstance.Functions()
	if err != nil {
		return appErrorInternal(err, api.ErrMsgLoadFunctions)
	}
	content := api.NewFunctionsInfo(fns)
	for _, fn := range content.Functions {
		isGeomFun := fn.Function.IsGeometryFunction()
		switch format {
		case api.FormatHTML:
			addFunctionURLs(fn, urlBase, isGeomFun)
		default:
			fn.Links = linksFunction(fn.Function.ID, urlBase, true, isGeomFun)
		}
	}

	switch format {
	case api.FormatHTML:
		context := ui.NewPageData()
		context.URLHome = urlPathFormat(urlBase, "", api.FormatHTML)
		context.URLJSON = urlPathFormat(urlBase, api.TagFunctions, api.FormatJSON)

		return writeHTML(w, content, context, ui.PageFunctions())
	default:
		content.Links = linksFunctions(urlBase)
		return writeJSON(w, api.ContentTypeJSON, content)
	}
}

func linksFunctions(urlBase string) []*api.Link {
	var links []*api.Link
	links = append(links, linkSelf(urlBase, api.TagFunctions, api.TitleDocument))
	links = append(links, linkAlt(urlBase, api.TagFunctions, api.TitleDocument))
	return links
}

func addFunctionURLs(content *api.FunctionSummary, urlBase string, isGeomFun bool) {
	name := content.Name
	path := api.PathFunction(name)
	pathItems := api.PathFunctionItems(name)
	content.URLMetadataJSON = urlPathFormat(urlBase, path, api.FormatJSON)
	content.URLMetadataHTML = urlPathFormat(urlBase, path, api.FormatHTML)
	// there is no HTML view for non-spatial (for now)
	content.URLItemsHTML = urlIfExists(isGeomFun, urlPathFormat(urlBase, pathItems, api.FormatHTML))
	content.URLItemsJSON = urlPathFormat(urlBase, pathItems, api.FormatJSON)
}

// urlIfExists returns the url, or blank if the document does not exist
func urlIfExists(isExists bool, url string) string {
	if isExists {
		return url
	}
	return ""
}

func linksFunction(id string, urlBase string, isSummary bool, isGeomFun bool) []*api.Link {
	path := api.PathFunction(id)
	pathItems := api.PathFunctionItems(id)

	titleDesc := api.TitleDocument
	if isSummary {
		titleDesc = api.TitleMetadata
	}

	var links []*api.Link
	links = append(links, linkSelf(urlBase, path, titleDesc))
	links = append(links, linkAlt(urlBase, path, titleDesc))

	dataTitle := api.TitleDataJSON
	conType := api.ContentTypeJSON
	if isGeomFun {
		dataTitle = api.TitleFeatuuresGeoJSON
		conType = api.ContentTypeGeoJSON
	}

	links = append(links, &api.Link{
		Href:  urlPath(urlBase, pathItems),
		Rel:   "items",
		Type:  conType,
		Title: dataTitle})
	return links
}

func handleFunction(w http.ResponseWriter, r *http.Request) *appError {
	format := api.RequestedFormat(r)
	urlBase := serveURLBase(r)

	shortName := getRequestVar(routeVarID, r)
	name := data.FunctionQualifiedId(shortName)

	fn, err := catalogInstance.FunctionByName(name)
	if fn == nil && err == nil {
		return appErrorNotFoundFmt(err, api.ErrMsgFunctionNotFound, name)
	}
	content := api.NewFunctionInfo(fn)
	isGeomFun := fn.IsGeometryFunction()
	content.Parameters = api.FunctionParameters(fn)
	content.Properties = api.FunctionProperties(fn)

	// --- encoding
	switch format {
	case api.FormatHTML:
		pathItems := api.PathFunctionItems(shortName)
		context := ui.NewPageData()
		context.URLHome = urlPathFormat(urlBase, "", api.FormatHTML)
		context.URLFunctions = urlPathFormat(urlBase, api.TagFunctions, api.FormatHTML)
		context.URLFunction = urlPathFormat(urlBase, api.PathFunction(name), api.FormatHTML)
		context.URLJSON = urlPathFormat(urlBase, api.PathFunction(name), api.FormatJSON)
		// there is no HTML view for non-spatial (for now)
		context.URLItems = urlIfExists(isGeomFun, urlPathFormat(urlBase, pathItems, api.FormatHTML))
		context.URLItemsJSON = urlPathFormat(urlBase, pathItems, api.FormatJSON)
		context.Title = fn.ID
		context.Function = fn

		return writeHTML(w, content, context, ui.PageFunction())
	default:
		content.Links = linksFunction(shortName, urlBase, false, isGeomFun)

		return writeJSON(w, api.ContentTypeJSON, content)
	}
}

func handleFunctionItems(w http.ResponseWriter, r *http.Request) *appError {
	// TODO: determine content from request header?
	format := api.RequestedFormat(r)
	urlBase := serveURLBase(r)

	//--- extract request parameters
	name := data.FunctionQualifiedId(getRequestVar(routeVarID, r))
	reqParam, err := parseRequestParams(r)
	if err != nil {
		return appErrorMsg(err, err.Error(), http.StatusBadRequest)
	}
	query := api.URLQuery(r.URL)

	fn, err := catalogInstance.FunctionByName(name)
	if fn == nil && err == nil {
		return appErrorNotFoundFmt(err, api.ErrMsgFunctionNotFound, name)
	}
	param, err := createQueryParams(&reqParam, fn.OutNames, data.SRID_4326)
	if err != nil {
		return appErrorBadRequest(err, err.Error())
	}
	fnArgs := restrict(reqParam.Values, fn.InNames)
	//log.Debugf("Function request args: %v ", fnArgs)

	ctx := r.Context()
	switch format {
	case api.FormatJSON:
		if fn.IsGeometryFunction() {
			return writeFunItemsGeoJSON(ctx, w, name, fnArgs, param, urlBase)
		}
		return writeFunItemsJSON(ctx, w, name, fnArgs, param)
	case api.FormatHTML:
		return writeFunItemsHTML(w, name, query, urlBase)
	case api.FormatText:
		return writeFunItemsText(ctx, w, api.ContentTypeText, name, fnArgs, param)
	case api.FormatSVG:
		return writeFunItemsText(ctx, w, api.ContentTypeSVG, name, fnArgs, param)
	}
	return nil
}

func writeFunItemsHTML(w http.ResponseWriter, name string, query string, urlBase string) *appError {
	fn, err1 := catalogInstance.FunctionByName(name)
	if err1 != nil {
		return appErrorInternalFmt(err1, api.ErrMsgFunctionAccess, name)
	}
	if fn == nil {
		return appErrorNotFoundFmt(err1, api.ErrMsgFunctionNotFound, name)
	}
	pathItems := api.PathFunctionItems(name)
	// --- encoding
	context := ui.NewPageData()
	context.URLHome = urlPathFormat(urlBase, "", api.FormatHTML)
	context.URLCollections = urlPathFormat(urlBase, api.TagFunctions, api.FormatHTML)
	context.URLCollection = urlPathFormat(urlBase, api.PathFunction(name), api.FormatHTML)
	context.URLItems = urlPathFormatQuery(urlBase, pathItems, api.FormatHTML, query)
	context.URLJSON = urlPathFormatQuery(urlBase, pathItems, api.FormatJSON, query)
	context.Group = "Functions"
	context.Title = fn.ID
	context.Function = fn
	context.IDColumn = data.FunctionIDColumnName

	// features are not needed for items page (page queries for them)
	return writeHTML(w, nil, context, ui.PageFunctionItems())
}

func writeFunItemsGeoJSON(ctx context.Context, w http.ResponseWriter, name string, args map[string]string, param *data.QueryParam, urlBase string) *appError {
	//--- query features data
	features, err := catalogInstance.FunctionFeatures(ctx, name, args, param)
	if err != nil {
		return appErrorInternalFmt(err, api.ErrMsgDataReadError, name)
	}
	if features == nil {
		return appErrorNotFoundFmt(err, api.ErrMsgNoDataRead, name)
	}

	//--- assemble resonse
	content := api.NewFeatureCollectionInfo(features)
	content.Links = linksItems(name, urlBase)

	return writeJSON(w, api.ContentTypeGeoJSON, content)
}

func writeFunItemsJSON(ctx context.Context, w http.ResponseWriter, name string, args map[string]string, param *data.QueryParam) *appError {
	//--- query features data
	features, err := catalogInstance.FunctionData(ctx, name, args, param)
	if err != nil {
		return appErrorInternalFmt(err, api.ErrMsgFunctionAccess, name)
	}
	if features == nil {
		return appErrorNotFoundFmt(err, api.ErrMsgNoDataRead, name)
	}
	return writeJSON(w, api.ContentTypeJSON, features)
}

func writeFunItemsText(ctx context.Context, w http.ResponseWriter, contentType string, name string, args map[string]string, param *data.QueryParam) *appError {
	//--- query features data
	features, err := catalogInstance.FunctionData(ctx, name, args, param)
	if err != nil {
		return appErrorInternalFmt(err, api.ErrMsgFunctionAccess, name)
	}
	if features == nil {
		return appErrorNotFoundFmt(err, api.ErrMsgNoDataRead, name)
	}
	content := writeFeaturesToByte(features)
	return writeText(w, contentType, content)
}

func writeFeaturesToByte(features []map[string]interface{}) []byte {
	var b bytes.Buffer
	for _, feat := range features {
		for _, val := range feat {
			s := fmt.Sprintf("%v", val)
			b.WriteString(s)
		}
	}
	return b.Bytes()
}
