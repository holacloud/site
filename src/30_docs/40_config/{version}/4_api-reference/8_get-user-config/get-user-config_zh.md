# 获取用户配置

获取当前用户的配置。此端点需要 API 密钥身份验证。

## 身份验证

需要 `Api-Key` 和 `Api-Secret` 标头。

## 请求

```bash
curl "https://api.hola.cloud/v1/config" \
  -H "Api-Key: 您的API密钥" \
  -H "Api-Secret: 您的API密钥密码"
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
    "new-checkout": false
  }
}
```

## 错误代码

| 状态 | 代码 | 描述 |
|--------|------|-------------|
| 401 | Unauthorized | 缺少或无效的 API 凭证 |
| 500 | Internal Server Error | 发生意外错误 |
