# List Collections

Returns a list of all collections in the KVNode.

## Authentication

Requires internal authentication. Pass credentials via `X-Glue-Authentication` header, or `apikey` and `secret` headers.

## Example Request

```bash
curl "https://api.hola.cloud/v1/collections" \
  -H "X-Glue-Authentication: YOUR_AUTH_TOKEN"
```

## Example Response

```http
HTTP/1.1 200 OK
Content-Type: application/json
```

```json
{
  "collections": [
    {
      "name": "users",
      "key_count": 1200,
      "created_at": "2025-06-01T10:00:00Z"
    },
    {
      "name": "sessions",
      "key_count": 450,
      "created_at": "2025-06-10T14:30:00Z"
    }
  ]
}
```

## Error Codes

| Status | Code | Description |
|--------|------|-------------|
| 401 | unauthorized | Missing or invalid authentication |
| 500 | internal_error | Internal server error |
