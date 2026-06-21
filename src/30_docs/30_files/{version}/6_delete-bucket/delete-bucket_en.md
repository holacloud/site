# DELETE /v1/buckets/{id}

Delete an empty bucket. The bucket must not contain any files.

## Authentication

Requires `Api-Key` and `Api-Secret` headers.

## Path Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| `id` | string | The bucket ID (e.g., `bkt_abc123`) |

## Request

```bash
curl -X DELETE "https://api.hola.cloud/v1/buckets/bkt_abc123" \
  -H "Api-Key: YOUR_API_KEY" \
  -H "Api-Secret: YOUR_API_SECRET"
```

## Response

```http
HTTP/1.1 204 No Content
```

## Error Codes

| Status | Code | Description |
|--------|------|-------------|
| 401 | Unauthorized | Missing or invalid API credentials |
| 404 | Not Found | The specified bucket does not exist |
| 409 | Conflict | Bucket is not empty; delete all files first |
| 500 | Internal Server Error | An unexpected error occurred |
