CREATE EXTENSION IF NOT EXISTS postgis;
ALTER SYSTEM SET log_statement = 'all';
SELECT pg_reload_conf();
