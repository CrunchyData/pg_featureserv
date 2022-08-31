package service

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
*/

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

// returns the all the body from an http response
func readBody(resp *httptest.ResponseRecorder) []byte {
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}

// do an http request to url with default method GET and expected status OK
func doRequest(t *testing.T, url string) *httptest.ResponseRecorder {
	return doRequestStatus(t, url, http.StatusOK)
}

func doPostRequest(t *testing.T, url string, data []byte, header http.Header) *httptest.ResponseRecorder {
	return doRequestMethodStatus(t, "POST", url, data, header, http.StatusCreated)
}

// do an http request to url with default method GET and a specific expected status
func doRequestStatus(t *testing.T, url string,
	statusExpected int) *httptest.ResponseRecorder {
	return doRequestMethodStatus(t, "GET", url, nil, nil, statusExpected)
}

// do an http request to url with a specific method and specific expected status
func doRequestMethodStatus(t *testing.T, method string, url string,
	data []byte, header http.Header, statusExpected int) *httptest.ResponseRecorder {

	req, err := http.NewRequest(method, basePath+url, bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}

	for k, v := range header {
		req.Header.Add(k, v[0])
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// Check the status code
	status := rr.Code
	if status != statusExpected {
		var bodyMsg string
		body, errBody := ioutil.ReadAll(rr.Body)
		if errBody == nil && len(body) != 0 {
			bodyMsg = "Error: " + string(body)
		}

		if bodyMsg != "" {
			equals(t,
				statusExpected, status,
				fmt.Errorf("handler returned wrong status code.\n\tCaused by: %v", bodyMsg).Error())
		} else {
			equals(t,
				statusExpected, status,
				"handler returned wrong status code.")
		}
	}
	return rr
}

//---- testing utilities from https://github.com/benbjohnson/testing

// assert fails the test if the condition is false.
func assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: "+msg+"\033[39m\n\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		tb.FailNow()
	}
}

// equals fails the test if exp is not equal to act.
func equals(tb testing.TB, exp, act interface{}, msg string) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d: %s - expected: %#v; got: %#v\n", filepath.Base(file), line, msg, exp, act)
		tb.FailNow()
	}
}
