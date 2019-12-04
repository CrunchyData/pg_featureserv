package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/dr-jts/pg_featureserv/api"
	"github.com/dr-jts/pg_featureserv/ui"
)

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
		fmt.Printf(err.Error())
		//fmt.Printf(templateStr)
	}
	return encodedContent, err
}

func encodeHTML(content interface{}, context interface{}, templ *template.Template) ([]byte, error) {
	encodedContent, err := ui.RenderHTML(templ, content, context)
	if err != nil {
		// TODO: encode error in format
		//jsonError(w, "EncodeError", err.Error(), http.StatusInternalServerError)
		fmt.Printf(err.Error())
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
	fmt.Printf("%v Request: %v\n", r.RemoteAddr, r.URL)
}
