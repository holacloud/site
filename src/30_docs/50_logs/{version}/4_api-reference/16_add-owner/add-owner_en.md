
# Add Logger Owner

Add an owner to a logger. Owners have management access to the logger.

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
| `user_id` | string | yes | The user ID to add as owner |

```json
{
  "user_id": "user_abc123"
}
```

## Request

```bash
curl -X POST "https://api.hola.cloud/v1/loggers/logger_xyz789/owners" \
  -H "Api-Key: LOGGER_API_KEY" \
  -H "Api-Secret: LOGGER_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "user_abc123"
  }'
```

```http
POST /v1/loggers/logger_xyz789/owners HTTP/1.1
Host: api.hola.cloud
Api-Key: LOGGER_API_KEY
Api-Secret: LOGGER_API_SECRET
Content-Type: application/json

{
  "user_id": "user_abc123"
}
```

## Response

Returns the updated list of owners.

```json
["user_def456", "user_abc123"]
```

## Error Codes

| Code | Description |
|------|-------------|
| 400 | Missing or invalid `user_id` field |
| 401 | Missing or invalid API credentials |
| 403 | API credentials do not have access to this logger |
| 404 | Logger not found |
