
# Logger Stats

Get statistics for a logger, including total events ingested and storage usage.

## Authentication

Requires data credentials:

- `X-Instantlogs-Event-Secret: <secret>` or `Authorization: Bearer <secret>`

## Path Parameters

| Parameter | Description |
|-----------|-------------|
| `id` | The unique identifier of the logger |

## Request

```bash
curl "https://api.hola.cloud/v1/loggers/logger_xyz789/stats" \
  -H "X-Instantlogs-Event-Secret: lgs_abc123def456"
```

```http
GET /v1/loggers/logger_xyz789/stats HTTP/1.1
Host: api.hola.cloud
X-Instantlogs-Event-Secret: lgs_abc123def456
```

## Response

```json
{
  "total_events": 15234,
  "storage_bytes": 4194304,
  "events_last_hour": 234,
  "events_last_day": 5601
}
```

## Error Codes

| Code | Description |
|------|-------------|
| 401 | Missing or invalid event secret |
| 404 | Logger not found |
