-- ============================================================
-- ADMIN SCHEMA: Support (Tickets & Escalations)
-- Database: public | Prefix: admin_
-- ============================================================

CREATE TYPE admin_ticket_status   AS ENUM ('open', 'in_progress', 'resolved', 'closed');
CREATE TYPE admin_ticket_priority AS ENUM ('low', 'medium', 'high', 'urgent');

CREATE TABLE IF NOT EXISTS public.admin_tickets (
    id          VARCHAR(36)           PRIMARY KEY DEFAULT gen_random_uuid()::TEXT,
    user_id     VARCHAR(36)           NOT NULL,
    user_type   VARCHAR(20)           NOT NULL DEFAULT 'user',
    category    VARCHAR(100)          NOT NULL DEFAULT '',
    subject     VARCHAR(300)          NOT NULL,
    description TEXT                  NOT NULL DEFAULT '',
    status      admin_ticket_status   NOT NULL DEFAULT 'open',
    priority    admin_ticket_priority NOT NULL DEFAULT 'medium',
    assigned_to VARCHAR(36),
    created_at  TIMESTAMPTZ           NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ           NOT NULL DEFAULT NOW(),
    resolved_at TIMESTAMPTZ,
    deleted_at  TIMESTAMPTZ,
    -- Full-text search vector (auto-updated via trigger below)
    search_tsv  TSVECTOR,
    CONSTRAINT chk_ticket_subject   CHECK (length(subject) >= 1),
    CONSTRAINT chk_ticket_user_type CHECK (user_type IN ('user', 'driver'))
);

CREATE TRIGGER set_updated_at BEFORE UPDATE ON public.admin_tickets
    FOR EACH ROW EXECUTE FUNCTION trigger_set_updated_at();

-- Auto-populate search_tsv on insert/update
CREATE OR REPLACE FUNCTION admin_ticket_search_tsv_trigger()
RETURNS TRIGGER AS $$
BEGIN
    NEW.search_tsv := to_tsvector('english', COALESCE(NEW.subject, '') || ' ' || COALESCE(NEW.description, ''));
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_search_tsv BEFORE INSERT OR UPDATE OF subject, description ON public.admin_tickets
    FOR EACH ROW EXECUTE FUNCTION admin_ticket_search_tsv_trigger();

CREATE INDEX idx_admin_tickets_user      ON public.admin_tickets (user_id, user_type);
CREATE INDEX idx_admin_tickets_status    ON public.admin_tickets (status);
CREATE INDEX idx_admin_tickets_priority  ON public.admin_tickets (priority);
CREATE INDEX idx_admin_tickets_assigned  ON public.admin_tickets (assigned_to) WHERE assigned_to IS NOT NULL;
CREATE INDEX idx_admin_tickets_composite ON public.admin_tickets (status, priority, created_at);
CREATE INDEX idx_admin_tickets_search    ON public.admin_tickets USING GIN (search_tsv);
CREATE INDEX idx_admin_tickets_deleted   ON public.admin_tickets (deleted_at) WHERE deleted_at IS NULL;

CREATE TYPE admin_escalation_status AS ENUM ('pending', 'acknowledged', 'resolved');

CREATE TABLE IF NOT EXISTS public.admin_escalations (
    id            VARCHAR(36)             PRIMARY KEY DEFAULT gen_random_uuid()::TEXT,
    ticket_id     VARCHAR(36)             NOT NULL,
    from_agent_id VARCHAR(36),
    to_agent_id   VARCHAR(36),
    reason        TEXT                    NOT NULL DEFAULT '',
    priority      admin_ticket_priority   NOT NULL DEFAULT 'high',
    status        admin_escalation_status NOT NULL DEFAULT 'pending',
    created_at    TIMESTAMPTZ             NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMPTZ             NOT NULL DEFAULT NOW()
);

CREATE TRIGGER set_updated_at BEFORE UPDATE ON public.admin_escalations
    FOR EACH ROW EXECUTE FUNCTION trigger_set_updated_at();

CREATE INDEX idx_admin_escalations_ticket ON public.admin_escalations (ticket_id);
CREATE INDEX idx_admin_escalations_status ON public.admin_escalations (status);
CREATE INDEX idx_admin_escalations_agents ON public.admin_escalations (to_agent_id) WHERE to_agent_id IS NOT NULL;
