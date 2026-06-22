# Run Lambda (Public)

Invokes a lambda publicly. No authentication is required. The endpoint accepts any HTTP method.

Use this route for webhooks and public API calls. The lambda ID comes from the lambda management endpoints.

## Path Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| `lambda_id` | string | Lambda identifier |

## Request Body

Send any payload you want the lambda to receive as `req.body`.

## HTTP Request

```http
POST /run/f1a2b3c4-d5e6-7890-abcd-ef0123456789 HTTP/1.1
Host: api.hola.cloud
Content-Type: application/json

{
  "event": "payment_received",
  "amount": 49.99,
  "currency": "USD"
}
```

## Example

```bash
curl -X POST "https://api.hola.cloud/run/f1a2b3c4-d5e6-7890-abcd-ef0123456789" \
  -H "Content-Type: application/json" \
  -d '{
    "event": "payment_received",
    "amount": 49.99,
    "currency": "USD"
  }'
```

## Response

The response is whatever the lambda handler returns.

```json
{
  "body": {
    "processed": true,
    "event": "payment_received"
  }
}
```

## Error Codes

| Code | Description |
|------|-------------|
| 400 | Invalid request body |
| 404 | Lambda not found |
| 500 | Lambda execution error |
