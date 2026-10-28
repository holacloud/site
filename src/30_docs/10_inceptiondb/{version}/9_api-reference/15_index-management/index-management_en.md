# Index management

Create, list, get and delete the indexes of a collection. All of them need an API key. [Indexes](/docs/inceptiondb/v2/indexes/) explains when to use each type.

## Create an index

`POST …/collections/{collection}:createIndex`

| Body field | Types | Meaning |
|---|---|---|
| `name` | all | Name of the index. |
| `type` | all | `map`, `btree` or `fts`. |
| `field` | `map`, `fts` | The indexed field. |
| `fields` | `btree` | Indexed fields, in order. `-field` for descending. |
| `unique` | `btree` | Reject repeated values. `map` indexes are always unique. |
| `sparse` | `map`, `btree` | Documents without the field are left out of the index. They are always stored. |

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:createIndex" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -d '{"name": "by_email", "type": "map", "field": "email"}'
```

Response `201 Created`:

```json
{"field":"email","name":"by_email","sparse":false,"type":"map"}
```

## List indexes

`POST …/collections/{collection}:listIndexes`

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:listIndexes" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET"
```

```json
[{"field":"email","name":"by_email","sparse":false,"type":"map"}]
```

## Get an index

`POST …/collections/{collection}:getIndex`

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:getIndex" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -d '{"name": "by_email"}'
```

```json
{"field":"email","name":"by_email","sparse":false,"type":"map"}
```

## Delete an index

`POST …/collections/{collection}:dropIndex`

The documents are not touched. Response `204 No Content`.

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:dropIndex" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -d '{"name": "by_email"}'
```
