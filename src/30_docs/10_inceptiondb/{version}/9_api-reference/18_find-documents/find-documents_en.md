# Find documents

`POST /v1/databases/{databaseId}/collections/{collection}:find` · API key

Request: JSON, `Content-Type: application/json` · Response: JSONL, or `204 No Content` with an empty body

| Body field | Default | Meaning |
|---|---|---|
| `filter` | `{}` | Conditions; see [Querying](/docs/inceptiondb/v2/querying/). |
| `limit` | `1` | Maximum number of documents. `-1` returns all. |
| `skip` | `0` | Matching documents to skip first. |
| `index` | none | Name of an [index](/docs/inceptiondb/v2/indexes/) to answer with. |
| `value` | | Value to look up, with a `map` index. |
| `from`, `to` | | Range limits, with a `btree` index. Example: `{"mrr": 100}`. |
| `reverse` | `false` | Descending order, with a `btree` index. |
| `match` | | Words to search, with an `fts` index. |

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:find" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"filter": {"plan": {"$in": ["pro", "team"]}}, "limit": 10}'
```

Response `200 OK`, the documents as JSON Lines. `204 No Content` when nothing matches.

```jsonl
{"id":"ada","name":"Ada Lovelace","plan":"pro","mrr":120}
```

Operators: `$eq`, `$ne`, `$gt`, `$ge`, `$lt`, `$le`, `$in`, `$nin`, `$contains`, `$and`, `$or`.

> **`$gte` and `$lte` do not exist:** use `$ge` and `$le`. An unknown operator is not reported as an error: the filter matches nothing and the answer is `204 No Content`.
