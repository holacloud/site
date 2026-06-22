# Listar Archivos

Lista archivos en un bucket. La ruta después de `/list/` se usa como filtro de prefijo.

## Autenticación

Requiere `X-Glue-Authentication`.

## Solicitud

```bash
curl "https://api.hola.cloud/v1/buckets/bucket-550e8400-e29b-41d4-a716-446655440000/list/images/" \
  -H 'X-Glue-Authentication: {"user":{"id":"user-123"}}'
```

## Respuesta

La respuesta es un array JSON de objetos file con los campos implementados.
