# GET /readyz

Returns the readiness status of the KVNode, including the health of its parent connection.

## Authentication

This endpoint is public. No authentication required.

## Example Request

```bash
curl "https://api.hola.cloud/readyz"
```

## Example Response

```http
HTTP/1.1 200 OK
Content-Type: application/json
```

```json
{
  "status": "ok",
  "parent_connected": true
}
```

## Error Codes

| Status | Code | Description |
|--------|------|-------------|
| 503 | unavailable | Node is not ready |
