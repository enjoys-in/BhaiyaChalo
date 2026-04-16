-- ============================================================
-- ADMIN SCHEMA: Surge Policies & Zones
-- Database: public | Prefix: admin_
-- Requires: 000_geo_regions.sql (regions/cities)
--
-- GEO-SHARDING: surge_history uses compound partition
--   Level 1: LIST (region_id)     → routes to regional shard
--   Level 2: RANGE (calculated_at) → monthly time partitions
-- ============================================================

CREATE TABLE IF NOT EXISTS public.admin_surge_policies (
    id                       VARCHAR(36)      PRIMARY KEY DEFAULT gen_random_uuid()::TEXT,
    city_id                  VARCHAR(36)      NOT NULL,
    region_id                VARCHAR(36)      NOT NULL DEFAULT '',
    min_demand_supply_ratio  DOUBLE PRECISION NOT NULL DEFAULT 1.5,
    max_multiplier           DOUBLE PRECISION NOT NULL DEFAULT 3.0,
    step_size                DOUBLE PRECISION NOT NULL DEFAULT 0.1,
    cooldown_minutes         INT              NOT NULL DEFAULT 10,
    active                   BOOLEAN          NOT NULL DEFAULT TRUE,
    created_at               TIMESTAMPTZ      NOT NULL DEFAULT NOW(),
    updated_at               TIMESTAMPTZ      NOT NULL DEFAULT NOW(),
    CONSTRAINT chk_surge_pol_ratio     CHECK (min_demand_supply_ratio > 0),
    CONSTRAINT chk_surge_pol_max_mult  CHECK (max_multiplier >= 1.0),
    CONSTRAINT chk_surge_pol_step      CHECK (step_size > 0),
    CONSTRAINT chk_surge_pol_cooldown  CHECK (cooldown_minutes >= 0)
);

CREATE TRIGGER set_updated_at BEFORE UPDATE ON public.admin_surge_policies
    FOR EACH ROW EXECUTE FUNCTION trigger_set_updated_at();

CREATE INDEX idx_admin_surge_policies_city   ON public.admin_surge_policies (city_id);
CREATE INDEX idx_admin_surge_policies_region ON public.admin_surge_policies (region_id);
CREATE INDEX idx_admin_surge_policies_active ON public.admin_surge_policies (active) WHERE active = TRUE;

CREATE TABLE IF NOT EXISTS public.admin_surge_zones (
    id                  VARCHAR(36)      PRIMARY KEY DEFAULT gen_random_uuid()::TEXT,
    city_id             VARCHAR(36)      NOT NULL,
    region_id           VARCHAR(36)      NOT NULL DEFAULT '',
    geofence_id         VARCHAR(36)      NOT NULL,
    current_multiplier  DOUBLE PRECISION NOT NULL DEFAULT 1.0,
    demand_count        INT              NOT NULL DEFAULT 0,
    supply_count        INT              NOT NULL DEFAULT 0,
    updated_at          TIMESTAMPTZ      NOT NULL DEFAULT NOW(),
    CONSTRAINT chk_surge_zone_mult   CHECK (current_multiplier >= 1.0),
    CONSTRAINT chk_surge_zone_demand CHECK (demand_count >= 0),
    CONSTRAINT chk_surge_zone_supply CHECK (supply_count >= 0)
);

CREATE TRIGGER set_updated_at BEFORE UPDATE ON public.admin_surge_zones
    FOR EACH ROW EXECUTE FUNCTION trigger_set_updated_at();

CREATE INDEX idx_admin_surge_zones_city      ON public.admin_surge_zones (city_id);
CREATE INDEX idx_admin_surge_zones_region    ON public.admin_surge_zones (region_id);
CREATE INDEX idx_admin_surge_zones_geofence  ON public.admin_surge_zones (geofence_id);
CREATE INDEX idx_admin_surge_zones_composite ON public.admin_surge_zones (region_id, city_id, geofence_id);

-- Compound partitioned: LIST(region_id) → RANGE(calculated_at)
CREATE TABLE IF NOT EXISTS public.admin_surge_history (
    id            VARCHAR(36)      NOT NULL DEFAULT gen_random_uuid()::TEXT,
    zone_id       VARCHAR(36)      NOT NULL,
    city_id       VARCHAR(36)      NOT NULL DEFAULT '',
    region_id     VARCHAR(36)      NOT NULL,
    multiplier    DOUBLE PRECISION NOT NULL,
    demand_count  INT              NOT NULL DEFAULT 0,
    supply_count  INT              NOT NULL DEFAULT 0,
    calculated_at TIMESTAMPTZ      NOT NULL DEFAULT NOW(),
    CONSTRAINT chk_surge_hist_mult   CHECK (multiplier >= 1.0),
    CONSTRAINT chk_surge_hist_demand CHECK (demand_count >= 0),
    CONSTRAINT chk_surge_hist_supply CHECK (supply_count >= 0)
) PARTITION BY LIST (region_id);

-- Bootstrap region partitions (each sub-partitioned by calculated_at monthly)
SELECT bootstrap_region_partitions('public.admin_surge_history', 'calculated_at');

CREATE INDEX idx_admin_surge_history_zone   ON public.admin_surge_history (zone_id);
CREATE INDEX idx_admin_surge_history_region ON public.admin_surge_history (region_id);
CREATE INDEX idx_admin_surge_history_time   ON public.admin_surge_history (calculated_at);
CREATE INDEX idx_admin_surge_history_comp   ON public.admin_surge_history (region_id, zone_id, calculated_at);
