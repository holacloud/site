# Save Environment

Saves environment variables for a repository.

## Request Body

```json
{
  "repository": "my-project/my-app",
  "env": [
    {"key": "LOG_LEVEL", "desired_value": "info"},
    {"key": "DATABASE_URL", "desired_value": "postgres://user:pass@host:5432/db"}
  ]
}
```

## Example

```bash
curl -X PUT "https://api.hola.cloud/api/console/env" \
  -H "Content-Type: application/json" \
  -d '{
    "repository": "my-project/my-app",
    "env": [
      {"key": "LOG_LEVEL", "desired_value": "info"},
      {"key": "DATABASE_URL", "desired_value": "postgres://user:pass@host:5432/db"}
    ]
  }'
```
