# Primeros Pasos

Las colas Tailon son colas simples en memoria.

```bash
curl -X POST "https://api.hola.cloud/v1/queues" \
  -H "Content-Type: application/json" \
  -d '{"name":"my-queue"}'
```

Crear devuelve `201` con cuerpo vacío.

```bash
curl -X POST "https://api.hola.cloud/v1/queues/queue_1:write" \
  --data-binary $'{"message":"hello"}\n'
```

```bash
curl "https://api.hola.cloud/v1/queues/queue_1:read"
```

Las lecturas devuelven NDJSON.
