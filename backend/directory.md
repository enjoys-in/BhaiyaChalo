# Go Workspace Directory Structure (30+ Microservices)

This file tracks only repository and Go workspace directory structure.

## Recommended Layout

```text
backend/
├─ go.work
├─ directory.md
├─ plan.md
├─ context.txt
│
├─ apps/
│  ├─ admin-api-gateway/
│  │  ├─ go.mod
│  │  ├─ cmd/server/main.go
│  │  ├─ internal/
│  │  │  ├─ config/
│  │  │  ├─ middleware/
│  │  │  ├─ route/
│  │  │  ├─ handler/
│  │  │  └─ client/
│  │  └─ deployments/helm/
│  │
│  ├─ user-api-gateway/
│  │  ├─ go.mod
│  │  ├─ cmd/server/main.go
│  │  ├─ internal/
│  │  │  ├─ config/
│  │  │  ├─ middleware/
│  │  │  ├─ route/
│  │  │  ├─ handler/
│  │  │  └─ client/
│  │  └─ deployments/helm/
│  │
│  └─ driver-api-gateway/
│     ├─ go.mod
│     ├─ cmd/server/main.go
│     ├─ internal/
│     │  ├─ config/
│     │  ├─ middleware/
│     │  ├─ route/
│     │  ├─ handler/
│     │  └─ client/
│     └─ deployments/helm/
│
├─ realtime/
│  └─ phoenix-websocket-gateway/
│     ├─ mix.exs
│     ├─ lib/
│     │  ├─ channels/
│     │  ├─ presence/
│     │  ├─ pubsub/
│     │  └─ telemetry/
│     ├─ config/
│     └─ deployments/helm/
│
├─ services/
│  ├─ identity/
│  │  ├─ auth-service/
│  │  ├─ session-service/
│  │  ├─ otp-service/
│  │  └─ iam-policy-service/
│  │
│  ├─ profile/
│  │  ├─ user-service/
│  │  ├─ driver-service/
│  │  ├─ vehicle-service/
│  │  └─ document-verification-service/
│  │
│  ├─ ride-core/
│  │  ├─ booking-service/
│  │  ├─ search-service/
│  │  ├─ matching-service/
│  │  ├─ dispatch-service/
│  │  ├─ tracking-service/
│  │  ├─ eta-service/
│  │  ├─ route-planning-service/
│  │  ├─ stop-planning-service/
│  │  ├─ trip-service/
│  │  ├─ fare-service/
│  │  ├─ pricing-service/
│  │  ├─ surge-policy-service/
│  │  └─ geofence-service/
│  │
│  ├─ realtime-core/
│  │  ├─ socket-session-service/
│  │  ├─ connection-manager-service/
│  │  ├─ subscription-service/
│  │  ├─ fanout-orchestrator-service/
│  │  ├─ location-ingest-service/
│  │  └─ delivery-ack-service/
│  │
│  ├─ driver-ops/
│  │  ├─ availability-service/
│  │  ├─ earnings-service/
│  │  ├─ incentive-service/
│  │  └─ payout-service/
│  │
│  ├─ payments/
│  │  ├─ payment-service/
│  │  ├─ wallet-service/
│  │  ├─ invoice-service/
│  │  └─ reconciliation-service/
│  │
│  ├─ growth/
│  │  ├─ promo-service/
│  │  ├─ campaign-service/
│  │  └─ referral-service/
│  │
│  ├─ engagement/
│  │  ├─ notification-service/
│  │  ├─ template-service/
│  │  ├─ rating-service/
│  │  └─ review-moderation-service/
│  │
│  ├─ trust-safety/
│  │  ├─ fraud-detection-service/
│  │  └─ risk-scoring-service/
│  │
│  ├─ support/
│  │  ├─ support-ticket-service/
│  │  └─ escalation-service/
│  │
│  └─ analytics/
│     ├─ analytics-ingestion-service/
│     ├─ analytics-query-service/
│     ├─ realtime-metrics-service/
│     └─ audit-log-service/
│
├─ libs/
│  ├─ common/
│  │  ├─ logger/
│  │  ├─ errors/
│  │  ├─ httpx/
│  │  ├─ telemetry/
│  │  ├─ idempotency/
│  │  ├─ types/
│  │  └─ constants/
│  ├─ grpc/
│  ├─ kafka/
│  ├─ redis/
│  ├─ postgres/
│  ├─ mongodb/
│  └─ clickhouse/
│
├─ contracts/
│  ├─ proto/
│  │  ├─ booking/
│  │  ├─ pricing/
│  │  ├─ promo/
│  │  └─ auth/
│  └─ events/
│     ├─ booking/
│     ├─ matching/
│     ├─ tracking/
│     ├─ trip/
│     ├─ payment/
│     ├─ notification/
│     ├─ rating/
│     └─ support/
│
├─ deployments/
│  ├─ helm/
│  │  ├─ base/
│  │  ├─ apps/
│  │  │  ├─ admin-api-gateway/
│  │  │  ├─ user-api-gateway/
│  │  │  └─ driver-api-gateway/
│  │  ├─ services/
│  │  ├─ environments/
│  │  │  ├─ dev/
│  │  │  ├─ staging/
│  │  │  └─ prod/
│  │  └─ city-tiers/
│  │     ├─ tier-1/
│  │     ├─ tier-2/
│  │     └─ tier-3/
│  └─ k8s/
│
├─ scripts/
│  ├─ local/
│  ├─ ci/
│  └─ release/
│
└─ docs/
   ├─ architecture/
   ├─ adr/
   ├─ runbooks/
   └─ api/
```

## Go Workspace Scope

Include at minimum:
- apps/admin-api-gateway
- apps/user-api-gateway
- apps/driver-api-gateway
- all services folders under services/
- shared libraries under libs/
- contracts/proto generated client modules

Phoenix scope:
- realtime/phoenix-websocket-gateway is an Elixir app and is not part of go.work.

## Directory Notes

- Keep naming consistent as `<domain>-service`.
- Keep each service folder independently buildable and deployable.
- Keep shared code only under libs/ to avoid cross-service folder imports.
- Keep websocket handling in realtime/phoenix-websocket-gateway and keep ride business logic in Go services.
- See plan.md for architecture, communication, database, and rollout instructions.

## Types and Interfaces Convention

Yes, you should keep explicit folders for these.

Recommended per service or gateway internal layout:
- `internal/types/` for domain types, DTOs, and shared structs inside that module.
- `internal/constants/` for fixed values, event names, and config keys.
- `internal/ports/` for interfaces (hexagonal style input/output ports).

Recommended shared layout:
- `libs/common/types/` for truly cross-service types only.
- `libs/common/constants/` for global constants used by multiple modules.

Rule:
- Prefer local (`internal/*`) first.
- Move to `libs/common/*` only when at least two modules genuinely share the same contract.
