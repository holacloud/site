
# List Clients

Lista todos los clientes de streaming actualmente conectados en todas las colas.

## Autenticación

Pública — no requiere autenticación.

## Solicitud

```bash
curl "https://api.hola.cloud/v1/clients"
```

```http
GET /v1/clients HTTP/1.1
Host: api.hola.cloud
```

## Respuesta

```json
[
  {
    "id": "client-001",
    "queue_id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
    "connected_at": "2025-06-21T12:00:00Z",
    "user_agent": "curl/7.88.1"
  },
  {
    "id": "client-002",
    "queue_id": "b2c3d4e5-f6a7-8901-bcde-f12345678901",
    "connected_at": "2025-06-21T12:00:05Z",
    "user_agent": "TailonClient/1.0"
  }
]
```

## Códigos de Error

| Código | Descripción |
|--------|-------------|
| 500 | Error interno del servidor |
