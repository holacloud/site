
# Tailon API Reference

Base URL: `https://api.hola.cloud`

## Authentication

Tailon endpoints are **public** by default. No authentication headers are required for most operations.

## Endpoints

| Method | Path | Description |
|--------|------|-------------|
| GET | `/v1/queues` | List all queues |
| POST | `/v1/queues` | Create a new queue |
| GET | `/v1/queues/{id}` | Get queue details |
| DELETE | `/v1/queues/{id}` | Delete a queue |
| POST | `/v1/queues/{id}` | Write messages to a queue |
| GET | `/v1/queues/{id}` | Read messages from a queue (long-poll) |
| GET | `/v1/clients` | List active streaming clients |

## Common Error Codes

| Code | Description |
|------|-------------|
| 400 | Bad request — malformed JSON or invalid parameters |
| 404 | Not found — the requested queue does not exist |
| 409 | Conflict — queue operation conflict |
| 500 | Internal server error |
