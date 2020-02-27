---
title: "Functions"
date:
draft: false
weight: 200
---

<!-- ## Expose Database Functions -->

## List Functions

`/functions`

`/functions.json`

## Query Function Metadata

`/functions/{funid}`

`/functions/{funid}.json`

## Query Function Features or Data

`/functions/{funid}/items`

`/functions/{funid}/items.json`

- Response is GeoJSON for result dataset

### Function arguments

`param=value`

- Omitted arguments will use the default specified in the function definition (if any).

### Limiting and paging results

- `limit=N` specifies the maximum number of records returned. The default is 10.

- `offset=N` specifies the number of rows to skip before beginning to return rows.

### Specify properties in result

`properties=PROP1,PROP2,PROP3...`

### Ordering results

`orderBy=PROP`

- `orderBy=PROP:A` returns results based on the specified property in ascending order.

- `orderBy=PROP:D` returns results based on the specified property in descending order.

### Filter by bbox

`bbox=MINX,MINY,MAXX,MAXY`

- Extent is in lon/lat (4326)
