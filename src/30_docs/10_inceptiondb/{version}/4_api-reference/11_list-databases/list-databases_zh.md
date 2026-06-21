
# GET /v1/databases

列出已认证项目可访问的所有数据库。

## 身份验证

需要 `Api-Key`、`Api-Secret` 和 `X-Project` 请求头。

## HTTP 请求

```http
GET /v1/databases HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
X-Project: acme-webapp
```

## 示例

```bash
curl -X GET "https://api.hola.cloud/v1/databases" \
  -H "Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d" \
  -H "Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf" \
  -H "X-Project: acme-webapp"
```

## 响应

```json
[
  {
    "id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
    "name": "production-db",
    "created_at": "2025-06-15T10:30:00Z",
    "collection_count": 12
  },
  {
    "id": "b2c3d4e5-f6a7-8901-bcde-f12345678901",
    "name": "staging-db",
    "created_at": "2025-06-10T08:00:00Z",
    "collection_count": 5
  }
]
```

## 错误码

| 状态码 | 描述 |
|--------|------|
| 401 | 缺少或无效的身份验证标头 |
| 403 | 项目访问被拒绝 |
