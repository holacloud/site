# Referencia de la API de Files

URL base: `https://api.hola.cloud`

Todos los endpoints requieren autenticación mediante los encabezados `Api-Key` y `Api-Secret`.

## Endpoints

| Método | Ruta | Descripción |
|--------|------|-------------|
| GET | `/v1/buckets` | Listar todos los buckets |
| POST | `/v1/buckets` | Crear un nuevo bucket |
| GET | `/v1/buckets/{id}` | Obtener detalles de un bucket |
| PATCH | `/v1/buckets/{id}` | Modificar metadatos de un bucket |
| DELETE | `/v1/buckets/{id}` | Eliminar un bucket |
| GET | `/v1/buckets/{id}/list/*` | Listar archivos en un bucket |
| PUT | `/v1/buckets/{id}/files/*` | Subir un archivo |
| GET | `/v1/buckets/{id}/files/*` | Descargar un archivo |
| DELETE | `/v1/buckets/{id}/files/*` | Eliminar un archivo |
| HEAD | `/v1/buckets/{id}/files/*` | Obtener metadatos de un archivo |

## Errores Comunes

| Estado | Código | Descripción |
|--------|--------|-------------|
| 400 | Bad Request | Cuerpo o parámetros de solicitud inválidos |
| 401 | Unauthorized | Credenciales API faltantes o inválidas |
| 403 | Forbidden | Permisos insuficientes |
| 404 | Not Found | El recurso especificado no existe |
| 409 | Conflict | El recurso ya existe o la operación entra en conflicto |
| 413 | Payload Too Large | El archivo excede el límite máximo de tamaño (5 GB) |
| 429 | Too Many Requests | Límite de velocidad excedido |
| 500 | Internal Server Error | Ocurrió un error inesperado |
