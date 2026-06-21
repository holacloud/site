
# POST /v1/loggers/{id}/apiKeys

Create a new API key for a logger. API keys allow management access to the logger.

## Authentication

Requires management credentials:

- `Api-Key` — Your API key
- `Api-Secret` — Your API secret

## Path Parameters

| Parameter | Description |
|-----------|-------------|
| `id` | The unique identifier of the logger |

## Request Body

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `name` | string | yes | A label to identify this API key |

```json
{
  "name": "ci-cd-key"
}
```

## Request

```bash
curl -X POST "https://api.hola.cloud/v1/loggers/logger_xyz789/apiKeys" \
  -H "Api-Key: your-api-key" \
  -H "Api-Secret: your-api-secret" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "ci-cd-key"
  }'
```

```http
POST /v1/loggers/logger_xyz789/apiKeys HTTP/1.1
Host: api.hola.cloud
Api-Key: your-api-key
Api-Secret: your-api-secret
Content-Type: application/json

{
  "name": "ci-cd-key"
}
```

## Response

```json
{
  "id": "ak_123456",
  "name": "ci-cd-key",
  "key": "lgs_api_key_value",
  "created_at": "2025-06-20T15:00:00Z"
}
```

## Error Codes

| Code | Description |
|------|-------------|
| 400 | Missing or invalid `name` field |
| 401 | Missing or invalid API credentials |
| 403 | API credentials do not have access to this logger |
| 404 | Logger not found |
