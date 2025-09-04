CREATE TABLE cities (
  id serial PRIMARY KEY,
  name text,
  geom geometry(Point, 4326)
);

INSERT INTO cities (name, geom) VALUES
  ('Paris', ST_SetSRID(ST_MakePoint(2.3522, 48.8566), 4326)),
  ('NYC',   ST_SetSRID(ST_MakePoint(-74.0060, 40.7128), 4326));

-- View with geometry and featureID column (no PK)
CREATE VIEW cities_view AS
  SELECT id AS id, name, geom FROM cities;

