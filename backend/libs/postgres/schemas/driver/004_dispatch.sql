-- ============================================================
-- DRIVER SCHEMA: Dispatch Offers & Rounds
-- Database: public | Prefix: driver_
-- ============================================================

CREATE TYPE driver_offer_status AS ENUM ('pending', 'accepted', 'rejected', 'expired');
CREATE TYPE driver_round_status AS ENUM ('active', 'completed', 'failed');

CREATE TABLE IF NOT EXISTS public.driver_dispatch_offers (
    id               VARCHAR(36)         PRIMARY KEY DEFAULT gen_random_uuid()::TEXT,
    booking_id       VARCHAR(36)         NOT NULL,
    driver_id        VARCHAR(36)         NOT NULL,
    city_id          VARCHAR(36)         NOT NULL,
    status           driver_offer_status NOT NULL DEFAULT 'pending',
    offer_expires_at TIMESTAMPTZ         NOT NULL,
    responded_at     TIMESTAMPTZ,
    created_at       TIMESTAMPTZ         NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_driver_dispatch_offers_booking   ON public.driver_dispatch_offers (booking_id);
CREATE INDEX idx_driver_dispatch_offers_driver    ON public.driver_dispatch_offers (driver_id);
CREATE INDEX idx_driver_dispatch_offers_status    ON public.driver_dispatch_offers (status);
CREATE INDEX idx_driver_dispatch_offers_composite ON public.driver_dispatch_offers (booking_id, driver_id, status);
CREATE INDEX idx_driver_dispatch_offers_pending   ON public.driver_dispatch_offers (driver_id, offer_expires_at)
    WHERE status = 'pending';

CREATE TABLE IF NOT EXISTS public.driver_dispatch_rounds (
    id                   VARCHAR(36)         PRIMARY KEY DEFAULT gen_random_uuid()::TEXT,
    booking_id           VARCHAR(36)         NOT NULL,
    round_number         INT                 NOT NULL DEFAULT 1,
    candidate_driver_ids JSONB               NOT NULL DEFAULT '[]',
    status               driver_round_status NOT NULL DEFAULT 'active',
    created_at           TIMESTAMPTZ         NOT NULL DEFAULT NOW(),
    CONSTRAINT chk_dispatch_round_num CHECK (round_number > 0)
);

CREATE INDEX idx_driver_dispatch_rounds_booking    ON public.driver_dispatch_rounds (booking_id);
CREATE INDEX idx_driver_dispatch_rounds_status     ON public.driver_dispatch_rounds (status);
CREATE INDEX idx_driver_dispatch_rounds_candidates ON public.driver_dispatch_rounds USING GIN (candidate_driver_ids);
CREATE UNIQUE INDEX uq_driver_dispatch_rounds_booking_round ON public.driver_dispatch_rounds (booking_id, round_number);
