# Update Scheduler

Replaces the configuration of an existing scheduler.

## Authentication

Requires authentication. Pass your API key via `X-API-Key` or `Authorization: Bearer` header.

## Path Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| id | string | The unique identifier of the scheduler |

## Request Body

```json
{
  "display_name": "renamed-scheduler"
}
```

## Example Request

```bash
curl -X PUT "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890" \
  -H "X-API-Key: YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "display_name": "renamed-scheduler"
  }'
```

## Example Response

```http
HTTP/1.1 200 OK
Content-Type: application/json
```

```json
{
  "id": "sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890",
  "display_name": "renamed-scheduler",
  "task_count": 5,
  "status": "active",
  "created_at": "2025-06-20T10:00:00Z",
  "updated_at": "2025-06-21T09:00:00Z"
}
```

## Error Codes

| Status | Code | Description |
|--------|------|-------------|
| 400 | invalid_request | Invalid request body |
| 401 | unauthorized | Missing or invalid API key |
| 404 | not_found | Scheduler not found |
| 500 | internal_error | Internal server error |
