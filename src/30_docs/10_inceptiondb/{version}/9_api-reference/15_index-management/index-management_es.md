# Gestión de índices

Crear, listar, obtener y borrar los índices de una colección. Todas necesitan una API key. En [Índices](/es/docs/inceptiondb/v2/indices/) se explica cuándo usar cada tipo.

## Crear un índice

`POST …/collections/{collection}:createIndex`

| Campo del cuerpo | Tipos | Significado |
|---|---|---|
| `name` | todos | Nombre del índice. |
| `type` | todos | `map`, `btree` o `fts`. |
| `field` | `map`, `fts` | El campo indexado. |
| `fields` | `btree` | Campos indexados, en orden. `-campo` para orden descendente. |
| `unique` | `btree` | Rechazar valores repetidos. Los índices `map` son siempre únicos. |
| `sparse` | `map`, `btree` | Los documentos sin el campo quedan fuera del índice. Siempre se guardan. |

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:createIndex" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"name": "by_email", "type": "map", "field": "email"}'
```

Respuesta `201 Created`:

```json
{"field":"email","name":"by_email","sparse":false,"type":"map"}
```

## Listar índices

`POST …/collections/{collection}:listIndexes`

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:listIndexes" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET"
```

```json
[{"field":"email","name":"by_email","sparse":false,"type":"map"}]
```

## Obtener un índice

`POST …/collections/{collection}:getIndex`

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:getIndex" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"name": "by_email"}'
```

```json
{"field":"email","name":"by_email","sparse":false,"type":"map"}
```

## Borrar un índice

`POST …/collections/{collection}:dropIndex`

Los documentos no se tocan. Respuesta `204 No Content`.

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:dropIndex" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"name": "by_email"}'
```
