# 删除文档

删除通过全表扫描或索引遍历选中的文档。端点以 JSON Lines 返回每个已删除文档。

## 认证

需要 `Api-Key` 和 `Api-Secret`，或者在数据库 owner 被允许时使用 Glue owner token。

## 请求体

| 字段 | 类型 | 描述 |
|------|------|------|
| `index` | string | 可选索引名称。省略时扫描整个集合。 |
| `filter` | object | 删除前应用的可选 Connor filter。 |
| `skip` | integer | 跳过的匹配文档数量。默认：`0`。 |
| `limit` | integer | 最多删除的文档数量。默认：`1`；`0` 不删除文档。 |

请求体中也可以包含索引专用字段。

## HTTP 请求

```http
POST /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections/users:remove HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
Content-Type: application/json

{"filter":{"id":"user-2"},"limit":1}
```

## 响应

```jsonl
{"id":"user-2","name":"Bob","role":"member"}
```
