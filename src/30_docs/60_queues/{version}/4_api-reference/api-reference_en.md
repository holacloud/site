# Tailon API Reference

Base URL: `https://api.hola.cloud`

Tailon is a simple in-memory queue service.

| Method | Path | Description |
|--------|------|-------------|
| POST | `/v1/queues` | Create queue; `201` empty body |
| GET | `/v1/queues` | List queue IDs as `[]string` |
| GET | `/v1/queues/{queue_id}` | Retrieve `name`, `len`, `reads`, `writes` |
| POST | `/v1/queues/{queue_id}:write` | Write messages |
| GET | `/v1/queues/{queue_id}:read` | Read messages as NDJSON |

The delete handler is not implemented.
