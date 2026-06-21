# Task Lifecycle

This document details the complete lifecycle of a task: enqueuing, listing, reserving, extending leases, acknowledging, rescheduling, and real-time streaming via SSE.

## Task States

A task moves through the following states:

- **pending** — Waiting for its scheduled time. Not yet visible to workers.
- **available** — Scheduled time has elapsed; ready for reservation.
- **reserved** — Held by a worker under a lease.
- **completed** — Acknowledged and removed by the worker.

If a lease expires without acknowledgement, the task returns to **available**.

## Enqueuing a Task

Create a new task with a payload, future time (or delay), and optional labels:

```bash
curl -X POST "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks" \
  -H "X-API-Key: YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "payload": {
      "action": "generate_report",
      "report_id": "rpt-5678"
    },
    "future": "2025-06-22T00:00:00Z",
    "labels": {
      "team": "analytics",
      "env": "production"
    }
  }'
```

```http
POST /schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks HTTP/1.1
Host: api.hola.cloud
X-API-Key: YOUR_API_KEY
Content-Type: application/json

{
  "payload": {
    "action": "generate_report",
    "report_id": "rpt-5678"
  },
  "future": "2025-06-22T00:00:00Z",
  "labels": {
    "team": "analytics",
    "env": "production"
  }
}
```

Alternatively, use `delay` (seconds from now) instead of `future`:

```json
{
  "payload": { "action": "send_reminder" },
  "delay": 3600,
  "labels": { "type": "email" }
}
```

Expected response:

```json
{
  "id": "task-001",
  "scheduler_id": "sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890",
  "payload": { "action": "generate_report", "report_id": "rpt-5678" },
  "state": "pending",
  "available_at": "2025-06-22T00:00:00Z",
  "labels": { "team": "analytics", "env": "production" }
}
```

## Listing Tasks with Pagination

List tasks with optional pagination and label filtering:

```bash
curl "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks?offset=0&limit=20"
```

```http
GET /schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks?offset=0&limit=20 HTTP/1.1
Host: api.hola.cloud
```

Expected response:

```json
{
  "tasks": [
    {
      "id": "task-001",
      "state": "pending",
      "available_at": "2025-06-22T00:00:00Z",
      "labels": { "team": "analytics" }
    },
    {
      "id": "task-002",
      "state": "available",
      "available_at": "2025-06-21T12:00:00Z",
      "labels": { "team": "analytics" }
    }
  ],
  "total": 2,
  "offset": 0,
  "limit": 20
}
```

## Reserving a Task

A worker reserves an available task by specifying a `worktime` — the lease duration in seconds:

```bash
curl -X POST "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks/reserve" \
  -H "X-API-Key: YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "worktime": 60
  }'
```

```http
POST /schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks/reserve HTTP/1.1
Host: api.hola.cloud
X-API-Key: YOUR_API_KEY
Content-Type: application/json

{
  "worktime": 60
}
```

Expected response:

```json
{
  "id": "task-002",
  "payload": { "action": "send_reminder" },
  "lease_expires_at": "2025-06-21T12:01:00Z"
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
    "worktime": 30
  }'
```

```http
POST /schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks/task-002/extend HTTP/1.1
Host: api.hola.cloud
X-API-Key: YOUR_API_KEY
Content-Type: application/json

{
  "worktime": 30
}
```

Expected response:

```json
{
  "id": "task-002",
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

If processing fails, reschedule the task with a new delay or future time:

```bash
curl -X POST "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks/task-002/reschedule" \
  -H "X-API-Key: YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "delay": 300
  }'
```

```http
POST /schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks/task-002/reschedule HTTP/1.1
Host: api.hola.cloud
X-API-Key: YOUR_API_KEY
Content-Type: application/json

{
  "delay": 300
}
```

Expected response:

```json
{
  "id": "task-002",
  "state": "pending",
  "available_at": "2025-06-21T12:06:00Z"
}
```

This is ideal for implementing retry logic with exponential backoff.

## SSE Stream for Real-Time Updates

Subscribe to real-time task events via Server-Sent Events:

```bash
curl "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks/stream"
```

The stream emits events in SSE format:

```
event: task_available
data: {"id":"task-003","payload":{"action":"process_image"},"labels":{}}

event: task_completed
data: {"id":"task-003"}

event: task_expired
data: {"id":"task-003"}
```

| Event | Description |
|-------|-------------|
| `task_available` | A task has become available for reservation |
| `task_completed` | A task was acknowledged by a worker |
| `task_expired` | A lease expired and the task is available again |

SSE is useful for building responsive worker pools that react to tasks immediately without polling.
