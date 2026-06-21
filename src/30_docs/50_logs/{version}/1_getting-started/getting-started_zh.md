# InstantLogs快速入门

本指南将帮助您创建日志记录器、接收日志条目并实时流式传输事件。

## 第一步：创建日志记录器

首先，使用管理 API 创建一个日志记录器。您需要具有管理权限的 API 密钥。

```bash
curl -X POST "https://api.hola.cloud/v1/loggers" \
  -H "Api-Key: your-api-key" \
  -H "Api-Secret: your-api-secret" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "my-app-logger"
  }'
```

预期响应：

```json
{
  "id": "logger_xyz789",
  "name": "my-app-logger",
  "secret": "lgs_abc123def456",
  "created_at": "2025-06-20T14:22:00Z"
}
```

请安全保存 `secret` 值——它用于验证数据操作。

## 第二步：接收日志条目

使用 Logger Secret 向您的日志记录器发送日志条目：

```bash
curl -X POST "https://api.hola.cloud/v1/loggers/logger_xyz789/ingest" \
  -H "X-Instantlogs-Event-Secret: lgs_abc123def456" \
  -H "Content-Type: application/json" \
  -d '{
    "message": "来自我的应用的问候！",
    "level": "info",
    "service": "web",
    "timestamp": "2025-06-20T14:22:30Z"
  }'
```

预期响应：

```json
{
  "ingested": 1
}
```

## 第三步：实时流式传输日志

建立流式传输连接，实时查看到达的日志：

```bash
curl "https://api.hola.cloud/v1/loggers/logger_xyz789/events" \
  -H "X-Instantlogs-Event-Secret: lgs_abc123def456"
```

服务器将保持连接打开状态，并以 NDJSON 格式推送新的日志条目：

```
{"message":"来自我的应用的问候！","level":"info","service":"web","timestamp":"2025-06-20T14:22:30Z","id":"evt_001"}
```

使用 `Accept: text/plain` 请求头，您可以改为接收纯文本行。

## 第四步：过滤流

添加正则表达式模式来过滤事件：

```bash
curl "https://api.hola.cloud/v1/loggers/logger_xyz789/filter?pattern=error" \
  -H "X-Instantlogs-Event-Secret: lgs_abc123def456"
```

只有消息匹配 `error` 模式的日志条目才会被流式传输。

## 总结

您已创建了日志记录器、接收了日志、实时流式传输了事件，并应用了正则表达式过滤。InstantLogs 让集中化和实时观察应用程序日志变得简单。
