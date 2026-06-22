# Crear Cola

```bash
curl -X POST "https://api.hola.cloud/v1/queues" \
  -H "Content-Type: application/json" \
  -d '{"name":"my-queue"}'
```

Respuesta: `201 Created` con cuerpo vacío.
