# 删除数据库

根据 ID 删除数据库。

## 认证

需要 `X-Glue-Authentication`。

## HTTP 请求

```http
DELETE /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890 HTTP/1.1
Host: api.hola.cloud
X-Glue-Authentication: your-glue-token
```

## 响应

```json
{
  "id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
  "name": "production-db",
  "creation_date": "2025-06-15T10:30:00Z"
}
```
