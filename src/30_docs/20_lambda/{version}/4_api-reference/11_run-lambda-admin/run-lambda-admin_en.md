
# Run Lambda (Admin)

Invokes a lambda function using admin credentials. Requires authentication.

## Authentication

Requires `Api-Key` and `Api-Secret` headers.

## Path Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| id | uuid | The unique identifier of the lambda to run |

## Request Body

Any JSON payload you want to pass to the lambda function. The lambda receives it as the `req` parameter.

## HTTP Request

```http
POST /api/v0/run/f1a2b3c4-d5e6-7890-abcd-ef0123456789 HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
Content-Type: application/json

{
  "name": "Alice"
}
```

## Example

```bash
curl -X POST "https://api.hola.cloud/api/v0/run/f1a2b3c4-d5e6-7890-abcd-ef0123456789" \
  -H "Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d" \
  -H "Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Alice"
  }'
```

## Response

```json
{
  "status": 200,
  "body": {
    "message": "Hello, Alice!"
  }
}
```

## Error Codes

| Code | Description |
|------|-------------|
| 400 | Invalid request body |
| 401 | Missing or invalid authentication headers |
| 404 | Lambda not found or inactive |
| 500 | Lambda execution error |
