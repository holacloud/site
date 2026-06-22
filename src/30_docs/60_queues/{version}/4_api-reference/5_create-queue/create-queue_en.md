
# Create Queue

Create a new queue.

## Authentication

Public — no authentication required.

## Request Body

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `name` | string | yes | A name for the queue |

```json
{
  "name": "my-queue"
}
```

## Request

```bash
curl -X POST "https://api.hola.cloud/v1/queues" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "my-queue"
  }'
```

```http
POST /v1/queues HTTP/1.1
Host: api.hola.cloud
Content-Type: application/json

{
  "name": "my-queue"
}
```

## Response

```json
{
  "id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
  "name": "my-queue",
  "length": 0,
  "reads": 0,
  "writes": 0
}
```

## Error Codes

| Code | Description |
|------|-------------|
| 400 | Missing or invalid `name` field |
| 409 | A queue with this name already exists |
