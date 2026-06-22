
# Delete Logger

永久删除日志记录器及其所有关联的日志条目。

## 认证

需要管理凭据：

- `Api-Key` — 您的 API 密钥
- `Api-Secret` — 您的 API 密钥密码

## 路径参数

| 参数 | 描述 |
|------|-------------|
| `id` | 日志记录器的唯一标识符 |

## 请求

```bash
curl -X DELETE "https://api.hola.cloud/v1/loggers/logger_xyz789" \
  -H "Api-Key: your-api-key" \
  -H "Api-Secret: your-api-secret"
```

```http
DELETE /v1/loggers/logger_xyz789 HTTP/1.1
Host: api.hola.cloud
Api-Key: your-api-key
Api-Secret: your-api-secret
```

## 响应

HTTP `204 无内容`。

## 错误代码

| 代码 | 描述 |
|------|------|
| 401 | 缺少或无效的 API 凭据 |
| 403 | API 凭据无权访问此日志记录器 |
| 404 | 日志记录器未找到 |
