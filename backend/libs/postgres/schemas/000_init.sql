-- ============================================================
-- SHARED: Extensions & Utility Functions
-- Database: public
-- Run this FIRST before any other schema file.
-- ============================================================

-- Extensions
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "pgcrypto";
CREATE EXTENSION IF NOT EXISTS "postgis";

-- ============================================================
-- Auto-update updated_at on every row modification
-- Usage:
--   CREATE TRIGGER set_updated_at BEFORE UPDATE ON <table>
--   FOR EACH ROW EXECUTE FUNCTION trigger_set_updated_at();
-- ============================================================
CREATE OR REPLACE FUNCTION public.trigger_set_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- ============================================================
-- Soft-delete helper: set deleted_at instead of hard DELETE
-- Usage:
--   CREATE TRIGGER soft_delete BEFORE DELETE ON <table>
--   FOR EACH ROW EXECUTE FUNCTION trigger_soft_delete();
-- ============================================================
CREATE OR REPLACE FUNCTION public.trigger_soft_delete()
RETURNS TRIGGER AS $$
BEGIN
    UPDATE current_catalog SET deleted_at = NOW() WHERE ctid = OLD.ctid;
    RETURN NULL;  -- suppress the DELETE
END;
$$ LANGUAGE plpgsql;
