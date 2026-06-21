
# POST /v1/loggers/{id}/ingest

Ingest log entries into a logger.

## Authentication

Requires data credentials:

- `X-Instantlogs-Event-Secret: <secret>` or `Authorization: Bearer <secret>`

## Path Parameters

| Parameter | Description |
|-----------|-------------|
| `id` | The unique identifier of the logger |

## Request Body

Single log entry:

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `message` | string | yes | The log message content |
| `level` | string | no | Log level (e.g., `info`, `warn`, `error`) |
| `service` | string | no | Source service name |
| `timestamp` | string | no | ISO 8601 timestamp (defaults to server time) |

```json
{
  "message": "User login successful",
  "level": "info",
  "service": "web",
  "timestamp": "2025-06-20T14:22:30Z"
}
```

## Request

```bash
curl -X POST "https://api.hola.cloud/v1/loggers/logger_xyz789/ingest" \
  -H "X-Instantlogs-Event-Secret: lgs_abc123def456" \
  -H "Content-Type: application/json" \
  -d '{
    "message": "User login successful",
    "level": "info",
    "service": "web",
    "timestamp": "2025-06-20T14:22:30Z"
  }'
```

```http
POST /v1/loggers/logger_xyz789/ingest HTTP/1.1
Host: api.hola.cloud
X-Instantlogs-Event-Secret: lgs_abc123def456
Content-Type: application/json

{
  "message": "User login successful",
  "level": "info",
  "service": "web",
  "timestamp": "2025-06-20T14:22:30Z"
}
```

## Response

```json
{
  "ingested": 1
}
```

## Error Codes

| Code | Description |
|------|-------------|
| 400 | Missing or invalid request body |
| 401 | Missing or invalid event secret |
| 404 | Logger not found |
| 413 | Payload too large |
