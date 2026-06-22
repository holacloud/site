# Task Stream

Streams task snapshots using Server-Sent Events (SSE).

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
event: snapshot
data: {"scheduler_id":"sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890","generated_at":"2025-06-21T12:00:01Z","scheduled":[{"id":"task-x1y2z3","future":"2025-06-21T12:01:01Z","labels":["priority:high"]}],"inflight":[],"scheduled_meta":{"page":1,"per_page":25,"total":1,"total_pages":1},"inflight_meta":{"page":1,"per_page":25,"total":0,"total_pages":0},"health":{"status":"ok","ready":true,"scheduled":1,"inflight":0,"scheduler_id":"sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890"}}
```

## Error Codes

| Status | Code | Description |
|--------|------|-------------|
| 400 | validation_error | Invalid pagination or filters |
| 404 | not_found | Scheduler not found |
| 500 | internal_error | Internal server error |
