# Listar Buckets

Lista buckets del usuario autenticado.

## Autenticación

Requiere `X-Glue-Authentication`.

## Solicitud

```bash
curl "https://api.hola.cloud/v1/buckets" \
  -H 'X-Glue-Authentication: {"user":{"id":"user-123"}}'
```

## Respuesta

La respuesta es un array JSON con `id`, `name`, `description`, `created_timestamp`, `created_h` y `owners`.
