
# InceptionDB API Reference

Base URL: `https://api.hola.cloud`

## Authentication

InceptionDB management endpoints require authentication via two headers:

- `Api-Key` — Your API key (UUID)
- `Api-Secret` — Your API secret (UUID)
- `X-Project` — Scopes the request to a specific project

All requests must be sent over HTTPS.

## Endpoints

| Method | Path | Description |
|--------|------|-------------|
| GET | `/v1/databases` | List all databases |
| POST | `/v1/databases` | Create a new database |
| GET | `/v1/databases/{id}` | Get a database by ID |
| DELETE | `/v1/databases/{id}` | Drop a database |
| PATCH | `/v1/databases/{id}` | Update a database |
| GET | `/v1/databases/{id}/collections` | List collections in a database |
| POST | `/v1/databases/{id}/collections` | Create a collection |
| POST | `/v1/databases/{id}/collections/{col}` | Document operations (insert, find, remove, patch) |

## Common Error Codes

| Code | Description |
|------|-------------|
| 400 | Bad request — malformed syntax or invalid parameters |
| 401 | Unauthorized — missing or invalid API credentials |
| 403 | Forbidden — credentials do not have access to the resource |
| 404 | Not found — the requested resource does not exist |
| 409 | Conflict — resource already exists |
| 500 | Internal server error |
