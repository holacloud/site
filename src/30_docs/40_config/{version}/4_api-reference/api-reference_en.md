# Config API Reference

Base URL: `https://api.hola.cloud`

The Config service exposes two API surfaces:

- **Admin API** (`/v0/configs`) — publicly accessible for managing configurations
- **User API** (`/v1/config`) — requires `Api-Key` and `Api-Secret` headers for runtime configuration access

## Endpoints

| Method | Path | Description | Auth |
|--------|------|-------------|------|
| GET | `/v0/configs` | List all configs (admin) | Public |
| POST | `/v0/configs` | Create a new config (admin) | Public |
| GET | `/v0/configs/{id}` | Get a config by ID (admin) | Public |
| DELETE | `/v0/configs/{id}` | Delete a config (admin) | Public |
| PATCH | `/v0/configs/{id}` | Partial update of a config (admin) | Public |
| GET | `/v1/config` | Retrieve the current user config | API Key |
| PATCH | `/v1/config` | Update the current user config | API Key |

## Common Errors

| Status | Code | Description |
|--------|------|-------------|
| 400 | Bad Request | Invalid request body or parameters |
| 401 | Unauthorized | Missing or invalid API credentials |
| 404 | Not Found | The specified resource does not exist |
| 409 | Conflict | Resource already exists or operation conflicts |
| 500 | Internal Server Error | An unexpected error occurred |
