# Update User Config

Update the current user's configuration. Uses deep merge semantics: only the fields provided in the request body are modified. This endpoint requires API key authentication.

## Authentication

Requires `Api-Key` and `Api-Secret` headers.

## Request Body

```json
{
  "features": {
    "new-checkout": true
  }
}
```

Any subset of the configuration can be provided. The update is merged with the existing configuration.

## Request

```bash
curl -X PATCH "https://api.hola.cloud/v1/config" \
  -H "Api-Key: YOUR_API_KEY" \
  -H "Api-Secret: YOUR_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{
    "features": {
      "new-checkout": true
    }
  }'
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
    "new-checkout": true
  }
}
```

## Error Codes

| Status | Code | Description |
|--------|------|-------------|
| 400 | Bad Request | Invalid JSON in request body |
| 401 | Unauthorized | Missing or invalid API credentials |
| 500 | Internal Server Error | An unexpected error occurred |
