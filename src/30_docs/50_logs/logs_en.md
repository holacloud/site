# InstantLogs

InstantLogs is a real-time log aggregation and streaming service within the Hola.Cloud ecosystem. It provides multi-tenant log ingestion, regex-based filtering, live streaming, NDJSON export, and persistent storage — all with low latency.

## Key Benefits

### Real-Time Log Aggregation
Ingest logs from multiple sources in real time. InstantLogs collects, processes, and stores log entries as they arrive, giving you immediate visibility into your systems.

### Multi-Tenant Architecture
Organize logs into isolated loggers. Each logger has its own credentials, retention policy, and access controls, making InstantLogs suitable for teams, clients, or microservices operating in the same account.

### Regex Filtering
Filter log streams using regular expressions. Narrow down to entries matching a pattern, exclude noise, and focus on what matters without leaving the streaming endpoint.

### Live Streaming
Stream logs as they are ingested via long-lived HTTP connections. Choose between NDJSON and plain-text formats. Events appear in near real-time with minimal overhead.

### Persistent Storage
All ingested logs are durably stored and available for replay. You can re-stream historical data or export it for offline analysis.

### Dual Authentication Model
InstantLogs uses two authentication modes. Management operations (create, list, delete loggers, manage API keys) require an API key (`Api-Key` + `Api-Secret` headers). Data operations (ingest, stream, filter, stats) use a Logger Secret (`X-Instantlogs-Event-Secret` header or `Authorization: Bearer`).

## API Overview

### Management Endpoints (API Key Auth)

| Method | Path | Description |
|--------|------|-------------|
| GET | `/v1/loggers` | List all loggers |
| POST | `/v1/loggers` | Create a new logger |
| GET | `/v1/loggers/{id}` | Get logger details |
| DELETE | `/v1/loggers/{id}` | Delete a logger |
| POST | `/v1/loggers/{id}/apiKeys` | Create a new API key for a logger |
| DELETE | `/v1/loggers/{id}/apiKeys/{key}` | Delete an API key from a logger |

### Data Endpoints (Logger Secret Auth)

| Method | Path | Description |
|--------|------|-------------|
| POST | `/v1/loggers/{id}/ingest` | Ingest log entries |
| GET | `/v1/loggers/{id}/filter` | Stream and filter logs |
| GET | `/v1/loggers/{id}/events` | Stream events (NDJSON or text) |
| GET | `/v1/loggers/{id}/stats` | Get logger statistics |

### Framed Events (Event Secret Auth)

| Method | Path | Description |
|--------|------|-------------|
| POST | `/v1/ingest/events` | Ingest framed events |

Base URL: `https://api.hola.cloud`

## Best Use Cases

### Centralized Logging
Aggregate logs from all your services, containers, and servers into a single stream. Debug production issues faster with unified search and filtering.

### Real-Time Monitoring
Stream application and infrastructure logs live. Detect anomalies as they happen and trigger alerts based on log patterns.

### Audit Trails
Store immutable log records for compliance and auditing. Replay historical streams on demand.

### Multi-Tenant SaaS
Isolate logs per customer or team using separate loggers. Each tenant gets its own credentials and retention policies.
