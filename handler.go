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
	varCollectionID = "cid"
	varFeatureID    = "fid"

	errMsgEncoding      = "Error encoding response"
	errMsgNoCollections = "No collections loaded"
	errMsgFailData      = "Error loading data"
)

func initRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	addRoute(router, "/", handleRootJSON)
	addRoute(router, "/home{.fmt}", handleHome)

	addRoute(router, "/conformance", handleConformance)
	addRoute(router, "/conformance.{fmt}", handleConformance)

	addRoute(router, "/collections", handleCollections)
	addRoute(router, "/collections.{fmt}", handleCollections)

	addRoute(router, "/collections/{cid}", handleCollection)
	addRoute(router, "/collections/{cid}.{fmt}", handleCollection)

	addRoute(router, "/collections/{cid}/items", handleCollectionItems)
	addRoute(router, "/collections/{cid}/items.{fmt}", handleCollectionItems)

	addRoute(router, "/collections/{cid}/items/{fid}", handleItem)
	addRoute(router, "/collections/{cid}/items/{fid}.{fmt}", handleItem)

	return router
}

func addRoute(router *mux.Router, path string, handler func(http.ResponseWriter, *http.Request) *appError) {
	router.Handle(path, appHandler(handler))
}

// ServeHTTP is the base Handler for routed requests.
// Common handling logic is placed here
func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// --- log the request
	log.Printf("%v Request: %v\n", r.RemoteAddr, r.URL)

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
	format := api.PathFormat(r.URL)
	return doRoot(w, r, format)
}

func doRoot(w http.ResponseWriter, r *http.Request, format string) *appError {
	urlBase := serveURLBase(r)

	// --- create content
	content := api.NewRootInfo(&config.Configuration)
	content.Links = linksRoot(urlBase, format)

	switch format {
	case api.FormatHTML:
		context := NewPageData()
		context.URLHome = urlPathFormat(urlBase, "", api.FormatHTML)
		context.URLJSON = urlPathFormat(urlBase, "", api.FormatJSON)

		return writeHTML(w, content, context, ui.HTMLTemplate.Home)
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
	adt := altFormat(format)
	return &api.Link{
		Href:  urlPathFormat(urlBase, path, adt),
		Rel:   api.RelAlt,
		Type:  api.ContentType(adt),
		Title: "This document as " + strings.ToUpper(adt)}
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

	colls, err := catalogInstance.Layers()
	if err != nil {
		return appErrorInternal(err, errMsgNoCollections)
	}

	content := api.NewCollectionsInfo(colls)
	content.Links = linksCollections(urlBase, format)
	for _, coll := range content.Collections {
		coll.Links = linksCollection(coll.Name, urlBase, format)
	}

	switch format {
	case api.FormatHTML:
		context := NewPageData()
		context.URLHome = urlPathFormat(urlBase, "", api.FormatHTML)
		context.URLJSON = urlPathFormat(urlBase, api.TagCollections, api.FormatJSON)

		return writeHTML(w, content, context, ui.HTMLTemplate.Collections)
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
	pathItems := api.PathItems(name)

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

	name := getRequestVar(varCollectionID, r)

	layer, err := catalogInstance.LayerByName(name)
	if layer == nil && err == nil {
		return appErrorNotFoundFmt(err, api.ErrMsgLayerNotFound, name)
	}
	content := api.NewCollectionInfo(layer)
	content.Links = linksCollection(name, urlBase, format)

	// --- encoding
	switch format {
	case api.FormatHTML:
		context := NewPageData()
		context.URLHome = urlPathFormat(urlBase, "", api.FormatHTML)
		context.URLCollections = urlPathFormat(urlBase, api.TagCollections, api.FormatHTML)
		context.URLCollection = urlPathFormat(urlBase, api.PathCollection(name), api.FormatHTML)
		context.URLJSON = urlPathFormat(urlBase, api.PathCollection(name), api.FormatJSON)
		context.CollectionTitle = layer.Title
		context.Layer = layer

		return writeHTML(w, content, context, ui.HTMLTemplate.Collection)
	default:
		return writeJSON(w, api.ContentTypeJSON, content)
	}
}

func handleCollectionItems(w http.ResponseWriter, r *http.Request) *appError {
	// TODO: determine content from request header?
	format := api.PathFormat(r.URL)
	urlBase := serveURLBase(r)

	//--- extract request parameters
	name := getRequestVar(varCollectionID, r)
	param := parseRequestParams(r)

	//param := NewQueryParam()
	switch format {
	case api.FormatJSON:
		return writeItemsJSON(w, name, param, urlBase)
	case api.FormatHTML:
		return writeItemsHTML(w, name, param, urlBase)
	}
	return nil
}

func writeItemsHTML(w http.ResponseWriter, name string, param data.QueryParam, urlBase string) *appError {
	//--- query data for request
	layer, err1 := catalogInstance.LayerByName(name)
	if err1 != nil {
		return appErrorInternal(err1, errMsgFailData)
	}
	if layer == nil {
		return appErrorNotFoundFmt(err1, api.ErrMsgLayerNotFound, name)
	}
	features, err2 := catalogInstance.LayerFeatures(name, param)
	if err2 != nil {
		return appErrorInternal(err2, errMsgFailData)
	}

	//--- assemble resonse
	content := api.NewFeatureCollectionInfo(features)

	// --- encoding
	context := NewPageData()
	context.URLHome = urlPathFormat(urlBase, "", api.FormatHTML)
	context.URLCollections = urlPathFormat(urlBase, api.TagCollections, api.FormatHTML)
	context.URLCollection = urlPathFormat(urlBase, api.PathCollection(name), api.FormatHTML)
	context.URLItems = urlPathFormat(urlBase, api.PathItems(name), api.FormatHTML)
	context.URLJSON = urlPathFormat(urlBase, api.PathItems(name), api.FormatJSON)
	context.CollectionTitle = layer.Title
	context.UseMap = true

	return writeHTML(w, content, context, ui.HTMLTemplate.Items)
}

func writeItemsJSON(w http.ResponseWriter, name string, param data.QueryParam, urlBase string) *appError {
	//--- query data for request
	features, err := catalogInstance.LayerFeatures(name, param)
	if err != nil {
		return appErrorInternal(err, errMsgFailData)
	}
	if features == nil {
		return appErrorNotFoundFmt(err, api.ErrMsgLayerNotFound, name)
	}

	//--- assemble resonse
	content := api.NewFeatureCollectionInfo(features)
	content.Links = linksItems(name, urlBase, api.FormatJSON)

	return writeJSON(w, api.ContentTypeGeoJSON, content)
}

func linksItems(name string, urlBase string, format string) []*api.Link {
	path := api.PathItems(name)

	var links []*api.Link
	links = append(links, linkSelf(urlBase, path, format))
	links = append(links, linkAlt(urlBase, path, format))

	return links
}

func handleItem(w http.ResponseWriter, r *http.Request) *appError {
	// TODO: determine content from request header?
	format := api.PathFormat(r.URL)
	urlBase := serveURLBase(r)

	//--- extract request parameters
	name := getRequestVar(varCollectionID, r)
	fid := getRequestVar(varFeatureID, r)
	param := parseRequestParams(r)

	switch format {
	case api.FormatJSON:
		return writeItemJSON(w, name, fid, param, urlBase)
	case api.FormatHTML:
		return writeItemHTML(w, name, fid, param, urlBase)
	}
	return nil
}

func writeItemHTML(w http.ResponseWriter, name string, fid string, param data.QueryParam, urlBase string) *appError {
	//--- query data for request
	layer, err1 := catalogInstance.LayerByName(name)
	if err1 != nil {
		return appErrorInternal(err1, errMsgFailData)
	}
	if layer == nil {
		return appErrorNotFoundFmt(err1, api.ErrMsgLayerNotFound, name)
	}
	feature, err2 := catalogInstance.LayerFeature(name, fid, param)
	if err2 != nil {
		return appErrorInternal(err2, errMsgFailData)
	}

	//--- assemble resonse
	content := feature

	// --- encoding
	context := NewPageData()
	context.URLHome = urlPathFormat(urlBase, "", api.FormatHTML)
	context.URLCollections = urlPathFormat(urlBase, api.TagCollections, api.FormatHTML)
	context.URLCollection = urlPathFormat(urlBase, api.PathCollection(name), api.FormatHTML)
	context.URLItems = urlPathFormat(urlBase, api.PathItems(name), api.FormatHTML)
	context.URLJSON = urlPathFormat(urlBase, api.PathItem(name, fid), api.FormatJSON)
	context.CollectionTitle = layer.Title
	context.FeatureID = fid
	context.UseMap = true

	return writeHTML(w, content, context, ui.HTMLTemplate.Item)
}

func writeItemJSON(w http.ResponseWriter, name string, fid string, param data.QueryParam, urlBase string) *appError {
	//--- query data for request
	feature, err := catalogInstance.LayerFeature(name, fid, param)
	if err != nil {
		return appErrorInternal(err, errMsgFailData)
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
		context := NewPageData()
		context.URLHome = urlPathFormat(urlBase, "", api.FormatHTML)
		context.URLJSON = urlPathFormat(urlBase, api.TagConformance, api.FormatJSON)

		return writeHTML(w, content, context, ui.HTMLTemplate.Conformance)
	default:
		return writeJSON(w, api.ContentTypeJSON, content)
	}
}
