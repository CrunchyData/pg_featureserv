package data

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"text/template"
)

type catalogMock struct {
	collectionsData []*Layer
}

var layers []*Layer
var layerData map[string][]string
var catMock Catalog

func init() {
	layers = append(layers, &Layer{
		Name:        "mock_a",
		Title:       "Mock A",
		Description: "This dataset contains mock data about A",
		Extent:      Extent{Minx: 0, Miny: 0, Maxx: 80, Maxy: 90},
		Crs:         "crs1",
	})
	layers = append(layers, &Layer{
		Name:        "mock_b",
		Title:       "Mock B",
		Description: "This dataset contains mock data about B (100 points)",
		Extent:      Extent{Minx: -130, Miny: 40, Maxx: -120, Maxy: 60},
		Crs:         "crs1",
	})
	catMock = &catalogMock{
		collectionsData: layers,
	}

	features := []string{
		`{ "type": "Feature", "id": 1,  "geometry": {"type": "Point","coordinates": [  -75,	  45]  },
		  "properties": { "value": "89.9"  } }`,
		`{ "type": "Feature", "id": 2,  "geometry": {"type": "Point","coordinates": [  -75,	  40]  },
		  "properties": { "value": "89.9"  } }`,
		`{ "type": "Feature", "id": 3,  "geometry": {"type": "Point","coordinates": [  -75,	  35]  },
		  "properties": { "value": "89.9"  } }`,
	}
	layerData = map[string][]string{}
	layerData["mock_a"] = features
	//layerData["mock_b"] = features
	layerData["mock_b"] = makeFeatures(10, 10)
}

// InstanceCatMock tbd
func InstanceCatMock() Catalog {
	return catMock
}

func (cat *catalogMock) Layers() ([]*Layer, error) {
	return cat.collectionsData, nil
}

func (cat *catalogMock) LayerByName(name string) (*Layer, error) {
	for _, coll := range cat.collectionsData {
		if coll.Name == name {
			return coll, nil
		}
	}
	// not found
	return nil, fmt.Errorf("Invalid collection name: %v", name)
}

func (cat *catalogMock) LayerFeatures(name string) ([]string, error) {
	features := layerData[name]
	//fmt.Println("LayerFeatures: " + name)
	//fmt.Println(layerData)
	return features, nil
}

func (cat *catalogMock) LayerFeature(name string, id string) (string, error) {
	features, ok := layerData[name]
	if !ok {
		return "", errors.New("No such layer: " + name)
	}
	index, err := strconv.Atoi(id)
	if err != nil {
		return "", errors.New("No such feature id: " + id)
	}

	//fmt.Println("LayerFeatures: " + name)
	//fmt.Println(layerData)
	return features[index], nil
}

type featurePointMock struct {
	ID  int
	X   float64
	Y   float64
	Val string
}

var templateFeaturePoint = `{ "type": "Feature", "id": {{ .ID }},
"geometry": {"type": "Point","coordinates": [  {{ .X }}, {{ .Y }} ]  },
"properties": { "value": "{{ .Val }}"  } }`

func makeFeatures(nx int, ny int) []string {
	tmpl, err := template.New("feature").Parse(templateFeaturePoint)
	if err != nil {
		panic(err)
	}
	n := nx * ny
	features := make([]string, n)
	var tempOut bytes.Buffer
	index := 0
	for ix := 0; ix < nx; ix++ {
		for iy := 0; iy < ny; iy++ {
			x := -75 + 0.01*float64(ix)
			y := 45 + 0.01*float64(iy)
			val := fmt.Sprintf("data value %v", index)
			feat := featurePointMock{index, x, y, val}
			tempOut.Reset()
			tmpl.Execute(&tempOut, feat)
			features[index] = tempOut.String()
			//fmt.Println(features[index])

			index++
		}
	}
	return features
}
