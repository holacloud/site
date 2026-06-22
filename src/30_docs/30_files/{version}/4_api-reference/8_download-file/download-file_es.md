# Descargar Archivo

Descargar un archivo de un bucket. La ruta del archivo se especifica después de `/files/` en la URL.

## Autenticación

Requiere los encabezados `Api-Key` y `Api-Secret`.

## Parámetros de Ruta

| Parámetro | Tipo | Descripción |
|-----------|------|-------------|
| `id` | string | El ID del bucket (ej. `bkt_abc123`) |

## Parámetros de Consulta

| Parámetro | Tipo | Valor por Defecto | Descripción |
|-----------|------|-------------------|-------------|
| `metadata` | boolean | false | Establecer en `true` para devolver metadatos como JSON en lugar del cuerpo del archivo |

## Solicitud

```bash
curl "https://api.hola.cloud/v1/buckets/bkt_abc123/files/images/logo.png" \
  -H "Api-Key: SU_API_KEY" \
  -H "Api-Secret: SU_API_SECRET" \
  -o logo.png
```

## Respuesta

```http
HTTP/1.1 200 OK
Content-Type: image/png
Content-Length: 24576
Last-Modified: Sun, 21 Jun 2026 12:00:00 GMT
ETag: "abc123def456"
```

El cuerpo de la respuesta es el contenido del archivo en bruto.

### Respuesta de Metadatos

Cuando se especifica `?metadata=true`, la respuesta devuelve JSON en lugar del cuerpo del archivo:

```json
{
  "path": "images/logo.png",
  "size": 24576,
  "contentType": "image/png",
  "modifiedAt": "2026-06-21T12:00:00Z",
  "etag": "\"abc123def456\""
}
```

## Códigos de Error

| Estado | Código | Descripción |
|--------|--------|-------------|
| 401 | Unauthorized | Credenciales API faltantes o inválidas |
| 404 | Not Found | El bucket o archivo especificado no existe |
| 500 | Internal Server Error | Ocurrió un error inesperado |
