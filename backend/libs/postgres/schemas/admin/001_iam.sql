-- ============================================================
-- ADMIN SCHEMA: IAM (Identity & Access Management)
-- Database: public | Prefix: admin_
-- ============================================================

CREATE TABLE IF NOT EXISTS public.admin_roles (
    id          VARCHAR(36)  PRIMARY KEY DEFAULT gen_random_uuid()::TEXT,
    name        VARCHAR(100) NOT NULL UNIQUE,
    description TEXT         NOT NULL DEFAULT '',
    created_at  TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    deleted_at  TIMESTAMPTZ,
    CONSTRAINT chk_admin_roles_name_len CHECK (length(name) >= 2)
);

CREATE TRIGGER set_updated_at BEFORE UPDATE ON public.admin_roles
    FOR EACH ROW EXECUTE FUNCTION trigger_set_updated_at();

CREATE INDEX idx_admin_roles_deleted ON public.admin_roles (deleted_at) WHERE deleted_at IS NULL;

CREATE TABLE IF NOT EXISTS public.admin_permissions (
    id          VARCHAR(36)  PRIMARY KEY DEFAULT gen_random_uuid()::TEXT,
    resource    VARCHAR(100) NOT NULL,
    action      VARCHAR(50)  NOT NULL,
    description TEXT         NOT NULL DEFAULT '',
    CONSTRAINT uq_admin_permissions_res_act UNIQUE (resource, action),
    CONSTRAINT chk_admin_permissions_resource CHECK (length(resource) >= 1),
    CONSTRAINT chk_admin_permissions_action   CHECK (length(action) >= 1)
);

CREATE TABLE IF NOT EXISTS public.admin_role_permissions (
    role_id       VARCHAR(36) NOT NULL REFERENCES public.admin_roles(id) ON DELETE CASCADE,
    permission_id VARCHAR(36) NOT NULL REFERENCES public.admin_permissions(id) ON DELETE CASCADE,
    assigned_at   TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (role_id, permission_id)
);

CREATE TABLE IF NOT EXISTS public.admin_policies (
    id           VARCHAR(36)  PRIMARY KEY DEFAULT gen_random_uuid()::TEXT,
    subject_type VARCHAR(50)  NOT NULL,
    subject_id   VARCHAR(36)  NOT NULL,
    role_id      VARCHAR(36)  NOT NULL REFERENCES public.admin_roles(id) ON DELETE CASCADE,
    scope        VARCHAR(100) NOT NULL DEFAULT '',
    created_at   TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    CONSTRAINT chk_admin_policies_subtype CHECK (subject_type IN ('user', 'driver', 'admin', 'service'))
);

CREATE INDEX idx_admin_policies_subject   ON public.admin_policies (subject_type, subject_id);
CREATE INDEX idx_admin_policies_role      ON public.admin_policies (role_id);
CREATE INDEX idx_admin_policies_composite ON public.admin_policies (subject_type, subject_id, scope);
