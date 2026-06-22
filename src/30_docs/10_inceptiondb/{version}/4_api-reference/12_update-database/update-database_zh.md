# 更新数据库

更新数据库属性。

## 认证

需要 `X-Glue-Authentication`。

## HTTP 请求

```http
PATCH /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890 HTTP/1.1
Host: api.hola.cloud
X-Glue-Authentication: your-glue-token
Content-Type: application/json

{
  "name": "renamed-database"
}
```

## 响应

```json
{
  "id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
  "name": "renamed-database",
  "creation_date": "2025-06-15T10:30:00Z",
  "owners": ["owner-id"],
  "api_keys": ["api-key-id"],
  "owners_length": 1
}
```
