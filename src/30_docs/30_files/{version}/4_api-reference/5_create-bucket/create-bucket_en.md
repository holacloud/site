# Create Bucket

Create a bucket owned by the authenticated user.

## Authentication

Requires `X-Glue-Authentication`.

## Request Body

```json
{
  "name": "my-new-bucket",
  "description": "Optional bucket description"
}
```

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `name` | string | no | Letters, digits, `_`, and `-`, up to 64 characters. If empty, the generated bucket ID is used. |
| `description` | string | no | Description up to 4096 characters. |

## Request

```bash
curl -X POST "https://api.hola.cloud/v1/buckets" \
  -H 'X-Glue-Authentication: {"user":{"id":"user-123"}}' \
  -H "Content-Type: application/json" \
  -d '{"name":"my-new-bucket","description":"Optional bucket description"}'
```

## Response

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
| 409 | Bucket already exists |
| 500 | Persistence error |
