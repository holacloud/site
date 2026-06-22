# List Databases

Lists databases available to the authenticated Glue user.

## Authentication

Requires `X-Glue-Authentication`.

## HTTP Request

```http
GET /v1/databases HTTP/1.1
Host: api.hola.cloud
X-Glue-Authentication: your-glue-token
```

## Response

```json
[
  {
    "id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
    "name": "production-db",
    "creation_date": "2025-06-15T10:30:00Z",
    "owners_length": 1
  }
]
```
