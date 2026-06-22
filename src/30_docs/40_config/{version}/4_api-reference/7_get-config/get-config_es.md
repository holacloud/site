# Obtener Config

Obtener una configuración específica por su ID. Este endpoint es de acceso público (API de administración).

## Autenticación

No requiere autenticación. Este es un endpoint público.

## Parámetros de Ruta

| Parámetro | Tipo | Descripción |
|-----------|------|-------------|
| `id` | string | El ID de la configuración (ej. `cfg_abc123`) |

## Solicitud

```bash
curl "https://api.hola.cloud/v0/configs/cfg_abc123"
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
| 404 | Not Found | La configuración especificada no existe |
| 500 | Internal Server Error | Ocurrió un error inesperado |
