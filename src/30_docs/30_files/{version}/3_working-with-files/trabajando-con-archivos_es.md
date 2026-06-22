# Trabajando Con Archivos

Los archivos se identifican por la ruta después de `/files/`. El listado usa la ruta después de `/list/` como filtro de prefijo.

## Subir un Archivo

```bash
curl -X PUT "https://api.hola.cloud/v1/buckets/bucket-550e8400-e29b-41d4-a716-446655440000/files/notas/readme.txt" \
  -H 'X-Glue-Authentication: {"user":{"id":"user-123"}}' \
  -H "Content-Type: text/plain" \
  --data-binary @readme.txt
```

La respuesta es un objeto file con los campos implementados.

## Descargar un Archivo

```bash
curl "https://api.hola.cloud/v1/buckets/bucket-550e8400-e29b-41d4-a716-446655440000/files/notas/readme.txt" \
  -H 'X-Glue-Authentication: {"user":{"id":"user-123"}}' \
  -o readme.txt
```

El cuerpo de respuesta es el contenido guardado. `Content-Type` se toma de `mime_type`.

## Obtener Metadata JSON

Use `?metadata` en el endpoint de descarga para obtener el objeto file como JSON.

```bash
curl "https://api.hola.cloud/v1/buckets/bucket-550e8400-e29b-41d4-a716-446655440000/files/notas/readme.txt?metadata" \
  -H 'X-Glue-Authentication: {"user":{"id":"user-123"}}'
```

## Listar Archivos

```bash
curl "https://api.hola.cloud/v1/buckets/bucket-550e8400-e29b-41d4-a716-446655440000/list/notas/" \
  -H 'X-Glue-Authentication: {"user":{"id":"user-123"}}'
```

La respuesta es un array JSON de objetos file.

## HEAD de Archivo

```bash
curl -I "https://api.hola.cloud/v1/buckets/bucket-550e8400-e29b-41d4-a716-446655440000/files/notas/readme.txt" \
  -H 'X-Glue-Authentication: {"user":{"id":"user-123"}}'
```

El handler define `Last-Modified` y `Content-Length`.

## Eliminar un Archivo

```bash
curl -X DELETE "https://api.hola.cloud/v1/buckets/bucket-550e8400-e29b-41d4-a716-446655440000/files/notas/readme.txt" \
  -H 'X-Glue-Authentication: {"user":{"id":"user-123"}}'
```
