# GET /api/console

Devuelve repositorios, imágenes y estado de contenedores en ejecución para el panel de la Consola.

## Descripción

Este endpoint agrega todos los datos relacionados con contenedores que necesita la Consola de HolaCloud para mostrar el estado actual de los servicios desplegados.

## Autenticación

Ninguna. Este endpoint es público.

## Solicitud

No se requiere cuerpo en la solicitud.

## Ejemplo

```bash
curl -X GET "https://api.hola.cloud/api/console"
```

## Respuesta

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

## Códigos de Error

| Código | Descripción |
|--------|-------------|
| 200 | Datos de la consola devueltos correctamente |
