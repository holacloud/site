# Insertar documentos

Inserta uno o más documentos JSON en una colección. La colección se crea automáticamente si no existe.

## Autenticación

Requiere `Api-Key` y `Api-Secret`, o un token Glue de owner cuando el propietario de la base de datos está permitido.

## Solicitud HTTP

```http
POST /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections/users:insert HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
Content-Type: application/jsonl

{"id":"user-1","name":"Alice","role":"admin"}
{"id":"user-2","name":"Bob","role":"member"}
```

## Respuesta

Estado: `201 Created` cuando se inserta al menos un documento. La respuesta es JSON Lines con los documentos insertados.

```jsonl
{"id":"user-1","name":"Alice","role":"admin"}
{"id":"user-2","name":"Bob","role":"member"}
```

Si el cuerpo está vacío, la respuesta es `204 No Content`. Los conflictos de índices únicos devuelven `409 Conflict` antes de escribir cuerpo de respuesta.
