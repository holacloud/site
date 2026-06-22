# Crear Bucket

Crea un bucket para el usuario autenticado.

## Autenticación

Requiere `X-Glue-Authentication`.

## Cuerpo de Solicitud

```json
{
  "name": "mi-bucket",
  "description": "Descripción opcional"
}
```

Campos: `name` y `description`.

## Solicitud

```bash
curl -X POST "https://api.hola.cloud/v1/buckets" \
  -H 'X-Glue-Authentication: {"user":{"id":"user-123"}}' \
  -H "Content-Type: application/json" \
  -d '{"name":"mi-bucket","description":"Descripción opcional"}'
```

## Respuesta

```json
{
  "id": "bucket-550e8400-e29b-41d4-a716-446655440000",
  "project_id": "",
  "created_timestamp": 1782045600000000000,
  "owners": ["user-123"],
  "name": "mi-bucket",
  "description": "Descripción opcional"
}
```
