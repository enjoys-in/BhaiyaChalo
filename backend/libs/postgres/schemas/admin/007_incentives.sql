-- ============================================================
-- ADMIN SCHEMA: Incentives
-- Database: public | Prefix: admin_
-- ============================================================

CREATE TYPE admin_incentive_type AS ENUM ('trip_count', 'revenue_target', 'peak_hours', 'referral', 'streak');

CREATE TABLE IF NOT EXISTS public.admin_incentives (
    id           VARCHAR(36)          PRIMARY KEY DEFAULT gen_random_uuid()::TEXT,
    name         VARCHAR(200)         NOT NULL,
    description  TEXT                 NOT NULL DEFAULT '',
    city_id      VARCHAR(36)          NOT NULL DEFAULT '',
    target_type  VARCHAR(30)          NOT NULL DEFAULT 'driver',
    type         admin_incentive_type NOT NULL DEFAULT 'trip_count',
    min_trips    INT                  NOT NULL DEFAULT 0,
    bonus_amount NUMERIC(12,2)        NOT NULL DEFAULT 0,
    currency     VARCHAR(10)          NOT NULL DEFAULT 'INR',
    valid_from   TIMESTAMPTZ          NOT NULL DEFAULT NOW(),
    valid_until  TIMESTAMPTZ          NOT NULL DEFAULT NOW(),
    active       BOOLEAN              NOT NULL DEFAULT TRUE,
    created_at   TIMESTAMPTZ          NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMPTZ          NOT NULL DEFAULT NOW(),
    deleted_at   TIMESTAMPTZ,
    CONSTRAINT chk_incentive_name       CHECK (length(name) >= 1),
    CONSTRAINT chk_incentive_min_trips  CHECK (min_trips >= 0),
    CONSTRAINT chk_incentive_bonus      CHECK (bonus_amount >= 0),
    CONSTRAINT chk_incentive_valid_range CHECK (valid_until >= valid_from),
    CONSTRAINT chk_incentive_target     CHECK (target_type IN ('driver', 'user'))
);

CREATE TRIGGER set_updated_at BEFORE UPDATE ON public.admin_incentives
    FOR EACH ROW EXECUTE FUNCTION trigger_set_updated_at();

CREATE INDEX idx_admin_incentives_city    ON public.admin_incentives (city_id);
CREATE INDEX idx_admin_incentives_active  ON public.admin_incentives (active) WHERE active = TRUE;
CREATE INDEX idx_admin_incentives_valid   ON public.admin_incentives (valid_from, valid_until);
CREATE INDEX idx_admin_incentives_deleted ON public.admin_incentives (deleted_at) WHERE deleted_at IS NULL;
