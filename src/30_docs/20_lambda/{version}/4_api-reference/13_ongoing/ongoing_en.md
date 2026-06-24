# List Ongoing Invocations

Returns a list of currently running lambda invocations for the authenticated account.

## Authentication

Requires `X-Glue-Authentication`.

## HTTP Request

```http
GET /api/v0/ongoing HTTP/1.1
Host: api.hola.cloud
X-Glue-Authentication: YOUR_TOKEN
```

## Example

```bash
curl -X GET "https://api.hola.cloud/api/v0/ongoing" \
  -H "X-Glue-Authentication: YOUR_TOKEN"
```

## Response

```json
[
  {
    "id": "inv_abc123",
    "lambda_id": "f1a2b3c4-d5e6-7890-abcd-ef0123456789",
    "lambda_name": "hello-world",
    "started_at": "2025-07-01T12:00:00Z",
    "method": "POST",
    "path": "/hello-world"
  }
]
```

## Error Codes

| Code | Description |
|------|-------------|
| 401 | Missing or invalid authentication |
