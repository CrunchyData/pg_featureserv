package utiltest

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

 Date     : September 2022
 Authors  : Benoit De Mezzo (benoit dot de dot mezzo at oslandia dot com)
*/

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/CrunchyData/pg_featureserv/internal/conf"
	"github.com/gorilla/mux"
)

type HttpTesting struct {
	UrlBase    string
	BasePath   string
	AssetsPath string
	Router     *mux.Router
}

func MakeHttpTesting(urlBase string, basePath string, assetsPath string, router *mux.Router) HttpTesting {
	hTest := HttpTesting{
		UrlBase:    urlBase,
		BasePath:   basePath,
		AssetsPath: assetsPath,
		Router:     router,
	}
	hTest.Setup()
	return hTest
}

func (hTest *HttpTesting) Setup() {
	conf.Configuration = conf.Config{
		Server: conf.Server{
			HttpHost:   "0.0.0.0",
			HttpPort:   9000,
			UrlBase:    hTest.UrlBase,
			BasePath:   hTest.BasePath,
			AssetsPath: hTest.AssetsPath,
			TransformFunctions: []string{
				"ST_Centroid",
				"ST_PointOnSurface",
			},
		},
		Paging: conf.Paging{
			LimitDefault: 10,
			LimitMax:     1000,
		},
		Metadata: conf.Metadata{
			Title:       "test",
			Description: "test",
		},
	}
}

// returns the all the body from an http response
func (hTest *HttpTesting) ReadBody(resp *httptest.ResponseRecorder) []byte {
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}

// do an http request to url with default method GET and expected status OK
func (hTest *HttpTesting) DoRequest(t *testing.T, url string) *httptest.ResponseRecorder {
	return hTest.DoRequestStatus(t, url, http.StatusOK)
}

func (hTest *HttpTesting) DoPostRequest(t *testing.T, url string, data []byte, header http.Header) *httptest.ResponseRecorder {
	return hTest.DoRequestMethodStatus(t, "POST", url, data, header, http.StatusCreated)
}

// do an http request to url with default method GET and a specific expected status
func (hTest *HttpTesting) DoRequestStatus(t *testing.T, url string,
	statusExpected int) *httptest.ResponseRecorder {
	return hTest.DoRequestMethodStatus(t, "GET", url, nil, nil, statusExpected)
}

// do an http request to url with DELETE method and an expected status
func (hTest *HttpTesting) DoDeleteRequestStatus(t *testing.T, url string, statusExpected int) *httptest.ResponseRecorder {
	return hTest.DoRequestMethodStatus(t, "DELETE", url, nil, nil, statusExpected)
}

// do an http request to url with a specific method and specific expected status
func (hTest *HttpTesting) DoRequestMethodStatus(t *testing.T, method string, url string,
	data []byte, header http.Header, statusExpected int) *httptest.ResponseRecorder {

	req, err := http.NewRequest(method, hTest.BasePath+url, bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}

	for k, v := range header {
		req.Header.Add(k, v[0])
	}

	rr := httptest.NewRecorder()
	hTest.Router.ServeHTTP(rr, req)

	// Check the status code
	status := rr.Code
	if status != statusExpected {
		var bodyMsg string
		body, errBody := ioutil.ReadAll(rr.Body)
		if errBody == nil && len(body) != 0 {
			bodyMsg = "Error: " + string(body)
		}

		if bodyMsg != "" {
			Equals(t,
				statusExpected, status,
				fmt.Errorf("handler returned wrong status code.\n\tCaused by: %v", bodyMsg).Error())
		} else {
			Equals(t,
				statusExpected, status,
				"handler returned wrong status code.")
		}
	}
	return rr
}
