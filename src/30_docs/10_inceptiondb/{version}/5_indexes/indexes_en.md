# Indexes

Without an index, a query reads every document of the collection. That is fast for thousands of documents. For millions, or to get documents in order, or to make sure a value never repeats, create an index on the field you search by.

**Before you start:** this guide uses the `customers` collection as [Getting started](/docs/inceptiondb/v2/getting-started/) leaves it, with Ada and Grace, and its environment variables. It creates four indexes and deletes them at the end. Documents are never changed.

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
  -H "Content-Type: application/json" \
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
  -H "Content-Type: application/json" \
  -d '{"index": "by_email", "value": "grace@example.com"}'
```

```jsonl
{"id":"grace","name":"Grace Hopper","email":"grace@example.com","plan":"team","country":"US","mrr":480,"tags":["vip"],"address":{"city":"New York","zip":"10001"}}
```

From now on, inserting another customer with an existing email fails with `409 Conflict`:

```json
{"error":{"description":"Unexpected error","message":"index add 'by_email': index conflict: field 'email' with value 'ada@example.com'"}}
```

`description` is a generic text; the useful detail is always in `message`: the index, the field and the repeated value. The document is not stored.

To make the `id` of the documents unique too, create a `map` index on `id`. Reads by id then use it: their `source` is `index` instead of `fullscan`.

## Ranges and order (btree)

Create an index on the monthly revenue:

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

The customers that pay between 100 and 500, from the lowest:

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

The top customers, from the highest: `"reverse": true`.

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:find" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
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
  -H "Content-Type: application/json" \
  -d '{"name": "search_name", "type": "fts", "field": "name"}'
```

Find the customers with `grace` in their name, in upper or lower case:

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

## Documents without the field

A `map` or `btree` index can only be created if every document already in the collection has the field. Otherwise the answer is a `500` with a message like ``error indexing existing data: field `company` is indexed and mandatory``.

For a field that only some documents have, like `company`, create the index with `"sparse": true`. Documents without the field are stored as usual, but they are not part of the index: a query that uses the index does not return them, and a query without `index` still finds them.

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:createIndex" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"name": "by_company", "type": "btree", "fields": ["company"], "sparse": true}'
```

## List the indexes

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:listIndexes" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET"
```

```json
[{"fields":["company"],"name":"by_company","sparse":true,"type":"btree","unique":false},{"field":"email","name":"by_email","sparse":false,"type":"map"},{"fields":["mrr"],"name":"by_mrr","sparse":false,"type":"btree","unique":false},{"field":"name","name":"search_name","type":"fts"}]
```

## Clean up

Deleting an index does not touch the documents. The answer is `204 No Content`. Delete the four indexes of this guide:

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

Creating an index with a name that already exists answers `500` with the message `index 'by_email' already exists`. If you ran this guide before without cleaning up, run these four commands and start again.

In the [console](https://inceptiondb.hola.cloud/), the **Schema** tab creates an index on any field in one click.
