# Datos de Console

Devuelve datos de Console para un repositorio.

## Request

```http
GET /api/console?repository=my-project/my-app
```

El query parameter `repository` es obligatorio.

## Ejemplo

```bash
curl "https://api.hola.cloud/api/console?repository=my-project/my-app"
```

## Respuesta

```json
{
  "repository": "my-project/my-app",
  "references": ["latest", "v1"],
  "env": [
    {"key": "LOG_LEVEL", "desired_value": "info"}
  ],
  "volumes": [
    {"name": "my-data", "target": "/data", "mode": "rw"}
  ]
}
```
