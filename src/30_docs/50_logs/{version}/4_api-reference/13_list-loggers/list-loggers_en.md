
# GET /v1/loggers

List all loggers associated with your account.

## Authentication

Requires management credentials:

- `Api-Key` — Your API key
- `Api-Secret` — Your API secret

## Request

```bash
curl "https://api.hola.cloud/v1/loggers" \
  -H "Api-Key: your-api-key" \
  -H "Api-Secret: your-api-secret"
```

```http
GET /v1/loggers HTTP/1.1
Host: api.hola.cloud
Api-Key: your-api-key
Api-Secret: your-api-secret
```

## Response

```json
[
  {
    "id": "logger_xyz789",
    "name": "my-app-logger",
    "created_at": "2025-06-20T14:22:00Z"
  },
  {
    "id": "logger_abc123",
    "name": "staging-logger",
    "created_at": "2025-06-19T10:00:00Z"
  }
]
```

## Error Codes

| Code | Description |
|------|-------------|
| 401 | Missing or invalid API credentials |
| 403 | API credentials do not have listing permissions |
