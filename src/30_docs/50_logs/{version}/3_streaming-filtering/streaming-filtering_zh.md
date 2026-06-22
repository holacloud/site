# StreamingFiltering

InstantLogs 提供长连接的 HTTP 流式传输端点，能够近乎实时地交付日志事件。您可以通过正则表达式模式过滤流，并选择 NDJSON 或纯文本输出格式。

## 流式传输端点

### 事件流

`GET /v1/loggers/{id}/events` 端点流式传输所有接收到的日志条目。默认情况下，事件以 NDJSON（换行符分隔的 JSON）格式返回。

```bash
curl "https://api.hola.cloud/v1/loggers/logger_xyz789/events" \
  -H "X-Instantlogs-Event-Secret: lgs_abc123def456"
```

响应体的每一行都是一个表示日志条目的 JSON 对象：

```
{"id":"evt_001","message":"请求开始","level":"info","service":"api-gateway","timestamp":"2025-06-20T15:00:00Z"}
{"id":"evt_002","message":"请求完成","level":"info","service":"api-gateway","timestamp":"2025-06-20T15:00:05Z"}
```

要接收纯文本输出，请将 `Accept` 请求头设置为 `text/plain`：

```bash
curl "https://api.hola.cloud/v1/loggers/logger_xyz789/events" \
  -H "X-Instantlogs-Event-Secret: lgs_abc123def456" \
  -H "Accept: text/plain"
```

### 过滤流

`GET /v1/loggers/{id}/filter` 端点在流式传输前应用正则表达式过滤。仅传递消息与模式匹配的日志条目。

```bash
curl "https://api.hola.cloud/v1/loggers/logger_xyz789/filter?pattern=error|critical" \
  -H "X-Instantlogs-Event-Secret: lgs_abc123def456"
```

这仅流式传输消息字段中包含 `error` 或 `critical` 的条目。pattern 参数接受任何有效的 Go 正则表达式。

### 组合过滤与格式

您可以将正则表达式过滤与您偏好的输出格式结合使用：

```bash
curl "https://api.hola.cloud/v1/loggers/logger_xyz789/filter?pattern=^\\[WARN\\]" \
  -H "X-Instantlogs-Event-Secret: lgs_abc123def456" \
  -H "Accept: text/plain"
```

## 日志记录器统计信息

`GET /v1/loggers/{id}/stats` 端点返回日志记录器的统计信息，无需消费事件流。

```bash
curl "https://api.hola.cloud/v1/loggers/logger_xyz789/stats" \
  -H "X-Instantlogs-Event-Secret: lgs_abc123def456"
```

预期响应：

```json
{
  "logger_id": "logger_xyz789",
  "total_events": 15234,
  "events_last_hour": 342,
  "storage_bytes": 1048576,
  "created_at": "2025-06-20T14:22:00Z"
}
```

## HTTP 请求参考

```http
GET /v1/loggers/{id}/events HTTP/1.1
Host: api.hola.cloud
X-Instantlogs-Event-Secret: lgs_abc123def456

GET /v1/loggers/{id}/filter?pattern=error HTTP/1.1
Host: api.hola.cloud
X-Instantlogs-Event-Secret: lgs_abc123def456

GET /v1/loggers/{id}/stats HTTP/1.1
Host: api.hola.cloud
X-Instantlogs-Event-Secret: lgs_abc123def456
```
