
# Delete Logger API Key

从日志记录器中删除 API 密钥。

## 认证

需要管理凭据：

- `Api-Key` — 您的 API 密钥
- `Api-Secret` — 您的 API 密钥密码

## 路径参数

| 参数 | 描述 |
|------|-------------|
| `id` | 日志记录器的唯一标识符 |
| `key` | 要删除的 API 密钥的 ID |

## 请求

```bash
curl -X DELETE "https://api.hola.cloud/v1/loggers/logger_xyz789/apiKeys/ak_123456" \
  -H "Api-Key: LOGGER_API_KEY" \
  -H "Api-Secret: LOGGER_API_SECRET"
```

```http
DELETE /v1/loggers/logger_xyz789/apiKeys/ak_123456 HTTP/1.1
Host: api.hola.cloud
Api-Key: LOGGER_API_KEY
Api-Secret: LOGGER_API_SECRET
```

## 响应

HTTP `204 无内容`。

## 错误代码

| 代码 | 描述 |
|------|------|
| 401 | 缺少或无效的 API 凭据 |
| 403 | API 凭据无权访问此日志记录器 |
| 404 | 日志记录器或 API 密钥未找到 |
