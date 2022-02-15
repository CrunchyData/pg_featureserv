---
title: "Example: Function to generate a Geographic Grid"
date:
draft: false
weight: 220
---

This example shows how to generate geometry data from a function, controlled by some input parameters.

This particular function does not query an existing table in the database; rather, it uses PostGIS functions to generate spatial data. Grids generated in this way could be used for data visualization, analysis, or clustering.

## Create a spatial function that generates a grid over a desired area

```sql
CREATE OR REPLACE FUNCTION postgisftw.geo_grid(
  num_x integer DEFAULT 10,
  num_y integer DEFAULT 10,
  lon_min numeric DEFAULT -180.0,
  lat_min numeric DEFAULT -90.0,
  lon_max numeric DEFAULT 180.0,
  lat_max numeric DEFAULT 90.0)
RETURNS TABLE(id text, geom geometry)
AS $$
DECLARE
    dlon numeric;
    dlat numeric;
BEGIN
    dlon := (lon_max - lon_min) / num_x;
    dlat := (lat_max - lat_min) / num_y;
	RETURN QUERY
		SELECT
			x.x::text || '_' || y.y::text AS id,
			ST_MakeEnvelope(
                lon_min + (x.x - 1) * dlon, lat_min + (y.y - 1) * dlat,
                lon_min + x.x * dlon, lat_min + y.y * dlat, 4326
             ) AS geom
		FROM generate_series(1, num_x) AS x(x)
        CROSS JOIN generate_series(1, num_y) AS y(y);
END;
$$
LANGUAGE 'plpgsql'
STABLE
STRICT;
```

Notes:

* The `geo_grid` function accepts a num_x and a num_y value to define the number of grid cells along the longitudinal (X) and latitudinal (Y) axes respectively. It also takes in minimum and maximum longitude and latitude values for the map area we want covered.
* The function first calculates the lengths of the sides of the grid (dlon and dlat).
* A CROSS JOIN on two generate_series() functions produces X and Y indices for each grid cell.
* The PostGIS function [ST_MakeEnvelope()](https://postgis.net/docs/ST_MakeEnvelope.html) contructs a rectangular polygon for each cell. An `id` value is also returned that encodes the grid index.

## Example of API query

`http://localhost:9000/functions/geo_grid/items?num_x=5&num_y=5&lon_min=-128&lat_min=25&lon_max=-65&lat_max=49&limit=50`

This generates a 5x5 grid over the United States.

The server returns a limited number of features by default, so we add a `limit` parameter in the call to ensure that we get all the grid cells. See _Limiting and Paging_ in [Executing Functions](/usage/query_function/) for more details on the `limit` parameter.

## Sample GeoJSON response

The function returns a feature collection of Polygons.

```json
{
  "type": "FeatureCollection",
  "features": [
    {
      "type": "Feature",
      "id": "1_1",
      "geometry": {
        "type": "Polygon",
        "coordinates": [
          [
            [
              -128,
              25
            ],
            [
              -128,
              29.8
            ],
            [
              -115.4,
              29.8
            ],
            [
              -115.4,
              25
            ],
            [
              -128,
              25
            ]
          ]
        ]
      },
      "properties": {
        "id": "1_1"
      }
    },
    ...
    {
      "type": "Feature",
      "id": "5_5",
      "geometry": {
        "type": "Polygon",
        "coordinates": [
          [
            [
              -77.6,
              44.2
            ],
            ...
            [
              -77.6,
              44.2
            ]
          ]
        ]
      },
      "properties": {
        "id": "5_5"
      }
    }
  ],
  "numberReturned": 25,
  "timeStamp": "2020-04-05T19:54:17Z",
  "links": [
    {
      "href": "http://localhost:9000/collections/geo_grid/items.json",
      "rel": "self",
      "type": "application/json",
      "title": "This document as JSON"
    },
    {
      "href": "http://localhost:9000/collections/geo_grid/items.html",
      "rel": "alternate",
      "type": "text/html",
      "title": "This document as HTML"
    }
  ]
}
```

Each cell has an `id` value that also indicates where it is on the grid. Since longitude and latitude values increase as you move east and north respectively, the cell with `id` 1_1 is the most southwestern corner of the grid, while cell 1_2 is immediately east and cell 2_1 immediately north.

## Web preview

![Geographic grid over the United States](/ex-query-grid.png)
