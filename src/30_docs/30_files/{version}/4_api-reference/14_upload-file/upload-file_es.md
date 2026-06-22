# Subir Archivo

Subir un archivo a un bucket. La ruta del archivo se especifica después de `/files/` en la URL. El tamaño máximo del archivo es de 5 GB.

## Autenticación

Requiere los encabezados `Api-Key` y `Api-Secret`.

## Parámetros de Ruta

| Parámetro | Tipo | Descripción |
|-----------|------|-------------|
| `id` | string | El ID del bucket (ej. `bkt_abc123`) |

## Encabezados de Solicitud

| Encabezado | Requerido | Descripción |
|------------|-----------|-------------|
| `Content-Type` | sí | Tipo MIME del archivo que se está subiendo |
| `Content-Length` | sí | Tamaño del archivo en bytes |

## Solicitud

```bash
curl -X PUT "https://api.hola.cloud/v1/buckets/bkt_abc123/files/images/logo.png" \
  -H "Api-Key: SU_API_KEY" \
  -H "Api-Secret: SU_API_SECRET" \
  -H "Content-Type: image/png" \
  --data-binary @logo.png
```

## Respuesta

```http
HTTP/1.1 201 Created
Content-Type: application/json
```

```json
{
  "path": "images/logo.png",
  "size": 24576,
  "contentType": "image/png",
  "uploadedAt": "2026-06-21T12:00:00Z",
  "etag": "\"abc123def456\""
}
```

## Códigos de Error

| Estado | Código | Descripción |
|--------|--------|-------------|
| 400 | Bad Request | Solicitud inválida o falta Content-Type |
| 401 | Unauthorized | Credenciales API faltantes o inválidas |
| 404 | Not Found | El bucket especificado no existe |
| 413 | Payload Too Large | El archivo excede el límite de 5 GB |
| 500 | Internal Server Error | Ocurrió un error inesperado |
