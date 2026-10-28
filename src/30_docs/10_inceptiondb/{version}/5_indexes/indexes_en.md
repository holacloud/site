# Indexes

Without an index, a query reads every document of the collection. That is fast for thousands of documents. For millions, or to get documents in order, or to make sure a value never repeats, create an index on the field you search by.

The examples use the `customers` collection and the environment variables of [Getting started](/docs/inceptiondb/v2/getting-started/).

## Which index do I need?

| Type | Use it to | Unique values | Query it with |
|---|---|---|---|
| `map` | Find a document by an exact value: an email, a username, an external id. | Yes, always | `value` |
| `btree` | Ranges and order: dates, amounts, scores. Values can repeat. | Optional (`unique`) | `from`, `to`, `reverse` |
| `fts` | Find documents by the words of a text: names, titles, descriptions. | No | `match` |

Indexes are kept in memory and updated on every write. `map` indexes, and `btree` indexes with `unique`, are checked before each write, so a repeated value is rejected.

## Look up by email (map)

Create the index:

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:createIndex" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -d '{"name": "by_email", "type": "map", "field": "email"}'
```

```json
{"field":"email","name":"by_email","sparse":false,"type":"map"}
```

Find a customer by email:

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:find" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -d '{"index": "by_email", "value": "grace@example.com"}'
```

```jsonl
{"id":"grace","name":"Grace Hopper","email":"grace@example.com","plan":"team","country":"US","mrr":480,"tags":["vip"],"address":{"city":"New York","zip":"10001"}}
```

From now on, inserting another customer with an existing email fails with `409 Conflict`:

```json
{"error":{"description":"Unexpected error","message":"index add 'by_email': index conflict: field 'email' with value 'ada@example.com'"}}
```

## Ranges and order (btree)

Create an index on the monthly revenue:

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:createIndex" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -d '{"name": "by_mrr", "type": "btree", "fields": ["mrr"]}'
```

```json
{"fields":["mrr"],"name":"by_mrr","sparse":false,"type":"btree","unique":false}
```

The customers that pay between 100 and 500, from the lowest:

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:find" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -d '{"index": "by_mrr", "from": {"mrr": 100}, "to": {"mrr": 500}, "limit": 10}'
```

```jsonl
{"id":"ada","name":"Ada Lovelace","email":"ada@example.com","plan":"pro","country":"GB","mrr":120,"tags":["vip","beta"],"address":{"city":"London","zip":"N1"}}
{"id":"grace","name":"Grace Hopper","email":"grace@example.com","plan":"team","country":"US","mrr":480,"tags":["vip"],"address":{"city":"New York","zip":"10001"}}
```

The top customers, from the highest: `"reverse": true`.

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:find" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -d '{"index": "by_mrr", "reverse": true, "limit": 3}'
```

Good to know:

- `from` and `to` are optional; leave one out for an open range.
- `fields` takes several fields, for example `["country", "-mrr"]`: by country, and inside each country from the highest revenue (`-` means descending).
- You can add a `filter` as well: the index chooses the order and the range, the filter discards documents.
- `"unique": true` rejects repeated values, like a `map` index.

## Word search (fts)

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:createIndex" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -d '{"name": "search_name", "type": "fts", "field": "name"}'
```

Find the customers with `grace` in their name, in upper or lower case:

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:find" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -d '{"index": "search_name", "match": "grace", "limit": 10}'
```

```jsonl
{"id":"grace","name":"Grace Hopper","email":"grace@example.com","plan":"team","country":"US","mrr":480,"tags":["vip"],"address":{"city":"New York","zip":"10001"}}
```

## Documents without the field

Documents that do not have the indexed field are stored as usual, but they are not part of the index: a query that uses the index does not return them. A query without `index` still finds them. Create indexes with `"sparse": true` to make that explicit:

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:createIndex" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -d '{"name": "by_company", "type": "btree", "fields": ["company"], "sparse": true}'
```

## List and delete indexes

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:listIndexes" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET"
```

```json
[{"fields":["company"],"name":"by_company","sparse":true,"type":"btree","unique":false},{"field":"email","name":"by_email","sparse":false,"type":"map"},{"fields":["mrr"],"name":"by_mrr","sparse":false,"type":"btree","unique":false},{"field":"name","name":"search_name","type":"fts"}]
```

Deleting an index does not touch the documents (`204 No Content`):

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:dropIndex" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -d '{"name": "by_mrr"}'
```

In the [console](https://inceptiondb.hola.cloud/), the **Schema** tab creates an index on any field in one click.
