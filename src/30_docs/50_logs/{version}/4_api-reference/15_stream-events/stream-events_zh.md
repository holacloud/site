
# Stream Events

实时流式传输日志事件。支持 NDJSON 和纯文本格式。

## 认证

需要数据凭据：

- `X-Instantlogs-Event-Secret: <secret>` 或 `Authorization: Bearer <secret>`

## 路径参数

| 参数 | 描述 |
|------|-------------|
| `id` | 日志记录器的唯一标识符 |

## 查询参数

| 参数 | 类型 | 默认值 | 描述 |
|------|------|--------|-------------|
| `regex` | string | — | 用于按消息内容过滤事件的正则表达式模式 |
| `follow` | bool | `true` | 如果为 `true`，保持连接打开并流式传输新事件 |
| `format` | string | `ndjson` | 输出格式：`ndjson` 或 `text` |

## 请求

```bash
# 以 NDJSON 格式流式传输所有事件
curl "https://api.hola.cloud/v1/loggers/logger_xyz789/events" \
  -H "X-Instantlogs-Event-Secret: lgs_abc123def456"

# 以纯文本格式流式传输
curl "https://api.hola.cloud/v1/loggers/logger_xyz789/events?format=text" \
  -H "X-Instantlogs-Event-Secret: lgs_abc123def456"

# 使用正则表达式过滤
curl "https://api.hola.cloud/v1/loggers/logger_xyz789/events?regex=error&format=ndjson" \
  -H "X-Instantlogs-Event-Secret: lgs_abc123def456"
```

```http
GET /v1/loggers/logger_xyz789/events?format=ndjson HTTP/1.1
Host: api.hola.cloud
X-Instantlogs-Event-Secret: lgs_abc123def456
```

## 响应

NDJSON 格式（默认）：

```
{"message":"用户登录成功","level":"info","service":"web","timestamp":"2025-06-20T14:22:30Z","id":"evt_001"}
{"message":"数据库连接错误","level":"error","service":"db","timestamp":"2025-06-20T14:22:31Z","id":"evt_002"}
```

文本格式：

```
14:22:30 [info] [web] 用户登录成功
14:22:31 [error] [db] 数据库连接错误
```

## 错误代码

| 代码 | 描述 |
|------|------|
| 400 | 无效的查询参数 |
| 401 | 缺少或无效的事件密钥 |
| 404 | 日志记录器未找到 |
