package util

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
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"github.com/CrunchyData/pg_featureserv/internal/api"
	"github.com/paulmach/orb/geojson"
)

// extracted from catalog_db.go
// TODO should be imported from catalog.go
type GeojsonFeatureData struct {
	Type  string                 `json:"type"`
	ID    string                 `json:"id,omitempty"`
	Geom  *geojson.Geometry      `json:"geometry"`
	Props map[string]interface{} `json:"properties"`
}

// Define a FeatureCollection structure for parsing test data
// TODO should be move to and imported from catalog.go
type FeatureCollection struct {
	Type           string                `json:"type"`
	Features       []*GeojsonFeatureData `json:"features"`
	NumberMatched  uint                  `json:"numberMatched,omitempty"`
	NumberReturned uint                  `json:"numberReturned"`
	TimeStamp      string                `json:"timeStamp,omitempty"`
	Links          []*api.Link           `json:"links"`
}

//---- testing utilities from https://github.com/benbjohnson/testing

// assert fails the test if the condition is false.
func Assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: "+msg+"\033[39m\n\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		tb.FailNow()
	}
}

// equals fails the test if exp is not equal to act.
func Equals(tb testing.TB, exp, act interface{}, msg string) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d: %s - expected: %#v; got: %#v\n", filepath.Base(file), line, msg, exp, act)
		tb.FailNow()
	}
}

// isError fails the test if err is nil, ie. test should have failed!
func AssertIsError(tb testing.TB, err error, msg string) {
	if err == nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d: %s - expected error\n", filepath.Base(file), line, msg)
		tb.FailNow()
	}
}
