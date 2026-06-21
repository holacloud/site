
# GET /v1/queues/{id}

Obtiene los detalles de una cola específica por su ID.

## Autenticación

Pública — no requiere autenticación.

## Parámetros de Ruta

| Parámetro | Descripción |
|-----------|-------------|
| `id` | El identificador único de la cola |

## Solicitud

```bash
curl "https://api.hola.cloud/v1/queues/a1b2c3d4-e5f6-7890-abcd-ef1234567890"
```

```http
GET /v1/queues/a1b2c3d4-e5f6-7890-abcd-ef1234567890 HTTP/1.1
Host: api.hola.cloud
```

## Respuesta

```json
{
  "id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
  "name": "my-queue",
  "length": 5,
  "reads": 10,
  "writes": 15
}
```

## Códigos de Error

| Código | Descripción |
|--------|-------------|
| 404 | Cola no encontrada |
