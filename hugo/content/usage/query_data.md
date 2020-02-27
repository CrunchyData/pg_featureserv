---
title: "Querying Features"
date:
draft: false
weight: 150
---

## Query features

`/collections/{collid}/items`

- Response is GeoJSON for result dataset.

### Limiting and paging results

- `limit=N` specifies the maximum number of records returned. The default is 10.

- `offset=N` specifies the number of rows to skip before beginning to return rows.

### Specify result properties

`properties=PROP1,PROP2,PROP3...`

### Ordering results

`orderBy=PROP`

- `orderBy=PROP:A` returns results based on the specified property in ascending order.

- `orderBy=PROP:D` returns results based on the specified property in descending order.

### Filter by bbox

`bbox=MINX,MINY,MAXX,MAXY`

- Extent is in lon/lat (4326).

## Query a single feature

`/collections/{collid}/items/{fid}`

### Specify properties in result

`properties=PROP1,PROP2,PROP3...`
