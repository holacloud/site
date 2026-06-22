
# Create Collection

Crea una nueva colección dentro de la base de datos especificada. Las colecciones contienen documentos.

## Autenticación

Requiere las cabeceras `Api-Key`, `Api-Secret` y `X-Project`.

## Parámetros de Ruta

| Parámetro | Tipo | Descripción |
|-----------|------|-------------|
| id | uuid | Identificador único de la base de datos |

## Cuerpo de la Solicitud

| Campo | Tipo | Descripción |
|-------|------|-------------|
| name | string | El nombre de la colección |

## Solicitud HTTP

```http
POST /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
X-Project: acme-webapp
Content-Type: application/json

{
  "name": "mi-colección"
}
```

## Ejemplo

```bash
curl -X POST "https://api.hola.cloud/v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections" \
  -H "Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d" \
  -H "Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf" \
  -H "X-Project: acme-webapp" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "mi-colección"
  }'
```

## Respuesta

```json
{
  "defaults": {
    "id": "uuid()"
  },
  "indexes": 0,
  "name": "mi-colección",
  "total": 0
}
```

## Códigos de Error

| Código | Descripción |
|--------|-------------|
| 400 | Campo `name` faltante o inválido |
| 401 | Cabeceras de autenticación faltantes o inválidas |
| 403 | Acceso denegado al proyecto |
| 404 | Base de datos no encontrada |
| 409 | Ya existe una colección con el mismo nombre |
