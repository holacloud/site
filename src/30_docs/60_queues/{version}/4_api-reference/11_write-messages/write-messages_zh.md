
# Write Messages

向队列写入消息。接受单个 JSON 对象或 JSON 对象流（NDJSON）。

## 认证

公开 — 无需认证。

## 路径参数

| 参数 | 描述 |
|------|-------------|
| `id` | 队列的唯一标识符 |

## 请求体

接受 `application/json`（单条消息）或 `application/x-ndjson`（多条消息，每行一条）。

每条消息是任意 JSON 对象。服务器在入队时自动添加 `id` 和 `timestamp` 字段。

```json
{
  "type": "order.created",
  "order_id": "ord-1234",
  "amount": 49.99
}
```

或多条消息作为 NDJSON：

```
{"type": "order.created", "order_id": "ord-1234", "amount": 49.99}
{"type": "pageview", "url": "/home", "user": "u1"}
```

## 请求

```bash
# 单条消息
curl -X POST "https://api.hola.cloud/v1/queues/a1b2c3d4-e5f6-7890-abcd-ef1234567890" \
  -H "Content-Type: application/json" \
  -d '{
    "type": "order.created",
    "order_id": "ord-1234",
    "amount": 49.99
  }'

# 多条消息 (NDJSON)
curl -X POST "https://api.hola.cloud/v1/queues/a1b2c3d4-e5f6-7890-abcd-ef1234567890" \
  -H "Content-Type: application/x-ndjson" \
  -d '{"type": "order.created", "order_id": "ord-1234", "amount": 49.99}
{"type": "pageview", "url": "/home", "user": "u1"}'
```

```http
POST /v1/queues/a1b2c3d4-e5f6-7890-abcd-ef1234567890 HTTP/1.1
Host: api.hola.cloud
Content-Type: application/x-ndjson

{"type": "order.created", "order_id": "ord-1234", "amount": 49.99}
{"type": "pageview", "url": "/home", "user": "u1"}
```

## 响应

```json
{
  "written": 2
}
```

## 错误代码

| 代码 | 描述 |
|------|------|
| 400 | JSON 格式错误或内容类型无效 |
| 404 | 队列未找到 |
| 413 | 负载过大 |
