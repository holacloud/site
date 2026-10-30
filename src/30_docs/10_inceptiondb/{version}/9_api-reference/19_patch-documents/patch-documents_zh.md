# 更新文档

`POST /v1/databases/{databaseId}/collections/{collection}:patch` · API key

请求：JSON，`Content-Type: application/json` · 响应：JSONL，或响应体为空的 `204 No Content`

对匹配的文档应用 JSON merge patch：`patch` 中的字段会被添加或替换，`null` 会删除字段。`id` 也可以修改；如果 `id` 上有 `map` 索引，改成已存在的 id 的 patch 不会被应用。

| 请求体字段 | 默认值 | 含义 |
|---|---|---|
| `filter` | `{}` | 选择哪些文档；语法与 `:find` 相同。 |
| `patch` | | 要做的修改。 |
| `limit` | `1` | 最多修改的文档数。`-1` 修改所有匹配的文档。 |

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:patch" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"filter": {"id": "alan"}, "patch": {"plan": "pro", "mrr": 120}}'
```

响应 `200 OK`，以 JSON Lines 返回修改后的文档：

```jsonl
{"id":"alan","name":"Alan Turing","plan":"pro","mrr":120}
```
