---
title: "Example: Generate a Geographic Grid"
date:
draft: false
weight: 220
---

This example shows how to generate geometry data from a function, controlled by some input parameters. 

This particular function does not query an existing table in the database; rather, it uses PostGIS functions to generate spatial data.

Scenarios in which this type of spatial function comes in handy are:

* Data visualization
* Data analysis

### Create a spatial function that generates a grid over a desired area 

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

