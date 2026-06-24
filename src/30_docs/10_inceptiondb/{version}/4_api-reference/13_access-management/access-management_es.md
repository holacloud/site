# Gestión de acceso

Las acciones de gestión de acceso de una base de datos crean API keys y administran IDs de usuarios owner.

## Autenticación

Requiere `X-Glue-Authentication` de un usuario owner de la base de datos.

## Crear API key

```http
POST /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890:createApiKey HTTP/1.1
Host: api.hola.cloud
X-Glue-Authentication: glue-session-token
Content-Type: application/json

{"name":"production"}
```

Estado de respuesta: `201 Created`

```json
{
  "name": "production",
  "key": "1abbe476-6ad6-4b97-9cca-6deb6ab2901d",
  "secret": "4bda6d52-762b-4e5d-bed7-85614c13b8bf",
  "creation_date": "2026-06-24T12:00:00Z"
}
```

El secret solo se devuelve al crear la key.

## Eliminar API key

```http
POST /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890:deleteApiKey HTTP/1.1
Host: api.hola.cloud
X-Glue-Authentication: glue-session-token
Content-Type: application/json

{"key":"1abbe476-6ad6-4b97-9cca-6deb6ab2901d"}
```

Las respuestas correctas tienen cuerpo vacío.

## Agregar owner

```http
POST /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890:addOwner HTTP/1.1
Host: api.hola.cloud
X-Glue-Authentication: glue-session-token
Content-Type: application/json

{"owner_id":"user-123"}
```

Respuesta:

```json
["current-owner","user-123"]
```

## Eliminar owner

```http
POST /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890:deleteOwner HTTP/1.1
Host: api.hola.cloud
X-Glue-Authentication: glue-session-token
Content-Type: application/json

{"owner_id":"user-123"}
```

Respuesta:

```json
["current-owner"]
```

El usuario autenticado no puede eliminarse a sí mismo de la lista de owners.
