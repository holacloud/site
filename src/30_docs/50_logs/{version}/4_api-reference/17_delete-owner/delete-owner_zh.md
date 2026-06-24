
# 删除 Logger 所有者

从日志记录器中移除所有者。所有者不能移除自己。

## 认证

需要管理凭据：

- `Api-Key` — 您的 API 密钥
- `Api-Secret` — 您的 API 密钥密码

## 路径参数

| 参数 | 描述 |
|------|-------------|
| `id` | 日志记录器的唯一标识符 |
| `ownerId` | 要移除的所有者的用户 ID |

## 请求

```bash
curl -X DELETE "https://api.hola.cloud/v1/loggers/logger_xyz789/owners/user_abc123" \
  -H "Api-Key: LOGGER_API_KEY" \
  -H "Api-Secret: LOGGER_API_SECRET"
```

```http
DELETE /v1/loggers/logger_xyz789/owners/user_abc123 HTTP/1.1
Host: api.hola.cloud
Api-Key: LOGGER_API_KEY
Api-Secret: LOGGER_API_SECRET
```

## 响应

返回更新后的所有者列表。

```json
["user_def456"]
```

## 错误代码

| 代码 | 描述 |
|------|------|
| 400 | 所有者不能移除自己 |
| 401 | 缺少或无效的 API 凭据 |
| 403 | API 凭据无权访问此日志记录器 |
| 404 | 日志记录器未找到 |
