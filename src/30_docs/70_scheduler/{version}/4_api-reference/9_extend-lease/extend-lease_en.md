# Extend Lease

Extends the lease on a currently reserved task.

## Authentication

Requires authentication. Pass your API key via `X-API-Key` or `Authorization: Bearer` header.

## Path Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| id | string | The unique identifier of the scheduler |
| task | string | The unique identifier of the task |

## Request Body

| Field | Type | Description |
|-------|------|-------------|
| extension | integer | Additional lease time in seconds |

```json
{
  "extension": 30
}
```

## Example Request

```bash
curl -X POST "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks/task-x1y2z3/extend" \
  -H "X-API-Key: YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "extension": 30
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
  "lease_expires_at": "2025-06-21T12:02:31Z"
}
```

## Error Codes

| Status | Code | Description |
|--------|------|-------------|
| 400 | invalid_request | Missing or invalid extension value |
| 401 | unauthorized | Missing or invalid API key |
| 404 | not_found | Scheduler or task not found |
| 409 | conflict | Task is not currently reserved or lease already expired |
| 500 | internal_error | Internal server error |
