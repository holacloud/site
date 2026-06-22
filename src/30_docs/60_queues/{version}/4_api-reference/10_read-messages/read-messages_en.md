# Read Messages

Read messages with `GET /v1/queues/{queue_id}:read`. The response is NDJSON.

```bash
curl "https://api.hola.cloud/v1/queues/queue_1:read"
```

```json
{"message":"hello"}
```
