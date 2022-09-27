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
 Authors  : Nicolas Revelant (nicolas dot revelant at ign dot fr)
*/

import (
	"context"
	"net/http"
	"testing"

	"github.com/CrunchyData/pg_featureserv/internal/api"
	"github.com/CrunchyData/pg_featureserv/internal/data"
	util "github.com/CrunchyData/pg_featureserv/internal/utiltest"
)

func (t *DbTests) TestDeleteFeatureDb() {
	t.Test.Run("TestDeleteFeatureDb", func(t *testing.T) {
		//--- retrieve max feature id before delete
		var features []*api.GeojsonFeatureData
		params := data.QueryParam{Limit: 100000, Offset: 0, Crs: 4326}
		features, _ = cat.TableFeatures(context.Background(), "mock_b", &params)

		featuresNumbersBefore := len(features)

		// -- do the request call but we have to force the catalogInstance to db during this operation
		hTest.DoDeleteRequestStatus(t, "/collections/mock_b/items/1", http.StatusNoContent)

		//--- retrieve max feature id after delete
		features, _ = cat.TableFeatures(context.Background(), "mock_b", &params)
		featuresNumbersAfter := len(features)

		util.Assert(t, featuresNumbersBefore-1 == featuresNumbersAfter, "# feature still in db/not deleted")
	})
}
