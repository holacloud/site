# 创建集合

在指定数据库中创建集合。

## 认证

需要 `Api-Key` 和 `Api-Secret`；当数据库 owner 被允许时，也可以使用 Glue owner token。

## HTTP 请求

```http
POST /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
Content-Type: application/json

{
  "name": "my-collection"
}
```

## 响应

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
