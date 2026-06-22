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
  "collections": ["users", "sessions"]
}
```

## Error Codes

| Status | Code | Description |
|--------|------|-------------|
| 403 | forbidden | Missing authentication headers |
| 500 | internal_error | Internal server error |
