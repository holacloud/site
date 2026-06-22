# Tailon Queues

Tailon is a simple in-memory queue service. Queues live in process memory and expose a small HTTP API for create, list, inspect, write, and read.

## API Overview

| Method | Path | Description |
|--------|------|-------------|
| POST | `/v1/queues` | Create a queue; returns `201` with an empty body |
| GET | `/v1/queues` | List queue IDs as `[]string` |
| GET | `/v1/queues/{queue_id}` | Retrieve `name`, `len`, `reads`, and `writes` |
| POST | `/v1/queues/{queue_id}:write` | Write messages |
| GET | `/v1/queues/{queue_id}:read` | Read messages as NDJSON |

The delete handler is not implemented.
