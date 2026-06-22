# 处理数据

文档操作是集合 action。数据库访问使用 `Api-Key` 和 `Api-Secret`，当数据库 owner 被允许时也可以使用 Glue owner token。

## 插入文档

```http
POST /v1/databases/{databaseId}/collections/{collection}:insert
Content-Type: application/jsonl
Api-Key: your-api-key
Api-Secret: your-api-secret

{"id":"1","name":"Alice"}
{"id":"2","name":"Bob"}
```

响应为插入后的 JSONL 文档。

```jsonl
{"id":"1","name":"Alice"}
{"id":"2","name":"Bob"}
```

## 查询文档

```http
POST /v1/databases/{databaseId}/collections/{collection}:find
Content-Type: application/json
Api-Key: your-api-key
Api-Secret: your-api-secret

{
  "filter": {
    "name": "Alice"
  }
}
```

响应为匹配的 JSONL 文档。

```jsonl
{"id":"1","name":"Alice"}
```

## 根据 ID 获取文档

```http
GET /v1/databases/{databaseId}/collections/{collection}/documents/{documentId}
Api-Key: your-api-key
Api-Secret: your-api-secret
```

响应为文档 JSON。

```json
{
  "id": "1",
  "name": "Alice"
}
```

## 修改文档

```http
POST /v1/databases/{databaseId}/collections/{collection}:patch
Content-Type: application/json
Api-Key: your-api-key
Api-Secret: your-api-secret

{
  "filter": {
    "id": "1"
  },
  "patch": {
    "name": "Alice Updated"
  }
}
```

响应为修改后的 JSONL 文档。

```jsonl
{"id":"1","name":"Alice Updated"}
```

## 删除文档

```http
POST /v1/databases/{databaseId}/collections/{collection}:remove
Content-Type: application/json
Api-Key: your-api-key
Api-Secret: your-api-secret

{
  "filter": {
    "id": "1"
  }
}
```

响应为删除的 JSONL 文档。

```jsonl
{"id":"1","name":"Alice Updated"}
```
