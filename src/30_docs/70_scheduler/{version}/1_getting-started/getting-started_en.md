# Getting Started

This guide walks you through creating a scheduler, enqueuing a delayed task, listing tasks, reserving and acknowledging a task, and checking scheduler health.

## Prerequisites

- A HolaCloud account with API credentials.
- [curl](https://curl.se/) installed on your machine.

## Authentication

Some endpoints require authentication via API key. Pass it using the `X-API-Key` header or `Authorization: Bearer <key>`.

## Step 1: Create a Scheduler

Create a new scheduler:

```bash
curl -X POST "https://api.hola.cloud/schedulers" \
  -H "X-API-Key: YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "display_name": "my-scheduler"
  }'
```

Expected response:

```json
{
  "id": "sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890",
  "display_name": "my-scheduler",
  "task_count": 0,
  "status": "active"
}
```

Save the returned `id` — you will use it in subsequent requests.

## Step 2: Enqueue a Task with Delay

Enqueue a task that becomes available in 60 seconds:

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

Expected response:

```json
{
  "id": "task-001",
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

## Step 3: List Tasks

View all tasks in the scheduler:

```bash
curl "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks"
```

Expected response:

```json
{
  "tasks": [
    {
      "id": "task-001",
      "state": "pending",
      "available_at": "2025-06-21T12:01:01Z",
      "labels": { "project": "onboarding", "priority": "high" }
    }
  ],
  "total": 1
}
```

## Step 4: Reserve and Acknowledge a Task

Once the delay has elapsed, reserve the task:

```bash
curl -X POST "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks/reserve" \
  -H "X-API-Key: YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "worktime": 30
  }'
```

Expected response:

```json
{
  "id": "task-001",
  "payload": { "type": "send_email", "to": "user@example.com", "template": "welcome" },
  "lease_expires_at": "2025-06-21T12:02:01Z"
}
```

Process the task, then acknowledge (delete) it:

```bash
curl -X DELETE "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks/task-001" \
  -H "X-API-Key: YOUR_API_KEY"
```

Expected response: HTTP `204 No Content`.

## Step 5: Check Scheduler Health

Verify the scheduler is healthy:

```bash
curl "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/health"
```

Expected response:

```json
{
  "status": "ok",
  "uptime_seconds": 12345
}
```

## Next Steps

- Learn how to manage schedulers, including updates and monitoring, in [Managing Schedulers](../2_managing-schedulers/managing-schedulers_en.md).
- Dive deeper into the task lifecycle — including lease extension, reschedule, and SSE streaming — in [Task Lifecycle](../3_task-lifecycle/task-lifecycle_en.md).
