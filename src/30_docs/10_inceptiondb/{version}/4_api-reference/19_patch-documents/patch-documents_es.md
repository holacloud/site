# Modificar documentos

Modifica documentos seleccionados mediante escaneo completo o recorrido de índice. El endpoint devuelve cada documento modificado como JSON Lines.

## Autenticación

Requiere `Api-Key` y `Api-Secret`, o un token Glue de owner cuando el propietario de la base de datos está permitido.

## Cuerpo de la solicitud

| Campo | Tipo | Descripción |
|-------|------|-------------|
| `patch` | object | Objeto patch que se aplica a cada documento coincidente. |
| `index` | string | Nombre de índice opcional. Omítelo para escanear toda la colección. |
| `filter` | object | Filtro Connor opcional aplicado antes del patch. |
| `skip` | integer | Número de documentos coincidentes que se omiten. Default: `0`. |
| `limit` | integer | Número máximo de documentos a modificar. Default: `1`; `0` no modifica ninguno. |

También se pueden incluir campos específicos del índice en el cuerpo.

## Solicitud HTTP

```http
POST /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections/users:patch HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
Content-Type: application/json

{
  "filter": {"id": "user-1"},
  "patch": {"role": "owner"},
  "limit": 1
}
```

## Respuesta

```jsonl
{"id":"user-1","name":"Alice","role":"owner"}
```
