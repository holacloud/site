# Managing Buckets

Buckets are containers for files. The implemented bucket operations are create, list, get, and delete.

## Bucket Input

`POST /v1/buckets` accepts `name` and `description`.

- `name` is trimmed, may be empty, and can contain letters, digits, `_`, and `-` up to 64 characters.
- `description` can be up to 4096 characters.
- If `name` is empty, the generated bucket ID is used as the name.

## Create a Bucket

```bash
curl -X POST "https://api.hola.cloud/v1/buckets" \
  -H 'X-Glue-Authentication: {"user":{"id":"user-123"}}' \
  -H "Content-Type: application/json" \
  -d '{"name":"assets","description":"Application assets"}'
```

Response:

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

## List Buckets

```bash
curl "https://api.hola.cloud/v1/buckets" \
  -H 'X-Glue-Authentication: {"user":{"id":"user-123"}}'
```

Response:

```json
[
  {
    "id": "bucket-550e8400-e29b-41d4-a716-446655440000",
    "name": "assets",
    "description": "Application assets",
    "created_timestamp": 1782045600000000000,
    "created_h": "2026-06-21T10:00:00Z",
    "owners": ["user-123"]
  }
]
```

## Get Bucket Details

```bash
curl "https://api.hola.cloud/v1/buckets/bucket-550e8400-e29b-41d4-a716-446655440000" \
  -H 'X-Glue-Authentication: {"user":{"id":"user-123"}}'
```

Response:

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

## Modify Bucket

`PATCH /v1/buckets/{bucket_id}` is registered but not implemented.

## Delete Bucket

```bash
curl -X DELETE "https://api.hola.cloud/v1/buckets/bucket-550e8400-e29b-41d4-a716-446655440000" \
  -H 'X-Glue-Authentication: {"user":{"id":"user-123"}}'
```

The response body is the deleted bucket object.
