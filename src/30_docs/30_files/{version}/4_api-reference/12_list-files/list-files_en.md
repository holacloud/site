# List Files

List files in a bucket. The wildcard (`*`) can be replaced with an optional prefix path to filter results.

## Authentication

Requires `Api-Key` and `Api-Secret` headers.

## Path Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| `id` | string | The bucket ID (e.g., `bkt_abc123`) |

## Query Parameters

| Parameter | Type | Default | Description |
|-----------|------|---------|-------------|
| `prefix` | string | "" | Filter files starting with this prefix |
| `delimiter` | string | "" | Group results by this delimiter (e.g., `/`) |
| `maxKeys` | integer | 1000 | Maximum number of files to return |

## Request

```bash
curl "https://api.hola.cloud/v1/buckets/bkt_abc123/list/*?prefix=images/&maxKeys=50" \
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
  "files": [
    {
      "path": "images/logo.png",
      "size": 24576,
      "contentType": "image/png",
      "modifiedAt": "2026-06-21T11:00:00Z",
      "etag": "\"abc123def456\""
    },
    {
      "path": "images/banner.jpg",
      "size": 102400,
      "contentType": "image/jpeg",
      "modifiedAt": "2026-06-21T10:30:00Z",
      "etag": "\"789012ghi345\""
    }
  ],
  "prefix": "images/",
  "total": 2
}
```

## Error Codes

| Status | Code | Description |
|--------|------|-------------|
| 401 | Unauthorized | Missing or invalid API credentials |
| 404 | Not Found | The specified bucket does not exist |
| 500 | Internal Server Error | An unexpected error occurred |
