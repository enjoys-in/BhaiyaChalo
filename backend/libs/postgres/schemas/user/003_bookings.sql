-- ============================================================
-- USER SCHEMA: Bookings
-- Database: public | Prefix: user_
-- Requires: postgis extension (see 000_init.sql)
-- ============================================================

CREATE TYPE user_booking_status AS ENUM (
    'pending', 'confirmed', 'driver_assigned',
    'in_progress', 'completed', 'cancelled'
);

-- Partitioned by created_at for time-series queries & retention
CREATE TABLE IF NOT EXISTS public.user_bookings (
    id              VARCHAR(36)         NOT NULL DEFAULT gen_random_uuid()::TEXT,
    user_id         VARCHAR(36)         NOT NULL,
    city_id         VARCHAR(36)         NOT NULL,
    pickup_lat      DOUBLE PRECISION    NOT NULL,
    pickup_lng      DOUBLE PRECISION    NOT NULL,
    pickup_address  TEXT                NOT NULL DEFAULT '',
    pickup_geom     GEOMETRY(Point, 4326) GENERATED ALWAYS AS (
                        ST_SetSRID(ST_MakePoint(pickup_lng, pickup_lat), 4326)
                    ) STORED,
    drop_lat        DOUBLE PRECISION    NOT NULL,
    drop_lng        DOUBLE PRECISION    NOT NULL,
    drop_address    TEXT                NOT NULL DEFAULT '',
    drop_geom       GEOMETRY(Point, 4326) GENERATED ALWAYS AS (
                        ST_SetSRID(ST_MakePoint(drop_lng, drop_lat), 4326)
                    ) STORED,
    vehicle_type    VARCHAR(30)         NOT NULL,
    estimated_fare  NUMERIC(12,2)       NOT NULL DEFAULT 0,
    final_fare      NUMERIC(12,2)       NOT NULL DEFAULT 0,
    promo_code      VARCHAR(50)         NOT NULL DEFAULT '',
    discount_amount NUMERIC(12,2)       NOT NULL DEFAULT 0,
    status          user_booking_status NOT NULL DEFAULT 'pending',
    driver_id       VARCHAR(36)         NOT NULL DEFAULT '',
    payment_method  VARCHAR(30)         NOT NULL DEFAULT 'cash',
    created_at      TIMESTAMPTZ         NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ         NOT NULL DEFAULT NOW(),
    cancelled_at    TIMESTAMPTZ,
    cancel_reason   TEXT,
    CONSTRAINT chk_booking_est_fare   CHECK (estimated_fare  >= 0),
    CONSTRAINT chk_booking_final_fare CHECK (final_fare      >= 0),
    CONSTRAINT chk_booking_discount   CHECK (discount_amount >= 0),
    CONSTRAINT chk_booking_pickup_lat CHECK (pickup_lat BETWEEN -90 AND 90),
    CONSTRAINT chk_booking_pickup_lng CHECK (pickup_lng BETWEEN -180 AND 180),
    CONSTRAINT chk_booking_drop_lat   CHECK (drop_lat BETWEEN -90 AND 90),
    CONSTRAINT chk_booking_drop_lng   CHECK (drop_lng BETWEEN -180 AND 180)
) PARTITION BY RANGE (created_at);

CREATE TABLE public.user_bookings_default
    PARTITION OF public.user_bookings DEFAULT;

CREATE TRIGGER set_updated_at BEFORE UPDATE ON public.user_bookings
    FOR EACH ROW EXECUTE FUNCTION trigger_set_updated_at();

CREATE INDEX idx_user_bookings_user      ON public.user_bookings (user_id);
CREATE INDEX idx_user_bookings_driver    ON public.user_bookings (driver_id) WHERE driver_id != '';
CREATE INDEX idx_user_bookings_status    ON public.user_bookings (status);
CREATE INDEX idx_user_bookings_city      ON public.user_bookings (city_id);
CREATE INDEX idx_user_bookings_created   ON public.user_bookings (created_at);
CREATE INDEX idx_user_bookings_composite ON public.user_bookings (city_id, status, created_at);
CREATE INDEX idx_user_bookings_pickup    ON public.user_bookings USING GIST (pickup_geom);
CREATE INDEX idx_user_bookings_drop      ON public.user_bookings USING GIST (drop_geom);
