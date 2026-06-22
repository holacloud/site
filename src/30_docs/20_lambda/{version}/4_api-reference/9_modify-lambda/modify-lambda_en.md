# Modify Lambda

Updates supported fields on an existing lambda.

## Authentication

Requires `X-Glue-Authentication`.

## Path Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| `lambda_id` | string | Lambda identifier |

## Request Body

| Field | Type | Description |
|-------|------|-------------|
| `name` | string | New lambda name |
| `language` | string | `javascript`, `static-html`, `static-css`, or `static-js` |
| `code` | string | New source code or static content |
| `method` | string | New HTTP method |
| `path` | string | New HTTP path |

## HTTP Request

```http
PATCH /api/v0/lambdas/f1a2b3c4-d5e6-7890-abcd-ef0123456789 HTTP/1.1
Host: api.hola.cloud
X-Glue-Authentication: YOUR_TOKEN
Content-Type: application/json

{
  "name": "hello-updated",
  "method": "POST",
  "path": "/hello-updated",
  "code": "export default (req) => ({ body: { message: 'Updated lambda', data: req.body } })"
}
```

## Example

```bash
curl -X PATCH "https://api.hola.cloud/api/v0/lambdas/f1a2b3c4-d5e6-7890-abcd-ef0123456789" \
  -H "X-Glue-Authentication: YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "hello-updated",
    "method": "POST",
    "path": "/hello-updated",
    "code": "export default (req) => ({ body: { message: \"Updated lambda\", data: req.body } })"
  }'
```

## Response

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

## Error Codes

| Code | Description |
|------|-------------|
| 400 | Invalid request body |
| 401 | Missing or invalid authentication |
| 404 | Lambda not found |
