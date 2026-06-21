
# GET /v1/loggers

列出与您账户关联的所有日志记录器。

## 认证

需要管理凭据：

- `Api-Key` — 您的 API 密钥
- `Api-Secret` — 您的 API 密钥密码

## 请求

```bash
curl "https://api.hola.cloud/v1/loggers" \
  -H "Api-Key: your-api-key" \
  -H "Api-Secret: your-api-secret"
```

```http
GET /v1/loggers HTTP/1.1
Host: api.hola.cloud
Api-Key: your-api-key
Api-Secret: your-api-secret
```

## 响应

```json
[
  {
    "id": "logger_xyz789",
    "name": "my-app-logger",
    "created_at": "2025-06-20T14:22:00Z"
  },
  {
    "id": "logger_abc123",
    "name": "staging-logger",
    "created_at": "2025-06-19T10:00:00Z"
  }
]
```

## 错误代码

| 代码 | 描述 |
|------|------|
| 401 | 缺少或无效的 API 凭据 |
| 403 | API 凭据没有列出权限 |
