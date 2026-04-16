-- ============================================================
-- SHARED: Geo-Sharding Foundation — Regions, Cities & Routing
-- Database: public
-- Run AFTER 000_init.sql, BEFORE any domain schema.
--
-- Strategy: LIST partition by region_id → RANGE partition by time
-- This enables future Citus/physical sharding by region with
-- zero schema changes.
-- ============================================================

-- ============================================================
-- 1. Regions (shard keys)
-- Each region maps to a logical shard. In future, each region
-- can be moved to its own PostgreSQL node via Citus or FDW.
-- ============================================================
CREATE TABLE IF NOT EXISTS public.geo_regions (
    id          VARCHAR(36)  PRIMARY KEY,
    name        VARCHAR(100) NOT NULL UNIQUE,
    code        VARCHAR(20)  NOT NULL UNIQUE,
    timezone    VARCHAR(50)  NOT NULL DEFAULT 'Asia/Kolkata',
    currency    VARCHAR(10)  NOT NULL DEFAULT 'INR',
    active      BOOLEAN      NOT NULL DEFAULT TRUE,
    created_at  TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    CONSTRAINT chk_geo_region_name CHECK (length(name) >= 2),
    CONSTRAINT chk_geo_region_code CHECK (length(code) >= 2)
);

CREATE TRIGGER set_updated_at BEFORE UPDATE ON public.geo_regions
    FOR EACH ROW EXECUTE FUNCTION trigger_set_updated_at();

-- ============================================================
-- 2. Cities (routing keys)
-- Every ride, search, dispatch happens within a city.
-- city → region mapping drives shard routing.
-- ============================================================
CREATE TABLE IF NOT EXISTS public.geo_cities (
    id              VARCHAR(36)  PRIMARY KEY,
    region_id       VARCHAR(36)  NOT NULL REFERENCES public.geo_regions(id),
    name            VARCHAR(200) NOT NULL,
    code            VARCHAR(20)  NOT NULL UNIQUE,
    state           VARCHAR(100) NOT NULL DEFAULT '',
    country         VARCHAR(10)  NOT NULL DEFAULT 'IN',
    timezone        VARCHAR(50)  NOT NULL DEFAULT 'Asia/Kolkata',
    center_lat      DOUBLE PRECISION NOT NULL DEFAULT 0,
    center_lng      DOUBLE PRECISION NOT NULL DEFAULT 0,
    -- PostGIS: auto-computed from center_lat/center_lng
    center_geom     GEOMETRY(Point, 4326) GENERATED ALWAYS AS (
                        ST_SetSRID(ST_MakePoint(center_lng, center_lat), 4326)
                    ) STORED,
    boundary        GEOMETRY(Polygon, 4326),
    radius_km       DOUBLE PRECISION NOT NULL DEFAULT 30.0,
    active          BOOLEAN      NOT NULL DEFAULT TRUE,
    launch_date     DATE,
    created_at      TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    CONSTRAINT chk_geo_city_name      CHECK (length(name) >= 2),
    CONSTRAINT chk_geo_city_code      CHECK (length(code) >= 2),
    CONSTRAINT chk_geo_city_lat       CHECK (center_lat BETWEEN -90 AND 90),
    CONSTRAINT chk_geo_city_lng       CHECK (center_lng BETWEEN -180 AND 180),
    CONSTRAINT chk_geo_city_radius    CHECK (radius_km > 0)
);

CREATE TRIGGER set_updated_at BEFORE UPDATE ON public.geo_cities
    FOR EACH ROW EXECUTE FUNCTION trigger_set_updated_at();

CREATE INDEX idx_geo_cities_region   ON public.geo_cities (region_id);
CREATE INDEX idx_geo_cities_active   ON public.geo_cities (active) WHERE active = TRUE;
CREATE INDEX idx_geo_cities_geom     ON public.geo_cities USING GIST (center_geom);
CREATE INDEX idx_geo_cities_boundary ON public.geo_cities USING GIST (boundary) WHERE boundary IS NOT NULL;
CREATE INDEX idx_geo_cities_country  ON public.geo_cities (country);

-- ============================================================
-- 3. Seed: Indian Regions
-- ============================================================
INSERT INTO public.geo_regions (id, name, code, timezone) VALUES
    ('reg_north',   'North India',   'north',   'Asia/Kolkata'),
    ('reg_south',   'South India',   'south',   'Asia/Kolkata'),
    ('reg_west',    'West India',    'west',    'Asia/Kolkata'),
    ('reg_east',    'East India',    'east',    'Asia/Kolkata'),
    ('reg_central', 'Central India', 'central', 'Asia/Kolkata')
ON CONFLICT (id) DO NOTHING;

-- ============================================================
-- 4. Seed: Major Indian Cities
-- ============================================================
INSERT INTO public.geo_cities (id, region_id, name, code, state, center_lat, center_lng, radius_km) VALUES
    -- North
    ('city_del', 'reg_north', 'Delhi NCR',     'DEL', 'Delhi',            28.6139,  77.2090, 50.0),
    ('city_jai', 'reg_north', 'Jaipur',        'JAI', 'Rajasthan',        26.9124,  75.7873, 25.0),
    ('city_lko', 'reg_north', 'Lucknow',       'LKO', 'Uttar Pradesh',    26.8467,  80.9462, 20.0),
    ('city_chd', 'reg_north', 'Chandigarh',    'CHD', 'Chandigarh',       30.7333,  76.7794, 15.0),
    -- South
    ('city_blr', 'reg_south', 'Bangalore',     'BLR', 'Karnataka',        12.9716,  77.5946, 40.0),
    ('city_chn', 'reg_south', 'Chennai',       'CHN', 'Tamil Nadu',       13.0827,  80.2707, 35.0),
    ('city_hyd', 'reg_south', 'Hyderabad',     'HYD', 'Telangana',        17.3850,  78.4867, 35.0),
    ('city_koc', 'reg_south', 'Kochi',         'KOC', 'Kerala',            9.9312,  76.2673, 15.0),
    -- West
    ('city_mum', 'reg_west',  'Mumbai',        'MUM', 'Maharashtra',      19.0760,  72.8777, 40.0),
    ('city_pun', 'reg_west',  'Pune',          'PUN', 'Maharashtra',      18.5204,  73.8567, 25.0),
    ('city_amd', 'reg_west',  'Ahmedabad',     'AMD', 'Gujarat',          23.0225,  72.5714, 25.0),
    ('city_goa', 'reg_west',  'Goa',           'GOA', 'Goa',              15.2993,  74.1240, 20.0),
    -- East
    ('city_kol', 'reg_east',  'Kolkata',       'KOL', 'West Bengal',      22.5726,  88.3639, 30.0),
    ('city_bbs', 'reg_east',  'Bhubaneswar',   'BBS', 'Odisha',           20.2961,  85.8245, 15.0),
    ('city_pat', 'reg_east',  'Patna',         'PAT', 'Bihar',            25.6093,  85.1376, 15.0),
    ('city_ghy', 'reg_east',  'Guwahati',      'GHY', 'Assam',            26.1445,  91.7362, 15.0),
    -- Central
    ('city_bpl', 'reg_central','Bhopal',       'BPL', 'Madhya Pradesh',   23.2599,  77.4126, 15.0),
    ('city_idr', 'reg_central','Indore',       'IDR', 'Madhya Pradesh',   22.7196,  75.8577, 15.0),
    ('city_nag', 'reg_central','Nagpur',       'NAG', 'Maharashtra',      21.1458,  79.0882, 20.0),
    ('city_rai', 'reg_central','Raipur',       'RAI', 'Chhattisgarh',     21.2514,  81.6296, 15.0)
ON CONFLICT (id) DO NOTHING;

-- ============================================================
-- 5. Shard Routing Functions
-- ============================================================

-- Resolve city_id → region_id (used by app layer for shard routing)
CREATE OR REPLACE FUNCTION public.get_region_for_city(p_city_id VARCHAR)
RETURNS VARCHAR AS $$
    SELECT region_id FROM public.geo_cities WHERE id = p_city_id;
$$ LANGUAGE sql STABLE;

-- Resolve lat/lng → nearest city_id (for auto-detecting city from coordinates)
CREATE OR REPLACE FUNCTION public.get_city_for_location(p_lat DOUBLE PRECISION, p_lng DOUBLE PRECISION)
RETURNS VARCHAR AS $$
    SELECT id
    FROM public.geo_cities
    WHERE active = TRUE
    ORDER BY center_geom <-> ST_SetSRID(ST_MakePoint(p_lng, p_lat), 4326)
    LIMIT 1;
$$ LANGUAGE sql STABLE;

-- Resolve lat/lng → region_id (combines both lookups)
CREATE OR REPLACE FUNCTION public.get_region_for_location(p_lat DOUBLE PRECISION, p_lng DOUBLE PRECISION)
RETURNS VARCHAR AS $$
    SELECT gc.region_id
    FROM public.geo_cities gc
    WHERE gc.active = TRUE
    ORDER BY gc.center_geom <-> ST_SetSRID(ST_MakePoint(p_lng, p_lat), 4326)
    LIMIT 1;
$$ LANGUAGE sql STABLE;

-- Check if a point is within a city's service boundary
CREATE OR REPLACE FUNCTION public.is_within_city(p_city_id VARCHAR, p_lat DOUBLE PRECISION, p_lng DOUBLE PRECISION)
RETURNS BOOLEAN AS $$
    SELECT CASE
        WHEN gc.boundary IS NOT NULL THEN
            ST_Contains(gc.boundary, ST_SetSRID(ST_MakePoint(p_lng, p_lat), 4326))
        ELSE
            ST_DWithin(
                gc.center_geom::geography,
                ST_SetSRID(ST_MakePoint(p_lng, p_lat), 4326)::geography,
                gc.radius_km * 1000
            )
    END
    FROM public.geo_cities gc
    WHERE gc.id = p_city_id;
$$ LANGUAGE sql STABLE;

-- ============================================================
-- 6. Auto-partition helper for compound partitioning
--    Creates monthly sub-partitions within a region partition.
--
-- Usage:
--   SELECT create_monthly_partitions(
--       'public.user_bookings_north',  -- parent (region partition)
--       'created_at',                   -- time column
--       '2024-01-01',                   -- start month
--       '2025-01-01'                    -- end month (exclusive)
--   );
-- ============================================================
CREATE OR REPLACE FUNCTION public.create_monthly_partitions(
    p_parent_table TEXT,
    p_time_column  TEXT,
    p_start_date   DATE,
    p_end_date     DATE
) RETURNS VOID AS $$
DECLARE
    v_month     DATE;
    v_next      DATE;
    v_part_name TEXT;
BEGIN
    v_month := date_trunc('month', p_start_date)::DATE;
    WHILE v_month < p_end_date LOOP
        v_next := (v_month + INTERVAL '1 month')::DATE;
        v_part_name := p_parent_table || '_' || to_char(v_month, 'YYYY_MM');

        EXECUTE format(
            'CREATE TABLE IF NOT EXISTS %I PARTITION OF %s FOR VALUES FROM (%L) TO (%L)',
            v_part_name, p_parent_table, v_month, v_next
        );

        v_month := v_next;
    END LOOP;
END;
$$ LANGUAGE plpgsql;

-- ============================================================
-- 7. Auto-create region partition for a given table
--
-- Usage:
--   SELECT create_region_partition(
--       'public.user_bookings',  -- parent table
--       'reg_north',             -- region_id value
--       'created_at'             -- time column for sub-partitioning
--   );
-- ============================================================
CREATE OR REPLACE FUNCTION public.create_region_partition(
    p_parent_table TEXT,
    p_region_id    TEXT,
    p_time_column  TEXT DEFAULT NULL
) RETURNS TEXT AS $$
DECLARE
    v_region_code TEXT;
    v_part_name   TEXT;
BEGIN
    SELECT code INTO v_region_code FROM public.geo_regions WHERE id = p_region_id;
    IF v_region_code IS NULL THEN
        RAISE EXCEPTION 'Region % not found', p_region_id;
    END IF;

    v_part_name := p_parent_table || '_' || v_region_code;

    IF p_time_column IS NOT NULL THEN
        EXECUTE format(
            'CREATE TABLE IF NOT EXISTS %s PARTITION OF %s FOR VALUES IN (%L) PARTITION BY RANGE (%I)',
            v_part_name, p_parent_table, p_region_id, p_time_column
        );
        -- Create a default sub-partition for the region
        EXECUTE format(
            'CREATE TABLE IF NOT EXISTS %s PARTITION OF %s DEFAULT',
            v_part_name || '_default', v_part_name
        );
    ELSE
        EXECUTE format(
            'CREATE TABLE IF NOT EXISTS %s PARTITION OF %s FOR VALUES IN (%L)',
            v_part_name, p_parent_table, p_region_id
        );
    END IF;

    RETURN v_part_name;
END;
$$ LANGUAGE plpgsql;

-- ============================================================
-- 8. Bootstrap all region partitions for a table
--
-- Usage:
--   SELECT bootstrap_region_partitions('public.user_bookings', 'created_at');
-- ============================================================
CREATE OR REPLACE FUNCTION public.bootstrap_region_partitions(
    p_parent_table TEXT,
    p_time_column  TEXT DEFAULT NULL
) RETURNS VOID AS $$
DECLARE
    v_region RECORD;
BEGIN
    FOR v_region IN SELECT id FROM public.geo_regions WHERE active = TRUE LOOP
        PERFORM public.create_region_partition(p_parent_table, v_region.id, p_time_column);
    END LOOP;
END;
$$ LANGUAGE plpgsql;
