
# InstantLogs API Reference

Base URL: `https://api.hola.cloud`

## Authentication

InstantLogs uses two authentication modes depending on the endpoint type.

### Management Authentication

Management endpoints (logger CRUD, API key management) require two headers:

| Header | Description |
|--------|-------------|
| `Api-Key` | Your API key identifier |
| `Api-Secret` | Your API secret key |

### Data Authentication

Data endpoints (ingest, filter, events, stats) authenticate using the logger secret. Provide it as either:

| Header | Description |
|--------|-------------|
| `X-Instantlogs-Event-Secret` | The logger secret value |
| `Authorization: Bearer <secret>` | Bearer token with the logger secret |

## Endpoints

| Method | Path | Description | Auth |
|--------|------|-------------|------|
| GET | `/v1/loggers` | List all loggers | Management |
| POST | `/v1/loggers` | Create a new logger | Management |
| GET | `/v1/loggers/{id}` | Get logger details | Management |
| DELETE | `/v1/loggers/{id}` | Delete a logger | Management |
| POST | `/v1/loggers/{id}/ingest` | Ingest log entries | Data |
| GET | `/v1/loggers/{id}/filter` | Stream and filter logs | Data |
| GET | `/v1/loggers/{id}/events` | Stream events | Data |
| GET | `/v1/loggers/{id}/stats` | Get logger statistics | Data |
| POST | `/v1/loggers/{id}/apiKeys` | Create an API key | Management |
| DELETE | `/v1/loggers/{id}/apiKeys/{key}` | Delete an API key | Management |
| POST | `/v1/ingest/events` | Ingest framed events | Data |

## Common Error Codes

| Code | Description |
|------|-------------|
| 400 | Bad request — malformed JSON or invalid parameters |
| 401 | Unauthorized — missing or invalid credentials |
| 403 | Forbidden — credentials do not have access |
| 404 | Not found — the requested resource does not exist |
| 409 | Conflict — resource already exists |
| 429 | Too many requests — rate limit exceeded |
| 500 | Internal server error |
