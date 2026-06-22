# InceptionDB 快速开始

本指南创建一个集合，插入三个 JSON 文档，并查询过滤后的文档。

数据库访问端点使用 `Api-Key` 和 `Api-Secret`。当数据库 owner 被允许时，也可以使用 Glue owner token。

## 步骤 1：创建集合

```bash
curl -X POST "https://api.hola.cloud/v1/databases/your-database-id/collections" \
  -H "Api-Key: your-api-key" \
  -H "Api-Secret: your-api-secret" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "my-collection"
  }'
```

预期响应：

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

## 步骤 2：插入文档

使用集合 action 端点并发送 JSONL：每行一个 JSON 文档。

```bash
curl -X POST "https://api.hola.cloud/v1/databases/your-database-id/collections/my-collection:insert" \
  -H "Api-Key: your-api-key" \
  -H "Api-Secret: your-api-secret" \
  -H "Content-Type: application/jsonl" \
  --data-binary '{"name":"Element 1","value":100}
{"name":"Element 2","value":200}
{"name":"Element 3","value":300}'
```

预期响应：

```jsonl
{"id":"document-id-1","name":"Element 1","value":100}
{"id":"document-id-2","name":"Element 2","value":200}
{"id":"document-id-3","name":"Element 3","value":300}
```

## 步骤 3：查询过滤文档

```bash
curl -X POST "https://api.hola.cloud/v1/databases/your-database-id/collections/my-collection:find" \
  -H "Api-Key: your-api-key" \
  -H "Api-Secret: your-api-secret" \
  -H "Content-Type: application/json" \
  -d '{
    "filter": {
      "value": { "$gte": 200 }
    }
  }'
```

预期响应：

```jsonl
{"id":"document-id-2","name":"Element 2","value":200}
{"id":"document-id-3","name":"Element 3","value":300}
```
