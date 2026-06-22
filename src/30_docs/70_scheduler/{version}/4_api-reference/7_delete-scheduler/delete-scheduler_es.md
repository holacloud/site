# Eliminar Scheduler

Elimina un scheduler y todas sus tareas.

## Autenticación

Requiere autenticación. Pasa tu clave API mediante el encabezado `X-API-Key` o `Authorization: Bearer`.

## Parámetros de Ruta

| Parámetro | Tipo | Descripción |
|-----------|------|-------------|
| id | string | Identificador único del scheduler |

## Ejemplo de Solicitud

```bash
curl -X DELETE "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890" \
  -H "X-API-Key: TU_API_KEY"
```

## Ejemplo de Respuesta

```http
HTTP/1.1 204 No Content
```

## Códigos de Error

| Estado | Código | Descripción |
|--------|--------|-------------|
| 401 | unauthorized | Clave API faltante o inválida |
| 404 | not_found | Scheduler no encontrado |
| 500 | internal_error | Error interno del servidor |
