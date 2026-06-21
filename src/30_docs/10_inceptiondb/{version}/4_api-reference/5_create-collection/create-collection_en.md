
# POST /v1/databases/{id}/collections

Creates a new collection within the specified database. Collections are containers for documents.

## Authentication

Requires `Api-Key`, `Api-Secret`, and `X-Project` headers.

## Path Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| id | uuid | The unique identifier of the database |

## Request Body

| Field | Type | Description |
|-------|------|-------------|
| name | string | The name of the collection |

## HTTP Request

```http
POST /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
X-Project: acme-webapp
Content-Type: application/json

{
  "name": "my-collection"
}
```

## Example

```bash
curl -X POST "https://api.hola.cloud/v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections" \
  -H "Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d" \
  -H "Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf" \
  -H "X-Project: acme-webapp" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "my-collection"
  }'
```

## Response

```json
{
  "defaults": {
    "id": "uuid()"
  },
  "indexes": 0,
  "name": "my-collection",
  "total": 0
}
```

## Error Codes

| Code | Description |
|------|-------------|
| 400 | Missing or invalid `name` field |
| 401 | Missing or invalid authentication headers |
| 403 | Project access denied |
| 404 | Database not found |
| 409 | A collection with the same name already exists |
