package db_test

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

 Date     : October 2022
 Authors  : Amaury Zarzelli (amaury dot zarzelli at ign dot fr)

*/

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/CrunchyData/pg_featureserv/internal/data"
	"github.com/CrunchyData/pg_featureserv/internal/util"
)

func (t *DbTests) TestCacheSizeIncreaseAfterCreate() {
	t.Test.Run("TestCacheSizeIncreaseAfterCreate", func(t *testing.T) {
		var header = make(http.Header)
		header.Add("Content-Type", "application/geo+json")

		//--- retrievecache size before insert
		var sizeBefore = len(cat.GetCache())

		//--- generate json from new object
		tables, _ := cat.Tables()
		var cols []string
		for _, tbl := range tables {
			if tbl.ID == "public.mock_a" {
				for _, c := range tbl.Columns {
					if c != "id" {
						cols = append(cols, c)
					}
				}
				break
			}
		}
		jsonStr := data.MakeFeatureMockPointAsJSON(99, 12, 34, cols)
		// -- do the request call but we have to force the catalogInstance to db during this operation
		_ = hTest.DoPostRequest(t, "/collections/mock_a/items", []byte(jsonStr), header)

		// Sleep in order to wait for the cache to update (parallel goroutine)
		time.Sleep(100 * time.Millisecond)

		//--- retrievecache size before insert
		var sizeAfter = len(cat.GetCache())

		util.Equals(t, sizeAfter, sizeBefore+1, "cache size augmented by 1 after one insert")
	})
}

func (t *DbTests) TestCacheSizeDecreaseAfterDelete() {
	t.Test.Run("TestCacheSizeDecreaseAfterDelete", func(t *testing.T) {
		var header = make(http.Header)
		header.Add("Content-Type", "application/geo+json")

		//--- generate json from new object
		tables, _ := cat.Tables()
		var cols []string
		for _, tbl := range tables {
			if tbl.ID == "public.mock_a" {
				for _, c := range tbl.Columns {
					if c != "id" {
						cols = append(cols, c)
					}
				}
				break
			}
		}
		jsonStr := data.MakeFeatureMockPointAsJSON(101, 12, 34, cols)
		// -- do the request call but we have to force the catalogInstance to db during this operation
		_ = hTest.DoPostRequest(t, "/collections/mock_a/items", []byte(jsonStr), header)
		rr := hTest.DoPostRequest(t, "/collections/mock_a/items", []byte(jsonStr), header)

		loc := rr.Header().Get("Location")
		var splittedLoc = strings.Split(loc, "/")
		var firstId, _ = strconv.Atoi(splittedLoc[len(splittedLoc)-1])

		time.Sleep(100 * time.Millisecond)

		//--- retrievecache size before insert
		var sizeBefore = len(cat.GetCache())

		fmt.Println(firstId)
		hTest.DoDeleteRequestStatus(t, fmt.Sprintf("/collections/mock_a/items/%v", firstId), http.StatusNoContent)

		// Sleep in order to wait for the cache to update (parallel goroutine)
		time.Sleep(100 * time.Millisecond)

		//--- retrievecache size before insert
		var sizeAfter = len(cat.GetCache())

		util.Equals(t, sizeAfter, sizeBefore-1, "cache size decreased by 1 after one delete")
	})
}

func (t *DbTests) TestCacheModifiedAfterUpdate() {
	t.Test.Run("TestCacheSizeDecreaseAfterDelete", func(t *testing.T) {
		var header = make(http.Header)
		header.Add("Content-Type", "application/geo+json")

		//--- generate json from new object
		jsonStr := `{
			"type": "Feature",
			"geometry": {
				"type": "Point",
				"coordinates": [
				-120,
				40
				]
			},
			"properties": {
				"prop_a": "POST",
				"prop_b": 1,
				"prop_c": "propC"
			}
		}`
		// -- do the request call but we have to force the catalogInstance to db during this operation
		rr := hTest.DoPostRequest(t, "/collections/mock_a/items", []byte(jsonStr), header)

		loc := rr.Header().Get("Location")
		var splittedLoc = strings.Split(loc, "/")
		var firstId, _ = strconv.Atoi(splittedLoc[len(splittedLoc)-1])

		jsonStr = fmt.Sprintf(`{
			"type": "Feature",
			"id": "%v",
			"geometry": {
				"type": "Point",
				"coordinates": [
				-120,
				40
				]
			},
			"properties": {
				"prop_a": "PUT",
				"prop_b": 1,
				"prop_c": "propC"
			}
		}`, firstId)

		path := fmt.Sprintf("/collections/mock_a/items/%v", firstId)
		hTest.DoRequestMethodStatus(t, "PUT", path, []byte(jsonStr), header, http.StatusNoContent)

		// Sleep in order to wait for the cache to update (parallel goroutine)
		time.Sleep(100 * time.Millisecond)

		// TODO: Test HERE if cache modified

		jsonStr = fmt.Sprintf(`{
			"type": "Feature",
			"id": "%v",
			"geometry": {
				"type": "Point",
				"coordinates": [
				-120,
				40
				]
			},
			"properties": {
				"prop_a": "PATCH",
				"prop_b": 1,
				"prop_c": "propC"
			}
		}`, firstId)

		hTest.DoRequestMethodStatus(t, "PATCH", path, []byte(jsonStr), header, http.StatusNoContent)

		// Sleep in order to wait for the cache to update (parallel goroutine)
		time.Sleep(100 * time.Millisecond)

		// TODO: Test HERE if cache modified
	})
}
