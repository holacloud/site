# Modificar Config

Actualizar parcialmente una configuración. Solo se modifican los campos proporcionados en el cuerpo de la solicitud. Este endpoint es de acceso público (API de administración).

## Autenticación

No requiere autenticación. Este es un endpoint público.

## Parámetros de Ruta

| Parámetro | Tipo | Descripción |
|-----------|------|-------------|
| `id` | string | El ID de la configuración (ej. `cfg_abc123`) |

## Cuerpo de la Solicitud

```json
{
  "data": {
    "features": {
      "new-checkout": true
    }
  }
}
```

| Campo | Tipo | Requerido | Descripción |
|-------|------|-----------|-------------|
| `project` | string | no | Nuevo nombre del proyecto |
| `environment` | string | no | Nuevo nombre del entorno |
| `data` | object | no | Datos parciales de configuración a fusionar |

## Solicitud

```bash
curl -X PATCH "https://api.hola.cloud/v0/configs/cfg_abc123" \
  -H "Content-Type: application/json" \
  -d '{
    "data": {
      "features": {
        "new-checkout": true
      }
    }
  }'
```

## Respuesta

```http
HTTP/1.1 200 OK
Content-Type: application/json
```

```json
{
  "id": "cfg_abc123",
  "project": "my-app",
  "environment": "production",
  "data": {
    "database": {
      "host": "db.example.com",
      "port": 5432
    },
    "features": {
      "new-checkout": true
    }
  },
  "createdAt": "2026-06-21T10:00:00Z",
  "updatedAt": "2026-06-21T11:00:00Z"
}
```

## Códigos de Error

| Estado | Código | Descripción |
|--------|--------|-------------|
| 400 | Bad Request | Cuerpo de solicitud inválido |
| 404 | Not Found | La configuración especificada no existe |
| 409 | Conflict | La actualización entra en conflicto con una configuración existente |
| 500 | Internal Server Error | Ocurrió un error inesperado |
