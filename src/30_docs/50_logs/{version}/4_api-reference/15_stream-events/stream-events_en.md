
# GET /v1/loggers/{id}/events

Stream log events in real time. Supports NDJSON and plain-text formats.

## Authentication

Requires data credentials:

- `X-Instantlogs-Event-Secret: <secret>` or `Authorization: Bearer <secret>`

## Path Parameters

| Parameter | Description |
|-----------|-------------|
| `id` | The unique identifier of the logger |

## Query Parameters

| Parameter | Type | Default | Description |
|-----------|------|---------|-------------|
| `regex` | string | — | Regex pattern to filter events by message content |
| `follow` | bool | `true` | If `true`, keep the connection open and stream new events |
| `format` | string | `ndjson` | Output format: `ndjson` or `text` |

## Request

```bash
# Stream all events as NDJSON
curl "https://api.hola.cloud/v1/loggers/logger_xyz789/events" \
  -H "X-Instantlogs-Event-Secret: lgs_abc123def456"

# Stream as plain text
curl "https://api.hola.cloud/v1/loggers/logger_xyz789/events?format=text" \
  -H "X-Instantlogs-Event-Secret: lgs_abc123def456"

# Filter with regex
curl "https://api.hola.cloud/v1/loggers/logger_xyz789/events?regex=error&format=ndjson" \
  -H "X-Instantlogs-Event-Secret: lgs_abc123def456"
```

```http
GET /v1/loggers/logger_xyz789/events?format=ndjson HTTP/1.1
Host: api.hola.cloud
X-Instantlogs-Event-Secret: lgs_abc123def456
```

## Response

NDJSON format (default):

```
{"message":"User login successful","level":"info","service":"web","timestamp":"2025-06-20T14:22:30Z","id":"evt_001"}
{"message":"Database connection error","level":"error","service":"db","timestamp":"2025-06-20T14:22:31Z","id":"evt_002"}
```

Text format:

```
14:22:30 [info] [web] User login successful
14:22:31 [error] [db] Database connection error
```

## Error Codes

| Code | Description |
|------|-------------|
| 400 | Invalid query parameter |
| 401 | Missing or invalid event secret |
| 404 | Logger not found |
