# Managing Schedulers

This document covers the scheduler lifecycle: creation with metadata, listing with search and filter, updating configuration, health monitoring, and deletion.

## Creating a Scheduler

Create a new scheduler with a display name:

```bash
curl -X POST "https://api.hola.cloud/schedulers" \
  -H "X-API-Key: YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "display_name": "email-scheduler"
  }'
```

```http
POST /schedulers HTTP/1.1
Host: api.hola.cloud
X-API-Key: YOUR_API_KEY
Content-Type: application/json

{
  "display_name": "email-scheduler"
}
```

Expected response:

```json
{
  "id": "sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890",
  "display_name": "email-scheduler",
  "task_count": 0,
  "status": "active"
}
```

## Listing Schedulers

List all schedulers on your account (public):

```bash
curl "https://api.hola.cloud/schedulers"
```

```http
GET /schedulers HTTP/1.1
Host: api.hola.cloud
```

Expected response:

```json
[
  {
    "id": "sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890",
    "display_name": "email-scheduler",
    "task_count": 5,
    "status": "active"
  },
  {
    "id": "sched-b2c3d4e5-f6a7-8901-bcde-f12345678901",
    "display_name": "report-generator",
    "task_count": 0,
    "status": "active"
  }
]
```

## Getting Scheduler Details

Retrieve a single scheduler by its ID:

```bash
curl "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890"
```

```http
GET /schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890 HTTP/1.1
Host: api.hola.cloud
```

Expected response:

```json
{
  "id": "sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890",
  "display_name": "email-scheduler",
  "task_count": 5,
  "status": "active"
}
```

## Updating a Scheduler

Update the display name or other metadata:

```bash
curl -X PATCH "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890" \
  -H "X-API-Key: YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "display_name": "email-scheduler-v2"
  }'
```

```http
PATCH /schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890 HTTP/1.1
Host: api.hola.cloud
X-API-Key: YOUR_API_KEY
Content-Type: application/json

{
  "display_name": "email-scheduler-v2"
}
```

Expected response:

```json
{
  "id": "sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890",
  "display_name": "email-scheduler-v2",
  "task_count": 5,
  "status": "active"
}
```

PUT can be used interchangeably with PATCH.

## Health Monitoring

Check the health status of a scheduler at any time:

```bash
curl "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/health"
```

```http
GET /schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/health HTTP/1.1
Host: api.hola.cloud
```

Expected response:

```json
{
  "status": "ok",
  "uptime_seconds": 123456
}
```

A non-`ok` status indicates the scheduler is unavailable or in a degraded state.

## Deleting a Scheduler

Permanently delete a scheduler and all its tasks:

```bash
curl -X DELETE "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890" \
  -H "X-API-Key: YOUR_API_KEY"
```

```http
DELETE /schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890 HTTP/1.1
Host: api.hola.cloud
X-API-Key: YOUR_API_KEY
```

Expected response: HTTP `204 No Content`.

Once deleted, the scheduler ID is no longer valid and all pending tasks are discarded.
