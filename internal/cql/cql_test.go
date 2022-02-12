package cql

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
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"
)

func TestPredicate(t *testing.T) {
	checkCQL(t, "", "")
	checkCQL(t, "id > 1", "\"id\" > 1")
	checkCQL(t, "id >= 1", "\"id\" >= 1")
	checkCQL(t, "id < 1", "\"id\" < 1")
	checkCQL(t, "id <= 1", "\"id\" <= 1")
	checkCQL(t, "id = 1", "\"id\" = 1")
	checkCQL(t, "id <> 1", "\"id\" <> 1")

	checkCQL(t, "id = -1.2345", "\"id\" = -1.2345")
	checkCQL(t, "id = id2", "\"id\" = \"id2\"")
	checkCQL(t, "id = 'foo'", "\"id\" = 'foo'")

	checkCQL(t, "id LIKE 'foo'", "\"id\" LIKE 'foo'")
	checkCQL(t, "id ILIKE 'foo'", "\"id\" ILIKE 'foo'")
	checkCQL(t, "id ILIKE '%Ca%'", "\"id\" ILIKE '%Ca%'")

	checkCQL(t, "id BETWEEN 1 and 2", "\"id\" BETWEEN 1 AND 2")
	checkCQL(t, "id NOT BETWEEN 1 and 2", "\"id\" NOT BETWEEN 1 AND 2")

	checkCQL(t, "id IN (1,2,3)", "\"id\" IN (1,2,3)")
	checkCQL(t, "id NOT IN (1,2,3)", "\"id\" NOT IN (1,2,3)")
	checkCQL(t, "id IN ('a','b','c')", "\"id\" IN ('a','b','c')")

	checkCQL(t, "id IS NULL", "\"id\" IS NULL")
	checkCQL(t, "id IS NOT NULL", "\"id\" IS NOT NULL")
}
func TestBooleanExpression(t *testing.T) {
	checkCQL(t, "x > 1 AND x < 9", "\"x\" > 1 AND \"x\" < 9")
	checkCQL(t, "x = 1 OR x = 2", "\"x\" = 1 OR \"x\" = 2")
	checkCQL(t, "(x = 1 OR x = 2) AND y < 4", "(\"x\" = 1 OR \"x\" = 2) AND \"y\" < 4")
	checkCQL(t, "NOT x IS NOT NULL", "NOT  \"x\" IS NOT NULL")
}

func checkCQL(t *testing.T, cqlStr string, sql string) {
	actual := TranspileToSQL(cqlStr)
	actual = strings.TrimSpace(actual)
	equals(t, sql, actual, "")
}

// equals fails the test if exp is not equal to act.
func equals(tb testing.TB, exp, act interface{}, msg string) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d: %s - expected: %#v; got: %#v\n", filepath.Base(file), line, msg, exp, act)
		tb.FailNow()
	}
}
