# Administrando Schedulers

Este documento cubre el ciclo de vida de los schedulers: creación con metadatos, listado con búsqueda y filtro, actualización de configuración, monitoreo de salud y eliminación.

## Crear un Scheduler

Crea un nuevo scheduler con un nombre visible:

```bash
curl -X POST "https://api.hola.cloud/schedulers" \
  -H "X-API-Key: TU_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "display_name": "scheduler-correos"
  }'
```

```http
POST /schedulers HTTP/1.1
Host: api.hola.cloud
X-API-Key: TU_API_KEY
Content-Type: application/json

{
  "display_name": "scheduler-correos"
}
```

Respuesta esperada:

```json
{
  "id": "sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890",
  "display_name": "scheduler-correos",
  "task_count": 0,
  "status": "active"
}
```

## Listar Schedulers

Lista todos los schedulers de tu cuenta (público):

```bash
curl "https://api.hola.cloud/schedulers"
```

```http
GET /schedulers HTTP/1.1
Host: api.hola.cloud
```

Respuesta esperada:

```json
[
  {
    "id": "sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890",
    "display_name": "scheduler-correos",
    "task_count": 5,
    "status": "active"
  },
  {
    "id": "sched-b2c3d4e5-f6a7-8901-bcde-f12345678901",
    "display_name": "generador-informes",
    "task_count": 0,
    "status": "active"
  }
]
```

## Obtener Detalles de un Scheduler

Recupera un scheduler por su ID:

```bash
curl "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890"
```

```http
GET /schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890 HTTP/1.1
Host: api.hola.cloud
```

Respuesta esperada:

```json
{
  "id": "sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890",
  "display_name": "scheduler-correos",
  "task_count": 5,
  "status": "active"
}
```

## Actualizar un Scheduler

Actualiza el nombre visible u otros metadatos:

```bash
curl -X PATCH "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890" \
  -H "X-API-Key: TU_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "display_name": "scheduler-correos-v2"
  }'
```

```http
PATCH /schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890 HTTP/1.1
Host: api.hola.cloud
X-API-Key: TU_API_KEY
Content-Type: application/json

{
  "display_name": "scheduler-correos-v2"
}
```

Respuesta esperada:

```json
{
  "id": "sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890",
  "display_name": "scheduler-correos-v2",
  "task_count": 5,
  "status": "active"
}
```

PUT puede usarse indistintamente con PATCH.

## Monitoreo de Salud

Verifica el estado de salud de un scheduler en cualquier momento:

```bash
curl "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/health"
```

```http
GET /schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/health HTTP/1.1
Host: api.hola.cloud
```

Respuesta esperada:

```json
{
  "status": "ok",
  "uptime_seconds": 123456
}
```

Un estado distinto de `ok` indica que el scheduler no está disponible o está en un estado degradado.

## Eliminar un Scheduler

Elimina permanentemente un scheduler y todas sus tareas:

```bash
curl -X DELETE "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890" \
  -H "X-API-Key: TU_API_KEY"
```

```http
DELETE /schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890 HTTP/1.1
Host: api.hola.cloud
X-API-Key: TU_API_KEY
```

Respuesta esperada: HTTP `204 No Content`.

Una vez eliminado, el ID del scheduler ya no es válido y todas las tareas pendientes se descartan.
