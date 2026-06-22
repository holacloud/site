# List Keys

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
  "items": [
    {
      "key": "user:1001",
      "value": { "name": "Alice" },
      "version": 1,
      "updatedAt": "2025-06-20T15:30:00Z"
    },
    {
      "key": "user:1002",
      "value": { "name": "Bob" },
      "version": 1,
      "updatedAt": "2025-06-21T08:00:00Z"
    }
  ],
  "next": "user:1002"
}
```

## Error Codes

| Status | Code | Description |
|--------|------|-------------|
| 403 | forbidden | Missing authentication headers |
| 500 | internal_error | Internal server error |
