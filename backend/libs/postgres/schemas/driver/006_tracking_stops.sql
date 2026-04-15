-- ============================================================
-- DRIVER SCHEMA: Tracking Sessions & Stops
-- Database: public | Prefix: driver_
-- Requires: postgis extension (see 000_init.sql)
-- ============================================================

CREATE TABLE IF NOT EXISTS public.driver_tracking_sessions (
    id         VARCHAR(36) PRIMARY KEY DEFAULT gen_random_uuid()::TEXT,
    trip_id    VARCHAR(36) NOT NULL,
    driver_id  VARCHAR(36) NOT NULL,
    user_id    VARCHAR(36) NOT NULL,
    active     BOOLEAN     NOT NULL DEFAULT TRUE,
    started_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    ended_at   TIMESTAMPTZ
);

CREATE INDEX idx_driver_tracking_sessions_trip      ON public.driver_tracking_sessions (trip_id);
CREATE INDEX idx_driver_tracking_sessions_driver    ON public.driver_tracking_sessions (driver_id);
CREATE INDEX idx_driver_tracking_sessions_active    ON public.driver_tracking_sessions (active) WHERE active = TRUE;
CREATE INDEX idx_driver_tracking_sessions_composite ON public.driver_tracking_sessions (driver_id, active, started_at);

-- Multi-stop trips

CREATE TYPE driver_stop_status AS ENUM ('pending', 'arrived', 'completed', 'skipped');

CREATE TABLE IF NOT EXISTS public.driver_stops (
    id          VARCHAR(36)        PRIMARY KEY DEFAULT gen_random_uuid()::TEXT,
    trip_id     VARCHAR(36)        NOT NULL,
    lat         DOUBLE PRECISION   NOT NULL,
    lng         DOUBLE PRECISION   NOT NULL,
    -- PostGIS: auto-computed from lat/lng for spatial queries
    geom        GEOMETRY(Point, 4326) GENERATED ALWAYS AS (
                    ST_SetSRID(ST_MakePoint(lng, lat), 4326)
                ) STORED,
    address     TEXT               NOT NULL DEFAULT '',
    stop_order  INT                NOT NULL DEFAULT 0,
    status      driver_stop_status NOT NULL DEFAULT 'pending',
    arrived_at  TIMESTAMPTZ,
    departed_at TIMESTAMPTZ,
    created_at  TIMESTAMPTZ        NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ        NOT NULL DEFAULT NOW(),
    CONSTRAINT chk_stop_lat   CHECK (lat BETWEEN -90 AND 90),
    CONSTRAINT chk_stop_lng   CHECK (lng BETWEEN -180 AND 180),
    CONSTRAINT chk_stop_order CHECK (stop_order >= 0)
);

CREATE TRIGGER set_updated_at BEFORE UPDATE ON public.driver_stops
    FOR EACH ROW EXECUTE FUNCTION trigger_set_updated_at();

CREATE INDEX idx_driver_stops_trip  ON public.driver_stops (trip_id);
CREATE INDEX idx_driver_stops_order ON public.driver_stops (trip_id, stop_order);
CREATE INDEX idx_driver_stops_geom  ON public.driver_stops USING GIST (geom);
