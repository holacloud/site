# 列出集合

列出数据库中的集合。

## 认证

需要 `Api-Key` 和 `Api-Secret`；当数据库 owner 被允许时，也可以使用 Glue owner token。

## HTTP 请求

```http
GET /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
```

## 响应

```json
[
  {
    "name": "users",
    "total": 5400,
    "indexes": 2,
    "defaults": {
      "id": "uuid()"
    }
  }
]
```
