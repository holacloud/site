# Create Scheduler

Creates a new scheduler.

## Authentication

Requires authentication. Pass your API key via `X-API-Key` or `Authorization: Bearer` header.

## Request Body

```json
{
  "id": "sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890",
  "display_name": "my-scheduler"
}
```

## Example Request

```bash
curl -X POST "https://api.hola.cloud/schedulers" \
  -H "X-API-Key: YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "id": "sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890",
    "display_name": "my-scheduler"
  }'
```

## Example Response

```http
HTTP/1.1 201 Created
Content-Type: application/json
```

```json
{
  "scheduler": {
    "id": "sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890",
    "ready": true,
    "scheduled": 0,
    "inflight": 0,
    "display_name": "my-scheduler",
    "created_at": "2025-06-21T10:00:00Z",
    "updated_at": "2025-06-21T10:00:00Z"
  }
}
```

## Error Codes

| Status | Code | Description |
|--------|------|-------------|
| 400 | invalid_json | Invalid JSON payload |
| 400 | validation_error | Missing or invalid id, display_name, or description |
| 401 | unauthorized | Missing or invalid API key |
| 409 | conflict | Scheduler already exists |
| 500 | internal_error | Internal server error |
