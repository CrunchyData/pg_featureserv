package main

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
	"net/http"

	"github.com/CrunchyData/pg_featureserv/api"
	"github.com/CrunchyData/pg_featureserv/config"
	"github.com/CrunchyData/pg_featureserv/data"
	"github.com/CrunchyData/pg_featureserv/ui"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

const (
	routeVarID        = "id"
	routeVarFeatureID = "fid"
)

func initRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

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

	addRoute(router, "/collections/{id}/items/{fid}", handleItem)
	addRoute(router, "/collections/{id}/items/{fid}.{fmt}", handleItem)

	addRoute(router, "/functions", handleFunctions)
	addRoute(router, "/functions.{fmt}", handleFunctions)

	addRoute(router, "/functions/{id}", handleFunction)
	addRoute(router, "/functions/{id}.{fmt}", handleFunction)

	addRoute(router, "/functions/{id}/items", handleFunctionItems)
	addRoute(router, "/functions/{id}/items.{fmt}", handleFunctionItems)

	return router
}

func addRoute(router *mux.Router, path string, handler func(http.ResponseWriter, *http.Request) *appError) {
	router.Handle(path, appHandler(handler))
}

// ServeHTTP is the base Handler for routed requests.
// Common handling logic is placed here
func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// --- log the request
	log.Printf("%v %v %v\n", r.RemoteAddr, r.Method, r.URL)

	// execute the handler
	e := fn(w, r)

	if e != nil { // e is *appError, not os.Error.
		// TODO: is this the desire behaviour?
		// perhaps detect format and emit accordingly?
		// log error here?
		// should log attached error?
		// panic on severe error?
		http.Error(w, e.Message, e.Code)
	}
}

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
	content := api.NewRootInfo(&config.Configuration)

	switch format {
	case api.FormatHTML:
		context := ui.NewPageData()
		context.URLHome = urlPathFormat(urlBase, "", api.FormatHTML)
		context.URLJSON = urlPathFormat(urlBase, "", api.FormatJSON)

		return writeHTML(w, content, context, ui.PageHome())
	default:
		content.Links = linksRoot(urlBase)
		return writeJSON(w, api.ContentTypeJSON, content)
	}
}

func linksRoot(urlBase string) []*api.Link {
	var links []*api.Link
	format := api.FormatJSON
	links = append(links, linkSelf(urlBase, "", api.TitleDocument))
	links = append(links, linkAlt(urlBase, "", api.TitleDocument))

	links = append(links, &api.Link{
		Href: urlPathFormat(urlBase, api.TagCollections, format),
		Rel:  api.RelData, Type: api.ContentType(format), Title: "collections"})
	links = append(links, &api.Link{
		Href: urlPathFormat(urlBase, api.TagFunctions, format),
		Rel:  api.RelFunctions, Type: api.ContentType(format), Title: "functions"})

	return links
}

func linkSelf(urlBase string, path string, desc string) *api.Link {
	return &api.Link{
		Href:  urlPathFormat(urlBase, path, api.FormatJSON),
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

func altFormat(format string) string {
	switch format {
	case api.FormatJSON:
		return api.FormatHTML
	case api.FormatHTML:
		return api.FormatJSON
	}
	// TODO: panic here?
	return ""
}

func handleCollections(w http.ResponseWriter, r *http.Request) *appError {
	format := api.PathFormat(r.URL)
	isJSON := format == api.FormatJSON
	urlBase := serveURLBase(r)

	colls, err := catalogInstance.Tables()
	if err != nil {
		return appErrorInternal(err, api.ErrMsgLoadCollections)
	}

	content := api.NewCollectionsInfo(colls)
	for _, coll := range content.Collections {
		addCollectionLinks(coll, urlBase, isJSON, true)
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

func addCollectionLinks(coll *api.CollectionInfo, urlBase string, isJSON bool, isInList bool) {
	name := coll.Name
	path := api.PathCollection(name)
	pathItems := api.PathCollectionItems(name)

	if isJSON {
		coll.Links = linksCollection(name, urlBase, isInList)
	} else {
		coll.URLMetadataJSON = urlPathFormat(urlBase, path, api.FormatJSON)
		coll.URLMetadataHTML = urlPathFormat(urlBase, path, api.FormatHTML)
		coll.URLItemsHTML = urlPathFormat(urlBase, pathItems, api.FormatHTML)
	}
}

func linksCollection(name string, urlBase string, isInList bool) []*api.Link {
	path := api.PathCollection(name)
	pathItems := api.PathCollectionItems(name)

	titleDesc := api.TitleDocument
	if isInList {
		titleDesc = api.TitleMetadata
	}

	var links []*api.Link
	links = append(links, linkSelf(urlBase, path, titleDesc))
	links = append(links, linkAlt(urlBase, path, titleDesc))

	linkItemsJSON := urlPathFormat(urlBase, pathItems, api.FormatJSON)
	links = append(links, &api.Link{
		Href:  linkItemsJSON,
		Rel:   api.RelItems,
		Type:  api.ContentTypeGeoJSON,
		Title: api.TitleFeatuuresGeoJSON})

	return links
}

func handleCollection(w http.ResponseWriter, r *http.Request) *appError {
	format := api.PathFormat(r.URL)
	urlBase := serveURLBase(r)

	name := getRequestVar(routeVarID, r)

	tbl, err := catalogInstance.TableByName(name)
	if tbl == nil && err == nil {
		return appErrorNotFoundFmt(err, api.ErrMsgCollectionNotFound, name)
	}
	content := api.NewCollectionInfo(tbl)
	isJSON := format == api.FormatJSON
	addCollectionLinks(content, urlBase, isJSON, false)

	// --- encoding
	switch format {
	case api.FormatHTML:
		context := ui.NewPageData()
		context.URLHome = urlPathFormat(urlBase, "", api.FormatHTML)
		context.URLCollections = urlPathFormat(urlBase, api.TagCollections, api.FormatHTML)
		context.URLCollection = urlPathFormat(urlBase, api.PathCollection(name), api.FormatHTML)
		context.URLJSON = urlPathFormat(urlBase, api.PathCollection(name), api.FormatJSON)
		context.Title = tbl.Title
		context.Table = tbl
		context.IDColumn = tbl.IDColumn

		return writeHTML(w, content, context, ui.PageCollection())
	default:
		return writeJSON(w, api.ContentTypeJSON, content)
	}
}

func handleCollectionItems(w http.ResponseWriter, r *http.Request) *appError {
	// TODO: determine content from request header?
	format := api.PathFormat(r.URL)
	urlBase := serveURLBase(r)
	query := api.URLQuery(r.URL)

	//--- extract request parameters
	name := getRequestVar(routeVarID, r)
	param, err := parseRequestParams(r)
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
	param.Columns = normalizePropNames(param.Properties, tbl.Columns)

	switch format {
	case api.FormatJSON:
		return writeItemsJSON(w, name, param, urlBase)
	case api.FormatHTML:
		return writeItemsHTML(w, tbl, name, query, urlBase)
	}
	return nil
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

func writeItemsJSON(w http.ResponseWriter, name string, param data.QueryParam, urlBase string) *appError {
	//--- query features data
	features, err := catalogInstance.TableFeatures(name, param)
	if err != nil {
		return appErrorInternalFmt(err, api.ErrMsgDataRead, name)
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
	format := api.PathFormat(r.URL)
	urlBase := serveURLBase(r)

	query := api.URLQuery(r.URL)
	//--- extract request parameters
	name := getRequestVar(routeVarID, r)
	fid := getRequestVar(routeVarFeatureID, r)
	param, err := parseRequestParams(r)
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

	param.Columns = normalizePropNames(param.Properties, tbl.Columns)

	switch format {
	case api.FormatJSON:
		return writeItemJSON(w, name, fid, param, urlBase)
	case api.FormatHTML:
		return writeItemHTML(w, tbl, name, fid, query, urlBase)
	}
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

	// feature is not needed for item page (page queries for them)
	return writeHTML(w, nil, context, ui.PageItem())
}

func writeItemJSON(w http.ResponseWriter, name string, fid string, param data.QueryParam, urlBase string) *appError {
	//--- query data for request
	feature, err := catalogInstance.TableFeature(name, fid, param)
	if err != nil {
		return appErrorInternalFmt(err, api.ErrMsgDataRead, name)
	}
	if len(feature) == 0 {
		return appErrorNotFoundFmt(nil, api.ErrCodeFeatureNotFound, fid)
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
	format := api.PathFormat(r.URL)
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
	format := api.PathFormat(r.URL)
	urlBase := serveURLBase(r)

	content := api.GetAPIContent()

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
	format := api.PathFormat(r.URL)
	isJSON := format == api.FormatJSON
	urlBase := serveURLBase(r)

	fns, err := catalogInstance.Functions()
	if err != nil {
		return appErrorInternal(err, api.ErrMsgLoadFunctions)
	}
	content := api.NewFunctionsInfo(fns)
	for _, fn := range content.Functions {
		isGeomFun := fn.Function.IsGeometryFunction()
		addFunctionLinks(fn, urlBase, isJSON, true, isGeomFun)
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

func addFunctionLinks(content *api.FunctionInfo, urlBase string, isJSON bool, isInList bool, isGeomFun bool) {
	name := content.Name
	if isJSON {
		content.Links = linksFunction(name, urlBase, isInList, isGeomFun)
	} else {
		path := api.PathFunction(name)
		pathItems := api.PathFunctionItems(name)
		content.URLMetadataJSON = urlPathFormat(urlBase, path, api.FormatJSON)
		content.URLMetadataHTML = urlPathFormat(urlBase, path, api.FormatHTML)
		content.URLItemsHTML = urlPathFormat(urlBase, pathItems, api.FormatHTML)
		content.URLItemsJSON = urlPathFormat(urlBase, pathItems, api.FormatJSON)
	}
}

func linksFunction(id string, urlBase string, isInList bool, isGeomFun bool) []*api.Link {
	path := api.PathFunction(id)
	pathItems := api.PathFunctionItems(id)

	titleDesc := api.TitleDocument
	if isInList {
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
		Href:  urlPathFormat(urlBase, pathItems, api.FormatJSON),
		Rel:   "items",
		Type:  conType,
		Title: dataTitle})
	return links
}

func handleFunction(w http.ResponseWriter, r *http.Request) *appError {
	format := api.PathFormat(r.URL)
	urlBase := serveURLBase(r)

	name := getRequestVar(routeVarID, r)

	fn, err := catalogInstance.FunctionByName(name)
	if fn == nil && err == nil {
		return appErrorNotFoundFmt(err, api.ErrMsgFunctionNotFound, name)
	}
	content := api.NewFunctionInfo(fn)
	isGeomFun := fn.IsGeometryFunction()
	isJSON := format == api.FormatJSON
	addFunctionLinks(content, urlBase, isJSON, false, isGeomFun)
	// --- encoding
	switch format {
	case api.FormatHTML:
		context := ui.NewPageData()
		context.URLHome = urlPathFormat(urlBase, "", api.FormatHTML)
		context.URLFunctions = urlPathFormat(urlBase, api.TagFunctions, api.FormatHTML)
		context.URLFunction = urlPathFormat(urlBase, api.PathFunction(name), api.FormatHTML)
		context.URLJSON = urlPathFormat(urlBase, api.PathFunction(name), api.FormatJSON)
		context.Title = fn.ID
		context.Function = fn

		return writeHTML(w, content, context, ui.PageFunction())
	default:

		return writeJSON(w, api.ContentTypeJSON, content)
	}
}

func handleFunctionItems(w http.ResponseWriter, r *http.Request) *appError {
	// TODO: determine content from request header?
	format := api.PathFormat(r.URL)
	urlBase := serveURLBase(r)

	//--- extract request parameters
	name := getRequestVar(routeVarID, r)
	param, err := parseRequestParams(r)
	if err != nil {
		return appErrorMsg(err, err.Error(), http.StatusBadRequest)
	}
	query := api.URLQuery(r.URL)

	fn, err := catalogInstance.FunctionByName(name)
	if fn == nil && err == nil {
		return appErrorNotFoundFmt(err, api.ErrMsgFunctionNotFound, name)
	}
	param.Columns = normalizePropNames(param.Properties, fn.OutNames)

	switch format {
	case api.FormatJSON:
		if fn.IsGeometryFunction() {
			return writeFunItemsGeoJSON(w, name, param, urlBase)
		}
		return writeFunItemsJSON(w, name, param)
	case api.FormatHTML:
		return writeFunItemsHTML(w, name, query, urlBase)
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

func writeFunItemsGeoJSON(w http.ResponseWriter, name string, param data.QueryParam, urlBase string) *appError {
	//--- query features data
	features, err := catalogInstance.FunctionFeatures(name, param)
	if err != nil {
		return appErrorInternalFmt(err, api.ErrMsgDataRead, name)
	}
	if features == nil {
		return appErrorNotFoundFmt(err, api.ErrMsgDataRead, name)
	}

	//--- assemble resonse
	content := api.NewFeatureCollectionInfo(features)
	content.Links = linksItems(name, urlBase)

	return writeJSON(w, api.ContentTypeGeoJSON, content)
}

func writeFunItemsJSON(w http.ResponseWriter, name string, param data.QueryParam) *appError {
	//--- query features data
	features, err := catalogInstance.FunctionData(name, param)
	if err != nil {
		return appErrorInternalFmt(err, api.ErrMsgFunctionAccess, name)
	}
	if features == nil {
		return appErrorNotFoundFmt(err, api.ErrMsgDataRead, name)
	}
	return writeJSON(w, api.ContentTypeJSON, features)
}
