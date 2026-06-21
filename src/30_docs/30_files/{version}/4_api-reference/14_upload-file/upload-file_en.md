# PUT /v1/buckets/{id}/files/*

Upload a file to a bucket. The file path is specified after `/files/` in the URL. The maximum file size is 5 GB.

## Authentication

Requires `Api-Key` and `Api-Secret` headers.

## Path Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| `id` | string | The bucket ID (e.g., `bkt_abc123`) |

## Request Headers

| Header | Required | Description |
|--------|----------|-------------|
| `Content-Type` | yes | MIME type of the file being uploaded |
| `Content-Length` | yes | Size of the file in bytes |

## Request

```bash
curl -X PUT "https://api.hola.cloud/v1/buckets/bkt_abc123/files/images/logo.png" \
  -H "Api-Key: YOUR_API_KEY" \
  -H "Api-Secret: YOUR_API_SECRET" \
  -H "Content-Type: image/png" \
  --data-binary @logo.png
```

## Response

```http
HTTP/1.1 201 Created
Content-Type: application/json
```

```json
{
  "path": "images/logo.png",
  "size": 24576,
  "contentType": "image/png",
  "uploadedAt": "2026-06-21T12:00:00Z",
  "etag": "\"abc123def456\""
}
```

## Error Codes

| Status | Code | Description |
|--------|------|-------------|
| 400 | Bad Request | Invalid request or missing Content-Type |
| 401 | Unauthorized | Missing or invalid API credentials |
| 404 | Not Found | The specified bucket does not exist |
| 413 | Payload Too Large | File exceeds the 5 GB size limit |
| 500 | Internal Server Error | An unexpected error occurred |
