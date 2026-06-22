# Task Lifecycle

This document details the complete lifecycle of a task: enqueuing, listing, reserving, extending leases, acknowledging, rescheduling, and SSE snapshots.

## Task States

A task moves through the following states:

- **pending** — Waiting for its scheduled time. Not yet visible to workers.
- **available** — Scheduled time has elapsed; ready for reservation.
- **reserved** — Held by a worker under a lease.
- **completed** — Acknowledged and removed by the worker.

If a lease expires without acknowledgement, the task returns to **available**.

## Enqueuing a Task

Create a new task with an id, payload, future time (or delay), and optional labels:

```bash
curl -X POST "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks" \
  -H "X-API-Key: YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "id": "task-001",
    "payload": {
      "action": "generate_report",
      "report_id": "rpt-5678"
    },
    "future": "2025-06-22T00:00:00Z",
    "labels": ["team:analytics", "env:production"]
  }'
```

```http
POST /schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks HTTP/1.1
Host: api.hola.cloud
X-API-Key: YOUR_API_KEY
Content-Type: application/json

{
  "id": "task-001",
  "payload": {
    "action": "generate_report",
    "report_id": "rpt-5678"
  },
  "future": "2025-06-22T00:00:00Z",
  "labels": ["team:analytics", "env:production"]
}
```

Alternatively, use `delay` as a Go duration string instead of `future`:

```json
{
  "payload": { "action": "send_reminder" },
  "delay": "1h",
  "labels": ["type:email"]
}
```

Expected response: HTTP `202 Accepted` with an empty body.

## Listing Tasks with Pagination

List tasks with optional pagination and label filtering:

```bash
curl "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks?scheduled_page=1&scheduled_per_page=20"
```

```http
GET /schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks?scheduled_page=1&scheduled_per_page=20 HTTP/1.1
Host: api.hola.cloud
```

Expected response:

```json
{
  "scheduled": [
    {
      "id": "task-001",
      "future": "2025-06-22T00:00:00Z",
      "labels": ["team:analytics"]
    }
  ],
  "inflight": [
    {
      "id": "task-002",
      "lease_expires_at": "2025-06-21T12:01:00Z",
      "labels": ["team:analytics"]
    }
  ],
  "scheduled_meta": { "page": 1, "per_page": 20, "total": 1, "total_pages": 1 },
  "inflight_meta": { "page": 1, "per_page": 25, "total": 1, "total_pages": 1 }
}
```

## Reserving a Task

A worker reserves an available task by specifying `worktime` as a Go duration string:

```bash
curl -X POST "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks/reserve" \
  -H "X-API-Key: YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "worktime": "60s"
  }'
```

```http
POST /schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks/reserve HTTP/1.1
Host: api.hola.cloud
X-API-Key: YOUR_API_KEY
Content-Type: application/json

{
  "worktime": "60s"
}
```

Expected response:

```json
{
  "id": "task-002",
  "payload": { "action": "send_reminder" },
  "lease_expires_at": "2025-06-21T12:01:00Z",
  "labels": ["type:email"]
}
```

If no task is available, the endpoint returns HTTP `204 No Content`.

## Extending a Lease

If the worker needs more time, extend the lease before it expires:

```bash
curl -X POST "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks/task-002/extend" \
  -H "X-API-Key: YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "extension": "30s"
  }'
```

```http
POST /schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks/task-002/extend HTTP/1.1
Host: api.hola.cloud
X-API-Key: YOUR_API_KEY
Content-Type: application/json

{
  "extension": "30s"
}
```

Expected response:

```json
{
  "lease_expires_at": "2025-06-21T12:01:30Z"
}
```

## Acknowledging (Deleting) a Task

After successful processing, acknowledge the task to remove it:

```bash
curl -X DELETE "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks/task-002" \
  -H "X-API-Key: YOUR_API_KEY"
```

```http
DELETE /schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks/task-002 HTTP/1.1
Host: api.hola.cloud
X-API-Key: YOUR_API_KEY
```

Expected response: HTTP `204 No Content`.

## Rescheduling a Task

Reschedule the task with a new delay or future time:

```bash
curl -X POST "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks/task-002/reschedule" \
  -H "X-API-Key: YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "delay": "5m"
  }'
```

```http
POST /schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks/task-002/reschedule HTTP/1.1
Host: api.hola.cloud
X-API-Key: YOUR_API_KEY
Content-Type: application/json

{
  "delay": "5m"
}
```

Expected response:

```json
{
  "id": "task-002",
  "future": "2025-06-21T12:06:00Z"
}
```

## SSE Stream for Snapshots

Subscribe to task snapshots via Server-Sent Events:

```bash
curl "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks/stream"
```

The stream emits events in SSE format:

```
event: snapshot
data: {"scheduler_id":"sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890","generated_at":"2025-06-21T12:00:01Z","scheduled":[{"id":"task-003","future":"2025-06-21T12:01:00Z","labels":["type:image"]}],"inflight":[],"scheduled_meta":{"page":1,"per_page":25,"total":1,"total_pages":1},"inflight_meta":{"page":1,"per_page":25,"total":0,"total_pages":0},"health":{"status":"ok","ready":true,"scheduled":1,"inflight":0,"scheduler_id":"sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890"}}
```

| Event | Description |
|-------|-------------|
| `snapshot` | Current scheduled and inflight task lists |

SSE is useful for monitoring current queue state without polling.
