# GET /v1/buckets/{id}

Get details for a specific bucket by its ID.

## Authentication

Requires `Api-Key` and `Api-Secret` headers.

## Path Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| `id` | string | The bucket ID (e.g., `bkt_abc123`) |

## Request

```bash
curl "https://api.hola.cloud/v1/buckets/bkt_abc123" \
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
  "id": "bkt_abc123",
  "name": "my-first-bucket",
  "createdAt": "2026-06-21T10:00:00Z",
  "size": 1048576,
  "fileCount": 5,
  "public": false
}
```

## Error Codes

| Status | Code | Description |
|--------|------|-------------|
| 401 | Unauthorized | Missing or invalid API credentials |
| 404 | Not Found | The specified bucket does not exist |
| 500 | Internal Server Error | An unexpected error occurred |
