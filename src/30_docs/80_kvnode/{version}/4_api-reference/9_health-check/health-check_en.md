# Health Check

Returns the health status of the KVNode.

## Authentication

This endpoint is public. No authentication required.

## Example Request

```bash
curl "https://api.hola.cloud/healthz"
```

## Example Response

```http
HTTP/1.1 200 OK
Content-Type: application/json
```

```json
{
  "status": "ok"
}
```

## Error Codes

| Status | Code | Description |
|--------|------|-------------|
| 500 | internal_error | Node is unhealthy |
