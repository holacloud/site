# Files

HolaCloud Files provides a REST API for storing files in buckets. Buckets group files, and each file is addressed by its path after `/files/` or filtered by prefix after `/list/`.

## Authentication

All `/v1` endpoints require the `X-Glue-Authentication` header. The header value is a JSON object with the authenticated user data:

```http
X-Glue-Authentication: {"user":{"id":"user-123"}}
```

## Data Model

Bucket objects use these fields: `id`, `project_id`, `created_timestamp`, `owners`, `name`, and `description`.

File objects use these fields: `id`, `uuid`, `created_timestamp`, `updated_timestamp`, `owners`, `status`, `size`, `name`, `bucket`, `hash_md5`, `hash_sha256`, and `mime_type`.

## API Reference

All endpoints are under the base URL `https://api.hola.cloud`.

| Method | Path | Description |
|--------|------|-------------|
| GET | `/v1/buckets` | List buckets owned by the authenticated user |
| POST | `/v1/buckets` | Create a bucket |
| GET | `/v1/buckets/{bucket_id}` | Get bucket details |
| PATCH | `/v1/buckets/{bucket_id}` | Not implemented |
| DELETE | `/v1/buckets/{bucket_id}` | Delete a bucket |
| GET | `/v1/buckets/{bucket_id}/list/*` | List files in a bucket, optionally filtering by prefix |
| PUT | `/v1/buckets/{bucket_id}/files/*` | Upload a file |
| GET | `/v1/buckets/{bucket_id}/files/*` | Download a file, or return JSON metadata with `?metadata` |
| DELETE | `/v1/buckets/{bucket_id}/files/*` | Delete a file |
| HEAD | `/v1/buckets/{bucket_id}/files/*` | Return file headers |
