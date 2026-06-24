# 获取当前用户

返回已认证用户的信息。

## 认证

需要 `X-Glue-Authentication`。

## HTTP 请求

```http
GET /api/v0/me HTTP/1.1
Host: api.hola.cloud
X-Glue-Authentication: YOUR_TOKEN
```

## 示例

```bash
curl -X GET "https://api.hola.cloud/api/v0/me" \
  -H "X-Glue-Authentication: YOUR_TOKEN"
```

## 响应

```json
{
  "id": "user_123",
  "email": "user@example.com",
  "created_at": "2025-01-15T08:00:00Z"
}
```

## 错误码

| 代码 | 描述 |
|------|------|
| 401 | 认证缺失或无效 |
