# Gestión de índices

Las acciones de índices administran índices secundarios de una colección.

## Autenticación

Requiere `Api-Key` y `Api-Secret`, o un token Glue de owner cuando el propietario de la base de datos está permitido.

## Crear índice

Los tipos de índice soportados son `map` y `btree`.

```http
POST /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections/users:createIndex HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
Content-Type: application/json

{"name":"by-email","type":"map","field":"email","unique":true}
```

Estado de respuesta: `201 Created`

```json
{"name":"by-email","type":"map","field":"email","unique":true}
```

## Listar índices

```http
POST /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections/users:listIndexes HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
```

Respuesta:

```json
[{"name":"by-email","type":"map","field":"email","unique":true}]
```

## Obtener índice

```http
POST /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections/users:getIndex HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
Content-Type: application/json

{"name":"by-email"}
```

Respuesta:

```json
{"name":"by-email","type":"map","field":"email","unique":true}
```

## Eliminar índice

```http
POST /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections/users:dropIndex HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
Content-Type: application/json

{"name":"by-email"}
```

Estado de respuesta: `204 No Content`
