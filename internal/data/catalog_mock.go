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

 Date     : October 2022
 Authors  : Benoit De Mezzo (benoit dot de dot mezzo at oslandia dot com)
        	Amaury Zarzelli (amaury dot zarzelli at ign dot fr)
			Jean-philippe Bazonnais (jean-philippe dot bazonnais at ign dot fr)
			Nicolas Revelant (nicolas dot revelant at ign dot fr)
*/

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"hash/fnv"
	"strconv"
	"time"

	"github.com/CrunchyData/pg_featureserv/internal/api"

	log "github.com/sirupsen/logrus"
)

type CatalogMock struct {
	TableDefs    []*api.Table
	tableData    map[string][]*featureMock
	FunctionDefs []*api.Function
	cache        Cacher
}

var instance CatalogMock

// CatMockInstance tbd
func CatMockInstance() *CatalogMock {
	log.Debug("Using Test Catalog data")
	// TODO: make a singleton
	instance = newCatalogMock()
	return &instance
}

func newCatalogMock() CatalogMock {
	// must be in synch with featureMock type
	propNames := []string{"prop_a", "prop_b", "prop_c", "prop_d"}
	types := map[string]api.Column{
		"prop_a": {Index: 0, Type: "text", IsRequired: true},
		"prop_b": {Index: 1, Type: "int", IsRequired: true},
		"prop_c": {Index: 2, Type: "text", IsRequired: false},
		"prop_d": {Index: 3, Type: "int", IsRequired: false},
	}
	jtypes := []api.JSONType{api.JSONTypeString, api.JSONTypeNumber, api.JSONTypeString, api.JSONTypeNumber}
	colDesc := []string{"Property A", "Property B", "Property C", "Property D"}

	layerA := &api.Table{
		ID:           "mock_a",
		Title:        "Mock A",
		Description:  "This dataset contains mock data about A (9 points)",
		Extent:       api.Extent{Minx: -120, Miny: 40, Maxx: -74, Maxy: 50},
		Srid:         4326,
		GeometryType: "Point",
		Columns:      propNames,
		DbTypes:      types,
		JSONTypes:    jtypes,
		ColDesc:      colDesc,
	}

	layerB := &api.Table{
		ID:           "mock_b",
		Title:        "Mock B",
		Description:  "This dataset contains mock data about B (100 points)",
		Extent:       api.Extent{Minx: -75, Miny: 45, Maxx: -74, Maxy: 46},
		Srid:         4326,
		GeometryType: "Point",
		Columns:      propNames,
		DbTypes:      types,
		JSONTypes:    jtypes,
		ColDesc:      colDesc,
	}

	layerC := &api.Table{
		ID:           "mock_c",
		Title:        "Mock C",
		Description:  "This dataset contains mock data about C (10000 points)",
		Extent:       api.Extent{Minx: -120, Miny: 40, Maxx: -74, Maxy: 60},
		Srid:         4326,
		GeometryType: "Point",
		Columns:      propNames,
		DbTypes:      types,
		JSONTypes:    jtypes,
		ColDesc:      colDesc,
	}

	tableData := map[string][]*featureMock{}
	tableData["mock_a"] = MakeFeaturesMockPoint(layerA.Extent, 3, 3)
	tableData["mock_b"] = MakeFeaturesMockPoint(layerB.Extent, 10, 10)
	tableData["mock_c"] = MakeFeaturesMockPoint(layerC.Extent, 100, 100)

	var tables []*api.Table
	tables = append(tables, layerA)
	tables = append(tables, layerB)
	tables = append(tables, layerC)

	funA := &api.Function{
		ID:          "fun_a",
		Schema:      "postgisftw",
		Name:        "fun_a",
		Description: "Function A",
		InNames:     []string{"in_param1"},
		InDbTypes:   []string{"text"},
		InTypeMap: map[string]api.PGType{
			"in_param1": api.PGTypeText,
		},
		InDefaults:   []string{"aa"},
		NumNoDefault: 0,
		OutNames:     []string{"out_param1"},
		OutDbTypes:   []string{"text"},
		OutJSONTypes: []api.JSONType{api.JSONTypeString},
		Types: map[string]api.PGType{
			"in_param1": api.PGTypeText,
		},
		GeometryColumn: "",
		IDColumn:       "",
	}
	funB := &api.Function{
		ID:          "fun_b",
		Schema:      "postgisftw",
		Name:        "fun_b",
		Description: "Function B",
		InNames:     []string{"in_param1"},
		InDbTypes:   []string{"int"},
		InTypeMap: map[string]api.PGType{
			"in_param1": api.PGTypeInt,
		},
		InDefaults:   []string{"999"},
		NumNoDefault: 0,
		OutNames:     []string{"out_geom", "out_id", "out_param1"},
		OutDbTypes:   []string{"geometry", "int", "text"},
		OutJSONTypes: []api.JSONType{api.JSONTypeGeometry, api.JSONTypeNumber, api.JSONTypeString},
		Types: map[string]api.PGType{
			"in_param1":  api.PGTypeInt,
			"out_geom":   api.PGTypeGeometry,
			"out_id":     api.PGTypeInt,
			"out_param1": api.PGTypeText,
		},
		GeometryColumn: "",
		IDColumn:       "",
	}
	funNoParam := &api.Function{
		ID:           "fun_noparam",
		Schema:       "postgisftw",
		Name:         "fun_noparam",
		Description:  "Function with no parameters",
		InNames:      []string{},
		InDbTypes:    []string{},
		InTypeMap:    map[string]api.PGType{},
		InDefaults:   []string{},
		NumNoDefault: 0,
		OutNames:     []string{"out_geom", "out_id", "out_param1"},
		OutDbTypes:   []string{"geometry", "int", "text"},
		OutJSONTypes: []api.JSONType{api.JSONTypeGeometry, api.JSONTypeNumber, api.JSONTypeString},
		Types: map[string]api.PGType{
			"in_param1":  api.PGTypeInt,
			"out_geom":   api.PGTypeGeometry,
			"out_id":     api.PGTypeInt,
			"out_param1": api.PGTypeText,
		},
		GeometryColumn: "",
		IDColumn:       "",
	}
	funDefs := []*api.Function{
		funA,
		funB,
		funNoParam,
	}
	cache := makeCache()
	catMock := CatalogMock{
		TableDefs:    tables,
		tableData:    tableData,
		FunctionDefs: funDefs,
		cache:        cache,
	}

	return catMock
}

// -------------------------------------------------

func (cat *CatalogMock) Initialize(includeList []string, excludeList []string) {
	// this is a no-op
}

func (cat *CatalogMock) Close() {
	// this is a no-op
}

func (cat *CatalogMock) GetCache() map[string]interface{} {
	return make(map[string]interface{})
}

func (cat *CatalogMock) Tables() ([]*api.Table, error) {
	return cat.TableDefs, nil
}

func (cat *CatalogMock) TableReload(name string) {
	// no-op for mock data
}

func (cat *CatalogMock) TableByName(name string) (*api.Table, error) {
	for _, lyr := range cat.TableDefs {
		if lyr.ID == name {
			return lyr, nil
		}
	}
	// not found - indicated by nil value returned
	return nil, nil
}

func (cat *CatalogMock) TableFeatures(ctx context.Context, name string, param *QueryParam) ([]*api.GeojsonFeatureData, error) {
	features, ok := cat.tableData[name]
	if !ok {
		// table not found - indicated by nil value returned
		return nil, nil
	}
	featFilt := doFilter(features, param.Filter)
	featuresLim := doLimit(featFilt, param.Limit, param.Offset)

	var propNames []string
	if len(param.Columns) > 0 {
		propNames = param.Columns
	}
	featureClones := make([]*api.GeojsonFeatureData, len(featuresLim))
	for i, f := range featuresLim {
		featureClones[i] = f.newPropsFilteredFeature(propNames)
	}
	return featureClones, nil
}

func (cat *CatalogMock) TableFeature(ctx context.Context, name string, id string, param *QueryParam) (*api.GeojsonFeatureData, error) {
	features, ok := cat.tableData[name]
	if !ok {
		// table not found - indicated by empty value returned
		return nil, nil
	}
	index, err := strconv.Atoi(id)
	if err != nil {
		// a malformed int is treated as feature not found
		return nil, nil
	}

	index--

	// return not found if index out of range
	if index < 0 || index >= len(features) {
		return nil, nil
	}

	var propNames []string
	if len(param.Columns) > 0 {
		propNames = param.Columns
	}

	for elementIdx, feature := range features {
		if feature.ID == id {

			feature_data := features[elementIdx].newPropsFilteredFeature(propNames)
			weakEtag := feature.WeakEtag
			feature_data.WeakEtag = weakEtag
			cat.cache.AddWeakEtag(weakEtag, map[string]interface{}{"lastModified": time.Now().String()})
			return feature_data, nil
		}
	}

	// not found
	return nil, nil

}

// returns the number of feature for a specific table
func (cat *CatalogMock) TableSize(tableName string) int64 {
	return int64(len(cat.tableData[tableName]))
}

func (cat *CatalogMock) AddTableFeature(ctx context.Context, tableName string, jsonData []byte) (int64, error) {
	var newFeature featureMock

	var schemaObject api.GeojsonFeatureData
	err := json.Unmarshal(jsonData, &schemaObject)
	if err != nil {
		return 0, err
	}

	maxId := cat.TableSize(tableName)

	newFeature.Type = "Feature"
	newFeature.ID = fmt.Sprintf("%d", maxId+1)
	newFeature.Geom = schemaObject.Geom
	newFeature.Props = make(map[string]interface{})
	newFeature.Props["prop_a"] = schemaObject.Props["prop_a"].(string)
	newFeature.Props["prop_b"] = int(schemaObject.Props["prop_b"].(float64))
	newFeature.Props["prop_c"] = schemaObject.Props["prop_c"].(string)
	newFeature.Props["prop_d"] = int(schemaObject.Props["prop_d"].(float64))

	sum := fnv.New32a()
	encodedContent, _ := json.Marshal(schemaObject.Geom)
	sum.Write(encodedContent)
	weakEtag := fmt.Sprint(sum.Sum32())
	newFeature.WeakEtag = weakEtag

	cat.cache.AddWeakEtag(weakEtag, map[string]interface{}{"lastModified": time.Now().String()})

	cat.tableData[tableName] = append(cat.tableData[tableName], &newFeature)
	return maxId + 1, nil
}

func (cat *CatalogMock) PartialUpdateTableFeature(ctx context.Context, tableName string, id string, jsonData []byte) error {

	var schemaObject api.GeojsonFeatureData
	err1 := json.Unmarshal(jsonData, &schemaObject)
	if err1 != nil {
		return err1
	}

	index, err2 := strconv.ParseInt(id, 10, 64)
	if err2 != nil {
		return err2
	}

	oldFeature := cat.tableData[tableName][index-1]

	// update values if exist into json !
	if schemaObject.Geom != nil {
		oldFeature.Geom = schemaObject.Geom
	}

	if schemaObject.Props["prop_a"] != nil {
		oldFeature.Props["prop_a"] = schemaObject.Props["prop_a"].(string)
	}

	if schemaObject.Props["prop_b"] != nil {
		oldFeature.Props["prop_b"] = int(schemaObject.Props["prop_b"].(float64))
	}

	if schemaObject.Props["prop_c"] != nil {
		oldFeature.Props["prop_c"] = schemaObject.Props["prop_c"].(string)
	}

	if schemaObject.Props["prop_d"] != nil {
		oldFeature.Props["prop_d"] = int(schemaObject.Props["prop_d"].(float64))
	}

	propNames := cat.TableDefs[0].Columns
	jsonStr := oldFeature.toJSON(propNames)
	if jsonStr == "" {
		return fmt.Errorf("Error marshalling feature into JSON:: %v", tableName)
	}

	return nil
}

func (cat *CatalogMock) ReplaceTableFeature(ctx context.Context, tableName string, id string, jsonData []byte) error {
	var schemaObject api.GeojsonFeatureData
	err1 := json.Unmarshal(jsonData, &schemaObject)
	if err1 != nil {
		return err1
	}

	index, err2 := strconv.ParseInt(id, 10, 64)
	if err2 != nil {
		return nil
	}

	oldFeature := cat.tableData[tableName][index-1]

	// update values into json
	// all required values should exist sinc it's replace and schema has been checked at this point
	if schemaObject.Geom == nil {
		return fmt.Errorf("Error missing geometry:: %v", tableName)
	}
	oldFeature.Geom = schemaObject.Geom

	if schemaObject.Props["prop_a"] == nil {
		// fails because required property
		return fmt.Errorf("Error missing property:: %v", tableName)
	}
	oldFeature.Props["prop_a"] = schemaObject.Props["prop_a"].(string)

	if schemaObject.Props["prop_b"] == nil {
		// fails because required property
		return fmt.Errorf("Error missing property:: %v", tableName)
	}
	oldFeature.Props["prop_b"] = int(schemaObject.Props["prop_b"].(float64))

	// property not required and should be set to nil (NULL in db)
	if schemaObject.Props["prop_c"] != nil {
		oldFeature.Props["prop_c"] = schemaObject.Props["prop_c"].(string)
	} else {
		delete(oldFeature.Props, "prop_c")
	}

	// property not required and should be set to nil (NULL in db)
	if schemaObject.Props["prop_d"] != nil {
		oldFeature.Props["prop_d"] = int(schemaObject.Props["prop_d"].(float64))
	} else {
		delete(oldFeature.Props, "prop_d")
	}

	propNames := cat.TableDefs[0].Columns
	jsonStr := oldFeature.toJSON(propNames)
	if jsonStr == "" {
		return fmt.Errorf("Error marshalling feature into JSON:: %v", tableName)
	}

	return nil
}

func (cat *CatalogMock) DeleteTableFeature(ctx context.Context, tableName string, id string) error {

	features, ok := cat.tableData[tableName]
	if !ok {
		return errors.New("Table not found")
	}

	for elementIdx, feature := range features {
		if feature.ID == id {
			cat.tableData[tableName] = append(features[:elementIdx], features[(elementIdx+1):]...)
			return nil
		}
	}
	return errors.New("Feature not found")
}

func (cat *CatalogMock) CheckStrongEtags(etagsList []string) (bool, error) {
	for _, strongEtag := range etagsList {
		found, err := cat.cache.ContainsWeakEtag(strongEtag)
		if err != nil {
			return false, err
		}
		if found {
			return true, nil
		}
	}
	return false, nil
}

func (cat *CatalogMock) Functions() ([]*api.Function, error) {
	return cat.FunctionDefs, nil
}

func (cat *CatalogMock) FunctionByName(name string) (*api.Function, error) {
	for _, fn := range cat.FunctionDefs {
		if fn.Schema+"."+fn.ID == name {
			return fn, nil
		}
	}
	// not found - indicated by nil value returned
	return nil, nil
}

func (cat *CatalogMock) FunctionFeatures(ctx context.Context, name string, args map[string]string, param *QueryParam) ([]*api.GeojsonFeatureData, error) {
	// TODO:
	return nil, nil
}

func (cat *CatalogMock) FunctionData(ctx context.Context, name string, args map[string]string, param *QueryParam) ([]map[string]interface{}, error) {
	// TODO:
	return nil, nil
}
