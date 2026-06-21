# GET /schedulers/{id}/health

Devuelve el estado de salud de un scheduler.

## Autenticación

Este endpoint es público. No requiere autenticación.

## Parámetros de Ruta

| Parámetro | Tipo | Descripción |
|-----------|------|-------------|
| id | string | Identificador único del scheduler |

## Ejemplo de Solicitud

```bash
curl "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/health"
```

## Ejemplo de Respuesta

```http
HTTP/1.1 200 OK
Content-Type: application/json
```

```json
{
  "status": "ok",
  "uptime_seconds": 12345
}
```

## Códigos de Error

| Estado | Código | Descripción |
|--------|--------|-------------|
| 404 | not_found | Scheduler no encontrado |
| 500 | internal_error | Error interno del servidor |
