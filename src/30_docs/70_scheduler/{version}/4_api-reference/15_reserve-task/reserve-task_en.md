# POST /schedulers/{id}/tasks/reserve

Reserves one available task for processing. The task enters a lease period defined by `worktime`.

## Authentication

Requires authentication. Pass your API key via `X-API-Key` or `Authorization: Bearer` header.

## Path Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| id | string | The unique identifier of the scheduler |

## Request Body

| Field | Type | Description |
|-------|------|-------------|
| worktime | integer | Lease duration in seconds |

```json
{
  "worktime": 30
}
```

## Example Request

```bash
curl -X POST "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks/reserve" \
  -H "X-API-Key: YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "worktime": 30
  }'
```

## Example Response

```http
HTTP/1.1 200 OK
Content-Type: application/json
```

```json
{
  "id": "task-x1y2z3",
  "payload": {
    "type": "send_email",
    "to": "user@example.com",
    "template": "welcome"
  },
  "lease_expires_at": "2025-06-21T12:02:01Z"
}
```

## Error Codes

| Status | Code | Description |
|--------|------|-------------|
| 401 | unauthorized | Missing or invalid API key |
| 404 | not_found | Scheduler not found |
| 409 | conflict | No tasks available for reservation |
| 500 | internal_error | Internal server error |
