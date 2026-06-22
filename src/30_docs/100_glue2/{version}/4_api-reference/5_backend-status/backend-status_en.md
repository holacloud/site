# Backend Status

Returns the health status of all backend services.

## Description

This endpoint performs health checks against each registered backend service and reports their current status. It is used by monitoring systems and the Console dashboard.

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
{
  "services": [
    {
      "name": "inceptiondb",
      "status": "healthy",
      "latency_ms": 5,
      "last_checked": "2026-06-21T10:00:00Z"
    },
    {
      "name": "lambda",
      "status": "healthy",
      "latency_ms": 12,
      "last_checked": "2026-06-21T10:00:00Z"
    },
    {
      "name": "files",
      "status": "degraded",
      "latency_ms": 350,
      "last_checked": "2026-06-21T10:00:00Z"
    },
    {
      "name": "kvnode",
      "status": "healthy",
      "latency_ms": 3,
      "last_checked": "2026-06-21T10:00:00Z"
    }
  ],
  "overall": "degraded"
}
```

## Error Codes

| Code | Description |
|------|-------------|
| 200 | Status information returned successfully |
