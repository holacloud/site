
# Delete Logger API Key

Delete an API key from a logger.

## Authentication

Requires management credentials:

- `Api-Key` — Your API key
- `Api-Secret` — Your API secret

## Path Parameters

| Parameter | Description |
|-----------|-------------|
| `id` | The unique identifier of the logger |
| `key` | The ID of the API key to delete |

## Request

```bash
curl -X DELETE "https://api.hola.cloud/v1/loggers/logger_xyz789/apiKeys/ak_123456" \
  -H "Api-Key: LOGGER_API_KEY" \
  -H "Api-Secret: LOGGER_API_SECRET"
```

```http
DELETE /v1/loggers/logger_xyz789/apiKeys/ak_123456 HTTP/1.1
Host: api.hola.cloud
Api-Key: LOGGER_API_KEY
Api-Secret: LOGGER_API_SECRET
```

## Response

HTTP `204 No Content`.

## Error Codes

| Code | Description |
|------|-------------|
| 401 | Missing or invalid API credentials |
| 403 | API credentials do not have access to this logger |
| 404 | Logger or API key not found |
