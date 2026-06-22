# Getting Started

Tailon queues are simple in-memory queues.

```bash
curl -X POST "https://api.hola.cloud/v1/queues" \
  -H "Content-Type: application/json" \
  -d '{"name":"my-queue"}'
```

Create returns `201` with an empty body.

```bash
curl -X POST "https://api.hola.cloud/v1/queues/queue_1:write" \
  --data-binary $'{"message":"hello"}\n'
```

```bash
curl "https://api.hola.cloud/v1/queues/queue_1:read"
```

Reads return NDJSON.
