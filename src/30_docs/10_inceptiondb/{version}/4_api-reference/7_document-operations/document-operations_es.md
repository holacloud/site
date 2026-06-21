
# POST /v1/databases/{id}/collections/{col}

Realiza operaciones con documentos en una colección. La operación específica se determina mediante un sufijo separado por dos puntos en el nombre de la colección.

## Autenticación

Requiere las cabeceras `Api-Key`, `Api-Secret` y `X-Project`.

## Parámetros de Ruta

| Parámetro | Tipo | Descripción |
|-----------|------|-------------|
| id | uuid | Identificador único de la base de datos |
| col | string | Nombre de la colección con el sufijo de operación |

## Operaciones

### Insertar — `{collection}:insert`

Inserta uno o más documentos en la colección.

```http
POST /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections/users:insert HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
X-Project: acme-webapp
Content-Type: application/json

{
  "name": "Alice",
  "email": "alice@example.com",
  "role": "admin"
}
```

```bash
curl -X POST "https://api.hola.cloud/v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections/users:insert" \
  -H "Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d" \
  -H "Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf" \
  -H "X-Project: acme-webapp" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Alice",
    "email": "alice@example.com",
    "role": "admin"
  }'
```

Respuesta:
```json
{
  "id": "d7e8f9a0-b1c2-3456-7890-123456789abc",
  "message": "Documento insertado exitosamente"
}
```

### Buscar — `{collection}:find`

Consulta documentos utilizando un filtro.

```http
POST /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections/users:find HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
X-Project: acme-webapp
Content-Type: application/json

{
  "filter": {
    "role": "admin"
  }
}
```

```bash
curl -X POST "https://api.hola.cloud/v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections/users:find" \
  -H "Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d" \
  -H "Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf" \
  -H "X-Project: acme-webapp" \
  -H "Content-Type: application/json" \
  -d '{
    "filter": {
      "role": "admin"
    }
  }'
```

Respuesta:
```json
{
  "documents": [
    {
      "id": "d7e8f9a0-b1c2-3456-7890-123456789abc",
      "name": "Alice",
      "email": "alice@example.com",
      "role": "admin"
    }
  ]
}
```

### Eliminar — `{collection}:remove`

Elimina documentos que coinciden con un filtro.

```http
POST /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections/users:remove HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
X-Project: acme-webapp
Content-Type: application/json

{
  "filter": {
    "email": "alice@example.com"
  }
}
```

```bash
curl -X POST "https://api.hola.cloud/v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections/users:remove" \
  -H "Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d" \
  -H "Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf" \
  -H "X-Project: acme-webapp" \
  -H "Content-Type: application/json" \
  -d '{
    "filter": {
      "email": "alice@example.com"
    }
  }'
```

Respuesta:
```json
{
  "deleted": 1,
  "message": "Documentos eliminados exitosamente"
}
```

### Modificar — `{collection}:patch`

Actualiza documentos que coinciden con un filtro.

```http
POST /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections/users:patch HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
X-Project: acme-webapp
Content-Type: application/json

{
  "filter": {
    "role": "admin"
  },
  "patch": {
    "role": "superadmin"
  }
}
```

```bash
curl -X POST "https://api.hola.cloud/v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections/users:patch" \
  -H "Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d" \
  -H "Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf" \
  -H "X-Project: acme-webapp" \
  -H "Content-Type: application/json" \
  -d '{
    "filter": {
      "role": "admin"
    },
    "patch": {
      "role": "superadmin"
    }
  }'
```

Respuesta:
```json
{
  "updated": 1,
  "message": "Documentos modificados exitosamente"
}
```

## Códigos de Error

| Código | Descripción |
|--------|-------------|
| 400 | Nombre de operación inválido o cuerpo de solicitud mal formado |
| 401 | Cabeceras de autenticación faltantes o inválidas |
| 403 | Acceso denegado al proyecto |
| 404 | Base de datos o colección no encontrada |
