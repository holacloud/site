
# GET /v1/databases/{id}

Retrieves details about a specific database by its ID.

## Authentication

Requires `Api-Key`, `Api-Secret`, and `X-Project` headers.

## Path Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| id | uuid | The unique identifier of the database |

## HTTP Request

```http
GET /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890 HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
X-Project: acme-webapp
```

## Example

```bash
curl -X GET "https://api.hola.cloud/v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890" \
  -H "Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d" \
  -H "Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf" \
  -H "X-Project: acme-webapp"
```

## Response

```json
{
  "id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
  "name": "production-db",
  "created_at": "2025-06-15T10:30:00Z",
  "collection_count": 12,
  "collections": [
    {
      "name": "users",
      "document_count": 5400
    },
    {
      "name": "orders",
      "document_count": 12300
    }
  ]
}
```

## Error Codes

| Code | Description |
|------|-------------|
| 401 | Missing or invalid authentication headers |
| 403 | Project access denied |
| 404 | Database not found |
