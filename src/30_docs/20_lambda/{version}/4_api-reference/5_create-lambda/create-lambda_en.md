# Create Lambda

Creates a lambda with source code or static content and route metadata.

## Authentication

Requires `X-Glue-Authentication`.

## Request Body

| Field | Type | Description |
|-------|------|-------------|
| `name` | string | Human-readable lambda name |
| `language` | string | `javascript`, `static-html`, `static-css`, or `static-js` |
| `code` | string | Source code or static content |
| `method` | string | HTTP method for the lambda route |
| `path` | string | HTTP path for the lambda route |

## HTTP Request

```http
POST /api/v0/lambdas HTTP/1.1
Host: api.hola.cloud
X-Glue-Authentication: YOUR_TOKEN
Content-Type: application/json

{
  "name": "hello-world",
  "language": "javascript",
  "method": "GET",
  "path": "/hello-world",
  "code": "export default (req) => ({ body: { message: 'Hello, World!' } })"
}
```

## Example

```bash
curl -X POST "https://api.hola.cloud/api/v0/lambdas" \
  -H "X-Glue-Authentication: YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "hello-world",
    "language": "javascript",
    "method": "GET",
    "path": "/hello-world",
    "code": "export default (req) => ({ body: { message: \"Hello, World!\" } })"
  }'
```

## Response

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

## Error Codes

| Code | Description |
|------|-------------|
| 400 | Missing or invalid required fields |
| 401 | Missing or invalid authentication |
| 409 | A lambda with the same name already exists |
