# PATCH /v1/buckets/{id}

Modify metadata for an existing bucket. Only the fields provided in the request body are updated.

## Authentication

Requires `Api-Key` and `Api-Secret` headers.

## Path Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| `id` | string | The bucket ID (e.g., `bkt_abc123`) |

## Request Body

```json
{
  "name": "renamed-bucket",
  "public": true
}
```

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `name` | string | no | New globally unique bucket name |
| `public` | boolean | no | Set to `true` to allow public access to files |

## Request

```bash
curl -X PATCH "https://api.hola.cloud/v1/buckets/bkt_abc123" \
  -H "Api-Key: YOUR_API_KEY" \
  -H "Api-Secret: YOUR_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "renamed-bucket",
    "public": true
  }'
```

## Response

```http
HTTP/1.1 200 OK
Content-Type: application/json
```

```json
{
  "id": "bkt_abc123",
  "name": "renamed-bucket",
  "createdAt": "2026-06-21T10:00:00Z",
  "size": 1048576,
  "fileCount": 5,
  "public": true
}
```

## Error Codes

| Status | Code | Description |
|--------|------|-------------|
| 400 | Bad Request | Invalid request body |
| 401 | Unauthorized | Missing or invalid API credentials |
| 404 | Not Found | The specified bucket does not exist |
| 409 | Conflict | The new name is already taken |
| 500 | Internal Server Error | An unexpected error occurred |
