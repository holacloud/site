
# GET /v1/queues/{id}

Read messages from a queue using long-poll. The request blocks until at least one message is available or the server timeout is reached.

## Authentication

Public — no authentication required.

## Path Parameters

| Parameter | Description |
|-----------|-------------|
| `id` | The unique identifier of the queue |

## Request Headers

| Header | Description |
|--------|-------------|
| `Limit` | Maximum number of messages to return (default varies by server) |

## Request

```bash
curl -X GET "https://api.hola.cloud/v1/queues/a1b2c3d4-e5f6-7890-abcd-ef1234567890" \
  -H "Limit: 10"
```

```http
GET /v1/queues/a1b2c3d4-e5f6-7890-abcd-ef1234567890 HTTP/1.1
Host: api.hola.cloud
Limit: 10
```

## Response

Messages are returned in FIFO order and removed from the queue upon delivery.

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

If the queue is empty and no messages arrive before the timeout, the server returns HTTP `204 No Content`.

## Error Codes

| Code | Description |
|------|-------------|
| 204 | No messages available (long-poll timed out) |
| 400 | Invalid `Limit` value |
| 404 | Queue not found |
