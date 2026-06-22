# Set Key

Sets the value of a key in a collection. The wildcard `*` in the path is replaced with the key name.

## Authentication

Requires internal authentication. Pass credentials via `X-Glue-Authentication` header, or `apikey` and `secret` headers.

## Path Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| col | string | The name of the collection |

## Request Body

| Field | Type | Description |
|-------|------|-------------|
| value | any | The JSON value to store |

```json
{
  "value": {
    "name": "Alice",
    "email": "alice@example.com",
    "role": "admin"
  }
}
```

## Example Request

```bash
curl -X POST "https://api.hola.cloud/v1/collections/users/keys/user:1001" \
  -H "X-Glue-Authentication: YOUR_AUTH_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "value": {
      "name": "Alice",
      "email": "alice@example.com",
      "role": "admin"
    }
  }'
```

## Example Response

```http
HTTP/1.1 200 OK
Content-Type: application/json
```

```json
{
  "key": "user:1001",
  "collection": "users",
  "created_at": "2025-06-21T10:00:00Z"
}
```

## Error Codes

| Status | Code | Description |
|--------|------|-------------|
| 400 | invalid_request | Missing or invalid value |
| 401 | unauthorized | Missing or invalid authentication |
| 404 | not_found | Collection not found |
| 500 | internal_error | Internal server error |
