-- ============================================================
-- ADMIN SCHEMA: Reconciliation & Templates
-- Database: public | Prefix: admin_
-- ============================================================

CREATE TYPE admin_recon_status AS ENUM ('pending', 'matched', 'mismatched', 'resolved');

CREATE TABLE IF NOT EXISTS public.admin_reconciliations (
    id             VARCHAR(36)        PRIMARY KEY DEFAULT gen_random_uuid()::TEXT,
    payment_id     VARCHAR(36)        NOT NULL,
    gateway_id     VARCHAR(100)       NOT NULL DEFAULT '',
    gateway_amount NUMERIC(12,2)      NOT NULL DEFAULT 0,
    system_amount  NUMERIC(12,2)      NOT NULL DEFAULT 0,
    currency       VARCHAR(10)        NOT NULL DEFAULT 'INR',
    status         admin_recon_status NOT NULL DEFAULT 'pending',
    discrepancy    NUMERIC(12,2)      NOT NULL DEFAULT 0,
    reconciled_at  TIMESTAMPTZ,
    created_at     TIMESTAMPTZ        NOT NULL DEFAULT NOW(),
    updated_at     TIMESTAMPTZ        NOT NULL DEFAULT NOW(),
    CONSTRAINT chk_recon_gateway_amt CHECK (gateway_amount >= 0),
    CONSTRAINT chk_recon_system_amt  CHECK (system_amount  >= 0)
);

CREATE TRIGGER set_updated_at BEFORE UPDATE ON public.admin_reconciliations
    FOR EACH ROW EXECUTE FUNCTION trigger_set_updated_at();

CREATE INDEX idx_admin_reconciliations_payment ON public.admin_reconciliations (payment_id);
CREATE INDEX idx_admin_reconciliations_status  ON public.admin_reconciliations (status);
CREATE INDEX idx_admin_reconciliations_comp    ON public.admin_reconciliations (status, created_at);

CREATE TYPE admin_template_channel AS ENUM ('push', 'sms', 'email', 'in_app');

CREATE TABLE IF NOT EXISTS public.admin_templates (
    id         VARCHAR(36)            PRIMARY KEY DEFAULT gen_random_uuid()::TEXT,
    name       VARCHAR(200)           NOT NULL UNIQUE,
    type       VARCHAR(50)            NOT NULL DEFAULT 'transactional',
    channel    admin_template_channel NOT NULL DEFAULT 'push',
    subject    VARCHAR(300)           NOT NULL DEFAULT '',
    body       TEXT                   NOT NULL DEFAULT '',
    variables  JSONB                  NOT NULL DEFAULT '[]',
    active     BOOLEAN                NOT NULL DEFAULT TRUE,
    created_at TIMESTAMPTZ            NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ            NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ,
    CONSTRAINT chk_template_name CHECK (length(name) >= 1)
);

CREATE TRIGGER set_updated_at BEFORE UPDATE ON public.admin_templates
    FOR EACH ROW EXECUTE FUNCTION trigger_set_updated_at();

CREATE INDEX idx_admin_templates_channel   ON public.admin_templates (channel);
CREATE INDEX idx_admin_templates_active    ON public.admin_templates (active) WHERE active = TRUE;
CREATE INDEX idx_admin_templates_variables ON public.admin_templates USING GIN (variables);
CREATE INDEX idx_admin_templates_deleted   ON public.admin_templates (deleted_at) WHERE deleted_at IS NULL;
