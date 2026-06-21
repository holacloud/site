
# PATCH /api/v0/lambdas/{id}

Updates the code, runtime, or active status of an existing lambda function.

## Authentication

Requires `Api-Key` and `Api-Secret` headers.

## Path Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| id | uuid | The unique identifier of the lambda |

## Request Body

| Field | Type | Description |
|-------|------|-------------|
| name | string | New name (optional) |
| runtime | string | New runtime (optional) |
| code | string | New source code (optional) |
| active | boolean | Whether the lambda is active (optional) |

## HTTP Request

```http
PATCH /api/v0/lambdas/f1a2b3c4-d5e6-7890-abcd-ef0123456789 HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
Content-Type: application/json

{
  "code": "export default async (req) => { return { status: 200, body: { message: 'Updated lambda' } }; }",
  "active": false
}
```

## Example

```bash
curl -X PATCH "https://api.hola.cloud/api/v0/lambdas/f1a2b3c4-d5e6-7890-abcd-ef0123456789" \
  -H "Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d" \
  -H "Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf" \
  -H "Content-Type: application/json" \
  -d '{
    "code": "export default async (req) => { return { status: 200, body: { message: \"Updated lambda\" } }; }",
    "active": false
  }'
```

## Response

```json
{
  "id": "f1a2b3c4-d5e6-7890-abcd-ef0123456789",
  "name": "hello-world",
  "runtime": "javascript",
  "active": false,
  "updated_at": "2025-07-02T09:15:00Z"
}
```

## Error Codes

| Code | Description |
|------|-------------|
| 400 | Invalid request body |
| 401 | Missing or invalid authentication headers |
| 404 | Lambda not found |
