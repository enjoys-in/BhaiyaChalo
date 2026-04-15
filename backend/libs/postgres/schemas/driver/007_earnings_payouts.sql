-- ============================================================
-- DRIVER SCHEMA: Earnings, Daily/Weekly Summaries, Payouts
-- Database: public | Prefix: driver_
-- ============================================================

-- Partitioned by earned_at for time-series queries & retention
CREATE TABLE IF NOT EXISTS public.driver_earnings (
    id              VARCHAR(36)   NOT NULL DEFAULT gen_random_uuid()::TEXT,
    driver_id       VARCHAR(36)   NOT NULL,
    trip_id         VARCHAR(36)   NOT NULL,
    fare_amount     NUMERIC(12,2) NOT NULL DEFAULT 0,
    commission      NUMERIC(12,2) NOT NULL DEFAULT 0,
    incentive_bonus NUMERIC(12,2) NOT NULL DEFAULT 0,
    tip_amount      NUMERIC(12,2) NOT NULL DEFAULT 0,
    net_earning     NUMERIC(12,2) NOT NULL DEFAULT 0,
    currency        VARCHAR(10)   NOT NULL DEFAULT 'INR',
    earned_at       TIMESTAMPTZ   NOT NULL DEFAULT NOW(),
    created_at      TIMESTAMPTZ   NOT NULL DEFAULT NOW(),
    CONSTRAINT chk_earning_fare       CHECK (fare_amount     >= 0),
    CONSTRAINT chk_earning_commission CHECK (commission      >= 0),
    CONSTRAINT chk_earning_incentive  CHECK (incentive_bonus >= 0),
    CONSTRAINT chk_earning_tip        CHECK (tip_amount      >= 0)
) PARTITION BY RANGE (earned_at);

CREATE TABLE public.driver_earnings_default
    PARTITION OF public.driver_earnings DEFAULT;

CREATE INDEX idx_driver_earnings_driver ON public.driver_earnings (driver_id);
CREATE INDEX idx_driver_earnings_trip   ON public.driver_earnings (trip_id);
CREATE INDEX idx_driver_earnings_time   ON public.driver_earnings (earned_at);
CREATE INDEX idx_driver_earnings_comp   ON public.driver_earnings (driver_id, earned_at);

CREATE TABLE IF NOT EXISTS public.driver_daily_summaries (
    driver_id        VARCHAR(36)   NOT NULL,
    date             DATE          NOT NULL,
    total_trips      INT           NOT NULL DEFAULT 0,
    total_fare       NUMERIC(12,2) NOT NULL DEFAULT 0,
    total_commission NUMERIC(12,2) NOT NULL DEFAULT 0,
    total_incentive  NUMERIC(12,2) NOT NULL DEFAULT 0,
    total_tips       NUMERIC(12,2) NOT NULL DEFAULT 0,
    net_earning      NUMERIC(12,2) NOT NULL DEFAULT 0,
    PRIMARY KEY (driver_id, date),
    CONSTRAINT chk_daily_trips      CHECK (total_trips      >= 0),
    CONSTRAINT chk_daily_fare       CHECK (total_fare       >= 0),
    CONSTRAINT chk_daily_commission CHECK (total_commission >= 0),
    CONSTRAINT chk_daily_incentive  CHECK (total_incentive  >= 0),
    CONSTRAINT chk_daily_tips       CHECK (total_tips       >= 0)
);

CREATE TABLE IF NOT EXISTS public.driver_weekly_summaries (
    driver_id        VARCHAR(36)   NOT NULL,
    week_start       DATE          NOT NULL,
    week_end         DATE          NOT NULL,
    total_trips      INT           NOT NULL DEFAULT 0,
    total_fare       NUMERIC(12,2) NOT NULL DEFAULT 0,
    total_commission NUMERIC(12,2) NOT NULL DEFAULT 0,
    total_incentive  NUMERIC(12,2) NOT NULL DEFAULT 0,
    total_tips       NUMERIC(12,2) NOT NULL DEFAULT 0,
    net_earning      NUMERIC(12,2) NOT NULL DEFAULT 0,
    PRIMARY KEY (driver_id, week_start),
    CONSTRAINT chk_weekly_range      CHECK (week_end > week_start),
    CONSTRAINT chk_weekly_trips      CHECK (total_trips      >= 0),
    CONSTRAINT chk_weekly_fare       CHECK (total_fare       >= 0),
    CONSTRAINT chk_weekly_commission CHECK (total_commission >= 0),
    CONSTRAINT chk_weekly_incentive  CHECK (total_incentive  >= 0),
    CONSTRAINT chk_weekly_tips       CHECK (total_tips       >= 0)
);

-- Payouts

CREATE TYPE driver_payout_status AS ENUM ('pending', 'processing', 'completed', 'failed');

CREATE TABLE IF NOT EXISTS public.driver_payouts (
    id              VARCHAR(36)          PRIMARY KEY DEFAULT gen_random_uuid()::TEXT,
    driver_id       VARCHAR(36)          NOT NULL,
    amount          NUMERIC(12,2)        NOT NULL DEFAULT 0,
    currency        VARCHAR(10)          NOT NULL DEFAULT 'INR',
    method          VARCHAR(30)          NOT NULL DEFAULT 'bank_transfer',
    bank_account_id VARCHAR(100)         NOT NULL DEFAULT '',
    status          driver_payout_status NOT NULL DEFAULT 'pending',
    reference       VARCHAR(100)         NOT NULL DEFAULT '',
    failure_reason  TEXT,
    initiated_at    TIMESTAMPTZ          NOT NULL DEFAULT NOW(),
    completed_at    TIMESTAMPTZ,
    created_at      TIMESTAMPTZ          NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ          NOT NULL DEFAULT NOW(),
    CONSTRAINT chk_payout_amount CHECK (amount > 0)
);

CREATE TRIGGER set_updated_at BEFORE UPDATE ON public.driver_payouts
    FOR EACH ROW EXECUTE FUNCTION trigger_set_updated_at();

CREATE INDEX idx_driver_payouts_driver    ON public.driver_payouts (driver_id);
CREATE INDEX idx_driver_payouts_status    ON public.driver_payouts (status);
CREATE INDEX idx_driver_payouts_time      ON public.driver_payouts (initiated_at);
CREATE INDEX idx_driver_payouts_composite ON public.driver_payouts (driver_id, status, initiated_at);
