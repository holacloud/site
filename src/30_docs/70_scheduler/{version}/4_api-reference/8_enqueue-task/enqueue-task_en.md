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
| id | string | Optional client-provided task ID |
| future | string | ISO 8601 timestamp for when the task should become available |
| delay | integer | Delay in seconds from now (alternative to future) |
| payload | object | Arbitrary JSON payload for the worker |
| labels | object | Optional key-value pairs for filtering |

```json
{
  "payload": {
    "type": "send_email",
    "to": "user@example.com",
    "template": "welcome"
  },
  "delay": 60,
  "labels": {
    "project": "onboarding",
    "priority": "high"
  }
}
```

## Example Request

```bash
curl -X POST "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks" \
  -H "X-API-Key: YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "payload": {
      "type": "send_email",
      "to": "user@example.com",
      "template": "welcome"
    },
    "delay": 60,
    "labels": {
      "project": "onboarding",
      "priority": "high"
    }
  }'
```

## Example Response

```http
HTTP/1.1 201 Created
Content-Type: application/json
```

```json
{
  "id": "task-x1y2z3",
  "scheduler_id": "sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890",
  "payload": {
    "type": "send_email",
    "to": "user@example.com",
    "template": "welcome"
  },
  "state": "pending",
  "available_at": "2025-06-21T12:01:01Z",
  "labels": {
    "project": "onboarding",
    "priority": "high"
  }
}
```

## Error Codes

| Status | Code | Description |
|--------|------|-------------|
| 400 | invalid_request | Missing or invalid request body |
| 401 | unauthorized | Missing or invalid API key |
| 404 | not_found | Scheduler not found |
| 500 | internal_error | Internal server error |
