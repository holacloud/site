# Insert documents

`POST /v1/databases/{databaseId}/collections/{collection}:insert` · API key

The body is JSON Lines: one JSON object per line. Missing fields take the default values of the collection.

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:insert" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  --data-binary '{"id":"ada","name":"Ada Lovelace","plan":"pro","mrr":120}
{"id":"alan","name":"Alan Turing","plan":"free","mrr":0}'
```

Response `201 Created`, the stored documents as JSON Lines:

```jsonl
{"id":"ada","name":"Ada Lovelace","plan":"pro","mrr":120}
{"id":"alan","name":"Alan Turing","plan":"free","mrr":0}
```

Documents are stored in order. The insert stops at the first rejected document; the previous ones stay stored.

| Code | When |
|---|---|
| 400 | A line is not a JSON object, or has repeated field names. |
| 409 | A unique index already has the value. |
