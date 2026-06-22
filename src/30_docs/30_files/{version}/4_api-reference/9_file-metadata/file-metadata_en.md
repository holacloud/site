# File Headers

Return headers for a file with `HEAD /v1/buckets/{bucket_id}/files/*`.

## Authentication

Requires `X-Glue-Authentication`.

## Path Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| `bucket_id` | string | The bucket ID |
| `*` | string | File path |

## Request

```bash
curl -I "https://api.hola.cloud/v1/buckets/bucket-550e8400-e29b-41d4-a716-446655440000/files/images/logo.png" \
  -H 'X-Glue-Authentication: {"user":{"id":"user-123"}}'
```

## Response

The handler sets `Last-Modified` and `Content-Length`.

```http
HTTP/1.1 200 OK
Last-Modified: 2026-06-21T10:01:00Z
Content-Length: 24576
```

## Error Codes

| Status | Description |
|--------|-------------|
| 401 | Missing or invalid `X-Glue-Authentication` |
| 404 | File not found |
| 500 | Persistence or filesystem error |
