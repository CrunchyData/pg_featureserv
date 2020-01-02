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
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/CrunchyData/pg_featureserv/api"
	"github.com/CrunchyData/pg_featureserv/config"
	"github.com/CrunchyData/pg_featureserv/ui"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

//--- simple request handler error framework
// see https://blog.golang.org/error-handling-and-go

type appError struct {
	Error   error
	Message string
	Code    int
}

type appHandler func(http.ResponseWriter, *http.Request) *appError

func appErrorMsg(err error, msg string, code int) *appError {
	return &appError{err, msg, code}
}

func appErrorInternal(err error, msg string) *appError {
	return &appError{err, msg, http.StatusInternalServerError}
}

func appErrorNoFound(err error, msg string) *appError {
	return &appError{err, msg, http.StatusNotFound}
}

func appErrorNotFoundFmt(err error, format string, v string) *appError {
	msg := fmt.Sprintf(format, v)
	return &appError{err, msg, http.StatusNotFound}
}

//========================

func serveURLBase(r *http.Request) string {
	// Preferred host:port
	php := r.Host
	php = strings.TrimRight(php, "/")

	// Preferred scheme
	ps := "http"

	// Preferred base path
	pbp := "/"

	// Preferred scheme / host / port / base
	pshpb := fmt.Sprintf("%v://%v%v", ps, php, pbp)
	return pshpb
}

func getRequestVar(varname string, r *http.Request) string {
	vars := mux.Vars(r)
	nameFull := vars[varname]
	name := api.PathStripFormat(nameFull)
	return name
}

// urlPathFormat provides a URL for the given base, path and format
func urlPathFormat(urlBase string, path string, format string) string {
	var pathType string
	if path == "" {
		pathType = ""
		if format == api.FormatHTML {
			pathType = "home.html"
		}
	} else {
		pathType = path + "." + format
	}
	url := fmt.Sprintf("%v%v", urlBase, pathType)
	/*
		if !supportedContentType(contentType) {
			panic(fmt.Sprintf("unsupported content type: %v", contentType))
		}
	*/
	return url
}

func urlPathFormatQuery(urlBase string, path string, format string, query string) string {
	url := urlPathFormat(urlBase, path, format)
	if query != "" {
		url = fmt.Sprintf("%v?%v", url, query)
	}
	return url
}

func writeJSON(w http.ResponseWriter, contype string, content interface{}) *appError {
	encodedContent, err := json.Marshal(content)
	if err != nil {
		log.Printf("JSON encoding error: %v", err.Error())
		return appErrorInternal(err, errMsgEncoding)
	}
	writeResponse(w, contype, encodedContent)
	return nil
}

func writeHTML(w http.ResponseWriter, content interface{}, context interface{}, templ *template.Template) *appError {
	encodedContent, err := ui.RenderHTML(templ, content, context)
	if err != nil {
		log.Printf("HTML encoding error: %v", err.Error())
		return appErrorInternal(err, errMsgEncoding)
	}
	writeResponse(w, api.ContentTypeHTML, encodedContent)
	return nil
}

func writeResponse(w http.ResponseWriter, contype string, encodedContent []byte) {
	w.Header().Set("Content-Type", contype) //api.ContentType(format))
	w.WriteHeader(http.StatusOK)
	w.Write(encodedContent)
}

// Sets response 'status', and writes a json-encoded object with property "description" having value "msg".
func writeError(w http.ResponseWriter, code string, msg string, status int) {
	w.WriteHeader(status)

	result, err := json.Marshal(struct {
		Code        string `json:"code"`
		Description string `json:"description"`
	}{
		Code:        code,
		Description: msg,
	})

	if err != nil {
		w.Write([]byte(fmt.Sprintf("problem marshaling error: %v", msg)))
	} else {
		w.Write(result)
	}
}

// NewPageData create a page context initialized with globals.
func NewPageData() *ui.PageData {
	con := ui.PageData{}
	con.AppName = config.AppConfig.Name
	con.AppVersion = config.AppConfig.Version
	return &con
}
