# Actualizar base de datos

Actualiza propiedades de una base de datos.

## Autenticación

Requiere `X-Glue-Authentication`.

## Solicitud HTTP

```http
PATCH /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890 HTTP/1.1
Host: api.hola.cloud
X-Glue-Authentication: tu-token-glue
Content-Type: application/json

{
  "name": "base-renombrada"
}
```

## Respuesta

```json
{
  "id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
  "name": "base-renombrada",
  "creation_date": "2025-06-15T10:30:00Z",
  "owners": ["owner-id"],
  "api_keys": ["api-key-id"],
  "owners_length": 1
}
```
