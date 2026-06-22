# Rollback de imagen

Revierte un repositorio a una referencia o digest subido.

## Body

```json
{
  "repository": "my-project/my-app",
  "reference": "v1.0.0"
}
```

También se puede usar `digest`:

```json
{
  "repository": "my-project/my-app",
  "digest": "sha256:a1b2c3d4..."
}
```

## Ejemplo

```bash
curl -X POST "https://api.hola.cloud/api/console/rollback" \
  -H "Content-Type: application/json" \
  -d '{"repository": "my-project/my-app", "reference": "v1.0.0"}'
```
