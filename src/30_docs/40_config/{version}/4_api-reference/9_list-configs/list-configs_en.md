# List Configs

List all configurations. This endpoint is publicly accessible (admin API).

## Authentication

No authentication required. This is a public endpoint.

## Query Parameters

| Parameter | Type | Default | Description |
|-----------|------|---------|-------------|
| `limit` | integer | 100 | Maximum number of configs to return |
| `offset` | integer | 0 | Number of configs to skip |
| `project` | string | "" | Filter configs by project name |

## Request

```bash
curl "https://api.hola.cloud/v0/configs?limit=10&offset=0"
```

## Response

```http
HTTP/1.1 200 OK
Content-Type: application/json
```

```json
{
  "configs": [
    {
      "id": "cfg_abc123",
      "project": "my-app",
      "environment": "production",
      "data": {
        "database": {
          "host": "db.example.com",
          "port": 5432
        }
      },
      "createdAt": "2026-06-21T10:00:00Z",
      "updatedAt": "2026-06-21T10:00:00Z"
    }
  ],
  "total": 1
}
```

## Error Codes

| Status | Code | Description |
|--------|------|-------------|
| 500 | Internal Server Error | An unexpected error occurred |
