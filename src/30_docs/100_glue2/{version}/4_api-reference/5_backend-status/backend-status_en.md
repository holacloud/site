# Backend Status

Returns the health status of all backend services.

## Description

This endpoint checks project hosts and returns an array with one status object per project host.

## Authentication

None. This endpoint is public.

## Request

No request body required.

## Example

```bash
curl -X GET "https://api.hola.cloud/v0/status"
```

## Response

```json
[
  {
    "id": "project-123",
    "name": "My Project",
    "host": "my-project.hola.cloud",
    "status": 200,
    "statusText": "200 OK"
  }
]
```

## Error Codes

| Code | Description |
|------|-------------|
| 200 | Status information returned successfully |
