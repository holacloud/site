# Modificar Lambda

Actualiza campos soportados de una lambda existente.

## Autenticación

Requiere `X-Glue-Authentication`.

## Parámetros de Path

| Parámetro | Tipo | Descripción |
|-----------|------|-------------|
| `lambda_id` | string | Identificador de la lambda |

## Cuerpo de la Solicitud

| Campo | Tipo | Descripción |
|-------|------|-------------|
| `name` | string | Nuevo nombre de la lambda |
| `language` | string | `javascript`, `static-html`, `static-css` o `static-js` |
| `code` | string | Nuevo código fuente o contenido estático |
| `method` | string | Nuevo método HTTP |
| `path` | string | Nuevo path HTTP |

## Solicitud HTTP

```http
PATCH /api/v0/lambdas/f1a2b3c4-d5e6-7890-abcd-ef0123456789 HTTP/1.1
Host: api.hola.cloud
X-Glue-Authentication: TU_TOKEN
Content-Type: application/json

{
  "name": "hello-updated",
  "method": "POST",
  "path": "/hello-updated",
  "code": "export default (req) => ({ body: { message: 'Updated lambda', data: req.body } })"
}
```

## Ejemplo

```bash
curl -X PATCH "https://api.hola.cloud/api/v0/lambdas/f1a2b3c4-d5e6-7890-abcd-ef0123456789" \
  -H "X-Glue-Authentication: TU_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "hello-updated",
    "method": "POST",
    "path": "/hello-updated",
    "code": "export default (req) => ({ body: { message: \"Updated lambda\", data: req.body } })"
  }'
```

## Respuesta

```json
{
  "id": "f1a2b3c4-d5e6-7890-abcd-ef0123456789",
  "created_timestamp": 1751378400,
  "owner": "user_123",
  "project_id": "project_456",
  "name": "hello-updated",
  "language": "javascript",
  "code": "export default (req) => ({ body: { message: \"Updated lambda\", data: req.body } })",
  "method": "POST",
  "path": "/hello-updated"
}
```

## Códigos de Error

| Código | Descripción |
|--------|-------------|
| 400 | Cuerpo de solicitud inválido |
| 401 | Autenticación faltante o inválida |
| 404 | Lambda no encontrada |
