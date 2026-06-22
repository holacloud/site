
# Run Lambda (Public)

Invokes a lambda function publicly. No authentication required.

This endpoint is designed for webhooks and public API endpoints. The lambda ID can be obtained from the lambda management pages.

## Path Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| id | uuid | The unique identifier of the lambda to run |

## Request Body

Any JSON payload you want to pass to the lambda function.

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

```json
{
  "status": 200,
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
| 404 | Lambda not found or inactive |
| 500 | Lambda execution error |
