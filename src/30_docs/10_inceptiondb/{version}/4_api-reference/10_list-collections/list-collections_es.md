# Listar colecciones

Lista las colecciones de una base de datos.

## Autenticación

Requiere `Api-Key` y `Api-Secret`, o un token Glue de owner cuando el propietario de la base de datos está permitido.

## Solicitud HTTP

```http
GET /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
```

## Respuesta

```json
[
  {
    "name": "users",
    "total": 5400,
    "indexes": 2,
    "defaults": {
      "id": "uuid()"
    }
  }
]
```
