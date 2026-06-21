# POST /schedulers/{id}/tasks/{task}/extend

Extiende la concesión de una tarea actualmente reservada.

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
| extension | integer | Tiempo adicional de concesión en segundos |

```json
{
  "extension": 30
}
```

## Ejemplo de Solicitud

```bash
curl -X POST "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks/tarea-x1y2z3/extend" \
  -H "X-API-Key: TU_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "extension": 30
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
  "lease_expires_at": "2025-06-21T12:02:31Z"
}
```

## Códigos de Error

| Estado | Código | Descripción |
|--------|--------|-------------|
| 400 | invalid_request | Valor de extensión faltante o inválido |
| 401 | unauthorized | Clave API faltante o inválida |
| 404 | not_found | Scheduler o tarea no encontrados |
| 409 | conflict | La tarea no está reservada o la concesión ya expiró |
| 500 | internal_error | Error interno del servidor |
