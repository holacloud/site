
# Get Database

根据数据库 ID 获取特定数据库的详细信息。

## 身份验证

需要 `Api-Key`、`Api-Secret` 和 `X-Project` 请求头。

## 路径参数

| 参数 | 类型 | 描述 |
|-----------|------|------|
| id | uuid | 数据库的唯一标识符 |

## HTTP 请求

```http
GET /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890 HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
X-Project: acme-webapp
```

## 示例

```bash
curl -X GET "https://api.hola.cloud/v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890" \
  -H "Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d" \
  -H "Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf" \
  -H "X-Project: acme-webapp"
```

## 响应

```json
{
  "id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
  "name": "production-db",
  "created_at": "2025-06-15T10:30:00Z",
  "collection_count": 12,
  "collections": [
    {
      "name": "users",
      "document_count": 5400
    },
    {
      "name": "orders",
      "document_count": 12300
    }
  ]
}
```

## 错误码

| 状态码 | 描述 |
|--------|------|
| 401 | 缺少或无效的身份验证标头 |
| 403 | 项目访问被拒绝 |
| 404 | 未找到数据库 |
