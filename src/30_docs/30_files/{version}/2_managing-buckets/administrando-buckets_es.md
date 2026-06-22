# Administrando Buckets

Los buckets son contenedores para archivos. Las operaciones implementadas son crear, listar, obtener y eliminar.

## Entrada de Bucket

`POST /v1/buckets` acepta `name` y `description`.

- `name` se recorta, puede estar vacío y puede contener letras, dígitos, `_` y `-` hasta 64 caracteres.
- `description` puede tener hasta 4096 caracteres.
- Si `name` está vacío, se usa el ID generado como nombre.

## Crear un Bucket

```bash
curl -X POST "https://api.hola.cloud/v1/buckets" \
  -H 'X-Glue-Authentication: {"user":{"id":"user-123"}}' \
  -H "Content-Type: application/json" \
  -d '{"name":"assets","description":"Activos de la aplicación"}'
```

## Listar Buckets

```bash
curl "https://api.hola.cloud/v1/buckets" \
  -H 'X-Glue-Authentication: {"user":{"id":"user-123"}}'
```

La respuesta es un array JSON con `id`, `name`, `description`, `created_timestamp`, `created_h` y `owners`.

## Obtener Detalles

```bash
curl "https://api.hola.cloud/v1/buckets/bucket-550e8400-e29b-41d4-a716-446655440000" \
  -H 'X-Glue-Authentication: {"user":{"id":"user-123"}}'
```

La respuesta es un objeto bucket con `id`, `project_id`, `created_timestamp`, `owners`, `name` y `description`.

## Modificar Bucket

`PATCH /v1/buckets/{bucket_id}` está registrado pero no implementado.

## Eliminar Bucket

```bash
curl -X DELETE "https://api.hola.cloud/v1/buckets/bucket-550e8400-e29b-41d4-a716-446655440000" \
  -H 'X-Glue-Authentication: {"user":{"id":"user-123"}}'
```

El cuerpo de respuesta es el objeto bucket eliminado.
