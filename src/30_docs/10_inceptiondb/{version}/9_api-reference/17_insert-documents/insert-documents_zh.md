# 插入文档

`POST /v1/databases/{databaseId}/collections/{collection}:insert` · API key

请求：JSON Lines（JSONL），`Content-Type: application/x-ndjson` · 响应：JSONL

请求体为 JSON Lines：每行一个 JSON 对象。缺少的字段使用集合的默认值。

```bash
curl -X POST "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections/customers:insert" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET" \
  -H "Content-Type: application/x-ndjson" \
  --data-binary '{"id":"ada","name":"Ada Lovelace","plan":"pro","mrr":120}
{"id":"alan","name":"Alan Turing","plan":"free","mrr":0}'
```

响应 `201 Created`，以 JSON Lines 返回已保存的文档：

```jsonl
{"id":"ada","name":"Ada Lovelace","plan":"pro","mrr":120}
{"id":"alan","name":"Alan Turing","plan":"free","mrr":0}
```

文档按顺序保存。插入在第一个被拒绝的文档处停止；之前的文档仍会保存。

- JSON 数组（`[{...},{...}]`）不是 JSON Lines：会以 `400` 拒绝。
- 如果集合不存在，插入会创建它。
- 除非集合在 `id` 上有 `map` 索引，否则 `id` 不保证唯一。

| 状态码 | 何时 |
|---|---|
| 400 | 某行不是 JSON 对象（例如请求体是 JSON 数组），或包含重复的字段名。 |
| 409 | 唯一索引中已存在该值。 |
