# DELETE /schedulers/{id}/tasks/{task}

Confirma y elimina una tarea del scheduler. Debe llamarse después de que un trabajador procese exitosamente una tarea reservada.

## Autenticación

Requiere autenticación. Pasa tu clave API mediante el encabezado `X-API-Key` o `Authorization: Bearer`.

## Parámetros de Ruta

| Parámetro | Tipo | Descripción |
|-----------|------|-------------|
| id | string | Identificador único del scheduler |
| task | string | Identificador único de la tarea |

## Ejemplo de Solicitud

```bash
curl -X DELETE "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/tasks/tarea-x1y2z3" \
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
| 404 | not_found | Scheduler o tarea no encontrados |
| 409 | conflict | La tarea no está reservada actualmente |
| 500 | internal_error | Error interno del servidor |
