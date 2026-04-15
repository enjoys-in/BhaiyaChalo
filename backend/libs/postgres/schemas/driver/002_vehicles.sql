-- ============================================================
-- DRIVER SCHEMA: Vehicles
-- Database: public | Prefix: driver_
-- ============================================================

CREATE TYPE driver_vehicle_type   AS ENUM ('auto', 'mini', 'sedan', 'suv', 'premium');
CREATE TYPE driver_vehicle_status AS ENUM ('pending', 'approved', 'rejected', 'expired');

CREATE TABLE IF NOT EXISTS public.driver_vehicles (
    id               VARCHAR(36)           PRIMARY KEY DEFAULT gen_random_uuid()::TEXT,
    driver_id        VARCHAR(36)           NOT NULL,
    make             VARCHAR(100)          NOT NULL DEFAULT '',
    model            VARCHAR(100)          NOT NULL DEFAULT '',
    year             INT                   NOT NULL DEFAULT 0,
    color            VARCHAR(50)           NOT NULL DEFAULT '',
    plate_number     VARCHAR(30)           NOT NULL UNIQUE,
    vehicle_type     driver_vehicle_type   NOT NULL DEFAULT 'mini',
    insurance_expiry TIMESTAMPTZ           NOT NULL,
    fitness_expiry   TIMESTAMPTZ           NOT NULL,
    status           driver_vehicle_status NOT NULL DEFAULT 'pending',
    created_at       TIMESTAMPTZ           NOT NULL DEFAULT NOW(),
    updated_at       TIMESTAMPTZ           NOT NULL DEFAULT NOW(),
    deleted_at       TIMESTAMPTZ,
    CONSTRAINT chk_vehicle_year  CHECK (year >= 1990 AND year <= 2100),
    CONSTRAINT chk_vehicle_plate CHECK (length(plate_number) >= 4),
    CONSTRAINT chk_vehicle_make  CHECK (length(make)  >= 1),
    CONSTRAINT chk_vehicle_model CHECK (length(model) >= 1)
);

CREATE TRIGGER set_updated_at BEFORE UPDATE ON public.driver_vehicles
    FOR EACH ROW EXECUTE FUNCTION trigger_set_updated_at();

CREATE INDEX idx_driver_vehicles_driver  ON public.driver_vehicles (driver_id);
CREATE INDEX idx_driver_vehicles_plate   ON public.driver_vehicles (plate_number);
CREATE INDEX idx_driver_vehicles_status  ON public.driver_vehicles (status);
CREATE INDEX idx_driver_vehicles_expiry  ON public.driver_vehicles (insurance_expiry, fitness_expiry);
CREATE INDEX idx_driver_vehicles_deleted ON public.driver_vehicles (deleted_at) WHERE deleted_at IS NULL;
CREATE INDEX idx_driver_vehicles_comp    ON public.driver_vehicles (driver_id, status);
