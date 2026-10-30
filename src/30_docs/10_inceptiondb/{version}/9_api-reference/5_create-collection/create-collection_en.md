# Create a collection

`POST /v1/databases/{databaseId}/collections` · API key

Creates an empty collection. Its default values give every document without `id` a random one.

| Body field | Type | Required | Meaning |
|---|---|---|---|
| `name` | string | yes | Name of the collection, used in its URLs. |

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"name": "customers"}'
```

Response `201 Created`:

```json
{"defaults":{"id":"uuid()"},"indexes":0,"name":"customers","total":0}
```

If a collection with that name already exists, the answer is `500` and the error says that it already exists.
