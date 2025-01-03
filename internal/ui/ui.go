package ui

/*
 Copyright 2019 - 2025 Crunchy Data Solutions, Inc.
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
	"html/template"

	"github.com/CrunchyData/pg_featureserv/internal/conf"
	"github.com/CrunchyData/pg_featureserv/internal/data"
	log "github.com/sirupsen/logrus"
)

// PageData - data used on the HTML pages
type PageData struct {
	AppName    string
	AppVersion string
	// URLHome - URL for the service home page
	URLHome        string
	URLCollections string
	URLCollection  string
	URLFunctions   string
	URLFunction    string
	// URLItems is url for items HTML page
	URLItems string
	// URLItemsJSON is url for items JSON
	URLItemsJSON string
	// URLJSON is the url for the current page in JSON
	URLJSON         string
	Group           string
	Title           string
	Table           *data.Table
	IDColumn        string
	Function        *data.Function
	FeatureID       string
	ShowFeatureLink bool
}

var htmlTemp struct {
	home          *template.Template
	conformance   *template.Template
	api           *template.Template
	collections   *template.Template
	collection    *template.Template
	items         *template.Template
	item          *template.Template
	functions     *template.Template
	function      *template.Template
	functionItems *template.Template
}

// HTMLDynamicLoad sets whether HTML templates are loaded every time they are used
// this allows rapid prototyping of changes
var HTMLDynamicLoad bool

func init() {
	HTMLDynamicLoad = false
}

// NewPageData create a page context initialized with globals.
func NewPageData() *PageData {
	con := PageData{}
	con.AppName = conf.AppConfig.Name
	con.AppVersion = conf.AppConfig.Version
	return &con
}

func createTemplate(filename ...string) *template.Template {
	tmpl, err := template.ParseFiles(filename...)
	if err != nil {
		log.Fatalf("Failure loading templates from %v: %v", filename, err)
	}
	return tmpl
}

func loadTemplate(curr *template.Template, filename ...string) *template.Template {
	if curr != nil && !HTMLDynamicLoad {
		return curr
	}
	return createTemplate(filename...)
}

func loadPageTemplate(curr *template.Template, filename string) *template.Template {
	if curr != nil && !HTMLDynamicLoad {
		return curr
	}
	return createTemplate(
		conf.Configuration.Server.AssetsPath+"/page.gohtml",
		conf.Configuration.Server.AssetsPath+"/"+filename)
}

func loadMapPageTemplate(curr *template.Template, filename string) *template.Template {
	if curr != nil && !HTMLDynamicLoad {
		return curr
	}
	return createTemplate(
		conf.Configuration.Server.AssetsPath+"/page.gohtml",
		conf.Configuration.Server.AssetsPath+"/map_script.gohtml",
		conf.Configuration.Server.AssetsPath+"/"+filename)
}

func PageHome() *template.Template {
	htmlTemp.home = loadPageTemplate(htmlTemp.home, "home.gohtml")
	return htmlTemp.home
}
func PageConformance() *template.Template {
	htmlTemp.conformance = loadPageTemplate(htmlTemp.conformance, "conformance.gohtml")
	return htmlTemp.conformance
}
func PageAPI() *template.Template {
	htmlTemp.api = loadTemplate(htmlTemp.api, conf.Configuration.Server.AssetsPath+"/api.gohtml")
	return htmlTemp.api
}
func PageCollections() *template.Template {
	htmlTemp.collections = loadPageTemplate(htmlTemp.collections, "collections.gohtml")
	return htmlTemp.collections
}
func PageCollection() *template.Template {
	htmlTemp.collection = loadPageTemplate(htmlTemp.collection, "collection.gohtml")
	return htmlTemp.collection
}
func PageItems() *template.Template {
	htmlTemp.items = loadMapPageTemplate(htmlTemp.items, "items.gohtml")
	return htmlTemp.items
}
func PageItem() *template.Template {
	htmlTemp.item = loadMapPageTemplate(htmlTemp.item, "item.gohtml")
	return htmlTemp.item
}
func PageFunctions() *template.Template {
	htmlTemp.functions = loadPageTemplate(htmlTemp.functions, "functions.gohtml")
	return htmlTemp.functions
}
func PageFunction() *template.Template {
	htmlTemp.function = loadPageTemplate(htmlTemp.function, "function.gohtml")
	return htmlTemp.function
}
func PageFunctionItems() *template.Template {
	htmlTemp.functionItems = loadTemplate(htmlTemp.functionItems,
		conf.Configuration.Server.AssetsPath+"/page.gohtml",
		conf.Configuration.Server.AssetsPath+"/items.gohtml",
		conf.Configuration.Server.AssetsPath+"/map_script.gohtml",
		conf.Configuration.Server.AssetsPath+"/fun_script.gohtml")
	return htmlTemp.functionItems
}

// RenderHTML tbd
func RenderHTML(temp *template.Template, content interface{}, context interface{}) ([]byte, error) {
	bodyData := map[string]interface{}{
		"config":  conf.Configuration,
		"context": context,
		"data":    content}
	contentBytes, err := renderTemplate(temp, bodyData)
	if err != nil {
		return contentBytes, err
	}
	return contentBytes, err
}

func renderTemplate(temp *template.Template, data map[string]interface{}) ([]byte, error) {
	var buf bytes.Buffer

	if err := temp.ExecuteTemplate(&buf, "page", data); err != nil {
		return buf.Bytes(), err
	}
	return buf.Bytes(), nil
}
