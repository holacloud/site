
# Create Logger API Key

为日志记录器创建新的 API 密钥。API 密钥允许对日志记录器进行管理访问。

## 认证

需要管理凭据：

- `Api-Key` — 您的 API 密钥
- `Api-Secret` — 您的 API 密钥密码

## 路径参数

| 参数 | 描述 |
|------|-------------|
| `id` | 日志记录器的唯一标识符 |

## 请求体

| 字段 | 类型 | 必填 | 描述 |
|------|------|------|-------------|
| `name` | string | 是 | 用于标识此 API 密钥的标签 |

```json
{
  "name": "ci-cd-key"
}
```

## 请求

```bash
curl -X POST "https://api.hola.cloud/v1/loggers/logger_xyz789/apiKeys" \
  -H "Api-Key: LOGGER_API_KEY" \
  -H "Api-Secret: LOGGER_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "ci-cd-key"
  }'
```

```http
POST /v1/loggers/logger_xyz789/apiKeys HTTP/1.1
Host: api.hola.cloud
Api-Key: LOGGER_API_KEY
Api-Secret: LOGGER_API_SECRET
Content-Type: application/json

{
  "name": "ci-cd-key"
}
```

## 响应

```json
{
  "id": "ak_123456",
  "name": "ci-cd-key",
  "key": "lgs_api_key_value",
  "created_at": "2025-06-20T15:00:00Z"
}
```

## 错误代码

| 代码 | 描述 |
|------|------|
| 400 | `name` 字段缺失或无效 |
| 401 | 缺少或无效的 API 凭据 |
| 403 | API 凭据无权访问此日志记录器 |
| 404 | 日志记录器未找到 |
