-- ============================================================
-- USER SCHEMA: Profile, Addresses, Emergency Contacts
-- Database: public | Prefix: user_
-- Requires: postgis extension (see 000_init.sql)
-- ============================================================

CREATE TYPE user_profile_status AS ENUM ('active', 'inactive', 'suspended');
CREATE TYPE user_address_label  AS ENUM ('home', 'work', 'other');

CREATE TABLE IF NOT EXISTS public.user_profiles (
    id            UUID                PRIMARY KEY DEFAULT gen_random_uuid(),
    first_name    VARCHAR(100)        NOT NULL DEFAULT '',
    last_name     VARCHAR(100)        NOT NULL DEFAULT '',
    phone         VARCHAR(20)         NOT NULL UNIQUE,
    email         VARCHAR(255)        NOT NULL DEFAULT '',
    avatar_url    TEXT                NOT NULL DEFAULT '',
    gender        VARCHAR(20)         NOT NULL DEFAULT '',
    date_of_birth DATE,
    city_id       VARCHAR(36)         NOT NULL DEFAULT '',
    status        user_profile_status NOT NULL DEFAULT 'active',
    created_at    TIMESTAMPTZ         NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMPTZ         NOT NULL DEFAULT NOW(),
    deleted_at    TIMESTAMPTZ,
    CONSTRAINT chk_profile_phone CHECK (length(phone) >= 10),
    CONSTRAINT chk_profile_dob   CHECK (date_of_birth IS NULL OR date_of_birth < CURRENT_DATE)
);

CREATE TRIGGER set_updated_at BEFORE UPDATE ON public.user_profiles
    FOR EACH ROW EXECUTE FUNCTION trigger_set_updated_at();

CREATE INDEX idx_user_profiles_phone   ON public.user_profiles (phone);
CREATE INDEX idx_user_profiles_email   ON public.user_profiles (email) WHERE email != '';
CREATE INDEX idx_user_profiles_city    ON public.user_profiles (city_id);
CREATE INDEX idx_user_profiles_status  ON public.user_profiles (status);
CREATE INDEX idx_user_profiles_deleted ON public.user_profiles (deleted_at) WHERE deleted_at IS NULL;

CREATE TABLE IF NOT EXISTS public.user_addresses (
    id                UUID               PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id           UUID               NOT NULL,
    label             user_address_label NOT NULL DEFAULT 'other',
    lat               DOUBLE PRECISION   NOT NULL DEFAULT 0,
    lng               DOUBLE PRECISION   NOT NULL DEFAULT 0,
    -- PostGIS: auto-computed from lat/lng for spatial queries
    geom              GEOMETRY(Point, 4326) GENERATED ALWAYS AS (
                          ST_SetSRID(ST_MakePoint(lng, lat), 4326)
                      ) STORED,
    formatted_address TEXT               NOT NULL DEFAULT '',
    is_default        BOOLEAN            NOT NULL DEFAULT FALSE,
    created_at        TIMESTAMPTZ        NOT NULL DEFAULT NOW(),
    CONSTRAINT chk_addr_lat CHECK (lat BETWEEN -90 AND 90),
    CONSTRAINT chk_addr_lng CHECK (lng BETWEEN -180 AND 180)
);

CREATE INDEX idx_user_addresses_user ON public.user_addresses (user_id);
CREATE INDEX idx_user_addresses_geom ON public.user_addresses USING GIST (geom);

CREATE TABLE IF NOT EXISTS public.user_emergency_contacts (
    id       UUID         PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id  UUID         NOT NULL,
    name     VARCHAR(100) NOT NULL,
    phone    VARCHAR(20)  NOT NULL,
    relation VARCHAR(50)  NOT NULL DEFAULT '',
    CONSTRAINT chk_emg_name  CHECK (length(name) >= 1),
    CONSTRAINT chk_emg_phone CHECK (length(phone) >= 10)
);

CREATE INDEX idx_user_emergency_contacts_user ON public.user_emergency_contacts (user_id);
