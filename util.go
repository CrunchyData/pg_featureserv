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
	"github.com/CrunchyData/pg_featureserv/conf"
	"github.com/CrunchyData/pg_featureserv/data"
	"github.com/CrunchyData/pg_featureserv/ui"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/theckman/httpforwarded"
)

//--- simple request handler error framework
// see https://blog.golang.org/error-handling-and-go
// Allows handlers to return structure error messages,
// and centralizes common logic in the ServeHTTP method

type appError struct {
	Error   error
	Message string
	Code    int
}

// appHandler is a named function type which is augmented
// with a ServeHTTP method
// This allows it to be used as an http.Handler.
// When the ServeHTTP method is called
// the bound function is called, and
// can return a structured error message.
type appHandler func(http.ResponseWriter, *http.Request) *appError

// ServeHTTP is the base Handler for routed requests.
// Common handling logic is placed here
// See also https://golang.org/pkg/net/http/#Handler
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

func appErrorMsg(err error, msg string, code int) *appError {
	return &appError{err, msg, code}
}

func appErrorInternal(err error, msg string) *appError {
	return &appError{err, msg, http.StatusInternalServerError}
}

func appErrorInternalFmt(err error, format string, v ...interface{}) *appError {
	msg := fmt.Sprintf(format, v...)
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
	// Use configuration file settings if we have them
	configUrl := conf.Configuration.Server.UrlBase

	if configUrl != "" {
		return configUrl + "/"
	}
	// Preferred protocol
	ps := "http"
	// Preferred host:port
	ph := strings.TrimRight(r.Host, "/")

	// Check IETF standard "Forwarded" header
	// for reverse proxy information
	xf := http.CanonicalHeaderKey("Forwarded")
	if f, ok := r.Header[xf]; ok {
		if fm, err := httpforwarded.Parse(f); err == nil {
			ph = fm["host"][0]
			ps = fm["proto"][0]
			return fmt.Sprintf("%v://%v/", ps, ph)
		}
	}

	// Check X-Forwarded-Host and X-Forwarded-Proto headers
	xfh := http.CanonicalHeaderKey("X-Forwarded-Host")
	if fh, ok := r.Header[xfh]; ok {
		ph = fh[0]
	}
	xfp := http.CanonicalHeaderKey("X-Forwarded-Proto")
	if fp, ok := r.Header[xfp]; ok {
		ps = fp[0]
	}

	return fmt.Sprintf("%v://%v/", ps, ph)
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
			pathType = api.RootPageName + ".html"
		}
	} else {
		pathType = path + "." + format
	}
	url := fmt.Sprintf("%v%v", urlBase, pathType)
	return url
}

func urlPathFormatQuery(urlBase string, path string, format string, query string) string {
	url := urlPathFormat(urlBase, path, format)
	if query != "" {
		url = fmt.Sprintf("%v?%v", url, query)
	}
	return url
}

func createQueryParams(requestParam *api.RequestParam, colNames []string) *data.QueryParam {
	param := data.QueryParam{
		Limit:         requestParam.Limit,
		Offset:        requestParam.Offset,
		Bbox:          requestParam.Bbox,
		OrderBy:       requestParam.OrderBy,
		Precision:     requestParam.Precision,
		TransformFuns: requestParam.TransformFuns,
	}
	param.Columns = normalizePropNames(requestParam.Properties, colNames)
	return &param
}

// inputArgs extracts function arguments from any provided in the query parameters
func restrict(inMap map[string]string, names []string) map[string]string {
	outMap := make(map[string]string)
	for _, name := range names {
		log.Debugf("testing request param %v", name)
		if val, ok := inMap[name]; ok {
			outMap[name] = val
			log.Debugf("copying request param %v = %v ", name, val)
		}
	}
	return outMap
}

func writeJSON(w http.ResponseWriter, contype string, content interface{}) *appError {
	encodedContent, err := json.Marshal(content)
	if err != nil {
		log.Printf("JSON encoding error: %v", err.Error())
		return appErrorInternal(err, api.ErrMsgEncoding)
	}
	writeResponse(w, contype, encodedContent)
	return nil
}

func writeHTML(w http.ResponseWriter, content interface{}, context interface{}, templ *template.Template) *appError {
	encodedContent, err := ui.RenderHTML(templ, content, context)
	if err != nil {
		log.Printf("HTML encoding error: %v", err.Error())
		return appErrorInternal(err, api.ErrMsgEncoding)
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
