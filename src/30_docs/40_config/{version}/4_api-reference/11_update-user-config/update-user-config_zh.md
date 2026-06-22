# 更新用户配置

更新当前用户的配置。使用深度合并语义：仅修改请求体中提供的字段。此端点需要 API 密钥身份验证。

## 身份验证

需要 `Api-Key` 和 `Api-Secret` 标头。

## 请求体

```json
{
  "features": {
    "new-checkout": true
  }
}
```

可以提供配置的任何子集。更新将与现有配置合并。

## 请求

```bash
curl -X PATCH "https://api.hola.cloud/v1/config" \
  -H "Api-Key: 您的API密钥" \
  -H "Api-Secret: 您的API密钥密码" \
  -H "Content-Type: application/json" \
  -d '{
    "features": {
      "new-checkout": true
    }
  }'
```

## 响应

```http
HTTP/1.1 200 OK
Content-Type: application/json
```

```json
{
  "project": "my-app",
  "environment": "production",
  "services": {
    "database": {
      "host": "db.example.com",
      "port": 5432
    },
    "cache": {
      "host": "redis.example.com",
      "port": 6379
    }
  },
  "features": {
    "new-checkout": true
  }
}
```

## 错误代码

| 状态 | 代码 | 描述 |
|--------|------|-------------|
| 400 | Bad Request | 请求体中 JSON 无效 |
| 401 | Unauthorized | 缺少或无效的 API 凭证 |
| 500 | Internal Server Error | 发生意外错误 |
