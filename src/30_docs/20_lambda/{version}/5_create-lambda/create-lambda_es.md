
# POST /api/v0/lambdas

Crea una nueva función lambda con el código y la configuración de ejecución especificados.

## Autenticación

Requiere las cabeceras `Api-Key` y `Api-Secret`.

## Cuerpo de la Solicitud

| Campo | Tipo | Descripción |
|-------|------|-------------|
| name | string | Un nombre legible para la lambda |
| runtime | string | Entorno de ejecución: `javascript`, `python` o `go` |
| code | string | El código fuente de la función |
| active | boolean | Si la lambda está activa (opcional, por defecto `true`) |

## Solicitud HTTP

```http
POST /api/v0/lambdas HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
Content-Type: application/json

{
  "name": "hello-world",
  "runtime": "javascript",
  "active": true,
  "code": "export default async (req) => { return { status: 200, body: { message: 'Hello, World!' } }; }"
}
```

## Ejemplo

```bash
curl -X POST "https://api.hola.cloud/api/v0/lambdas" \
  -H "Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d" \
  -H "Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "hello-world",
    "runtime": "javascript",
    "active": true,
    "code": "export default async (req) => { return { status: 200, body: { message: \"Hello, World!\" } }; }"
  }'
```

## Respuesta

```json
{
  "id": "f1a2b3c4-d5e6-7890-abcd-ef0123456789",
  "name": "hello-world",
  "runtime": "javascript",
  "active": true,
  "created_at": "2025-07-01T14:00:00Z"
}
```

## Códigos de Error

| Código | Descripción |
|--------|-------------|
| 400 | Campos obligatorios faltantes o inválidos |
| 401 | Cabeceras de autenticación faltantes o inválidas |
| 409 | Ya existe una lambda con el mismo nombre |
