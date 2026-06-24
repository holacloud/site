# Gestión de colecciones

Estos endpoints de acción administran una colección existente dentro de una base de datos.

## Autenticación

Requiere `Api-Key` y `Api-Secret`, o un token Glue de owner cuando el propietario de la base de datos está permitido.

## Obtener colección

```http
GET /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections/users HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
```

Respuesta:

```json
{"name":"users","total":5400,"indexes":2,"defaults":{"id":"uuid()"}}
```

## Eliminar colección

```http
POST /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections/users:dropCollection HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
```

Las respuestas correctas tienen cuerpo vacío.

## Definir defaults

Los defaults se mezclan con documentos nuevos al insertarlos. Envía `null` para eliminar el default de un campo.

```http
POST /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections/users:setDefaults HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
Content-Type: application/json

{"id":"uuid()","status":"active","legacy":null}
```

Respuesta:

```json
{"id":"uuid()","status":"active"}
```

## Tamaño

`size` es experimental y devuelve métricas de almacenamiento de la colección.

```http
POST /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections/users:size HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
```

Respuesta:

```json
{"memory":123456,"disk":234567,"index.id":12345}
```
