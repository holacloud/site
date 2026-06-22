# Create Database

Creates a new database.

## Authentication

Requires `X-Glue-Authentication`.

## Request Body

| Field | Type | Description |
|-------|------|-------------|
| name | string | Database name |

## HTTP Request

```http
POST /v1/databases HTTP/1.1
Host: api.hola.cloud
X-Glue-Authentication: your-glue-token
Content-Type: application/json

{
  "name": "my-database"
}
```

## Response

```json
{
  "id": "c3d4e5f6-a7b8-9012-cdef-123456789012",
  "name": "my-database",
  "creation_date": "2025-07-01T14:00:00Z",
  "owners": ["owner-id"],
  "api_keys": [],
  "owners_length": 1
}
```
