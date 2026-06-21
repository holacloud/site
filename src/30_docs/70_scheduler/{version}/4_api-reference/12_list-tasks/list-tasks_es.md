# GET /schedulers/{id}/tasks

Lista las tareas de un scheduler, con paginación separada para tareas programadas y en vuelo.

## Autenticación

Este endpoint es público. No requiere autenticación.

## Parámetros de Ruta

| Parámetro | Tipo | Descripción |
|-----------|------|-------------|
| id | string | Identificador único del scheduler |

## Parámetros de Consulta

| Parámetro | Tipo | Descripción |
|-----------|------|-------------|
| scheduled_page | integer | Número de página para tareas programadas (predeterminado: 1) |
| scheduled_per_page | integer | Elementos por página para tareas programadas (predeterminado: 20) |
| inflight_page | integer | Número de página para tareas en vuelo (predeterminado: 1) |
| inflight_per_page | integer | Elementos por página para tareas en vuelo (predeterminado: 20) |
| search | string | Buscar tareas por contenido del payload |
| label | string | Filtrar por etiqueta (formato: clave:valor) |

## Ejemplo de Solicitud

```bash
curl "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks?label=prioridad:alta"
```

## Ejemplo de Respuesta

```http
HTTP/1.1 200 OK
Content-Type: application/json
```

```json
{
  "scheduled": {
    "tasks": [
      {
        "id": "tarea-x1y2z3",
        "state": "pending",
        "available_at": "2025-06-21T12:01:01Z",
        "labels": { "proyecto": "incorporacion", "prioridad": "alta" }
      }
    ],
    "total": 1,
    "page": 1,
    "per_page": 20
  },
  "inflight": {
    "tasks": [],
    "total": 0,
    "page": 1,
    "per_page": 20
  }
}
```

## Códigos de Error

| Estado | Código | Descripción |
|--------|--------|-------------|
| 404 | not_found | Scheduler no encontrado |
| 500 | internal_error | Error interno del servidor |
