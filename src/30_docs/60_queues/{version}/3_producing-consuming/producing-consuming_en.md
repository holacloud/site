# Producing Consuming

This document covers the message format, writing strategies, reading with long-poll, error handling, and common consumer patterns.

## Message Format

Each message in Tailon is an arbitrary JSON object. Tailon does not enforce a schema — you decide the structure. However, every message is assigned metadata by the server:

```json
{
  "id": "msg-001",
  "type": "order.created",
  "order_id": "ord-1234",
  "amount": 49.99,
  "timestamp": "2025-06-21T12:00:01Z"
}
```

The `id` and `timestamp` fields are automatically added when the message is enqueued.

## Writing Messages

### Write a Single Message

Send a single JSON object with `Content-Type: application/json`:

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

Expected response:

```json
{
  "written": 1
}
```

### Write Multiple Messages (NDJSON)

Use `Content-Type: application/x-ndjson` to send multiple messages in a single request, one JSON object per line:

```bash
curl -X POST "https://api.hola.cloud/v1/queues/a1b2c3d4-e5f6-7890-abcd-ef1234567890" \
  -H "Content-Type: application/x-ndjson" \
  -d '{"type": "pageview", "url": "/home", "user": "u1"}
{"type": "pageview", "url": "/pricing", "user": "u2"}
{"type": "pageview", "url": "/docs", "user": "u1"}'
```

Expected response:

```json
{
  "written": 3
}
```

## Reading Messages

Tailon uses **long-poll** — the request blocks until at least one message is available or a server timeout is reached. Use the `Limit` header to control the batch size.

```bash
curl -X GET "https://api.hola.cloud/v1/queues/a1b2c3d4-e5f6-7890-abcd-ef1234567890" \
  -H "Limit: 10"
```

```http
GET /v1/queues/a1b2c3d4-e5f6-7890-abcd-ef1234567890 HTTP/1.1
Host: api.hola.cloud
Limit: 10
```

Expected response (array of messages):

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

Messages are returned in FIFO order and removed from the queue upon delivery.

## Error Handling

| Status Code | Meaning |
|-------------|---------|
| 200 | Success (messages returned or written) |
| 204 | No content — queue exists but is empty (long-poll timed out) |
| 400 | Invalid request body or headers |
| 404 | Queue not found |
| 409 | Queue operation conflict |
| 500 | Internal server error |

Producer errors are typically caused by malformed JSON. Consumer errors often indicate a missing queue or an invalid `Limit` value.

## Consumer Patterns

### Single Consumer

A single worker polls the queue in a loop, processing one batch at a time:

```bash
while true; do
  curl -X GET "https://api.hola.cloud/v1/queues/my-queue-id" -H "Limit: 5"
  # process each message
done
```

### Multiple Consumers

Multiple workers can poll the same queue. Each message is delivered to exactly one consumer because messages are removed on read.

### Fan-Out via Multiple Queues

If you need the same message to reach multiple consumers, create separate queues and have the producer write to each one, or build a distributor process that reads from one queue and writes to others.

### Priority via Dedicated Queues

Use separate queues for different message priorities. High-priority workers poll the high-priority queue more frequently.
