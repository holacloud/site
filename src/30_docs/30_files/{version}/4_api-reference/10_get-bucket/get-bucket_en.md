# Get Bucket

Get a bucket by ID for the authenticated user.

## Authentication

Requires `X-Glue-Authentication`.

## Path Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| `bucket_id` | string | The bucket ID |

## Request

```bash
curl "https://api.hola.cloud/v1/buckets/bucket-550e8400-e29b-41d4-a716-446655440000" \
  -H 'X-Glue-Authentication: {"user":{"id":"user-123"}}'
```

## Response

```json
{
  "id": "bucket-550e8400-e29b-41d4-a716-446655440000",
  "project_id": "",
  "created_timestamp": 1782045600000000000,
  "owners": ["user-123"],
  "name": "assets",
  "description": "Application assets"
}
```

## Error Codes

| Status | Description |
|--------|-------------|
| 401 | Missing or invalid `X-Glue-Authentication` |
| 404 | Bucket not found |
| 500 | Persistence error |
