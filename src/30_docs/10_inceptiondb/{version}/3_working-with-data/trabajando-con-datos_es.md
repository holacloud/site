# Trabajando con datos

Insertar, leer, actualizar y borrar documentos.

**Antes de empezar:** esta guía usa la colección `customers` tal como la deja [Primeros pasos](/es/docs/inceptiondb/v2/primeros-pasos/), con Ada y Grace, y sus variables de entorno. Añade dos clientes temporales marcados con `"trial": true` y los borra al final, así que `customers` queda como estaba.

## Insertar documentos

Las inserciones usan **JSON Lines** (JSONL): un objeto JSON por línea. Una petición puede llevar un documento o millones.

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:insert" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/x-ndjson" \
  --data-binary '{"id":"linus","name":"Linus Torvalds","plan":"free","country":"FI","mrr":0,"trial":true}
{"id":"barbara","name":"Barbara Liskov","plan":"team","country":"US","mrr":300,"trial":true}'
```

La respuesta (`201 Created`) contiene los documentos guardados, con los valores por defecto de la colección ya aplicados:

```jsonl
{"id":"linus","name":"Linus Torvalds","plan":"free","country":"FI","mrr":0,"trial":true}
{"id":"barbara","name":"Barbara Liskov","plan":"team","country":"US","mrr":300,"trial":true}
```

Un array JSON **no** es JSONL. Este cuerpo se rechaza con `400 Bad Request` y el mensaje `document must be a JSON object`:

```text
[{"id":"linus"},{"id":"barbara"}]
```

Los mismos dos documentos en JSONL:

```text
{"id":"linus"}
{"id":"barbara"}
```

Conviene saber:

- Sin `id`, el documento recibe uno de los [valores por defecto](/es/docs/inceptiondb/v2/manejando-colecciones/) de la colección (un UUID aleatorio salvo que lo cambies).
- InceptionDB no comprueba que los ids sean únicos: insertar dos veces el mismo `id` guarda dos documentos. Para evitarlo, crea un [índice `map`](/es/docs/inceptiondb/v2/indices/) sobre `id`. Entonces un id repetido responde `409 Conflict`, y las lecturas por id son instantáneas.
- Los documentos se guardan en orden y la inserción **se detiene en el primer documento rechazado**. Los anteriores quedan guardados.
- Para cargar un archivo, usa `--data-binary @customers.jsonl`. `-d` quitaría los saltos de línea.
- Cada documento debe ser un objeto JSON sin nombres de campo repetidos. Si no, la respuesta es `400 Bad Request`.
- Si un [índice único](/es/docs/inceptiondb/v2/indices/) ya tiene el valor, la respuesta es `409 Conflict`.
- Insertar en una colección que no existe la crea. Comprueba el nombre de la colección en la URL.

## Leer un documento por id

```bash
curl "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers/documents/linus" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET"
```

```json
{"id":"linus","document":{"country":"FI","id":"linus","mrr":0,"name":"Linus Torvalds","plan":"free","trial":true},"source":{"type":"fullscan"}}
```

`source` indica cómo se encontró el documento: `index` cuando respondió un índice sobre `id`, `fullscan` cuando se recorrió la colección. Un documento que no existe responde `404 Not Found`.

Para leer muchos documentos a la vez, usa una [consulta](/es/docs/inceptiondb/v2/consultas/).

## Actualizar documentos

Un patch cambia solo los campos que envías (un [JSON merge patch](https://www.rfc-editor.org/rfc/rfc7386)): los campos nuevos se añaden, los existentes se sustituyen y `null` elimina un campo.

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:patch" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"filter": {"id": "linus"}, "patch": {"plan": "pro", "mrr": 49, "country": null}}'
```

La respuesta contiene los documentos tal como quedan después del cambio:

```jsonl
{"id":"linus","name":"Linus Torvalds","plan":"pro","mrr":49,"trial":true}
```

`filter` elige los documentos, con la misma sintaxis que las [consultas](/es/docs/inceptiondb/v2/consultas/). Igual que las consultas, un patch cambia **un solo documento salvo que indiques `limit`**. Marca a todos los clientes de prueba con `priority`:

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:patch" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"filter": {"trial": true}, "patch": {"priority": true}, "limit": -1}'
```

```jsonl
{"id":"linus","name":"Linus Torvalds","plan":"pro","mrr":49,"trial":true,"priority":true}
{"id":"barbara","name":"Barbara Liskov","plan":"team","country":"US","mrr":300,"trial":true,"priority":true}
```

`"limit": -1` significa todos los documentos que coinciden. Si nada coincide, la respuesta es `204 No Content`.

Un patch también puede cambiar el `id` de un documento. Si la colección tiene un índice `map` sobre `id`, un patch a un id que ya existe no se aplica.

## Borrar documentos

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:remove" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"filter": {"id": "linus"}}'
```

La respuesta contiene los documentos borrados:

```jsonl
{"id":"linus","name":"Linus Torvalds","plan":"pro","mrr":49,"trial":true,"priority":true}
```

Borra un solo documento salvo que indiques `limit`. Borra todos los clientes de prueba, lo que además limpia esta guía:

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:remove" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"filter": {"trial": true}, "limit": -1}'
```

```jsonl
{"id":"barbara","name":"Barbara Liskov","plan":"team","country":"US","mrr":300,"trial":true,"priority":true}
```

Los documentos borrados no se pueden recuperar. Ejecuta antes el mismo filtro con `:find` para ver qué coincide.

## Exportar una colección

`:find` con `"limit": -1` devuelve todos los documentos en JSONL, listos para guardarlos en un archivo o insertarlos en otro sitio:

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:find" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"limit": -1}' \
  -o customers.jsonl
```

`customers.jsonl` contiene a Ada y Grace, una por línea: la colección ha vuelto a como estaba antes de esta guía.
