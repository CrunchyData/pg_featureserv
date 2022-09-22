package data

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
	"fmt"
	"strconv"

	"github.com/CrunchyData/pg_featureserv/internal/api"
	orb "github.com/paulmach/orb"
	"github.com/paulmach/orb/geojson"
)

// mock object used in catalog_mock
type featureMock struct {
	api.GeojsonFeatureData
}

func makeFeatureMockPoint(id int, x float64, y float64) *featureMock {
	geom := geojson.NewGeometry(orb.Point{x, y})

	idstr := strconv.Itoa(id)
	feat := featureMock{
		GeojsonFeatureData: api.GeojsonFeatureData{
			Type:  "Feature",
			ID:    idstr,
			Geom:  geom,
			Props: map[string]interface{}{"prop_a": "propA", "prop_b": id, "prop_c": "propC", "prop_d": id % 10},
		},
	}
	return &feat
}

func (fm *featureMock) toJSON(propNames []string) string {
	props := fm.extractProperties(propNames)
	return api.MakeGeojsonFeatureJSON(fm.ID, *fm.Geom, props)
}

func (fm *featureMock) extractProperties(propNames []string) map[string]interface{} {
	props := make(map[string]interface{})
	for _, name := range propNames {
		val, err := fm.getProperty(name)
		if err != nil {
			// panic to avoid having to return error
			panic(fmt.Errorf("Unknown property: %v", name))
		}
		props[name] = val
	}
	return props
}

func (fm *featureMock) getProperty(name string) (interface{}, error) {
	if name == "prop_a" || name == "prop_b" || name == "prop_c" || name == "prop_d" {
		return fm.Props[name], nil
	}
	return nil, fmt.Errorf("Unknown property: %v", name)
}

func (fm *featureMock) newPropsFilteredFeature(props []string) *api.GeojsonFeatureData {
	f := api.GeojsonFeatureData{
		Type:  fm.Type,
		ID:    fm.ID,
		Geom:  fm.Geom,
		Props: map[string]interface{}{},
	}

	for _, p := range props {
		f.Props[p] = fm.Props[p]
	}

	return &f
}

func (fm *featureMock) isFilterMatches(filter []*PropertyFilter) bool {
	for _, cond := range filter {
		val, _ := fm.getProperty(cond.Name)
		valStr := fmt.Sprintf("%v", val)
		if cond.Value != valStr {
			return false
		}
	}
	return true
}

func doFilter(features []*featureMock, filter []*PropertyFilter) []*featureMock {
	var result []*featureMock
	for _, feat := range features {
		if feat.isFilterMatches(filter) {
			result = append(result, feat)
		}
	}
	return result
}

func doLimit(features []*featureMock, limit int, offset int) []*featureMock {
	start := 0
	end := len(features)
	// handle limit/offset (offset is only respected if limit present)
	if limit < len(features) {
		start = offset
		end = offset + limit
		if end >= len(features) {
			end = len(features)
		}
	}
	return features[start:end]
}

func MakeFeatureMockPointAsJSON(id int, x float64, y float64, columns []string) string {
	feat := makeFeatureMockPoint(id, x, y)
	return feat.toJSON(columns)
}

func MakeFeaturesMockPoint(extent api.Extent, nx int, ny int) []*featureMock {
	basex := extent.Minx
	basey := extent.Miny
	dx := (extent.Maxx - extent.Minx) / float64(nx)
	dy := (extent.Maxy - extent.Miny) / float64(ny)

	n := nx * ny
	features := make([]*featureMock, n)
	index := 0
	for ix := 0; ix < nx; ix++ {
		for iy := 0; iy < ny; iy++ {
			id := index + 1
			x := basex + dx*float64(ix)
			y := basey + dy*float64(iy)
			features[index] = makeFeatureMockPoint(id, x, y)
			//fmt.Println(features[index])

			index++
		}
	}
	return features
}
