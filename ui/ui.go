package ui

import (
	"bytes"
	"html/template"

	"github.com/dr-jts/pg_featureserv/config"
)

// PageContext - data used on the HTML pages
type PageContext struct {
	// URLHome - URL for the service home page
	UrlHome         string
	UrlCollections  string
	UrlCollection   string
	UrlItems        string
	UrlJSON         string
	CollectionTitle string
	FeatureID       string
	UseMap          bool
}

// HTMLTemplate UI templates repo (inited on startup)
var HTMLTemplate struct {
	Page        *template.Template
	Home        *template.Template
	Conformance *template.Template
	Collections *template.Template
	Collection  *template.Template
	Items       *template.Template
	Item        *template.Template
}

func init() {
	HTMLTemplate.Page = createTemplate("page", templatePage)
	HTMLTemplate.Home = createTemplate("home", templateHome)
	HTMLTemplate.Conformance = createTemplate("conformance", templateConformance)
	HTMLTemplate.Collections = createTemplate("collections", templateCollections)
	HTMLTemplate.Collection = createTemplate("collection", templateCollection)
	HTMLTemplate.Items = createTemplate("items", templateItems+mapCode)
	HTMLTemplate.Item = createTemplate("item", templateItem+mapCode)
}

func createTemplate(name string, templateStr string) *template.Template {
	return template.Must(template.New(name).Parse(templateStr))
}

// RenderHTML tbd
func RenderHTML(temp *template.Template, content interface{}, context interface{}) ([]byte, error) {
	bodyData := map[string]interface{}{
		"config":  config.Configuration,
		"context": context,
		"data":    content}
	contentBytes, err := renderTemplate(temp, bodyData)
	if err != nil {
		return contentBytes, err
	}

	data := map[string]interface{}{
		"config":  config.Configuration,
		"context": context,
		"body":    template.HTML(contentBytes)}
	return renderTemplate(HTMLTemplate.Page, data)
}

func renderTemplate(temp *template.Template, data map[string]interface{}) ([]byte, error) {
	var tpl bytes.Buffer

	if err := temp.Execute(&tpl, data); err != nil {
		return tpl.Bytes(), err
	}
	return tpl.Bytes(), nil
}
