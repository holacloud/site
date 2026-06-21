
# GET /v1/loggers/{id}

Get details of a specific logger by its ID.

## Authentication

Requires management credentials:

- `Api-Key` — Your API key
- `Api-Secret` — Your API secret

## Path Parameters

| Parameter | Description |
|-----------|-------------|
| `id` | The unique identifier of the logger |

## Request

```bash
curl "https://api.hola.cloud/v1/loggers/logger_xyz789" \
  -H "Api-Key: your-api-key" \
  -H "Api-Secret: your-api-secret"
```

```http
GET /v1/loggers/logger_xyz789 HTTP/1.1
Host: api.hola.cloud
Api-Key: your-api-key
Api-Secret: your-api-secret
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
| 401 | Missing or invalid API credentials |
| 403 | API credentials do not have access to this logger |
| 404 | Logger not found |
