# GET /v1/buckets/{id}/files/*

Download a file from a bucket. The file path is specified after `/files/` in the URL.

## Authentication

Requires `Api-Key` and `Api-Secret` headers.

## Path Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| `id` | string | The bucket ID (e.g., `bkt_abc123`) |

## Query Parameters

| Parameter | Type | Default | Description |
|-----------|------|---------|-------------|
| `metadata` | boolean | false | Set to `true` to return file metadata as JSON instead of the file body |

## Request

```bash
curl "https://api.hola.cloud/v1/buckets/bkt_abc123/files/images/logo.png" \
  -H "Api-Key: YOUR_API_KEY" \
  -H "Api-Secret: YOUR_API_SECRET" \
  -o logo.png
```

## Response

```http
HTTP/1.1 200 OK
Content-Type: image/png
Content-Length: 24576
Last-Modified: Sun, 21 Jun 2026 12:00:00 GMT
ETag: "abc123def456"
```

The response body is the raw file content.

### Metadata Response

When `?metadata=true` is specified, the response returns JSON instead of the file body:

```json
{
  "path": "images/logo.png",
  "size": 24576,
  "contentType": "image/png",
  "modifiedAt": "2026-06-21T12:00:00Z",
  "etag": "\"abc123def456\""
}
```

## Error Codes

| Status | Code | Description |
|--------|------|-------------|
| 401 | Unauthorized | Missing or invalid API credentials |
| 404 | Not Found | The specified bucket or file does not exist |
| 500 | Internal Server Error | An unexpected error occurred |
