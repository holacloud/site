# Find Documents

Finds documents in a collection by full scan or by index traversal.

## Authentication

Requires `Api-Key` and `Api-Secret`, or a Glue owner token where the database owner is allowed.

## Request Body

| Field | Type | Description |
|-------|------|-------------|
| `index` | string | Optional index name. Omit it for a full collection scan. |
| `filter` | object | Optional Connor filter applied to documents. |
| `skip` | integer | Number of matched documents to skip. Default: `0`. |
| `limit` | integer | Maximum number of documents to return. Default: `1`; `0` returns none. |

Index-specific fields can also be included in the body. For example, a `map` index traversal commonly uses `value`.

## HTTP Request

```http
POST /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections/users:find HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
Content-Type: application/json

{"filter":{"role":"admin"},"limit":10}
```

## Response

The response is JSON Lines with matched documents.

```jsonl
{"id":"user-1","name":"Alice","role":"admin"}
```

## Indexed Lookup

```json
{"index":"by-email","value":"alice@example.com","limit":1}
```
