# GET /schedulers/{id}/tasks/stream

Transmite eventos de tareas en tiempo real mediante Server-Sent Events (SSE).

## Autenticación

Este endpoint es público. No requiere autenticación.

## Parámetros de Ruta

| Parámetro | Tipo | Descripción |
|-----------|------|-------------|
| id | string | Identificador único del scheduler |

## Ejemplo de Solicitud

```bash
curl -N "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks/stream"
```

## Ejemplo de Respuesta

```http
HTTP/1.1 200 OK
Content-Type: text/event-stream
Cache-Control: no-cache
Connection: keep-alive
```

```
event: task_available
data: {"id":"tarea-x1y2z3","state":"pending","available_at":"2025-06-21T12:01:01Z"}

event: task_reserved
data: {"id":"tarea-x1y2z3","state":"inflight","lease_expires_at":"2025-06-21T12:02:01Z"}

event: task_completed
data: {"id":"tarea-x1y2z3","state":"completed"}

event: lease_expired
data: {"id":"tarea-x1y2z3","state":"pending"}
```

## Códigos de Error

| Estado | Código | Descripción |
|--------|--------|-------------|
| 404 | not_found | Scheduler no encontrado |
| 500 | internal_error | Error interno del servidor |
