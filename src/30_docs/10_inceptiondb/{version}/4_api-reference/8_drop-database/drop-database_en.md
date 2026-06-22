
# Drop Database

Permanently drops a database and all its collections and documents. This action cannot be undone.

## Authentication

Requires `Api-Key`, `Api-Secret`, and `X-Project` headers.

## Path Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| id | uuid | The unique identifier of the database to drop |

## HTTP Request

```http
DELETE /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890 HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
X-Project: acme-webapp
```

## Example

```bash
curl -X DELETE "https://api.hola.cloud/v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890" \
  -H "Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d" \
  -H "Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf" \
  -H "X-Project: acme-webapp"
```

## Response

```json
{
  "message": "Database dropped successfully",
  "id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890"
}
```

## Error Codes

| Code | Description |
|------|-------------|
| 401 | Missing or invalid authentication headers |
| 403 | Project access denied |
| 404 | Database not found |
