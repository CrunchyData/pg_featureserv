package mock_test

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
	"strings"
	"testing"

	"github.com/CrunchyData/pg_featureserv/internal/service"
	util "github.com/CrunchyData/pg_featureserv/internal/utiltest"
)

func TestRootEmptyBasePath(t *testing.T) {

	hTestBadPath := util.MakeHttpTesting("http://test", "", "../../../assets", service.InitRouter(""))
	service.Initialize()

	testCases := []string{
		"/",
		"/index.html",
	}
	for _, tc := range testCases {
		title := fmt.Sprintf("route works with empty base path (%v)", strings.TrimPrefix(tc, "/"))
		t.Run(title, func(t *testing.T) {
			resp := hTestBadPath.DoRequest(t, tc)
			util.Assert(t, resp.Code == 200, "Status must be 200")
		})
	}

}
