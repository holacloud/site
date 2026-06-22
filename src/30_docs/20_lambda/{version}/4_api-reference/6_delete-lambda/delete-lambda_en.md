# Delete Lambda

Permanently deletes a lambda.

## Authentication

Requires `X-Glue-Authentication`.

## Path Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| `lambda_id` | string | Lambda identifier |

## HTTP Request

```http
DELETE /api/v0/lambdas/f1a2b3c4-d5e6-7890-abcd-ef0123456789 HTTP/1.1
Host: api.hola.cloud
X-Glue-Authentication: YOUR_TOKEN
```

## Example

```bash
curl -X DELETE "https://api.hola.cloud/api/v0/lambdas/f1a2b3c4-d5e6-7890-abcd-ef0123456789" \
  -H "X-Glue-Authentication: YOUR_TOKEN"
```

## Response

```json
{
  "id": "f1a2b3c4-d5e6-7890-abcd-ef0123456789"
}
```

## Error Codes

| Code | Description |
|------|-------------|
| 401 | Missing or invalid authentication |
| 404 | Lambda not found |
