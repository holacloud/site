# Delete documents

`POST /v1/databases/{databaseId}/collections/{collection}:remove` · API key

Request: JSON, `Content-Type: application/json` · Response: JSONL, or `204 No Content` with an empty body

| Body field | Default | Meaning |
|---|---|---|
| `filter` | `{}` | Which documents; same syntax as `:find`. |
| `limit` | `1` | Maximum number of documents to delete. `-1` deletes all that match. |

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:remove" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"filter": {"id": "alan"}}'
```

Response `200 OK`, the deleted documents as JSON Lines:

```jsonl
{"id":"alan","name":"Alan Turing","plan":"free","mrr":0}
```
