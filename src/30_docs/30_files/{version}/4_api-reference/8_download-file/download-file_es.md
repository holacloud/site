# Descargar Archivo

Descarga un archivo de un bucket. La ruta del archivo va después de `/files/`.

## Autenticación

Requiere `X-Glue-Authentication`.

## Parámetros de Consulta

| Parámetro | Descripción |
|-----------|-------------|
| `metadata` | Si está presente, devuelve el objeto file como JSON en vez del cuerpo del archivo. |

## Solicitud

```bash
curl "https://api.hola.cloud/v1/buckets/bucket-550e8400-e29b-41d4-a716-446655440000/files/images/logo.png" \
  -H 'X-Glue-Authentication: {"user":{"id":"user-123"}}' \
  -o logo.png
```

## Respuesta

El cuerpo de respuesta es el contenido del archivo. `Content-Type` se toma de `mime_type`.

## Respuesta con Metadata

```bash
curl "https://api.hola.cloud/v1/buckets/bucket-550e8400-e29b-41d4-a716-446655440000/files/images/logo.png?metadata" \
  -H 'X-Glue-Authentication: {"user":{"id":"user-123"}}'
```

Devuelve el objeto file con los campos implementados.
