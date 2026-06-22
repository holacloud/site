# Get Database

Retrieves details about a database by ID.

## Authentication

Requires `X-Glue-Authentication`.

## Path Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| databaseId | uuid | The unique identifier of the database |

## HTTP Request

```http
GET /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890 HTTP/1.1
Host: api.hola.cloud
X-Glue-Authentication: your-glue-token
```

## Response

```json
{
  "id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
  "name": "production-db",
  "creation_date": "2025-06-15T10:30:00Z",
  "owners": ["owner-id"],
  "api_keys": ["api-key-id"],
  "owners_length": 1
}
```
