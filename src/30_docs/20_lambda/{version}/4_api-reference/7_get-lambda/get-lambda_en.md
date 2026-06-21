
# GET /api/v0/lambdas/{id}

Retrieves details about a specific lambda function by its ID.

## Authentication

Requires `Api-Key` and `Api-Secret` headers.

## Path Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| id | uuid | The unique identifier of the lambda |

## HTTP Request

```http
GET /api/v0/lambdas/f1a2b3c4-d5e6-7890-abcd-ef0123456789 HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
```

## Example

```bash
curl -X GET "https://api.hola.cloud/api/v0/lambdas/f1a2b3c4-d5e6-7890-abcd-ef0123456789" \
  -H "Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d" \
  -H "Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf"
```

## Response

```json
{
  "id": "f1a2b3c4-d5e6-7890-abcd-ef0123456789",
  "name": "hello-world",
  "runtime": "javascript",
  "active": true,
  "code": "export default async (req) => { return { status: 200, body: { message: 'Hello, World!' } }; }",
  "created_at": "2025-07-01T14:00:00Z",
  "updated_at": "2025-07-01T14:00:00Z"
}
```

## Error Codes

| Code | Description |
|------|-------------|
| 401 | Missing or invalid authentication headers |
| 404 | Lambda not found |
