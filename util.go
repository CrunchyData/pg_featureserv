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
	"net/url"
	"strconv"
	"strings"

	"github.com/CrunchyData/pg_featureserv/api"
	"github.com/CrunchyData/pg_featureserv/config"
	"github.com/CrunchyData/pg_featureserv/data"
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

func AppError(err error, msg string, code int) *appError {
	return &appError{err, msg, code}
}

func AppErrorInternal(err error, msg string) *appError {
	return &appError{err, msg, http.StatusInternalServerError}
}

func AppErrorNoFound(err error, msg string) *appError {
	return &appError{err, msg, http.StatusNotFound}
}

func AppErrorNotFoundFmt(err error, format string, v string) *appError {
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

// Provides a link for the given content type
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

func writeJSON(w http.ResponseWriter, contype string, content interface{}) *appError {
	encodedContent, err := json.Marshal(content)
	if err != nil {
		log.Printf("JSON encoding error: %v", err.Error())
		return AppErrorInternal(err, errMsgEncoding)
	}
	writeResponse(w, contype, encodedContent)
	return nil
}

func encodeJSON(content interface{}) ([]byte, error) {
	encodedContent, err := json.Marshal(content)
	if err != nil {
		log.Printf("JSON encoding error: %v", err.Error())
	}
	return encodedContent, err
}

func writeHTML(w http.ResponseWriter, content interface{}, context interface{}, templ *template.Template) *appError {
	encodedContent, err := ui.RenderHTML(templ, content, context)
	if err != nil {
		log.Printf("HTML encoding error: %v", err.Error())
		return AppErrorInternal(err, errMsgEncoding)
	}
	writeResponse(w, api.ContentTypeHTML, encodedContent)
	return nil
}

func encodeHTML(content interface{}, context interface{}, templ *template.Template) ([]byte, error) {
	encodedContent, err := ui.RenderHTML(templ, content, context)
	if err != nil {
		log.Printf("HTML encoding error: %v", err.Error())
		//fmt.Printf(templateStr)
	}
	return encodedContent, err
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

func logRequest(r *http.Request) {
	log.Printf("%v Request: %v\n", r.RemoteAddr, r.URL)
}

// NewPageData create a page context initialized with globals.
func NewPageData() *ui.PageData {
	con := ui.PageData{}
	con.AppName = config.AppConfig.Name
	con.AppVersion = config.AppConfig.Version
	return &con
}

func parseRequestParams(r *http.Request) data.QueryParam {
	param := data.QueryParam{
		Limit: config.Configuration.Server.DefaultLimit,
	}

	queryValues := r.URL.Query()

	//-- parse limit
	param.Limit = parseLimit(queryValues)
	param.TransformFuns = parseTransform(queryValues, 0)

	return param
}

func parseLimit(values url.Values) int {
	val := values.Get(api.ParamLimit)
	if len(val) < 1 {
		return config.Configuration.Server.DefaultLimit
	}
	limit, err := strconv.Atoi(val)
	if err != nil {
		return config.Configuration.Server.DefaultLimit
	}
	if limit < 0 || limit > config.Configuration.Server.MaxLimit {
		return config.Configuration.Server.MaxLimit
	}
	return limit
}

const transformParamSep = ","
const transformFunSep = "|"

func parseTransform(values url.Values, index int) []data.TransformFunction {
	val := values.Get(api.ParamTransform)
	if len(val) < 1 {
		return nil
	}
	funDefs := strings.Split(val, transformFunSep)

	funList := make([]data.TransformFunction, 0)
	for _, fun := range funDefs {
		name, arg := parseTransformFun(fun)
		if name != "" {
			funList = append(funList, data.TransformFunction{Name: name, Arg: arg})
		}
	}
	return funList
}

func parseTransformFun(def string) (string, string) {
	funName := def
	arg := ""
	// check for function parameter
	sepIndex := strings.Index(def, transformParamSep)
	if sepIndex >= 0 {
		funName = def[:sepIndex]
		arg = def[sepIndex+1:]
	}
	// TODO: harden this by checking arg is a valid number
	// TODO: have whitelist for function names?
	return funName, arg
}
