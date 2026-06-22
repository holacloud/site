# Primeros Pasos

Esta guía te explica cómo crear un scheduler, encolar una tarea con retraso, listar tareas, reservar y confirmar una tarea, y verificar el estado del scheduler.

## Prerrequisitos

- Una cuenta de HolaCloud con credenciales API.
- [curl](https://curl.se/) instalado en tu máquina.

## Autenticación

Algunos endpoints requieren autenticación mediante clave API. Pásala usando el encabezado `X-API-Key` o `Authorization: Bearer <key>`.

## Paso 1: Crear un Scheduler

Crea un nuevo scheduler:

```bash
curl -X POST "https://api.hola.cloud/schedulers" \
  -H "X-API-Key: TU_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "display_name": "mi-scheduler"
  }'
```

Respuesta esperada:

```json
{
  "id": "sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890",
  "display_name": "mi-scheduler",
  "task_count": 0,
  "status": "active"
}
```

Guarda el `id` devuelto — lo usarás en las siguientes solicitudes.

## Paso 2: Encolar una Tarea con Retraso

Encola una tarea que estará disponible en 60 segundos:

```bash
curl -X POST "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks" \
  -H "X-API-Key: TU_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "payload": {
      "type": "enviar_correo",
      "to": "usuario@example.com",
      "template": "bienvenida"
    },
    "delay": 60,
    "labels": {
      "proyecto": "incorporacion",
      "prioridad": "alta"
    }
  }'
```

Respuesta esperada:

```json
{
  "id": "tarea-001",
  "scheduler_id": "sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890",
  "payload": {
    "type": "enviar_correo",
    "to": "usuario@example.com",
    "template": "bienvenida"
  },
  "state": "pending",
  "available_at": "2025-06-21T12:01:01Z",
  "labels": {
    "proyecto": "incorporacion",
    "prioridad": "alta"
  }
}
```

## Paso 3: Listar Tareas

Ve todas las tareas del scheduler:

```bash
curl "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks"
```

Respuesta esperada:

```json
{
  "tasks": [
    {
      "id": "tarea-001",
      "state": "pending",
      "available_at": "2025-06-21T12:01:01Z",
      "labels": { "proyecto": "incorporacion", "prioridad": "alta" }
    }
  ],
  "total": 1
}
```

## Paso 4: Reservar y Confirmar una Tarea

Una vez que haya transcurrido el retraso, reserva la tarea:

```bash
curl -X POST "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks/reserve" \
  -H "X-API-Key: TU_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "worktime": 30
  }'
```

Respuesta esperada:

```json
{
  "id": "tarea-001",
  "payload": { "type": "enviar_correo", "to": "usuario@example.com", "template": "bienvenida" },
  "lease_expires_at": "2025-06-21T12:02:01Z"
}
```

Procesa la tarea y luego confírmala (elimínala):

```bash
curl -X DELETE "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks/tarea-001" \
  -H "X-API-Key: TU_API_KEY"
```

Respuesta esperada: HTTP `204 No Content`.

## Paso 5: Verificar el Estado del Scheduler

Verifica que el scheduler esté saludable:

```bash
curl "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/health"
```

Respuesta esperada:

```json
{
  "status": "ok",
  "uptime_seconds": 12345
}
```

## Siguientes Pasos

- Aprende a gestionar schedulers, incluyendo actualizaciones y monitoreo, en [Administrando Schedulers](../2_managing-schedulers/administrando-schedulers_es.md).
- Profundiza en el ciclo de vida de las tareas — incluyendo extensión de concesión, reprogramación y streaming SSE — en [Ciclo de Vida de las Tareas](../3_task-lifecycle/ciclo-de-vida-tareas_es.md).
