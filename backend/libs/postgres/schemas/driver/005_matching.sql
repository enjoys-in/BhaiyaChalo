-- ============================================================
-- DRIVER SCHEMA: Match Requests
-- Database: public | Prefix: driver_
-- Requires: postgis extension (see 000_init.sql)
-- Requires: 000_geo_regions.sql (regions/cities)
--
-- GEO-SHARDING: region_id for shard routing.
-- Matching is inherently geo-local — only search nearby drivers.
-- ============================================================

CREATE TYPE driver_match_status AS ENUM ('searching', 'matched', 'failed');

CREATE TABLE IF NOT EXISTS public.driver_match_requests (
    id           VARCHAR(36)         PRIMARY KEY DEFAULT gen_random_uuid()::TEXT,
    booking_id   VARCHAR(36)         NOT NULL,
    city_id      VARCHAR(36)         NOT NULL,
    region_id    VARCHAR(36)         NOT NULL,
    pickup_lat   DOUBLE PRECISION    NOT NULL,
    pickup_lng   DOUBLE PRECISION    NOT NULL,
    pickup_geom  GEOMETRY(Point, 4326) GENERATED ALWAYS AS (
                     ST_SetSRID(ST_MakePoint(pickup_lng, pickup_lat), 4326)
                 ) STORED,
    vehicle_type VARCHAR(30)         NOT NULL,
    radius_km    DOUBLE PRECISION    NOT NULL DEFAULT 5.0,
    status       driver_match_status NOT NULL DEFAULT 'searching',
    created_at   TIMESTAMPTZ         NOT NULL DEFAULT NOW(),
    CONSTRAINT chk_match_req_radius     CHECK (radius_km > 0),
    CONSTRAINT chk_match_req_pickup_lat CHECK (pickup_lat BETWEEN -90 AND 90),
    CONSTRAINT chk_match_req_pickup_lng CHECK (pickup_lng BETWEEN -180 AND 180)
);

CREATE INDEX idx_driver_match_requests_booking ON public.driver_match_requests (booking_id);
CREATE INDEX idx_driver_match_requests_region  ON public.driver_match_requests (region_id);
CREATE INDEX idx_driver_match_requests_city    ON public.driver_match_requests (city_id);
CREATE INDEX idx_driver_match_requests_status  ON public.driver_match_requests (status);
CREATE INDEX idx_driver_match_requests_geom    ON public.driver_match_requests USING GIST (pickup_geom);
CREATE INDEX idx_driver_match_requests_comp    ON public.driver_match_requests (region_id, city_id, vehicle_type, status);
