# Consultas

`:find` devuelve los documentos que cumplen un filtro, como JSON Lines (un documento por línea). Los ejemplos usan la colección `customers` y las variables de entorno de [Primeros pasos](/es/docs/inceptiondb/v2/primeros-pasos/).

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:find" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -d '{"filter": {"plan": "pro"}, "limit": 10}'
```

```jsonl
{"id":"ada","name":"Ada Lovelace","email":"ada@example.com","plan":"pro","country":"GB","mrr":120,"tags":["vip","beta"],"address":{"city":"London","zip":"N1"}}
```

El cuerpo de la petición:

| Campo | Por defecto | Significado |
|---|---|---|
| `filter` | `{}` (todo) | Condiciones que deben cumplir los documentos. |
| `limit` | `1` | Cuántos documentos devolver. `-1` los devuelve todos. |
| `skip` | `0` | Cuántos documentos que coinciden saltar primero, para paginar. |
| `index` | ninguno | Responder la consulta con un [índice](/es/docs/inceptiondb/v2/indices/). |

Cuando nada coincide, la respuesta es `204 No Content` con el cuerpo vacío.

## Filtros

Un filtro es un objeto JSON. Cada campo es una condición y deben cumplirse todas.

**Igual a un valor**: `{"plan": "pro", "country": "GB"}`.

**Comparar** con un operador:

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:find" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -d '{"filter": {"mrr": {"$ge": 100, "$lt": 500}}, "limit": 10}'
```

| Operador | Coincide cuando el campo | Ejemplo |
|---|---|---|
| `$eq` | es igual al valor | `{"plan": {"$eq": "pro"}}` |
| `$ne` | no es igual | `{"plan": {"$ne": "free"}}` |
| `$gt` | es mayor que | `{"mrr": {"$gt": 100}}` |
| `$ge` | es mayor o igual que | `{"mrr": {"$ge": 100}}` |
| `$lt` | es menor que | `{"mrr": {"$lt": 100}}` |
| `$le` | es menor o igual que | `{"mrr": {"$le": 100}}` |
| `$in` | es uno de los valores | `{"plan": {"$in": ["pro", "team"]}}` |
| `$nin` | no es ninguno de los valores | `{"country": {"$nin": ["US", "GB"]}}` |
| `$contains` | es un texto que contiene el valor | `{"email": {"$contains": "@example.com"}}` |

Usa exactamente estos nombres: `$ge` y `$le`, no `$gte` ni `$lte`. Un filtro con un operador que no existe no coincide con ningún documento.

Las fechas escritas como texto en formato ISO (`"2026-10-28T09:30:00Z"`) se comparan en orden cronológico: `{"created_at": {"$gt": "2026-01-01"}}`.

## Combinar condiciones

Varios campos deben cumplirse todos. Para que coincida **cualquiera** de varias condiciones, usa `$or`:

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:find" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -d '{"filter": {"$or": [{"plan": "team"}, {"country": "US"}]}, "limit": 10}'
```

```jsonl
{"id":"grace","name":"Grace Hopper","email":"grace@example.com","plan":"team","country":"US","mrr":480,"tags":["vip"],"address":{"city":"New York","zip":"10001"}}
```

`$and` también recibe una lista, y sirve para aplicar dos condiciones al mismo campo.

## Campos anidados

Llega a un campo dentro de un objeto con un punto:

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:find" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -d '{"filter": {"address.city": "London"}, "limit": 10}'
```

```jsonl
{"id":"ada","name":"Ada Lovelace","email":"ada@example.com","plan":"pro","country":"GB","mrr":120,"tags":["vip","beta"],"address":{"city":"London","zip":"N1"}}
```

## Arrays

Una condición sobre un campo que es un array se cumple cuando **cualquiera de sus elementos** la cumple. Los clientes con la etiqueta `vip`:

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:find" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -d '{"filter": {"tags": "vip"}, "limit": 10}'
```

```jsonl
{"id":"ada","name":"Ada Lovelace","email":"ada@example.com","plan":"pro","country":"GB","mrr":120,"tags":["vip","beta"],"address":{"city":"London","zip":"N1"}}
{"id":"grace","name":"Grace Hopper","email":"grace@example.com","plan":"team","country":"US","mrr":480,"tags":["vip"],"address":{"city":"New York","zip":"10001"}}
```

## Paginación

Combina `skip` y `limit`. La segunda página de 25 documentos:

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:find" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -d '{"filter": {}, "skip": 25, "limit": 25}'
```

Sin índice, los documentos no llegan en ningún orden concreto. Para recibirlos ordenados, o para que las consultas sobre colecciones grandes sean rápidas, usa un [índice](/es/docs/inceptiondb/v2/indices/).

## Pruébalo en la consola

La [consola](https://inceptiondb.hola.cloud/) ejecuta tu filtro mientras escribes, sugiere los nombres de campo de tus datos y muestra filtros de ejemplo sacados de ellos. La pestaña **Schema** lista cada campo con su tipo y sus valores más comunes.
