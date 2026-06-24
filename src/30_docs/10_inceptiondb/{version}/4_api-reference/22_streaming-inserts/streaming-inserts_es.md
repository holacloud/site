# Inserciones por stream

`insertStream` e `insertFullduplex` son endpoints experimentales de inserción para entrada JSON en stream.

## Autenticación

Requiere `Api-Key` y `Api-Secret`, o un token Glue de owner cuando el propietario de la base de datos está permitido.

## Insert Stream

`insertStream` toma control de la conexión HTTP y responde con `202 Accepted` usando salida chunked.

```http
POST /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections/events:insertStream HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
Content-Type: application/jsonl

{"id":"evt-1","type":"signup"}
{"id":"evt-2","type":"login"}
```

## Insert Full Duplex

`insertFullduplex` activa HTTP full duplex y codifica cada documento insertado como JSON mientras lee el cuerpo de la solicitud.

```http
POST /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections/events:insertFullduplex HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
Content-Type: application/jsonl

{"id":"evt-1","type":"signup"}
{"id":"evt-2","type":"login"}
```

El cuerpo de respuesta de ambos endpoints es JSON Lines con los documentos insertados, o strings de error por documento en `insertStream`.
