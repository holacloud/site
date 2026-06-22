# Modificar Bucket

Modificar los metadatos de un bucket existente. Solo se actualizan los campos proporcionados en el cuerpo de la solicitud.

## Autenticación

Requiere los encabezados `Api-Key` y `Api-Secret`.

## Parámetros de Ruta

| Parámetro | Tipo | Descripción |
|-----------|------|-------------|
| `id` | string | El ID del bucket (ej. `bkt_abc123`) |

## Cuerpo de la Solicitud

```json
{
  "name": "bucket-renombrado",
  "public": true
}
```

| Campo | Tipo | Requerido | Descripción |
|-------|------|-----------|-------------|
| `name` | string | no | Nuevo nombre único global para el bucket |
| `public` | boolean | no | Establecer en `true` para permitir acceso público a los archivos |

## Solicitud

```bash
curl -X PATCH "https://api.hola.cloud/v1/buckets/bkt_abc123" \
  -H "Api-Key: SU_API_KEY" \
  -H "Api-Secret: SU_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "bucket-renombrado",
    "public": true
  }'
```

## Respuesta

```http
HTTP/1.1 200 OK
Content-Type: application/json
```

```json
{
  "id": "bkt_abc123",
  "name": "bucket-renombrado",
  "createdAt": "2026-06-21T10:00:00Z",
  "size": 1048576,
  "fileCount": 5,
  "public": true
}
```

## Códigos de Error

| Estado | Código | Descripción |
|--------|--------|-------------|
| 400 | Bad Request | Cuerpo de solicitud inválido |
| 401 | Unauthorized | Credenciales API faltantes o inválidas |
| 404 | Not Found | El bucket especificado no existe |
| 409 | Conflict | El nuevo nombre ya está en uso |
| 500 | Internal Server Error | Ocurrió un error inesperado |
