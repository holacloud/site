# Files

HolaCloud Files ofrece una API REST para almacenar archivos en buckets. Los buckets agrupan archivos, y cada archivo se identifica por la ruta después de `/files/` o se filtra por prefijo después de `/list/`.

## Autenticación

Todos los endpoints `/v1` requieren el header `X-Glue-Authentication`:

```http
X-Glue-Authentication: {"user":{"id":"user-123"}}
```

## Modelo de Datos

Campos de bucket: `id`, `project_id`, `created_timestamp`, `owners`, `name`, `description`.

Campos de file: `id`, `uuid`, `created_timestamp`, `updated_timestamp`, `owners`, `status`, `size`, `name`, `bucket`, `hash_md5`, `hash_sha256`, `mime_type`.

## Referencia API

Todos los endpoints usan la URL base `https://api.hola.cloud`.

| Método | Ruta | Descripción |
|--------|------|-------------|
| GET | `/v1/buckets` | Lista buckets del usuario autenticado |
| POST | `/v1/buckets` | Crea un bucket |
| GET | `/v1/buckets/{bucket_id}` | Obtiene detalles de un bucket |
| PATCH | `/v1/buckets/{bucket_id}` | No implementado |
| DELETE | `/v1/buckets/{bucket_id}` | Elimina un bucket |
| GET | `/v1/buckets/{bucket_id}/list/*` | Lista archivos, con filtro opcional por prefijo |
| PUT | `/v1/buckets/{bucket_id}/files/*` | Sube un archivo |
| GET | `/v1/buckets/{bucket_id}/files/*` | Descarga un archivo o devuelve metadata JSON con `?metadata` |
| DELETE | `/v1/buckets/{bucket_id}/files/*` | Elimina un archivo |
| HEAD | `/v1/buckets/{bucket_id}/files/*` | Devuelve headers del archivo |
