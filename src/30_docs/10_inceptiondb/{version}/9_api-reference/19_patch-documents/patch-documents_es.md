# Actualizar documentos

`POST /v1/databases/{databaseId}/collections/{collection}:patch` · API key

Petición: JSON, `Content-Type: application/json` · Respuesta: JSONL, o `204 No Content` con el cuerpo vacío

Aplica un JSON merge patch a los documentos que coinciden: los campos de `patch` se añaden o se sustituyen, y `null` elimina un campo. También se puede cambiar el `id`; con un índice `map` sobre `id`, un patch a un id que ya existe no se aplica.

| Campo del cuerpo | Por defecto | Significado |
|---|---|---|
| `filter` | `{}` | Qué documentos; misma sintaxis que `:find`. |
| `patch` | | Los cambios. |
| `limit` | `1` | Número máximo de documentos a cambiar. `-1` cambia todos los que coinciden. |

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:patch" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"filter": {"id": "alan"}, "patch": {"plan": "pro", "mrr": 120}}'
```

Respuesta `200 OK`, los documentos cambiados como JSON Lines:

```jsonl
{"id":"alan","name":"Alan Turing","plan":"pro","mrr":120}
```
