# Obtener Bucket

Obtiene un bucket por ID para el usuario autenticado.

## Autenticación

Requiere `X-Glue-Authentication`.

## Solicitud

```bash
curl "https://api.hola.cloud/v1/buckets/bucket-550e8400-e29b-41d4-a716-446655440000" \
  -H 'X-Glue-Authentication: {"user":{"id":"user-123"}}'
```

## Respuesta

Objeto bucket con `id`, `project_id`, `created_timestamp`, `owners`, `name` y `description`.
