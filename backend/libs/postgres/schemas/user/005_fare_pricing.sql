-- ============================================================
-- USER SCHEMA: Fare Calculations & Price Estimates
-- Database: public | Prefix: user_
-- ============================================================

CREATE TABLE IF NOT EXISTS public.user_fare_calculations (
    id               VARCHAR(36)   PRIMARY KEY DEFAULT gen_random_uuid()::TEXT,
    booking_id       VARCHAR(36)   NOT NULL,
    base_price       NUMERIC(12,2) NOT NULL DEFAULT 0,
    distance_charge  NUMERIC(12,2) NOT NULL DEFAULT 0,
    time_charge      NUMERIC(12,2) NOT NULL DEFAULT 0,
    surge_multiplier DOUBLE PRECISION NOT NULL DEFAULT 1.0,
    surge_amount     NUMERIC(12,2) NOT NULL DEFAULT 0,
    toll_charges     NUMERIC(12,2) NOT NULL DEFAULT 0,
    tax_amount       NUMERIC(12,2) NOT NULL DEFAULT 0,
    promo_discount   NUMERIC(12,2) NOT NULL DEFAULT 0,
    total_fare       NUMERIC(12,2) NOT NULL DEFAULT 0,
    currency         VARCHAR(10)   NOT NULL DEFAULT 'INR',
    city_id          VARCHAR(36)   NOT NULL,
    region_id        VARCHAR(36)   NOT NULL DEFAULT '',
    vehicle_type     VARCHAR(30)   NOT NULL,
    created_at       TIMESTAMPTZ   NOT NULL DEFAULT NOW(),
    CONSTRAINT chk_fare_calc_base       CHECK (base_price      >= 0),
    CONSTRAINT chk_fare_calc_distance   CHECK (distance_charge >= 0),
    CONSTRAINT chk_fare_calc_time       CHECK (time_charge     >= 0),
    CONSTRAINT chk_fare_calc_surge_mult CHECK (surge_multiplier >= 1.0),
    CONSTRAINT chk_fare_calc_surge_amt  CHECK (surge_amount    >= 0),
    CONSTRAINT chk_fare_calc_toll       CHECK (toll_charges    >= 0),
    CONSTRAINT chk_fare_calc_tax        CHECK (tax_amount      >= 0),
    CONSTRAINT chk_fare_calc_promo      CHECK (promo_discount  >= 0),
    CONSTRAINT chk_fare_calc_total      CHECK (total_fare      >= 0)
);

CREATE INDEX idx_user_fare_calculations_booking ON public.user_fare_calculations (booking_id);
CREATE INDEX idx_user_fare_calculations_region ON public.user_fare_calculations (region_id);
CREATE INDEX idx_user_fare_calculations_city    ON public.user_fare_calculations (city_id, vehicle_type);

CREATE TABLE IF NOT EXISTS public.user_price_estimates (
    id               VARCHAR(36)      PRIMARY KEY DEFAULT gen_random_uuid()::TEXT,
    city_id          VARCHAR(36)      NOT NULL,
    region_id        VARCHAR(36)      NOT NULL DEFAULT '',
    vehicle_type     VARCHAR(30)      NOT NULL,
    distance_km      DOUBLE PRECISION NOT NULL DEFAULT 0,
    duration_min     DOUBLE PRECISION NOT NULL DEFAULT 0,
    base_fare        NUMERIC(12,2)    NOT NULL DEFAULT 0,
    surge_multiplier DOUBLE PRECISION NOT NULL DEFAULT 1.0,
    estimated_fare   NUMERIC(12,2)    NOT NULL DEFAULT 0,
    currency         VARCHAR(10)      NOT NULL DEFAULT 'INR',
    created_at       TIMESTAMPTZ      NOT NULL DEFAULT NOW(),
    CONSTRAINT chk_price_est_distance  CHECK (distance_km      >= 0),
    CONSTRAINT chk_price_est_duration  CHECK (duration_min     >= 0),
    CONSTRAINT chk_price_est_base      CHECK (base_fare        >= 0),
    CONSTRAINT chk_price_est_surge     CHECK (surge_multiplier >= 1.0),
    CONSTRAINT chk_price_est_total     CHECK (estimated_fare   >= 0)
);

CREATE INDEX idx_user_price_estimates_region ON public.user_price_estimates (region_id);
CREATE INDEX idx_user_price_estimates_city ON public.user_price_estimates (city_id, vehicle_type);
