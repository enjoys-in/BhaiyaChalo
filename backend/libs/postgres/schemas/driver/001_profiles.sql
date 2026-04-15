-- ============================================================
-- DRIVER SCHEMA: Driver Profiles, Locations, Preferences
-- Database: public | Prefix: driver_
-- Requires: postgis extension (see 000_init.sql)
-- ============================================================

CREATE TYPE driver_profile_status AS ENUM ('pending', 'active', 'inactive', 'suspended', 'blocked');

CREATE TABLE IF NOT EXISTS public.driver_profiles (
    id              VARCHAR(36)           PRIMARY KEY DEFAULT gen_random_uuid()::TEXT,
    first_name      VARCHAR(100)          NOT NULL DEFAULT '',
    last_name       VARCHAR(100)          NOT NULL DEFAULT '',
    phone           VARCHAR(20)           NOT NULL UNIQUE,
    email           VARCHAR(255)          NOT NULL DEFAULT '',
    avatar_url      TEXT                  NOT NULL DEFAULT '',
    license_number  VARCHAR(50)           NOT NULL DEFAULT '',
    city_id         VARCHAR(36)           NOT NULL DEFAULT '',
    rating          DOUBLE PRECISION      NOT NULL DEFAULT 0,
    total_trips     INT                   NOT NULL DEFAULT 0,
    status          driver_profile_status NOT NULL DEFAULT 'pending',
    onboarding_step INT                   NOT NULL DEFAULT 0,
    created_at      TIMESTAMPTZ           NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ           NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ,
    CONSTRAINT chk_driver_phone       CHECK (length(phone) >= 10),
    CONSTRAINT chk_driver_rating      CHECK (rating BETWEEN 0 AND 5.0),
    CONSTRAINT chk_driver_total_trips CHECK (total_trips >= 0),
    CONSTRAINT chk_driver_onboarding  CHECK (onboarding_step >= 0)
);

CREATE TRIGGER set_updated_at BEFORE UPDATE ON public.driver_profiles
    FOR EACH ROW EXECUTE FUNCTION trigger_set_updated_at();

CREATE INDEX idx_driver_profiles_phone   ON public.driver_profiles (phone);
CREATE INDEX idx_driver_profiles_city    ON public.driver_profiles (city_id);
CREATE INDEX idx_driver_profiles_status  ON public.driver_profiles (status);
CREATE INDEX idx_driver_profiles_rating  ON public.driver_profiles (rating);
CREATE INDEX idx_driver_profiles_deleted ON public.driver_profiles (deleted_at) WHERE deleted_at IS NULL;
CREATE INDEX idx_driver_profiles_comp    ON public.driver_profiles (city_id, status, rating DESC);

CREATE TABLE IF NOT EXISTS public.driver_locations (
    driver_id  VARCHAR(36)      PRIMARY KEY,
    lat        DOUBLE PRECISION NOT NULL DEFAULT 0,
    lng        DOUBLE PRECISION NOT NULL DEFAULT 0,
    -- PostGIS: auto-computed from lat/lng for spatial queries
    geom       GEOMETRY(Point, 4326) GENERATED ALWAYS AS (
                   ST_SetSRID(ST_MakePoint(lng, lat), 4326)
               ) STORED,
    heading    DOUBLE PRECISION NOT NULL DEFAULT 0,
    speed      DOUBLE PRECISION NOT NULL DEFAULT 0,
    updated_at TIMESTAMPTZ      NOT NULL DEFAULT NOW(),
    CONSTRAINT chk_driver_loc_lat     CHECK (lat BETWEEN -90 AND 90),
    CONSTRAINT chk_driver_loc_lng     CHECK (lng BETWEEN -180 AND 180),
    CONSTRAINT chk_driver_loc_heading CHECK (heading BETWEEN 0 AND 360),
    CONSTRAINT chk_driver_loc_speed   CHECK (speed >= 0)
);

CREATE TRIGGER set_updated_at BEFORE UPDATE ON public.driver_locations
    FOR EACH ROW EXECUTE FUNCTION trigger_set_updated_at();

CREATE INDEX idx_driver_locations_geom ON public.driver_locations USING GIST (geom);

CREATE TABLE IF NOT EXISTS public.driver_preferences (
    driver_id       VARCHAR(36)      PRIMARY KEY,
    auto_accept     BOOLEAN          NOT NULL DEFAULT FALSE,
    max_distance    DOUBLE PRECISION NOT NULL DEFAULT 10.0,
    preferred_zones JSONB            NOT NULL DEFAULT '[]',
    updated_at      TIMESTAMPTZ      NOT NULL DEFAULT NOW(),
    CONSTRAINT chk_driver_pref_dist CHECK (max_distance > 0)
);

CREATE TRIGGER set_updated_at BEFORE UPDATE ON public.driver_preferences
    FOR EACH ROW EXECUTE FUNCTION trigger_set_updated_at();

CREATE INDEX idx_driver_preferences_zones ON public.driver_preferences USING GIN (preferred_zones);
