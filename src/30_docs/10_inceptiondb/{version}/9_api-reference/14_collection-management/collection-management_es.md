# Gestión de colecciones

Obtener, configurar, medir y borrar una colección. Todas necesitan una API key.

## Obtener una colección

`GET /v1/databases/{databaseId}/collections/{collection}`

```bash
curl "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET"
```

```json
{"defaults":{"id":"uuid()"},"indexes":0,"name":"customers","total":3}
```

## Valores por defecto

`POST …/collections/{collection}:setDefaults`

Los campos que envías se añaden a los valores por defecto actuales y sustituyen a los que tienen el mismo nombre; un campo enviado como `null` elimina ese valor. Además de valores fijos, `"uuid()"` genera un id aleatorio, `"auto()"` un número creciente y `"unixnano()"` el momento de la inserción en nanosegundos. Un valor por defecto se aplica cuando el documento no tiene el campo o lo tiene como `null`.

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:setDefaults" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"id": "uuid()", "plan": "free", "created": "unixnano()"}'
```

Respuesta `200 OK` con todos los valores por defecto en uso:

```json
{"created":"unixnano()","id":"uuid()","plan":"free"}
```

## Tamaño

`POST …/collections/{collection}:size`

Bytes de memoria que ocupan los documentos (`memory`) y cada índice (`index.<nombre>`), y tamaño en disco (`disk`).

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:size" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET"
```

```json
{"disk":630,"index.by_email":260,"memory":646}
```

## Borrar una colección

`POST …/collections/{collection}:dropCollection`

Borra la colección, sus documentos y sus índices. No se puede deshacer. Respuesta `204 No Content`.

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:dropCollection" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET"
```
