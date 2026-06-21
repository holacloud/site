# 生产与消费消息

本文档涵盖消息格式、写入策略、长轮询读取、错误处理以及常见的消费模式。

## 消息格式

Tailon中的每条消息都是一个任意的JSON对象。Tailon不强制规定模式——您决定结构。但是，服务器会为每条消息分配元数据：

```json
{
  "id": "msg-001",
  "type": "order.created",
  "order_id": "ord-1234",
  "amount": 49.99,
  "timestamp": "2025-06-21T12:00:01Z"
}
```

`id`和`timestamp`字段在消息入队时自动添加。

## 写入消息

### 写入单条消息

使用`Content-Type: application/json`发送单个JSON对象：

```bash
curl -X POST "https://api.hola.cloud/v1/queues/a1b2c3d4-e5f6-7890-abcd-ef1234567890" \
  -H "Content-Type: application/json" \
  -d '{
    "type": "order.created",
    "order_id": "ord-1234",
    "amount": 49.99
  }'
```

```http
POST /v1/queues/a1b2c3d4-e5f6-7890-abcd-ef1234567890 HTTP/1.1
Host: api.hola.cloud
Content-Type: application/json

{
  "type": "order.created",
  "order_id": "ord-1234",
  "amount": 49.99
}
```

预期响应：

```json
{
  "written": 1
}
```

### 写入多条消息（NDJSON）

使用`Content-Type: application/x-ndjson`在单个请求中发送多条消息，每行一个JSON对象：

```bash
curl -X POST "https://api.hola.cloud/v1/queues/a1b2c3d4-e5f6-7890-abcd-ef1234567890" \
  -H "Content-Type: application/x-ndjson" \
  -d '{"type": "pageview", "url": "/home", "user": "u1"}
{"type": "pageview", "url": "/pricing", "user": "u2"}
{"type": "pageview", "url": "/docs", "user": "u1"}'
```

预期响应：

```json
{
  "written": 3
}
```

## 读取消息

Tailon使用**长轮询**——请求会阻塞，直到至少有一条消息可用或达到服务器超时。使用`Limit`头部控制批量大小。

```bash
curl -X GET "https://api.hola.cloud/v1/queues/a1b2c3d4-e5f6-7890-abcd-ef1234567890" \
  -H "Limit: 10"
```

```http
GET /v1/queues/a1b2c3d4-e5f6-7890-abcd-ef1234567890 HTTP/1.1
Host: api.hola.cloud
Limit: 10
```

预期响应（消息数组）：

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

消息按FIFO顺序返回，并在投递时从队列中移除。

## 错误处理

| 状态码 | 含义 |
|--------|------|
| 200 | 成功（返回或写入了消息） |
| 204 | 无内容 — 队列存在但为空（长轮询超时） |
| 400 | 请求体或头部无效 |
| 404 | 队列未找到 |
| 409 | 队列操作冲突 |
| 500 | 服务器内部错误 |

生产者错误通常由格式错误的JSON引起。消费者错误通常表示队列不存在或`Limit`值无效。

## 消费模式

### 单一消费者

单个工作者循环轮询队列，一次处理一批：

```bash
while true; do
  curl -X GET "https://api.hola.cloud/v1/queues/my-queue-id" -H "Limit: 5"
  # 处理每条消息
done
```

### 多个消费者

多个工作者可以轮询同一个队列。每条消息恰好投递给一个消费者，因为消息在读取时即被移除。

### 通过多个队列实现扇出

如果需要同一条消息到达多个消费者，请创建独立的队列，让生产者向每个队列写入，或者构建一个分发器进程，从一个队列读取并写入其他队列。

### 通过专用队列实现优先级

为不同优先级的消息使用单独的队列。高优先级的工作者更频繁地轮询高优先级队列。
