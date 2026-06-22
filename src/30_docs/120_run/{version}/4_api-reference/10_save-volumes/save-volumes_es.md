# Guardar volúmenes

Guarda configuración de volúmenes para un repositorio.

```json
{
  "repository": "my-project/my-app",
  "volumes": [
    {"name": "my-data", "target": "/data", "mode": "rw"},
    {"name": "config", "target": "/config", "mode": "ro"}
  ]
}
```

```bash
curl -X PUT "https://api.hola.cloud/api/console/volumes" \
  -H "Content-Type: application/json" \
  -d '{"repository": "my-project/my-app", "volumes": [{"name": "my-data", "target": "/data", "mode": "rw"}]}'
```
