# GET /v1/buckets

List all buckets for the authenticated account.

## Authentication

Requires `Api-Key` and `Api-Secret` headers.

## Query Parameters

| Parameter | Type | Default | Description |
|-----------|------|---------|-------------|
| `limit` | integer | 100 | Maximum number of buckets to return |
| `offset` | integer | 0 | Number of buckets to skip |

## Request

```bash
curl "https://api.hola.cloud/v1/buckets?limit=10&offset=0" \
  -H "Api-Key: YOUR_API_KEY" \
  -H "Api-Secret: YOUR_API_SECRET"
```

## Response

```http
HTTP/1.1 200 OK
Content-Type: application/json
```

```json
{
  "buckets": [
    {
      "id": "bkt_abc123",
      "name": "my-first-bucket",
      "createdAt": "2026-06-21T10:00:00Z",
      "size": 1048576,
      "fileCount": 5,
      "public": false
    },
    {
      "id": "bkt_def456",
      "name": "assets",
      "createdAt": "2026-06-20T08:30:00Z",
      "size": 52428800,
      "fileCount": 42,
      "public": true
    }
  ],
  "total": 2
}
```

## Error Codes

| Status | Code | Description |
|--------|------|-------------|
| 401 | Unauthorized | Missing or invalid API credentials |
| 403 | Forbidden | Insufficient permissions |
| 500 | Internal Server Error | An unexpected error occurred |
