
# Delete Queue

Delete a queue and all its pending messages permanently.

## Authentication

Public — no authentication required.

## Path Parameters

| Parameter | Description |
|-----------|-------------|
| `id` | The unique identifier of the queue |

## Request

```bash
curl -X DELETE "https://api.hola.cloud/v1/queues/a1b2c3d4-e5f6-7890-abcd-ef1234567890"
```

```http
DELETE /v1/queues/a1b2c3d4-e5f6-7890-abcd-ef1234567890 HTTP/1.1
Host: api.hola.cloud
```

## Response

HTTP `204 No Content`.

## Error Codes

| Code | Description |
|------|-------------|
| 404 | Queue not found |
