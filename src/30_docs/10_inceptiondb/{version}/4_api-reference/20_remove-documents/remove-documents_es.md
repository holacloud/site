# Eliminar documentos

Elimina documentos seleccionados mediante escaneo completo o recorrido de índice. El endpoint devuelve cada documento eliminado como JSON Lines.

## Autenticación

Requiere `Api-Key` y `Api-Secret`, o un token Glue de owner cuando el propietario de la base de datos está permitido.

## Cuerpo de la solicitud

| Campo | Tipo | Descripción |
|-------|------|-------------|
| `index` | string | Nombre de índice opcional. Omítelo para escanear toda la colección. |
| `filter` | object | Filtro Connor opcional aplicado antes de eliminar. |
| `skip` | integer | Número de documentos coincidentes que se omiten. Default: `0`. |
| `limit` | integer | Número máximo de documentos a eliminar. Default: `1`; `0` no elimina ninguno. |

También se pueden incluir campos específicos del índice en el cuerpo.

## Solicitud HTTP

```http
POST /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections/users:remove HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
Content-Type: application/json

{"filter":{"id":"user-2"},"limit":1}
```

## Respuesta

```jsonl
{"id":"user-2","name":"Bob","role":"member"}
```
