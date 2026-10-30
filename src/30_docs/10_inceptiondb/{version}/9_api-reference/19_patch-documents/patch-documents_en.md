# Update documents

`POST /v1/databases/{databaseId}/collections/{collection}:patch` · API key

Request: JSON, `Content-Type: application/json` · Response: JSONL, or `204 No Content` with an empty body

Applies a JSON merge patch to the matching documents: fields in `patch` are added or replaced, and `null` removes a field. The `id` can change too; with a `map` index on `id`, a patch to an id that already exists is not applied.

| Body field | Default | Meaning |
|---|---|---|
| `filter` | `{}` | Which documents; same syntax as `:find`. |
| `patch` | | The changes. |
| `limit` | `1` | Maximum number of documents to change. `-1` changes all that match. |

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:patch" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"filter": {"id": "alan"}, "patch": {"plan": "pro", "mrr": 120}}'
```

Response `200 OK`, the changed documents as JSON Lines:

```jsonl
{"id":"alan","name":"Alan Turing","plan":"pro","mrr":120}
```
