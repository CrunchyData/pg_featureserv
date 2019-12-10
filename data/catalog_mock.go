package data

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
	"bytes"
	"fmt"
	"strconv"
	"text/template"

	log "github.com/sirupsen/logrus"
)

type catalogMock struct {
	layers    []*Layer
	layerData map[string][]string
}

var instance Catalog

// CatMockInstance tbd
func CatMockInstance() Catalog {
	log.Printf("Using Test Catalog data")
	// TODO: make a singleton
	instance = newCatalogMock()
	return instance
}

var templateFeaturePoint *template.Template

func init() {
	templateFeaturePoint = template.Must(template.New("feature").Parse(templateStrFeaturePoint))
}

func newCatalogMock() Catalog {

	layerA := &Layer{
		ID:          "mock_a",
		Title:       "Mock A",
		Description: "This dataset contains mock data about A (9 points)",
		Extent:      Extent{Minx: -120, Miny: 40, Maxx: -74, Maxy: 50},
		Srid:        4326,
	}

	layerB := &Layer{
		ID:          "mock_b",
		Title:       "Mock B",
		Description: "This dataset contains mock data about B (100 points)",
		Extent:      Extent{Minx: -75, Miny: 45, Maxx: -74, Maxy: 46},
		Srid:        4326,
	}

	layerC := &Layer{
		ID:          "mock_c",
		Title:       "Mock C",
		Description: "This dataset contains mock data about C (10000 points)",
		Extent:      Extent{Minx: -120, Miny: 40, Maxx: -74, Maxy: 60},
		Srid:        4326,
	}

	layerData := map[string][]string{}
	layerData["mock_a"] = makePointFeatures(layerA.Extent, 3, 3)
	layerData["mock_b"] = makePointFeatures(layerB.Extent, 10, 10)
	layerData["mock_c"] = makePointFeatures(layerC.Extent, 100, 100)

	var layers []*Layer
	layers = append(layers, layerA)
	layers = append(layers, layerB)
	layers = append(layers, layerC)

	catMock := catalogMock{
		layers:    layers,
		layerData: layerData,
	}

	return &catMock
}

func (cat *catalogMock) Layers() ([]*Layer, error) {
	return cat.layers, nil
}

func (cat *catalogMock) LayerByName(name string) (*Layer, error) {
	for _, lyr := range cat.layers {
		if lyr.ID == name {
			return lyr, nil
		}
	}
	// not found - indicated by nil value returned
	return nil, nil
}

func (cat *catalogMock) LayerFeatures(name string) ([]string, error) {
	features, ok := cat.layerData[name]
	if !ok {
		//		return []string{}, fmt.Errorf(errMsgBadLayerName, name)
		// layer ot found - indicated by nil value returned
		return nil, nil
	}
	//fmt.Println("LayerFeatures: " + name)
	//fmt.Println(layerData)
	return features, nil
}

func (cat *catalogMock) LayerFeature(name string, id string) (string, error) {
	features, ok := cat.layerData[name]
	if !ok {
		// layer ot found - indicated by empty value returned
		return "", nil
	}
	index, err := strconv.Atoi(id)
	if err != nil {
		// a malformed int is treated as feature not found
		return "", nil
	}

	//fmt.Println("LayerFeatures: " + name)
	//fmt.Println(layerData)

	// TODO: return not found if index out of range
	if index < 0 || index >= len(features) {
		return "", nil
	}
	return features[index], nil
}

var featuresMock = []string{
	`{ "type": "Feature", "id": 1,  "geometry": {"type": "Point","coordinates": [  -75,	  45]  },
	  "properties": { "value": "89.9"  } }`,
	`{ "type": "Feature", "id": 2,  "geometry": {"type": "Point","coordinates": [  -75,	  40]  },
	  "properties": { "value": "89.9"  } }`,
	`{ "type": "Feature", "id": 3,  "geometry": {"type": "Point","coordinates": [  -75,	  35]  },
	  "properties": { "value": "89.9"  } }`,
}

func makePointFeatures(extent Extent, nx int, ny int) []string {
	basex := extent.Minx
	basey := extent.Miny
	dx := (extent.Maxx - extent.Minx) / float64(nx)
	dy := (extent.Maxy - extent.Miny) / float64(ny)

	n := nx * ny
	features := make([]string, n)
	index := 0
	for ix := 0; ix < nx; ix++ {
		for iy := 0; iy < ny; iy++ {
			x := basex + dx*float64(ix)
			y := basey + dy*float64(iy)
			val := fmt.Sprintf("data value %v", index)
			features[index] = makeFeaturePoint(index, x, y, val)
			//fmt.Println(features[index])

			index++
		}
	}
	return features
}

type featurePointMock struct {
	ID  int
	X   float64
	Y   float64
	Val string
}

var templateStrFeaturePoint = `{ "type": "Feature", "id": {{ .ID }},
"geometry": {"type": "Point","coordinates": [  {{ .X }}, {{ .Y }} ]  },
"properties": { "value": "{{ .Val }}"  } }`

func makeFeaturePoint(id int, x float64, y float64, val string) string {
	feat := featurePointMock{id, x, y, val}
	var tempOut bytes.Buffer
	//	tempOut.Reset()
	templateFeaturePoint.Execute(&tempOut, feat)
	return tempOut.String()
}
