# Crear una colección

`POST /v1/databases/{databaseId}/collections` · API key

Crea una colección vacía. Sus valores por defecto dan un id aleatorio a cada documento que no lo tenga.

| Campo del cuerpo | Tipo | Obligatorio | Significado |
|---|---|---|---|
| `name` | string | sí | Nombre de la colección, usado en sus URLs. |

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -d '{"name": "customers"}'
```

Respuesta `201 Created`:

```json
{"defaults":{"id":"uuid()"},"indexes":0,"name":"customers","total":0}
```

Si ya existe una colección con ese nombre, la respuesta es `500` y el error indica que ya existe.
