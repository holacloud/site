# Run Lambda (Admin)

Invokes a lambda through the authenticated admin route. The endpoint accepts any HTTP method.

## Authentication

Requires `X-Glue-Authentication`.

## Path Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| `lambda_id` | string | Lambda identifier |

## Request Body

Send any payload you want the lambda to receive as `req.body`.

## HTTP Request

```http
POST /api/v0/run/f1a2b3c4-d5e6-7890-abcd-ef0123456789 HTTP/1.1
Host: api.hola.cloud
X-Glue-Authentication: YOUR_TOKEN
Content-Type: application/json

{
  "name": "Alice"
}
```

## Example

```bash
curl -X POST "https://api.hola.cloud/api/v0/run/f1a2b3c4-d5e6-7890-abcd-ef0123456789" \
  -H "X-Glue-Authentication: YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Alice"
  }'
```

## Response

The response is whatever the lambda handler returns.

```json
{
  "body": {
    "message": "Hello, Alice!"
  }
}
```

## Error Codes

| Code | Description |
|------|-------------|
| 400 | Invalid request body |
| 401 | Missing or invalid authentication |
| 404 | Lambda not found |
| 500 | Lambda execution error |
