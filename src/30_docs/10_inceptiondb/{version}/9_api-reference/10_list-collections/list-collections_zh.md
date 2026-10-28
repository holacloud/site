# 列出集合

`GET /v1/databases/{databaseId}/collections` · API key

返回数据库中的所有集合，以及每个集合的文档数、索引数和默认值。

```bash
curl "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET"
```

响应 `200 OK`：

```json
[{"defaults":{"id":"uuid()"},"indexes":0,"name":"customers","total":3}]
```
