
# GET /v1/loggers/{id}/filter

Stream and filter log entries in real time using a regex pattern.

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
| `pattern` | string | — | Regex pattern to filter log entries by message content |
| `follow` | bool | `false` | If `true`, keep the connection open and stream new entries |

## Request

```bash
curl "https://api.hola.cloud/v1/loggers/logger_xyz789/filter?pattern=error&follow=true" \
  -H "X-Instantlogs-Event-Secret: lgs_abc123def456"
```

```http
GET /v1/loggers/logger_xyz789/filter?pattern=error&follow=true HTTP/1.1
Host: api.hola.cloud
X-Instantlogs-Event-Secret: lgs_abc123def456
```

## Response

Returns matching log entries as NDJSON:

```
{"message":"Database connection error","level":"error","service":"db","timestamp":"2025-06-20T14:22:31Z","id":"evt_002"}
{"message":"Timeout exceeded","level":"error","service":"web","timestamp":"2025-06-20T14:22:32Z","id":"evt_003"}
```

With `follow=true`, the connection remains open and new matching entries are streamed as they arrive.

## Error Codes

| Code | Description |
|------|-------------|
| 400 | Invalid regex pattern |
| 401 | Missing or invalid event secret |
| 404 | Logger not found |
