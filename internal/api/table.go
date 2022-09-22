package api

import "fmt"

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

// Column holds metadata for column objects
type Column struct {
	Index      int
	Type       PGType
	IsRequired bool
}

// Table holds metadata for table/view objects
type Table struct {
	ID             string
	Schema         string
	Table          string
	Title          string
	Description    string
	GeometryType   string
	GeometryColumn string
	IDColumn       string
	Srid           int
	Extent         Extent
	Columns        []string
	DbTypes        map[string]Column
	JSONTypes      []JSONType
	ColDesc        []string
}

// Check the existence of table fields from json data
func (tbl *Table) CheckTableFields(props map[string]interface{}) (bool, error) {
	p := props["properties"]
	if p != nil {
		props := props["properties"].(map[string]interface{})
		for k := range props {
			if !func(s []string, e string) bool {
				for _, a := range s {
					if a == e {
						return true
					}
				}
				return false
			}(tbl.Columns, k) {
				return false, fmt.Errorf("Properties not conform with field table: %v", k)
			}
		}
	}
	return true, nil
}

func (tbl *Table) extendAsBbox() *Bbox {
	// extent bbox is always in 4326 for now
	crs := "http://www.opengis.net/def/crs/EPSG/0/4326"
	return &Bbox{
		Crs:    crs,
		Extent: []float64{tbl.Extent.Minx, tbl.Extent.Miny, tbl.Extent.Maxx, tbl.Extent.Maxy},
	}
}

func (tbl *Table) NewCollectionInfo() *CollectionInfo {
	doc := CollectionInfo{
		Name:        tbl.ID,
		Title:       tbl.Title,
		Description: tbl.Description,
		Extent: &CollectionExtent{
			Spatial: tbl.extendAsBbox(),
		},
	}
	return &doc
}

func (tbl *Table) TableProperties() []*Property {
	props := make([]*Property, len(tbl.Columns))
	for i, name := range tbl.Columns {
		props[i] = &Property{
			Name:        name,
			Type:        string(tbl.JSONTypes[i]),
			Description: tbl.ColDesc[i],
		}
	}
	return props
}
