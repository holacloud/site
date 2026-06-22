
# Read Messages

使用长轮询从队列读取消息。请求会阻塞，直到至少有一条消息可用或达到服务器超时时间。

## 认证

公开 — 无需认证。

## 路径参数

| 参数 | 描述 |
|------|-------------|
| `id` | 队列的唯一标识符 |

## 请求头

| 请求头 | 描述 |
|--------|------|
| `Limit` | 返回的最大消息数（默认值由服务器决定） |

## 请求

```bash
curl -X GET "https://api.hola.cloud/v1/queues/a1b2c3d4-e5f6-7890-abcd-ef1234567890" \
  -H "Limit: 10"
```

```http
GET /v1/queues/a1b2c3d4-e5f6-7890-abcd-ef1234567890 HTTP/1.1
Host: api.hola.cloud
Limit: 10
```

## 响应

消息按 FIFO 顺序返回，并在交付时从队列中移除。

```json
[
  {
    "id": "msg-001",
    "type": "order.created",
    "order_id": "ord-1234",
    "amount": 49.99,
    "timestamp": "2025-06-21T12:00:01Z"
  },
  {
    "id": "msg-002",
    "type": "pageview",
    "url": "/home",
    "user": "u1",
    "timestamp": "2025-06-21T12:00:02Z"
  }
]
```

如果队列为空且在超时前没有消息到达，服务器返回 HTTP `204 无内容`。

## 错误代码

| 代码 | 描述 |
|------|------|
| 204 | 没有可用消息（长轮询超时） |
| 400 | `Limit` 值无效 |
| 404 | 队列未找到 |
