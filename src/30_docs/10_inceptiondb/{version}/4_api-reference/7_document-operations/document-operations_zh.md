# 文档操作

文档操作是集合上的 action 端点。

## 认证

需要 `Api-Key` 和 `Api-Secret`；当数据库 owner 被允许时，也可以使用 Glue owner token。

## 插入：`{collection}:insert`

```http
POST /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections/users:insert HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
Content-Type: application/jsonl

{"id":"user-1","name":"Alice","role":"admin"}
{"id":"user-2","name":"Bob","role":"member"}
```

```jsonl
{"id":"user-1","name":"Alice","role":"admin"}
{"id":"user-2","name":"Bob","role":"member"}
```

## 查询：`{collection}:find`

```http
POST /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections/users:find HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
Content-Type: application/json

{
  "filter": {
    "role": "admin"
  }
}
```

```jsonl
{"id":"user-1","name":"Alice","role":"admin"}
```

## 根据 ID 获取

```http
GET /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections/users/documents/user-1 HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
```

```json
{"id":"user-1","name":"Alice","role":"admin"}
```

## 修改：`{collection}:patch`

```http
POST /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections/users:patch HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
Content-Type: application/json

{
  "filter": { "id": "user-1" },
  "patch": { "role": "owner" }
}
```

```jsonl
{"id":"user-1","name":"Alice","role":"owner"}
```

## 删除：`{collection}:remove`

```http
POST /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections/users:remove HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
Content-Type: application/json

{
  "filter": { "id": "user-2" }
}
```

```jsonl
{"id":"user-2","name":"Bob","role":"member"}
```
