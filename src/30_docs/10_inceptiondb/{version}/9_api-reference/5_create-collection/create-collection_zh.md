# 创建集合

`POST /v1/databases/{databaseId}/collections` · API key

创建一个空集合。其默认值会为每个没有 `id` 的文档生成随机 id。

| 请求体字段 | 类型 | 必填 | 含义 |
|---|---|---|---|
| `name` | string | 是 | 集合名称，用于其 URL。 |

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -d '{"name": "customers"}'
```

响应 `201 Created`：

```json
{"defaults":{"id":"uuid()"},"indexes":0,"name":"customers","total":0}
```

如果同名集合已存在，返回 `500`，错误信息说明它已存在。
