# Borrar documentos

`POST /v1/databases/{databaseId}/collections/{collection}:remove` · API key

| Campo del cuerpo | Por defecto | Significado |
|---|---|---|
| `filter` | `{}` | Qué documentos; misma sintaxis que `:find`. |
| `limit` | `1` | Número máximo de documentos a borrar. `-1` borra todos los que coinciden. |

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:remove" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -d '{"filter": {"id": "alan"}}'
```

Respuesta `200 OK`, los documentos borrados como JSON Lines:

```jsonl
{"id":"alan","name":"Alan Turing","plan":"free","mrr":0}
```
