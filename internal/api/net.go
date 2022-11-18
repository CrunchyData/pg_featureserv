package api

/*
 Copyright 2022 Crunchy Data Solutions, Inc.
 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at
      http://www.apache.org/licenses/LICENSE-2.0
 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.

 Date     : October 2022
 Authors  : Nicolas Revelant (nicolas dot revelant at ign dot fr)
*/

import (
	"net/http"
	"net/url"
	"strings"
	"time"
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

	// Extension value
	suffix := PathSuffix(path)
	if suffix != "" {
		switch suffix {
		case "json":
			return FormatJSON
		case "html":
			return FormatHTML
		case "txt":
			return FormatText
		case "svg":
			return FormatSVG
		default:
			return suffix
		}
	} else {
		// Accept header value
		hdrAcceptValue := r.Header.Get("Accept")
		if hdrAcceptValue != "" {
			// Accept header fields preferences:
			// -> https://www.rfc-editor.org/rfc/rfc9110.html#section-12.5.1

			// Examples:
			// "Accept: application/json"
			// "Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8"
			preferredFormats := strings.Split(hdrAcceptValue, ",")
			for _, value := range preferredFormats {
				mediaTypeValue := value
				lastSemicolon := strings.LastIndex(value, ";")
				if lastSemicolon > 0 {
					mediaTypeValue = value[:lastSemicolon] // 'q' quality parameter not used
				}
				switch mediaTypeValue {
				case ContentTypeJSON:
					return FormatJSON
				case ContentTypeSchemaJSON, ContentTypeSchemaPatchJSON:
					return FormatSchemaJSON
				case ContentTypeHTML:
					return FormatHTML
				case ContentTypeText:
					return FormatText
				case ContentTypeSVG:
					return FormatSVG
				}
			}
			return hdrAcceptValue
		}
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
	pos := strings.LastIndex(path, ".")
	if pos != -1 {
		path = path[:pos]
	}
	return path
}

// PathSuffix returns the format extension from a path following a dot character
func PathSuffix(path string) string {
	splittedPath := strings.Split(path, "/")
	pathEnd := splittedPath[len(splittedPath)-1]
	pos := strings.LastIndex(pathEnd, ".")
	if pos != -1 {
		return pathEnd[pos+1:]
	} else {
		return ""
	}
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

// Returns the current Http Date as a string
func GetCurrentHttpDate() string {

	// The date format here conforms to HTTP Date, always in GMT time zone like "Wed, 21 Oct 2015 07:28:00 GMT"
	// -> https://www.rfc-editor.org/rfc/rfc7231#section-7.1.1.1
	location, _ := time.LoadLocation("GMT")
	currentTime := time.Now().In(location)

	return currentTime.Format(time.RFC1123)
}
