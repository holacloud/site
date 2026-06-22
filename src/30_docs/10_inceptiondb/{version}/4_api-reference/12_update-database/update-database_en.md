
# Update Database

Updates the properties of an existing database.

## Authentication

Requires `Api-Key`, `Api-Secret`, and `X-Project` headers.

## Path Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| id | uuid | The unique identifier of the database |

## Request Body

| Field | Type | Description |
|-------|------|-------------|
| name | string | New name for the database |

## HTTP Request

```http
PATCH /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890 HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
X-Project: acme-webapp
Content-Type: application/json

{
  "name": "renamed-database"
}
```

## Example

```bash
curl -X PATCH "https://api.hola.cloud/v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890" \
  -H "Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d" \
  -H "Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf" \
  -H "X-Project: acme-webapp" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "renamed-database"
  }'
```

## Response

```json
{
  "id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
  "name": "renamed-database",
  "updated_at": "2025-07-02T09:15:00Z"
}
```

## Error Codes

| Code | Description |
|------|-------------|
| 400 | Missing or invalid request body |
| 401 | Missing or invalid authentication headers |
| 403 | Project access denied |
| 404 | Database not found |
| 409 | A database with the new name already exists |
