# Go Workspace Directory Structure (50+ Microservices)

This file tracks repository directory structure, MVC file layout, and conventions.

## Service Internal Layout (MVC Pattern)

Every Go microservice follows this identical internal structure. No file should exceed 300 lines except handler.go.

```text
<service-name>/
├─ go.mod
├─ Dockerfile
├─ .env
├─ cmd/server/main.go              ← entry point, wires dependencies
├─ internal/
│  ├─ config/config.go             ← env loading, Config struct
│  ├─ constants/constants.go       ← service name, topic names, fixed values
│  ├─ model/entity.go              ← domain entities (DB models)
│  ├─ dto/
│  │  ├─ request.go                ← incoming request structs
│  │  └─ response.go               ← outgoing response structs
│  ├─ ports/
│  │  ├─ repository.go             ← repository interface
│  │  ├─ service.go                ← service interface
│  │  └─ event.go                  ← event publisher interface
│  ├─ repository/
│  │  └─ postgres.go               ← (or mongodb.go / clickhouse.go)
│  ├─ service/service.go           ← business logic implementation
│  ├─ event/producer.go            ← Kafka event publisher
│  ├─ handler/
│  │  ├─ handler.go                ← HTTP handlers (can exceed 300 lines)
│  │  ├─ helpers.go                ← response helpers (successJSON/errorJSON)
│  │  └─ routes.go                 ← route registration
│  └─ client/                      ← gRPC clients to other services
└─ deployments/helm/               ← (optional per-service overrides)
```

## Gateway Internal Layout

```text
<gateway-name>/
├─ go.mod
├─ Dockerfile
├─ .env
├─ cmd/server/main.go
├─ internal/
│  ├─ config/config.go
│  ├─ constants/constants.go
│  ├─ middleware/
│  │  ├─ auth.go                   ← JWT token validation
│  │  ├─ ratelimit.go              ← Valkey-backed rate limiting
│  │  ├─ cors.go                   ← CORS headers
│  │  └─ logging.go                ← request/response logging
│  ├─ handler/health.go            ← health and readiness endpoints
│  ├─ route/router.go              ← route registration + middleware chain
│  └─ client/grpc.go               ← gRPC client stubs
└─ deployments/helm/
```

## Standard API Response Format

All APIs (user, admin, driver gateways) use a single response envelope:

```json
{
  "success": true,
  "message": "booking created successfully",
  "result": {}
}
```

- `success`: boolean — true on 2xx, false on error
- `message`: human-readable status message
- `result`: response payload — object, array, string, or null

Implemented in `libs/common/httpx/response.go` and per-service `handler/helpers.go`.

## Recommended Layout

```text
backend/
├─ go.work
├─ .github/
│  └─ workflows/
│     ├─ ci.yml                    ← build/test only changed services
│     └─ deploy.yml                ← build image + helm deploy per service
│
├─ apps/
│  ├─ admin-api-gateway/           ← admin-facing API edge (port 3001)
│  ├─ user-api-gateway/            ← rider-facing API edge (port 3002)
│  └─ driver-api-gateway/          ← driver-facing API edge (port 3003)
│
├─ realtime/
│  └─ phoenix-websocket-gateway/   ← Elixir/Phoenix (not in go.work)
│
├─ services/
│  ├─ identity/
│  │  ├─ auth-service/             ← port 8081
│  │  ├─ session-service/          ← port 8082
│  │  ├─ otp-service/              ← port 8083
│  │  └─ iam-policy-service/       ← port 8084
│  │
│  ├─ profile/
│  │  ├─ user-service/             ← port 8101
│  │  ├─ driver-service/           ← port 8102
│  │  ├─ vehicle-service/          ← port 8103
│  │  └─ document-verification-service/ ← port 8104
│  │
│  ├─ ride-core/
│  │  ├─ booking-service/          ← port 8201
│  │  ├─ search-service/           ← port 8202
│  │  ├─ matching-service/         ← port 8203
│  │  ├─ dispatch-service/         ← port 8204
│  │  ├─ tracking-service/         ← port 8205
│  │  ├─ eta-service/              ← port 8206
│  │  ├─ route-planning-service/   ← port 8207 (MongoDB)
│  │  ├─ stop-planning-service/    ← port 8208 (MongoDB)
│  │  ├─ trip-service/             ← port 8209 (MongoDB)
│  │  ├─ fare-service/             ← port 8210
│  │  ├─ pricing-service/          ← port 8211
│  │  ├─ surge-policy-service/     ← port 8212
│  │  └─ geofence-service/         ← port 8213
│  │
│  ├─ realtime-core/
│  │  ├─ socket-session-service/   ← port 8301
│  │  ├─ connection-manager-service/ ← port 8302
│  │  ├─ subscription-service/     ← port 8303
│  │  ├─ fanout-orchestrator-service/ ← port 8304
│  │  ├─ location-ingest-service/  ← port 8305
│  │  └─ delivery-ack-service/     ← port 8306
│  │
│  ├─ driver-ops/
│  │  ├─ availability-service/     ← port 8401
│  │  ├─ earnings-service/         ← port 8402
│  │  ├─ incentive-service/        ← port 8403
│  │  └─ payout-service/           ← port 8404
│  │
│  ├─ payments/
│  │  ├─ payment-service/          ← port 8501
│  │  ├─ wallet-service/           ← port 8502
│  │  ├─ invoice-service/          ← port 8503
│  │  └─ reconciliation-service/   ← port 8504
│  │
│  ├─ growth/
│  │  ├─ promo-service/            ← port 8601
│  │  ├─ campaign-service/         ← port 8602
│  │  └─ referral-service/         ← port 8603
│  │
│  ├─ engagement/
│  │  ├─ notification-service/     ← port 8701
│  │  ├─ template-service/         ← port 8702
│  │  ├─ rating-service/           ← port 8703
│  │  └─ review-moderation-service/ ← port 8704 (MongoDB)
│  │
│  ├─ trust-safety/
│  │  ├─ fraud-detection-service/  ← port 8801
│  │  └─ risk-scoring-service/     ← port 8802
│  │
│  ├─ support/
│  │  ├─ support-ticket-service/   ← port 8901 (MongoDB)
│  │  └─ escalation-service/       ← port 8902 (MongoDB)
│  │
│  └─ analytics/
│     ├─ analytics-ingestion-service/ ← port 9001 (ClickHouse)
│     ├─ analytics-query-service/   ← port 9002 (ClickHouse)
│     ├─ realtime-metrics-service/  ← port 9003 (ClickHouse)
│     └─ audit-log-service/         ← port 9004 (ClickHouse)
│
├─ libs/
│  ├─ common/
│  │  ├─ go.mod
│  │  ├─ logger/logger.go          ← structured JSON logger
│  │  ├─ errors/errors.go          ← AppError with HTTP status codes
│  │  ├─ httpx/
│  │  │  ├─ response.go            ← {success, message, result} envelope
│  │  │  └─ server.go              ← graceful shutdown HTTP server
│  │  ├─ telemetry/telemetry.go    ← OpenTelemetry setup
│  │  ├─ idempotency/idempotency.go ← idempotency key generation
│  │  ├─ types/types.go            ← Pagination, Coordinate, CityID
│  │  └─ constants/constants.go    ← global Kafka topics, header keys
│  ├─ grpc/grpc.go                 ← gRPC server/client helpers
│  ├─ kafka/kafka.go               ← Kafka producer/consumer wrappers
│  ├─ redis/redis.go               ← Valkey/Redis client wrapper
│  ├─ postgres/postgres.go         ← PostgreSQL connection pool
│  ├─ mongodb/mongodb.go           ← MongoDB client wrapper
│  └─ clickhouse/clickhouse.go     ← ClickHouse client wrapper
│
├─ contracts/
│  ├─ proto/                       ← gRPC protobuf definitions
│  └─ events/                      ← Kafka event schema definitions
│
├─ deployments/
│  ├─ helm/
│  │  ├─ base/                     ← shared Helm library chart
│  │  │  ├─ Chart.yaml
│  │  │  ├─ values.yaml
│  │  │  └─ templates/
│  │  │     ├─ _helpers.tpl
│  │  │     ├─ deployment.yaml
│  │  │     ├─ service.yaml
│  │  │     ├─ configmap.yaml
│  │  │     └─ hpa.yaml
│  │  ├─ apps/                     ← per-gateway Chart.yaml + values.yaml
│  │  │  ├─ admin-api-gateway/
│  │  │  ├─ user-api-gateway/
│  │  │  └─ driver-api-gateway/
│  │  ├─ services/                 ← per-service Chart.yaml + values.yaml
│  │  │  ├─ auth-service/
│  │  │  ├─ booking-service/
│  │  │  ├─ ... (one folder per service)
│  │  │  └─ audit-log-service/
│  │  ├─ environments/             ← env-specific value overrides
│  │  │  ├─ dev/values.yaml
│  │  │  ├─ staging/values.yaml
│  │  │  └─ prod/values.yaml
│  │  └─ city-tiers/
│  │     ├─ tier-1/
│  │     ├─ tier-2/
│  │     └─ tier-3/
│  ├─ k8s/
│  └─ local/
│     ├─ docker-compose.yml
│     └─ postgres/init/
│
├─ .github/workflows/
│  ├─ ci.yml                       ← per-service build/test on PR
│  └─ deploy.yml                   ← per-service image + helm deploy on main
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
- No file should exceed 300 lines except handler.go (handlers can be longer due to multiple endpoints).
- See plan.md for architecture, communication, database, and rollout instructions.

## MVC / Hexagonal Convention

| Layer | Folder | Responsibility |
|---|---|---|
| Model | `internal/model/` | DB entities, domain types |
| DTO | `internal/dto/` | Request/response structs (API layer) |
| Port | `internal/ports/` | Interfaces for repository, service, event |
| Repository | `internal/repository/` | Data access (Postgres/Mongo/ClickHouse) |
| Service | `internal/service/` | Business logic, orchestration |
| Event | `internal/event/` | Kafka publisher implementation |
| Handler | `internal/handler/` | HTTP handlers + routes + response helpers |
| Config | `internal/config/` | Environment config loader |
| Client | `internal/client/` | gRPC clients to other services |
| Constants | `internal/constants/` | Service-specific constants |

## Types and Interfaces Convention

Recommended per service or gateway internal layout:
- `internal/constants/` for fixed values, event names, and config keys.
- `internal/ports/` for interfaces (hexagonal style input/output ports).
- `internal/model/` for domain entities.
- `internal/dto/` for API request/response structs.

Recommended shared layout:
- `libs/common/types/` for truly cross-service types only.
- `libs/common/constants/` for global constants used by multiple modules.

Rule:
- Prefer local (`internal/*`) first.
- Move to `libs/common/*` only when at least two modules genuinely share the same contract.

## CI/CD Pipeline Convention

The GitHub Actions pipeline detects changes per-service using `git diff`:
- If code changes in `services/identity/auth-service/` → only auth-service is built, tested, and deployed.
- If code changes in `libs/` → all libs are tested (downstream services rebuild on next push).
- Each service has its own Helm chart under `deployments/helm/services/<name>/`.
- Deploy command: `helm upgrade --install <name> deployments/helm/services/<name>/ -f deployments/helm/environments/<env>/values.yaml`

## Helm Chart Convention

Every service Helm chart depends on `deployments/helm/base/` (library chart):
- base/ provides shared templates (Deployment, Service, ConfigMap, HPA).
- Each service chart has only `Chart.yaml` + `values.yaml` (service name, port, resources, env).
- Environment overrides in `deployments/helm/environments/{dev,staging,prod}/values.yaml`.
- To deploy a single service: `helm upgrade --install <svc-name> deployments/helm/services/<svc-name>/ -f deployments/helm/environments/dev/values.yaml`
