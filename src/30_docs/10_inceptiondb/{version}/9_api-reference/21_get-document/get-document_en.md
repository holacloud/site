# Get a document by id

`GET /v1/databases/{databaseId}/collections/{collection}/documents/{id}` · API key

```bash
curl "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers/documents/ada" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET"
```

Response `200 OK`. The document is in `document`; `source` says whether an index on `id` answered (`index`) or the collection was read (`fullscan`).

```json
{"id":"ada","document":{"id":"ada","mrr":120,"name":"Ada Lovelace","plan":"pro"},"source":{"type":"fullscan"}}
```

A missing document answers `404 Not Found`.
