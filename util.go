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
	log "github.com/sirupsen/logrus"
)

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

func encodeJSON(content interface{}) ([]byte, error) {
	encodedContent, err := json.Marshal(content)
	if err != nil {
		// TODO: encode error in format
		//jsonError(w, "EncodeError", err.Error(), http.StatusInternalServerError)
		log.Printf("JSON encoding error: %v", err.Error())
	}
	return encodedContent, err
}

func encodeHTML(content interface{}, context interface{}, templ *template.Template) ([]byte, error) {
	encodedContent, err := ui.RenderHTML(templ, content, context)
	if err != nil {
		// TODO: encode error in format
		//jsonError(w, "EncodeError", err.Error(), http.StatusInternalServerError)
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

	// testing only
	//param.TransformFunc = "ST_PointOnSurface"
	param.TransformFun, param.TransformArg = parseTransform(queryValues, 0)
	param.TransformFun2, param.TransformArg2 = parseTransform(queryValues, 1)

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

// parseTransform parses a transform function and optional argument:  transform=ST_Fun,arg
func parseTransform(values url.Values, index int) (string, string) {
	vals, ok := values[api.ParamTransform]
	if !ok || len(vals) <= index {
		return "", ""
	}
	val := vals[index]
	if len(val) < 1 {
		return "", ""
	}
	funName := val
	arg := ""
	// check for function parameter
	sepIndex := strings.Index(val, transformParamSep)
	if sepIndex >= 0 {
		funName = val[:sepIndex]
		arg = val[sepIndex+1:]
	}
	// TODO: harden this by checking arg is a valid number
	// TODO: have whitelist for function names?
	return funName, arg
}
