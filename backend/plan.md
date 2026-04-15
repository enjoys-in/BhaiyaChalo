# Ride Booking Backend Production Plan

## 1. Mission
Build a ride-booking platform using microservices, loose coupling, and event-driven architecture that can scale to large Indian city traffic patterns (rush hours, rain spikes, festivals, and stadium exits) without redesign.

## 2. Non-Negotiable Architecture Principles
- Event first: every domain change emits events to Kafka.
- Loose coupling: no direct database sharing across services.
- Contract-driven integration: gRPC for blocking decisions only, Kafka for async workflows.
- City-aware design: every hot path and topic uses `city_id` partitioning.
- Stateless compute: services are horizontally scalable and restart-safe.

## 3. Communication Contract (Sync vs Async)
Use synchronous gRPC only for request-path decisions that block ride flow:
- Booking -> Pricing (fare required before confirmation)
- Booking -> Promo (discount application before final fare)

Use Kafka for all decoupled workflows:
- Trip lifecycle transitions
- Driver status and availability propagation
- Payment triggers and retries
- Notifications
- Analytics ingestion and projections
- Post-trip tasks (rating, receipts, support hooks)

Rule:
- No service calls another service database.

## 4. Production-Grade Service Capacity Baseline
This baseline is a strong starting point for Indian metro launch traffic. Tune per city tier after load tests.

| Service | Role | Min Pods | Max Pods | CPU Request/Limit | Memory Request/Limit | Autoscaling Signal |
|---|---|---:|---:|---|---|---|
| API Gateway (Kong/Traefik) | North-south routing | 3 | 20 | 200m / 1000m | 256Mi / 512Mi | CPU > 60% |
| WebSocket Gateway | Tracking push fanout | 3 | 30 | 200m / 1000m | 512Mi / 2Gi | Memory > 65% |
| Matching | Redis geo-query + assignment | 4 | 40 | 500m / 2000m | 256Mi / 1Gi | CPU > 50% |
| Booking | Write-heavy booking orchestration | 3 | 20 | 200m / 1000m | 256Mi / 512Mi | CPU > 60% |
| Tracking | GPS ingestion and state write | 3 | 30 | 200m / 800m | 256Mi / 512Mi | CPU > 55% |
| Auth | JWT validation + identity flows | 3 | 15 | 100m / 500m | 128Mi / 256Mi | CPU > 65% |
| Trip | Trip state machine | 2 | 12 | 200m / 800m | 256Mi / 512Mi | CPU > 60% |
| Payment | External gateway integration | 2 | 10 | 100m / 500m | 128Mi / 256Mi | CPU > 60% |
| Pricing | Surge and fare computation | 2 | 10 | 200m / 800m | 128Mi / 256Mi | CPU > 60% |
| Driver | Profile/status management | 2 | 8 | 100m / 500m | 128Mi / 256Mi | CPU > 65% |
| User | Rider profile and preferences | 2 | 8 | 100m / 500m | 128Mi / 256Mi | CPU > 65% |
| Notification | Kafka consumer for push/SMS/email | 2 | 12 | 100m / 400m | 128Mi / 256Mi | Kafka lag > 50 |
| Rating | Async post-trip feedback | 2 | 6 | 100m / 300m | 128Mi / 256Mi | Kafka lag > 100 |
| Promo | Read-heavy campaign service | 2 | 6 | 100m / 300m | 128Mi / 256Mi | CPU > 70% |
| Support | Low-traffic support workflows | 1 | 4 | 100m / 300m | 128Mi / 256Mi | CPU > 70% |
| Analytics | Stream processing + warehouse load | 2 | 12 | 300m / 1200m | 512Mi / 2Gi | Kafka lag + batch duration |

## 5. Data and Storage Strategy
- PostgreSQL per service: Auth, User, Driver, Booking, Pricing, Promo, Rating, Payment.
- MongoDB per service: Trip, Support (flexible state transitions and documents).
- Redis: live geo state, availability, sessions, rate limits, short-lived projections.
- ClickHouse: BI dashboards, aggregates, cohort and city-level analytics.

## 6. Real-Time Tracking Design (Low Latency + Durability)
- Driver GPS updates arrive at WebSocket Gateway.
- Tracking service persists latest coordinates in Redis GEO.
- Rider sockets receive immediate updates from gateway fanout.
- Kafka receives periodic snapshots for replay, auditing, and analytics.
- Hot path uses Redis and sockets; analytics path uses Kafka and ClickHouse.

## 7. Kubernetes and Helm Production Standards
- Separate Helm values by environment and city tier.
- Enforce pod anti-affinity for critical services (Gateway, Matching, Tracking).
- Define PodDisruptionBudgets for user-facing and critical path services.
- Use HPA for CPU/memory services and KEDA for Kafka consumers.
- Keep readiness strict and startup/liveness probes service-specific.
- Use rolling update with safe `maxUnavailable` for zero-downtime deployments.

## 8. Database and Connection Protection
- PgBouncer required for all PostgreSQL-backed services.
- Use transaction pooling mode for high pod count.
- Cap per-pod DB pool sizes to prevent connection storms.
- Add circuit breakers and timeout budgets for external payment dependencies.

## 9. Kafka Design for Indian-Scale Spikes
- Partition topics by `city_id` and stable key (`driver_id`, `trip_id`, or `booking_id`).
- Use retry topics and dead-letter topics for poison messages.
- Define idempotent consumers for at-least-once delivery semantics.
- Monitor consumer lag by city and service.
- Keep event schema versioned and backward compatible.

## 10. City Rollout Strategy
Phase 1 (first launch city):
- Shared regional cluster with city-aware partitioning enabled from day one.

Phase 2 (multi-city expansion):
- Dedicated namespaces and node pools for heavy cities.
- Keep topic naming and Redis keyspace consistent so migrations remain config-level.

Phase 3 (mega-city isolation):
- Dedicated cluster for very high demand cities.
- Isolate noisy workloads and protect SLOs for all regions.

## 11. WebSocket Runtime Strategy
Recommended hybrid runtime:
- Go for core domain and orchestration services.
- BEAM stack (Elixir/Phoenix) for WebSocket connection management and pub/sub fanout.

Operational requirement:
- Use a shared Erlang cookie via Kubernetes Secret for reliable inter-node communication.

## 12. Reliability, Failure Handling, and DR
Highest risk:
- Redis outage (matching + tracking impact).

Mitigation:
- Highly available Redis deployment with failover.
- Regular failover drills.
- Graceful degradation: temporary fallback behavior for non-critical real-time features.

Additional safeguards:
- Multi-AZ node pools.
- Backup/restore runbooks for PostgreSQL, MongoDB, and ClickHouse.
- Explicit SLOs and error budgets per critical user journey.

## 13. Security and Compliance Basics
- mTLS/service auth inside cluster.
- Signed JWT with rotation.
- Secret management through Kubernetes Secret manager integration.
- PII encryption at rest and in transit.
- Audit logs for auth, payment, and admin operations.

## 14. Implementation Roadmap
Stage A: Platform foundation
- Define domain events, topic contracts, and schema version policy.
- Finalize service ownership and API boundaries.
- Create city-tier Helm values and scaling profiles.

Stage B: Core ride loop
- Build and harden Auth -> User -> Driver -> Booking/Pricing -> Matching -> Tracking -> Trip -> Payment.
- Validate booking-to-payment happy path and failure path handling.

Stage C: Production hardening
- Add PgBouncer, Redis HA, KEDA, surge pre-scaling, and rate limiting.
- Run load tests for rush-hour and event-exit scenarios.

Stage D: Growth layer
- Add Rating, Promo, Support, Analytics pipelines.
- Launch additional cities using the same platform blueprint.

## 15. Success Metrics for Production Readiness
- Booking confirmation latency at peak.
- Match success rate and match time.
- Real-time tracking freshness and update delivery ratio.
- Payment success + retry recovery success.
- Consumer lag thresholds by city/service.
- Uptime/SLO attainment during predictable and unpredictable spikes.

## 16. Next Planning Deliverables
- Service dependency and failure-impact matrix.
- Kafka topic catalog with partition and retention policy.
- Helm chart structure and values template per city tier.
- Capacity workbook for Tier-1, Tier-2, and Tier-3 Indian cities.
- Production launch checklist (observability, chaos drill, rollback, DR).

## 17. Step-by-Step Architecture Execution

Step 1: Enforce communication boundaries
- No tight coupling between services.
- No service reads another service database.
- gRPC is allowed only for blocking realtime decisions.
- Kafka is mandatory for fire-and-forget asynchronous workflows.

Step 2: Confirm gRPC-only paths
- booking-service -> pricing-service
- booking-service -> promo-service
- gateway -> auth/session checks only when request cannot proceed without response

Step 3: Make Kafka the async backbone
- Use events for lifecycle transitions, tracking streams, notifications, payment retries, rating, support, audit, and analytics.
- Apply topic versioning and backward-compatible schema policy.

Step 4: Implement 2-second driver location update contract
1. Driver app sends location every 2 seconds to driver-api-gateway.
2. Gateway forwards to tracking-service.
3. tracking-service writes latest point to Valkey with TTL.
4. tracking-service publishes Kafka event partitioned by city and trip.
5. user-api-gateway websocket pushes updates to rider UI.
6. analytics-ingestion-service consumes stream for replay, ETA quality, and heatmaps.

Step 5: Standardize tracking topic
- Topic: tracking.driver.location.v1
- Partition key: city_id:trip_id
- Delivery: at-least-once with idempotent consumers
- Retention: short hot-stream retention plus long-term analytics sink retention

Step 6: Apply realtime safeguards
- Drop duplicates with idempotency key: trip_id + driver_id + event_timestamp.
- Reject abnormal GPS jumps with speed/distance sanity checks.
- Use backpressure/coalescing per trip during spikes.
- Alert on per-city consumer lag and websocket fanout latency.

Step 7: Choose geo indexing strategy
- Primary: S2 Geometry cells for nearest-driver discovery.
- Secondary: GeoHash for prefix-based bucketing and compatibility.
- Avoid Quadtree as primary for high-churn moving-point workloads.

Step 8: Add search and multi-stop capabilities
- search-service handles pickup/drop search and candidate generation.
- stop-planning-service validates and orders intermediate stops.
- route-planning-service builds route graph for pickup/stops/destination.
- eta-service computes pickup ETA and trip ETA per candidate driver.

Step 9: Run dispatch cycle
- search-service requests eligible drivers from matching-service by S2 ring.
- matching-service expands ring cell-by-cell until SLA or max radius.
- dispatch-service handles offer, accept, timeout, and retry flow.

Step 10: Apply DB and cache allocation
- PostgreSQL: auth-service, session-service, user-service, driver-service, booking-service, payment-service, pricing-service, promo-service, rating-service.
- Valkey + Kafka + ClickHouse sink: tracking-service, matching-service, search-service, eta-service, dispatch-service.
- MongoDB (or PostgreSQL when strict relations are needed): trip-service, support-ticket-service, escalation-service, review-moderation-service.
- ClickHouse primary analytics store: analytics-ingestion-service, analytics-query-service, realtime-metrics-service.

Step 11: Set cache platform policy
- Default critical cache: Valkey.
- Optional acceleration tier: Dragonfly for non-critical read-heavy derived caches.
- Never keep source-of-truth business state only in cache.

Step 12: Handle optional specialized stores only when justified
- Cassandra: only for extreme append-heavy location history beyond PostgreSQL + ClickHouse cost/performance envelope.
- InfluxDB: optional for infra telemetry time-series, not transactional workflows.

## 18. Gateway Responsibility Steps

Step 1: Admin API Gateway
- Admin operations, elevated authorization, audit-heavy endpoints.

Step 2: User API Gateway
- Rider booking, trip, payment, promo, and trip-history endpoints.

Step 3: Driver API Gateway
- Driver availability, trip actions, earnings, and payout-facing endpoints.
