
# GET /v1/queues/{id}

Lee mensajes de una cola usando long-poll. La solicitud se bloquea hasta que al menos un mensaje esté disponible o se alcance el tiempo de espera del servidor.

## Autenticación

Pública — no requiere autenticación.

## Parámetros de Ruta

| Parámetro | Descripción |
|-----------|-------------|
| `id` | El identificador único de la cola |

## Encabezados de Solicitud

| Encabezado | Descripción |
|------------|-------------|
| `Limit` | Número máximo de mensajes a devolver (por defecto varía según el servidor) |

## Solicitud

```bash
curl -X GET "https://api.hola.cloud/v1/queues/a1b2c3d4-e5f6-7890-abcd-ef1234567890" \
  -H "Limit: 10"
```

```http
GET /v1/queues/a1b2c3d4-e5f6-7890-abcd-ef1234567890 HTTP/1.1
Host: api.hola.cloud
Limit: 10
```

## Respuesta

Los mensajes se devuelven en orden FIFO y se eliminan de la cola al ser entregados.

```json
[
  {
    "id": "msg-001",
    "type": "pedido.creado",
    "order_id": "ord-1234",
    "amount": 49.99,
    "timestamp": "2025-06-21T12:00:01Z"
  },
  {
    "id": "msg-002",
    "type": "vista_pagina",
    "url": "/home",
    "user": "u1",
    "timestamp": "2025-06-21T12:00:02Z"
  }
]
```

Si la cola está vacía y no llegan mensajes antes del tiempo de espera, el servidor devuelve HTTP `204 Sin Contenido`.

## Códigos de Error

| Código | Descripción |
|--------|-------------|
| 204 | No hay mensajes disponibles (tiempo de espera agotado) |
| 400 | Valor de `Limit` inválido |
| 404 | Cola no encontrada |
