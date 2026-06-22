# Produciendo y Consumiendo

Escribe con `POST /v1/queues/{queue_id}:write` y lee con `GET /v1/queues/{queue_id}:read`.

```bash
curl -X POST "https://api.hola.cloud/v1/queues/queue_1:write" \
  --data-binary $'{"message":"hello"}\n'
```

```bash
curl "https://api.hola.cloud/v1/queues/queue_1:read"
```

Las respuestas de lectura son NDJSON.
