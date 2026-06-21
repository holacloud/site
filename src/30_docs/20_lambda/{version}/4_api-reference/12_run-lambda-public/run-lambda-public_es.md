
# POST /run/{id}

Invoca una función lambda de forma pública. No requiere autenticación.

Este endpoint está diseñado para webhooks y endpoints de API públicos. El ID de la lambda se puede obtener desde las páginas de administración de lambdas.

## Parámetros de Ruta

| Parámetro | Tipo | Descripción |
|-----------|------|-------------|
| id | uuid | Identificador único de la lambda a ejecutar |

## Cuerpo de la Solicitud

Cualquier carga útil JSON que desee pasar a la función lambda.

## Solicitud HTTP

```http
POST /run/f1a2b3c4-d5e6-7890-abcd-ef0123456789 HTTP/1.1
Host: api.hola.cloud
Content-Type: application/json

{
  "event": "payment_received",
  "amount": 49.99,
  "currency": "USD"
}
```

## Ejemplo

```bash
curl -X POST "https://api.hola.cloud/run/f1a2b3c4-d5e6-7890-abcd-ef0123456789" \
  -H "Content-Type: application/json" \
  -d '{
    "event": "payment_received",
    "amount": 49.99,
    "currency": "USD"
  }'
```

## Respuesta

```json
{
  "status": 200,
  "body": {
    "processed": true,
    "event": "payment_received"
  }
}
```

## Códigos de Error

| Código | Descripción |
|--------|-------------|
| 400 | Cuerpo de solicitud inválido |
| 404 | Lambda no encontrada o inactiva |
| 500 | Error de ejecución de la lambda |
