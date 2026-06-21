# POST /schedulers/{id}/tasks

Encola una nueva tarea para ejecutarse en un horario programado.

## Autenticación

Requiere autenticación. Pasa tu clave API mediante el encabezado `X-API-Key` o `Authorization: Bearer`.

## Parámetros de Ruta

| Parámetro | Tipo | Descripción |
|-----------|------|-------------|
| id | string | Identificador único del scheduler |

## Cuerpo de la Solicitud

| Campo | Tipo | Descripción |
|-------|------|-------------|
| id | string | ID de tarea proporcionado por el cliente (opcional) |
| future | string | Marca de tiempo ISO 8601 para la disponibilidad |
| delay | integer | Retraso en segundos desde ahora (alternativa a future) |
| payload | object | Carga útil JSON arbitraria para el trabajador |
| labels | object | Pares clave-valor opcionales para filtrado |

```json
{
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
}
```

## Ejemplo de Solicitud

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

## Ejemplo de Respuesta

```http
HTTP/1.1 201 Created
Content-Type: application/json
```

```json
{
  "id": "tarea-x1y2z3",
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

## Códigos de Error

| Estado | Código | Descripción |
|--------|--------|-------------|
| 400 | invalid_request | Cuerpo de solicitud faltante o inválido |
| 401 | unauthorized | Clave API faltante o inválida |
| 404 | not_found | Scheduler no encontrado |
| 500 | internal_error | Error interno del servidor |
