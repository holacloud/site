# Save Volumes

Saves volume configuration for a container.

## Description

Updates the persistent volume configuration for a specified container. Volumes are backed by HolaCloud's distributed storage and survive container restarts.

## Authentication

None. This endpoint is public.

## Request Body

```json
{
  "container_id": "abc123def456",
  "volumes": [
    {
      "container_path": "/data",
      "size_gb": 20,
      "mount_path": "/mnt/data"
    },
    {
      "container_path": "/config",
      "size_gb": 5,
      "mount_path": "/mnt/config"
    }
  ]
}
```

## Example

```bash
curl -X PUT "https://api.hola.cloud/api/console/volumes" \
  -H "Content-Type: application/json" \
  -d '{
    "container_id": "abc123def456",
    "volumes": [
      {"container_path": "/data", "size_gb": 20}
    ]
  }'
```

## Response

```json
{
  "container_id": "abc123def456",
  "volumes": [
    {
      "volume_id": "vol-001",
      "container_path": "/data",
      "size_gb": 20,
      "mount_path": "/mnt/data",
      "status": "active"
    }
  ],
  "updated": true
}
```

## Error Codes

| Code | Description |
|------|-------------|
| 200 | Volume configuration saved successfully |
| 400 | Invalid request body |
| 404 | Container not found |
