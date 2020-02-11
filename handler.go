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
	"fmt"
	"net/http"
	"strings"

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

	addRoute(router, "/", handleHome)
	addRoute(router, "/home{.fmt}", handleHome)
	// consistent with pg_tileserv
	addRoute(router, "/index{.fmt}", handleHome)

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

func handleHome(w http.ResponseWriter, r *http.Request) *appError {
	format := api.RequestedFormat(r)
	return doRoot(w, r, format)
}

func doRoot(w http.ResponseWriter, r *http.Request, format string) *appError {
	//log.Printf("Content-Type: %v  Accept: %v", r.Header.Get("Content-Type"), r.Header.Get("Accept"))
	urlBase := serveURLBase(r)

	// --- create content
	content := api.NewRootInfo(&config.Configuration)
	content.Links = linksRoot(urlBase, format)

	switch format {
	case api.FormatHTML:
		context := ui.NewPageData()
		context.URLHome = urlPathFormat(urlBase, "", api.FormatHTML)
		context.URLJSON = urlPathFormat(urlBase, "", api.FormatJSON)

		return writeHTML(w, content, context, ui.PageHome())
	default:
		return writeJSON(w, api.ContentTypeJSON, content)
	}
}

func linksRoot(urlBase string, format string) []*api.Link {
	var links []*api.Link

	links = append(links, linkSelf(urlBase, "", format))
	links = append(links, linkAlt(urlBase, "", format))

	links = append(links, &api.Link{
		Href: urlPathFormat(urlBase, api.TagCollections, format),
		Rel:  api.RelData, Type: api.ContentType(format), Title: "collections"})
	links = append(links, &api.Link{
		Href: urlPathFormat(urlBase, api.TagFunctions, format),
		Rel:  api.RelData, Type: api.ContentType(format), Title: "functions"})

	return links
}

func linkSelf(urlBase string, path string, format string) *api.Link {
	return &api.Link{
		Href:  urlPathFormat(urlBase, path, format),
		Rel:   api.RelSelf,
		Type:  api.ContentType(format),
		Title: "This document as " + strings.ToUpper(format)}
}

func linkAlt(urlBase string, path string, format string) *api.Link {
	alt := altFormat(format)
	return &api.Link{
		Href:  urlPathFormat(urlBase, path, alt),
		Rel:   api.RelAlt,
		Type:  api.ContentType(alt),
		Title: "This document as " + strings.ToUpper(alt)}
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
	urlBase := serveURLBase(r)

	colls, err := catalogInstance.Tables()
	if err != nil {
		return appErrorInternal(err, api.ErrMsgLoadCollections)
	}

	content := api.NewCollectionsInfo(colls)
	content.Links = linksCollections(urlBase, format)
	for _, coll := range content.Collections {
		coll.Links = linksCollection(coll.Name, urlBase, format)
	}

	switch format {
	case api.FormatHTML:
		context := ui.NewPageData()
		context.URLHome = urlPathFormat(urlBase, "", api.FormatHTML)
		context.URLJSON = urlPathFormat(urlBase, api.TagCollections, api.FormatJSON)

		return writeHTML(w, content, context, ui.PageCollections())
	default:
		return writeJSON(w, api.ContentTypeJSON, content)
	}
}

func linksCollections(urlBase string, format string) []*api.Link {
	var links []*api.Link
	links = append(links, linkSelf(urlBase, api.TagCollections, format))
	links = append(links, linkAlt(urlBase, api.TagCollections, format))
	return links
}

func linksCollection(name string, urlBase string, format string) []*api.Link {
	path := fmt.Sprintf("%v/%v", api.TagCollections, name)
	pathItems := api.PathItems(api.TagCollections, name)

	var links []*api.Link
	links = append(links, linkSelf(urlBase, path, format))
	links = append(links, linkAlt(urlBase, path, format))

	links = append(links, &api.Link{
		Href: urlPathFormat(urlBase, pathItems, api.FormatJSON),
		Rel:  "items", Type: api.ContentTypeJSON, Title: "Features as GeoJSON"})
	links = append(links, &api.Link{
		Href: urlPathFormat(urlBase, pathItems, api.FormatHTML),
		Rel:  "items", Type: api.ContentTypeHTML, Title: "Features as HTML"})

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
	content.Links = linksCollection(name, urlBase, format)

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

	pathItems := api.PathItems(api.TagCollections, name)
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
	content.Links = linksItems(name, urlBase, api.FormatJSON)

	return writeJSON(w, api.ContentTypeGeoJSON, content)
}

func linksItems(name string, urlBase string, format string) []*api.Link {
	path := api.PathItems(api.TagCollections, name)

	var links []*api.Link
	links = append(links, linkSelf(urlBase, path, format))
	links = append(links, linkAlt(urlBase, path, format))

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

	pathItems := api.PathItems(api.TagCollections, name)
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
	urlBase := serveURLBase(r)

	fns, err := catalogInstance.Functions()
	if err != nil {
		return appErrorInternal(err, api.ErrMsgLoadFunctions)
	}

	content := api.NewFunctionsInfo(fns)
	content.Links = linksFunctions(urlBase, format)

	for _, fn := range content.Functions {
		fn.Links = linksFunction(fn.Name, urlBase, format, fn.Function.IsGeometryFunction())
	}

	switch format {
	case api.FormatHTML:
		context := ui.NewPageData()
		context.URLHome = urlPathFormat(urlBase, "", api.FormatHTML)
		context.URLJSON = urlPathFormat(urlBase, api.TagFunctions, api.FormatJSON)

		return writeHTML(w, content, context, ui.PageFunctions())
	default:
		return writeJSON(w, api.ContentTypeJSON, content)
	}
}

func linksFunctions(urlBase string, format string) []*api.Link {
	var links []*api.Link
	links = append(links, linkSelf(urlBase, api.TagFunctions, format))
	links = append(links, linkAlt(urlBase, api.TagFunctions, format))
	return links
}

func linksFunction(id string, urlBase string, format string, isGeomFun bool) []*api.Link {
	path := fmt.Sprintf("%v/%v", api.TagFunctions, id)
	pathItems := api.PathItems(api.TagFunctions, id)

	var links []*api.Link
	links = append(links, linkSelf(urlBase, path, format))
	links = append(links, linkAlt(urlBase, path, format))

	dataTitle := "Data as JSON"
	if isGeomFun {
		dataTitle = "Features as GeoJSON"
	}

	links = append(links, &api.Link{
		Href: urlPathFormat(urlBase, pathItems, api.FormatJSON),
		Rel:  "items", Type: api.ContentTypeJSON, Title: dataTitle})
	if isGeomFun {
		links = append(links, &api.Link{
			Href: urlPathFormat(urlBase, pathItems, api.FormatHTML),
			Rel:  "items", Type: api.ContentTypeHTML, Title: "Features as HTML"})
	}
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
	content.Links = linksFunction(name, urlBase, format, fn.IsGeometryFunction())

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
	pathItems := api.PathItems(api.TagFunctions, name)
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
	content.Links = linksItems(name, urlBase, api.FormatJSON)

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
