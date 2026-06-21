
# POST /api/v0/lambdas

Creates a new lambda function with the specified code and runtime configuration.

## Authentication

Requires `Api-Key` and `Api-Secret` headers.

## Request Body

| Field | Type | Description |
|-------|------|-------------|
| name | string | A human-readable name for the lambda |
| runtime | string | Runtime environment: `javascript`, `python`, or `go` |
| code | string | The source code of the function |
| active | boolean | Whether the lambda is active (optional, default `true`) |

## HTTP Request

```http
POST /api/v0/lambdas HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
Content-Type: application/json

{
  "name": "hello-world",
  "runtime": "javascript",
  "active": true,
  "code": "export default async (req) => { return { status: 200, body: { message: 'Hello, World!' } }; }"
}
```

## Example

```bash
curl -X POST "https://api.hola.cloud/api/v0/lambdas" \
  -H "Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d" \
  -H "Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "hello-world",
    "runtime": "javascript",
    "active": true,
    "code": "export default async (req) => { return { status: 200, body: { message: \"Hello, World!\" } }; }"
  }'
```

## Response

```json
{
  "id": "f1a2b3c4-d5e6-7890-abcd-ef0123456789",
  "name": "hello-world",
  "runtime": "javascript",
  "active": true,
  "created_at": "2025-07-01T14:00:00Z"
}
```

## Error Codes

| Code | Description |
|------|-------------|
| 400 | Missing or invalid required fields |
| 401 | Missing or invalid authentication headers |
| 409 | A lambda with the same name already exists |
