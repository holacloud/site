# POST /v0/configs

Create a new configuration. This endpoint is publicly accessible (admin API).

## Authentication

No authentication required. This is a public endpoint.

## Request Body

```json
{
  "project": "my-app",
  "environment": "production",
  "data": {
    "database": {
      "host": "db.example.com",
      "port": 5432
    },
    "features": {
      "new-checkout": false
    }
  }
}
```

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `project` | string | yes | Project name identifier |
| `environment` | string | yes | Environment name (e.g., production, staging, development) |
| `data` | object | yes | Configuration data as a JSON object |

## Request

```bash
curl -X POST "https://api.hola.cloud/v0/configs" \
  -H "Content-Type: application/json" \
  -d '{
    "project": "my-app",
    "environment": "production",
    "data": {
      "database": {
        "host": "db.example.com",
        "port": 5432
      },
      "features": {
        "new-checkout": false
      }
    }
  }'
```

## Response

```http
HTTP/1.1 201 Created
Content-Type: application/json
```

```json
{
  "id": "cfg_abc123",
  "project": "my-app",
  "environment": "production",
  "data": {
    "database": {
      "host": "db.example.com",
      "port": 5432
    },
    "features": {
      "new-checkout": false
    }
  },
  "createdAt": "2026-06-21T10:00:00Z",
  "updatedAt": "2026-06-21T10:00:00Z"
}
```

## Error Codes

| Status | Code | Description |
|--------|------|-------------|
| 400 | Bad Request | Missing required fields or invalid JSON |
| 409 | Conflict | A config with the same project and environment already exists |
| 500 | Internal Server Error | An unexpected error occurred |
