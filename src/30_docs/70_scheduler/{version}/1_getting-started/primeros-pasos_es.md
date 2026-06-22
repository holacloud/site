# Primeros Pasos

Esta guia crea un scheduler, encola una tarea retrasada, lista tareas, reserva con lease y consulta health.

## Crear Scheduler

```bash
curl -X POST "https://api.hola.cloud/schedulers" \
  -H "X-API-Key: TU_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "id": "sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890",
    "display_name": "mi-scheduler"
  }'
```

```json
{
  "scheduler": {
    "id": "sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890",
    "ready": true,
    "scheduled": 0,
    "inflight": 0,
    "display_name": "mi-scheduler"
  }
}
```

## Encolar Tarea

```bash
curl -X POST "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks" \
  -H "X-API-Key: TU_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "id": "tarea-001",
    "payload": { "type": "enviar_correo" },
    "delay": "60s",
    "labels": ["proyecto:incorporacion", "prioridad:alta"]
  }'
```

Respuesta: HTTP `202 Accepted` con cuerpo vacio.

## Listar Tareas

```json
{
  "scheduled": [
    {
      "id": "tarea-001",
      "future": "2025-06-21T12:01:01Z",
      "labels": ["proyecto:incorporacion", "prioridad:alta"]
    }
  ],
  "inflight": [],
  "scheduled_meta": { "page": 1, "per_page": 25, "total": 1, "total_pages": 1 },
  "inflight_meta": { "page": 1, "per_page": 25, "total": 0, "total_pages": 0 }
}
```

## Reservar y Confirmar

```json
{ "worktime": "30s" }
```

```json
{
  "id": "tarea-001",
  "payload": { "type": "enviar_correo" },
  "lease_expires_at": "2025-06-21T12:02:01Z",
  "labels": ["proyecto:incorporacion", "prioridad:alta"]
}
```

Para confirmar, elimina la tarea. La respuesta es HTTP `204 No Content`.

## Health

```json
{
  "status": "ok",
  "ready": true,
  "scheduled": 0,
  "inflight": 0,
  "scheduler_id": "sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890"
}
```
