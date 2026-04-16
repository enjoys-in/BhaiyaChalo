-- ============================================================
-- SHARED: Partition Maintenance & Cron Jobs
-- Database: public
-- Run AFTER all domain schemas are created.
--
-- This file provides:
--   1. Auto-create upcoming monthly partitions (run via pg_cron)
--   2. Detach/archive old partitions (cold storage)
--   3. Partition health checks
-- ============================================================

-- ============================================================
-- 1. Auto-create monthly partitions for the next N months
--    across all compound-partitioned tables.
--
-- Run monthly via pg_cron:
--   SELECT cron.schedule('partition-maintenance', '0 0 25 * *',
--       $$SELECT maintain_monthly_partitions(3)$$);
-- ============================================================
CREATE OR REPLACE FUNCTION public.maintain_monthly_partitions(
    p_months_ahead INT DEFAULT 3
) RETURNS TABLE(table_name TEXT, partition_name TEXT, range_start DATE, range_end DATE) AS $$
DECLARE
    v_region       RECORD;
    v_month        DATE;
    v_next         DATE;
    v_part_name    TEXT;
    v_parent       TEXT;
    v_tables       TEXT[] := ARRAY[
        'public.user_bookings',
        'public.driver_earnings',
        'public.admin_surge_history',
        'public.driver_availability_logs'
    ];
    v_time_cols    TEXT[] := ARRAY[
        'created_at',
        'earned_at',
        'calculated_at',
        'timestamp'
    ];
    v_idx          INT;
BEGIN
    FOR v_region IN SELECT id, code FROM public.geo_regions WHERE active = TRUE LOOP
        FOR v_idx IN 1..array_length(v_tables, 1) LOOP
            v_parent := v_tables[v_idx] || '_' || v_region.code;

            -- Ensure region partition exists
            BEGIN
                EXECUTE format(
                    'CREATE TABLE IF NOT EXISTS %s PARTITION OF %s FOR VALUES IN (%L) PARTITION BY RANGE (%I)',
                    v_parent, v_tables[v_idx], v_region.id, v_time_cols[v_idx]
                );
            EXCEPTION WHEN duplicate_table THEN
                NULL; -- already exists
            END;

            v_month := date_trunc('month', NOW())::DATE;
            WHILE v_month < (date_trunc('month', NOW()) + (p_months_ahead || ' months')::INTERVAL)::DATE LOOP
                v_next := (v_month + INTERVAL '1 month')::DATE;
                v_part_name := v_parent || '_' || to_char(v_month, 'YYYY_MM');

                BEGIN
                    EXECUTE format(
                        'CREATE TABLE IF NOT EXISTS %s PARTITION OF %s FOR VALUES FROM (%L) TO (%L)',
                        v_part_name, v_parent, v_month, v_next
                    );

                    table_name     := v_tables[v_idx];
                    partition_name := v_part_name;
                    range_start    := v_month;
                    range_end      := v_next;
                    RETURN NEXT;
                EXCEPTION WHEN duplicate_table THEN
                    NULL; -- already exists
                END;

                v_month := v_next;
            END LOOP;
        END LOOP;
    END LOOP;
END;
$$ LANGUAGE plpgsql;

-- ============================================================
-- 2. Archive old partitions (detach partitions older than N months)
--    Detached tables can be moved to cold storage (S3 via pg_dump).
--
-- Usage:
--   SELECT archive_old_partitions(6);  -- detach partitions > 6 months old
--
-- WARNING: This DETACHES partitions. Data is NOT deleted but
-- queries won't hit those tables. Run pg_dump on detached tables
-- before dropping them.
-- ============================================================
CREATE OR REPLACE FUNCTION public.archive_old_partitions(
    p_months_retain INT DEFAULT 6
) RETURNS TABLE(detached_partition TEXT, detached_at TIMESTAMPTZ) AS $$
DECLARE
    v_cutoff DATE;
    v_part   RECORD;
BEGIN
    v_cutoff := (date_trunc('month', NOW()) - (p_months_retain || ' months')::INTERVAL)::DATE;

    -- Find all leaf partitions with range bounds before cutoff
    FOR v_part IN
        SELECT
            c.relname AS partition_name,
            n.nspname AS schema_name,
            pg_get_expr(c.relpartbound, c.oid) AS bound_expr
        FROM pg_class c
        JOIN pg_namespace n ON n.oid = c.relnamespace
        JOIN pg_inherits i ON i.inhrelid = c.oid
        WHERE c.relkind = 'r'
          AND n.nspname = 'public'
          AND c.relname ~ '_\d{4}_\d{2}$'  -- matches _YYYY_MM suffix
    LOOP
        -- Extract the year/month from partition name
        DECLARE
            v_year  INT;
            v_month INT;
            v_date  DATE;
            v_parent_name TEXT;
        BEGIN
            v_year  := substring(v_part.partition_name FROM '_(\d{4})_\d{2}$')::INT;
            v_month := substring(v_part.partition_name FROM '_\d{4}_(\d{2})$')::INT;
            v_date  := make_date(v_year, v_month, 1);

            IF v_date < v_cutoff THEN
                -- Find the parent table
                SELECT p.relname INTO v_parent_name
                FROM pg_inherits i
                JOIN pg_class p ON p.oid = i.inhparent
                WHERE i.inhrelid = (
                    SELECT oid FROM pg_class
                    WHERE relname = v_part.partition_name
                      AND relnamespace = (SELECT oid FROM pg_namespace WHERE nspname = 'public')
                );

                EXECUTE format(
                    'ALTER TABLE public.%I DETACH PARTITION public.%I',
                    v_parent_name, v_part.partition_name
                );

                detached_partition := v_part.partition_name;
                detached_at        := NOW();
                RETURN NEXT;
            END IF;
        EXCEPTION WHEN OTHERS THEN
            RAISE NOTICE 'Could not detach %: %', v_part.partition_name, SQLERRM;
        END;
    END LOOP;
END;
$$ LANGUAGE plpgsql;

-- ============================================================
-- 3. Partition health check — list all partitioned tables
--    with their partition count and estimated row counts.
--
-- Usage:
--   SELECT * FROM partition_health_check();
-- ============================================================
CREATE OR REPLACE FUNCTION public.partition_health_check()
RETURNS TABLE(
    parent_table   TEXT,
    region         TEXT,
    partition_name TEXT,
    estimated_rows BIGINT
) AS $$
BEGIN
    RETURN QUERY
    SELECT
        p.relname::TEXT    AS parent_table,
        CASE
            WHEN c.relname ~ '_north'   THEN 'north'
            WHEN c.relname ~ '_south'   THEN 'south'
            WHEN c.relname ~ '_west'    THEN 'west'
            WHEN c.relname ~ '_east'    THEN 'east'
            WHEN c.relname ~ '_central' THEN 'central'
            ELSE 'default'
        END                AS region,
        c.relname::TEXT    AS partition_name,
        c.reltuples::BIGINT AS estimated_rows
    FROM pg_inherits i
    JOIN pg_class c ON c.oid = i.inhrelid
    JOIN pg_class p ON p.oid = i.inhparent
    JOIN pg_namespace n ON n.oid = p.relnamespace
    WHERE n.nspname = 'public'
    ORDER BY p.relname, c.relname;
END;
$$ LANGUAGE plpgsql;

-- ============================================================
-- 4. Region traffic stats — useful for rebalancing decisions
--
-- Usage:
--   SELECT * FROM region_traffic_stats('2024-01-01', '2024-02-01');
-- ============================================================
CREATE OR REPLACE FUNCTION public.region_traffic_stats(
    p_start DATE,
    p_end   DATE
) RETURNS TABLE(
    region_id      TEXT,
    region_name    TEXT,
    booking_count  BIGINT,
    active_drivers BIGINT
) AS $$
BEGIN
    RETURN QUERY
    SELECT
        gr.id::TEXT,
        gr.name::TEXT,
        COALESCE(b.cnt, 0)::BIGINT,
        COALESCE(d.cnt, 0)::BIGINT
    FROM public.geo_regions gr
    LEFT JOIN (
        SELECT ub.region_id AS rid, COUNT(*) AS cnt
        FROM public.user_bookings ub
        WHERE ub.created_at >= p_start AND ub.created_at < p_end
        GROUP BY ub.region_id
    ) b ON b.rid = gr.id
    LEFT JOIN (
        SELECT dp.region_id AS rid, COUNT(*) AS cnt
        FROM public.driver_profiles dp
        WHERE dp.status = 'active'
        GROUP BY dp.region_id
    ) d ON d.rid = gr.id
    ORDER BY gr.name;
END;
$$ LANGUAGE plpgsql;
