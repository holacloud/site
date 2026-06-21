
# POST /v1/databases/{id}/collections/{col}

Performs document operations on a collection. The specific operation is determined by a colon-suffixed action on the collection name.

## Authentication

Requires `Api-Key`, `Api-Secret`, and `X-Project` headers.

## Path Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| id | uuid | The unique identifier of the database |
| col | string | The collection name with operation suffix |

## Operations

### Insert — `{collection}:insert`

Inserts one or more documents into the collection.

```http
POST /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections/users:insert HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
X-Project: acme-webapp
Content-Type: application/json

{
  "name": "Alice",
  "email": "alice@example.com",
  "role": "admin"
}
```

```bash
curl -X POST "https://api.hola.cloud/v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections/users:insert" \
  -H "Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d" \
  -H "Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf" \
  -H "X-Project: acme-webapp" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Alice",
    "email": "alice@example.com",
    "role": "admin"
  }'
```

Response:
```json
{
  "id": "d7e8f9a0-b1c2-3456-7890-123456789abc",
  "message": "Document inserted successfully"
}
```

### Find — `{collection}:find`

Queries documents using a filter.

```http
POST /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections/users:find HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
X-Project: acme-webapp
Content-Type: application/json

{
  "filter": {
    "role": "admin"
  }
}
```

```bash
curl -X POST "https://api.hola.cloud/v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections/users:find" \
  -H "Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d" \
  -H "Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf" \
  -H "X-Project: acme-webapp" \
  -H "Content-Type: application/json" \
  -d '{
    "filter": {
      "role": "admin"
    }
  }'
```

Response:
```json
{
  "documents": [
    {
      "id": "d7e8f9a0-b1c2-3456-7890-123456789abc",
      "name": "Alice",
      "email": "alice@example.com",
      "role": "admin"
    }
  ]
}
```

### Remove — `{collection}:remove`

Deletes documents matching a filter.

```http
POST /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections/users:remove HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
X-Project: acme-webapp
Content-Type: application/json

{
  "filter": {
    "email": "alice@example.com"
  }
}
```

```bash
curl -X POST "https://api.hola.cloud/v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections/users:remove" \
  -H "Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d" \
  -H "Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf" \
  -H "X-Project: acme-webapp" \
  -H "Content-Type: application/json" \
  -d '{
    "filter": {
      "email": "alice@example.com"
    }
  }'
```

Response:
```json
{
  "deleted": 1,
  "message": "Documents removed successfully"
}
```

### Patch — `{collection}:patch`

Updates documents matching a filter.

```http
POST /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections/users:patch HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
X-Project: acme-webapp
Content-Type: application/json

{
  "filter": {
    "role": "admin"
  },
  "patch": {
    "role": "superadmin"
  }
}
```

```bash
curl -X POST "https://api.hola.cloud/v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections/users:patch" \
  -H "Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d" \
  -H "Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf" \
  -H "X-Project: acme-webapp" \
  -H "Content-Type: application/json" \
  -d '{
    "filter": {
      "role": "admin"
    },
    "patch": {
      "role": "superadmin"
    }
  }'
```

Response:
```json
{
  "updated": 1,
  "message": "Documents patched successfully"
}
```

## Error Codes

| Code | Description |
|------|-------------|
| 400 | Invalid operation name or malformed request body |
| 401 | Missing or invalid authentication headers |
| 403 | Project access denied |
| 404 | Database or collection not found |
