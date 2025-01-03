package data

/*
 Copyright 2019 - 2025 Crunchy Data Solutions, Inc.
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
	"context"
	"fmt"
	"strconv"

	log "github.com/sirupsen/logrus"
)

type CatalogMock struct {
	TableDefs    []*Table
	tableData    map[string][]*featureMock
	FunctionDefs []*Function
}

var instance CatalogMock

// CatMockInstance tbd
func CatMockInstance() *CatalogMock {
	log.Printf("Using Test Catalog data")
	// TODO: make a singleton
	instance = newCatalogMock()
	return &instance
}

func newCatalogMock() CatalogMock {
	// must be in synch with featureMock type
	propNames := []string{"prop_a", "prop_b", "prop_c", "prop_d"}
	types := map[string]string{
		"prop_a": "text",
		"prop_b": "int",
		"prop_c": "text",
		"prop_d": "int",
	}
	jtypes := []string{"string", "number", "string", "number"}
	colDesc := []string{"Property A", "Property B", "Property C", "Property D"}

	layerA := &Table{
		ID:          "mock_a",
		Title:       "Mock A",
		Description: "This dataset contains mock data about A (9 points)",
		Extent:      Extent{Minx: -120, Miny: 40, Maxx: -74, Maxy: 50},
		Srid:        4326,
		Columns:     propNames,
		DbTypes:     types,
		JSONTypes:   jtypes,
		ColDesc:     colDesc,
	}

	layerB := &Table{
		ID:          "mock_b",
		Title:       "Mock B",
		Description: "This dataset contains mock data about B (100 points)",
		Extent:      Extent{Minx: -75, Miny: 45, Maxx: -74, Maxy: 46},
		Srid:        4326,
		Columns:     propNames,
		DbTypes:     types,
		JSONTypes:   jtypes,
		ColDesc:     colDesc,
	}

	layerC := &Table{
		ID:          "mock_c",
		Title:       "Mock C",
		Description: "This dataset contains mock data about C (10000 points)",
		Extent:      Extent{Minx: -120, Miny: 40, Maxx: -74, Maxy: 60},
		Srid:        4326,
		Columns:     propNames,
		DbTypes:     types,
		JSONTypes:   jtypes,
		ColDesc:     colDesc,
	}

	tableData := map[string][]*featureMock{}
	tableData["mock_a"] = makePointFeatures(layerA.Extent, 3, 3)
	tableData["mock_b"] = makePointFeatures(layerB.Extent, 10, 10)
	tableData["mock_c"] = makePointFeatures(layerC.Extent, 100, 100)

	var tables []*Table
	tables = append(tables, layerA)
	tables = append(tables, layerB)
	tables = append(tables, layerC)

	funA := &Function{
		ID:          "fun_a",
		Schema:      "postgisftw",
		Name:        "fun_a",
		Description: "Function A",
		InNames:     []string{"in_param1"},
		InDbTypes:   []string{"text"},
		InTypeMap: map[string]string{
			"in_param1": "text",
		},
		InDefaults:   []string{"aa"},
		NumNoDefault: 0,
		OutNames:     []string{"out_param1"},
		OutDbTypes:   []string{"text"},
		OutJSONTypes: []string{"string"},
		Types: map[string]string{
			"in_param1": "text",
		},
		GeometryColumn: "",
		IDColumn:       "",
	}
	funB := &Function{
		ID:          "fun_b",
		Schema:      "postgisftw",
		Name:        "fun_b",
		Description: "Function B",
		InNames:     []string{"in_param1"},
		InDbTypes:   []string{"int"},
		InTypeMap: map[string]string{
			"in_param1": "int",
		},
		InDefaults:   []string{"999"},
		NumNoDefault: 0,
		OutNames:     []string{"out_geom", "out_id", "out_param1"},
		OutDbTypes:   []string{"geometry", "int", "text"},
		OutJSONTypes: []string{"geometry", "int", "string"},
		Types: map[string]string{
			"in_param1":  "int",
			"out_geom":   "geometry",
			"out_id":     "int",
			"out_param1": "text",
		},
		GeometryColumn: "",
		IDColumn:       "",
	}
	funNoParam := &Function{
		ID:           "fun_noparam",
		Schema:       "postgisftw",
		Name:         "fun_noparam",
		Description:  "Function with no parameters",
		InNames:      []string{},
		InDbTypes:    []string{},
		InTypeMap:    map[string]string{},
		InDefaults:   []string{},
		NumNoDefault: 0,
		OutNames:     []string{"out_geom", "out_id", "out_param1"},
		OutDbTypes:   []string{"geometry", "int", "text"},
		OutJSONTypes: []string{"geometry", "int", "string"},
		Types: map[string]string{
			"in_param1":  "int",
			"out_geom":   "geometry",
			"out_id":     "int",
			"out_param1": "text",
		},
		GeometryColumn: "",
		IDColumn:       "",
	}
	funDefs := []*Function{
		funA,
		funB,
		funNoParam,
	}
	catMock := CatalogMock{
		TableDefs:    tables,
		tableData:    tableData,
		FunctionDefs: funDefs,
	}

	return catMock
}

func (cat *CatalogMock) SetIncludeExclude(includeList []string, excludeList []string) {
}

func (cat *CatalogMock) Close() {
	// this is a no-op
}

func (cat *CatalogMock) Tables() ([]*Table, error) {
	return cat.TableDefs, nil
}

func (cat *CatalogMock) TableReload(name string) {
	// no-op for mock data
}

func (cat *CatalogMock) TableByName(name string) (*Table, error) {
	for _, lyr := range cat.TableDefs {
		if lyr.ID == name {
			return lyr, nil
		}
	}
	// not found - indicated by nil value returned
	return nil, nil
}

func (cat *CatalogMock) TableFeatures(ctx context.Context, name string, param *QueryParam) ([]string, error) {
	features, ok := cat.tableData[name]
	if !ok {
		// table not found - indicated by nil value returned
		return nil, nil
	}
	featFilt := doFilter(features, param.Filter)
	featuresLim := doLimit(featFilt, param.Limit, param.Offset)
	// handle empty property list
	propNames := cat.TableDefs[0].Columns
	if len(param.Columns) > 0 {
		propNames = param.Columns
	}
	return featuresToJSON(featuresLim, propNames), nil
}

func (cat *CatalogMock) TableFeature(ctx context.Context, name string, id string, param *QueryParam) (string, error) {
	features, ok := cat.tableData[name]
	if !ok {
		// table not found - indicated by empty value returned
		return "", nil
	}
	index, err := strconv.Atoi(id)
	if err != nil {
		// a malformed int is treated as feature not found
		return "", nil
	}

	// TODO: return not found if index out of range
	if index < 0 || index >= len(features) {
		return "", nil
	}
	// handle empty property list
	propNames := cat.TableDefs[0].Columns
	if len(param.Columns) > 0 {
		propNames = param.Columns
	}

	return features[index].toJSON(propNames), nil
}

func (cat *CatalogMock) Functions() ([]*Function, error) {
	return cat.FunctionDefs, nil
}

func (cat *CatalogMock) FunctionByName(name string) (*Function, error) {
	for _, fn := range cat.FunctionDefs {
		if fn.Schema+"."+fn.ID == name {
			return fn, nil
		}
	}
	// not found - indicated by nil value returned
	return nil, nil
}

func (cat *CatalogMock) FunctionFeatures(ctx context.Context, name string, args map[string]string, param *QueryParam) ([]string, error) {
	// TODO:
	return nil, nil
}

func (cat *CatalogMock) FunctionData(ctx context.Context, name string, args map[string]string, param *QueryParam) ([]map[string]interface{}, error) {
	// TODO:
	return nil, nil
}

func makePointFeatures(extent Extent, nx int, ny int) []*featureMock {
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

type featureMock struct {
	ID    string
	Geom  string
	PropA string
	PropB int
	PropC string
	PropD int
}

func makeFeatureMockPoint(id int, x float64, y float64) *featureMock {
	geomFmt := `{"type": "Point","coordinates": [ %v, %v ]  }`
	geomStr := fmt.Sprintf(geomFmt, x, y)

	idstr := strconv.Itoa(id)
	feat := featureMock{idstr, geomStr, "propA", id, "propC", id % 10}
	return &feat
}

func (fm *featureMock) toJSON(propNames []string) string {
	props := fm.extractProperties(propNames)
	return makeFeatureJSON(fm.ID, fm.Geom, props)
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
	if name == "prop_a" {
		return fm.PropA, nil
	}
	if name == "prop_b" {
		return fm.PropB, nil
	}
	if name == "prop_c" {
		return fm.PropC, nil
	}
	if name == "prop_d" {
		return fm.PropD, nil
	}
	return nil, fmt.Errorf("Unknown property: %v", name)
}

func doFilter(features []*featureMock, filter []*PropertyFilter) []*featureMock {
	var result []*featureMock
	for _, feat := range features {
		if isFilterMatches(feat, filter) {
			result = append(result, feat)
		}
	}
	return result
}

func isFilterMatches(feature *featureMock, filter []*PropertyFilter) bool {
	for _, cond := range filter {
		val, _ := feature.getProperty(cond.Name)
		valStr := fmt.Sprintf("%v", val)
		if cond.Value != valStr {
			return false
		}
	}
	return true
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

func featuresToJSON(features []*featureMock, propNames []string) []string {
	n := len(features)
	featJSON := make([]string, n)
	for i := 0; i < n; i++ {
		featJSON[i] = features[i].toJSON(propNames)
	}
	return featJSON
}
