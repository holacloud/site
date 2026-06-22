# Colas Tailon

Tailon es un servicio simple de colas en memoria. Las colas viven en la memoria del proceso y exponen una API HTTP pequeña para crear, listar, inspeccionar, escribir y leer.

## Resumen de API

| Método | Ruta | Descripción |
|--------|------|-------------|
| POST | `/v1/queues` | Crear una cola; devuelve `201` con cuerpo vacío |
| GET | `/v1/queues` | Listar IDs de colas como `[]string` |
| GET | `/v1/queues/{queue_id}` | Recuperar `name`, `len`, `reads` y `writes` |
| POST | `/v1/queues/{queue_id}:write` | Escribir mensajes |
| GET | `/v1/queues/{queue_id}:read` | Leer mensajes como NDJSON |

El handler de delete no está implementado.
