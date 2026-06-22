# List Tasks

Lists tasks for a scheduler, with separate pagination for scheduled and in-flight tasks.

## Authentication

This endpoint is public. No authentication required.

## Path Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| id | string | The unique identifier of the scheduler |

## Query Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| scheduled_page | integer | Page number for scheduled tasks (default: 1) |
| scheduled_per_page | integer | Items per page for scheduled tasks (default: 20) |
| inflight_page | integer | Page number for in-flight tasks (default: 1) |
| inflight_per_page | integer | Items per page for in-flight tasks (default: 20) |
| search | string | Search tasks by payload content |
| label | string | Filter by label (format: key:value) |

## Example Request

```bash
curl "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks?label=priority:high"
```

## Example Response

```http
HTTP/1.1 200 OK
Content-Type: application/json
```

```json
{
  "scheduled": [
    {
      "id": "task-x1y2z3",
      "future": "2025-06-21T12:01:01Z",
      "labels": ["project:onboarding", "priority:high"]
    }
  ],
  "inflight": [],
  "scheduled_meta": {
    "total": 1,
    "page": 1,
    "per_page": 25,
    "total_pages": 1
  },
  "inflight_meta": {
    "total": 0,
    "page": 1,
    "per_page": 25,
    "total_pages": 0
  }
}
```

## Error Codes

| Status | Code | Description |
|--------|------|-------------|
| 400 | validation_error | Invalid pagination or filters |
| 404 | not_found | Scheduler not found |
| 500 | internal_error | Internal server error |
