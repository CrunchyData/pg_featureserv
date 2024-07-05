# pg_featureserv Test Requests

HTTP requests to test `pg_featureserv` capabilities.

### Usage

* Initialize the database with the DDL in `pgfs_test.sql`.
* Run `pg_featureserv`
* Paste queries into a browser

## CRS handling

```
http://localhost:9000/collections/pgfs_test.test_crs/items.json?crs=3005
```

```
http://localhost:9000/collections/pgfs_test.test_crs/items.json?crs=3005&bbox-crs=3005&bbox=1000000,400000,1010000,410000
```
Expected Response: 1 feature

```
http://localhost:9000/collections/pgfs_test.test_crs/items.json?crs=3005&bbox-crs=3005&bbox=1000000,400000,1030000,430000
```
Expected Response: 4 features


## CQL Filters
```
http://localhost:9000/collections/ne.admin_0_countries/items.json?properties=name,pop_est&filter=name%20BETWEEN%20%27Chile%27%20AND%20%27Den%27
```

```
http://localhost:9000/collections/ne.admin_0_countries/items.json?properties=name,pop_est&filter=name%20LIKE%20%27%25Ch%25%27
```

```
http://localhost:9000/collections/pgfs_test.test_crs/items.json?filter=NOT%20id%20In%20(1,2,3)
```

### CQL Spatial Operators
```
http://localhost:9000/collections/ne.admin_0_countries/items.json?properties=name,pop_est&filter=INTERSECTS(geom, POINT(1 1))
```

```
http://localhost:9000/collections/ne.admin_0_countries/items.json?properties=name,pop_est&filter=INTERSECTS(geom, LineString(1 1, 2 2))
```

```
http://localhost:9000/collections/ne.admin_0_countries/items.html?properties=name,pop_est&filter=crosses(geom,%20LineString(-5 35,48 -20))
```

### CQL Distance Operators
```
http://localhost:9000/collections/ne.admin_0_countries/items.html?properties=name,pop_est&filter=dwithin(geom,Point(0%2050),10)
```

```
http://localhost:9000/collections/public.geonames/items.html?filter=dwithin(geom,Point(-100 40),0.5)&limit=1000
```

```
http://localhost:9000/collections/public.geonames/items.html?filter=dwithin(geom,Linestring(-100%2040,-98%2042),0.1)&limit=1000
```
### CQL Temporal operators
```
http://localhost:9000/collections/public.trop_storm/items.html?filter=time_start%20between%202005-01-01%20and%20NOW()&limit=100
```

```
http://localhost:9000/collections/public.trop_storm/items.html?filter=time_start%20between%202005-01-01%20and%20NOW() AND intersects(geom, POLYGON ((-81.4067 30.8422, -79.6862 25.3781, -81.1609 24.7731, -83.9591 30.0292, -85.2258 29.6511, -87.5892 29.9914, -87.5514 31.0123, -81.4067 30.8422)) )&limit=100
```

### Filter-Crs
```
http://localhost:9000/collections/pgfs_test.test_crs/items.html?filter=DWITHIN(geom,POINT(-124.6 49.3),40000)&limit=100
```

```
http://localhost:9000/collections/pgfs_test.test_crs/items.html?filter=DWITHIN(geom,POINT(1000000 400000),60000)&filter-crs=3005&limit=100
```

```
http://localhost:9000/collections/pgfs_test.test_crs/items.html?filter=intersects(geom,%20ENVELOPE(1000000,400000,1100000,500000))&filter-crs=3005&limit=100
```

### Functions
```
http://localhost:9000/functions/postgisftw.countries_name/items.json?name_prefix=C&filter=continent%20ILIKE%20%27%25america%27
```

## Catalog handling

### Handle all column names
```
http://localhost:9000/collections/pgfs_test.test_names/items.json?properties=id,colCamelCase
```
