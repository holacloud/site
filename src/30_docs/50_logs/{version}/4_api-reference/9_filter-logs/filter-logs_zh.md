
# Filter Logs

使用正则表达式模式实时流式传输和过滤日志条目。

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
| `pattern` | string | — | 用于按消息内容过滤日志条目的正则表达式模式 |
| `follow` | bool | `false` | 如果为 `true`，保持连接打开并流式传输新条目 |

## 请求

```bash
curl "https://api.hola.cloud/v1/loggers/logger_xyz789/filter?pattern=error&follow=true" \
  -H "X-Instantlogs-Event-Secret: lgs_abc123def456"
```

```http
GET /v1/loggers/logger_xyz789/filter?pattern=error&follow=true HTTP/1.1
Host: api.hola.cloud
X-Instantlogs-Event-Secret: lgs_abc123def456
```

## 响应

返回匹配的日志条目为 NDJSON 格式：

```
{"message":"数据库连接错误","level":"error","service":"db","timestamp":"2025-06-20T14:22:31Z","id":"evt_002"}
{"message":"超时","level":"error","service":"web","timestamp":"2025-06-20T14:22:32Z","id":"evt_003"}
```

当 `follow=true` 时，连接保持打开状态，新的匹配条目到达时会被流式传输。

## 错误代码

| 代码 | 描述 |
|------|------|
| 400 | 无效的正则表达式模式 |
| 401 | 缺少或无效的事件密钥 |
| 404 | 日志记录器未找到 |
