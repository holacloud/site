# KVNode API Reference

Base URL: `https://api.hola.cloud`

| Method | Path | Description | Auth |
|--------|------|-------------|------|
| GET | /healthz | Health check | Public |
| GET | /readyz | Readiness check | Public |
| GET | /v1/status | Node status | Internal |
| GET | /v1/metrics | Node metrics | Internal |
| GET | /v1/collections | List collections | Internal |
| POST | /v1/collections | Create collection | Internal |
| DELETE | /v1/collections/{col} | Delete collection | Internal |
| GET | /v1/collections/{col}/keys | List keys | Internal |
| POST | /v1/collections/{col}/keys/* | Set key | Internal |
| GET | /v1/collections/{col}/keys/* | Get key | Internal |
| DELETE | /v1/collections/{col}/keys/* | Delete key | Internal |

## Authentication

Health check endpoints are public. All other endpoints require authentication via one of:

- `X-Glue-Authentication` header
- `apikey` and `secret` headers

## Error Codes

| Status | Code | Description |
|--------|------|-------------|
| 400 | invalid_request | Invalid request body or parameters |
| 401 | unauthorized | Missing or invalid authentication |
| 404 | not_found | Collection or key not found |
| 409 | conflict | Collection already exists |
| 500 | internal_error | Internal server error |
