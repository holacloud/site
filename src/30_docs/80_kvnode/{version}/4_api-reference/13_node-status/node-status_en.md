# Node Status

Returns the current status of the KVNode, including collections, replication state, and uptime.

## Authentication

Requires internal authentication. Pass credentials via `X-Glue-Authentication` header, or `apikey` and `secret` headers.

## Example Request

```bash
curl "https://api.hola.cloud/v1/status" \
  -H "X-Glue-Authentication: YOUR_AUTH_TOKEN"
```

## Example Response

```http
HTTP/1.1 200 OK
Content-Type: application/json
```

```json
{
  "node_id": "node-abc123",
  "collections": 3,
  "uptime_seconds": 86400,
  "replication": {
    "role": "primary",
    "connected_replicas": 2
  }
}
```

## Error Codes

| Status | Code | Description |
|--------|------|-------------|
| 401 | unauthorized | Missing or invalid authentication |
| 500 | internal_error | Internal server error |
