
# Create Logger

创建新的日志记录器。

## 认证

需要管理凭据：

- `Api-Key` — 您的 API 密钥
- `Api-Secret` — 您的 API 密钥密码

## 请求体

| 字段 | 类型 | 必填 | 描述 |
|------|------|------|-------------|
| `name` | string | 是 | 日志记录器的人类可读名称 |

```json
{
  "name": "my-app-logger"
}
```

## 请求

```bash
curl -X POST "https://api.hola.cloud/v1/loggers" \
  -H "Api-Key: your-api-key" \
  -H "Api-Secret: your-api-secret" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "my-app-logger"
  }'
```

```http
POST /v1/loggers HTTP/1.1
Host: api.hola.cloud
Api-Key: your-api-key
Api-Secret: your-api-secret
Content-Type: application/json

{
  "name": "my-app-logger"
}
```

## 响应

```json
{
  "id": "logger_xyz789",
  "name": "my-app-logger",
  "secret": "lgs_abc123def456",
  "created_at": "2025-06-20T14:22:00Z"
}
```

## 错误代码

| 代码 | 描述 |
|------|------|
| 400 | `name` 字段缺失或无效 |
| 401 | 缺少或无效的 API 凭据 |
| 409 | 此名称的日志记录器已存在 |
