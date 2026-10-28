# Find documents

`POST /v1/databases/{databaseId}/collections/{collection}:find` · API key

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
  -d '{"filter": {"plan": {"$in": ["pro", "team"]}}, "limit": 10}'
```

Response `200 OK`, the documents as JSON Lines. `204 No Content` when nothing matches.

```jsonl
{"id":"ada","name":"Ada Lovelace","plan":"pro","mrr":120}
```

Operators: `$eq`, `$ne`, `$gt`, `$ge`, `$lt`, `$le`, `$in`, `$nin`, `$contains`, `$and`, `$or`. A filter with an unknown operator matches no document.
