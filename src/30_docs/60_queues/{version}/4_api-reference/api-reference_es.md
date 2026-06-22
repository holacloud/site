# Referencia API de Tailon

URL base: `https://api.hola.cloud`

Tailon es un servicio simple de colas en memoria.

| Método | Ruta | Descripción |
|--------|------|-------------|
| POST | `/v1/queues` | Crear cola; `201` con cuerpo vacío |
| GET | `/v1/queues` | Listar IDs de colas como `[]string` |
| GET | `/v1/queues/{queue_id}` | Recuperar `name`, `len`, `reads`, `writes` |
| POST | `/v1/queues/{queue_id}:write` | Escribir mensajes |
| GET | `/v1/queues/{queue_id}:read` | Leer mensajes como NDJSON |

El handler de delete no está implementado.
