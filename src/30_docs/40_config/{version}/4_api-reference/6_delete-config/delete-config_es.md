# Eliminar Config

Eliminar una configuración por su ID. Este endpoint es de acceso público (API de administración).

## Autenticación

No requiere autenticación. Este es un endpoint público.

## Parámetros de Ruta

| Parámetro | Tipo | Descripción |
|-----------|------|-------------|
| `id` | string | El ID de la configuración (ej. `cfg_abc123`) |

## Solicitud

```bash
curl -X DELETE "https://api.hola.cloud/v0/configs/cfg_abc123"
```

## Respuesta

```http
HTTP/1.1 204 No Content
```

## Códigos de Error

| Estado | Código | Descripción |
|--------|--------|-------------|
| 404 | Not Found | La configuración especificada no existe |
| 500 | Internal Server Error | Ocurrió un error inesperado |
