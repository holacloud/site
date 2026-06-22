# Crear Scheduler

Crea un nuevo scheduler.

## Autenticación

Requiere autenticación. Pasa tu clave API mediante el encabezado `X-API-Key` o `Authorization: Bearer`.

## Cuerpo de la Solicitud

```json
{
  "display_name": "mi-scheduler"
}
```

## Ejemplo de Solicitud

```bash
curl -X POST "https://api.hola.cloud/schedulers" \
  -H "X-API-Key: TU_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "display_name": "mi-scheduler"
  }'
```

## Ejemplo de Respuesta

```http
HTTP/1.1 202 Accepted
Content-Type: application/json
```

```json
{
  "id": "sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890",
  "display_name": "mi-scheduler",
  "ready": true,
  "scheduled": 0,
  "inflight": 0,
  "created_at": "2025-06-21T10:00:00Z"
}
```

## Códigos de Error

| Estado | Código | Descripción |
|--------|--------|-------------|
| 400 | validation_error | display_name faltante o inválido |
| 401 | unauthorized | Clave API faltante o inválida |
| 500 | internal_error | Error interno del servidor |
