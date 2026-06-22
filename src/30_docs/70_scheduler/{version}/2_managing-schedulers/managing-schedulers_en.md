# Managing Schedulers

This document covers the scheduler lifecycle: creation with metadata, listing with search and filter, updating configuration, health monitoring, and deletion.

## Creating a Scheduler

Create a new scheduler with an id and display name:

```bash
curl -X POST "https://api.hola.cloud/schedulers" \
  -H "X-API-Key: YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "id": "sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890",
    "display_name": "email-scheduler"
  }'
```

```http
POST /schedulers HTTP/1.1
Host: api.hola.cloud
X-API-Key: YOUR_API_KEY
Content-Type: application/json

{
  "id": "sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890",
  "display_name": "email-scheduler"
}
```

Expected response:

```json
{
  "scheduler": {
    "id": "sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890",
    "ready": true,
    "scheduled": 0,
    "inflight": 0,
    "display_name": "email-scheduler"
  }
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
{
  "default_scheduler_id": "default",
  "schedulers": [
    {
      "id": "sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890",
      "ready": true,
      "scheduled": 5,
      "inflight": 0,
      "display_name": "email-scheduler"
    }
  ]
}
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
  "scheduler": {
    "id": "sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890",
    "ready": true,
    "scheduled": 5,
    "inflight": 0,
    "display_name": "email-scheduler"
  }
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
  "scheduler": {
    "id": "sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890",
    "ready": true,
    "scheduled": 5,
    "inflight": 0,
    "display_name": "email-scheduler-v2"
  }
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
  "ready": true,
  "scheduled": 5,
  "inflight": 0,
  "scheduler_id": "sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890"
}
```

The health response reports the scheduler readiness and current scheduled/inflight counts.

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
