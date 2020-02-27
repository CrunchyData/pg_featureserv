---
title: "Querying Features"
date:
draft: false
weight: 150
---

## Query Features

`/collections/{collid}/items`

- response is GeoJSON for result dataset

### Limiting and paging results

`limit=N`

`offset=N`

### Ordering results

`orderBy=PROP`

`orderBy=PROP:A`

`orderBy=PROP:D`

### Filter by bbox

`bbox=MINX,MINY,MAXX,MAXY`

- extent is in lon/lat (4326)

### Specify properties in result

`properties=PROP1,PROP2,PROP3...`

## Query a single feature

`/collections/{collid}/items/{fid}`

### Specify properties in result

`properties=PROP1,PROP2,PROP3...`
