# Get Key

Gets the value of a key in a collection. The wildcard `*` in the path is replaced with the key name.

## Authentication

Requires internal authentication. Pass credentials via `X-Glue-Authentication` header, or `apikey` and `secret` headers.

## Path Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| col | string | The name of the collection |

## Example Request

```bash
curl "https://api.hola.cloud/v1/collections/users/keys/user:1001" \
  -H "X-Glue-Authentication: YOUR_AUTH_TOKEN"
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
  "value": {
    "name": "Alice",
    "email": "alice@example.com",
    "role": "admin"
  },
  "created_at": "2025-06-21T10:00:00Z",
  "updated_at": "2025-06-21T10:00:00Z"
}
```

## Error Codes

| Status | Code | Description |
|--------|------|-------------|
| 401 | unauthorized | Missing or invalid authentication |
| 404 | not_found | Collection or key not found |
| 500 | internal_error | Internal server error |
