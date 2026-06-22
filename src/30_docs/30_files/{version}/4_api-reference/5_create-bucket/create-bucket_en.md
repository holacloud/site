# Create Bucket

Create a new bucket. The bucket name must be globally unique across all accounts.

## Authentication

Requires `Api-Key` and `Api-Secret` headers.

## Request Body

```json
{
  "name": "my-new-bucket",
  "public": false
}
```

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `name` | string | yes | Globally unique bucket name (3-63 characters, lowercase letters, numbers, hyphens) |
| `public` | boolean | no | Whether files in this bucket can be accessed without authentication (default: `false`) |

## Request

```bash
curl -X POST "https://api.hola.cloud/v1/buckets" \
  -H "Api-Key: YOUR_API_KEY" \
  -H "Api-Secret: YOUR_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "my-new-bucket",
    "public": false
  }'
```

## Response

```http
HTTP/1.1 201 Created
Content-Type: application/json
```

```json
{
  "id": "bkt_xyz789",
  "name": "my-new-bucket",
  "createdAt": "2026-06-21T12:00:00Z",
  "size": 0,
  "fileCount": 0,
  "public": false
}
```

## Error Codes

| Status | Code | Description |
|--------|------|-------------|
| 400 | Bad Request | Invalid bucket name or missing required fields |
| 401 | Unauthorized | Missing or invalid API credentials |
| 409 | Conflict | A bucket with this name already exists |
| 500 | Internal Server Error | An unexpected error occurred |
