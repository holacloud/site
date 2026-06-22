# Listar Files

Listar archivos en un bucket. El comodín (`*`) puede reemplazarse con un prefijo de ruta opcional para filtrar resultados.

## Autenticación

Requiere los encabezados `Api-Key` y `Api-Secret`.

## Parámetros de Ruta

| Parámetro | Tipo | Descripción |
|-----------|------|-------------|
| `id` | string | El ID del bucket (ej. `bkt_abc123`) |

## Parámetros de Consulta

| Parámetro | Tipo | Valor por Defecto | Descripción |
|-----------|------|-------------------|-------------|
| `prefix` | string | "" | Filtrar archivos que comienzan con este prefijo |
| `delimiter` | string | "" | Agrupar resultados por este delimitador (ej. `/`) |
| `maxKeys` | integer | 1000 | Número máximo de archivos a devolver |

## Solicitud

```bash
curl "https://api.hola.cloud/v1/buckets/bkt_abc123/list/*?prefix=images/&maxKeys=50" \
  -H "Api-Key: SU_API_KEY" \
  -H "Api-Secret: SU_API_SECRET"
```

## Respuesta

```http
HTTP/1.1 200 OK
Content-Type: application/json
```

```json
{
  "files": [
    {
      "path": "images/logo.png",
      "size": 24576,
      "contentType": "image/png",
      "modifiedAt": "2026-06-21T11:00:00Z",
      "etag": "\"abc123def456\""
    },
    {
      "path": "images/banner.jpg",
      "size": 102400,
      "contentType": "image/jpeg",
      "modifiedAt": "2026-06-21T10:30:00Z",
      "etag": "\"789012ghi345\""
    }
  ],
  "prefix": "images/",
  "total": 2
}
```

## Códigos de Error

| Estado | Código | Descripción |
|--------|--------|-------------|
| 401 | Unauthorized | Credenciales API faltantes o inválidas |
| 404 | Not Found | El bucket especificado no existe |
| 500 | Internal Server Error | Ocurrió un error inesperado |
