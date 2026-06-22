# Create Collection

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
HTTP/1.1 200 OK
Content-Type: application/json
```

```json
{
  "ok": true,
  "collection": "users"
}
```

## Error Codes

| Status | Code | Description |
|--------|------|-------------|
| 400 | invalid_json | Invalid JSON payload |
| 400 | missing_collection | Collection is required |
| 403 | forbidden | Missing authentication headers |
| 500 | internal_error | Internal server error |
