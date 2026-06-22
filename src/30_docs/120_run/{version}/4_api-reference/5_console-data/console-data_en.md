# Console Data

Returns repositories, images, and running container state for the Console dashboard.

## Description

This endpoint aggregates all container-related data needed by the HolaCloud Console to display the current state of deployed services.

## Authentication

None. This endpoint is public.

## Request

No request body required.

## Example

```bash
curl -X GET "https://api.hola.cloud/api/console"
```

## Response

```json
{
  "repositories": [
    {
      "name": "my-project/my-app",
      "tags": ["latest", "v1.0.0", "v1.1.0"],
      "pull_count": 142
    }
  ],
  "images": [
    {
      "repository": "my-project/my-app",
      "tag": "latest",
      "digest": "sha256:a1b2c3d4...",
      "size_bytes": 72450000,
      "created": "2026-06-20T10:00:00Z"
    }
  ],
  "containers": [
    {
      "id": "abc123def456",
      "image": "my-project/my-app:latest",
      "status": "running",
      "ports": {"80/tcp": 8080},
      "started_at": "2026-06-21T08:00:00Z",
      "env": {
        "LOG_LEVEL": "info",
        "DATABASE_URL": "postgres://..."
      }
    }
  ]
}
```

## Error Codes

| Code | Description |
|------|-------------|
| 200 | Console data returned successfully |
