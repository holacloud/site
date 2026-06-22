# Revertir Image

Revierte un contenedor a una versión de imagen anterior.

## Descripción

Reemplaza la imagen actual de un contenedor en ejecución o detenido con una versión anterior especificada y lo reinicia.

## Autenticación

Ninguna. Este endpoint es público.

## Cuerpo de la Solicitud

```json
{
  "container_id": "abc123def456",
  "tag": "v1.0.0"
}
```

## Ejemplo

```bash
curl -X POST "https://api.hola.cloud/api/console/rollback" \
  -H "Content-Type: application/json" \
  -d '{
    "container_id": "abc123def456",
    "tag": "v1.0.0"
  }'
```

## Respuesta

```json
{
  "container_id": "abc123def456",
  "previous_image": "my-project/my-app:latest",
  "current_image": "my-project/my-app:v1.0.0",
  "status": "running",
  "rolled_back_at": "2026-06-21T12:00:00Z"
}
```

## Códigos de Error

| Código | Descripción |
|--------|-------------|
| 200 | Reversión completada correctamente |
| 400 | Cuerpo de solicitud inválido |
| 404 | Contenedor o etiqueta de imagen no encontrada |
