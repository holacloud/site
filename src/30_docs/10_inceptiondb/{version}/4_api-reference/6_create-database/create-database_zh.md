
# Create Database

在已认证项目中创建新数据库。

## 身份验证

需要 `Api-Key`、`Api-Secret` 和 `X-Project` 请求头。

## 请求体

| 字段 | 类型 | 描述 |
|-------|------|------|
| name | string | 数据库的可读名称 |

## HTTP 请求

```http
POST /v1/databases HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
X-Project: acme-webapp
Content-Type: application/json

{
  "name": "my-database"
}
```

## 示例

```bash
curl -X POST "https://api.hola.cloud/v1/databases" \
  -H "Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d" \
  -H "Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf" \
  -H "X-Project: acme-webapp" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "my-database"
  }'
```

## 响应

```json
{
  "id": "c3d4e5f6-a7b8-9012-cdef-123456789012",
  "name": "my-database",
  "created_at": "2025-07-01T14:00:00Z",
  "collection_count": 0
}
```

## 错误码

| 状态码 | 描述 |
|--------|------|
| 400 | 缺少或无效的 `name` 字段 |
| 401 | 缺少或无效的身份验证标头 |
| 403 | 项目访问被拒绝 |
| 409 | 同名数据库已存在 |
