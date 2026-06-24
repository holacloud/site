# Replicate

Replicates data from a parent node to this KVNode.

## Authentication

Requires internal authentication. Pass credentials via `X-Glue-Authentication` header, or `apikey` and `secret` headers.

## Example Request

```bash
curl -X POST "https://api.hola.cloud/v1/replicate" \
  -H "X-Glue-Authentication: YOUR_AUTH_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "entries": [
      {
        "collection": "users",
        "key": "user:1001",
        "value": {"name": "Alice"},
        "seq": 42
      }
    ]
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
  "applied": 1
}
```

## Error Codes

| Status | Code | Description |
|--------|------|-------------|
| 400 | invalid_json | Invalid JSON payload |
| 403 | forbidden | Missing authentication headers |
| 502 | parent_unavailable | Parent node is not reachable |
| 500 | internal_error | Internal server error |
