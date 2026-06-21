# GET /schedulers/{id}/tasks/stream

Streams task events in real time using Server-Sent Events (SSE).

## Authentication

This endpoint is public. No authentication required.

## Path Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| id | string | The unique identifier of the scheduler |

## Example Request

```bash
curl -N "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks/stream"
```

## Example Response

```http
HTTP/1.1 200 OK
Content-Type: text/event-stream
Cache-Control: no-cache
Connection: keep-alive
```

```
event: task_available
data: {"id":"task-x1y2z3","state":"pending","available_at":"2025-06-21T12:01:01Z"}

event: task_reserved
data: {"id":"task-x1y2z3","state":"inflight","lease_expires_at":"2025-06-21T12:02:01Z"}

event: task_completed
data: {"id":"task-x1y2z3","state":"completed"}

event: lease_expired
data: {"id":"task-x1y2z3","state":"pending"}
```

## Error Codes

| Status | Code | Description |
|--------|------|-------------|
| 404 | not_found | Scheduler not found |
| 500 | internal_error | Internal server error |
