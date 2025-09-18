CREATE TABLE cities (
  name text PRIMARY KEY,
  geom geometry(Point, 4326)
);

CREATE TABLE trips (
  id serial PRIMARY KEY,
  city text REFERENCES cities(name),
  time timestamptz,
  start_time timestamptz,
  end_time timestamptz
);

CREATE TABLE receipts (
  id serial PRIMARY KEY,
  trip_id int REFERENCES trips(id),
  time timestamptz,
  amount numeric
);


INSERT INTO cities (name, geom) VALUES
  ('Paris', ST_SetSRID(ST_MakePoint(2.3522, 48.8566), 4326)),
  ('London', ST_SetSRID(ST_MakePoint(-0.1276, 51.5074), 4326)),
  ('Tokyo', ST_SetSRID(ST_MakePoint(139.6917, 35.6895), 4326)),
  ('Sydney', ST_SetSRID(ST_MakePoint(151.2093, -33.8688), 4326)),
  ('NYC',   ST_SetSRID(ST_MakePoint(-74.0060, 40.7128), 4326));



INSERT INTO trips (city, time, start_time, end_time) VALUES
  ('Paris', '2025-01-01 12:00:00', '2024-01-01 12:00:00', '2025-01-01 12:00:00'),
  ('London', '2025-02-01 12:00:00', '2024-02-01 12:00:00', '2025-02-01 12:00:00'),
  ('Tokyo', '2025-03-01 12:00:00', '2024-03-01 12:00:00', '2025-03-01 12:00:00'),
  ('Sydney', '2025-04-01 12:00:00', '2024-04-01 12:00:00', '2025-04-01 12:00:00'),
  ('NYC',   '2025-05-01 12:00:00', '2024-05-01 12:00:00', '2025-05-01 12:00:00');


INSERT INTO receipts (trip_id, time, amount) VALUES
  (1, '2024-06-01 12:00:00', 100.00),
  (1, '2024-07-01 12:00:00', 150.00),
  (2, '2024-06-15 12:00:00', 200.00),
  (3, '2024-08-01 12:00:00', 250.00),
  (4, '2024-09-01 12:00:00', 300.00),
  (5, '2024-10-01 12:00:00', 350.00);

-- View with geometry and featureID column (no PK)
CREATE VIEW cities_view AS
  SELECT * FROM cities;

CREATE VIEW trips_view AS
  SELECT trips.*, cities.geom FROM trips LEFT JOIN cities ON trips.city = cities.name;

CREATE VIEW receipts_view AS
  SELECT receipts.*, cities.geom FROM receipts LEFT JOIN trips ON receipts.trip_id = trips.id LEFT JOIN cities ON trips.city = cities.name;
