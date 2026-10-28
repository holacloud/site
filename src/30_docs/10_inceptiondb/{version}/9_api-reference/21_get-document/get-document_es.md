# Obtener un documento por id

`GET /v1/databases/{databaseId}/collections/{collection}/documents/{id}` · API key

```bash
curl "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers/documents/ada" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET"
```

Respuesta `200 OK`. El documento está en `document`; `source` indica si respondió un índice sobre `id` (`index`) o se recorrió la colección (`fullscan`).

```json
{"id":"ada","document":{"id":"ada","mrr":120,"name":"Ada Lovelace","plan":"pro"},"source":{"type":"fullscan"}}
```

Si el documento no existe, la respuesta es `404 Not Found`.
