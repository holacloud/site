# GET /v1/config

Retrieve the current user's configuration. This endpoint requires API key authentication.

## Authentication

Requires `Api-Key` and `Api-Secret` headers.

## Request

```bash
curl "https://api.hola.cloud/v1/config" \
  -H "Api-Key: YOUR_API_KEY" \
  -H "Api-Secret: YOUR_API_SECRET"
```

## Response

```http
HTTP/1.1 200 OK
Content-Type: application/json
```

```json
{
  "project": "my-app",
  "environment": "production",
  "services": {
    "database": {
      "host": "db.example.com",
      "port": 5432
    },
    "cache": {
      "host": "redis.example.com",
      "port": 6379
    }
  },
  "features": {
    "new-checkout": false
  }
}
```

## Error Codes

| Status | Code | Description |
|--------|------|-------------|
| 401 | Unauthorized | Missing or invalid API credentials |
| 500 | Internal Server Error | An unexpected error occurred |
