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

KVNode only checks that one of these header forms is present. Missing headers return `403 forbidden`.

## Error Codes

| Status | Code | Description |
|--------|------|-------------|
| 400 | invalid_json | Invalid JSON payload |
| 400 | missing_collection | Collection is required |
| 400 | invalid_project | Project id contains invalid characters |
| 403 | forbidden | Missing authentication headers |
| 404 | not_found | Collection or key not found |
| 404 | missing_collection | Collection does not exist |
| 502 | parent_unavailable | Parent node is not reachable |
| 500 | internal_error | Internal server error |
