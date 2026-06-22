# Download File

Download a file from a bucket. The file path is specified after `/files/`.

## Authentication

Requires `X-Glue-Authentication`.

## Path Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| `bucket_id` | string | The bucket ID |
| `*` | string | File path |

## Query Parameters

| Parameter | Description |
|-----------|-------------|
| `metadata` | If present, returns the file object as JSON instead of the file body. |

## Request

```bash
curl "https://api.hola.cloud/v1/buckets/bucket-550e8400-e29b-41d4-a716-446655440000/files/images/logo.png" \
  -H 'X-Glue-Authentication: {"user":{"id":"user-123"}}' \
  -o logo.png
```

## Response

The response body is the raw file content. `Content-Type` is set from the stored `mime_type`.

### Metadata Response

```bash
curl "https://api.hola.cloud/v1/buckets/bucket-550e8400-e29b-41d4-a716-446655440000/files/images/logo.png?metadata" \
  -H 'X-Glue-Authentication: {"user":{"id":"user-123"}}'
```

```json
{
  "id": "file-9f0b7b3c-1d2e-4a5f-8b9c-0123456789ab",
  "uuid": "9f0b7b3c-1d2e-4a5f-8b9c-0123456789ab",
  "created_timestamp": 1782045660000000000,
  "updated_timestamp": 1782045660000000000,
  "owners": ["user-123"],
  "status": "available",
  "size": 24576,
  "name": "images/logo.png",
  "bucket": "bucket-550e8400-e29b-41d4-a716-446655440000",
  "hash_md5": "example-md5",
  "hash_sha256": "example-sha256",
  "mime_type": "image/png"
}
```

## Error Codes

| Status | Description |
|--------|-------------|
| 401 | Missing or invalid `X-Glue-Authentication` |
| 404 | File not found |
| 500 | Persistence or filesystem error |
