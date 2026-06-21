
# Lambda API Reference

Base URL: `https://api.hola.cloud`

## Authentication

Management endpoints require `Api-Key` and `Api-Secret` headers. Public invocation endpoints (`/run/{id}` and `/mux/{owner}/*`) do not require authentication.

All requests must be sent over HTTPS.

## Endpoints

| Method | Path | Description |
|--------|------|-------------|
| GET | `/api/v0/lambdas` | List all lambdas |
| POST | `/api/v0/lambdas` | Create a new lambda |
| GET | `/api/v0/lambdas/{id}` | Get a lambda by ID |
| PATCH | `/api/v0/lambdas/{id}` | Modify a lambda |
| DELETE | `/api/v0/lambdas/{id}` | Delete a lambda |
| POST | `/api/v0/run/{id}` | Run a lambda (admin) |
| POST | `/run/{id}` | Run a lambda (public) |
| GET | `/mux/{owner}/*` | Route to owner's lambdas |

## Common Error Codes

| Code | Description |
|------|-------------|
| 400 | Bad request — malformed syntax or invalid parameters |
| 401 | Unauthorized — missing or invalid API credentials |
| 403 | Forbidden — credentials do not have access to the resource |
| 404 | Not found — the requested resource does not exist |
| 409 | Conflict — resource already exists |
| 500 | Internal server error |
