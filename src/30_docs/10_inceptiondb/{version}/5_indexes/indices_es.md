# Índices

Sin índice, una consulta lee todos los documentos de la colección. Eso es rápido con miles de documentos. Con millones, para recibir los documentos ordenados o para asegurarte de que un valor no se repite nunca, crea un índice sobre el campo por el que buscas.

**Antes de empezar:** esta guía usa la colección `customers` tal como la deja [Primeros pasos](/es/docs/inceptiondb/v2/primeros-pasos/), con Ada y Grace, y sus variables de entorno. Crea cuatro índices y los borra al final. Los documentos no se modifican nunca.

## ¿Qué índice necesito?

| Tipo | Úsalo para | Valores únicos | Se consulta con |
|---|---|---|---|
| `map` | Encontrar un documento por un valor exacto: un email, un usuario, un id externo. | Sí, siempre | `value` |
| `btree` | Rangos y orden: fechas, importes, puntuaciones. Los valores se pueden repetir. | Opcional (`unique`) | `from`, `to`, `reverse` |
| `fts` | Encontrar documentos por las palabras de un texto: nombres, títulos, descripciones. | No | `match` |

Los índices se guardan en memoria y se actualizan en cada escritura. Los índices `map`, y los `btree` con `unique`, se comprueban antes de cada escritura, así que un valor repetido se rechaza.

## Buscar por email (map)

Crea el índice:

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:createIndex" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"name": "by_email", "type": "map", "field": "email"}'
```

```json
{"field":"email","name":"by_email","sparse":false,"type":"map"}
```

Busca un cliente por su email:

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:find" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"index": "by_email", "value": "grace@example.com"}'
```

```jsonl
{"id":"grace","name":"Grace Hopper","email":"grace@example.com","plan":"team","country":"US","mrr":480,"tags":["vip"],"address":{"city":"New York","zip":"10001"}}
```

A partir de ahora, insertar otro cliente con un email que ya existe falla con `409 Conflict`:

```json
{"error":{"description":"Unexpected error","message":"index add 'by_email': index conflict: field 'email' with value 'ada@example.com'"}}
```

`description` es un texto genérico; el detalle útil está siempre en `message`: el índice, el campo y el valor repetido. El documento no se guarda.

Para que el `id` de los documentos también sea único, crea un índice `map` sobre `id`. Las lecturas por id lo usan entonces: su `source` es `index` en lugar de `fullscan`.

## Rangos y orden (btree)

Crea un índice sobre los ingresos mensuales:

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:createIndex" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"name": "by_mrr", "type": "btree", "fields": ["mrr"]}'
```

```json
{"fields":["mrr"],"name":"by_mrr","sparse":false,"type":"btree","unique":false}
```

Los clientes que pagan entre 100 y 500, de menor a mayor:

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:find" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"index": "by_mrr", "from": {"mrr": 100}, "to": {"mrr": 500}, "limit": 10}'
```

```jsonl
{"id":"ada","name":"Ada Lovelace","email":"ada@example.com","plan":"pro","country":"GB","mrr":120,"tags":["vip","beta"],"address":{"city":"London","zip":"N1"}}
{"id":"grace","name":"Grace Hopper","email":"grace@example.com","plan":"team","country":"US","mrr":480,"tags":["vip"],"address":{"city":"New York","zip":"10001"}}
```

Los mejores clientes, de mayor a menor: `"reverse": true`.

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:find" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"index": "by_mrr", "reverse": true, "limit": 3}'
```

Conviene saber que:

- `from` y `to` son opcionales; quita uno para un rango abierto.
- `fields` admite varios campos, por ejemplo `["country", "-mrr"]`: por país y, dentro de cada país, de mayor a menor ingreso (`-` significa descendente).
- También puedes añadir un `filter`: el índice decide el orden y el rango, y el filtro descarta documentos.
- `"unique": true` rechaza valores repetidos, como un índice `map`.

## Búsqueda por palabras (fts)

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:createIndex" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"name": "search_name", "type": "fts", "field": "name"}'
```

Encuentra los clientes con `grace` en su nombre, en mayúsculas o minúsculas:

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:find" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"index": "search_name", "match": "grace", "limit": 10}'
```

```jsonl
{"id":"grace","name":"Grace Hopper","email":"grace@example.com","plan":"team","country":"US","mrr":480,"tags":["vip"],"address":{"city":"New York","zip":"10001"}}
```

## Documentos sin el campo

Un índice `map` o `btree` solo se puede crear si todos los documentos que ya hay en la colección tienen el campo. Si no, la respuesta es un `500` con un mensaje como ``error indexing existing data: field `company` is indexed and mandatory``.

Para un campo que solo tienen algunos documentos, como `company`, crea el índice con `"sparse": true`. Los documentos sin el campo se guardan con normalidad, pero no forman parte del índice: una consulta que usa el índice no los devuelve, y una consulta sin `index` sí los encuentra.

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:createIndex" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"name": "by_company", "type": "btree", "fields": ["company"], "sparse": true}'
```

## Listar los índices

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:listIndexes" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET"
```

```json
[{"fields":["company"],"name":"by_company","sparse":true,"type":"btree","unique":false},{"field":"email","name":"by_email","sparse":false,"type":"map"},{"fields":["mrr"],"name":"by_mrr","sparse":false,"type":"btree","unique":false},{"field":"name","name":"search_name","type":"fts"}]
```

## Limpieza

Borrar un índice no toca los documentos. La respuesta es `204 No Content`. Borra los cuatro índices de esta guía:

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:dropIndex" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"name": "by_mrr"}'
```

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:dropIndex" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"name": "by_email"}'
```

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:dropIndex" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"name": "search_name"}'
```

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:dropIndex" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"name": "by_company"}'
```

Crear un índice con un nombre que ya existe responde `500` con el mensaje `index 'by_email' already exists`. Si ya seguiste esta guía sin limpiar, ejecuta estos cuatro comandos y empieza de nuevo.

En la [consola](https://inceptiondb.hola.cloud/), la pestaña **Schema** crea un índice sobre cualquier campo con un clic.
