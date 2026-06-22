# Create Collection

Creates a new collection within the specified database.

## Authentication

Requires `Api-Key` and `Api-Secret`, or a Glue owner token where the database owner is allowed.

## Path Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| databaseId | uuid | The unique identifier of the database |

## Request Body

| Field | Type | Description |
|-------|------|-------------|
| name | string | The collection name |

## HTTP Request

```http
POST /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
Content-Type: application/json

{
  "name": "my-collection"
}
```

## Response

```json
{
  "name": "my-collection",
  "total": 0,
  "indexes": 0,
  "defaults": {
    "id": "uuid()"
  }
}
```
