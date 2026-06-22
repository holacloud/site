# Crear Lambda

Crea una lambda con código fuente o contenido estático y metadatos de ruta.

## Autenticación

Requiere `X-Glue-Authentication`.

## Cuerpo de la Solicitud

| Campo | Tipo | Descripción |
|-------|------|-------------|
| `name` | string | Nombre legible de la lambda |
| `language` | string | `javascript`, `static-html`, `static-css` o `static-js` |
| `code` | string | Código fuente o contenido estático |
| `method` | string | Método HTTP para la ruta de la lambda |
| `path` | string | Path HTTP para la ruta de la lambda |

## Solicitud HTTP

```http
POST /api/v0/lambdas HTTP/1.1
Host: api.hola.cloud
X-Glue-Authentication: TU_TOKEN
Content-Type: application/json

{
  "name": "hello-world",
  "language": "javascript",
  "method": "GET",
  "path": "/hello-world",
  "code": "export default (req) => ({ body: { message: 'Hello, World!' } })"
}
```

## Ejemplo

```bash
curl -X POST "https://api.hola.cloud/api/v0/lambdas" \
  -H "X-Glue-Authentication: TU_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "hello-world",
    "language": "javascript",
    "method": "GET",
    "path": "/hello-world",
    "code": "export default (req) => ({ body: { message: \"Hello, World!\" } })"
  }'
```

## Respuesta

```json
{
  "id": "f1a2b3c4-d5e6-7890-abcd-ef0123456789",
  "created_timestamp": 1751378400,
  "owner": "user_123",
  "project_id": "project_456",
  "name": "hello-world",
  "language": "javascript",
  "code": "export default (req) => ({ body: { message: \"Hello, World!\" } })",
  "method": "GET",
  "path": "/hello-world"
}
```

## Códigos de Error

| Código | Descripción |
|--------|-------------|
| 400 | Campos requeridos faltantes o inválidos |
| 401 | Autenticación faltante o inválida |
| 409 | Ya existe una lambda con el mismo nombre |
