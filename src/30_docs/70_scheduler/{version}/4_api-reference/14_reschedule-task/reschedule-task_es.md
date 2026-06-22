# Reprogramar Tarea

Reprograma una tarea con un nuevo retraso o tiempo futuro.

## Autenticación

Requiere autenticación. Pasa tu clave API mediante el encabezado `X-API-Key` o `Authorization: Bearer`.

## Parámetros de Ruta

| Parámetro | Tipo | Descripción |
|-----------|------|-------------|
| id | string | Identificador único del scheduler |
| task | string | Identificador único de la tarea |

## Cuerpo de la Solicitud

| Campo | Tipo | Descripción |
|-------|------|-------------|
| future | string | Marca de tiempo ISO 8601 para la disponibilidad |
| delay | integer | Retraso en segundos desde ahora (alternativa a future) |

```json
{
  "delay": "120s"
}
```

## Ejemplo de Solicitud

```bash
curl -X POST "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks/tarea-x1y2z3/reschedule" \
  -H "X-API-Key: TU_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "delay": "120s"
  }'
```

## Ejemplo de Respuesta

```http
HTTP/1.1 200 OK
Content-Type: application/json
```

```json
{
  "id": "tarea-x1y2z3",      "future": "2025-06-21T12:03:01Z"
}
```

## Códigos de Error

| Estado | Código | Descripción |
|--------|--------|-------------|
| 400 | validation_error | future/delay faltante o inválido |
| 401 | unauthorized | Clave API faltante o inválida |
| 404 | not_found | Scheduler o tarea no encontrados |
| 500 | internal_error | Error interno del servidor |
