---
title: "Tables and Views"
date:
draft: false
weight: 100
---

## Expose Database tables as Feature Collections



## List Feature Collections

`/collections`

## Query Feature Collection metadata

`/collections/{collid}`

## Query Feature Collection features

`/collections/{collid}/items`

- response is GeoJSON

### Limiting and paging results

`limit`
`offset`

### Ordering results

`orderBy=PROP`
`orderBy=PROP:A` or `orderBy=PROP:D`

### Filter by bbox

`bbox=MINX,MINY,MAXX,MAXY`

- lon/lat

### Specify properties in result

`properties=PROP1,PROP2,PROP3...`

## Query single feature

`/collections/{collid}/items/{fid}`

### Specify properties in result
