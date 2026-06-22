# Trabajando Con Archivos

Esta guía cubre las operaciones con archivos en HolaCloud Files: subir, descargar, listar, obtener metadatos y eliminar archivos.

## Rutas de Archivos

Las rutas de los archivos se especifican en la URL después de `/files/` o `/list/`. Las rutas pueden incluir directorios usando barras inclinadas:

```
/v1/buckets/{id}/files/images/logo.png
/v1/buckets/{id}/files/backups/db-2026-06-21.sql.gz
```

No es necesario crear directorios previamente -- la ruta del archivo se trata como una clave plana con barras.

## Subir un Archivo

Use PUT para subir un archivo. Establezca el encabezado `Content-Type` al tipo MIME apropiado.

### Subir un Archivo de Texto

```bash
curl -X PUT "https://api.hola.cloud/v1/buckets/bkt_abc123/files/notas/leame.txt" \
  -H "Api-Key: SU_API_KEY" \
  -H "Api-Secret: SU_API_SECRET" \
  -H "Content-Type: text/plain" \
  --data-binary @leame.txt
```

### Subir un Archivo Binario (Imagen)

```bash
curl -X PUT "https://api.hola.cloud/v1/buckets/bkt_abc123/files/images/foto.jpg" \
  -H "Api-Key: SU_API_KEY" \
  -H "Api-Secret: SU_API_SECRET" \
  -H "Content-Type: image/jpeg" \
  --data-binary @foto.jpg
```

### Subir con Metadatos

Puede adjuntar metadatos personalizados a los archivos usando el encabezado `File-Metadata`:

```bash
curl -X PUT "https://api.hola.cloud/v1/buckets/bkt_abc123/files/informe.pdf" \
  -H "Api-Key: SU_API_KEY" \
  -H "Api-Secret: SU_API_SECRET" \
  -H "Content-Type: application/pdf" \
  -H "File-Metadata: {\"author\": \"Ana García\", \"department\": \"ingeniería\"}" \
  --data-binary @informe.pdf
```

Respuesta esperada:

```json
{
  "path": "informe.pdf",
  "size": 245760,
  "contentType": "application/pdf",
  "uploadedAt": "2026-06-21T12:00:00Z",
  "etag": "\"abc123def456\"",
  "metadata": {
    "author": "Ana García",
    "department": "ingeniería"
  }
}
```

## Descargar un Archivo

Use GET para descargar un archivo por su ruta:

```bash
curl "https://api.hola.cloud/v1/buckets/bkt_abc123/files/images/foto.jpg" \
  -H "Api-Key: SU_API_KEY" \
  -H "Api-Secret: SU_API_SECRET" \
  -o foto.jpg
```

Solicitud HTTP:

```http
GET /v1/buckets/bkt_abc123/files/images/foto.jpg HTTP/1.1
Host: api.hola.cloud
Api-Key: SU_API_KEY
Api-Secret: SU_API_SECRET
```

La respuesta incluye el contenido del archivo con los encabezados `Content-Type` y `Content-Length` apropiados.

### Descargas Parciales (Solicitudes de Rango)

Files soporta solicitudes de rango HTTP para descargas parciales:

```bash
curl "https://api.hola.cloud/v1/buckets/bkt_abc123/files/video.mp4" \
  -H "Api-Key: SU_API_KEY" \
  -H "Api-Secret: SU_API_SECRET" \
  -H "Range: bytes=0-1023" \
  -o fragmento_video.mp4
```

## Listar Archivos

Liste archivos en un bucket con un filtro de prefijo opcional. El `*` después de `/list/` es obligatorio e indica todos los archivos. Use una ruta específica para filtrar por prefijo.

### Listar Todos los Archivos

```bash
curl "https://api.hola.cloud/v1/buckets/bkt_abc123/list/*" \
  -H "Api-Key: SU_API_KEY" \
  -H "Api-Secret: SU_API_SECRET"
```

### Listar Archivos en un Directorio

Liste archivos bajo el prefijo `images/`:

```bash
curl "https://api.hola.cloud/v1/buckets/bkt_abc123/list/images/*" \
  -H "Api-Key: SU_API_KEY" \
  -H "Api-Secret: SU_API_SECRET"
```

Respuesta:

```json
{
  "files": [
    {
      "path": "images/foto.jpg",
      "size": 245760,
      "contentType": "image/jpeg",
      "modifiedAt": "2026-06-21T12:00:00Z"
    },
    {
      "path": "images/banner.png",
      "size": 512000,
      "contentType": "image/png",
      "modifiedAt": "2026-06-21T11:30:00Z"
    }
  ],
  "prefix": "images/",
  "total": 2
}
```

### Listar con Paginación

```bash
curl "https://api.hola.cloud/v1/buckets/bkt_abc123/list/*?limit=10&offset=0" \
  -H "Api-Key: SU_API_KEY" \
  -H "Api-Secret: SU_API_SECRET"
```

## Obtener Metadatos de un Archivo

Use HEAD para obtener metadatos de un archivo sin descargar el contenido:

```bash
curl -I "https://api.hola.cloud/v1/buckets/bkt_abc123/files/informe.pdf" \
  -H "Api-Key: SU_API_KEY" \
  -H "Api-Secret: SU_API_SECRET"
```

Encabezados de respuesta:

```http
HTTP/1.1 200 OK
Content-Type: application/pdf
Content-Length: 245760
ETag: "abc123def456"
Last-Modified: Sun, 21 Jun 2026 12:00:00 GMT
File-Metadata: {"author": "Ana García", "department": "ingeniería"}
```

## Eliminar un Archivo

Elimine un archivo por su ruta:

```bash
curl -X DELETE "https://api.hola.cloud/v1/buckets/bkt_abc123/files/notas/leame.txt" \
  -H "Api-Key: SU_API_KEY" \
  -H "Api-Secret: SU_API_SECRET"
```

Respuesta esperada: HTTP 204 Sin Contenido.

### Error: Archivo No Encontrado

```json
{
  "error": {
    "code": "FILE_NOT_FOUND",
    "message": "Archivo 'notas/leame.txt' no encontrado en el bucket 'bkt_abc123'."
  }
}
```

## Manejo de Content-Type

Al subir un archivo, debe establecer el encabezado `Content-Type` al tipo MIME correcto para su archivo. Si se omite, el servicio utiliza `application/octet-stream` por defecto.

Tipos MIME comunes:

| Extensión | Content-Type |
|-----------|-------------|
| .txt | text/plain |
| .html | text/html |
| .css | text/css |
| .js | application/javascript |
| .json | application/json |
| .jpg / .jpeg | image/jpeg |
| .png | image/png |
| .gif | image/gif |
| .pdf | application/pdf |
| .zip | application/zip |
| .gz | application/gzip |
| .mp4 | video/mp4 |

## Convenciones de Rutas de Archivos

- Las rutas de archivos distinguen entre mayúsculas y minúsculas
- No se permiten barras inclinadas iniciales: `notas/leame.txt` (correcto), `/notas/leame.txt` (incorrecto)
- Las rutas pueden tener hasta 1024 caracteres
- Use barras inclinadas (`/`) como separadores de directorio
- Los caracteres especiales en los segmentos de ruta deben codificarse como URL

## Resumen de Códigos de Estado HTTP

| Código | Descripción |
|--------|-------------|
| 200 | Éxito (GET, HEAD, PUT) |
| 204 | Sin Contenido (DELETE) |
| 400 | Solicitud Incorrecta |
| 404 | No Encontrado -- el archivo o bucket no existe |
| 416 | Rango No Satisfactorio |
| 500 | Error Interno del Servidor |
