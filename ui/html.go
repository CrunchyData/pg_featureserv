package ui

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

var templatePage = `<!doctype html>
<html lang="en">
<head>
<meta charset="utf-8">
<title>{{ .config.Metadata.Title }}</title>

{{ if .context.UseMap }}
<link rel="stylesheet" href="https://openlayers.org/en/v4.6.5/css/ol.css" type="text/css">
<style>
.map {
	height: 400px;
	width: 400px;
	margin-bottom: 10px;
}
.arrow_box {
	border-radius: 5px;
	padding: 10px;
}
.arrow_box {
	position: relative;
	background: #fff;
	border: 1px solid #003c88;
}
.arrow_box:after, .arrow_box:before {
	top: 100%;
	left: 50%;
	border: solid transparent;
	content: " ";
	height: 0;
	width: 0;
	position: absolute;
	pointer-events: none;
}
.arrow_box:after {
	border-color: rgba(255, 255, 255, 0);
	border-top-color: #fff;
	border-width: 10px;
	margin-left: -10px;
}
.arrow_box:before {
	border-color: rgba(153, 153, 153, 0);
	border-top-color: #003c88;
	border-width: 11px;
	margin-left: -11px;
}
</style>
{{ end }}

{{ range .links }}
<link rel="{{ .Rel }}" type="application/json" href="{{ .Href }}"/>
{{ end }}
<style>
body {	font-family: sans-serif; }
a { color:blue; text-decoration: none;}
a:hover, a:focus { text-decoration: underline; }
.page-title { font-family: monospace; font-size: 20pt; }
.coll-title { font-weight: bold; padding-right: 10px }
.json-link { font-size: 12px; background-color: lightgray;  color: blue;
	padding: 2px 8px 2px 8px; border-radius: 10px;
}
.coll-desc { font-style: italic; }
.coll-meta-field { font-weight: bold; }
.crumbs { font-size: 11pt; margin: 8px; }
</style>
</head>

<body>
	<header style='color: white; padding: 6px; background-color: #0167B4'>
		<div class='page-title'><a href="{{ .context.UrlHome }}"><span style='color:white'>{{ .config.Metadata.Title }}</span></a></div>
	</header>
	{{ .body }}
	<hr/>
	<div style='font-size: 10px; font-style: italic;'>{{ .context.AppName }} - Version {{ .context.AppVersion }}</div>
</body>
</html>`

var templateHome = `
<div class='crumbs'>Home
<a style='margin-left: 20px' class='json-link' href='{{ .context.UrlJSON }}' title='JSON document for this page'>JSON</a>
</div>
<hr>
<div>{{ .config.Metadata.Description }}</div>
<h2>Collections</h2>
<a href="collections.html">View the collections</a>

<h2>Conformance</h2>
<a href="conformance.html">View the conformance</a>
`

var templateConformance = `
<div class='crumbs'><a href="{{ .context.UrlHome }}">Home</a>
/ Conformance
<a style='margin-left: 20px' class='json-link' href='{{ .context.UrlJSON }}' title='JSON document for this page'>JSON</a>
</div>
<hr>
<h2>Conformance</h2>
<ul>
{{ range .data.ConformsTo }}
	<li><a href="{{ . }}">{{ . }}</a></li>
{{ end }}
</ul>
`

var templateCollections = `
<div class='crumbs'><a href="{{ .context.UrlHome }}">Home</a>
/ Collections
<a style='margin-left: 20px' class='json-link' href='{{ .context.UrlJSON }}' title='JSON document for this page'>JSON</a>
</div>
<hr>
<h2>Feature Collections</h2>

{{ range .data.Collections }}
{{$collTitle := .Title}}
	<div >

	{{ range .Links }}
		{{ if (eq .Rel "self") }}
		<a href="{{ .Href }}"><span class='coll-title'>{{ $collTitle }}</span></a>
		{{ end }}
	{{ end }}
	{{ range .Links }}
		{{ if (eq .Rel "alternate") }}
			<a  class='json-link' href="{{ .Href }}">JSON</a></li>
		{{ end }}
	{{ end }}
	</div>
	<div class='coll-desc'>{{ .Description }}</div>
	<p>
{{ end }}
`

var templateCollection = `
<div class='crumbs'><a href="{{ .context.UrlHome }}">Home</a>
/ <a href="{{ .context.UrlCollections }}">Collections</a>
/ {{ .data.Title }}
<a style='margin-left: 20px' class='json-link' href='{{ .context.UrlJSON }}' title='JSON document for this page'>JSON</a>
</div>
<hr>
<h2>Feature Collection: {{ .data.Title }}</h2>

<div class='coll-desc'>{{ .data.Description }}</div>
<p>
<div ><span class='coll-meta-field'>ID column:</span> {{ .context.Layer.IDColumn }}</div>
<div ><span class='coll-meta-field'>Geometry column:</span> {{ .context.Layer.GeometryColumn }}</div>
<div ><span class='coll-meta-field'>Geometry type:</span> {{ .context.Layer.GeometryType }}</div>
<div ><span class='coll-meta-field'>SRID:</span> {{ .context.Layer.Srid }}</div>
<div ><span class='coll-meta-field'>Extent:</span> {{ .data.Extent }}</div>
<div ><span class='coll-meta-field'>Columns:</span> </div>
{{ range $name, $typ := .context.Layer.Properties }}
<div><span style='padding-left: 10px;'>{{ $name }}</span> :: {{ $typ }}</div>
{{ end }}
</p>

<h3>Features</h3>
	{{ range .data.Links }}
	<div>
		{{ if (eq .Rel "items") }}
			<a href="{{ .Href }}">{{ .Title }}</a>
		{{ end }}
	</div>
	{{ end }}
<p>
`

var templateItems = `
<div class='crumbs'><a href="{{ .context.UrlHome }}">Home</a>
/ <a href="{{ .context.UrlCollections }}">Collections</a>
/ <a href="{{ .context.UrlCollection }}">{{ .context.CollectionTitle }}</a>
/ Items
<a style='margin-left: 20px' class='json-link' href='{{ .context.UrlJSON }}' title='JSON document for this page'>JSON</a>
</div>
<hr>
<h2>Features: {{ .context.CollectionTitle }}</h2>

<div id="map" class="map"></div>
<div id="popup-container" class="arrow_box"></div>

<script>
var geojsonObject = {{ .data }};
</script>
`

var templateItem = `
<div class='crumbs'><a href="{{ .context.UrlHome }}">Home</a>
/ <a href="{{ .context.UrlCollections }}">Collections</a>
/ <a href="{{ .context.UrlCollection }}">{{ .context.CollectionTitle }}</a>
/ <a href="{{ .context.UrlItems }}">Items</a>
/ {{ .context.FeatureID }}
<a style='margin-left: 20px' class='json-link' href='{{ .context.UrlJSON }}' title='JSON document for this page'>JSON</a>
</div>
<hr>
<h2>Feature: {{ .context.FeatureID }}</h2>

<div id="map" class="map"></div>
<div id="popup-container" class="arrow_box"></div>

<script>
var geojsonObject = {{ .data }};
</script>
`

var mapCode = `
<script src="https://openlayers.org/en/v4.6.5/build/ol.js"></script>
<script>
var image = new ol.style.Circle({
	radius: 5,
	fill: new ol.style.Fill({
		color: 'rgb(255, 0, 0)'
	}),
	stroke: new ol.style.Stroke({color: 'red', width: 1})
});
var styles = {
	'Point': new ol.style.Style({
		image: image
	}),
	'LineString': new ol.style.Style({
		stroke: new ol.style.Stroke({
			color: 'red',
			width: 2
		})
	}),
	'MultiLineString': new ol.style.Style({
		stroke: new ol.style.Stroke({
			color: 'red',
			width: 2
		})
	}),
	'MultiPoint': new ol.style.Style({
		image: image
	}),
	'MultiPolygon': new ol.style.Style({
		stroke: new ol.style.Stroke({
			color: 'blue',
			width: 2
		}),
		fill: new ol.style.Fill({
			color: 'rgba(0, 0, 255, 0.1)'
		})
	}),
	'Polygon': new ol.style.Style({
		stroke: new ol.style.Stroke({
			color: 'blue',
			lineDash: [4],
			width: 2
		}),
		fill: new ol.style.Fill({
			color: 'rgba(0, 0, 255, 0.1)'
		})
	}),
	'GeometryCollection': new ol.style.Style({
		stroke: new ol.style.Stroke({
			color: 'magenta',
			width: 2
		}),
		fill: new ol.style.Fill({
			color: 'magenta'
		}),
		image: new ol.style.Circle({
			radius: 10,
			fill: null,
			stroke: new ol.style.Stroke({
				color: 'magenta'
			})
		})
	}),
	'Circle': new ol.style.Style({
		stroke: new ol.style.Stroke({
			color: 'red',
			width: 2
		}),
		fill: new ol.style.Fill({
			color: 'rgba(255,0,0,0.2)'
		})
	})
};
var styleFunction = function(feature) {
	return styles[feature.getGeometry().getType()];
};
var vectorSource = new ol.source.Vector({
	features: (new ol.format.GeoJSON()).readFeatures(geojsonObject, {
		dataProjection: "EPSG:4326",
		featureProjection: "EPSG:3857"
	})
});
var vectorLayer = new ol.layer.Vector({
	source: vectorSource,
	style: styleFunction,
});
var map = new ol.Map({
	layers: [
		new ol.layer.Tile({
			source: new ol.source.OSM({
				"url" : "https://maps.wikimedia.org/osm-intl/{z}/{x}/{y}.png"
			})
		}),
		vectorLayer
	],
	target: 'map',
	controls: ol.control.defaults({
		attributionOptions: {
			collapsible: false
		}
	}),
	view: new ol.View({
		zoom: -10
	})
});
map.getView().fit(vectorLayer.getSource().getExtent(), map.getSize());
var overlay = new ol.Overlay({
	element: document.getElementById('popup-container'),
	positioning: 'bottom-center',
	offset: [0, -10]
});
map.addOverlay(overlay);
map.on('click', function(evt) {
	overlay.setPosition();
	var features = map.getFeaturesAtPixel(evt.pixel);
	var loc = evt.coordinate
	if (features) {
		var identifier = features[0].getId();
		var popup = 'Id: ' + identifier
		if (! document.location.href.includes('/items/')) {
			popup = '<a href="items/' + identifier + '.html">' + popup + '</a>'
		}
		overlay.getElement().innerHTML = popup;
		overlay.setPosition(loc);
	}
});
</script>
`
