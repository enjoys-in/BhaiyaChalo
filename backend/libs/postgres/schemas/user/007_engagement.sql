-- ============================================================
-- USER SCHEMA: Promos, Referrals & Ratings
-- Database: public | Prefix: user_
-- ============================================================

CREATE TABLE IF NOT EXISTS public.user_promo_usages (
    id               VARCHAR(36)   PRIMARY KEY DEFAULT gen_random_uuid()::TEXT,
    promo_id         VARCHAR(36)   NOT NULL,
    user_id          VARCHAR(36)   NOT NULL,
    booking_id       VARCHAR(36)   NOT NULL,
    discount_applied NUMERIC(12,2) NOT NULL DEFAULT 0,
    used_at          TIMESTAMPTZ   NOT NULL DEFAULT NOW(),
    CONSTRAINT chk_promo_usage_discount CHECK (discount_applied >= 0)
);

CREATE INDEX idx_user_promo_usages_user    ON public.user_promo_usages (user_id);
CREATE INDEX idx_user_promo_usages_promo   ON public.user_promo_usages (promo_id);
CREATE INDEX idx_user_promo_usages_booking ON public.user_promo_usages (booking_id);
-- Prevent double-use of same promo on same booking
CREATE UNIQUE INDEX uq_user_promo_usages_booking_promo ON public.user_promo_usages (booking_id, promo_id);

-- Referrals

CREATE TYPE user_referral_status AS ENUM ('pending', 'completed', 'rewarded');
CREATE TYPE user_reward_type     AS ENUM ('referrer', 'referee');

CREATE TABLE IF NOT EXISTS public.user_referral_codes (
    id         VARCHAR(36)  PRIMARY KEY DEFAULT gen_random_uuid()::TEXT,
    user_id    VARCHAR(36)  NOT NULL,
    code       VARCHAR(20)  NOT NULL UNIQUE,
    active     BOOLEAN      NOT NULL DEFAULT TRUE,
    created_at TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    CONSTRAINT chk_referral_code_len CHECK (length(code) >= 4)
);

CREATE INDEX idx_user_referral_codes_user   ON public.user_referral_codes (user_id);
CREATE INDEX idx_user_referral_codes_code   ON public.user_referral_codes (code);
CREATE INDEX idx_user_referral_codes_active ON public.user_referral_codes (active) WHERE active = TRUE;

CREATE TABLE IF NOT EXISTS public.user_referrals (
    id               VARCHAR(36)          PRIMARY KEY DEFAULT gen_random_uuid()::TEXT,
    referrer_id      VARCHAR(36)          NOT NULL,
    referee_id       VARCHAR(36)          NOT NULL,
    referral_code_id VARCHAR(36)          NOT NULL,
    status           user_referral_status NOT NULL DEFAULT 'pending',
    completed_at     TIMESTAMPTZ,
    created_at       TIMESTAMPTZ          NOT NULL DEFAULT NOW(),
    CONSTRAINT chk_referral_not_self CHECK (referrer_id != referee_id)
);

CREATE INDEX idx_user_referrals_referrer ON public.user_referrals (referrer_id);
CREATE INDEX idx_user_referrals_referee  ON public.user_referrals (referee_id);
CREATE INDEX idx_user_referrals_status   ON public.user_referrals (status);
-- One referral per referee
CREATE UNIQUE INDEX uq_user_referrals_referee ON public.user_referrals (referee_id);

CREATE TABLE IF NOT EXISTS public.user_referral_rewards (
    id          VARCHAR(36)      PRIMARY KEY DEFAULT gen_random_uuid()::TEXT,
    referral_id VARCHAR(36)      NOT NULL,
    user_id     VARCHAR(36)      NOT NULL,
    type        user_reward_type NOT NULL,
    amount      NUMERIC(12,2)    NOT NULL DEFAULT 0,
    currency    VARCHAR(10)      NOT NULL DEFAULT 'INR',
    credited_at TIMESTAMPTZ,
    CONSTRAINT chk_ref_reward_amount CHECK (amount > 0)
);

CREATE INDEX idx_user_referral_rewards_referral ON public.user_referral_rewards (referral_id);
CREATE INDEX idx_user_referral_rewards_user     ON public.user_referral_rewards (user_id);

-- Ratings

CREATE TABLE IF NOT EXISTS public.user_ratings (
    id           VARCHAR(36)      PRIMARY KEY DEFAULT gen_random_uuid()::TEXT,
    trip_id      VARCHAR(36)      NOT NULL,
    booking_id   VARCHAR(36)      NOT NULL,
    from_user_id VARCHAR(36)      NOT NULL,
    to_user_id   VARCHAR(36)      NOT NULL,
    from_role    VARCHAR(20)      NOT NULL DEFAULT 'user',
    to_role      VARCHAR(20)      NOT NULL DEFAULT 'driver',
    score        DOUBLE PRECISION NOT NULL DEFAULT 5.0,
    comment      TEXT             NOT NULL DEFAULT '',
    created_at   TIMESTAMPTZ      NOT NULL DEFAULT NOW(),
    CONSTRAINT chk_rating_score      CHECK (score BETWEEN 1.0 AND 5.0),
    CONSTRAINT chk_rating_not_self   CHECK (from_user_id != to_user_id),
    CONSTRAINT chk_rating_from_role  CHECK (from_role IN ('user', 'driver')),
    CONSTRAINT chk_rating_to_role    CHECK (to_role IN ('user', 'driver'))
);

CREATE INDEX idx_user_ratings_trip    ON public.user_ratings (trip_id);
CREATE INDEX idx_user_ratings_from    ON public.user_ratings (from_user_id);
CREATE INDEX idx_user_ratings_to      ON public.user_ratings (to_user_id);
CREATE INDEX idx_user_ratings_booking ON public.user_ratings (booking_id);
-- One rating per direction per trip
CREATE UNIQUE INDEX uq_user_ratings_trip_direction ON public.user_ratings (trip_id, from_user_id, to_user_id);

-- Notifications

CREATE TABLE IF NOT EXISTS public.user_notifications (
    id         VARCHAR(36)  PRIMARY KEY DEFAULT gen_random_uuid()::TEXT,
    user_id    VARCHAR(36)  NOT NULL,
    title      VARCHAR(300) NOT NULL,
    body       TEXT         NOT NULL DEFAULT '',
    type       VARCHAR(50)  NOT NULL DEFAULT 'general',
    channel    VARCHAR(30)  NOT NULL DEFAULT 'push',
    read       BOOLEAN      NOT NULL DEFAULT FALSE,
    data       JSONB        NOT NULL DEFAULT '{}',
    created_at TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    read_at    TIMESTAMPTZ,
    CONSTRAINT chk_notif_title CHECK (length(title) >= 1)
);

CREATE INDEX idx_user_notifications_user   ON public.user_notifications (user_id);
CREATE INDEX idx_user_notifications_unread ON public.user_notifications (user_id, read) WHERE read = FALSE;
CREATE INDEX idx_user_notifications_data   ON public.user_notifications USING GIN (data);
CREATE INDEX idx_user_notifications_comp   ON public.user_notifications (user_id, type, created_at);
