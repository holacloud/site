# POST /v1/collections

Creates a new collection in the KVNode.

## Authentication

Requires internal authentication. Pass credentials via `X-Glue-Authentication` header, or `apikey` and `secret` headers.

## Request Body

```json
{
  "name": "users"
}
```

## Example Request

```bash
curl -X POST "https://api.hola.cloud/v1/collections" \
  -H "X-Glue-Authentication: YOUR_AUTH_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "users"
  }'
```

## Example Response

```http
HTTP/1.1 201 Created
Content-Type: application/json
```

```json
{
  "name": "users",
  "key_count": 0,
  "created_at": "2025-06-21T10:00:00Z"
}
```

## Error Codes

| Status | Code | Description |
|--------|------|-------------|
| 400 | invalid_request | Missing or invalid collection name |
| 401 | unauthorized | Missing or invalid authentication |
| 409 | conflict | Collection already exists |
| 500 | internal_error | Internal server error |
