# Delete Bucket

Delete a bucket by ID.

## Authentication

Requires `X-Glue-Authentication`.

## Path Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| `bucket_id` | string | The bucket ID |

## Request

```bash
curl -X DELETE "https://api.hola.cloud/v1/buckets/bucket-550e8400-e29b-41d4-a716-446655440000" \
  -H 'X-Glue-Authentication: {"user":{"id":"user-123"}}'
```

## Response

The response body is the deleted bucket object.

```json
{
  "id": "bucket-550e8400-e29b-41d4-a716-446655440000",
  "project_id": "",
  "created_timestamp": 1782045600000000000,
  "owners": ["user-123"],
  "name": "my-new-bucket",
  "description": "Optional bucket description"
}
```

## Error Codes

| Status | Description |
|--------|-------------|
| 401 | Missing or invalid `X-Glue-Authentication` |
| 404 | Bucket not found |
| 500 | Persistence error |
