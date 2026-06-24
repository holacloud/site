# Obtener documento

Obtiene un documento por `id`. La búsqueda usa un índice `map` sobre el campo `id` cuando existe; si no, cae a un escaneo de colección.

## Autenticación

Requiere `Api-Key` y `Api-Secret`, o un token Glue de owner cuando el propietario de la base de datos está permitido.

## Solicitud HTTP

```http
GET /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections/users/documents/user-1 HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
```

## Respuesta

```json
{
  "id": "user-1",
  "document": {"id":"user-1","name":"Alice","role":"admin"},
  "source": {"type":"index","name":"id"}
}
```

`source` se omite cuando el endpoint usa escaneo de colección. IDs vacíos devuelven `400 Bad Request`; colecciones o documentos inexistentes devuelven `404 Not Found`.
