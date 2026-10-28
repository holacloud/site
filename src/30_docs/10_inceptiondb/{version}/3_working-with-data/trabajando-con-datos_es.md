# Trabajando con datos

Insertar, leer, actualizar y borrar documentos. Los ejemplos usan la colección `customers` y las variables de entorno de [Primeros pasos](/es/docs/inceptiondb/v2/primeros-pasos/).

## Insertar documentos

Envía un objeto JSON por línea (JSON Lines). Una petición puede llevar un documento o millones.

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:insert" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  --data-binary '{"id":"linus","name":"Linus Torvalds","plan":"free","country":"FI","mrr":0}
{"id":"barbara","name":"Barbara Liskov","plan":"team","country":"US","mrr":300}'
```

La respuesta (`201 Created`) trae los documentos guardados, ya con los valores por defecto de la colección:

```jsonl
{"id":"linus","name":"Linus Torvalds","plan":"free","country":"FI","mrr":0}
{"id":"barbara","name":"Barbara Liskov","plan":"team","country":"US","mrr":300}
```

Conviene saber que:

- Sin `id`, el documento recibe uno de los [valores por defecto](/es/docs/inceptiondb/v2/manejando-colecciones/) de la colección (un UUID aleatorio, salvo que lo cambies).
- Los documentos se guardan en orden y la inserción **se detiene en el primer documento rechazado**. Los anteriores quedan guardados.
- Para cargar un fichero, usa `--data-binary @customers.jsonl`.
- Cada documento debe ser un objeto JSON sin nombres de campo repetidos. Si no, la respuesta es `400 Bad Request`.
- Si un [índice único](/es/docs/inceptiondb/v2/indices/) ya tiene el valor, la respuesta es `409 Conflict`.

## Leer un documento por id

```bash
curl "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers/documents/linus" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET"
```

```json
{"id":"linus","document":{"country":"FI","id":"linus","mrr":0,"name":"Linus Torvalds","plan":"free"},"source":{"type":"fullscan"}}
```

`source` indica cómo se encontró el documento: `index` cuando respondió un índice sobre `id`, `fullscan` cuando se recorrió la colección. Si el documento no existe, la respuesta es `404 Not Found`.

Para leer muchos documentos a la vez, usa una [consulta](/es/docs/inceptiondb/v2/consultas/).

## Actualizar documentos

Un patch cambia solo los campos que envías (un [JSON merge patch](https://www.rfc-editor.org/rfc/rfc7386)): los campos nuevos se añaden, los que existen se sustituyen y `null` elimina un campo.

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:patch" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -d '{"filter": {"id": "linus"}, "patch": {"plan": "pro", "mrr": 49, "country": null}}'
```

La respuesta trae los documentos tal como quedan después del cambio:

```jsonl
{"id":"linus","name":"Linus Torvalds","plan":"pro","mrr":49}
```

`filter` elige los documentos, con la misma sintaxis que las [consultas](/es/docs/inceptiondb/v2/consultas/). Igual que las consultas, un patch cambia **un solo documento salvo que indiques `limit`**. Marca como `priority` a todos los clientes del plan `team`:

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:patch" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -d '{"filter": {"plan": "team"}, "patch": {"priority": true}, "limit": -1}'
```

`"limit": -1` significa todos los documentos que coinciden. El `id` de un documento no se puede cambiar.

## Borrar documentos

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:remove" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -d '{"filter": {"id": "linus"}}'
```

La respuesta trae los documentos borrados:

```jsonl
{"id":"linus","name":"Linus Torvalds","plan":"pro","mrr":49}
```

Borra un solo documento salvo que indiques `limit`. Borra todos los clientes del plan `free`:

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:remove" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -d '{"filter": {"plan": "free"}, "limit": -1}'
```

Los documentos borrados no se pueden recuperar. Ejecuta antes el mismo filtro con `:find` para ver qué documentos coinciden.

## Exportar una colección

`:find` con `"limit": -1` devuelve todos los documentos como JSON Lines, listos para guardar en un fichero o insertar en otro sitio:

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:find" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -d '{"limit": -1}' \
  -o customers.jsonl
```
