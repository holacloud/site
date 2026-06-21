# GET /schedulers/{id}

Returns details of a specific scheduler.

## Authentication

This endpoint is public. No authentication required.

## Path Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| id | string | The unique identifier of the scheduler |

## Example Request

```bash
curl "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890"
```

## Example Response

```http
HTTP/1.1 200 OK
Content-Type: application/json
```

```json
{
  "id": "sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890",
  "display_name": "my-scheduler",
  "task_count": 5,
  "status": "active",
  "created_at": "2025-06-20T10:00:00Z",
  "updated_at": "2025-06-21T08:30:00Z"
}
```

## Error Codes

| Status | Code | Description |
|--------|------|-------------|
| 404 | not_found | Scheduler not found |
| 500 | internal_error | Internal server error |
