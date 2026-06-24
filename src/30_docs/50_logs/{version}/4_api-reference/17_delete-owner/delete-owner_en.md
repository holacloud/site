
# Delete Logger Owner

Remove an owner from a logger. Owners cannot remove themselves.

## Authentication

Requires management credentials:

- `Api-Key` — Your API key
- `Api-Secret` — Your API secret

## Path Parameters

| Parameter | Description |
|-----------|-------------|
| `id` | The unique identifier of the logger |
| `ownerId` | The user ID of the owner to remove |

## Request

```bash
curl -X DELETE "https://api.hola.cloud/v1/loggers/logger_xyz789/owners/user_abc123" \
  -H "Api-Key: LOGGER_API_KEY" \
  -H "Api-Secret: LOGGER_API_SECRET"
```

```http
DELETE /v1/loggers/logger_xyz789/owners/user_abc123 HTTP/1.1
Host: api.hola.cloud
Api-Key: LOGGER_API_KEY
Api-Secret: LOGGER_API_SECRET
```

## Response

Returns the updated list of owners.

```json
["user_def456"]
```

## Error Codes

| Code | Description |
|------|-------------|
| 400 | Owner cannot be removed by themselves |
| 401 | Missing or invalid API credentials |
| 403 | API credentials do not have access to this logger |
| 404 | Logger not found |
