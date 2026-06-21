# 发送日志

InstantLogs 支持两种接收日志数据的方法：直接向日志记录器接收和帧事件接收。两种方法都需要相应的密钥进行认证。

## 认证

数据接收端点使用 Logger Secret 进行认证。您可以通过以下两种方式之一提供：

- **请求头**：`X-Instantlogs-Event-Secret: <您的日志记录器密钥>`
- **Bearer 令牌**：`Authorization: Bearer <您的日志记录器密钥>`

## 方法一：直接接收

使用 `POST /v1/loggers/{id}/ingest` 端点将日志条目直接发送到特定日志记录器。

### 单条条目

```bash
curl -X POST "https://api.hola.cloud/v1/loggers/logger_xyz789/ingest" \
  -H "X-Instantlogs-Event-Secret: lgs_abc123def456" \
  -H "Content-Type: application/json" \
  -d '{
    "message": "用户登录成功",
    "level": "info",
    "service": "auth-service",
    "timestamp": "2025-06-20T15:00:00Z",
    "metadata": {
      "user_id": "usr_123",
      "ip": "192.168.1.1"
    }
  }'
```

预期响应：

```json
{
  "ingested": 1
}
```

### 批量条目

您可以通过提供数组在单个请求中发送多条日志条目：

```bash
curl -X POST "https://api.hola.cloud/v1/loggers/logger_xyz789/ingest" \
  -H "X-Instantlogs-Event-Secret: lgs_abc123def456" \
  -H "Content-Type: application/json" \
  -d '[
    {
      "message": "请求开始",
      "level": "info",
      "service": "api-gateway",
      "timestamp": "2025-06-20T15:00:00Z"
    },
    {
      "message": "请求完成",
      "level": "info",
      "service": "api-gateway",
      "timestamp": "2025-06-20T15:00:05Z"
    }
  ]'
```

预期响应：

```json
{
  "ingested": 2
}
```

## 方法二：帧事件接收

对于高吞吐量场景或使用二进制协议时，请使用帧事件端点 `POST /v1/ingest/events`。此端点使用 Event Secret 进行认证，并接受帧事件流。

```bash
curl -X POST "https://api.hola.cloud/v1/ingest/events" \
  -H "X-Instantlogs-Event-Secret: evt_secret_789" \
  -H "Content-Type: application/json" \
  -d '{
    "logger_id": "logger_xyz789",
    "events": [
      {
        "message": "帧事件 1",
        "level": "warn",
        "timestamp": "2025-06-20T15:01:00Z"
      },
      {
        "message": "帧事件 2",
        "level": "error",
        "timestamp": "2025-06-20T15:01:01Z"
      }
    ]
  }'
```

## HTTP 请求参考

```http
POST /v1/loggers/{id}/ingest HTTP/1.1
Host: api.hola.cloud
X-Instantlogs-Event-Secret: lgs_abc123def456
Content-Type: application/json

{
  "message": "示例日志",
  "level": "info",
  "service": "web",
  "timestamp": "2025-06-20T15:00:00Z"
}

POST /v1/ingest/events HTTP/1.1
Host: api.hola.cloud
X-Instantlogs-Event-Secret: evt_secret_789
Content-Type: application/json

{
  "logger_id": "logger_xyz789",
  "events": [ ... ]
}
```
