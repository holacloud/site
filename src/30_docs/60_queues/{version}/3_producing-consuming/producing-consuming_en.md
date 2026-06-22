# Producing and Consuming

Write with `POST /v1/queues/{queue_id}:write` and read with `GET /v1/queues/{queue_id}:read`.

```bash
curl -X POST "https://api.hola.cloud/v1/queues/queue_1:write" \
  --data-binary $'{"message":"hello"}\n'
```

```bash
curl "https://api.hola.cloud/v1/queues/queue_1:read"
```

Read responses are NDJSON.
