# PATCH /v0/configs/{id}

Partially update a configuration. Only the fields provided in the request body are modified. This endpoint is publicly accessible (admin API).

## Authentication

No authentication required. This is a public endpoint.

## Path Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| `id` | string | The config ID (e.g., `cfg_abc123`) |

## Request Body

```json
{
  "data": {
    "features": {
      "new-checkout": true
    }
  }
}
```

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `project` | string | no | New project name |
| `environment` | string | no | New environment name |
| `data` | object | no | Partial configuration data to merge |

## Request

```bash
curl -X PATCH "https://api.hola.cloud/v0/configs/cfg_abc123" \
  -H "Content-Type: application/json" \
  -d '{
    "data": {
      "features": {
        "new-checkout": true
      }
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
  "id": "cfg_abc123",
  "project": "my-app",
  "environment": "production",
  "data": {
    "database": {
      "host": "db.example.com",
      "port": 5432
    },
    "features": {
      "new-checkout": true
    }
  },
  "createdAt": "2026-06-21T10:00:00Z",
  "updatedAt": "2026-06-21T11:00:00Z"
}
```

## Error Codes

| Status | Code | Description |
|--------|------|-------------|
| 400 | Bad Request | Invalid request body |
| 404 | Not Found | The specified config does not exist |
| 409 | Conflict | Update conflicts with an existing config |
| 500 | Internal Server Error | An unexpected error occurred |
