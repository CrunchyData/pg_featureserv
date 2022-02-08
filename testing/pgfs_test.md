# pg_fetaureserv Test Queries

## CRS handling

'''
http://localhost:9000/collections/pgfs_test.test_crs/items.json?crs=3005

http://localhost:9000/collections/pgfs_test.test_crs/items.json?crs=3005&bbox-crs=3005&bbox=1000000,400000,1010000,410000
Response: 1 feature

http://localhost:9000/collections/pgfs_test.test_crs/items.json?crs=3005&bbox-crs=3005&bbox=1000000,400000,1030000,430000
Response: 4 features
'''
