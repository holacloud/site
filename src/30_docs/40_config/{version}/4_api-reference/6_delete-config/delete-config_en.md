# DELETE /v0/configs/{id}

Delete a configuration by its ID. This endpoint is publicly accessible (admin API).

## Authentication

No authentication required. This is a public endpoint.

## Path Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| `id` | string | The config ID (e.g., `cfg_abc123`) |

## Request

```bash
curl -X DELETE "https://api.hola.cloud/v0/configs/cfg_abc123"
```

## Response

```http
HTTP/1.1 204 No Content
```

## Error Codes

| Status | Code | Description |
|--------|------|-------------|
| 404 | Not Found | The specified config does not exist |
| 500 | Internal Server Error | An unexpected error occurred |
