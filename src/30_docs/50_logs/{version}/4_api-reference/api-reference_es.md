
# Referencia de la API de InstantLogs

URL base: `https://api.hola.cloud`

## Autenticación

InstantLogs utiliza dos modos de autenticación según el tipo de endpoint.

### Autenticación de Gestión

Los endpoints de gestión (CRUD de loggers, administración de claves API) requieren dos encabezados:

| Encabezado | Descripción |
|------------|-------------|
| `Api-Key` | Identificador de tu clave API |
| `Api-Secret` | Tu clave secreta de API |

### Autenticación de Datos

Los endpoints de datos (ingesta, filtrado, eventos, estadísticas) se autentican usando el secreto del logger. Provéelo como:

| Encabezado | Descripción |
|------------|-------------|
| `X-Instantlogs-Event-Secret` | El valor secreto del logger |
| `Authorization: Bearer <secret>` | Token Bearer con el secreto del logger |

## Endpoints

| Método | Ruta | Descripción | Autenticación |
|--------|------|-------------|---------------|
| GET | `/v1/loggers` | Listar todos los loggers | Gestión |
| POST | `/v1/loggers` | Crear un nuevo logger | Gestión |
| GET | `/v1/loggers/{id}` | Obtener detalles de un logger | Gestión |
| DELETE | `/v1/loggers/{id}` | Eliminar un logger | Gestión |
| POST | `/v1/loggers/{id}/ingest` | Ingresar entradas de log | Datos |
| GET | `/v1/loggers/{id}/filter` | Transmitir y filtrar logs | Datos |
| GET | `/v1/loggers/{id}/events` | Transmitir eventos | Datos |
| GET | `/v1/loggers/{id}/stats` | Obtener estadísticas del logger | Datos |
| POST | `/v1/loggers/{id}/apiKeys` | Crear una clave API | Gestión |
| DELETE | `/v1/loggers/{id}/apiKeys/{key}` | Eliminar una clave API | Gestión |
| POST | `/v1/ingest/events` | Ingresar eventos enmarcados | Datos |

## Códigos de Error Comunes

| Código | Descripción |
|--------|-------------|
| 400 | Solicitud incorrecta — JSON malformado o parámetros inválidos |
| 401 | No autorizado — credenciales faltantes o inválidas |
| 403 | Prohibido — las credenciales no tienen acceso |
| 404 | No encontrado — el recurso solicitado no existe |
| 409 | Conflicto — el recurso ya existe |
| 429 | Demasiadas solicitudes — límite de tasa excedido |
| 500 | Error interno del servidor |
