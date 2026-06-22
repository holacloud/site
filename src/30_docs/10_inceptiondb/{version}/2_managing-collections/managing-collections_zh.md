# 管理集合

集合端点使用数据库访问凭据：`Api-Key` 和 `Api-Secret`；当数据库 owner 被允许时，也可以使用 Glue owner token。

## 创建集合

```bash
curl -X POST "https://api.hola.cloud/v1/databases/{databaseId}/collections" \
  -H "Api-Key: {api_key}" \
  -H "Api-Secret: {api_secret}" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "my-collection"
  }'
```

```json
{
  "name": "my-collection",
  "total": 0,
  "indexes": 0,
  "defaults": {
    "id": "uuid()"
  }
}
```

## 列出集合

```bash
curl "https://api.hola.cloud/v1/databases/{databaseId}/collections" \
  -H "Api-Key: {api_key}" \
  -H "Api-Secret: {api_secret}"
```

```json
[
  {
    "name": "my-collection",
    "total": 3,
    "indexes": 0,
    "defaults": {
      "id": "uuid()"
    }
  }
]
```

## 插入文档

```bash
curl -X POST "https://api.hola.cloud/v1/databases/{databaseId}/collections/my-collection:insert" \
  -H "Api-Key: {api_key}" \
  -H "Api-Secret: {api_secret}" \
  -H "Content-Type: application/jsonl" \
  --data-binary '{"id":"1","name":"Alfonso"}
{"id":"2","name":"Gerardo"}'
```

响应为插入文档的 JSONL。

```jsonl
{"id":"1","name":"Alfonso"}
{"id":"2","name":"Gerardo"}
```
