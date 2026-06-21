# Referencia de la API de Scheduler

URL base: `https://api.hola.cloud`

| Método | Ruta | Descripción | Autenticación |
|--------|------|-------------|---------------|
| GET | /schedulers | Listar schedulers | Pública |
| POST | /schedulers | Crear scheduler | Privada |
| GET | /schedulers/{id} | Obtener scheduler | Pública |
| PUT | /schedulers/{id} | Actualizar scheduler | Privada |
| PATCH | /schedulers/{id} | Actualización parcial | Privada |
| DELETE | /schedulers/{id} | Eliminar scheduler | Privada |
| GET | /schedulers/{id}/health | Verificar estado | Pública |
| POST | /schedulers/{id}/tasks | Encolar tarea | Privada |
| GET | /schedulers/{id}/tasks | Listar tareas | Pública |
| GET | /schedulers/{id}/tasks/stream | Stream SSE de tareas | Pública |
| POST | /schedulers/{id}/tasks/reserve | Reservar tarea | Privada |
| DELETE | /schedulers/{id}/tasks/{task} | Confirmar tarea | Privada |
| POST | /schedulers/{id}/tasks/{task}/extend | Extender concesión | Privada |
| POST | /schedulers/{id}/tasks/{task}/reschedule | Reprogramar tarea | Privada |

## Autenticación

Los endpoints privados requieren una clave API. Pásala usando:

- Encabezado `X-API-Key`: `X-API-Key: TU_API_KEY`
- Encabezado `Authorization`: `Authorization: Bearer TU_API_KEY`

Los endpoints de solo lectura (GET) son públicos y no requieren autenticación.

## Códigos de Error

| Estado | Código | Descripción |
|--------|--------|-------------|
| 400 | invalid_request | Cuerpo o parámetros de solicitud inválidos |
| 401 | unauthorized | Clave API faltante o inválida |
| 404 | not_found | Scheduler o tarea no encontrados |
| 409 | conflict | Tarea ya reservada o conflicto de concesión |
| 500 | internal_error | Error interno del servidor |
