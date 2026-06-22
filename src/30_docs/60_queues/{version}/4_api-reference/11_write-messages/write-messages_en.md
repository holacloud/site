
# Write Messages

Write messages to a queue. Accepts a single JSON object or a stream of JSON objects (NDJSON).

## Authentication

Public — no authentication required.

## Path Parameters

| Parameter | Description |
|-----------|-------------|
| `id` | The unique identifier of the queue |

## Request Body

Accepts either `application/json` (single message) or `application/x-ndjson` (multiple messages, one per line).

Each message is an arbitrary JSON object. The server automatically adds `id` and `timestamp` fields upon enqueue.

```json
{
  "type": "order.created",
  "order_id": "ord-1234",
  "amount": 49.99
}
```

Or multiple messages as NDJSON:

```
{"type": "order.created", "order_id": "ord-1234", "amount": 49.99}
{"type": "pageview", "url": "/home", "user": "u1"}
```

## Request

```bash
# Single message
curl -X POST "https://api.hola.cloud/v1/queues/a1b2c3d4-e5f6-7890-abcd-ef1234567890" \
  -H "Content-Type: application/json" \
  -d '{
    "type": "order.created",
    "order_id": "ord-1234",
    "amount": 49.99
  }'

# Multiple messages (NDJSON)
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

## Response

```json
{
  "written": 2
}
```

## Error Codes

| Code | Description |
|------|-------------|
| 400 | Malformed JSON or invalid content type |
| 404 | Queue not found |
| 413 | Payload too large |
