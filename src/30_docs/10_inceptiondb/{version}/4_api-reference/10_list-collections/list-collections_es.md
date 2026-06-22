
# List Collections

Lista todas las colecciones dentro de una base de datos específica.

## Autenticación

Requiere las cabeceras `Api-Key`, `Api-Secret` y `X-Project`.

## Parámetros de Ruta

| Parámetro | Tipo | Descripción |
|-----------|------|-------------|
| id | uuid | Identificador único de la base de datos |

## Solicitud HTTP

```http
GET /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
X-Project: acme-webapp
```

## Ejemplo

```bash
curl -X GET "https://api.hola.cloud/v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections" \
  -H "Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d" \
  -H "Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf" \
  -H "X-Project: acme-webapp"
```

## Respuesta

```json
[
  {
    "name": "users",
    "document_count": 5400,
    "indexes": 2
  },
  {
    "name": "orders",
    "document_count": 12300,
    "indexes": 3
  },
  {
    "name": "products",
    "document_count": 890,
    "indexes": 1
  }
]
```

## Códigos de Error

| Código | Descripción |
|--------|-------------|
| 401 | Cabeceras de autenticación faltantes o inválidas |
| 403 | Acceso denegado al proyecto |
| 404 | Base de datos no encontrada |
