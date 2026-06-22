# Save Volumes

Saves volume configuration for a repository.

## Request Body

```json
{
  "repository": "my-project/my-app",
  "volumes": [
    {"name": "my-data", "target": "/data", "mode": "rw"},
    {"name": "config", "target": "/config", "mode": "ro"}
  ]
}
```

## Example

```bash
curl -X PUT "https://api.hola.cloud/api/console/volumes" \
  -H "Content-Type: application/json" \
  -d '{
    "repository": "my-project/my-app",
    "volumes": [
      {"name": "my-data", "target": "/data", "mode": "rw"}
    ]
  }'
```
