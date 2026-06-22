# Obtener Scheduler

Devuelve los detalles de un scheduler específico.

## Autenticación

Este endpoint es público. No requiere autenticación.

## Parámetros de Ruta

| Parámetro | Tipo | Descripción |
|-----------|------|-------------|
| id | string | Identificador único del scheduler |

## Ejemplo de Solicitud

```bash
curl "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890"
```

## Ejemplo de Respuesta

```http
HTTP/1.1 200 OK
Content-Type: application/json
```

```json
{
  "id": "sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890",
  "display_name": "mi-scheduler",
  "task_count": 5,
  "status": "active",
  "created_at": "2025-06-20T10:00:00Z",
  "updated_at": "2025-06-21T08:30:00Z"
}
```

## Códigos de Error

| Estado | Código | Descripción |
|--------|--------|-------------|
| 404 | not_found | Scheduler no encontrado |
| 500 | internal_error | Error interno del servidor |
