# Administrando Colas

Este documento cubre el ciclo de vida de las colas: creación con configuración opcional, listado, inspección, monitoreo de clientes y eliminación.

## Crear una Cola

Crea una cola enviando una solicitud POST:

```bash
curl -X POST "https://api.hola.cloud/v1/queues" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "mi-cola"
  }'
```

Respuesta esperada:

```json
{
  "id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
  "name": "mi-cola",
  "length": 0,
  "reads": 0,
  "writes": 0
}
```

La solicitud HTTP correspondiente se ve así:

```http
POST /v1/queues HTTP/1.1
Host: api.hola.cloud
Content-Type: application/json

{
  "name": "mi-cola"
}
```

## Listar Todas las Colas

Recupera una lista de todas las colas de tu cuenta:

```bash
curl "https://api.hola.cloud/v1/queues"
```

Respuesta esperada:

```json
[
  {
    "id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
    "name": "mi-cola",
    "length": 5,
    "reads": 10,
    "writes": 15
  },
  {
    "id": "b2c3d4e5-f6a7-8901-bcde-f12345678901",
    "name": "cola-eventos",
    "length": 0,
    "reads": 42,
    "writes": 42
  }
]
```

## Obtener Detalles de una Cola

Inspecciona una cola por su ID para ver la longitud actual, lecturas totales y escrituras totales:

```bash
curl "https://api.hola.cloud/v1/queues/a1b2c3d4-e5f6-7890-abcd-ef1234567890"
```

```http
GET /v1/queues/a1b2c3d4-e5f6-7890-abcd-ef1234567890 HTTP/1.1
Host: api.hola.cloud
```

Respuesta esperada:

```json
{
  "id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
  "name": "mi-cola",
  "length": 5,
  "reads": 10,
  "writes": 15
}
```

## Monitorear Clientes Activos

Lista todos los clientes de streaming actualmente conectados:

```bash
curl "https://api.hola.cloud/v1/clients"
```

```http
GET /v1/clients HTTP/1.1
Host: api.hola.cloud
```

Respuesta esperada:

```json
[
  {
    "id": "cliente-001",
    "queue_id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
    "connected_at": "2025-06-21T12:00:00Z",
    "user_agent": "curl/7.88.1"
  }
]
```

## Eliminar una Cola

Elimina permanentemente una cola y todos sus mensajes pendientes:

```bash
curl -X DELETE "https://api.hola.cloud/v1/queues/a1b2c3d4-e5f6-7890-abcd-ef1234567890"
```

```http
DELETE /v1/queues/a1b2c3d4-e5f6-7890-abcd-ef1234567890 HTTP/1.1
Host: api.hola.cloud
```

Respuesta esperada: HTTP `204 No Content`.

Una vez eliminada, el ID de la cola ya no es válido. Los lectores en curso recibirán un error en su siguiente solicitud.
