# Document Operations

Document operations are action endpoints on a collection.

## Authentication

Requires `Api-Key` and `Api-Secret`, or a Glue owner token where the database owner is allowed.

## Path Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| databaseId | uuid | The unique identifier of the database |
| collection | string | The collection name |
| documentId | string | The document ID, for get-by-ID requests |

## Insert: `{collection}:insert`

```http
POST /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections/users:insert HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
Content-Type: application/jsonl

{"id":"user-1","name":"Alice","role":"admin"}
{"id":"user-2","name":"Bob","role":"member"}
```

Response:

```jsonl
{"id":"user-1","name":"Alice","role":"admin"}
{"id":"user-2","name":"Bob","role":"member"}
```

## Find: `{collection}:find`

```http
POST /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections/users:find HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
Content-Type: application/json

{
  "filter": {
    "role": "admin"
  }
}
```

Response:

```jsonl
{"id":"user-1","name":"Alice","role":"admin"}
```

## Get by ID

```http
GET /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections/users/documents/user-1 HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
```

Response:

```json
{
  "id": "user-1",
  "document": {
    "id": "user-1",
    "name": "Alice",
    "role": "admin"
  },
  "source": {
    "type": "index",
    "name": "id"
  }
}
```

The `source` field is omitted when the lookup falls back to a collection scan.

## Streaming Inserts

`insertStream` and `insertFullduplex` are experimental action endpoints that accept streamed JSON input and emit inserted documents as they are processed.

```http
POST /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections/events:insertStream HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
Content-Type: application/jsonl

{"id":"evt-1","type":"signup"}
{"id":"evt-2","type":"login"}
```

Response:

```jsonl
{"id":"evt-1","type":"signup"}
{"id":"evt-2","type":"login"}
```

## Patch: `{collection}:patch`

```http
POST /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections/users:patch HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
Content-Type: application/json

{
  "filter": {
    "id": "user-1"
  },
  "patch": {
    "role": "owner"
  }
}
```

Response:

```jsonl
{"id":"user-1","name":"Alice","role":"owner"}
```

## Remove: `{collection}:remove`

```http
POST /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections/users:remove HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
Content-Type: application/json

{
  "filter": {
    "id": "user-2"
  }
}
```

Response:

```jsonl
{"id":"user-2","name":"Bob","role":"member"}
```
