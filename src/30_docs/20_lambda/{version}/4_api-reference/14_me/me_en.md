# Get Current User

Returns information about the authenticated user.

## Authentication

Requires `X-Glue-Authentication`.

## HTTP Request

```http
GET /api/v0/me HTTP/1.1
Host: api.hola.cloud
X-Glue-Authentication: YOUR_TOKEN
```

## Example

```bash
curl -X GET "https://api.hola.cloud/api/v0/me" \
  -H "X-Glue-Authentication: YOUR_TOKEN"
```

## Response

```json
{
  "id": "user_123",
  "email": "user@example.com",
  "created_at": "2025-01-15T08:00:00Z"
}
```

## Error Codes

| Code | Description |
|------|-------------|
| 401 | Missing or invalid authentication |
