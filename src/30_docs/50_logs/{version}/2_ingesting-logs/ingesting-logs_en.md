# Ingesting Logs

InstantLogs supports two methods for ingesting log data: direct ingestion into a logger and framed events ingestion. Both methods require the appropriate secret for authentication.

## Authentication

Data ingestion endpoints use the Logger Secret for authentication. You can provide it in one of two ways:

- **Header**: `X-Instantlogs-Event-Secret: <your-logger-secret>`
- **Bearer Token**: `Authorization: Bearer <your-logger-secret>`

## Method 1: Direct Ingestion

Send log entries directly to a specific logger using the `POST /v1/loggers/{id}/ingest` endpoint.

### Single Entry

```bash
curl -X POST "https://api.hola.cloud/v1/loggers/logger_xyz789/ingest" \
  -H "X-Instantlogs-Event-Secret: lgs_abc123def456" \
  -H "Content-Type: application/json" \
  -d '{
    "message": "User login successful",
    "level": "info",
    "service": "auth-service",
    "timestamp": "2025-06-20T15:00:00Z",
    "metadata": {
      "user_id": "usr_123",
      "ip": "192.168.1.1"
    }
  }'
```

Expected response:

```json
{
  "ingested": 1
}
```

### Batch Entries

You can send multiple log entries in a single request by providing an array:

```bash
curl -X POST "https://api.hola.cloud/v1/loggers/logger_xyz789/ingest" \
  -H "X-Instantlogs-Event-Secret: lgs_abc123def456" \
  -H "Content-Type: application/json" \
  -d '[
    {
      "message": "Request started",
      "level": "info",
      "service": "api-gateway",
      "timestamp": "2025-06-20T15:00:00Z"
    },
    {
      "message": "Request completed",
      "level": "info",
      "service": "api-gateway",
      "timestamp": "2025-06-20T15:00:05Z"
    }
  ]'
```

Expected response:

```json
{
  "ingested": 2
}
```

## Method 2: Framed Events Ingestion

For high-throughput scenarios or when using binary protocols, use the framed events endpoint `POST /v1/ingest/events`. This endpoint uses the Event Secret for authentication and accepts a stream of framed events.

```bash
curl -X POST "https://api.hola.cloud/v1/ingest/events" \
  -H "X-Instantlogs-Event-Secret: evt_secret_789" \
  -H "Content-Type: application/json" \
  -d '{
    "logger_id": "logger_xyz789",
    "events": [
      {
        "message": "Framed event 1",
        "level": "warn",
        "timestamp": "2025-06-20T15:01:00Z"
      },
      {
        "message": "Framed event 2",
        "level": "error",
        "timestamp": "2025-06-20T15:01:01Z"
      }
    ]
  }'
```

## HTTP Request Reference

```http
POST /v1/loggers/{id}/ingest HTTP/1.1
Host: api.hola.cloud
X-Instantlogs-Event-Secret: lgs_abc123def456
Content-Type: application/json

{
  "message": "example log",
  "level": "info",
  "service": "web",
  "timestamp": "2025-06-20T15:00:00Z"
}

POST /v1/ingest/events HTTP/1.1
Host: api.hola.cloud
X-Instantlogs-Event-Secret: evt_secret_789
Content-Type: application/json

{
  "logger_id": "logger_xyz789",
  "events": [ ... ]
}
```
