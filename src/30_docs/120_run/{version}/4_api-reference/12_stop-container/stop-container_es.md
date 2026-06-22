# Detener repositorio

Detiene un repositorio.

```json
{
  "repository": "my-project/my-app"
}
```

```bash
curl -X POST "https://api.hola.cloud/api/console/stop" \
  -H "Content-Type: application/json" \
  -d '{"repository": "my-project/my-app"}'
```
