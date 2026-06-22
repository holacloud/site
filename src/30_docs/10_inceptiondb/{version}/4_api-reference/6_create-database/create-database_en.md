
# Create Database

Creates a new database within the authenticated project.

## Authentication

Requires `Api-Key`, `Api-Secret`, and `X-Project` headers.

## Request Body

| Field | Type | Description |
|-------|------|-------------|
| name | string | A human-readable name for the database |

## HTTP Request

```http
POST /v1/databases HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
X-Project: acme-webapp
Content-Type: application/json

{
  "name": "my-database"
}
```

## Example

```bash
curl -X POST "https://api.hola.cloud/v1/databases" \
  -H "Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d" \
  -H "Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf" \
  -H "X-Project: acme-webapp" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "my-database"
  }'
```

## Response

```json
{
  "id": "c3d4e5f6-a7b8-9012-cdef-123456789012",
  "name": "my-database",
  "created_at": "2025-07-01T14:00:00Z",
  "collection_count": 0
}
```

## Error Codes

| Code | Description |
|------|-------------|
| 400 | Missing or invalid `name` field |
| 401 | Missing or invalid authentication headers |
| 403 | Project access denied |
| 409 | A database with the same name already exists |
