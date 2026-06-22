# List Files

List files in a bucket. The path after `/list/` is used as a prefix filter.

## Authentication

Requires `X-Glue-Authentication`.

## Path Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| `bucket_id` | string | The bucket ID |
| `*` | string | Prefix filter |

## Request

```bash
curl "https://api.hola.cloud/v1/buckets/bucket-550e8400-e29b-41d4-a716-446655440000/list/images/" \
  -H 'X-Glue-Authentication: {"user":{"id":"user-123"}}'
```

## Response

The response is a JSON array.

```json
[
  {
    "id": "file-9f0b7b3c-1d2e-4a5f-8b9c-0123456789ab",
    "name": "images/logo.png",
    "bucket": "bucket-550e8400-e29b-41d4-a716-446655440000",
    "created_timestamp": 1782045660000000000,
    "updated_timestamp": 1782045660000000000,
    "size": 24576,
    "mime_type": "image/png",
    "hash_md5": "example-md5",
    "hash_sha256": "example-sha256",
    "status": "available",
    "owners": ["user-123"]
  }
]
```

## Error Codes

| Status | Description |
|--------|-------------|
| 401 | Missing or invalid `X-Glue-Authentication` |
| 500 | Persistence error |
