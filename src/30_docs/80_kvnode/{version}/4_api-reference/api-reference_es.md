# Referencia de la API de KVNode

URL base: `https://api.hola.cloud`

| Método | Ruta | Descripción | Autenticación |
|--------|------|-------------|---------------|
| GET | /healthz | Verificación de salud | Pública |
| GET | /readyz | Verificación de disponibilidad | Pública |
| GET | /v1/status | Estado del nodo | Interna |
| GET | /v1/metrics | Métricas del nodo | Interna |
| GET | /v1/collections | Listar colecciones | Interna |
| POST | /v1/collections | Crear colección | Interna |
| DELETE | /v1/collections/{col} | Eliminar colección | Interna |
| GET | /v1/collections/{col}/keys | Listar claves | Interna |
| POST | /v1/collections/{col}/keys/* | Establecer clave | Interna |
| GET | /v1/collections/{col}/keys/* | Obtener clave | Interna |
| DELETE | /v1/collections/{col}/keys/* | Eliminar clave | Interna |

## Autenticación

Los endpoints de verificación de salud son públicos. Todos los demás endpoints requieren autenticación mediante uno de:

- Encabezado `X-Glue-Authentication`
- Encabezados `apikey` y `secret`

## Códigos de Error

| Estado | Código | Descripción |
|--------|--------|-------------|
| 400 | invalid_json | Cuerpo o parámetros de solicitud inválidos |
| 403 | forbidden | Missing authentication headers |
| 404 | not_found | Colección o clave no encontrada |
| 409 | conflict | La colección ya existe |
| 500 | internal_error | Error interno del servidor |
