
# DELETE /v1/loggers/{id}

Delete a logger and all its associated log entries permanently.

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
curl -X DELETE "https://api.hola.cloud/v1/loggers/logger_xyz789" \
  -H "Api-Key: your-api-key" \
  -H "Api-Secret: your-api-secret"
```

```http
DELETE /v1/loggers/logger_xyz789 HTTP/1.1
Host: api.hola.cloud
Api-Key: your-api-key
Api-Secret: your-api-secret
```

## Response

HTTP `204 No Content`.

## Error Codes

| Code | Description |
|------|-------------|
| 401 | Missing or invalid API credentials |
| 403 | API credentials do not have access to this logger |
| 404 | Logger not found |
