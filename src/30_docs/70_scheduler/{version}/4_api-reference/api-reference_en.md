# Scheduler API Reference

Base URL: `https://api.hola.cloud`

| Method | Path | Description | Auth |
|--------|------|-------------|------|
| GET | /schedulers | List schedulers | Public |
| POST | /schedulers | Create scheduler | Private |
| GET | /schedulers/{id} | Get scheduler | Public |
| PUT | /schedulers/{id} | Update scheduler | Private |
| PATCH | /schedulers/{id} | Partial update scheduler | Private |
| DELETE | /schedulers/{id} | Delete scheduler | Private |
| GET | /schedulers/{id}/health | Health check | Public |
| POST | /schedulers/{id}/tasks | Enqueue task | Private |
| GET | /schedulers/{id}/tasks | List tasks | Public |
| GET | /schedulers/{id}/tasks/stream | SSE task stream | Public |
| POST | /schedulers/{id}/tasks/reserve | Reserve task | Private |
| DELETE | /schedulers/{id}/tasks/{task} | Acknowledge task | Private |
| POST | /schedulers/{id}/tasks/{task}/extend | Extend lease | Private |
| POST | /schedulers/{id}/tasks/{task}/reschedule | Reschedule task | Private |

## Authentication

Private endpoints require an API key. Pass it using either:

- `X-API-Key` header: `X-API-Key: YOUR_API_KEY`
- `Authorization` header: `Authorization: Bearer YOUR_API_KEY`

Read endpoints (GET) are public and do not require authentication.

## Error Codes

| Status | Code | Description |
|--------|------|-------------|
| 400 | invalid_json | Invalid JSON payload |
| 400 | validation_error | Invalid request body or parameters |
| 401 | unauthorized | Missing or invalid API key |
| 404 | not_found | Scheduler or task not found |
| 404 | task_not_found | Task not found |
| 409 | conflict | Scheduler already exists |
| 409 | task_already_exists | Task already exists |
| 409 | task_in_flight | Task is currently leased |
| 500 | internal_error | Internal server error |
