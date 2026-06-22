# Reschedule Task

Reschedules a task with a new delay or future time.

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
| future | string | ISO 8601 timestamp for when the task should become available |
| delay | string | Go duration string from now, such as `120s` or `5m` (alternative to future) |

```json
{
  "delay": "120s"
}
```

## Example Request

```bash
curl -X POST "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks/task-x1y2z3/reschedule" \
  -H "X-API-Key: YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "delay": "120s"
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
  "future": "2025-06-21T12:03:01Z"
}
```

## Error Codes

| Status | Code | Description |
|--------|------|-------------|
| 400 | invalid_json | Invalid JSON payload |
| 400 | validation_error | Missing or invalid future/delay |
| 401 | unauthorized | Missing or invalid API key |
| 404 | task_not_found | Task not found |
| 409 | task_in_flight | Task is currently leased |
| 500 | internal_error | Internal server error |
