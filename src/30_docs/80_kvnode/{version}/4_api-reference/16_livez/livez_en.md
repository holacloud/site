# Live Check

Returns the live status of the KVNode.

## Authentication

This endpoint is public. No authentication required.

## Example Request

```bash
curl "https://api.hola.cloud/livez"
```

## Example Response

```http
HTTP/1.1 200 OK
Content-Type: application/json
```

```json
{
  "ok": true,
  "node": "node-abc123",
  "role": "leader"
}
```

## Error Codes

| Status | Code | Description |
|--------|------|-------------|
| 500 | internal_error | Node is unhealthy |
