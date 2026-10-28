# Querying

`:find` returns the documents that match a filter, as JSON Lines (one document per line). The examples use the `customers` collection and the environment variables of [Getting started](/docs/inceptiondb/v2/getting-started/).

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:find" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -d '{"filter": {"plan": "pro"}, "limit": 10}'
```

```jsonl
{"id":"ada","name":"Ada Lovelace","email":"ada@example.com","plan":"pro","country":"GB","mrr":120,"tags":["vip","beta"],"address":{"city":"London","zip":"N1"}}
```

The body of the request:

| Field | Default | Meaning |
|---|---|---|
| `filter` | `{}` (everything) | Conditions the documents must meet. |
| `limit` | `1` | How many documents to return. `-1` returns all of them. |
| `skip` | `0` | How many matching documents to skip first, for pagination. |
| `index` | none | Answer the query with an [index](/docs/inceptiondb/v2/indexes/). |

When nothing matches, the answer is `204 No Content` with an empty body.

## Filters

A filter is a JSON object. Each field is a condition and all of them must be met.

**Equal to a value**: `{"plan": "pro", "country": "GB"}`.

**Compare** with an operator:

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:find" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -d '{"filter": {"mrr": {"$ge": 100, "$lt": 500}}, "limit": 10}'
```

| Operator | Matches when the field | Example |
|---|---|---|
| `$eq` | is equal to the value | `{"plan": {"$eq": "pro"}}` |
| `$ne` | is not equal | `{"plan": {"$ne": "free"}}` |
| `$gt` | is greater than | `{"mrr": {"$gt": 100}}` |
| `$ge` | is greater than or equal | `{"mrr": {"$ge": 100}}` |
| `$lt` | is less than | `{"mrr": {"$lt": 100}}` |
| `$le` | is less than or equal | `{"mrr": {"$le": 100}}` |
| `$in` | is one of the values | `{"plan": {"$in": ["pro", "team"]}}` |
| `$nin` | is none of the values | `{"country": {"$nin": ["US", "GB"]}}` |
| `$contains` | is a text that contains the value | `{"email": {"$contains": "@example.com"}}` |

Use these exact names: `$ge` and `$le`, not `$gte` and `$lte`. A filter with an operator that does not exist matches no document.

Dates written as text in ISO format (`"2026-10-28T09:30:00Z"`) compare in time order: `{"created_at": {"$gt": "2026-01-01"}}`.

## Combine conditions

Several fields must all match. To match **any** of several conditions, use `$or`:

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:find" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -d '{"filter": {"$or": [{"plan": "team"}, {"country": "US"}]}, "limit": 10}'
```

```jsonl
{"id":"grace","name":"Grace Hopper","email":"grace@example.com","plan":"team","country":"US","mrr":480,"tags":["vip"],"address":{"city":"New York","zip":"10001"}}
```

`$and` takes a list too, and is useful to apply two conditions to the same field.

## Nested fields

Reach a field inside an object with a dot:

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

A condition on an array field matches when **any element** meets it. The customers tagged `vip`:

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

## Pagination

Combine `skip` and `limit`. The second page of 25 documents:

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:find" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -d '{"filter": {}, "skip": 25, "limit": 25}'
```

Without an index, documents come in no particular order. To get them sorted, or to make queries on large collections fast, use an [index](/docs/inceptiondb/v2/indexes/).

## Try it in the console

The [console](https://inceptiondb.hola.cloud/) runs your filter while you type, suggests the field names of your data and shows example filters made from it. The **Schema** tab lists every field with its type and most common values.
