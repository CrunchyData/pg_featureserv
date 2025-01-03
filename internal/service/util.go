package service

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
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/CrunchyData/pg_featureserv/internal/api"
	"github.com/CrunchyData/pg_featureserv/internal/conf"
	"github.com/CrunchyData/pg_featureserv/internal/ui"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/theckman/httpforwarded"
)

const (
	schemeHTTP  = "http"
	schemeHTTPS = "https"
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

	// signal for normal completion of handler
	handlerDone := make(chan struct{})
	start := time.Now()

	// monitor context status and log anything abnormal
	go func() {
		select {
		case <-handlerDone:
			log.Debugf("---- Request complete in %v", time.Since(start))
		case <-r.Context().Done():
			// log cancelations
			switch r.Context().Err() {
			case context.DeadlineExceeded:
				log.Warnf("---- Request processing terminated by write timeout after %v", time.Since(start))
			case context.Canceled:
				log.Debugf("---- Request cancelled by client after %v", time.Since(start))
			}
		}
	}()

	// execute the handler
	e := fn(w, r)

	if e != nil { // e is *appError, not os.Error.
		// TODO: is this the desire behaviour?
		// perhaps detect format and emit accordingly?
		// log error here?
		// should log attached error?
		// panic on severe error?
		log.Debugf("Request processing error: %v (%v)\n", e.Message, e.Code)
		http.Error(w, e.Message, e.Code)
	}
	close(handlerDone)
}

// FatalAfter aborts by logging a fatal message, after a time delay.
// The abort can be cancelled by closing the returned channel
func FatalAfter(delaySec int, msg string) chan struct{} {
	chanCancel := make(chan struct{})
	go func() {
		select {
		case <-chanCancel:
			// do nothing if cancelled
			return
		case <-time.After(time.Duration(delaySec) * time.Second):
			// terminate with extreme predjudice
			log.Fatalln(msg)
		}
	}()
	return chanCancel
}

func appErrorMsg(err error, msg string, code int) *appError {
	return &appError{err, msg, code}
}

func appErrorInternal(err error, msg string) *appError {
	return &appError{err, msg, http.StatusInternalServerError}
}

func appErrorBadRequest(err error, msg string) *appError {
	return &appError{err, msg, http.StatusBadRequest}
}

func appErrorInternalFmt(err error, format string, v ...interface{}) *appError {
	msg := fmt.Sprintf(format, v...)
	return &appError{err, msg, http.StatusInternalServerError}
}

func appErrorNotFoundFmt(err error, format string, v string) *appError {
	msg := fmt.Sprintf(format, v)
	return &appError{err, msg, http.StatusNotFound}
}

//========================

func serveURLBase(r *http.Request) string {
	// Use configuration file settings if we have them
	configURL := conf.Configuration.Server.UrlBase

	if configURL != "" {
		return configURL + "/"
	}
	// Preferred scheme
	ps := schemeHTTP
	// check if HTTPS (TLS) is being used
	if r.TLS != nil {
		ps = schemeHTTPS
	}
	// Preferred host:port
	ph := strings.TrimRight(r.Host, "/")

	// Check IETF standard "Forwarded" header
	// for reverse proxy information
	xf := http.CanonicalHeaderKey("Forwarded")
	if f, ok := r.Header[xf]; ok {
		if fm, err := httpforwarded.Parse(f); err == nil {
			if len(fm["host"]) > 0 && len(fm["proto"]) > 0 {
				ph = fm["host"][0]
				ps = fm["proto"][0]
				return fmt.Sprintf("%v://%v/", ps, ph)
			}
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

	path := conf.Configuration.Server.BasePath
	return fmt.Sprintf("%v://%v%v/", ps, ph, path)
}

func getRequestVar(varname string, r *http.Request) string {
	vars := mux.Vars(r)
	nameFull := vars[varname]
	name := api.PathStripFormat(nameFull)
	return name
}

// urlPathFormat provides a URL for the given base and path
func urlPath(urlBase string, path string) string {
	url := fmt.Sprintf("%v%v", urlBase, path)
	return url
}

// urlPathFormat provides a URL for the given base, path and format
func urlPathFormat(urlBase string, path string, format string) string {
	var pathFormat string
	if path == "" {
		pathFormat = ""
		if format == api.FormatHTML {
			pathFormat = api.RootPageName + ".html"
		}
	} else {
		pathFormat = path + "." + format
	}
	url := fmt.Sprintf("%v%v", urlBase, pathFormat)
	return url
}

func urlPathFormatQuery(urlBase string, path string, format string, query string) string {
	url := urlPathFormat(urlBase, path, format)
	if query != "" {
		url = fmt.Sprintf("%v?%v", url, query)
	}
	return url
}

// formatBaseURL takes a hostname (baseHost) and a base path
// and joins them.  Both are parsed as URLs (using net/url) and
// then joined to ensure a properly formed URL.
// net/url does not support parsing hostnames without a scheme
// (e.g. example.com is invalid; http://example.com is valid).
// serverURLHost ensures a scheme is added.
func formatBaseURL(protocol string, baseHost string, basePath string) string {
	urlHost, err := url.Parse(protocol + baseHost)
	if err != nil {
		log.Fatal(err)
	}
	urlPath, err := url.Parse(basePath)
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimRight(urlHost.ResolveReference(urlPath).String(), "/")
}

// restrict creates a map containing only entries in names
func restrict(inMap map[string]string, names []string) map[string]string {
	outMap := make(map[string]string)
	for _, name := range names {
		//log.Debugf("testing request param %v", name)
		if val, ok := inMap[name]; ok {
			outMap[name] = val
			//log.Debugf("copying request param %v = %v ", name, val)
		}
	}
	return outMap
}

// removeNames removes a list of names from a map (map is modified)
//nolint:unused
func removeNames(inMap map[string]string, names []string) {
	for _, name := range names {
		delete(inMap, name)
	}
}

func writeJSON(w http.ResponseWriter, contype string, content interface{}) *appError {
	encodedContent, err := json.Marshal(content)
	if err != nil {
		log.Printf("JSON encoding error: %v", err.Error())
		return appErrorInternal(err, api.ErrMsgEncoding)
	}
	//fmt.Println(string(encodedContent))
	writeResponse(w, contype, encodedContent)
	return nil
}

func writeText(w http.ResponseWriter, contype string, encodedContent []byte) *appError {
	//fmt.Println(string(encodedContent))
	writeResponse(w, contype, encodedContent)
	return nil
}

func writeHTML(w http.ResponseWriter, content interface{}, context interface{}, templ *template.Template) *appError {
	encodedContent, err := ui.RenderHTML(templ, content, context)
	if err != nil {
		log.Printf("HTML encoding error: %v", err.Error())
		return appErrorInternal(err, api.ErrMsgEncoding)
	}
	return writeResponse(w, api.ContentTypeHTML, encodedContent)
}

func writeResponse(w http.ResponseWriter, contype string, encodedContent []byte) *appError {
	w.Header().Set("Content-Type", contype) //api.ContentType(format))
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(encodedContent)
	if err != nil {
		return appErrorInternal(err, api.ErrMsgDataWriteError)
	}
	return nil
}

// Sets response 'status', and writes a json-encoded object with property "description" having value "msg".
//nolint:all
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
