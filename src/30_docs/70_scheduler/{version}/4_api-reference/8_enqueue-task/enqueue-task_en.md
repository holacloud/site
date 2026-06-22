# Enqueue Task

Enqueues a new task to be executed at a scheduled time.

## Authentication

Requires authentication. Pass your API key via `X-API-Key` or `Authorization: Bearer` header.

## Path Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| id | string | The unique identifier of the scheduler |

## Request Body

| Field | Type | Description |
|-------|------|-------------|
| id | string | Required task ID |
| future | string | ISO 8601 timestamp for when the task should become available |
| delay | string | Go duration string from now, such as `60s` or `5m` (alternative to future) |
| payload | object | Arbitrary JSON payload for the worker |
| labels | array of strings | Optional labels for filtering |

```json
{
  "id": "task-x1y2z3",
  "payload": {
    "type": "send_email",
    "to": "user@example.com",
    "template": "welcome"
  },
  "delay": "60s",
  "labels": ["project:onboarding", "priority:high"]
}
```

## Example Request

```bash
curl -X POST "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks" \
  -H "X-API-Key: YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "id": "task-x1y2z3",
    "payload": {
      "type": "send_email",
      "to": "user@example.com",
      "template": "welcome"
    },
    "delay": "60s",
    "labels": ["project:onboarding", "priority:high"]
  }'
```

## Example Response

```http
HTTP/1.1 202 Accepted
```

The response body is empty.

## Error Codes

| Status | Code | Description |
|--------|------|-------------|
| 400 | invalid_json | Invalid JSON payload |
| 400 | validation_error | Missing id, invalid future/delay, or invalid labels |
| 401 | unauthorized | Missing or invalid API key |
| 409 | task_already_exists | Task already exists |
| 500 | internal_error | Internal server error |
