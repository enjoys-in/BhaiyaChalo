-- ============================================================
-- USER SCHEMA: Auth & Sessions
-- Database: public | Prefix: user_
-- ============================================================

CREATE TYPE user_role   AS ENUM ('user', 'driver', 'admin');
CREATE TYPE user_status AS ENUM ('active', 'inactive', 'blocked');

CREATE TABLE IF NOT EXISTS public.user_auth_users (
    id         VARCHAR(36)  PRIMARY KEY DEFAULT gen_random_uuid()::TEXT,
    phone      VARCHAR(20)  NOT NULL UNIQUE,
    email      VARCHAR(255),
    role       user_role    NOT NULL DEFAULT 'user',
    status     user_status  NOT NULL DEFAULT 'active',
    created_at TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ,
    CONSTRAINT chk_auth_phone CHECK (length(phone) >= 10)
);

CREATE TRIGGER set_updated_at BEFORE UPDATE ON public.user_auth_users
    FOR EACH ROW EXECUTE FUNCTION trigger_set_updated_at();

CREATE INDEX idx_user_auth_users_phone   ON public.user_auth_users (phone);
CREATE INDEX idx_user_auth_users_email   ON public.user_auth_users (email) WHERE email IS NOT NULL;
CREATE INDEX idx_user_auth_users_status  ON public.user_auth_users (status);
CREATE INDEX idx_user_auth_users_deleted ON public.user_auth_users (deleted_at) WHERE deleted_at IS NULL;

CREATE TABLE IF NOT EXISTS public.user_tokens (
    id            VARCHAR(36) PRIMARY KEY DEFAULT gen_random_uuid()::TEXT,
    user_id       VARCHAR(36) NOT NULL,
    refresh_token TEXT        NOT NULL,
    expires_at    TIMESTAMPTZ NOT NULL,
    created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT chk_token_refresh CHECK (length(refresh_token) >= 1)
);

CREATE INDEX idx_user_tokens_user   ON public.user_tokens (user_id);
CREATE INDEX idx_user_tokens_expiry ON public.user_tokens (expires_at);
CREATE INDEX idx_user_tokens_comp   ON public.user_tokens (user_id, expires_at);

CREATE TYPE user_otp_purpose AS ENUM ('login', 'phone_verify', 'password_reset');

CREATE TABLE IF NOT EXISTS public.user_otps (
    id         VARCHAR(36)      PRIMARY KEY DEFAULT gen_random_uuid()::TEXT,
    phone      VARCHAR(20)      NOT NULL,
    code       VARCHAR(10)      NOT NULL,
    purpose    user_otp_purpose NOT NULL DEFAULT 'login',
    verified   BOOLEAN          NOT NULL DEFAULT FALSE,
    attempts   INT              NOT NULL DEFAULT 0,
    expires_at TIMESTAMPTZ      NOT NULL,
    created_at TIMESTAMPTZ      NOT NULL DEFAULT NOW(),
    CONSTRAINT chk_otp_attempts  CHECK (attempts >= 0 AND attempts <= 10),
    CONSTRAINT chk_otp_code_len  CHECK (length(code) >= 4),
    CONSTRAINT chk_otp_phone_len CHECK (length(phone) >= 10)
);

CREATE INDEX idx_user_otps_phone  ON public.user_otps (phone);
CREATE INDEX idx_user_otps_expiry ON public.user_otps (expires_at);
CREATE INDEX idx_user_otps_comp   ON public.user_otps (phone, purpose, verified);

CREATE TABLE IF NOT EXISTS public.user_sessions (
    id         VARCHAR(36)  PRIMARY KEY DEFAULT gen_random_uuid()::TEXT,
    user_id    VARCHAR(36)  NOT NULL,
    role       VARCHAR(20)  NOT NULL DEFAULT 'user',
    device_id  VARCHAR(100) NOT NULL DEFAULT '',
    ip         INET,
    user_agent TEXT         NOT NULL DEFAULT '',
    active     BOOLEAN      NOT NULL DEFAULT TRUE,
    expires_at TIMESTAMPTZ  NOT NULL,
    created_at TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE TRIGGER set_updated_at BEFORE UPDATE ON public.user_sessions
    FOR EACH ROW EXECUTE FUNCTION trigger_set_updated_at();

CREATE INDEX idx_user_sessions_user      ON public.user_sessions (user_id);
CREATE INDEX idx_user_sessions_active    ON public.user_sessions (active) WHERE active = TRUE;
CREATE INDEX idx_user_sessions_composite ON public.user_sessions (user_id, active, expires_at);
