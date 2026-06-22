# List Buckets

List buckets owned by the authenticated user.

## Authentication

Requires `X-Glue-Authentication`.

## Request

```bash
curl "https://api.hola.cloud/v1/buckets" \
  -H 'X-Glue-Authentication: {"user":{"id":"user-123"}}'
```

## Response

The response is a JSON array.

```json
[
  {
    "id": "bucket-550e8400-e29b-41d4-a716-446655440000",
    "name": "assets",
    "description": "Application assets",
    "created_timestamp": 1782045600000000000,
    "created_h": "2026-06-21T10:00:00Z",
    "owners": ["user-123"]
  }
]
```

## Error Codes

| Status | Description |
|--------|-------------|
| 401 | Missing or invalid `X-Glue-Authentication` |
| 500 | Persistence error |
