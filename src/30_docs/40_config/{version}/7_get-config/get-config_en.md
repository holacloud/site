# GET /v0/configs/{id}

Get a specific configuration by its ID. This endpoint is publicly accessible (admin API).

## Authentication

No authentication required. This is a public endpoint.

## Path Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| `id` | string | The config ID (e.g., `cfg_abc123`) |

## Request

```bash
curl "https://api.hola.cloud/v0/configs/cfg_abc123"
```

## Response

```http
HTTP/1.1 200 OK
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
| 404 | Not Found | The specified config does not exist |
| 500 | Internal Server Error | An unexpected error occurred |
