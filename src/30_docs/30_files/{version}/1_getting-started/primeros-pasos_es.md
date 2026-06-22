# Primeros Pasos

Esta guía muestra cómo crear un bucket, subir un archivo, listar archivos, descargar y eliminar.

## Requisitos

- `curl` instalado localmente
- Un header `X-Glue-Authentication` válido para el usuario autenticado

```http
X-Glue-Authentication: {"user":{"id":"user-123"}}
```

Todas las solicitudes usan `https://api.hola.cloud`.

## Paso 1: Crear un Bucket

```bash
curl -X POST "https://api.hola.cloud/v1/buckets" \
  -H 'X-Glue-Authentication: {"user":{"id":"user-123"}}' \
  -H "Content-Type: application/json" \
  -d '{"name":"mi-primer-bucket","description":"Primer bucket de prueba"}'
```

Respuesta esperada:

```json
{
  "id": "bucket-550e8400-e29b-41d4-a716-446655440000",
  "project_id": "",
  "created_timestamp": 1782045600000000000,
  "owners": ["user-123"],
  "name": "mi-primer-bucket",
  "description": "Primer bucket de prueba"
}
```

## Paso 2: Subir un Archivo

```bash
curl -X PUT "https://api.hola.cloud/v1/buckets/bucket-550e8400-e29b-41d4-a716-446655440000/files/hola.txt" \
  -H 'X-Glue-Authentication: {"user":{"id":"user-123"}}' \
  -H "Content-Type: text/plain" \
  --data-binary "Hola, HolaCloud Files!"
```

Respuesta esperada: un objeto file con `id`, `uuid`, `created_timestamp`, `updated_timestamp`, `owners`, `status`, `size`, `name`, `bucket`, `hash_md5`, `hash_sha256` y `mime_type`.

## Paso 3: Listar Archivos

```bash
curl "https://api.hola.cloud/v1/buckets/bucket-550e8400-e29b-41d4-a716-446655440000/list/*" \
  -H 'X-Glue-Authentication: {"user":{"id":"user-123"}}'
```

La respuesta es un array JSON de objetos file.

## Paso 4: Descargar el Archivo

```bash
curl "https://api.hola.cloud/v1/buckets/bucket-550e8400-e29b-41d4-a716-446655440000/files/hola.txt" \
  -H 'X-Glue-Authentication: {"user":{"id":"user-123"}}'
```

## Paso 5: Eliminar el Archivo

```bash
curl -X DELETE "https://api.hola.cloud/v1/buckets/bucket-550e8400-e29b-41d4-a716-446655440000/files/hola.txt" \
  -H 'X-Glue-Authentication: {"user":{"id":"user-123"}}'
```

## Paso 6: Eliminar el Bucket

```bash
curl -X DELETE "https://api.hola.cloud/v1/buckets/bucket-550e8400-e29b-41d4-a716-446655440000" \
  -H 'X-Glue-Authentication: {"user":{"id":"user-123"}}'
```
