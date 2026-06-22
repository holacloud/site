# Kvnode

KVNode is a key-value node for low-latency JSON storage with an HTTP API and NDJSON replication stream. It is part of the HolaCloud ecosystem.

## Key Features

### WAL-Based Persistence
Every write is recorded in a Write-Ahead Log (WAL) before being applied to the in-memory store. The WAL feeds the replication stream used by child nodes.

### Pluggable WAL Backends
KVNode supports multiple WAL backends — memory, Kafka, PostgreSQL, Redis, and MongoDB — so you can choose the persistence layer that best fits your infrastructure.

### Multi-Language SDKs
Official SDKs are available for Go, Python, Java, JavaScript, Kotlin, PHP, and Node.js, making it easy to integrate KVNode into any stack.

### Replication Stream
Replication uses `/v1/replicate` and uppercase NDJSON commands such as `SET`, `DELETE`, `BASELINE_BEGIN`, `BASELINE_END`, and `PING`.

## Use Cases

- **Configuration Storage**: Store application configuration as JSON key-value pairs.
- **Feature Flags**: Manage feature toggles centrally and propagate changes in real time.
- **Service Discovery**: Register and discover service endpoints with low-latency reads.
- **Session Store**: Store user sessions with fast lookups.

## API Overview

KVNode exposes a RESTful API at `https://api.hola.cloud`. Internal endpoints require the presence of `X-Glue-Authentication` or both `apikey` and `secret` headers; missing headers return `403 forbidden`.

| Method | Path | Description | Auth |
|--------|------|-------------|------|
| GET | /healthz | Health check | Public |
| GET | /readyz | Readiness (checks parent connection) | Public |
| GET | /v1/status | Node status (collections, replication, uptime) | Internal |
| GET | /v1/metrics | Node metrics (writes, reads) | Internal |
| GET | /v1/collections | List collections | Internal |
| POST | /v1/collections | Create collection | Internal |
| DELETE | /v1/collections/{col} | Delete collection | Internal |
| GET | /v1/collections/{col}/keys | List keys (limit, prefix) | Internal |
| GET | /v1/collections/{col}/keys/* | Get key value | Internal |
| POST | /v1/collections/{col}/keys/* | Set key value | Internal |
| DELETE | /v1/collections/{col}/keys/* | Delete key | Internal |
| POST | /v1/replicate | Replication stream (NDJSON) | Internal |
