# Crear base de datos

Crea una nueva base de datos.

## Autenticación

Requiere `X-Glue-Authentication`.

## Solicitud HTTP

```http
POST /v1/databases HTTP/1.1
Host: api.hola.cloud
X-Glue-Authentication: tu-token-glue
Content-Type: application/json

{
  "name": "mi-base-de-datos"
}
```

## Respuesta

```json
{
  "id": "c3d4e5f6-a7b8-9012-cdef-123456789012",
  "name": "mi-base-de-datos",
  "creation_date": "2025-07-01T14:00:00Z",
  "owners": ["owner-id"],
  "api_keys": [],
  "owners_length": 1
}
```
