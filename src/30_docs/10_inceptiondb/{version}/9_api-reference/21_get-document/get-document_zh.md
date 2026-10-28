# 按 id 获取文档

`GET /v1/databases/{databaseId}/collections/{collection}/documents/{id}` · API key

```bash
curl "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers/documents/ada" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET"
```

响应 `200 OK`。文档位于 `document` 中；`source` 说明是 `id` 上的索引响应的（`index`），还是遍历了集合（`fullscan`）。

```json
{"id":"ada","document":{"id":"ada","mrr":120,"name":"Ada Lovelace","plan":"pro"},"source":{"type":"fullscan"}}
```

文档不存在时返回 `404 Not Found`。
