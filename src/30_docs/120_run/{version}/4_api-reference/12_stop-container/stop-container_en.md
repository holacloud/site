# Stop Repository

Stops a repository.

## Request Body

```json
{
  "repository": "my-project/my-app"
}
```

## Example

```bash
curl -X POST "https://api.hola.cloud/api/console/stop" \
  -H "Content-Type: application/json" \
  -d '{"repository": "my-project/my-app"}'
```
