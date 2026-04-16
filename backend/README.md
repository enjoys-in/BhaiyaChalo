# BhaiyaChalo Backend

Microservice-based ride-booking platform built for Indian-city scale.

## Architecture Summary

| Principle | What It Means |
|---|---|
| Loose coupling | Services talk via contracts and events, never shared databases |
| Event-driven | Kafka is the async backbone; sync gRPC only when blocking is unavoidable |
| Isolated realtime | WebSocket fanout runs on a separate Phoenix cluster, not inside Go APIs |
| Independent scaling | Each service has its own HPA profile and resource limits |
| OSS-first | Every component is open-source and self-hostable |

## Tech Stack — What and Why

### Languages and Frameworks

| Tech | What | Why | Why Not the Alternative |
|---|---|---|---|
| Go | API gateways and all domain microservices | Fast compile, low memory, goroutine concurrency, strong ecosystem for microservices | Node.js has higher memory per connection; Java has slower cold starts |
| Elixir / Phoenix | WebSocket gateway cluster | BEAM VM handles millions of lightweight processes; native distributed pubsub across nodes | Go goroutines work but cost ~8KB each vs ~300 bytes per BEAM process at idle; sticky sessions needed in Go |

### Communication

| Tech | What | Why | Why Not the Alternative |
|---|---|---|---|
| Kafka | Async event bus for all decoupled workflows | Durable, replayable, partitioned by city; supports at-least-once with idempotent consumers | RabbitMQ lacks native replay and partition-level ordering |
| gRPC | Sync calls only for blocking decisions (booking→pricing, booking→promo) | Binary protocol, strict contracts, low latency | REST is fine for external APIs but slower and loosely typed for internal service-to-service |

### Databases

| Tech | What | Why | Why Not the Alternative |
|---|---|---|---|
| PostgreSQL 18 | Primary transactional DB for most services | ACID, mature replication, rich indexing, trusted for financial data | MySQL lacks transactional DDL and advanced JSON/GIS support |
| PostgreSQL read replicas | Offload read-heavy queries from primary | Reduces primary load without app-level sharding | — |
| PgBouncer | Connection pooler in front of PostgreSQL | Collapses hundreds of pod connections into a small pool; prevents connection storms | Pgpool-II is heavier and harder to operate for pure pooling |
| MongoDB | Trip state machine, support tickets, review moderation | Flexible document schemas for evolving state; no rigid migrations needed | PostgreSQL JSONB works but loses native document query ergonomics at scale |
| ClickHouse | Analytics — distance, earnings, heatmaps, city KPIs | Columnar storage, fast aggregates on billions of rows, native Kafka engine | Druid/Pinot are alternatives but more operationally complex for the same workload |

### Cache and Live State

| Tech | What | Why | Why Not the Alternative |
|---|---|---|---|
| Valkey | Primary cache for matching, tracking, sessions, rate limits, geo index | Redis-compatible, OSS-governed (no license risk), battle-tested data structures including GEO commands | Redis has BSL license since v7.4; Valkey is the community fork with identical API |
| Dragonfly (optional) | Secondary cache for non-critical read-heavy derived data (promo lookups, search result caching) | Higher throughput per node for simple GET/SET workloads | Not proven enough for critical mutable state like live location or matching; use only after benchmarking |

### Geo Indexing

| Tech | What | Why | Why Not the Alternative |
|---|---|---|---|
| S2 Geometry | Primary spatial index for nearest-driver discovery | Uniform cell coverage globally, clean ring expansion for radius search, strong city partitioning | GeoHash has edge distortion at cell boundaries |
| GeoHash | Secondary index for prefix-based cache keys and topic bucketing | Fast string-prefix lookups, good interop with existing tools | Quadtree is harder to maintain for high-write moving-point workloads |

### Infra and Orchestration

| Tech | What | Why | Why Not the Alternative |
|---|---|---|---|
| Docker Compose | Local development orchestration | Single command to run full infra stack locally | — |
| Kubernetes + Helm | Production deployment target | Industry standard for container orchestration, HPA, rolling updates, pod disruption budgets | Docker Swarm lacks mature autoscaling and ecosystem |
| Traefik / Kong OSS | API ingress gateway | OSS, lightweight, native K8s ingress support | Nginx Ingress works but has less dynamic config and plugin ecosystem |

### Observability (planned)

| Tech | What | Why |
|---|---|---|
| Prometheus + Alertmanager | Metrics and alerting | Standard for K8s workloads |
| Grafana | Dashboards | Connects to Prometheus, ClickHouse, and PostgreSQL |
| Loki | Log aggregation | Lightweight, label-based, pairs with Grafana |
| Jaeger / Tempo | Distributed tracing | OpenTelemetry compatible |

## Docker Compose Services

Single compose file: `docker-compose.yml`

### Core Infrastructure

| Service | Image | Exposed Port | Purpose |
|---|---|---|---|
| `postgres` | postgres:18 | 5432 | Primary database (master) |
| `postgres-replica-1` | postgres:18 | 5433 | Read replica 1 |
| `postgres-replica-2` | postgres:18 | 5434 | Read replica 2 |
| `pgbouncer-rw` | edoburu/pgbouncer | 6432 | Connection pool → primary (writes) |
| `pgbouncer-ro` | edoburu/pgbouncer | 7432 | Connection pool → replica (reads) |
| `mongo` | mongo:7 | 27017 | Document DB for trip/support/review services |
| `valkey` | valkey/valkey:7.2 | 6379 | Cache, geo state, sessions, rate limits |
| `clickhouse` | clickhouse-server:24.8 | 8123, 9000 | Analytics database |
| `kafka` | bitnami/kafka:3.7 | 9094 | Event streaming backbone (KRaft, no Zookeeper) |
| `kafka-ui` | provectuslabs/kafka-ui | 8080 | Kafka topic browser and consumer lag monitor |
| `adminer` | adminer:4 | 8088 | Lightweight DB admin UI |

### App Placeholders (profile: `apps`)

| Service | Purpose |
|---|---|
| `admin-api-gateway` | Admin-facing API edge |
| `user-api-gateway` | Rider-facing API edge |
| `driver-api-gateway` | Driver-facing API edge |
| `phoenix-websocket-gateway` | Realtime WebSocket cluster |

Start infra only: `docker compose up -d`
Start infra + apps: `docker compose --profile apps up -d`

## Folder Guide

| Folder | Contents |
|---|---|
| `apps/` | Client-facing API gateways (admin, user, driver) |
| `realtime/` | Phoenix WebSocket gateway scaffold |
| `services/` | Domain microservices grouped by bounded context |
| `libs/` | Shared Go utilities, infra adapters, types, constants |
| `contracts/` | gRPC protobuf definitions and Kafka event schemas |
| `deployments/` | Helm charts, K8s manifests, local Docker configs |
| `scripts/` | Local dev, CI, and release scripts |
| `docs/` | Architecture decisions, runbooks, API docs |

## Related Docs

- [plan.md](plan.md) — architecture decisions, build sequence, and step-by-step execution
- [directory.md](directory.md) — Go workspace directory structure
- [techstack.md](techstack.md) — service-level database and cache mapping

## PostgreSQL Schema

### Execution Order

```
000_init.sql                 ← Extensions (uuid-ossp, pgcrypto, postgis) + trigger functions
000_geo_regions.sql          ← Regions, cities, shard routing functions (run 2nd)
admin/001_iam.sql            ← Roles, permissions, policies
admin/002_documents.sql      ← Document verification
admin/003_geofences.sql      ← Geofence management (PostGIS)
admin/004_surge.sql          ← Surge policies, zones, history (geo-sharded)
admin/005_fare_pricing.sql   ← Fare configs, pricing rules, match configs
admin/006_campaigns_promos.sql ← Campaigns, promo codes
admin/007_incentives.sql     ← Driver/user incentives
admin/008_trust_safety.sql   ← Fraud signals, risk scores, audit logs
admin/009_support.sql        ← Tickets, escalations (full-text search)
admin/010_reconciliation_templates.sql ← Payment reconciliation, templates
user/001_auth_sessions.sql   ← Auth users, tokens, OTPs, sessions
user/002_profiles.sql        ← Profiles, addresses, emergency contacts (PostGIS)
user/003_bookings.sql        ← Bookings (geo-sharded, compound partitioned)
user/004_search.sql          ← Search queries & results (PostGIS)
user/005_fare_pricing.sql    ← Fare calculations, price estimates
user/006_payments.sql        ← Payments, refunds, wallets, invoices
user/007_engagement.sql      ← Promos, referrals, ratings, notifications
driver/001_profiles.sql      ← Profiles, locations, preferences (PostGIS)
driver/002_vehicles.sql      ← Vehicles
driver/003_availability.sql  ← Availability, logs (geo-sharded, compound partitioned)
driver/004_dispatch.sql      ← Dispatch offers & rounds
driver/005_matching.sql      ← Match requests (PostGIS)
driver/006_tracking_stops.sql ← Tracking sessions, multi-stop trips (PostGIS)
driver/007_earnings_payouts.sql ← Earnings (geo-sharded), summaries, payouts
000_shard_maintenance.sql    ← Partition management, archival, health checks (run last)
```

### Geo-Sharding Strategy

All high-volume, location-bound data is partitioned by region for horizontal scalability.

**Shard key flow:**
```
User opens app → lat/lng → get_city_for_location() → city_id
                                    ↓
              get_region_for_city(city_id) → region_id
                                    ↓
              INSERT INTO user_bookings (region_id, city_id, ...)
                                    ↓
              PostgreSQL routes to → user_bookings_north → user_bookings_north_2026_04
```

**5 Regions (logical shards):**

| Region | Code | Cities |
|--------|------|--------|
| North India | `north` | Delhi NCR, Jaipur, Lucknow, Chandigarh |
| South India | `south` | Bangalore, Chennai, Hyderabad, Kochi |
| West India | `west` | Mumbai, Pune, Ahmedabad, Goa |
| East India | `east` | Kolkata, Bhubaneswar, Patna, Guwahati |
| Central India | `central` | Bhopal, Indore, Nagpur, Raipur |

**Compound partitioned tables (LIST region → RANGE monthly):**

| Table | Level 1 (Region) | Level 2 (Time) |
|-------|-----------------|----------------|
| `user_bookings` | `region_id` | `created_at` |
| `driver_earnings` | `region_id` | `earned_at` |
| `admin_surge_history` | `region_id` | `calculated_at` |
| `driver_availability_logs` | `region_id` | `timestamp` |

**All 18 geo-aware tables carry `region_id`** for shard-affinity queries:

- User: profiles, bookings, search, fare_calculations, price_estimates, payments, invoices
- Driver: profiles, availability, availability_logs, dispatch_offers, dispatch_rounds, match_requests, tracking_sessions, earnings, daily/weekly_summaries, payouts
- Admin: geofences, surge_policies, surge_zones, surge_history, fare_configs, pricing_rules, campaigns, promo_codes, incentives

**Shard routing functions:**

| Function | Purpose |
|----------|---------|
| `get_region_for_city(city_id)` | Resolve city → region (shard key) |
| `get_city_for_location(lat, lng)` | Auto-detect nearest city from GPS coordinates |
| `get_region_for_location(lat, lng)` | Resolve GPS → region directly |
| `is_within_city(city_id, lat, lng)` | Check if point is inside city boundary/radius |
| `bootstrap_region_partitions(table, time_col)` | Create all region partitions for a table |
| `create_monthly_partitions(parent, col, start, end)` | Create monthly sub-partitions |
| `maintain_monthly_partitions(months_ahead)` | Auto-create upcoming partitions (pg_cron) |
| `archive_old_partitions(months_retain)` | Detach old partitions for cold storage |
| `partition_health_check()` | List all partitions with estimated row counts |
| `region_traffic_stats(start, end)` | Booking + driver counts per region |

**Future Citus migration (zero schema changes):**
```sql
SELECT create_distributed_table('user_bookings', 'region_id');
-- Same partition structure, now distributed across physical nodes.
```

### Schema Stats

| Metric | Count |
|--------|-------|
| Total tables | 67+ |
| CHECK constraints | 160+ |
| Triggers | 34 |
| PostGIS columns | 12 |
| GIN indexes | 12 |
| GiST indexes | 12 |
| Compound partitioned tables | 4 |
| Geo-sharded tables | 18 |
| Soft-delete tables | 14 |
| Seeded regions | 5 |
| Seeded cities | 20 |

## Scaffold Status

- All directories and Go module skeletons created.
- Every gateway and service has `go.mod`, `Dockerfile`, and `.env` placeholders.
- Production-grade PostgreSQL schemas with geo-sharding.
- No business code added yet — ready for implementation.
