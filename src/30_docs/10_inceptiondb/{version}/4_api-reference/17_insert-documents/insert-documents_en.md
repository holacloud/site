# Insert Documents

Inserts one or more JSON documents into a collection. The collection is created automatically if it does not exist.

## Authentication

Requires `Api-Key` and `Api-Secret`, or a Glue owner token where the database owner is allowed.

## HTTP Request

```http
POST /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections/users:insert HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
Content-Type: application/jsonl

{"id":"user-1","name":"Alice","role":"admin"}
{"id":"user-2","name":"Bob","role":"member"}
```

## Response

Status: `201 Created` when at least one document is inserted. The response is JSON Lines with the inserted documents.

```jsonl
{"id":"user-1","name":"Alice","role":"admin"}
{"id":"user-2","name":"Bob","role":"member"}
```

If the request body is empty, the response is `204 No Content`. Unique index conflicts return `409 Conflict` before any response body is written.
