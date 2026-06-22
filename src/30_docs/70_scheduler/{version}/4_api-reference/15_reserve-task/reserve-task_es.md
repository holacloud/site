# Reservar Tarea

Reserva una tarea disponible para procesamiento. La tarea entra en un período de concesión definido por `worktime`.

## Autenticación

Requiere autenticación. Pasa tu clave API mediante el encabezado `X-API-Key` o `Authorization: Bearer`.

## Parámetros de Ruta

| Parámetro | Tipo | Descripción |
|-----------|------|-------------|
| id | string | Identificador único del scheduler |

## Cuerpo de la Solicitud

| Campo | Tipo | Descripción |
|-------|------|-------------|
| worktime | integer | Duración de la concesión en segundos |

```json
{
  "worktime": "30s"
}
```

## Ejemplo de Solicitud

```bash
curl -X POST "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks/reserve" \
  -H "X-API-Key: TU_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "worktime": "30s"
  }'
```

## Ejemplo de Respuesta

```http
HTTP/1.1 200 OK
Content-Type: application/json
```

```json
{
  "id": "tarea-x1y2z3",
  "payload": {
    "type": "enviar_correo",
    "to": "usuario@example.com",
    "template": "bienvenida"
  },
  "lease_expires_at": "2025-06-21T12:02:01Z"
}
```

## Códigos de Error

| Estado | Código | Descripción |
|--------|--------|-------------|
| 401 | unauthorized | Clave API faltante o inválida |
| 404 | not_found | Scheduler no encontrado |
| 409 | conflict | No hay tareas disponibles para reservar |
| 500 | internal_error | Error interno del servidor |
