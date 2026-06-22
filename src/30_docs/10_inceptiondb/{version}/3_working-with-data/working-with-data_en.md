# Working with Data

Document operations are collection actions. Use `Api-Key` and `Api-Secret` headers for database access, or a Glue owner token where the database owner is allowed.

## Insert Documents

```http
POST /v1/databases/{databaseId}/collections/{collection}:insert
Content-Type: application/jsonl
Api-Key: your-api-key
Api-Secret: your-api-secret

{"id":"1","name":"Alice"}
{"id":"2","name":"Bob"}
```

The response is the inserted documents as JSONL.

```jsonl
{"id":"1","name":"Alice"}
{"id":"2","name":"Bob"}
```

## Find Documents

```http
POST /v1/databases/{databaseId}/collections/{collection}:find
Content-Type: application/json
Api-Key: your-api-key
Api-Secret: your-api-secret

{
  "filter": {
    "name": "Alice"
  }
}
```

The response is matching documents as JSONL.

```jsonl
{"id":"1","name":"Alice"}
```

## Get a Document by ID

```http
GET /v1/databases/{databaseId}/collections/{collection}/documents/{documentId}
Api-Key: your-api-key
Api-Secret: your-api-secret
```

The response is the document JSON.

```json
{
  "id": "1",
  "name": "Alice"
}
```

## Patch Documents

```http
POST /v1/databases/{databaseId}/collections/{collection}:patch
Content-Type: application/json
Api-Key: your-api-key
Api-Secret: your-api-secret

{
  "filter": {
    "id": "1"
  },
  "patch": {
    "name": "Alice Updated"
  }
}
```

The response is patched documents as JSONL.

```jsonl
{"id":"1","name":"Alice Updated"}
```

## Remove Documents

```http
POST /v1/databases/{databaseId}/collections/{collection}:remove
Content-Type: application/json
Api-Key: your-api-key
Api-Secret: your-api-secret

{
  "filter": {
    "id": "1"
  }
}
```

The response is removed documents as JSONL.

```jsonl
{"id":"1","name":"Alice Updated"}
```
