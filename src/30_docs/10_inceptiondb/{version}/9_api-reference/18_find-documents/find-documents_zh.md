# 查找文档

`POST /v1/databases/{databaseId}/collections/{collection}:find` · API key

请求：JSON，`Content-Type: application/json` · 响应：JSONL，或响应体为空的 `204 No Content`

| 请求体字段 | 默认值 | 含义 |
|---|---|---|
| `filter` | `{}` | 条件；见 [查询](/zh/docs/inceptiondb/v2/querying/)。 |
| `limit` | `1` | 最多返回的文档数。`-1` 返回全部。 |
| `skip` | `0` | 先跳过的匹配文档数。 |
| `index` | 无 | 用于响应查询的 [索引](/zh/docs/inceptiondb/v2/indexes/) 名称。 |
| `value` | | 要查找的值，配合 `map` 索引。 |
| `from`、`to` | | 范围边界，配合 `btree` 索引。例如：`{"mrr": 100}`。 |
| `reverse` | `false` | 降序，配合 `btree` 索引。 |
| `match` | | 要搜索的词，配合 `fts` 索引。 |

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:find" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"filter": {"plan": {"$in": ["pro", "team"]}}, "limit": 10}'
```

响应 `200 OK`，以 JSON Lines 返回文档。没有匹配时返回 `204 No Content`。

```jsonl
{"id":"ada","name":"Ada Lovelace","plan":"pro","mrr":120}
```

运算符：`$eq`、`$ne`、`$gt`、`$ge`、`$lt`、`$le`、`$in`、`$nin`、`$contains`、`$and`、`$or`。

> **`$gte` 和 `$lte` 不存在：** 请使用 `$ge` 和 `$le`。未知运算符不会报错：过滤器不匹配任何文档，响应为 `204 No Content`。
