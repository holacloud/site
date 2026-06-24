# 修改文档

修改通过全表扫描或索引遍历选中的文档。端点以 JSON Lines 返回每个已修改文档。

## 认证

需要 `Api-Key` 和 `Api-Secret`，或者在数据库 owner 被允许时使用 Glue owner token。

## 请求体

| 字段 | 类型 | 描述 |
|------|------|------|
| `patch` | object | 应用于每个匹配文档的 patch 对象。 |
| `index` | string | 可选索引名称。省略时扫描整个集合。 |
| `filter` | object | patch 前应用的可选 Connor filter。 |
| `skip` | integer | 跳过的匹配文档数量。默认：`0`。 |
| `limit` | integer | 最多修改的文档数量。默认：`1`；`0` 不修改文档。 |

请求体中也可以包含索引专用字段。

## HTTP 请求

```http
POST /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections/users:patch HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
Content-Type: application/json

{
  "filter": {"id": "user-1"},
  "patch": {"role": "owner"},
  "limit": 1
}
```

## 响应

```jsonl
{"id":"user-1","name":"Alice","role":"owner"}
```
