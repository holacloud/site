# List Schedulers

Returns a list of all schedulers.

## Authentication

This endpoint is public. No authentication required.

## Query Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| search | string | Filter schedulers by display name (partial match) |
| q | string | General search query |

## Example Request

```bash
curl "https://api.hola.cloud/schedulers?search=my"
```

## Example Response

```http
HTTP/1.1 200 OK
Content-Type: application/json
```

```json
{
  "schedulers": [
    {
      "id": "sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890",
      "display_name": "my-scheduler",
      "task_count": 5,
      "status": "active",
      "created_at": "2025-06-20T10:00:00Z",
      "updated_at": "2025-06-21T08:30:00Z"
    }
  ],
  "total": 1
}
```

## Error Codes

| Status | Code | Description |
|--------|------|-------------|
| 500 | internal_error | Internal server error |
