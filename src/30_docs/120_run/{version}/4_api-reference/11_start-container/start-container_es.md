# Iniciar repositorio

Inicia un repositorio con una referencia o digest subido.

```json
{
  "repository": "my-project/my-app",
  "reference": "latest"
}
```

```bash
curl -X POST "https://api.hola.cloud/api/console/start" \
  -H "Content-Type: application/json" \
  -d '{"repository": "my-project/my-app", "reference": "latest"}'
```

Forma con digest:

```json
{
  "repository": "my-project/my-app",
  "digest": "sha256:a1b2c3d4..."
}
```
