
# 添加 Logger 所有者

为日志记录器添加所有者。所有者拥有对日志记录器的管理访问权限。

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
|------|------|------|------|
| `user_id` | string | 是 | 要添加为所有者的用户 ID |

```json
{
  "user_id": "user_abc123"
}
```

## 请求

```bash
curl -X POST "https://api.hola.cloud/v1/loggers/logger_xyz789/owners" \
  -H "Api-Key: LOGGER_API_KEY" \
  -H "Api-Secret: LOGGER_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "user_abc123"
  }'
```

```http
POST /v1/loggers/logger_xyz789/owners HTTP/1.1
Host: api.hola.cloud
Api-Key: LOGGER_API_KEY
Api-Secret: LOGGER_API_SECRET
Content-Type: application/json

{
  "user_id": "user_abc123"
}
```

## 响应

返回更新后的所有者列表。

```json
["user_def456", "user_abc123"]
```

## 错误代码

| 代码 | 描述 |
|------|------|
| 400 | `user_id` 字段缺失或无效 |
| 401 | 缺少或无效的 API 凭据 |
| 403 | API 凭据无权访问此日志记录器 |
| 404 | 日志记录器未找到 |
