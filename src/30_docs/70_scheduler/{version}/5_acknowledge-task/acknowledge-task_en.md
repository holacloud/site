# DELETE /schedulers/{id}/tasks/{task}

Acknowledges and removes a task from the scheduler. This should be called after a worker successfully processes a reserved task.

## Authentication

Requires authentication. Pass your API key via `X-API-Key` or `Authorization: Bearer` header.

## Path Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| id | string | The unique identifier of the scheduler |
| task | string | The unique identifier of the task |

## Example Request

```bash
curl -X DELETE "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks/task-x1y2z3" \
  -H "X-API-Key: YOUR_API_KEY"
```

## Example Response

```http
HTTP/1.1 204 No Content
```

## Error Codes

| Status | Code | Description |
|--------|------|-------------|
| 401 | unauthorized | Missing or invalid API key |
| 404 | not_found | Scheduler or task not found |
| 409 | conflict | Task is not currently reserved |
| 500 | internal_error | Internal server error |
