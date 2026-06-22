# Eliminar Bucket

Elimina un bucket por ID.

## Autenticación

Requiere `X-Glue-Authentication`.

## Solicitud

```bash
curl -X DELETE "https://api.hola.cloud/v1/buckets/bucket-550e8400-e29b-41d4-a716-446655440000" \
  -H 'X-Glue-Authentication: {"user":{"id":"user-123"}}'
```

## Respuesta

El cuerpo de respuesta es el objeto bucket eliminado con `id`, `project_id`, `created_timestamp`, `owners`, `name` y `description`.
