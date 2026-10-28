# List collections

`GET /v1/databases/{databaseId}/collections` · API key

Returns every collection of the database with its number of documents, indexes and default values.

```bash
curl "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET"
```

Response `200 OK`:

```json
[{"defaults":{"id":"uuid()"},"indexes":0,"name":"customers","total":3}]
```
