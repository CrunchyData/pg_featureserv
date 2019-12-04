package data

// Catalog tbd
type Catalog interface {
	Layers() ([]*Layer, error)
	LayerByName(name string) (*Layer, error)
	LayerFeatures(name string) ([]string, error)
	LayerFeature(name string, id string) (string, error)
}

// Layer tbd
type Layer struct {
	Name        string
	Title       string
	Description string
	Extent      Extent
	Crs         string
}

// Extent of a layer
type Extent struct {
	Minx, Miny, Maxx, Maxy float64
}
