# Managing Collections

Collection endpoints use database access credentials: `Api-Key` and `Api-Secret`, or a Glue owner token where the database owner is allowed.

## Create a Collection

```bash
curl -X POST "https://api.hola.cloud/v1/databases/{databaseId}/collections" \
  -H "Api-Key: {api_key}" \
  -H "Api-Secret: {api_secret}" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "my-collection"
  }'
```

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

## List Collections

```bash
curl "https://api.hola.cloud/v1/databases/{databaseId}/collections" \
  -H "Api-Key: {api_key}" \
  -H "Api-Secret: {api_secret}"
```

```json
[
  {
    "name": "my-collection",
    "total": 3,
    "indexes": 0,
    "defaults": {
      "id": "uuid()"
    }
  }
]
```

## Insert Documents

Collections can receive JSONL documents through the `:insert` action.

```bash
curl -X POST "https://api.hola.cloud/v1/databases/{databaseId}/collections/my-collection:insert" \
  -H "Api-Key: {api_key}" \
  -H "Api-Secret: {api_secret}" \
  -H "Content-Type: application/jsonl" \
  --data-binary '{"id":"1","name":"Alfonso"}
{"id":"2","name":"Gerardo"}'
```

The response is JSONL with the inserted documents.

```jsonl
{"id":"1","name":"Alfonso"}
{"id":"2","name":"Gerardo"}
```
