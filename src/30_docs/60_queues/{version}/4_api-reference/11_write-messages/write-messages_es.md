
# Write Messages

Escribe mensajes en una cola. Acepta un solo objeto JSON o un flujo de objetos JSON (NDJSON).

## Autenticación

Pública — no requiere autenticación.

## Parámetros de Ruta

| Parámetro | Descripción |
|-----------|-------------|
| `id` | El identificador único de la cola |

## Cuerpo de la Solicitud

Acepta `application/json` (mensaje único) o `application/x-ndjson` (múltiples mensajes, uno por línea).

Cada mensaje es un objeto JSON arbitrario. El servidor agrega automáticamente los campos `id` y `timestamp` al encolar.

```json
{
  "type": "pedido.creado",
  "order_id": "ord-1234",
  "amount": 49.99
}
```

O múltiples mensajes como NDJSON:

```
{"type": "pedido.creado", "order_id": "ord-1234", "amount": 49.99}
{"type": "vista_pagina", "url": "/home", "user": "u1"}
```

## Solicitud

```bash
# Mensaje único
curl -X POST "https://api.hola.cloud/v1/queues/a1b2c3d4-e5f6-7890-abcd-ef1234567890" \
  -H "Content-Type: application/json" \
  -d '{
    "type": "pedido.creado",
    "order_id": "ord-1234",
    "amount": 49.99
  }'

# Múltiples mensajes (NDJSON)
curl -X POST "https://api.hola.cloud/v1/queues/a1b2c3d4-e5f6-7890-abcd-ef1234567890" \
  -H "Content-Type: application/x-ndjson" \
  -d '{"type": "pedido.creado", "order_id": "ord-1234", "amount": 49.99}
{"type": "vista_pagina", "url": "/home", "user": "u1"}'
```

```http
POST /v1/queues/a1b2c3d4-e5f6-7890-abcd-ef1234567890 HTTP/1.1
Host: api.hola.cloud
Content-Type: application/x-ndjson

{"type": "pedido.creado", "order_id": "ord-1234", "amount": 49.99}
{"type": "vista_pagina", "url": "/home", "user": "u1"}
```

## Respuesta

```json
{
  "written": 2
}
```

## Códigos de Error

| Código | Descripción |
|--------|-------------|
| 400 | JSON malformado o tipo de contenido inválido |
| 404 | Cola no encontrada |
| 413 | Carga útil demasiado grande |
