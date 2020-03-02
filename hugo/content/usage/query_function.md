---
title: "Querying Functions"
date:
draft: false
weight: 250
---

## Query Function Features or Data

`/functions/{funid}/items`


- response is GeoJSON for result dataset

### Function arguments

`param=value`

Omitted arguments will use the default specified in the function definition (if any).

### Filter by bbox

`bbox=MINX,MINY,MAXX,MAXY`

- extent is in lon/lat (4326)

### Specify properties in result

`properties=PROP1,PROP2,PROP3...`

### Limiting and paging results

`limit=N`

`offset=N`

### Ordering results

`orderBy=PROP`

`orderBy=PROP:A`

`orderBy=PROP:D`
