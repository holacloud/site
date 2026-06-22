# Leer Mensajes

Lee mensajes con `GET /v1/queues/{queue_id}:read`. La respuesta es NDJSON.

```bash
curl "https://api.hola.cloud/v1/queues/queue_1:read"
```

```json
{"message":"hello"}
```
