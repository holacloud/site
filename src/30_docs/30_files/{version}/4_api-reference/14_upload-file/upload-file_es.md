# Subir Archivo

Sube un archivo a un bucket. La ruta del archivo va después de `/files/`.

## Autenticación

Requiere `X-Glue-Authentication`.

## Solicitud

```bash
curl -X PUT "https://api.hola.cloud/v1/buckets/bucket-550e8400-e29b-41d4-a716-446655440000/files/images/logo.png" \
  -H 'X-Glue-Authentication: {"user":{"id":"user-123"}}' \
  -H "Content-Type: image/png" \
  --data-binary @logo.png
```

## Respuesta

Devuelve un objeto file con `id`, `uuid`, `created_timestamp`, `updated_timestamp`, `owners`, `status`, `size`, `name`, `bucket`, `hash_md5`, `hash_sha256` y `mime_type`.
