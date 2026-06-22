
# Create Logger

Create a new logger.

## Authentication

Requires management credentials:

- `Api-Key` — Your API key
- `Api-Secret` — Your API secret

## Request Body

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `name` | string | yes | A human-readable name for the logger |

```json
{
  "name": "my-app-logger"
}
```

## Request

```bash
curl -X POST "https://api.hola.cloud/v1/loggers" \
  -H "Api-Key: your-api-key" \
  -H "Api-Secret: your-api-secret" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "my-app-logger"
  }'
```

```http
POST /v1/loggers HTTP/1.1
Host: api.hola.cloud
Api-Key: your-api-key
Api-Secret: your-api-secret
Content-Type: application/json

{
  "name": "my-app-logger"
}
```

## Response

```json
{
  "id": "logger_xyz789",
  "name": "my-app-logger",
  "secret": "lgs_abc123def456",
  "created_at": "2025-06-20T14:22:00Z"
}
```

## Error Codes

| Code | Description |
|------|-------------|
| 400 | Missing or invalid `name` field |
| 401 | Missing or invalid API credentials |
| 409 | A logger with this name already exists |
