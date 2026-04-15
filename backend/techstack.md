# Ride Booking Tech Stack

## 1. Stack Principles
- OSS-first and self-hosted friendly.
- Loose coupling across services.
- Event-driven architecture with Kafka as the async backbone.
- gRPC only for realtime blocking decisions.
- Database-per-service ownership.

## 2. Core Platform Stack
- Language: Go (all core services and gateways)
- Container orchestration: Kubernetes
- Packaging and deployment: Helm
- API ingress gateway: Traefik OSS or Kong OSS
- Internal communication:
  - Sync: gRPC (only blocking paths)
  - Async: Kafka (fire-and-forget events)

## 3. Data and Infra Stack
- Transactional DB: PostgreSQL
- Flexible document DB: MongoDB
- Cache and live state: Valkey
- Optional cache acceleration: Dragonfly (non-critical read-heavy caches)
- Analytics DB: ClickHouse
- Optional extreme write history: Cassandra
- Optional infra time-series: InfluxDB
- Routing and maps: OpenStreetMap plus OSRM or Valhalla
- Geo indexing: S2 Geometry (primary), GeoHash (secondary)

## 4. Service-Level Tech Stack Mapping

| Service | Primary Store | Cache/Live State | Event/Stream | Sync Calls | Notes |
|---|---|---|---|---|---|
| admin-api-gateway | None | Valkey (rate-limit/session cache) | Kafka producer/consumer as needed | gRPC to auth/session | Admin-facing API edge |
| user-api-gateway | None | Valkey (rate-limit/session cache) | Kafka producer/consumer as needed | gRPC to auth/session | Rider-facing API edge |
| driver-api-gateway | None | Valkey (rate-limit/session cache) | Kafka producer/consumer as needed | gRPC to auth/session | Driver-facing API edge |
| auth-service | PostgreSQL | Valkey | Kafka | gRPC server for auth validation | Identity and token lifecycle |
| session-service | PostgreSQL | Valkey | Kafka | gRPC server for session checks | Session and refresh state |
| otp-service | PostgreSQL | Valkey | Kafka | Optional gRPC | OTP issue and verify |
| iam-policy-service | PostgreSQL | Valkey | Kafka | Optional gRPC | Policy and role rules |
| user-service | PostgreSQL | Valkey | Kafka | Optional gRPC | Rider profile |
| driver-service | PostgreSQL | Valkey | Kafka | Optional gRPC | Driver profile |
| vehicle-service | PostgreSQL | Valkey | Kafka | Optional gRPC | Vehicle metadata |
| document-verification-service | PostgreSQL | Valkey | Kafka | Optional gRPC | KYC and document checks |
| booking-service | PostgreSQL | Valkey | Kafka | gRPC to pricing and promo | Booking orchestration |
| search-service | PostgreSQL (query metadata) | Valkey | Kafka | gRPC to eta and pricing | Cab search results |
| matching-service | PostgreSQL (rules/config) | Valkey (S2/GeoHash live index) | Kafka | Optional gRPC | Nearest driver and assignment |
| dispatch-service | PostgreSQL | Valkey | Kafka | Optional gRPC | Offer, timeout, retry flow |
| tracking-service | PostgreSQL (minimal metadata) | Valkey (latest location with TTL) | Kafka | Optional gRPC | 2-second driver location ingestion |
| eta-service | PostgreSQL (model config) | Valkey | Kafka | gRPC to search/matching | Pickup and trip ETA |
| route-planning-service | MongoDB | Valkey | Kafka | Optional gRPC | Route graph and path plans |
| stop-planning-service | MongoDB | Valkey | Kafka | Optional gRPC | Multi-stop validation and ordering |
| trip-service | MongoDB | Valkey | Kafka | Optional gRPC | Trip state machine |
| fare-service | PostgreSQL | Valkey | Kafka | gRPC to booking/search | Fare composition |
| pricing-service | PostgreSQL | Valkey | Kafka | gRPC server for booking/search | Pricing and surge inputs |
| surge-policy-service | PostgreSQL | Valkey | Kafka | Optional gRPC | Surge rule evaluation |
| geofence-service | PostgreSQL | Valkey | Kafka | Optional gRPC | Zone and boundary enforcement |
| availability-service | PostgreSQL | Valkey | Kafka | Optional gRPC | Driver availability state |
| earnings-service | PostgreSQL | Valkey | Kafka | Optional gRPC | Driver earnings snapshots |
| incentive-service | PostgreSQL | Valkey | Kafka | Optional gRPC | Incentive programs |
| payout-service | PostgreSQL | Valkey | Kafka | Optional gRPC | Driver payouts |
| payment-service | PostgreSQL | Valkey | Kafka | gRPC for critical payment checks | Payment lifecycle |
| wallet-service | PostgreSQL | Valkey | Kafka | Optional gRPC | Wallet balance and ledger |
| invoice-service | PostgreSQL | Valkey | Kafka | Optional gRPC | Invoice generation |
| reconciliation-service | PostgreSQL | Valkey | Kafka | Optional gRPC | Settlement and reconciliation |
| promo-service | PostgreSQL | Valkey | Kafka | gRPC server for booking | Promo and discount rules |
| campaign-service | PostgreSQL | Valkey | Kafka | Optional gRPC | Campaign targeting |
| referral-service | PostgreSQL | Valkey | Kafka | Optional gRPC | Referral logic |
| notification-service | PostgreSQL (template refs/log refs) | Valkey | Kafka consumer | None required | Push, SMS, email delivery |
| template-service | PostgreSQL | Valkey | Kafka | Optional gRPC | Notification templates |
| rating-service | PostgreSQL | Valkey | Kafka | Optional gRPC | Post-trip ratings |
| review-moderation-service | MongoDB | Valkey | Kafka | Optional gRPC | Review and moderation workflows |
| fraud-detection-service | PostgreSQL | Valkey | Kafka | Optional gRPC | Fraud signals and rules |
| risk-scoring-service | PostgreSQL | Valkey | Kafka | Optional gRPC | Dynamic risk scoring |
| support-ticket-service | MongoDB | Valkey | Kafka | Optional gRPC | Ticket lifecycle |
| escalation-service | MongoDB | Valkey | Kafka | Optional gRPC | Escalation workflows |
| analytics-ingestion-service | ClickHouse | Valkey (optional query cache) | Kafka consumer | None required | Stream ingest and transformation |
| analytics-query-service | ClickHouse | Valkey (optional query cache) | Kafka optional | Optional gRPC | Analytics API layer |
| realtime-metrics-service | ClickHouse | Valkey (optional) | Kafka consumer | Optional gRPC | Realtime KPIs |
| audit-log-service | ClickHouse | None or Valkey | Kafka consumer | None required | Immutable event audit queries |

## 5. Analytics Use Cases and Store Choice
- User distance travelled: ClickHouse
- Driver movement history: ClickHouse
- Driver earnings and daily payouts: ClickHouse (source from payment and earnings events)
- City-level ride and demand metrics: ClickHouse
- Operational dashboards: Grafana or Apache Superset on top of ClickHouse

## 6. Event Backbone Standards
- Primary broker: Kafka
- Topic key pattern: city_id plus entity_id
- Event delivery model: at-least-once
- Consumer requirement: idempotent processing
- Dead-letter and retry topics: required for all critical flows

## 7. Realtime Location Standards
- Driver location publish interval: every 2 seconds
- Live state: Valkey with TTL
- Durable stream: Kafka topic tracking.driver.location.v1
- UI push: WebSocket from user-api-gateway or dedicated websocket gateway
- Fallback: snapshot polling endpoint on socket reconnect failure

## 8. Cost-Control and Scale Guidance
- Start with PostgreSQL, Valkey, Kafka, ClickHouse.
- Add Cassandra only for very high append-only location history pressure.
- Add InfluxDB only for infra telemetry if Prometheus retention is not enough.
- Keep Dragonfly optional and benchmark-driven.
