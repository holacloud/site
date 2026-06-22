# Crear Config

Crear una nueva configuración. Este endpoint es de acceso público (API de administración).

## Autenticación

No requiere autenticación. Este es un endpoint público.

## Cuerpo de la Solicitud

```json
{
  "project": "my-app",
  "environment": "production",
  "data": {
    "database": {
      "host": "db.example.com",
      "port": 5432
    },
    "features": {
      "new-checkout": false
    }
  }
}
```

| Campo | Tipo | Requerido | Descripción |
|-------|------|-----------|-------------|
| `project` | string | sí | Identificador del nombre del proyecto |
| `environment` | string | sí | Nombre del entorno (ej. producción, staging, desarrollo) |
| `data` | object | sí | Datos de configuración como objeto JSON |

## Solicitud

```bash
curl -X POST "https://api.hola.cloud/v0/configs" \
  -H "Content-Type: application/json" \
  -d '{
    "project": "my-app",
    "environment": "production",
    "data": {
      "database": {
        "host": "db.example.com",
        "port": 5432
      },
      "features": {
        "new-checkout": false
      }
    }
  }'
```

## Respuesta

```http
HTTP/1.1 201 Created
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
      "new-checkout": false
    }
  },
  "createdAt": "2026-06-21T10:00:00Z",
  "updatedAt": "2026-06-21T10:00:00Z"
}
```

## Códigos de Error

| Estado | Código | Descripción |
|--------|--------|-------------|
| 400 | Bad Request | Faltan campos requeridos o JSON inválido |
| 409 | Conflict | Ya existe una configuración con el mismo proyecto y entorno |
| 500 | Internal Server Error | Ocurrió un error inesperado |
