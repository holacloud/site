# Listar colecciones

`GET /v1/databases/{databaseId}/collections` · API key

Devuelve todas las colecciones de la base de datos con su número de documentos, índices y valores por defecto.

```bash
curl "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET"
```

Respuesta `200 OK`:

```json
[{"defaults":{"id":"uuid()"},"indexes":0,"name":"customers","total":3}]
```
