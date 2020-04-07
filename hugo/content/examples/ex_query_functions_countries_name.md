---
title: "Example: Query Countries by Name"
date:
draft: false
weight: 210
---

This is the same spatial function example shown in the [Usage](/usage/functions/) section, but we'll show a sample GeoJSON response, as well as the web UI preview. 

### Create a spatial function that returns a filtered set of countries 

```sql
CREATE OR REPLACE FUNCTION postgisftw.countries_name(
	name_prefix text DEFAULT 'A')
RETURNS TABLE(name text, abbrev text, continent text, geom geometry)
AS $$
BEGIN
	RETURN QUERY
		SELECT t.name::text,
            t.abbrev::text,
            t.continent::text,
            t.geom
    FROM ne.admin_0_countries t
    WHERE t.name ILIKE name_prefix || '%';
END;
$$
LANGUAGE 'plpgsql' STABLE PARALLEL SAFE;

COMMENT ON FUNCTION postgisftw.countries_name IS 'Filters the countries table by the initial letters of the name using the "name_prefix" parameter.';
```

### Example of API query

The function can be called via the API by providing a value for the `name_prefix` parameter. 

```http://localhost:9000/functions/countries_name/items?name_prefix=Mo```

Since a default value is included in the function declaration, you could omit the parameter in the call -- a random sample of features will be returned.

### Sample GeoJSON response

The response is a GeoJSON document containing the 7 countries starting with the letters 'Mo'.

```json
{
      "type":"FeatureCollection",
      "features":[
         {
           "type":"Feature",
           "geometry":{
              "type":"MultiPolygon",
              "coordinates":[
                 [
                     [
                        [
                           -62.1484375,
                           16.74033203125
                        ],
                        [
                          -62.154248046875,
                          16.681201171875
                       ],
                      ...
                      [
                        -62.1484375,
                        16.74033203125
                      ]
                  ]
              ]
          ]
     },
     "properties":{
         "abbrev":"Monts.",
         "continent":"North America",
         "name":"Montserrat"
     }
  },
  ...
 ],
 "numberReturned":7,
 "timeStamp":"2020-03-18T03:15:15Z",
 "links":[
 {
"href":"http://localhost:9000/collections/countries_name/items.json",
          "rel":"self",
          "type":"application/json",
          "title":"This document as JSON"
       },
       {
"href":"http://localhost:9000/collections/countries_name/items.html",
          "rel":"alternate",
          "type":"text/html",
          "title":"This document as HTML"
      }
  ]
}
```

### Web preview

![Countries starting with 'Mo'](/ex-query-countries.png)