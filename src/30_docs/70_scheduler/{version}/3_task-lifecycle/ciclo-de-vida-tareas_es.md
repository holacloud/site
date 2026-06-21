# Ciclo de Vida de las Tareas

Este documento detalla el ciclo de vida completo de una tarea: encolado, listado, reserva, extensión de concesión, confirmación, reprogramación y streaming en tiempo real mediante SSE.

## Estados de una Tarea

Una tarea pasa por los siguientes estados:

- **pending** — Esperando su hora programada. Aún no visible para los trabajadores.
- **available** — La hora programada ha llegado; lista para ser reservada.
- **reserved** — En manos de un trabajador bajo una concesión.
- **completed** — Confirmada y eliminada por el trabajador.

Si una concesión expira sin confirmación, la tarea vuelve a **available**.

## Encolar una Tarea

Crea una nueva tarea con un payload, hora futura (o retraso) y etiquetas opcionales:

```bash
curl -X POST "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks" \
  -H "X-API-Key: TU_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "payload": {
      "accion": "generar_informe",
      "informe_id": "rpt-5678"
    },
    "future": "2025-06-22T00:00:00Z",
    "labels": {
      "equipo": "analitica",
      "entorno": "produccion"
    }
  }'
```

```http
POST /schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks HTTP/1.1
Host: api.hola.cloud
X-API-Key: TU_API_KEY
Content-Type: application/json

{
  "payload": {
    "accion": "generar_informe",
    "informe_id": "rpt-5678"
  },
  "future": "2025-06-22T00:00:00Z",
  "labels": {
    "equipo": "analitica",
    "entorno": "produccion"
  }
}
```

Alternativamente, usa `delay` (segundos desde ahora) en lugar de `future`:

```json
{
  "payload": { "accion": "enviar_recordatorio" },
  "delay": 3600,
  "labels": { "tipo": "correo" }
}
```

Respuesta esperada:

```json
{
  "id": "tarea-001",
  "scheduler_id": "sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890",
  "payload": { "accion": "generar_informe", "informe_id": "rpt-5678" },
  "state": "pending",
  "available_at": "2025-06-22T00:00:00Z",
  "labels": { "equipo": "analitica", "entorno": "produccion" }
}
```

## Listar Tareas con Paginación

Lista tareas con paginación opcional y filtrado por etiquetas:

```bash
curl "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks?offset=0&limit=20"
```

```http
GET /schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks?offset=0&limit=20 HTTP/1.1
Host: api.hola.cloud
```

Respuesta esperada:

```json
{
  "tasks": [
    {
      "id": "tarea-001",
      "state": "pending",
      "available_at": "2025-06-22T00:00:00Z",
      "labels": { "equipo": "analitica" }
    },
    {
      "id": "tarea-002",
      "state": "available",
      "available_at": "2025-06-21T12:00:00Z",
      "labels": { "equipo": "analitica" }
    }
  ],
  "total": 2,
  "offset": 0,
  "limit": 20
}
```

## Reservar una Tarea

Un trabajador reserva una tarea disponible especificando un `worktime` — la duración de la concesión en segundos:

```bash
curl -X POST "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks/reserve" \
  -H "X-API-Key: TU_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "worktime": 60
  }'
```

```http
POST /schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks/reserve HTTP/1.1
Host: api.hola.cloud
X-API-Key: TU_API_KEY
Content-Type: application/json

{
  "worktime": 60
}
```

Respuesta esperada:

```json
{
  "id": "tarea-002",
  "payload": { "accion": "enviar_recordatorio" },
  "lease_expires_at": "2025-06-21T12:01:00Z"
}
```

Si no hay tareas disponibles, el endpoint devuelve HTTP `204 No Content`.

## Extender una Concesión

Si el trabajador necesita más tiempo, extiende la concesión antes de que expire:

```bash
curl -X POST "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks/tarea-002/extend" \
  -H "X-API-Key: TU_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "worktime": 30
  }'
```

```http
POST /schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks/tarea-002/extend HTTP/1.1
Host: api.hola.cloud
X-API-Key: TU_API_KEY
Content-Type: application/json

{
  "worktime": 30
}
```

Respuesta esperada:

```json
{
  "id": "tarea-002",
  "lease_expires_at": "2025-06-21T12:01:30Z"
}
```

## Confirmar (Eliminar) una Tarea

Después de procesar exitosamente, confirma la tarea para eliminarla:

```bash
curl -X DELETE "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks/tarea-002" \
  -H "X-API-Key: TU_API_KEY"
```

```http
DELETE /schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks/tarea-002 HTTP/1.1
Host: api.hola.cloud
X-API-Key: TU_API_KEY
```

Respuesta esperada: HTTP `204 No Content`.

## Reprogramar una Tarea

Si el procesamiento falla, reprograma la tarea con un nuevo retraso u hora futura:

```bash
curl -X POST "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks/tarea-002/reschedule" \
  -H "X-API-Key: TU_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "delay": 300
  }'
```

```http
POST /schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks/tarea-002/reschedule HTTP/1.1
Host: api.hola.cloud
X-API-Key: TU_API_KEY
Content-Type: application/json

{
  "delay": 300
}
```

Respuesta esperada:

```json
{
  "id": "tarea-002",
  "state": "pending",
  "available_at": "2025-06-21T12:06:00Z"
}
```

Esto es ideal para implementar lógica de reintento con retroceso exponencial.

## Flujo SSE para Actualizaciones en Tiempo Real

Suscríbete a eventos de tareas en tiempo real mediante Server-Sent Events:

```bash
curl "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks/stream"
```

El flujo emite eventos en formato SSE:

```
event: task_available
data: {"id":"tarea-003","payload":{"accion":"procesar_imagen"},"labels":{}}

event: task_completed
data: {"id":"tarea-003"}

event: task_expired
data: {"id":"tarea-003"}
```

| Evento | Descripción |
|--------|-------------|
| `task_available` | Una tarea está disponible para ser reservada |
| `task_completed` | Una tarea fue confirmada por un trabajador |
| `task_expired` | Una concesión expiró y la tarea está disponible nuevamente |

SSE es útil para construir grupos de trabajadores receptivos que reaccionan a las tareas inmediatamente sin necesidad de sondeo.
