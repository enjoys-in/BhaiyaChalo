-- ============================================================
-- ADMIN SCHEMA: Campaigns & Promos
-- Database: public | Prefix: admin_
-- ============================================================

CREATE TYPE admin_campaign_type   AS ENUM ('push', 'sms', 'email', 'in_app');
CREATE TYPE admin_campaign_status AS ENUM ('draft', 'scheduled', 'active', 'completed', 'paused');

CREATE TABLE IF NOT EXISTS public.admin_campaigns (
    id              VARCHAR(36)           PRIMARY KEY DEFAULT gen_random_uuid()::TEXT,
    name            VARCHAR(200)          NOT NULL,
    description     TEXT                  NOT NULL DEFAULT '',
    type            admin_campaign_type   NOT NULL,
    target_audience VARCHAR(100)          NOT NULL DEFAULT '',
    city_id         VARCHAR(36)           NOT NULL DEFAULT '',
    region_id       VARCHAR(36)           NOT NULL DEFAULT '',
    promo_code_id   VARCHAR(36),
    status          admin_campaign_status NOT NULL DEFAULT 'draft',
    scheduled_at    TIMESTAMPTZ,
    started_at      TIMESTAMPTZ,
    completed_at    TIMESTAMPTZ,
    created_at      TIMESTAMPTZ           NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ           NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ,
    CONSTRAINT chk_campaign_name CHECK (length(name) >= 1)
);

CREATE TRIGGER set_updated_at BEFORE UPDATE ON public.admin_campaigns
    FOR EACH ROW EXECUTE FUNCTION trigger_set_updated_at();

CREATE INDEX idx_admin_campaigns_status  ON public.admin_campaigns (status);
CREATE INDEX idx_admin_campaigns_region ON public.admin_campaigns (region_id);
CREATE INDEX idx_admin_campaigns_city    ON public.admin_campaigns (city_id);
CREATE INDEX idx_admin_campaigns_sched   ON public.admin_campaigns (scheduled_at) WHERE scheduled_at IS NOT NULL;
CREATE INDEX idx_admin_campaigns_deleted ON public.admin_campaigns (deleted_at) WHERE deleted_at IS NULL;

CREATE TABLE IF NOT EXISTS public.admin_campaign_stats (
    campaign_id VARCHAR(36) PRIMARY KEY,
    sent        INT         NOT NULL DEFAULT 0,
    delivered   INT         NOT NULL DEFAULT 0,
    opened      INT         NOT NULL DEFAULT 0,
    clicked     INT         NOT NULL DEFAULT 0,
    converted   INT         NOT NULL DEFAULT 0,
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT chk_camp_stats_sent      CHECK (sent      >= 0),
    CONSTRAINT chk_camp_stats_delivered CHECK (delivered >= 0),
    CONSTRAINT chk_camp_stats_opened    CHECK (opened    >= 0),
    CONSTRAINT chk_camp_stats_clicked   CHECK (clicked   >= 0),
    CONSTRAINT chk_camp_stats_converted CHECK (converted >= 0)
);

CREATE TRIGGER set_updated_at BEFORE UPDATE ON public.admin_campaign_stats
    FOR EACH ROW EXECUTE FUNCTION trigger_set_updated_at();

CREATE TYPE admin_promo_type AS ENUM ('flat', 'percentage');

CREATE TABLE IF NOT EXISTS public.admin_promo_codes (
    id              VARCHAR(36)      PRIMARY KEY DEFAULT gen_random_uuid()::TEXT,
    code            VARCHAR(50)      NOT NULL UNIQUE,
    city_id         VARCHAR(36)      NOT NULL DEFAULT '',
    region_id       VARCHAR(36)      NOT NULL DEFAULT '',
    type            admin_promo_type NOT NULL DEFAULT 'flat',
    discount_value  NUMERIC(12,2)    NOT NULL DEFAULT 0,
    max_discount    NUMERIC(12,2)    NOT NULL DEFAULT 0,
    min_order_value NUMERIC(12,2)    NOT NULL DEFAULT 0,
    usage_limit     INT              NOT NULL DEFAULT 0,
    used_count      INT              NOT NULL DEFAULT 0,
    valid_from      TIMESTAMPTZ      NOT NULL DEFAULT NOW(),
    valid_until     TIMESTAMPTZ      NOT NULL DEFAULT NOW(),
    active          BOOLEAN          NOT NULL DEFAULT TRUE,
    created_at      TIMESTAMPTZ      NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ      NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ,
    CONSTRAINT chk_promo_discount    CHECK (discount_value  >= 0),
    CONSTRAINT chk_promo_max_disc    CHECK (max_discount    >= 0),
    CONSTRAINT chk_promo_min_order   CHECK (min_order_value >= 0),
    CONSTRAINT chk_promo_usage_limit CHECK (usage_limit     >= 0),
    CONSTRAINT chk_promo_used_count  CHECK (used_count      >= 0),
    CONSTRAINT chk_promo_code_len    CHECK (length(code)    >= 3),
    CONSTRAINT chk_promo_valid_range CHECK (valid_until >= valid_from)
);

CREATE TRIGGER set_updated_at BEFORE UPDATE ON public.admin_promo_codes
    FOR EACH ROW EXECUTE FUNCTION trigger_set_updated_at();

CREATE INDEX idx_admin_promo_codes_code    ON public.admin_promo_codes (code);
CREATE INDEX idx_admin_promo_codes_active  ON public.admin_promo_codes (active) WHERE active = TRUE;
CREATE INDEX idx_admin_promo_codes_region ON public.admin_promo_codes (region_id);
CREATE INDEX idx_admin_promo_codes_city    ON public.admin_promo_codes (city_id);
CREATE INDEX idx_admin_promo_codes_valid   ON public.admin_promo_codes (valid_from, valid_until);
CREATE INDEX idx_admin_promo_codes_deleted ON public.admin_promo_codes (deleted_at) WHERE deleted_at IS NULL;
