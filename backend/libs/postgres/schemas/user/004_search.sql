-- ============================================================
-- USER SCHEMA: Search Queries & Results
-- Database: public | Prefix: user_
-- Requires: postgis extension (see 000_init.sql)
-- ============================================================

CREATE TABLE IF NOT EXISTS public.user_search_queries (
    id         VARCHAR(36)      PRIMARY KEY DEFAULT gen_random_uuid()::TEXT,
    user_id    VARCHAR(36)      NOT NULL,
    city_id    VARCHAR(36)      NOT NULL,
    pickup_lat DOUBLE PRECISION NOT NULL,
    pickup_lng DOUBLE PRECISION NOT NULL,
    pickup_geom GEOMETRY(Point, 4326) GENERATED ALWAYS AS (
                    ST_SetSRID(ST_MakePoint(pickup_lng, pickup_lat), 4326)
                ) STORED,
    drop_lat   DOUBLE PRECISION NOT NULL,
    drop_lng   DOUBLE PRECISION NOT NULL,
    drop_geom  GEOMETRY(Point, 4326) GENERATED ALWAYS AS (
                    ST_SetSRID(ST_MakePoint(drop_lng, drop_lat), 4326)
                ) STORED,
    created_at TIMESTAMPTZ      NOT NULL DEFAULT NOW(),
    CONSTRAINT chk_search_pickup_lat CHECK (pickup_lat BETWEEN -90 AND 90),
    CONSTRAINT chk_search_pickup_lng CHECK (pickup_lng BETWEEN -180 AND 180),
    CONSTRAINT chk_search_drop_lat   CHECK (drop_lat BETWEEN -90 AND 90),
    CONSTRAINT chk_search_drop_lng   CHECK (drop_lng BETWEEN -180 AND 180)
);

CREATE INDEX idx_user_search_queries_user   ON public.user_search_queries (user_id);
CREATE INDEX idx_user_search_queries_city   ON public.user_search_queries (city_id);
CREATE INDEX idx_user_search_queries_pickup ON public.user_search_queries USING GIST (pickup_geom);
CREATE INDEX idx_user_search_queries_drop   ON public.user_search_queries USING GIST (drop_geom);
CREATE INDEX idx_user_search_queries_comp   ON public.user_search_queries (user_id, city_id, created_at);

CREATE TABLE IF NOT EXISTS public.user_search_results (
    query_id          VARCHAR(36)   NOT NULL,
    vehicle_type      VARCHAR(30)   NOT NULL,
    estimated_fare    NUMERIC(12,2) NOT NULL DEFAULT 0,
    eta_minutes       INT           NOT NULL DEFAULT 0,
    available_drivers INT           NOT NULL DEFAULT 0,
    surge_multiplier  DOUBLE PRECISION NOT NULL DEFAULT 1.0,
    PRIMARY KEY (query_id, vehicle_type),
    CONSTRAINT chk_search_res_fare  CHECK (estimated_fare    >= 0),
    CONSTRAINT chk_search_res_eta   CHECK (eta_minutes       >= 0),
    CONSTRAINT chk_search_res_avail CHECK (available_drivers  >= 0),
    CONSTRAINT chk_search_res_surge CHECK (surge_multiplier   >= 1.0)
);
