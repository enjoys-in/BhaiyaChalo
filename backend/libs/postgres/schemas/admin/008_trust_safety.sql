-- ============================================================
-- ADMIN SCHEMA: Trust & Safety (Fraud, Risk, Audit)
-- Database: public | Prefix: admin_
-- ============================================================

CREATE TYPE admin_severity AS ENUM ('low', 'medium', 'high', 'critical');
CREATE TYPE admin_fraud_status AS ENUM ('open', 'investigating', 'confirmed', 'dismissed');

CREATE TABLE IF NOT EXISTS public.admin_fraud_signals (
    id          VARCHAR(36)        PRIMARY KEY DEFAULT gen_random_uuid()::TEXT,
    user_id     VARCHAR(36)        NOT NULL,
    user_type   VARCHAR(20)        NOT NULL DEFAULT 'user',
    signal_type VARCHAR(100)       NOT NULL,
    severity    admin_severity     NOT NULL DEFAULT 'low',
    status      admin_fraud_status NOT NULL DEFAULT 'open',
    description TEXT               NOT NULL DEFAULT '',
    trip_id     VARCHAR(36),
    metadata    JSONB              NOT NULL DEFAULT '{}',
    created_at  TIMESTAMPTZ        NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ        NOT NULL DEFAULT NOW(),
    CONSTRAINT chk_fraud_user_type CHECK (user_type IN ('user', 'driver'))
);

CREATE TRIGGER set_updated_at BEFORE UPDATE ON public.admin_fraud_signals
    FOR EACH ROW EXECUTE FUNCTION trigger_set_updated_at();

CREATE INDEX idx_admin_fraud_signals_user     ON public.admin_fraud_signals (user_id, user_type);
CREATE INDEX idx_admin_fraud_signals_severity ON public.admin_fraud_signals (severity);
CREATE INDEX idx_admin_fraud_signals_type     ON public.admin_fraud_signals (signal_type);
CREATE INDEX idx_admin_fraud_signals_status   ON public.admin_fraud_signals (status);
CREATE INDEX idx_admin_fraud_signals_meta     ON public.admin_fraud_signals USING GIN (metadata);
CREATE INDEX idx_admin_fraud_signals_comp     ON public.admin_fraud_signals (user_id, severity, created_at);

CREATE TABLE IF NOT EXISTS public.admin_risk_scores (
    id            VARCHAR(36)      PRIMARY KEY DEFAULT gen_random_uuid()::TEXT,
    user_id       VARCHAR(36)      NOT NULL,
    user_type     VARCHAR(20)      NOT NULL DEFAULT 'user',
    score         DOUBLE PRECISION NOT NULL DEFAULT 0,
    factors       JSONB            NOT NULL DEFAULT '{}',
    calculated_at TIMESTAMPTZ      NOT NULL DEFAULT NOW(),
    created_at    TIMESTAMPTZ      NOT NULL DEFAULT NOW(),
    CONSTRAINT chk_risk_score_range CHECK (score BETWEEN 0 AND 100),
    CONSTRAINT chk_risk_user_type   CHECK (user_type IN ('user', 'driver'))
);

CREATE INDEX idx_admin_risk_scores_user    ON public.admin_risk_scores (user_id, user_type);
CREATE INDEX idx_admin_risk_scores_score   ON public.admin_risk_scores (score);
CREATE INDEX idx_admin_risk_scores_factors ON public.admin_risk_scores USING GIN (factors);
CREATE INDEX idx_admin_risk_scores_comp    ON public.admin_risk_scores (user_id, user_type, calculated_at DESC);

-- Partitioned by created_at for time-series retention
CREATE TABLE IF NOT EXISTS public.admin_audit_logs (
    id          VARCHAR(36)  NOT NULL DEFAULT gen_random_uuid()::TEXT,
    actor_id    VARCHAR(36)  NOT NULL,
    actor_type  VARCHAR(20)  NOT NULL DEFAULT 'admin',
    action      VARCHAR(100) NOT NULL,
    resource    VARCHAR(100) NOT NULL,
    resource_id VARCHAR(36)  NOT NULL DEFAULT '',
    old_value   JSONB,
    new_value   JSONB,
    ip_address  INET,
    user_agent  TEXT         NOT NULL DEFAULT '',
    created_at  TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    CONSTRAINT chk_audit_actor_type CHECK (actor_type IN ('admin', 'system', 'service'))
) PARTITION BY RANGE (created_at);

CREATE TABLE public.admin_audit_logs_default
    PARTITION OF public.admin_audit_logs DEFAULT;

CREATE INDEX idx_admin_audit_logs_actor    ON public.admin_audit_logs (actor_id, actor_type);
CREATE INDEX idx_admin_audit_logs_resource ON public.admin_audit_logs (resource, resource_id);
CREATE INDEX idx_admin_audit_logs_time     ON public.admin_audit_logs (created_at);
CREATE INDEX idx_admin_audit_logs_comp     ON public.admin_audit_logs (resource, action, created_at);
CREATE INDEX idx_admin_audit_logs_old_val  ON public.admin_audit_logs USING GIN (old_value) WHERE old_value IS NOT NULL;
CREATE INDEX idx_admin_audit_logs_new_val  ON public.admin_audit_logs USING GIN (new_value) WHERE new_value IS NOT NULL;
