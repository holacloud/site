
# PATCH /v1/databases/{id}

Actualiza las propiedades de una base de datos existente.

## Autenticación

Requiere las cabeceras `Api-Key`, `Api-Secret` y `X-Project`.

## Parámetros de Ruta

| Parámetro | Tipo | Descripción |
|-----------|------|-------------|
| id | uuid | Identificador único de la base de datos |

## Cuerpo de la Solicitud

| Campo | Tipo | Descripción |
|-------|------|-------------|
| name | string | Nuevo nombre para la base de datos |

## Solicitud HTTP

```http
PATCH /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890 HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
X-Project: acme-webapp
Content-Type: application/json

{
  "name": "base-de-datos-renombrada"
}
```

## Ejemplo

```bash
curl -X PATCH "https://api.hola.cloud/v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890" \
  -H "Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d" \
  -H "Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf" \
  -H "X-Project: acme-webapp" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "base-de-datos-renombrada"
  }'
```

## Respuesta

```json
{
  "id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
  "name": "base-de-datos-renombrada",
  "updated_at": "2025-07-02T09:15:00Z"
}
```

## Códigos de Error

| Código | Descripción |
|--------|-------------|
| 400 | Cuerpo de solicitud faltante o inválido |
| 401 | Cabeceras de autenticación faltantes o inválidas |
| 403 | Acceso denegado al proyecto |
| 404 | Base de datos no encontrada |
| 409 | Ya existe una base de datos con el nuevo nombre |
