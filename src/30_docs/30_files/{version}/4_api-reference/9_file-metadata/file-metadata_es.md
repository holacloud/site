# Archivo Metadata

Obtener los metadatos de un archivo sin descargar su contenido. La ruta del archivo se especifica después de `/files/` en la URL.

## Autenticación

Requiere los encabezados `Api-Key` y `Api-Secret`.

## Parámetros de Ruta

| Parámetro | Tipo | Descripción |
|-----------|------|-------------|
| `id` | string | El ID del bucket (ej. `bkt_abc123`) |

## Solicitud

```bash
curl -I "https://api.hola.cloud/v1/buckets/bkt_abc123/files/images/logo.png" \
  -H "Api-Key: SU_API_KEY" \
  -H "Api-Secret: SU_API_SECRET"
```

## Respuesta

```http
HTTP/1.1 200 OK
Content-Type: image/png
Content-Length: 24576
Last-Modified: Sun, 21 Jun 2026 12:00:00 GMT
ETag: "abc123def456"
Accept-Ranges: bytes
```

La respuesta contiene solo encabezados; no se devuelve ningún cuerpo.

## Códigos de Error

| Estado | Código | Descripción |
|--------|--------|-------------|
| 401 | Unauthorized | Credenciales API faltantes o inválidas |
| 404 | Not Found | El bucket o archivo especificado no existe |
| 500 | Internal Server Error | Ocurrió un error inesperado |
