# Scheduler Health

Returns the health status of a scheduler.

## Authentication

This endpoint is public. No authentication required.

## Path Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| id | string | The unique identifier of the scheduler |

## Example Request

```bash
curl "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/health"
```

## Example Response

```http
HTTP/1.1 200 OK
Content-Type: application/json
```

```json
{
  "status": "ok",
  "ready": true,
  "scheduled": 3,
  "inflight": 1,
  "scheduler_id": "sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890"
}
```

## Error Codes

| Status | Code | Description |
|--------|------|-------------|
| 404 | not_found | Scheduler not found |
| 500 | internal_error | Internal server error |
