# 获取文档

按 `id` 获取一个文档。当存在字段为 `id` 的 `map` 索引时会使用该索引，否则回退到集合扫描。

## 认证

需要 `Api-Key` 和 `Api-Secret`，或者在数据库 owner 被允许时使用 Glue owner token。

## HTTP 请求

```http
GET /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections/users/documents/user-1 HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
```

## 响应

```json
{
  "id": "user-1",
  "document": {"id":"user-1","name":"Alice","role":"admin"},
  "source": {"type":"index","name":"id"}
}
```

端点使用集合扫描时会省略 `source`。空文档 ID 返回 `400 Bad Request`；集合或文档不存在返回 `404 Not Found`。
