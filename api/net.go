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
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const (
	// ContentTypeJSON tbd
	ContentTypeJSON = "application/json"

	// ContentTypeGeoJSON tbd
	ContentTypeGeoJSON = "application/geo+json"

	// ContentTypeHTML tbd
	ContentTypeHTML = "text/html"

	FormatJSON = "json"

	FormatHTML = "html"
)

// ContentType tbd
func ContentType(format string) string {
	switch format {
	case FormatJSON:
		return ContentTypeJSON
	case FormatHTML:
		return ContentTypeHTML
	}
	return ""
}

func PathFormat(url *url.URL) string {
	path := url.EscapedPath()
	if strings.HasSuffix(path, ".html") {
		return FormatHTML
	}
	return FormatJSON
}

func RequestedFormat(r *http.Request) string {
	// first check explicit path
	path := r.URL.EscapedPath()
	if strings.HasSuffix(path, ".html") {
		return FormatHTML
	}
	if strings.HasSuffix(path, ".json") {
		return FormatJSON
	}
	// Use Accept header if present
	hdrAccept := r.Header.Get("Accept")
	fmt.Println(hdrAccept)
	if strings.Index(hdrAccept, ContentTypeHTML) >= 0 {
		return FormatHTML
	}
	return FormatJSON
}

func PathStripFormat(path string) string {
	if strings.HasSuffix(path, ".html") || strings.HasSuffix(path, ".json") {
		return path[0 : len(path)-5]
	}
	return path
}

func URLQuery(url *url.URL) string {
	uri := url.RequestURI()
	qloc := strings.Index(uri, "?")
	if qloc < 0 {
		return ""
	}
	query := uri[qloc+1:]
	return query
}
