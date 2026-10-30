# 删除文档

`POST /v1/databases/{databaseId}/collections/{collection}:remove` · API key

请求：JSON，`Content-Type: application/json` · 响应：JSONL，或响应体为空的 `204 No Content`

| 请求体字段 | 默认值 | 含义 |
|---|---|---|
| `filter` | `{}` | 选择哪些文档；语法与 `:find` 相同。 |
| `limit` | `1` | 最多删除的文档数。`-1` 删除所有匹配的文档。 |

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:remove" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"filter": {"id": "alan"}}'
```

响应 `200 OK`，以 JSON Lines 返回被删除的文档：

```jsonl
{"id":"alan","name":"Alan Turing","plan":"free","mrr":0}
```
