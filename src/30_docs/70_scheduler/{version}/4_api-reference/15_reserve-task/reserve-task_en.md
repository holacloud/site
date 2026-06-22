# Reserve Task

Reserves one available task for processing. The task enters a lease period defined by `worktime`.

## Authentication

Requires authentication. Pass your API key via `X-API-Key` or `Authorization: Bearer` header.

## Path Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| id | string | The unique identifier of the scheduler |

## Request Body

| Field | Type | Description |
|-------|------|-------------|
| worktime | string | Lease duration as a Go duration string, such as `30s` or `2m` |

```json
{
  "worktime": "30s"
}
```

## Example Request

```bash
curl -X POST "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks/reserve" \
  -H "X-API-Key: YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "worktime": "30s"
  }'
```

## Example Response

```http
HTTP/1.1 200 OK
Content-Type: application/json
```

```json
{
  "id": "task-x1y2z3",
  "payload": {
    "type": "send_email",
    "to": "user@example.com",
    "template": "welcome"
  },
  "lease_expires_at": "2025-06-21T12:02:01Z",
  "labels": ["project:onboarding", "priority:high"]
}
```

If no task is available before the work window ends, the endpoint returns `204 No Content` with an empty body.

## Error Codes

| Status | Code | Description |
|--------|------|-------------|
| 400 | invalid_json | Invalid JSON payload |
| 400 | validation_error | Missing, invalid, or non-positive worktime |
| 401 | unauthorized | Missing or invalid API key |
| 500 | internal_error | Internal server error |
