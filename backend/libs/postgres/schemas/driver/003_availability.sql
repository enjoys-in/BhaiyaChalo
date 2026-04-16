-- ============================================================
-- DRIVER SCHEMA: Availability & Availability Logs
-- Database: public | Prefix: driver_
-- Requires: postgis extension (see 000_init.sql)
-- Requires: 000_geo_regions.sql (regions/cities)
--
-- GEO-SHARDING: driver_availability_logs uses compound partition
--   Level 1: LIST (region_id)  → routes to regional shard
--   Level 2: RANGE (timestamp) → monthly time partitions
-- ============================================================

CREATE TYPE driver_action_type AS ENUM ('went_online', 'went_offline', 'trip_started', 'trip_ended');

CREATE TABLE IF NOT EXISTS public.driver_availability (
    driver_id    VARCHAR(36)      PRIMARY KEY,
    city_id      VARCHAR(36)      NOT NULL DEFAULT '',
    region_id    VARCHAR(36)      NOT NULL DEFAULT '',
    online       BOOLEAN          NOT NULL DEFAULT FALSE,
    on_trip      BOOLEAN          NOT NULL DEFAULT FALSE,
    vehicle_type VARCHAR(30)      NOT NULL DEFAULT '',
    lat          DOUBLE PRECISION NOT NULL DEFAULT 0,
    lng          DOUBLE PRECISION NOT NULL DEFAULT 0,
    -- PostGIS: auto-computed from lat/lng for spatial queries
    geom         GEOMETRY(Point, 4326) GENERATED ALWAYS AS (
                     ST_SetSRID(ST_MakePoint(lng, lat), 4326)
                 ) STORED,
    last_seen_at TIMESTAMPTZ      NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMPTZ      NOT NULL DEFAULT NOW(),
    CONSTRAINT chk_avail_lat CHECK (lat BETWEEN -90 AND 90),
    CONSTRAINT chk_avail_lng CHECK (lng BETWEEN -180 AND 180)
);

CREATE TRIGGER set_updated_at BEFORE UPDATE ON public.driver_availability
    FOR EACH ROW EXECUTE FUNCTION trigger_set_updated_at();

CREATE INDEX idx_driver_availability_city   ON public.driver_availability (city_id);
CREATE INDEX idx_driver_availability_region ON public.driver_availability (region_id);
CREATE INDEX idx_driver_availability_online ON public.driver_availability (online) WHERE online = TRUE;
CREATE INDEX idx_driver_availability_geom   ON public.driver_availability USING GIST (geom);
CREATE INDEX idx_driver_availability_comp   ON public.driver_availability (region_id, city_id, online, vehicle_type) WHERE online = TRUE;

-- Compound partitioned: LIST(region_id) → RANGE(timestamp)
CREATE TABLE IF NOT EXISTS public.driver_availability_logs (
    id        VARCHAR(36)        NOT NULL DEFAULT gen_random_uuid()::TEXT,
    driver_id VARCHAR(36)        NOT NULL,
    city_id   VARCHAR(36)        NOT NULL DEFAULT '',
    region_id VARCHAR(36)        NOT NULL,
    action    driver_action_type NOT NULL,
    timestamp TIMESTAMPTZ        NOT NULL DEFAULT NOW()
) PARTITION BY LIST (region_id);

-- Bootstrap region partitions (each sub-partitioned by timestamp monthly)
SELECT bootstrap_region_partitions('public.driver_availability_logs', 'timestamp');

CREATE INDEX idx_driver_availability_logs_driver ON public.driver_availability_logs (driver_id);
CREATE INDEX idx_driver_availability_logs_region ON public.driver_availability_logs (region_id);
CREATE INDEX idx_driver_availability_logs_time   ON public.driver_availability_logs (timestamp);
CREATE INDEX idx_driver_availability_logs_comp   ON public.driver_availability_logs (region_id, driver_id, action, timestamp);
