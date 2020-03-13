
wget http://download.osgeo.org/livedvd/data/osm/Boston_MA/Boston_MA.osm.bz2



osm2pgrouting --password centos --host 127.0.0.1 --username centos --dbname routing --file Boston_MA.osm


