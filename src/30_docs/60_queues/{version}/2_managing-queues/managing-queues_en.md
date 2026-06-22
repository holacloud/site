# Managing Queues

Tailon exposes in-memory queues.

Use `POST /v1/queues` to create, `GET /v1/queues` to list queue IDs as `[]string`, and `GET /v1/queues/{queue_id}` to retrieve `name`, `len`, `reads`, and `writes`.

The delete handler is not implemented.
