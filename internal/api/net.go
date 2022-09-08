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
	"net/http"
	"net/url"
	"strings"
)

const (
	// ContentTypeJSON
	ContentTypeJSON = "application/json"

	// ContentTypeGeoJSON
	ContentTypeGeoJSON = "application/geo+json"

	// ContentTypeGML
	ContentTypeGML = "application/gml+xml"

	// ContentTypeSchemaJSON
	ContentTypeSchemaJSON = "application/schema+json"

	// ContentTypeSchemaPatchJSON
	ContentTypeSchemaPatchJSON = "application/merge-patch+json"

	// ContentTypeHTML
	ContentTypeHTML = "text/html"

	// ContentTypeText
	ContentTypeText = "text/plain"

	// ContentTypeSVG
	ContentTypeSVG = "image/svg+xml"

	// ContentTypeHTML
	ContentTypeOpenAPI = "application/vnd.oai.openapi+json;version=3.0"

	// FormatJSON code and extension for JSON
	FormatJSON = "json"

	// FormatHTML code and extension for HTML
	FormatHTML = "html"

	// FormatText code and extension for Text
	FormatText = "text"

	// FormatText code and extension for Text
	FormatSVG = "svg"

	// FormatJSON code and extension for JSON
	FormatSchemaJSON = "schema+json"

	// FormatXML code and extension for XML/GML
	FormatXML = "xml"
)

// RequestedFormat gets the format for a request from extension or headers
func RequestedFormat(r *http.Request) string {
	// first check explicit path
	path := r.URL.EscapedPath()
	if strings.HasSuffix(path, ".html") {
		return FormatHTML
	}
	if strings.HasSuffix(path, ".json") {
		return FormatJSON
	}
	if strings.HasSuffix(path, ".txt") {
		return FormatText
	}
	if strings.HasSuffix(path, ".svg") {
		return FormatSVG
	}
	// Use Accept header if present
	hdrAccept := r.Header.Get("Accept")
	//fmt.Println("Accept:" + hdrAccept)
	if strings.Contains(hdrAccept, ContentTypeSchemaJSON) {
		return FormatSchemaJSON
	}
	if strings.Contains(hdrAccept, ContentTypeHTML) {
		return FormatHTML
	}
	return FormatJSON
}

// RequestedFormat gets the format for a request from extension or headers
func SentDataFormat(r *http.Request) string {
	// Use ContentType header if present
	hdrContentType := r.Header.Get("Content-Type")
	if strings.Contains(hdrContentType, ContentTypeGeoJSON) {
		return FormatJSON
	}
	if strings.Contains(hdrContentType, ContentTypeGML) {
		return FormatXML
	}
	return FormatJSON
}

// PathStripFormat removes a format extension from a path
func PathStripFormat(path string) string {
	if strings.HasSuffix(path, ".html") || strings.HasSuffix(path, ".json") {
		return path[0 : len(path)-5]
	}
	return path
}

// URLQuery gets the query part of a URL
func URLQuery(url *url.URL) string {
	uri := url.RequestURI()
	qloc := strings.Index(uri, "?")
	if qloc < 0 {
		return ""
	}
	query := uri[qloc+1:]
	return query
}
