
# List Clients

列出当前连接到所有队列的流式客户端。

## 认证

公开 — 无需认证。

## 请求

```bash
curl "https://api.hola.cloud/v1/clients"
```

```http
GET /v1/clients HTTP/1.1
Host: api.hola.cloud
```

## 响应

```json
[
  {
    "id": "client-001",
    "queue_id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
    "connected_at": "2025-06-21T12:00:00Z",
    "user_agent": "curl/7.88.1"
  },
  {
    "id": "client-002",
    "queue_id": "b2c3d4e5-f6a7-8901-bcde-f12345678901",
    "connected_at": "2025-06-21T12:00:05Z",
    "user_agent": "TailonClient/1.0"
  }
]
```

## 错误代码

| 代码 | 描述 |
|------|------|
| 500 | 服务器内部错误 |
