# Getting Started with Tailon

This guide walks you through creating a queue, writing messages, reading them via long-poll, checking queue stats, and cleaning up.

## Prerequisites

- A HolaCloud account.
- [curl](https://curl.se/) installed on your machine.

## Step 1: Create a Queue

Create a new queue by sending a POST request with a name:

```bash
curl -X POST "https://api.hola.cloud/v1/queues" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "my-queue"
  }'
```

Expected response:

```json
{
  "id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
  "name": "my-queue",
  "length": 0,
  "reads": 0,
  "writes": 0
}
```

Save the returned `id` — you will use it in subsequent requests.

## Step 2: Write Messages

Tailon accepts a stream of JSON objects, one per line (NDJSON). Write three messages:

```bash
curl -X POST "https://api.hola.cloud/v1/queues/a1b2c3d4-e5f6-7890-abcd-ef1234567890" \
  -H "Content-Type: application/x-ndjson" \
  -d '{"type": "greeting", "text": "Hello"}
{"type": "greeting", "text": "Hola"}
{"type": "command", "action": "ping"}'
```

Expected response:

```json
{
  "written": 3
}
```

## Step 3: Read Messages

Use a long-poll GET to retrieve messages. The `Limit` header controls how many messages to return:

```bash
curl -X GET "https://api.hola.cloud/v1/queues/a1b2c3d4-e5f6-7890-abcd-ef1234567890" \
  -H "Limit: 2"
```

Expected response:

```json
[
  {
    "id": "m1",
    "type": "greeting",
    "text": "Hello",
    "timestamp": "2025-06-21T12:00:01Z"
  },
  {
    "id": "m2",
    "type": "greeting",
    "text": "Hola",
    "timestamp": "2025-06-21T12:00:01Z"
  }
]
```

Messages are returned in FIFO order. Repeat the request to retrieve the remaining messages.

## Step 4: Check Queue Stats

Inspect the queue details, including its current length and total reads/writes:

```bash
curl "https://api.hola.cloud/v1/queues/a1b2c3d4-e5f6-7890-abcd-ef1234567890"
```

Expected response:

```json
{
  "id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
  "name": "my-queue",
  "length": 1,
  "reads": 2,
  "writes": 3
}
```

## Step 5: Delete the Queue

When you are done, delete the queue to free resources:

```bash
curl -X DELETE "https://api.hola.cloud/v1/queues/a1b2c3d4-e5f6-7890-abcd-ef1234567890"
```

Expected response: HTTP `204 No Content`.

## Next Steps

- Learn how to manage multiple queues and monitor active clients in [Managing Queues](../2_managing-queues/managing-queues_en.md).
- Explore advanced producing and consuming patterns in [Producing and Consuming Messages](../3_producing-consuming/producing-consuming_en.md).
