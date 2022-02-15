# pg_fetaureserv Test Queries

## CRS handling

'''
http://localhost:9000/collections/pgfs_test.test_crs/items.json?crs=3005

http://localhost:9000/collections/pgfs_test.test_crs/items.json?crs=3005&bbox-crs=3005&bbox=1000000,400000,1010000,410000
Response: 1 feature

http://localhost:9000/collections/pgfs_test.test_crs/items.json?crs=3005&bbox-crs=3005&bbox=1000000,400000,1030000,430000
Response: 4 features
'''

## Filter CQL
```
http://localhost:9000/collections/ne.admin_0_countries/items.json?properties=name,pop_est&filter=name%20BETWEEN%20%27Chile%27%20AND%20%27Den%27

http://localhost:9000/collections/ne.admin_0_countries/items.json?properties=name,pop_est&filter=name%20LIKE%20%27%25Ch%25%27

http://localhost:9000/collections/pgfs_test.test_crs/items.json?filter=NOT%20id%20In%20(1,2,3)
```

### Functions
```
http://localhost:9000/functions/postgisftw.countries_name/items.json?name_prefix=C&filter=continent%20ILIKE%20%27%25america%27
```
