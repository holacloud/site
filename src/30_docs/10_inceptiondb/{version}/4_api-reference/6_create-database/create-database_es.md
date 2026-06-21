
# POST /v1/databases

Crea una nueva base de datos dentro del proyecto autenticado.

## Autenticación

Requiere las cabeceras `Api-Key`, `Api-Secret` y `X-Project`.

## Cuerpo de la Solicitud

| Campo | Tipo | Descripción |
|-------|------|-------------|
| name | string | Un nombre legible para la base de datos |

## Solicitud HTTP

```http
POST /v1/databases HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
X-Project: acme-webapp
Content-Type: application/json

{
  "name": "mi-base-de-datos"
}
```

## Ejemplo

```bash
curl -X POST "https://api.hola.cloud/v1/databases" \
  -H "Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d" \
  -H "Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf" \
  -H "X-Project: acme-webapp" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "mi-base-de-datos"
  }'
```

## Respuesta

```json
{
  "id": "c3d4e5f6-a7b8-9012-cdef-123456789012",
  "name": "mi-base-de-datos",
  "created_at": "2025-07-01T14:00:00Z",
  "collection_count": 0
}
```

## Códigos de Error

| Código | Descripción |
|--------|-------------|
| 400 | Campo `name` faltante o inválido |
| 401 | Cabeceras de autenticación faltantes o inválidas |
| 403 | Acceso denegado al proyecto |
| 409 | Ya existe una base de datos con el mismo nombre |
