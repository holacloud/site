# Headers de Archivo

Devuelve headers para un archivo con `HEAD /v1/buckets/{bucket_id}/files/*`.

## Autenticación

Requiere `X-Glue-Authentication`.

## Solicitud

```bash
curl -I "https://api.hola.cloud/v1/buckets/bucket-550e8400-e29b-41d4-a716-446655440000/files/images/logo.png" \
  -H 'X-Glue-Authentication: {"user":{"id":"user-123"}}'
```

## Respuesta

El handler define `Last-Modified` y `Content-Length`.
