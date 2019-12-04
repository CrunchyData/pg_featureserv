package api

import "strings"

const (
	// ContentTypeJSON tbd
	ContentTypeJSON = "application/json"

	// ContentTypeGeoJSON tbd
	ContentTypeGeoJSON = "application/geo+json"

	// ContentTypeHTML tbd
	ContentTypeHTML = "text/html"

	FormatJSON = "json"

	FormatHTML = "html"
)

// ContentType tbd
func ContentType(format string) string {
	switch format {
	case FormatJSON:
		return ContentTypeJSON
	case FormatHTML:
		return ContentTypeHTML
	}
	return ""
}

func PathFormat(path string) string {
	if strings.HasSuffix(path, ".html") {
		return FormatHTML
	}
	return FormatJSON
}

func PathStripFormat(path string) string {
	if strings.HasSuffix(path, ".html") || strings.HasSuffix(path, ".json") {
		return path[0 : len(path)-5]
	}
	return path
}
