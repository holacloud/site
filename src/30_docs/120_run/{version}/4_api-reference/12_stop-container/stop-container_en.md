# Stop Container

Stops a running container.

## Description

Gracefully stops a container by its ID. The container is given a grace period before being forcibly terminated.

## Authentication

None. This endpoint is public.

## Request Body

```json
{
  "container_id": "abc123def456"
}
```

## Example

```bash
curl -X POST "https://api.hola.cloud/api/console/stop" \
  -H "Content-Type: application/json" \
  -d '{"container_id": "abc123def456"}'
```

## Response

```json
{
  "container_id": "abc123def456",
  "status": "stopped",
  "stopped_at": "2026-06-21T11:00:00Z"
}
```

## Error Codes

| Code | Description |
|------|-------------|
| 200 | Container stopped successfully |
| 400 | Invalid request body |
| 404 | Container not found |
