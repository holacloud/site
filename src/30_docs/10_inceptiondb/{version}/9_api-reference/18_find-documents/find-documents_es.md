# Buscar documentos

`POST /v1/databases/{databaseId}/collections/{collection}:find` · API key

Petición: JSON, `Content-Type: application/json` · Respuesta: JSONL, o `204 No Content` con el cuerpo vacío

| Campo del cuerpo | Por defecto | Significado |
|---|---|---|
| `filter` | `{}` | Condiciones; ver [Consultas](/es/docs/inceptiondb/v2/consultas/). |
| `limit` | `1` | Número máximo de documentos. `-1` los devuelve todos. |
| `skip` | `0` | Documentos que coinciden que se saltan primero. |
| `index` | ninguno | Nombre de un [índice](/es/docs/inceptiondb/v2/indices/) con el que responder. |
| `value` | | Valor a buscar, con un índice `map`. |
| `from`, `to` | | Límites del rango, con un índice `btree`. Ejemplo: `{"mrr": 100}`. |
| `reverse` | `false` | Orden descendente, con un índice `btree`. |
| `match` | | Palabras a buscar, con un índice `fts`. |

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:find" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"filter": {"plan": {"$in": ["pro", "team"]}}, "limit": 10}'
```

Respuesta `200 OK`, los documentos como JSON Lines. `204 No Content` cuando nada coincide.

```jsonl
{"id":"ada","name":"Ada Lovelace","plan":"pro","mrr":120}
```

Operadores: `$eq`, `$ne`, `$gt`, `$ge`, `$lt`, `$le`, `$in`, `$nin`, `$contains`, `$and`, `$or`.

> **`$gte` y `$lte` no existen:** usa `$ge` y `$le`. Un operador desconocido no se notifica como error: el filtro no coincide con nada y la respuesta es `204 No Content`.
