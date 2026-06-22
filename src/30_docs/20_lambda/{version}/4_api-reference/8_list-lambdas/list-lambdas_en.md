# List Lambdas

Lists lambdas belonging to the authenticated account.

## Authentication

Requires `X-Glue-Authentication`.

## HTTP Request

```http
GET /api/v0/lambdas HTTP/1.1
Host: api.hola.cloud
X-Glue-Authentication: YOUR_TOKEN
```

## Example

```bash
curl -X GET "https://api.hola.cloud/api/v0/lambdas" \
  -H "X-Glue-Authentication: YOUR_TOKEN"
```

## Response

```json
[
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
  },
  {
    "id": "a2b3c4d5-e6f7-8901-bcde-f12345678901",
    "created_timestamp": 1750408200,
    "owner": "user_123",
    "project_id": "project_456",
    "name": "site-style",
    "language": "static-css",
    "code": "body { color: #111; }",
    "method": "GET",
    "path": "/site.css"
  }
]
```

## Error Codes

| Code | Description |
|------|-------------|
| 401 | Missing or invalid authentication |
