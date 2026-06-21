# GET /v1/collections/{col}/keys

Lists keys in a collection with optional pagination and prefix filtering.

## Authentication

Requires internal authentication. Pass credentials via `X-Glue-Authentication` header, or `apikey` and `secret` headers.

## Path Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| col | string | The name of the collection |

## Query Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| limit | integer | Maximum number of keys to return (default: 100) |
| prefix | string | Filter keys by prefix |

## Example Request

```bash
curl "https://api.hola.cloud/v1/collections/users/keys?prefix=user:&limit=10" \
  -H "X-Glue-Authentication: YOUR_AUTH_TOKEN"
```

## Example Response

```http
HTTP/1.1 200 OK
Content-Type: application/json
```

```json
{
  "keys": [
    {
      "key": "user:1001",
      "created_at": "2025-06-15T10:00:00Z",
      "updated_at": "2025-06-20T15:30:00Z"
    },
    {
      "key": "user:1002",
      "created_at": "2025-06-16T09:00:00Z",
      "updated_at": "2025-06-21T08:00:00Z"
    }
  ],
  "total": 2
}
```

## Error Codes

| Status | Code | Description |
|--------|------|-------------|
| 401 | unauthorized | Missing or invalid authentication |
| 404 | not_found | Collection not found |
| 500 | internal_error | Internal server error |
