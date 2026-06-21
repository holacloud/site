# Referencia de la API de Config

URL base: `https://api.hola.cloud`

El servicio Config expone dos superficies de API:

- **API de administración** (`/v0/configs`) — acceso público para gestionar configuraciones
- **API de usuario** (`/v1/config`) — requiere los encabezados `Api-Key` y `Api-Secret` para acceso a la configuración en tiempo de ejecución

## Endpoints

| Método | Ruta | Descripción | Autenticación |
|--------|------|-------------|---------------|
| GET | `/v0/configs` | Listar todas las configuraciones (admin) | Pública |
| POST | `/v0/configs` | Crear una nueva configuración (admin) | Pública |
| GET | `/v0/configs/{id}` | Obtener una configuración por ID (admin) | Pública |
| DELETE | `/v0/configs/{id}` | Eliminar una configuración (admin) | Pública |
| PATCH | `/v0/configs/{id}` | Actualización parcial de una configuración (admin) | Pública |
| GET | `/v1/config` | Obtener la configuración del usuario actual | API Key |
| PATCH | `/v1/config` | Actualizar la configuración del usuario actual | API Key |

## Errores Comunes

| Estado | Código | Descripción |
|--------|--------|-------------|
| 400 | Bad Request | Cuerpo o parámetros de solicitud inválidos |
| 401 | Unauthorized | Credenciales API faltantes o inválidas |
| 404 | Not Found | El recurso especificado no existe |
| 409 | Conflict | El recurso ya existe o la operación entra en conflicto |
| 500 | Internal Server Error | Ocurrió un error inesperado |
