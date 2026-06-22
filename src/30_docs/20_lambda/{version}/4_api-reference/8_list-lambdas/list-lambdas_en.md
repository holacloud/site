
# List Lambdas

Lists all lambda functions belonging to the authenticated account.

## Authentication

Requires `Api-Key` and `Api-Secret` headers.

## HTTP Request

```http
GET /api/v0/lambdas HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
```

## Example

```bash
curl -X GET "https://api.hola.cloud/api/v0/lambdas" \
  -H "Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d" \
  -H "Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf"
```

## Response

```json
[
  {
    "id": "f1a2b3c4-d5e6-7890-abcd-ef0123456789",
    "name": "hello-world",
    "runtime": "javascript",
    "active": true,
    "created_at": "2025-05-10T12:00:00Z"
  },
  {
    "id": "a2b3c4d5-e6f7-8901-bcde-f12345678901",
    "name": "process-orders",
    "runtime": "python",
    "active": true,
    "created_at": "2025-06-20T08:30:00Z"
  }
]
```

## Error Codes

| Code | Description |
|------|-------------|
| 401 | Missing or invalid authentication headers |
