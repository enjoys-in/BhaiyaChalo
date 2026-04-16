-- ============================================================
-- ADMIN SCHEMA: Fare & Pricing Configuration
-- Database: public | Prefix: admin_
-- ============================================================

CREATE TABLE IF NOT EXISTS public.admin_fare_configs (
    city_id            VARCHAR(36)  NOT NULL,
    region_id          VARCHAR(36)  NOT NULL DEFAULT '',
    vehicle_type       VARCHAR(30)  NOT NULL,
    base_price_per_km  NUMERIC(12,2) NOT NULL DEFAULT 0,
    base_price_per_min NUMERIC(12,2) NOT NULL DEFAULT 0,
    min_fare           NUMERIC(12,2) NOT NULL DEFAULT 0,
    cancellation_fee   NUMERIC(12,2) NOT NULL DEFAULT 0,
    created_at         TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at         TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    PRIMARY KEY (city_id, vehicle_type),
    CONSTRAINT chk_fare_cfg_base_km  CHECK (base_price_per_km  >= 0),
    CONSTRAINT chk_fare_cfg_base_min CHECK (base_price_per_min >= 0),
    CONSTRAINT chk_fare_cfg_min_fare CHECK (min_fare           >= 0),
    CONSTRAINT chk_fare_cfg_cancel   CHECK (cancellation_fee   >= 0)
);

CREATE TRIGGER set_updated_at BEFORE UPDATE ON public.admin_fare_configs
    FOR EACH ROW EXECUTE FUNCTION trigger_set_updated_at();

CREATE TABLE IF NOT EXISTS public.admin_pricing_rules (
    id                VARCHAR(36)   PRIMARY KEY DEFAULT gen_random_uuid()::TEXT,
    city_id           VARCHAR(36)   NOT NULL,
    region_id         VARCHAR(36)   NOT NULL DEFAULT '',
    vehicle_type      VARCHAR(30)   NOT NULL,
    base_fare_per_km  NUMERIC(12,2) NOT NULL DEFAULT 0,
    base_fare_per_min NUMERIC(12,2) NOT NULL DEFAULT 0,
    min_fare          NUMERIC(12,2) NOT NULL DEFAULT 0,
    max_fare          NUMERIC(12,2) NOT NULL DEFAULT 0,
    booking_fee       NUMERIC(12,2) NOT NULL DEFAULT 0,
    active            BOOLEAN       NOT NULL DEFAULT TRUE,
    created_at        TIMESTAMPTZ   NOT NULL DEFAULT NOW(),
    updated_at        TIMESTAMPTZ   NOT NULL DEFAULT NOW(),
    CONSTRAINT chk_pricing_base_km  CHECK (base_fare_per_km  >= 0),
    CONSTRAINT chk_pricing_base_min CHECK (base_fare_per_min >= 0),
    CONSTRAINT chk_pricing_min_fare CHECK (min_fare >= 0),
    CONSTRAINT chk_pricing_max_fare CHECK (max_fare >= 0),
    CONSTRAINT chk_pricing_minmax   CHECK (max_fare >= min_fare OR max_fare = 0),
    CONSTRAINT chk_pricing_book_fee CHECK (booking_fee >= 0)
);

CREATE TRIGGER set_updated_at BEFORE UPDATE ON public.admin_pricing_rules
    FOR EACH ROW EXECUTE FUNCTION trigger_set_updated_at();

CREATE INDEX idx_admin_pricing_rules_region ON public.admin_pricing_rules (region_id);
CREATE INDEX idx_admin_pricing_rules_city    ON public.admin_pricing_rules (city_id, vehicle_type);
CREATE INDEX idx_admin_pricing_rules_active  ON public.admin_pricing_rules (active) WHERE active = TRUE;

CREATE TABLE IF NOT EXISTS public.admin_match_configs (
    city_id        VARCHAR(36)      NOT NULL PRIMARY KEY,
    max_radius     DOUBLE PRECISION NOT NULL DEFAULT 5.0,
    min_score      DOUBLE PRECISION NOT NULL DEFAULT 0.5,
    max_candidates INT              NOT NULL DEFAULT 10,
    updated_at     TIMESTAMPTZ      NOT NULL DEFAULT NOW(),
    CONSTRAINT chk_match_cfg_radius CHECK (max_radius > 0),
    CONSTRAINT chk_match_cfg_score  CHECK (min_score BETWEEN 0 AND 1),
    CONSTRAINT chk_match_cfg_cands  CHECK (max_candidates > 0)
);

CREATE TRIGGER set_updated_at BEFORE UPDATE ON public.admin_match_configs
    FOR EACH ROW EXECUTE FUNCTION trigger_set_updated_at();
