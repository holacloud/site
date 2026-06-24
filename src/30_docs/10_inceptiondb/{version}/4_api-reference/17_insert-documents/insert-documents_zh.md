# 插入文档

向集合插入一个或多个 JSON 文档。如果集合不存在，会自动创建。

## 认证

需要 `Api-Key` 和 `Api-Secret`，或者在数据库 owner 被允许时使用 Glue owner token。

## HTTP 请求

```http
POST /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections/users:insert HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
Content-Type: application/jsonl

{"id":"user-1","name":"Alice","role":"admin"}
{"id":"user-2","name":"Bob","role":"member"}
```

## 响应

至少插入一个文档时状态为 `201 Created`。响应是包含已插入文档的 JSON Lines。

```jsonl
{"id":"user-1","name":"Alice","role":"admin"}
{"id":"user-2","name":"Bob","role":"member"}
```

如果请求体为空，响应为 `204 No Content`。唯一索引冲突会在写入响应体前返回 `409 Conflict`。
