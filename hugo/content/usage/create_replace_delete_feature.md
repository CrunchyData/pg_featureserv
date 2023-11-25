---
title: "Create Replace Delete Feature"
date:
draft: false
weight: 150
---

Transaction on Feature collections is supported.

## Create feature

POST query to the path `/collections/{collid}/items` allows to create
a new feature in a feature collection.

The geojson feature must be part of the request body. 
If the geometry geometry crs is different from the storage crs, the geometry will be transformed.
Missing properties will be ignored and the table default value for the column will be applied.
The id specified in the body is ignored and the database default value is used to create the feature.

#### Example
```
curl -i --request "POST" 'http://localhost:9000/collections/public.tramway_stations/items' -d '{"type":"Feature","id":"129","geometry":{"type":"Point","coordinates":[-71.222868058,46.836016945,0]},"properties":{"description":null,"diffusion":"Publique","niveau_rstc":"Tramway","nom":"Hôpital Enfant-Jésus","objectid":129,"type_station":"Reguliere"}}'
```

## Replace feature

PUT query to the path `/collections/{collid}/items/{fid}` allows to replace
a feature in a feature collection.

The geojson feature must be part of the request body. 
If the geometry geometry crs is different from the storage crs, the geometry will be transformed.
Missing properties will be replaced with null (unless a database trigger is applied)
The id specified in the body is ignored.

#### Example
```
curl -i --request "PUT" 'http://localhost:9000/collections/public.tramway_stations/items/129.json' -d '{"type":"Feature","id":"129","geometry":{"type":"Point","coordinates":[-71.222868058,46.836016945,0]},"properties":{"description":null,"diffusion":"Publique","niveau_rstc":"Tramway","nom":"Hôpital Enfant-Jésus","objectid":129,"type_station":"Reguliere"}}'
```

## Delete feature

DELETE query to the path `/collections/{collid}/items/{fid}` allows to delete
a feature in a feature collection.

#### Example
```
curl -i --request "Delete" 'http://localhost:9000/collections/public.tramway_stations/items/129.json'
```

