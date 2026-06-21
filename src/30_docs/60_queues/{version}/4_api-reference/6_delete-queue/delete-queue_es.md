
# DELETE /v1/queues/{id}

Elimina una cola y todos sus mensajes pendientes de forma permanente.

## Autenticación

Pública — no requiere autenticación.

## Parámetros de Ruta

| Parámetro | Descripción |
|-----------|-------------|
| `id` | El identificador único de la cola |

## Solicitud

```bash
curl -X DELETE "https://api.hola.cloud/v1/queues/a1b2c3d4-e5f6-7890-abcd-ef1234567890"
```

```http
DELETE /v1/queues/a1b2c3d4-e5f6-7890-abcd-ef1234567890 HTTP/1.1
Host: api.hola.cloud
```

## Respuesta

HTTP `204 Sin Contenido`.

## Códigos de Error

| Código | Descripción |
|--------|-------------|
| 404 | Cola no encontrada |
