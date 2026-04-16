-- ============================================================
-- ADMIN SCHEMA: Geofence Management
-- Database: public | Prefix: admin_
-- Requires: postgis extension (see 000_init.sql)
-- ============================================================

CREATE TYPE admin_fence_type AS ENUM ('city_boundary', 'zone', 'airport', 'restricted', 'surge_zone');

CREATE TABLE IF NOT EXISTS public.admin_geofences (
    id         VARCHAR(36)      PRIMARY KEY DEFAULT gen_random_uuid()::TEXT,
    city_id    VARCHAR(36)      NOT NULL,
    region_id  VARCHAR(36)      NOT NULL DEFAULT '',
    name       VARCHAR(200)     NOT NULL,
    type       admin_fence_type NOT NULL,
    polygon    JSONB            NOT NULL DEFAULT '[]',
    center_lat DOUBLE PRECISION NOT NULL DEFAULT 0,
    center_lng DOUBLE PRECISION NOT NULL DEFAULT 0,
    radius_km  DOUBLE PRECISION NOT NULL DEFAULT 0,
    -- PostGIS: auto-computed from center_lat/center_lng for spatial queries
    geom       GEOMETRY(Point, 4326) GENERATED ALWAYS AS (
                   ST_SetSRID(ST_MakePoint(center_lng, center_lat), 4326)
               ) STORED,
    active     BOOLEAN          NOT NULL DEFAULT TRUE,
    metadata   JSONB            NOT NULL DEFAULT '{}',
    created_at TIMESTAMPTZ      NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ      NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ,
    CONSTRAINT chk_admin_geofences_name      CHECK (length(name) >= 1),
    CONSTRAINT chk_admin_geofences_radius    CHECK (radius_km >= 0),
    CONSTRAINT chk_admin_geofences_lat_range CHECK (center_lat BETWEEN -90 AND 90),
    CONSTRAINT chk_admin_geofences_lng_range CHECK (center_lng BETWEEN -180 AND 180)
);

CREATE TRIGGER set_updated_at BEFORE UPDATE ON public.admin_geofences
    FOR EACH ROW EXECUTE FUNCTION trigger_set_updated_at();

CREATE INDEX idx_admin_geofences_region   ON public.admin_geofences (region_id);
CREATE INDEX idx_admin_geofences_city     ON public.admin_geofences (city_id);
CREATE INDEX idx_admin_geofences_type     ON public.admin_geofences (type);
CREATE INDEX idx_admin_geofences_active   ON public.admin_geofences (active) WHERE active = TRUE;
CREATE INDEX idx_admin_geofences_geom     ON public.admin_geofences USING GIST (geom);
CREATE INDEX idx_admin_geofences_metadata ON public.admin_geofences USING GIN (metadata);
CREATE INDEX idx_admin_geofences_deleted  ON public.admin_geofences (deleted_at) WHERE deleted_at IS NULL;
