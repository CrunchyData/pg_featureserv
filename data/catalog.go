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

// Catalog tbd
type Catalog interface {
	Layers() ([]*Layer, error)
	IsLayer(name string) (bool, error)
	LayerByName(name string) (*Layer, error)
	LayerFeatures(name string) ([]string, error)
	LayerFeature(name string, id string) (string, error)
}

// Layer tbd
type Layer struct {
	ID           string
	Schema       string
	Table        string
	Title        string
	Description  string
	GeometryType string
	//Properties     map[string]string
	GeometryColumn string
	Srid           int
	IDColumn       string
	Extent         Extent
}

// Extent of a layer
type Extent struct {
	Minx, Miny, Maxx, Maxy float64
}

const (
	errMsgLayerNotFound   = "Layer not found: %v"
	errMsgFeatureNotFound = "Feature not found: %v"
)
