# Files API Reference

Base URL: `https://api.hola.cloud`

All `/v1` endpoints require `X-Glue-Authentication`.

```http
X-Glue-Authentication: {"user":{"id":"user-123"}}
```

## Endpoints

| Method | Path | Description |
|--------|------|-------------|
| GET | `/v1/buckets` | List buckets |
| POST | `/v1/buckets` | Create a bucket |
| GET | `/v1/buckets/{bucket_id}` | Get bucket details |
| PATCH | `/v1/buckets/{bucket_id}` | Not implemented |
| DELETE | `/v1/buckets/{bucket_id}` | Delete a bucket |
| GET | `/v1/buckets/{bucket_id}/list/*` | List files by prefix |
| PUT | `/v1/buckets/{bucket_id}/files/*` | Upload a file |
| GET | `/v1/buckets/{bucket_id}/files/*` | Download a file |
| DELETE | `/v1/buckets/{bucket_id}/files/*` | Delete a file |
| HEAD | `/v1/buckets/{bucket_id}/files/*` | Return file headers |

## Object Fields

Bucket fields: `id`, `project_id`, `created_timestamp`, `owners`, `name`, `description`.

File fields: `id`, `uuid`, `created_timestamp`, `updated_timestamp`, `owners`, `status`, `size`, `name`, `bucket`, `hash_md5`, `hash_sha256`, `mime_type`.

## Common Errors

| Status | Description |
|--------|-------------|
| 401 | Missing or invalid `X-Glue-Authentication` |
| 404 | Bucket or file not found |
| 409 | Bucket already exists |
| 500 | Persistence or filesystem error |
