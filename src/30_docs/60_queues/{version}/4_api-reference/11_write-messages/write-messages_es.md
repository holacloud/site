# Escribir Mensajes

Escribe mensajes con `POST /v1/queues/{queue_id}:write`.

```bash
curl -X POST "https://api.hola.cloud/v1/queues/queue_1:write" \
  -H "Content-Type: application/x-ndjson" \
  --data-binary $'{"message":"hello"}\n'
```
