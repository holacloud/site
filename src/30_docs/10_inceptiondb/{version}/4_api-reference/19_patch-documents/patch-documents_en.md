# Patch Documents

Patches documents selected by full scan or by index traversal. The endpoint returns each patched document as JSON Lines.

## Authentication

Requires `Api-Key` and `Api-Secret`, or a Glue owner token where the database owner is allowed.

## Request Body

| Field | Type | Description |
|-------|------|-------------|
| `patch` | object | Patch object to apply to each matched document. |
| `index` | string | Optional index name. Omit it for a full collection scan. |
| `filter` | object | Optional Connor filter applied before patching. |
| `skip` | integer | Number of matched documents to skip. Default: `0`. |
| `limit` | integer | Maximum number of documents to patch. Default: `1`; `0` patches none. |

Index-specific fields can also be included in the body.

## HTTP Request

```http
POST /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections/users:patch HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
Content-Type: application/json

{
  "filter": {"id": "user-1"},
  "patch": {"role": "owner"},
  "limit": 1
}
```

## Response

```jsonl
{"id":"user-1","name":"Alice","role":"owner"}
```
