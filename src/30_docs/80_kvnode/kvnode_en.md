# KVNode

KVNode is a replicated, distributed key-value store designed for high availability and low-latency access. It is part of the HolaCloud ecosystem and provides a simple HTTP API for storing and retrieving JSON values.

## Key Features

### WAL-Based Persistence
Every write is recorded in a Write-Ahead Log (WAL) before being applied to the in-memory store. This ensures durability and enables seamless replication across nodes.

### Pluggable WAL Backends
KVNode supports multiple WAL backends — memory, Kafka, PostgreSQL, Redis, and MongoDB — so you can choose the persistence layer that best fits your infrastructure.

### Multi-Language SDKs
Official SDKs are available for Go, Python, Java, JavaScript, Kotlin, PHP, and Node.js, making it easy to integrate KVNode into any stack.

### Strong Consistency
Writes are replicated synchronously or asynchronously depending on the WAL backend, with linearizable semantics for single-key operations.

## Use Cases

- **Configuration Storage**: Store application configuration as key-value pairs with automatic replication across nodes.
- **Feature Flags**: Manage feature toggles centrally and propagate changes in real time.
- **Service Discovery**: Register and discover service endpoints with low-latency reads.
- **Session Store**: Store user sessions with fast lookups and built-in replication for fault tolerance.

## API Overview

KVNode exposes a RESTful API at `https://api.hola.cloud`. All endpoints that modify data require internal authentication via API key and secret, or the Glue gateway authentication header.

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
