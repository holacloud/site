# Files API Reference

Base URL: `https://api.hola.cloud`

All endpoints require authentication via the `Api-Key` and `Api-Secret` headers.

## Endpoints

| Method | Path | Description |
|--------|------|-------------|
| GET | `/v1/buckets` | List all buckets |
| POST | `/v1/buckets` | Create a new bucket |
| GET | `/v1/buckets/{id}` | Get bucket details |
| PATCH | `/v1/buckets/{id}` | Modify bucket metadata |
| DELETE | `/v1/buckets/{id}` | Delete a bucket |
| GET | `/v1/buckets/{id}/list/*` | List files in a bucket |
| PUT | `/v1/buckets/{id}/files/*` | Upload a file |
| GET | `/v1/buckets/{id}/files/*` | Download a file |
| DELETE | `/v1/buckets/{id}/files/*` | Delete a file |
| HEAD | `/v1/buckets/{id}/files/*` | Get file metadata |

## Common Errors

| Status | Code | Description |
|--------|------|-------------|
| 400 | Bad Request | Invalid request body or parameters |
| 401 | Unauthorized | Missing or invalid API credentials |
| 403 | Forbidden | Insufficient permissions |
| 404 | Not Found | The specified resource does not exist |
| 409 | Conflict | Resource already exists or operation conflicts |
| 413 | Payload Too Large | File exceeds maximum size limit (5 GB) |
| 429 | Too Many Requests | Rate limit exceeded |
| 500 | Internal Server Error | An unexpected error occurred |
