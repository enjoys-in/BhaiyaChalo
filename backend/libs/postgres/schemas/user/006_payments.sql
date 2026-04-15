-- ============================================================
-- USER SCHEMA: Payments, Refunds, Wallets, Invoices
-- Database: public | Prefix: user_
-- ============================================================

CREATE TYPE user_payment_method AS ENUM ('card', 'upi', 'wallet', 'cash');
CREATE TYPE user_payment_status AS ENUM ('initiated', 'authorized', 'captured', 'failed', 'refunded');
CREATE TYPE user_refund_status  AS ENUM ('pending', 'processed', 'failed');

CREATE TABLE IF NOT EXISTS public.user_payments (
    id             VARCHAR(36)         PRIMARY KEY DEFAULT gen_random_uuid()::TEXT,
    booking_id     VARCHAR(36)         NOT NULL,
    user_id        VARCHAR(36)         NOT NULL,
    amount         NUMERIC(12,2)       NOT NULL DEFAULT 0,
    currency       VARCHAR(10)         NOT NULL DEFAULT 'INR',
    method         user_payment_method NOT NULL DEFAULT 'cash',
    gateway_id     VARCHAR(100)        NOT NULL DEFAULT '',
    gateway_status VARCHAR(50)         NOT NULL DEFAULT '',
    status         user_payment_status NOT NULL DEFAULT 'initiated',
    failure_reason TEXT                NOT NULL DEFAULT '',
    created_at     TIMESTAMPTZ         NOT NULL DEFAULT NOW(),
    updated_at     TIMESTAMPTZ         NOT NULL DEFAULT NOW(),
    CONSTRAINT chk_payment_amount CHECK (amount >= 0)
);

CREATE TRIGGER set_updated_at BEFORE UPDATE ON public.user_payments
    FOR EACH ROW EXECUTE FUNCTION trigger_set_updated_at();

CREATE INDEX idx_user_payments_booking   ON public.user_payments (booking_id);
CREATE INDEX idx_user_payments_user      ON public.user_payments (user_id);
CREATE INDEX idx_user_payments_status    ON public.user_payments (status);
CREATE INDEX idx_user_payments_composite ON public.user_payments (user_id, status, created_at);

CREATE TABLE IF NOT EXISTS public.user_refunds (
    id         VARCHAR(36)        PRIMARY KEY DEFAULT gen_random_uuid()::TEXT,
    payment_id VARCHAR(36)        NOT NULL,
    amount     NUMERIC(12,2)      NOT NULL DEFAULT 0,
    reason     TEXT               NOT NULL DEFAULT '',
    status     user_refund_status NOT NULL DEFAULT 'pending',
    created_at TIMESTAMPTZ        NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ        NOT NULL DEFAULT NOW(),
    CONSTRAINT chk_refund_amount CHECK (amount > 0)
);

CREATE TRIGGER set_updated_at BEFORE UPDATE ON public.user_refunds
    FOR EACH ROW EXECUTE FUNCTION trigger_set_updated_at();

CREATE INDEX idx_user_refunds_payment ON public.user_refunds (payment_id);
CREATE INDEX idx_user_refunds_status  ON public.user_refunds (status);

-- Wallets

CREATE TYPE user_wallet_status   AS ENUM ('active', 'frozen');
CREATE TYPE user_transaction_type AS ENUM ('credit', 'debit');

CREATE TABLE IF NOT EXISTS public.user_wallets (
    id         VARCHAR(36)        PRIMARY KEY DEFAULT gen_random_uuid()::TEXT,
    user_id    VARCHAR(36)        NOT NULL UNIQUE,
    balance    NUMERIC(12,2)      NOT NULL DEFAULT 0,
    currency   VARCHAR(10)        NOT NULL DEFAULT 'INR',
    status     user_wallet_status NOT NULL DEFAULT 'active',
    created_at TIMESTAMPTZ        NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ        NOT NULL DEFAULT NOW(),
    CONSTRAINT chk_wallet_balance CHECK (balance >= 0)
);

CREATE TRIGGER set_updated_at BEFORE UPDATE ON public.user_wallets
    FOR EACH ROW EXECUTE FUNCTION trigger_set_updated_at();

CREATE INDEX idx_user_wallets_user ON public.user_wallets (user_id);

CREATE TABLE IF NOT EXISTS public.user_wallet_transactions (
    id             VARCHAR(36)           PRIMARY KEY DEFAULT gen_random_uuid()::TEXT,
    wallet_id      VARCHAR(36)           NOT NULL,
    type           user_transaction_type NOT NULL,
    amount         NUMERIC(12,2)         NOT NULL DEFAULT 0,
    reference      VARCHAR(100)          NOT NULL DEFAULT '',
    description    TEXT                  NOT NULL DEFAULT '',
    balance_before NUMERIC(12,2)         NOT NULL DEFAULT 0,
    balance_after  NUMERIC(12,2)         NOT NULL DEFAULT 0,
    created_at     TIMESTAMPTZ           NOT NULL DEFAULT NOW(),
    CONSTRAINT chk_wallet_txn_amount      CHECK (amount > 0),
    CONSTRAINT chk_wallet_txn_bal_before  CHECK (balance_before >= 0),
    CONSTRAINT chk_wallet_txn_bal_after   CHECK (balance_after  >= 0)
);

CREATE INDEX idx_user_wallet_txns_wallet ON public.user_wallet_transactions (wallet_id);
CREATE INDEX idx_user_wallet_txns_ref    ON public.user_wallet_transactions (reference) WHERE reference != '';
CREATE INDEX idx_user_wallet_txns_comp   ON public.user_wallet_transactions (wallet_id, created_at);

-- Invoices

CREATE TYPE user_invoice_status AS ENUM ('generated', 'sent', 'paid');

CREATE TABLE IF NOT EXISTS public.user_invoices (
    id             VARCHAR(36)         PRIMARY KEY DEFAULT gen_random_uuid()::TEXT,
    booking_id     VARCHAR(36)         NOT NULL,
    user_id        VARCHAR(36)         NOT NULL,
    invoice_number VARCHAR(50)         NOT NULL UNIQUE,
    sub_total      NUMERIC(12,2)       NOT NULL DEFAULT 0,
    tax_amount     NUMERIC(12,2)       NOT NULL DEFAULT 0,
    total_amount   NUMERIC(12,2)       NOT NULL DEFAULT 0,
    currency       VARCHAR(10)         NOT NULL DEFAULT 'INR',
    status         user_invoice_status NOT NULL DEFAULT 'generated',
    generated_at   TIMESTAMPTZ         NOT NULL DEFAULT NOW(),
    sent_at        TIMESTAMPTZ,
    CONSTRAINT chk_invoice_subtotal CHECK (sub_total    >= 0),
    CONSTRAINT chk_invoice_tax      CHECK (tax_amount   >= 0),
    CONSTRAINT chk_invoice_total    CHECK (total_amount  >= 0)
);

CREATE INDEX idx_user_invoices_booking ON public.user_invoices (booking_id);
CREATE INDEX idx_user_invoices_user    ON public.user_invoices (user_id);
CREATE INDEX idx_user_invoices_number  ON public.user_invoices (invoice_number);
CREATE INDEX idx_user_invoices_comp    ON public.user_invoices (user_id, status);

CREATE TABLE IF NOT EXISTS public.user_invoice_items (
    id          VARCHAR(36)   PRIMARY KEY DEFAULT gen_random_uuid()::TEXT,
    invoice_id  VARCHAR(36)   NOT NULL,
    description TEXT          NOT NULL DEFAULT '',
    quantity    INT           NOT NULL DEFAULT 1,
    unit_price  NUMERIC(12,2) NOT NULL DEFAULT 0,
    amount      NUMERIC(12,2) NOT NULL DEFAULT 0,
    CONSTRAINT chk_inv_item_qty    CHECK (quantity   > 0),
    CONSTRAINT chk_inv_item_price  CHECK (unit_price >= 0),
    CONSTRAINT chk_inv_item_amount CHECK (amount     >= 0)
);

CREATE INDEX idx_user_invoice_items_invoice ON public.user_invoice_items (invoice_id);
