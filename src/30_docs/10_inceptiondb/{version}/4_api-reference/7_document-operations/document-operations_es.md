# Operaciones de documentos

Las operaciones de documentos son endpoints de acción sobre una colección.

## Autenticación

Requiere `Api-Key` y `Api-Secret`, o un token Glue de owner cuando el propietario de la base de datos está permitido.

## Insertar: `{collection}:insert`

```http
POST /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections/users:insert HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
Content-Type: application/jsonl

{"id":"user-1","name":"Alice","role":"admin"}
{"id":"user-2","name":"Bob","role":"member"}
```

```jsonl
{"id":"user-1","name":"Alice","role":"admin"}
{"id":"user-2","name":"Bob","role":"member"}
```

## Buscar: `{collection}:find`

```http
POST /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections/users:find HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
Content-Type: application/json

{
  "filter": {
    "role": "admin"
  }
}
```

```jsonl
{"id":"user-1","name":"Alice","role":"admin"}
```

## Obtener por ID

```http
GET /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections/users/documents/user-1 HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
```

```json
{"id":"user-1","name":"Alice","role":"admin"}
```

## Modificar: `{collection}:patch`

```http
POST /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections/users:patch HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
Content-Type: application/json

{
  "filter": { "id": "user-1" },
  "patch": { "role": "owner" }
}
```

```jsonl
{"id":"user-1","name":"Alice","role":"owner"}
```

## Eliminar: `{collection}:remove`

```http
POST /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections/users:remove HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
Content-Type: application/json

{
  "filter": { "id": "user-2" }
}
```

```jsonl
{"id":"user-2","name":"Bob","role":"member"}
```
