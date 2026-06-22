# Quick Start Guide with InceptionDB

This guide creates a collection, inserts three JSON documents, and queries filtered documents.

Database access endpoints accept `Api-Key` and `Api-Secret` headers. A Glue owner token can also be used where the database owner is allowed.

## Step 1: Create a Collection

```bash
curl -X POST "https://api.hola.cloud/v1/databases/your-database-id/collections" \
  -H "Api-Key: your-api-key" \
  -H "Api-Secret: your-api-secret" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "my-collection"
  }'
```

Expected response:

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

## Step 2: Insert Documents

Use the collection action endpoint and send JSONL: one JSON document per line.

```bash
curl -X POST "https://api.hola.cloud/v1/databases/your-database-id/collections/my-collection:insert" \
  -H "Api-Key: your-api-key" \
  -H "Api-Secret: your-api-secret" \
  -H "Content-Type: application/jsonl" \
  --data-binary '{"name":"Element 1","value":100}
{"name":"Element 2","value":200}
{"name":"Element 3","value":300}'
```

Expected response:

```jsonl
{"id":"document-id-1","name":"Element 1","value":100}
{"id":"document-id-2","name":"Element 2","value":200}
{"id":"document-id-3","name":"Element 3","value":300}
```

## Step 3: Find Filtered Documents

```bash
curl -X POST "https://api.hola.cloud/v1/databases/your-database-id/collections/my-collection:find" \
  -H "Api-Key: your-api-key" \
  -H "Api-Secret: your-api-secret" \
  -H "Content-Type: application/json" \
  -d '{
    "filter": {
      "value": { "$gte": 200 }
    }
  }'
```

Expected response:

```jsonl
{"id":"document-id-2","name":"Element 2","value":200}
{"id":"document-id-3","name":"Element 3","value":300}
```

## Summary

You created a collection, inserted JSONL documents through `:insert`, and queried documents through `:find`.
