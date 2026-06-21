# Managing Queues

This document covers queue lifecycle operations: creation with optional configuration, listing, inspection, client monitoring, and deletion.

## Creating a Queue

Create a queue by sending a POST request:

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

The corresponding HTTP request looks like this:

```http
POST /v1/queues HTTP/1.1
Host: api.hola.cloud
Content-Type: application/json

{
  "name": "my-queue"
}
```

## Listing All Queues

Retrieve a list of all queues on your account:

```bash
curl "https://api.hola.cloud/v1/queues"
```

Expected response:

```json
[
  {
    "id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
    "name": "my-queue",
    "length": 5,
    "reads": 10,
    "writes": 15
  },
  {
    "id": "b2c3d4e5-f6a7-8901-bcde-f12345678901",
    "name": "events-queue",
    "length": 0,
    "reads": 42,
    "writes": 42
  }
]
```

## Getting Queue Details

Inspect a single queue by its ID to see current length, total reads, and total writes:

```bash
curl "https://api.hola.cloud/v1/queues/a1b2c3d4-e5f6-7890-abcd-ef1234567890"
```

```http
GET /v1/queues/a1b2c3d4-e5f6-7890-abcd-ef1234567890 HTTP/1.1
Host: api.hola.cloud
```

Expected response:

```json
{
  "id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
  "name": "my-queue",
  "length": 5,
  "reads": 10,
  "writes": 15
}
```

## Monitoring Active Clients

List all currently connected streaming clients:

```bash
curl "https://api.hola.cloud/v1/clients"
```

```http
GET /v1/clients HTTP/1.1
Host: api.hola.cloud
```

Expected response:

```json
[
  {
    "id": "client-001",
    "queue_id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
    "connected_at": "2025-06-21T12:00:00Z",
    "user_agent": "curl/7.88.1"
  }
]
```

## Deleting a Queue

Permanently delete a queue and all its pending messages:

```bash
curl -X DELETE "https://api.hola.cloud/v1/queues/a1b2c3d4-e5f6-7890-abcd-ef1234567890"
```

```http
DELETE /v1/queues/a1b2c3d4-e5f6-7890-abcd-ef1234567890 HTTP/1.1
Host: api.hola.cloud
```

Expected response: HTTP `204 No Content`.

Once deleted, the queue ID is no longer valid. Any in-flight readers will receive an error on their next request.
