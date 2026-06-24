# Get Document

Gets one document by `id`. The lookup uses a `map` index on field `id` when available, otherwise it falls back to a collection scan.

## Authentication

Requires `Api-Key` and `Api-Secret`, or a Glue owner token where the database owner is allowed.

## HTTP Request

```http
GET /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections/users/documents/user-1 HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
```

## Response

```json
{
  "id": "user-1",
  "document": {"id":"user-1","name":"Alice","role":"admin"},
  "source": {"type":"index","name":"id"}
}
```

`source` is omitted when the endpoint uses a collection scan. Empty document IDs return `400 Bad Request`; missing collections or documents return `404 Not Found`.
