# File Metadata

Retrieve metadata for a file without downloading its contents. The file path is specified after `/files/` in the URL.

## Authentication

Requires `Api-Key` and `Api-Secret` headers.

## Path Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| `id` | string | The bucket ID (e.g., `bkt_abc123`) |

## Request

```bash
curl -I "https://api.hola.cloud/v1/buckets/bkt_abc123/files/images/logo.png" \
  -H "Api-Key: YOUR_API_KEY" \
  -H "Api-Secret: YOUR_API_SECRET"
```

## Response

```http
HTTP/1.1 200 OK
Content-Type: image/png
Content-Length: 24576
Last-Modified: Sun, 21 Jun 2026 12:00:00 GMT
ETag: "abc123def456"
Accept-Ranges: bytes
```

The response contains only headers; no body is returned.

## Error Codes

| Status | Code | Description |
|--------|------|-------------|
| 401 | Unauthorized | Missing or invalid API credentials |
| 404 | Not Found | The specified bucket or file does not exist |
| 500 | Internal Server Error | An unexpected error occurred |
