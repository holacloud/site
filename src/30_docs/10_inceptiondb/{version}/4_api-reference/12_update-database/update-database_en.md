# Update Database

Updates database properties.

## Authentication

Requires `X-Glue-Authentication`.

## Path Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| databaseId | uuid | The unique identifier of the database |

## Request Body

| Field | Type | Description |
|-------|------|-------------|
| name | string | New database name |

## HTTP Request

```http
PATCH /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890 HTTP/1.1
Host: api.hola.cloud
X-Glue-Authentication: your-glue-token
Content-Type: application/json

{
  "name": "renamed-database"
}
```

## Response

```json
{
  "id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
  "name": "renamed-database",
  "creation_date": "2025-06-15T10:30:00Z",
  "owners": ["owner-id"],
  "api_keys": ["api-key-id"],
  "owners_length": 1
}
```
