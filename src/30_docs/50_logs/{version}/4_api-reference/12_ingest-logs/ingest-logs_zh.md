
# POST /v1/loggers/{id}/ingest

向日志记录器接收日志条目。

## 认证

需要数据凭据：

- `X-Instantlogs-Event-Secret: <secret>` 或 `Authorization: Bearer <secret>`

## 路径参数

| 参数 | 描述 |
|------|-------------|
| `id` | 日志记录器的唯一标识符 |

## 请求体

单条日志条目：

| 字段 | 类型 | 必填 | 描述 |
|------|------|------|-------------|
| `message` | string | 是 | 日志消息内容 |
| `level` | string | 否 | 日志级别（如 `info`、`warn`、`error`） |
| `service` | string | 否 | 来源服务名称 |
| `timestamp` | string | 否 | ISO 8601 时间戳（默认为服务器时间） |

```json
{
  "message": "用户登录成功",
  "level": "info",
  "service": "web",
  "timestamp": "2025-06-20T14:22:30Z"
}
```

## 请求

```bash
curl -X POST "https://api.hola.cloud/v1/loggers/logger_xyz789/ingest" \
  -H "X-Instantlogs-Event-Secret: lgs_abc123def456" \
  -H "Content-Type: application/json" \
  -d '{
    "message": "用户登录成功",
    "level": "info",
    "service": "web",
    "timestamp": "2025-06-20T14:22:30Z"
  }'
```

```http
POST /v1/loggers/logger_xyz789/ingest HTTP/1.1
Host: api.hola.cloud
X-Instantlogs-Event-Secret: lgs_abc123def456
Content-Type: application/json

{
  "message": "用户登录成功",
  "level": "info",
  "service": "web",
  "timestamp": "2025-06-20T14:22:30Z"
}
```

## 响应

```json
{
  "ingested": 1
}
```

## 错误代码

| 代码 | 描述 |
|------|------|
| 400 | 请求体缺失或无效 |
| 401 | 缺少或无效的事件密钥 |
| 404 | 日志记录器未找到 |
| 413 | 负载过大 |
