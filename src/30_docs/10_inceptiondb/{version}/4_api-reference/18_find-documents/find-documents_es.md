# Buscar documentos

Busca documentos en una colección mediante escaneo completo o recorrido de índice.

## Autenticación

Requiere `Api-Key` y `Api-Secret`, o un token Glue de owner cuando el propietario de la base de datos está permitido.

## Cuerpo de la solicitud

| Campo | Tipo | Descripción |
|-------|------|-------------|
| `index` | string | Nombre de índice opcional. Omítelo para escanear toda la colección. |
| `filter` | object | Filtro Connor opcional aplicado a los documentos. |
| `skip` | integer | Número de documentos coincidentes que se omiten. Default: `0`. |
| `limit` | integer | Número máximo de documentos a devolver. Default: `1`; `0` no devuelve ninguno. |

También se pueden incluir campos específicos del índice en el cuerpo. Por ejemplo, un recorrido de índice `map` suele usar `value`.

## Solicitud HTTP

```http
POST /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections/users:find HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
Content-Type: application/json

{"filter":{"role":"admin"},"limit":10}
```

## Respuesta

La respuesta es JSON Lines con los documentos coincidentes.

```jsonl
{"id":"user-1","name":"Alice","role":"admin"}
```

## Búsqueda con índice

```json
{"index":"by-email","value":"alice@example.com","limit":1}
```
