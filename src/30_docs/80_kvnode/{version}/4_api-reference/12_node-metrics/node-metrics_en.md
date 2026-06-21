# GET /v1/metrics

Returns operational metrics for the KVNode, including write and read counts.

## Authentication

Requires internal authentication. Pass credentials via `X-Glue-Authentication` header, or `apikey` and `secret` headers.

## Example Request

```bash
curl "https://api.hola.cloud/v1/metrics" \
  -H "X-Glue-Authentication: YOUR_AUTH_TOKEN"
```

## Example Response

```http
HTTP/1.1 200 OK
Content-Type: application/json
```

```json
{
  "writes_total": 15000,
  "reads_total": 98000,
  "collections": 3,
  "keys_total": 4500,
  "uptime_seconds": 86400
}
```

## Error Codes

| Status | Code | Description |
|--------|------|-------------|
| 401 | unauthorized | Missing or invalid authentication |
| 500 | internal_error | Internal server error |
