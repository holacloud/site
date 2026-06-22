# Extend Lease

Extends the lease on a currently reserved task.

## Authentication

Requires authentication. Pass your API key via `X-API-Key` or `Authorization: Bearer` header.

## Path Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| id | string | The unique identifier of the scheduler |
| task | string | The unique identifier of the task |

## Request Body

| Field | Type | Description |
|-------|------|-------------|
| extension | string | Additional lease time as a Go duration string, such as `30s` or `2m` |

```json
{
  "extension": "30s"
}
```

## Example Request

```bash
curl -X POST "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks/task-x1y2z3/extend" \
  -H "X-API-Key: YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "extension": "30s"
  }'
```

## Example Response

```http
HTTP/1.1 200 OK
Content-Type: application/json
```

```json
{
  "lease_expires_at": "2025-06-21T12:02:31Z"
}
```

## Error Codes

| Status | Code | Description |
|--------|------|-------------|
| 400 | invalid_json | Invalid JSON payload |
| 400 | validation_error | Missing or invalid extension value |
| 400 | invalid_lease | Extension must be positive |
| 401 | unauthorized | Missing or invalid API key |
| 404 | task_not_found | Task not found |
| 500 | internal_error | Internal server error |
