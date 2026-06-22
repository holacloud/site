# Rollback Image

Rolls back a container to a previous image version.

## Description

Replaces the current image of a running or stopped container with a specified previous version and restarts it.

## Authentication

None. This endpoint is public.

## Request Body

```json
{
  "container_id": "abc123def456",
  "tag": "v1.0.0"
}
```

## Example

```bash
curl -X POST "https://api.hola.cloud/api/console/rollback" \
  -H "Content-Type: application/json" \
  -d '{
    "container_id": "abc123def456",
    "tag": "v1.0.0"
  }'
```

## Response

```json
{
  "container_id": "abc123def456",
  "previous_image": "my-project/my-app:latest",
  "current_image": "my-project/my-app:v1.0.0",
  "status": "running",
  "rolled_back_at": "2026-06-21T12:00:00Z"
}
```

## Error Codes

| Code | Description |
|------|-------------|
| 200 | Rollback completed successfully |
| 400 | Invalid request body |
| 404 | Container or image tag not found |
