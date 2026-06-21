# Replication & High Availability

KVNode achieves high availability through a **leader-follower replication** model. Changes written to a primary (leader) node are streamed in real time to replica (follower) nodes via an NDJSON-based replication protocol.

## How Replication Works

1. A client writes a key-value pair to any node in the cluster.
2. The receiving node persists the write in its WAL and applies it to the in-memory store.
3. If the node has a **parent** configured, it forwards the write upstream via the `/v1/replicate` endpoint.
4. If the node has **children**, it broadcasts the write to all connected replicas over the replication stream.
5. Each replica applies the change to its own store, maintaining an eventually consistent copy.

## Replication Protocol

The replication endpoint (`POST /v1/replicate`) uses **NDJSON** (Newline-Delimited JSON) as the wire format. The protocol supports these command types:

| Command | Description |
|---------|-------------|
| `set` | A key was created or updated |
| `delete` | A key was removed |
| `baseline_begin` | Start of a full collection snapshot |
| `baseline_end` | End of a full collection snapshot |
| `ping` | Keep-alive heartbeat |

### Example: Replication Stream

```
{"type":"baseline_begin","collection":"my-collection","seq":0}
{"type":"set","collection":"my-collection","key":"user:alice","value":{"role":"admin"},"version":1,"timestamp":1700000000}
{"type":"baseline_end","collection":"my-collection","seq":1}
{"type":"set","collection":"my-collection","key":"user:bob","value":{"role":"user"},"version":1,"timestamp":1700000001}
{"type":"ping","timestamp":1700000010}
```

## Node Status

The `/v1/status` endpoint exposes the replication topology:

```bash
curl "https://api.hola.cloud/v1/status" \
  -H "apikey: your-api-key" \
  -H "secret: your-api-secret"
```

Response:

```json
{
  "node": "node-1",
  "role": "primary",
  "parent": "",
  "collections": {
    "my-collection": {
      "keys": 42,
      "lastSeq": 128,
      "walSizeBytes": 65536
    }
  },
  "children": 2,
  "uptimeSeconds": 86400,
  "replication": {
    "enabled": true,
    "parent_connected": false
  }
}
```

## Node Metrics

Monitor performance with the `/v1/metrics` endpoint:

```bash
curl "https://api.hola.cloud/v1/metrics" \
  -H "apikey: your-api-key" \
  -H "secret: your-api-secret"
```

Response:

```json
{
  "writes_total": 15000,
  "reads_total": 82000,
  "replication_commands_sent_total": 15000,
  "replication_commands_received_total": 0,
  "children_connected": 2,
  "parent_connected": false
}
```

## Readiness Checks

The `/readyz` endpoint verifies that a node is ready to serve traffic. For replicas, this means the parent connection must be established and the WAL fully replayed. For primaries, readiness is immediate after startup.

```bash
curl "https://api.hola.cloud/readyz"
```

Response when not ready:

```json
{"ok":false,"node":"node-2","role":"replica","ready":false,"checks":{"wal_replayed":true,"parent_connected":false}}
```

## Pluggable WAL Backends

KVNode supports multiple WAL backends, selectable at startup:

| Backend | Description | Use Case |
|---------|-------------|----------|
| **Memory** | In-process buffer (non-durable) | Development, testing, ephemeral caches |
| **Kafka** | Durable log via Apache Kafka | High-throughput production deployments |
| **PostgreSQL** | WAL stored in PostgreSQL tables | Environments already using Postgres |
| **Redis** | WAL stored in Redis streams | Low-latency, Redis-centric architectures |
| **MongoDB** | WAL stored in MongoDB collections | MongoDB-centric deployments |

Each backend implements the same WAL interface, so switching between them requires no application-level changes.
