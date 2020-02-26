---
title: "Functions"
date:
draft: false
weight: 200
---

## Expose Database Functions



## List Functions

`/functions`
`/functions.json`

## Query Function metadata

`/functions/{funid}`
`/functions/{funid}.json`

## Query Function features or data

`/functions/{funid}/items`
`/functions/{funid}/items.json`

- response is GeoJSON for result dataset

### Function arguments

`param=value`

Omitted arguments will use the default specified in the function definition (if any).

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
