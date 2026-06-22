# Lambda API Reference

Base URL: `https://api.hola.cloud`

## Authentication

Management endpoints require the `X-Glue-Authentication` header. Public invocation endpoints (`/run/{lambda_id}` and `/mux/{owner_id}/*`) do not require authentication.

All requests must be sent over HTTPS.

## Lambda Object

| Field | Type | Description |
|-------|------|-------------|
| `id` | string | Lambda identifier |
| `created_timestamp` | number | Creation time as a Unix timestamp |
| `owner` | string | Owner identifier |
| `project_id` | string | Project identifier |
| `name` | string | Lambda name |
| `language` | string | One of `javascript`, `static-html`, `static-css`, or `static-js` |
| `code` | string | Source code or static content |
| `method` | string | HTTP method associated with the lambda |
| `path` | string | HTTP path associated with the lambda |

## Endpoints

| Method | Path | Description |
|--------|------|-------------|
| POST | `/api/v0/lambdas` | Create a lambda |
| GET | `/api/v0/lambdas` | List lambdas |
| GET | `/api/v0/lambdas/{lambda_id}` | Get a lambda by ID |
| PATCH | `/api/v0/lambdas/{lambda_id}` | Modify a lambda |
| DELETE | `/api/v0/lambdas/{lambda_id}` | Delete a lambda |
| ANY | `/api/v0/run/{lambda_id}` | Run a lambda with authentication |
| ANY | `/run/{lambda_id}` | Run a lambda publicly |
| ANY | `/mux/{owner_id}/*` | Route a request through an owner scope |
| GET | `/ongoing` | List currently running invocations |
| GET | `/me` | Return the authenticated user |
| GET | `/openapi` | Return the OpenAPI document |

## Common Error Codes

| Code | Description |
|------|-------------|
| 400 | Bad request |
| 401 | Missing or invalid authentication |
| 403 | Forbidden |
| 404 | Resource not found |
| 500 | Internal server error |
