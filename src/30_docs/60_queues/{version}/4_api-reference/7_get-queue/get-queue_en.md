
# Get Queue

Get details of a specific queue by its ID.

## Authentication

Public — no authentication required.

## Path Parameters

| Parameter | Description |
|-----------|-------------|
| `id` | The unique identifier of the queue |

## Request

```bash
curl "https://api.hola.cloud/v1/queues/a1b2c3d4-e5f6-7890-abcd-ef1234567890"
```

```http
GET /v1/queues/a1b2c3d4-e5f6-7890-abcd-ef1234567890 HTTP/1.1
Host: api.hola.cloud
```

## Response

```json
{
  "id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
  "name": "my-queue",
  "length": 5,
  "reads": 10,
  "writes": 15
}
```

## Error Codes

| Code | Description |
|------|-------------|
| 404 | Queue not found |
