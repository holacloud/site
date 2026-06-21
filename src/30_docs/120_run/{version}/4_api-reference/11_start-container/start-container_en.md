# POST /api/console/start

Starts a new container from a specified image.

## Description

Creates and starts a container from the given image with optional environment variables, volumes, and port mappings.

## Authentication

None. This endpoint is public.

## Request Body

```json
{
  "image": "my-project/my-app:latest",
  "env": {
    "LOG_LEVEL": "debug",
    "DATABASE_URL": "postgres://user:pass@host:5432/db"
  },
  "volumes": [
    {
      "container_path": "/data",
      "size_gb": 10
    }
  ],
  "ports": {
    "80/tcp": 8080
  }
}
```

## Example

```bash
curl -X POST "https://api.hola.cloud/api/console/start" \
  -H "Content-Type: application/json" \
  -d '{
    "image": "my-project/my-app:latest",
    "env": {"LOG_LEVEL": "debug"},
    "ports": {"80/tcp": 8080}
  }'
```

## Response

```json
{
  "container_id": "abc123def456",
  "image": "my-project/my-app:latest",
  "status": "running",
  "ports": {"80/tcp": 8080},
  "started_at": "2026-06-21T10:00:00Z"
}
```

## Error Codes

| Code | Description |
|------|-------------|
| 200 | Container started successfully |
| 400 | Invalid request body |
| 404 | Image not found |
