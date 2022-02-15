---
title: "Example: Function to get Country Name at Coordinate"
date:
draft: false
weight: 220
---

Non-spatial functions (i.e. functions that don't return spatial data) can also be accessed via `pg_featureserv`, as long as they are published in a configured schema
(by default, all functions in the `postgisftw` schema are published).

The following function example can be used with the `ne.countries` collection created in the [Quick Start](/quickstart/) section. It shows a function that accepts longitude and latitude values, and returns the corresponding country (if any). Unlike the other function examples in this section, it does not return a table with a geometry type column.

Any kind of function can be published, which allows you very flexible access to data. You can create functions that return statistics, summary records, populate dropdown lists or autocomplete suggestions, and more.

## Create a non-spatial function that locates the country at a coordinate

```sql
CREATE OR REPLACE FUNCTION postgisftw.country_by_loc(
  lon numeric DEFAULT 0.0,
  lat numeric DEFAULT 0.0)
RETURNS TABLE(name text, abbrev text, postal text)
AS $$
BEGIN
     RETURN QUERY
   SELECT c.name::text, c.abbrev::text, c.postal::text
   FROM ne.countries c
   WHERE ST_Intersects(c.geom,
            ST_SetSRID(ST_MakePoint(lon, lat), 4326))
   LIMIT 1;
END;
$$
LANGUAGE 'plpgsql' STABLE STRICT;

COMMENT ON FUNCTION postgisftw.country_by_loc
IS 'Finds the country at a geographic location';
```

Notes:

* The function generates a [Point](https://postgis.net/docs/ST_MakePoint.html) based on the longitude and latitude values provided in the parameters.
* The `ne.countries` table is filtered based on whether the point [intersects](https://postgis.net/docs/ST_Intersects.html) a country polygon.
* It's possible that a point lies exactly on the boundary between two countries. Both country records will be included in the query result set, but `LIMIT 1` restricts the result to a single record.

## Example of API query

The geodetic coordinate (47,8) is passed into the function:

`http://localhost:9000/functions/postgisftw.country_by_loc/items.json?lat=47&lon=8`

## Sample JSON response

The service returns data from non-spatial functions in JSON, instead of GeoJSON.

```json
[
    {
        "abbrev":"Switz.",
        "name":"Switzerland",
        "postal":"CH"
    }
]
```
