# PATCH /schedulers/{id}

Actualiza parcialmente la configuración de un scheduler.

## Autenticación

Requiere autenticación. Pasa tu clave API mediante el encabezado `X-API-Key` o `Authorization: Bearer`.

## Parámetros de Ruta

| Parámetro | Tipo | Descripción |
|-----------|------|-------------|
| id | string | Identificador único del scheduler |

## Cuerpo de la Solicitud

```json
{
  "display_name": "scheduler-actualizado"
}
```

## Ejemplo de Solicitud

```bash
curl -X PATCH "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890" \
  -H "X-API-Key: TU_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "display_name": "scheduler-actualizado"
  }'
```

## Ejemplo de Respuesta

```http
HTTP/1.1 200 OK
Content-Type: application/json
```

```json
{
  "id": "sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890",
  "display_name": "scheduler-actualizado",
  "task_count": 5,
  "status": "active",
  "created_at": "2025-06-20T10:00:00Z",
  "updated_at": "2025-06-21T09:15:00Z"
}
```

## Códigos de Error

| Estado | Código | Descripción |
|--------|--------|-------------|
| 400 | invalid_request | Cuerpo de solicitud inválido |
| 401 | unauthorized | Clave API faltante o inválida |
| 404 | not_found | Scheduler no encontrado |
| 500 | internal_error | Error interno del servidor |
