# Referencia de la API Files

URL base: `https://api.hola.cloud`

Todos los endpoints `/v1` requieren `X-Glue-Authentication`.

```http
X-Glue-Authentication: {"user":{"id":"user-123"}}
```

## Endpoints

| Método | Ruta | Descripción |
|--------|------|-------------|
| GET | `/v1/buckets` | Lista buckets |
| POST | `/v1/buckets` | Crea un bucket |
| GET | `/v1/buckets/{bucket_id}` | Obtiene detalles de un bucket |
| PATCH | `/v1/buckets/{bucket_id}` | No implementado |
| DELETE | `/v1/buckets/{bucket_id}` | Elimina un bucket |
| GET | `/v1/buckets/{bucket_id}/list/*` | Lista archivos por prefijo |
| PUT | `/v1/buckets/{bucket_id}/files/*` | Sube un archivo |
| GET | `/v1/buckets/{bucket_id}/files/*` | Descarga un archivo |
| DELETE | `/v1/buckets/{bucket_id}/files/*` | Elimina un archivo |
| HEAD | `/v1/buckets/{bucket_id}/files/*` | Devuelve headers del archivo |

## Campos

Campos de bucket: `id`, `project_id`, `created_timestamp`, `owners`, `name`, `description`.

Campos de file: `id`, `uuid`, `created_timestamp`, `updated_timestamp`, `owners`, `status`, `size`, `name`, `bucket`, `hash_md5`, `hash_sha256`, `mime_type`.

## Errores Comunes

| Estado | Descripción |
|--------|-------------|
| 401 | Falta `X-Glue-Authentication` o es inválido |
| 404 | Bucket o archivo no encontrado |
| 409 | El bucket ya existe |
| 500 | Error de persistencia o filesystem |
