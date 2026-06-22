# Eliminar Archivo

Elimina un archivo de un bucket. La ruta del archivo va después de `/files/`.

## Autenticación

Requiere `X-Glue-Authentication`.

## Solicitud

```bash
curl -X DELETE "https://api.hola.cloud/v1/buckets/bucket-550e8400-e29b-41d4-a716-446655440000/files/images/logo.png" \
  -H 'X-Glue-Authentication: {"user":{"id":"user-123"}}'
```

## Respuesta

Devuelve una respuesta exitosa vacía cuando se elimina el archivo.
