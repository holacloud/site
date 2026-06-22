# Tarea Transmitir

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
event: snapshot
data: {"scheduler_id":"sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890","generated_at":"2025-06-21T12:00:01Z","scheduled":[{"id":"task-x1y2z3","future":"2025-06-21T12:01:01Z","labels":["priority:high"]}],"inflight":[],"scheduled_meta":{"page":1,"per_page":25,"total":1,"total_pages":1},"inflight_meta":{"page":1,"per_page":25,"total":0,"total_pages":0},"health":{"status":"ok","ready":true,"scheduled":1,"inflight":0,"scheduler_id":"sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890"}}
```

## Códigos de Error

| Estado | Código | Descripción |
|--------|--------|-------------|
| 404 | not_found | Scheduler no encontrado |
| 500 | internal_error | Error interno del servidor |
