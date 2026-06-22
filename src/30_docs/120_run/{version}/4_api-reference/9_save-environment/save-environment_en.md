# Save Environment

Saves environment variables for a container.

## Description

Updates the environment variable configuration for a specified container. The container must be restarted for changes to take effect.

## Authentication

None. This endpoint is public.

## Request Body

```json
{
  "container_id": "abc123def456",
  "env": {
    "LOG_LEVEL": "info",
    "DATABASE_URL": "postgres://user:pass@host:5432/db",
    "REDIS_URL": "redis://host:6379"
  }
}
```

## Example

```bash
curl -X PUT "https://api.hola.cloud/api/console/env" \
  -H "Content-Type: application/json" \
  -d '{
    "container_id": "abc123def456",
    "env": {
      "LOG_LEVEL": "info",
      "DATABASE_URL": "postgres://user:pass@host:5432/db"
    }
  }'
```

## Response

```json
{
  "container_id": "abc123def456",
  "env": {
    "LOG_LEVEL": "info",
    "DATABASE_URL": "postgres://user:pass@host:5432/db",
    "REDIS_URL": "redis://host:6379"
  },
  "updated": true
}
```

## Error Codes

| Code | Description |
|------|-------------|
| 200 | Environment variables saved successfully |
| 400 | Invalid request body |
| 404 | Container not found |
